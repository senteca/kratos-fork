package sql

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/gobuffalo/pop/v6"
	"github.com/gofrs/uuid"

	"github.com/ory/kratos/corp"
	"github.com/ory/kratos/identity"
	"github.com/ory/kratos/selfservice/flow/recovery"
	"github.com/ory/kratos/selfservice/strategy/code"
	"github.com/ory/kratos/selfservice/strategy/link"
	"github.com/ory/x/sqlcon"
)

var _ recovery.FlowPersister = new(Persister)
var _ link.RecoveryTokenPersister = new(Persister)

func (p *Persister) CreateRecoveryFlow(ctx context.Context, r *recovery.Flow) error {
	ctx, span := p.r.Tracer(ctx).Tracer().Start(ctx, "persistence.sql.CreateRecoveryFlow")
	defer span.End()

	r.NID = corp.ContextualizeNID(ctx, p.nid)
	return p.GetConnection(ctx).Create(r)
}

func (p *Persister) GetRecoveryFlow(ctx context.Context, id uuid.UUID) (*recovery.Flow, error) {
	ctx, span := p.r.Tracer(ctx).Tracer().Start(ctx, "persistence.sql.GetRecoveryFlow")
	defer span.End()

	var r recovery.Flow
	if err := p.GetConnection(ctx).Where("id = ? AND nid = ?", id, corp.ContextualizeNID(ctx, p.nid)).First(&r); err != nil {
		return nil, sqlcon.HandleError(err)
	}

	return &r, nil
}

func (p *Persister) UpdateRecoveryFlow(ctx context.Context, r *recovery.Flow) error {
	ctx, span := p.r.Tracer(ctx).Tracer().Start(ctx, "persistence.sql.UpdateRecoveryFlow")
	defer span.End()

	cp := *r
	cp.NID = corp.ContextualizeNID(ctx, p.nid)
	return p.update(ctx, cp)
}

func (p *Persister) CreateRecoveryToken(ctx context.Context, token *link.RecoveryToken) error {
	ctx, span := p.r.Tracer(ctx).Tracer().Start(ctx, "persistence.sql.CreateRecoveryToken")
	defer span.End()

	t := token.Token
	token.Token = p.hmacValue(ctx, t)
	token.NID = corp.ContextualizeNID(ctx, p.nid)

	// This should not create the request eagerly because otherwise we might accidentally create an address that isn't
	// supposed to be in the database.
	if err := p.GetConnection(ctx).Create(token); err != nil {
		return err
	}

	token.Token = t
	return nil
}

func (p *Persister) UseRecoveryToken(ctx context.Context, token string) (*link.RecoveryToken, error) {
	ctx, span := p.r.Tracer(ctx).Tracer().Start(ctx, "persistence.sql.UseRecoveryToken")
	defer span.End()

	var rt link.RecoveryToken

	nid := corp.ContextualizeNID(ctx, p.nid)
	if err := sqlcon.HandleError(p.Transaction(ctx, func(ctx context.Context, tx *pop.Connection) (err error) {
		for _, secret := range p.r.Config(ctx).SecretsSession() {
			if err = tx.Where("token = ? AND nid = ? AND NOT used", p.hmacValueWithSecret(ctx, token, secret), nid).First(&rt); err != nil {
				if !errors.Is(sqlcon.HandleError(err), sqlcon.ErrNoRows) {
					return err
				}
			} else {
				break
			}
		}
		if err != nil {
			return err
		}

		var ra identity.RecoveryAddress
		if err := tx.Where("id = ? AND nid = ?", rt.RecoveryAddressID, nid).First(&ra); err != nil {
			if !errors.Is(sqlcon.HandleError(err), sqlcon.ErrNoRows) {
				return err
			}
		}
		rt.RecoveryAddress = &ra

		/* #nosec G201 TableName is static */
		return tx.RawQuery(fmt.Sprintf("UPDATE %s SET used=true, used_at=? WHERE id=? AND nid = ?", rt.TableName(ctx)), time.Now().UTC(), rt.ID, nid).Exec()
	})); err != nil {
		return nil, err
	}

	return &rt, nil
}

func (p *Persister) DeleteRecoveryToken(ctx context.Context, token string) error {
	ctx, span := p.r.Tracer(ctx).Tracer().Start(ctx, "persistence.sql.DeleteRecoveryToken")
	defer span.End()

	/* #nosec G201 TableName is static */
	return p.GetConnection(ctx).RawQuery(fmt.Sprintf("DELETE FROM %s WHERE token=? AND nid = ?", new(link.RecoveryToken).TableName(ctx)), token, corp.ContextualizeNID(ctx, p.nid)).Exec()
}

func (p *Persister) DeleteExpiredRecoveryFlows(ctx context.Context, expiresAt time.Time, limit int) error {
	// #nosec G201
	err := p.GetConnection(ctx).RawQuery(fmt.Sprintf(
		"DELETE FROM %s WHERE id in (SELECT id FROM (SELECT id FROM %s c WHERE expires_at <= ? and nid = ? ORDER BY expires_at ASC LIMIT %d ) AS s )",
		new(recovery.Flow).TableName(ctx),
		new(recovery.Flow).TableName(ctx),
		limit,
	),
		expiresAt,
		corp.ContextualizeNID(ctx, p.nid),
	).Exec()
	if err != nil {
		return sqlcon.HandleError(err)
	}
	return nil
}

func (p *Persister) CreateRecoveryCode(ctx context.Context, recoveryCode *code.RecoveryCode) error {
	ctx, span := p.r.Tracer(ctx).Tracer().Start(ctx, "persistence.sql.CreateRecoveryCode")
	defer span.End()

	code := recoveryCode.Code
	recoveryCode.Code = p.hmacValue(ctx, code)
	recoveryCode.NID = corp.ContextualizeNID(ctx, p.nid)

	// TODO: This should not create the request eagerly because otherwise we might accidentally create an address that isn't
	// supposed to be in the database.
	if err := p.GetConnection(ctx).Create(recoveryCode); err != nil {
		return err
	}

	recoveryCode.Code = code
	return nil
}

func (p *Persister) UseRecoveryCode(ctx context.Context, codeVal string) (*code.RecoveryCode, error) {
	ctx, span := p.r.Tracer(ctx).Tracer().Start(ctx, "persistence.sql.UseRecoveryCode")
	defer span.End()

	var recoveryCode code.RecoveryCode

	nid := corp.ContextualizeNID(ctx, p.nid)
	if err := sqlcon.HandleError(p.Transaction(ctx, func(ctx context.Context, tx *pop.Connection) (err error) {
		for _, secret := range p.r.Config(ctx).SecretsSession() {
			if err = tx.Where("token = ? AND nid = ? AND NOT used", p.hmacValueWithSecret(ctx, codeStr, secret), nid).First(&recoveryCode); err != nil {
				if !errors.Is(sqlcon.HandleError(err), sqlcon.ErrNoRows) {
					return err
				}
			} else {
				break
			}
		}
		if err != nil {
			return err
		}

		var ra identity.RecoveryAddress
		if err := tx.Where("id = ? AND nid = ?", recoveryCode.RecoveryAddressID, nid).First(&ra); err != nil {
			if !errors.Is(sqlcon.HandleError(err), sqlcon.ErrNoRows) {
				return err
			}
		}
		recoveryCode.RecoveryAddress = &ra

		/* #nosec G201 TableName is static */
		return tx.RawQuery(fmt.Sprintf("UPDATE %s SET used=true, used_at=? WHERE id=? AND nid = ?", recoveryCode.TableName(ctx)), time.Now().UTC(), recoveryCode.ID, nid).Exec()
	})); err != nil {
		return nil, err
	}

	return &recoveryCode, nil
}

func (p *Persister) DeleteRecoveryCode(ctx context.Context, codeStr string) error {
	ctx, span := p.r.Tracer(ctx).Tracer().Start(ctx, "persistence.sql.DeleteRecoveryCode")
	defer span.End()

	/* #nosec G201 TableName is static */
	return p.GetConnection(ctx).RawQuery(fmt.Sprintf("DELETE FROM %s WHERE token=? AND nid = ?", new(code.RecoveryCode).TableName(ctx)), codeStr, corp.ContextualizeNID(ctx, p.nid)).Exec()
}
