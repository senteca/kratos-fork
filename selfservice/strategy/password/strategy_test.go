package password_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ory/viper"

	"github.com/ory/kratos/driver/configuration"
	"github.com/ory/kratos/identity"
	"github.com/ory/kratos/internal"
	"github.com/ory/kratos/selfservice/errorx"
	"github.com/ory/kratos/selfservice/strategy/password"
	"github.com/ory/kratos/session"
	"github.com/ory/kratos/x"
)

func newErrTs(t *testing.T, reg interface {
	errorx.PersistenceProvider
	x.WriterProvider
}) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		e, err := reg.SelfServiceErrorPersister().Read(r.Context(), x.ParseUUID(r.URL.Query().Get("error")))
		require.NoError(t, err)
		reg.Writer().Write(w, r, e.Errors)
	}))
}

func newReturnTs(t *testing.T, reg interface {
	session.ManagementProvider
	x.WriterProvider
}) *httptest.Server {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sess, err := reg.SessionManager().FetchFromRequest(r.Context(), r)
		require.NoError(t, err)
		reg.Writer().Write(w, r, sess)
	}))
	t.Cleanup(ts.Close)
	viper.Set(configuration.ViperKeyURLsDefaultReturnTo, ts.URL+"/return-ts")
	return ts
}

func TestCountActiveCredentials(t *testing.T) {
	conf, reg := internal.NewFastRegistryWithMocks(t)
	strategy := password.NewStrategy(reg, conf)

	hash, err := reg.Hasher().Generate([]byte("a password"))
	require.NoError(t, err)

	for k, tc := range []struct {
		in       identity.CredentialsCollection
		expected int
	}{
		{
			in: identity.CredentialsCollection{{
				Type:   strategy.ID(),
				Config: json.RawMessage{},
			}},
			expected: 0,
		},
		{
			in: identity.CredentialsCollection{{
				Type:   strategy.ID(),
				Config: json.RawMessage(`{"hashed_password": "` + string(hash) + `"}`),
			}},
			expected: 0,
		},
		{
			in: identity.CredentialsCollection{{
				Type:        strategy.ID(),
				Identifiers: []string{""},
				Config:      json.RawMessage(`{"hashed_password": "` + string(hash) + `"}`),
			}},
			expected: 0,
		},
		{
			in: identity.CredentialsCollection{{
				Type:        strategy.ID(),
				Identifiers: []string{"foo"},
				Config:      json.RawMessage(`{"hashed_password": "` + string(hash) + `"}`),
			}},
			expected: 1,
		},
		{
			in: identity.CredentialsCollection{{
				Type:   strategy.ID(),
				Config: json.RawMessage(`{"hashed_password": "asdf"}`),
			}},
			expected: 0,
		},
		{
			in: identity.CredentialsCollection{{
				Type:   strategy.ID(),
				Config: json.RawMessage(`{}`),
			}},
			expected: 0,
		},
		{
			in:       identity.CredentialsCollection{{}, {}},
			expected: 0,
		},
	} {
		t.Run(fmt.Sprintf("case=%d", k), func(t *testing.T) {
			cc := map[identity.CredentialsType]identity.Credentials{}
			for _, c := range tc.in {
				cc[c.Type] = c
			}

			actual, err := strategy.CountActiveCredentials(cc)
			require.NoError(t, err)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
