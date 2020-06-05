package configuration

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"

	"github.com/pkg/errors"

	"github.com/ory/x/tracing"
)

type HasherArgon2Config struct {
	Memory      uint32
	Iterations  uint32
	Parallelism uint8
	SaltLength  uint32
	KeyLength   uint32
}

type SelfServiceHook struct {
	Name   string          `json:"hook"`
	Config json.RawMessage `json:"config"`
}

type SelfServiceStrategy struct {
	Enabled bool            `json:"enabled"`
	Config  json.RawMessage `json:"config"`
}

type SchemaConfig struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}

type SchemaConfigs []SchemaConfig

func (s SchemaConfigs) FindSchemaByID(id string) (*SchemaConfig, error) {
	for _, sc := range s {
		if sc.ID == id {
			return &sc, nil
		}
	}

	return nil, errors.Errorf("could not find schema with id \"%s\"", id)
}

const DefaultIdentityTraitsSchemaID = "default"

type Provider interface {
	AdminListenOn() string
	PublicListenOn() string
	DSN() string

	SessionSecrets() [][]byte

	SelfPublicURL() *url.URL
	SelfAdminURL() *url.URL

	DefaultReturnToURL() *url.URL

	SettingsURL() *url.URL
	LoginURL() *url.URL
	VerificationURL() *url.URL
	ErrorURL() *url.URL
	MultiFactorURL() *url.URL
	RecoveryURL() *url.URL

	SessionLifespan() time.Duration
	SelfServiceSettingsRequestLifespan() time.Duration
	SelfServiceVerificationRequestLifespan() time.Duration
	SelfServiceLoginRequestLifespan() time.Duration
	SelfServiceRegistrationRequestLifespan() time.Duration
	SelfServiceRecoveryRequestLifespan() time.Duration

	SelfServiceStrategy(strategy string) *SelfServiceStrategy
	SelfServiceLoginBeforeHooks() []SelfServiceHook
	SelfServiceRegistrationBeforeHooks() []SelfServiceHook
	SelfServiceLoginAfterHooks(strategy string) []SelfServiceHook
	SelfServiceLoginReturnTo(strategy string) *url.URL
	SelfServiceRegistrationAfterHooks(strategy string) []SelfServiceHook
	SelfServiceRegistrationReturnTo(strategy string) *url.URL
	SelfServiceSettingsAfterHooks(strategy string) []SelfServiceHook
	SelfServiceSettingsReturnTo(strategy string, defaultReturnTo *url.URL) *url.URL
	SelfServiceLogoutRedirectURL() *url.URL

	SelfServicePrivilegedSessionMaxAge() time.Duration
	SelfServiceVerificationReturnTo() *url.URL

	CourierSMTPFrom() string
	CourierSMTPURL() *url.URL
	CourierTemplatesRoot() string

	DefaultIdentityTraitsSchemaURL() *url.URL
	IdentityTraitsSchemas() SchemaConfigs

	WhitelistedReturnToDomains() []url.URL

	RegisterURL() *url.URL

	HasherArgon2() *HasherArgon2Config

	TracingServiceName() string
	TracingProvider() string
	TracingJaegerConfig() *tracing.JaegerConfig

	IsInsecureDevMode() bool

	SessionSameSiteMode() http.SameSite
}
