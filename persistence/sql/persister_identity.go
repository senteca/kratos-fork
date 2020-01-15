package sql

import (
	"context"
	"database/sql"
	"encoding/json"
	"strings"

	"github.com/ory/kratos/driver/configuration"

	"github.com/gobuffalo/pop"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"

	"github.com/ory/herodot"
	"github.com/ory/x/errorsx"
	"github.com/ory/x/sqlcon"

	"github.com/ory/kratos/identity"
	"github.com/ory/kratos/schema"
)

var _ identity.Pool = new(Persister)

func (p *Persister) FindByCredentialsIdentifier(ctx context.Context, ct identity.CredentialsType, match string) (*identity.Identity, *identity.Credentials, error) {
	var cts []identity.CredentialsTypeTable
	if err := p.c.All(&cts); err != nil {
		return nil, nil, sqlcon.HandleError(err)
	}

	var find struct {
		IdentityID uuid.UUID `db:"identity_id"`
	}

	if err := p.c.RawQuery(`SELECT
    ic.identity_id
FROM identity_credentials ic
         INNER JOIN identity_credential_types ict on ic.identity_credential_type_id = ict.id
         INNER JOIN identity_credential_identifiers ici on ic.id = ici.identity_credential_id
WHERE ici.identifier = ?
  AND ict.name = ?`, match, ct).First(&find); err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, nil, herodot.ErrNotFound.WithTrace(err).WithReasonf(`No identity matching credentials identifier "%s" could be found.`, match)
		}

		return nil, nil, sqlcon.HandleError(err)
	}

	i, err := p.GetIdentityConfidential(ctx, find.IdentityID)
	if err != nil {
		return nil, nil, err
	}

	creds, ok := i.GetCredentials(ct)
	if !ok {
		return nil, nil, errors.WithStack(herodot.ErrInternalServerError.WithReasonf("The SQL adapter failed to return the appropriate credentials_type \"%s\". This is a bug in the code.", ct))
	}

	return i.CopyWithoutCredentials(), creds, nil
}

func findOrCreateIdentityCredentialsType(_ context.Context, tx *pop.Connection, ct identity.CredentialsType) (*identity.CredentialsTypeTable, error) {
	var m identity.CredentialsTypeTable
	if err := tx.Where("name = ?", ct).First(&m); err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			m.Name = ct
			if err := sqlcon.HandleError(tx.Create(&m)); err != nil {
				return nil, err
			}
			return &m, nil
		}
		return nil, sqlcon.HandleError(err)
	}

	return &m, nil
}

func createIdentityCredentials(ctx context.Context, tx *pop.Connection, i *identity.Identity) error {
	for k, cred := range i.Credentials {
		cred.IdentityID = i.ID
		if len(cred.Config) == 0 {
			cred.Config = json.RawMessage("{}")
		}

		ct, err := findOrCreateIdentityCredentialsType(ctx, tx, cred.Type)
		if err != nil {
			return err
		}

		cred.CredentialTypeID = ct.ID
		if err := tx.Create(&cred); err != nil {
			return err
		}

		for _, ids := range cred.Identifiers {
			// Force case-insensitivity for email addresses
			if strings.Contains(ids, "@") && cred.Type == identity.CredentialsTypePassword {
				ids = strings.ToLower(ids)
			}

			if len(ids) == 0 {
				return errors.WithStack(herodot.ErrInternalServerError.WithReasonf("Unable to create identity credentials with missing or empty identifier."))
			}

			ci := &identity.CredentialIdentifier{
				Identifier:            ids,
				IdentityCredentialsID: cred.ID,
			}
			if err := tx.Create(ci); err != nil {
				return err
			}
		}

		i.Credentials[k] = cred
	}

	return nil
}

func (p *Persister) CreateIdentity(ctx context.Context, i *identity.Identity) error {
	if i.TraitsSchemaID == "" {
		i.TraitsSchemaID = configuration.DefaultIdentityTraitsSchemaID
	}

	if len(i.Traits) == 0 {
		i.Traits = identity.Traits("{}")
	}

	if err := p.injectTraitsSchemaURL(i); err != nil {
		return err
	}

	if err := p.validateIdentity(i); err != nil {
		return err
	}

	return sqlcon.HandleError(p.c.Transaction(func(tx *pop.Connection) error {
		if err := tx.Create(i); err != nil {
			return err
		}

		return createIdentityCredentials(ctx, tx, i)
	}))
}

func (p *Persister) ListIdentities(ctx context.Context, limit, offset int) ([]identity.Identity, error) {
	is := make([]identity.Identity, 0)
	if err := sqlcon.HandleError(p.c.RawQuery("SELECT * FROM identities LIMIT ? OFFSET ?", limit, offset).All(&is)); err != nil {
		return nil, err
	}

	for i := range is {
		if err := p.injectTraitsSchemaURL(&(is[i])); err != nil {
			return nil, err
		}
	}

	return is, nil
}

func (p *Persister) UpdateIdentityConfidential(ctx context.Context, i *identity.Identity) error {
	if err := p.validateIdentity(i); err != nil {
		return err
	}

	return sqlcon.HandleError(p.c.Transaction(func(tx *pop.Connection) error {
		if err := tx.RawQuery(`DELETE FROM identity_credentials WHERE identity_id = ?`, i.ID).Exec(); err != nil {
			return err
		}

		if err := tx.Update(i); err != nil {
			return err
		}

		return createIdentityCredentials(ctx, tx, i)
	}))
}

func (p *Persister) UpdateIdentity(ctx context.Context, i *identity.Identity) error {
	if err := p.validateIdentity(i); err != nil {
		return err
	}

	fs, err := p.GetIdentityConfidential(ctx, i.ID)
	if err != nil {
		return err
	}

	// If credential identifiers have changed we need to block this action UNLESS
	// the identity has been authenticated in that request:
	//
	// - https://security.stackexchange.com/questions/24291/why-do-we-ask-for-a-users-existing-password-when-changing-their-password
	if !i.CredentialsEqual(fs.Credentials) {

		// !! WARNING !!
		//
		// This will leak the credential options which may include the hashed password. Do not use seriously:
		//
		//	p.r.Logger().
		//	 	WithField("original_credentials", fmt.Sprintf("%+v", fs.Credentials)).
		//	 	WithField("updated_credentials", fmt.Sprintf("%+v", i.Credentials)).
		//	 	Trace("Credentials changed unexpectedly in UpdateIdentity.")

		return errors.WithStack(
			herodot.ErrInternalServerError.
				WithReasonf(`A field was modified that updates one or more credentials-related settings. This action was blocked because a unprivileged DBAL method was used to execute the update. This is either a configuration issue, or a bug.`))
	}

	return sqlcon.HandleError(p.c.Update(i))
}

func (p *Persister) DeleteIdentity(_ context.Context, id uuid.UUID) error {
	count, err := p.c.RawQuery("DELETE FROM identities WHERE id = ?", id).ExecWithCount()
	if err != nil {
		return sqlcon.HandleError(err)
	}
	if count == 0 {
		return sqlcon.ErrNoRows
	}
	return nil
}

func (p *Persister) GetIdentity(_ context.Context, id uuid.UUID) (*identity.Identity, error) {
	var i identity.Identity
	if err := p.c.Find(&i, id); err != nil {
		return nil, sqlcon.HandleError(err)
	}
	i.Credentials = nil
	if err := p.injectTraitsSchemaURL(&i); err != nil {
		return nil, err
	}

	return &i, nil
}

func (p *Persister) GetIdentityConfidential(_ context.Context, id uuid.UUID) (*identity.Identity, error) {
	var i identity.Identity
	if err := p.c.Eager().Find(&i, id); err != nil {
		return nil, sqlcon.HandleError(err)
	}

	var cts []identity.CredentialsTypeTable
	if err := p.c.All(&cts); err != nil {
		return nil, sqlcon.HandleError(err)
	}

	i.Credentials = map[identity.CredentialsType]identity.Credentials{}
	for _, creds := range i.CredentialsCollection {
		var cs identity.CredentialIdentifierCollection
		if err := p.c.Where("identity_credential_id = ?", creds.ID).All(&cs); err != nil {
			return nil, sqlcon.HandleError(err)
		}

		creds.CredentialIdentifierCollection = nil
		creds.Identifiers = make([]string, len(cs))
		for k := range cs {
			for _, ct := range cts {
				if ct.ID == creds.CredentialTypeID {
					creds.Type = ct.Name
				}
			}
			creds.Identifiers[k] = cs[k].Identifier
		}
		i.Credentials[creds.Type] = creds
	}
	i.CredentialsCollection = nil
	if err := p.injectTraitsSchemaURL(&i); err != nil {
		return nil, err
	}

	return &i, nil
}

func (p *Persister) validateIdentity(i *identity.Identity) error {
	if err := p.r.IdentityValidator().Validate(i); err != nil {
		if _, ok := errorsx.Cause(err).(schema.ResultErrors); ok {
			return errors.WithStack(herodot.ErrBadRequest.WithReasonf("%s", err))
		}
		return err
	}

	return nil
}

func (p *Persister) injectTraitsSchemaURL(i *identity.Identity) error {
	s, err := p.r.IdentityTraitsSchemas().GetByID(i.TraitsSchemaID)
	if err != nil {
		return errors.WithStack(
			herodot.ErrInternalServerError.
				WithReasonf(
					`The JSON Schema "%s" for this identity's traits could not be found.`, i.TraitsSchemaID))
	}
	i.TraitsSchemaURL = s.SchemaURL(p.cf.SelfPublicURL()).String()
	return nil
}
