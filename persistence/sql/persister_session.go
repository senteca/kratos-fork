package sql

import (
	"context"

	"github.com/gofrs/uuid"

	"github.com/ory/x/sqlcon"

	"github.com/ory/kratos/session"
)

var _ session.Persister = new(Persister)

func (p *Persister) GetSession(ctx context.Context, sid uuid.UUID) (*session.Session, error) {
	var s session.Session
	if err := p.GetConnection(ctx).Eager("Identity").Find(&s, sid); err != nil {
		return nil, sqlcon.HandleError(err)
	}
	return &s, nil
}

func (p *Persister) CreateSession(ctx context.Context, s *session.Session) error {
	return p.GetConnection(ctx).Create(s) // This must not be eager or identities will be created / updated
}

func (p *Persister) DeleteSession(ctx context.Context, sid uuid.UUID) error {
	return p.GetConnection(ctx).Destroy(&session.Session{ID: sid}) // This must not be eager or identities will be created / updated
}

func (p *Persister) DeleteSessionsFor(ctx context.Context, identityID uuid.UUID) error {
	if err := p.GetConnection(ctx).RawQuery("DELETE FROM sessions WHERE identity_id = ?", identityID).Exec(); err != nil {
		return sqlcon.HandleError(err)
	}
	return nil
}

func (p *Persister) GetSessionFromToken(ctx context.Context, token string) (*session.Session, error) {
	var s session.Session
	if err := p.GetConnection(ctx).Eager("Identity").Where("token = ?", token).First(&s); err != nil {
		return nil, sqlcon.HandleError(err)
	}
	return &s, nil
}

func (p *Persister) DeleteSessionFromToken(ctx context.Context, token string) error {
	if err := p.GetConnection(ctx).RawQuery("DELETE FROM sessions WHERE token = ?", token).Exec(); err != nil {
		return sqlcon.HandleError(err)
	}
	return nil
}
