package driver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/ory/kratos/driver/configuration"
	"github.com/ory/kratos/identity"
	"github.com/ory/kratos/selfservice/hook"
)

func (m *RegistryDefault) hooksPost(credentialsType identity.CredentialsType, configs []configuration.SelfServiceHook) postHooks {
	var i postHooks

	for _, h := range configs {
		switch h.Run {
		case hook.KeySessionIssuer:
			i = append(
				i,
				hook.NewSessionIssuer(m),
			)
		case hook.KeyRedirector:
			var rc struct {
				R string `json:"default_redirect_url"`
				A bool   `json:"allow_user_defined_redirect"`
			}

			if err := json.NewDecoder(bytes.NewBuffer(h.Config)).Decode(&rc); err != nil {
				m.l.WithError(err).
					WithField("type", credentialsType).
					WithField("hook", h.Run).
					WithField("config", fmt.Sprintf("%s", h.Config)).
					Errorf("The after hook is misconfigured.")
				continue
			}

			rcr, err := url.ParseRequestURI(rc.R)
			if err != nil {
				m.l.WithError(err).
					WithField("type", credentialsType).
					WithField("hook", h.Run).
					WithField("config", fmt.Sprintf("%s", h.Config)).
					Errorf("The after hook is misconfigured.")
				continue
			}

			i = append(
				i,
				hook.NewRedirector(
					func() *url.URL {
						return rcr
					},
					m.c.WhitelistedReturnToDomains,
					func() bool {
						return rc.A
					},
				),
			)
		default:
			m.l.
				WithField("type", credentialsType).
				WithField("hook", h.Run).
				Errorf("A unknown post login hook was requested and can therefore not be used.")
		}
	}

	return i
}
