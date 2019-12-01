package session

import (
	"context"
	"testing"

	"github.com/bxcodec/faker"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ory/viper"

	"github.com/ory/kratos/driver/configuration"
	"github.com/ory/kratos/identity"
	"github.com/ory/kratos/x"
)

type PersistenceProvider interface {
	SessionPersister() Persister
}

type Persister interface {
	// Get retrieves a session from the store.
	GetSession(ctx context.Context, sid uuid.UUID) (*Session, error)

	// Create adds a session to the store.
	CreateSession(ctx context.Context, s *Session) error

	// Delete removes a session from the store
	DeleteSession(ctx context.Context, sid uuid.UUID) error
}

func TestPersister(p interface {
	Persister
	identity.Pool
}) func(t *testing.T) {
	return func(t *testing.T) {
		viper.Set(configuration.ViperKeyDefaultIdentityTraitsSchemaURL, "file://./stub/identity.schema.json")

		t.Run("case=not found", func(t *testing.T) {
			_, err := p.GetSession(context.Background(), x.NewUUID())
			require.Error(t, err)
		})

		t.Run("case=create session", func(t *testing.T) {
			var expected Session
			require.NoError(t, faker.FakeData(&expected))
			require.NoError(t, p.CreateIdentity(context.Background(), expected.Identity))

			now := expected.ID
			t.Logf("now: %s", now)
			assert.NotEqual(t, uuid.Nil, expected.ID)
			require.NoError(t, p.CreateSession(context.Background(), &expected))
			assert.NotEqual(t, uuid.Nil, expected.ID)
			later := expected.ID
			t.Logf("later: %s", later)

			actual, err := p.GetSession(context.Background(), expected.ID)
			require.NoError(t, err)
			assert.Equal(t, expected.Identity.ID, actual.Identity.ID)
			assert.Equal(t, expected.ID, actual.ID)
			assert.EqualValues(t, expected.ExpiresAt.Unix(), actual.ExpiresAt.Unix())
			assert.Equal(t, expected.AuthenticatedAt.Unix(), actual.AuthenticatedAt.Unix())
			assert.Equal(t, expected.IssuedAt.Unix(), actual.IssuedAt.Unix())
		})

		t.Run("case=delete session", func(t *testing.T) {
			var expected Session
			require.NoError(t, faker.FakeData(&expected))
			require.NoError(t, p.CreateIdentity(context.Background(), expected.Identity))
			require.NoError(t, p.CreateSession(context.Background(), &expected))

			require.NoError(t, p.DeleteSession(context.Background(), expected.ID))
			_, err := p.GetSession(context.Background(), expected.ID)
			require.Error(t, err)
		})
	}
}
