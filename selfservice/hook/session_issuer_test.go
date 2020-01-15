package hook_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ory/viper"

	"github.com/ory/kratos/driver/configuration"
	"github.com/ory/kratos/identity"
	"github.com/ory/kratos/internal"
	"github.com/ory/kratos/selfservice/hook"
	"github.com/ory/kratos/session"
	"github.com/ory/kratos/x"
)

func TestSessionIssuer(t *testing.T) {
	_, reg := internal.NewRegistryDefault(t)
	viper.Set(configuration.ViperKeyURLsSelfPublic, "http://localhost/")
	viper.Set(configuration.ViperKeyDefaultIdentityTraitsSchemaURL, "file://./stub/stub.schema.json")

	var r http.Request
	h := hook.NewSessionIssuer(reg)

	t.Run("method=sign-in", func(t *testing.T) {
		w := httptest.NewRecorder()
		sid := x.NewUUID()

		i := identity.NewIdentity(configuration.DefaultIdentityTraitsSchemaID)
		require.NoError(t, reg.IdentityPool().CreateIdentity(context.Background(), i))
		require.NoError(t, h.ExecuteLoginPostHook(w, &r, nil, &session.Session{ID: sid, Identity: i}))

		got, err := reg.SessionPersister().GetSession(context.Background(), sid)
		require.NoError(t, err)
		assert.Equal(t, sid, got.ID)
	})

	t.Run("method=sign-up", func(t *testing.T) {
		w := httptest.NewRecorder()
		sid := x.NewUUID()

		i := identity.NewIdentity(configuration.DefaultIdentityTraitsSchemaID)
		require.NoError(t, reg.IdentityPool().CreateIdentity(context.Background(), i))
		require.NoError(t, h.ExecuteRegistrationPostHook(w, &r, nil, &session.Session{ID: sid, Identity: i}))

		got, err := reg.SessionPersister().GetSession(context.Background(), sid)
		require.NoError(t, err)
		assert.Equal(t, sid, got.ID)
	})
}
