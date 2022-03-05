package webauthn_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/ory/kratos/driver/config"
	"github.com/ory/kratos/selfservice/strategy/webauthn"
	"github.com/ory/kratos/session"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ory/kratos/identity"
	"github.com/ory/kratos/internal"
)

func TestCompletedAuthenticationMethod(t *testing.T) {
	conf, reg := internal.NewFastRegistryWithMocks(t)
	strategy := webauthn.NewStrategy(reg)

	assert.Equal(t, session.AuthenticationMethod{
		Method: strategy.ID(),
		AAL:    identity.AuthenticatorAssuranceLevel2,
	}, strategy.CompletedAuthenticationMethod(context.Background()))

	conf.MustSet(config.ViperKeyWebAuthnPasswordless, true)
	assert.Equal(t, session.AuthenticationMethod{
		Method: strategy.ID(),
		AAL:    identity.AuthenticatorAssuranceLevel1,
	}, strategy.CompletedAuthenticationMethod(context.Background()))
}

func TestCountActiveFirstFactorCredentials(t *testing.T) {
	_, reg := internal.NewFastRegistryWithMocks(t)
	strategy := webauthn.NewStrategy(reg)

	for k, tc := range []struct {
		in       identity.CredentialsCollection
		expected int
	}{
		{
			in: identity.CredentialsCollection{{
				Type:   strategy.ID(),
				Config: []byte{},
			}},
			expected: 0,
		},
		{
			in: identity.CredentialsCollection{{
				Type:   strategy.ID(),
				Config: []byte(`{"credentials": []}`),
			}},
			expected: 0,
		},
		{
			in: identity.CredentialsCollection{{
				Type:        strategy.ID(),
				Identifiers: []string{"foo"},
				Config:      []byte(`{"credentials": [{}]}`),
			}},
			expected: 0,
		},
		{
			in: identity.CredentialsCollection{{
				Type:        strategy.ID(),
				Identifiers: []string{"foo"},
				Config:      []byte(`{"credentials": [{"is_passwordless": true}]}`),
			}},
			expected: 1,
		},
		{
			in: identity.CredentialsCollection{{
				Type:        strategy.ID(),
				Identifiers: []string{"foo"},
				Config:      []byte(`{"credentials": [{"is_passwordless": true}, {"is_passwordless": true}]}`),
			}},
			expected: 2,
		},
		{
			in: identity.CredentialsCollection{{
				Type:   strategy.ID(),
				Config: []byte(`{}`),
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

			actual, err := strategy.CountActiveFirstFactorCredentials(cc)
			require.NoError(t, err)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
