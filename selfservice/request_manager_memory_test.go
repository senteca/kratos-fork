package selfservice_test

import (
	"context"
	"testing"

	"github.com/bxcodec/faker"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ory/hive/identity"
	. "github.com/ory/hive/selfservice"
	"github.com/ory/hive/selfservice/password"
)

func TestRequestManagerMemory(t *testing.T) {
	t.Run("suite=sign-in", func(t *testing.T) {
		m := NewRequestManagerMemory()

		var r LoginRequest
		require.NoError(t, faker.FakeData(&r))
		r.Methods = map[identity.CredentialsType]*DefaultRequestMethod{
			identity.CredentialsTypePassword: {},
		}

		require.NoError(t, m.CreateLoginRequest(context.Background(), &r))
		g, err := m.GetLoginRequest(context.Background(), r.ID)
		require.NoError(t, err)
		assert.EqualValues(t, r, *g)

		require.NoError(t, m.UpdateLoginRequest(context.Background(), r.ID, identity.CredentialsTypePassword, &password.RequestMethodConfig{Action: "foo"}))

		g, err = m.GetLoginRequest(context.Background(), r.ID)
		require.NoError(t, err)
		assert.EqualValues(t, r, *g)
		assert.EqualValues(t, identity.CredentialsTypePassword, r.Active)
		assert.EqualValues(t, "foo", r.Methods[identity.CredentialsTypePassword].Config.(*password.RequestMethodConfig).Action)
	})

	t.Run("suite=sign-up", func(t *testing.T) {
		m := NewRequestManagerMemory()

		var r RegistrationRequest
		require.NoError(t, faker.FakeData(&r))
		r.Methods = map[identity.CredentialsType]*DefaultRequestMethod{
			identity.CredentialsTypePassword: {},
		}

		require.NoError(t, m.CreateRegistrationRequest(context.Background(), &r))
		g, err := m.GetRegistrationRequest(context.Background(), r.ID)
		require.NoError(t, err)
		assert.EqualValues(t, r, *g)

		require.NoError(t, m.UpdateRegistrationRequest(context.Background(), r.ID, identity.CredentialsTypePassword, &password.RequestMethodConfig{Action: "foo"}))

		g, err = m.GetRegistrationRequest(context.Background(), r.ID)
		require.NoError(t, err)
		assert.EqualValues(t, r, *g)
		assert.EqualValues(t, identity.CredentialsTypePassword, r.Active)
		assert.EqualValues(t, "foo", r.Methods[identity.CredentialsTypePassword].Config.(*password.RequestMethodConfig).Action)
	})
}
