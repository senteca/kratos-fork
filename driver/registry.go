package driver

import (
	"github.com/go-errors/errors"
	"github.com/gorilla/sessions"
	"github.com/justinas/nosurf"
	"github.com/sirupsen/logrus"

	"github.com/ory/herodot"

	"github.com/ory/hive/selfservice"

	"github.com/ory/x/dbal"

	"github.com/ory/hive/driver/configuration"
	"github.com/ory/hive/errorx"
	"github.com/ory/hive/identity"
	password2 "github.com/ory/hive/selfservice/password"
	"github.com/ory/hive/session"
)

type Registry interface {
	dbal.Driver

	Init() error

	WithConfig(c configuration.Provider) Registry
	WithLogger(l logrus.FieldLogger) Registry

	BuildVersion() string
	BuildDate() string
	BuildHash() string
	WithBuildInfo(version, hash, date string) Registry

	Logger() logrus.FieldLogger
	Writer() herodot.Writer

	ErrorManager() errorx.Manager
	ErrorHandler() *errorx.Handler

	IdentityHandler() *identity.Handler
	IdentityPool() identity.Pool

	PasswordHasher() password2.Hasher
	PasswordValidator() password2.Validator

	SessionHandler() *session.Handler
	SessionManager() session.Manager

	StrategyHandler() *selfservice.StrategyHandler
	SelfServiceStrategies() []selfservice.Strategy

	CookieManager() sessions.Store

	WithCSRFHandler(c *nosurf.CSRFHandler)
	CSRFHandler() *nosurf.CSRFHandler

	AuthHookRegistrationPreExecutors() []selfservice.HookRegistrationPreExecutor
	AuthHookLoginPreExecutors() []selfservice.HookLoginPreExecutor
	LoginExecutor() *selfservice.LoginExecutor
	PostLoginHooks(credentialsType identity.CredentialsType) []selfservice.HookLoginPostExecutor
	RegistrationExecutor() *selfservice.RegistrationExecutor
	PostRegistrationHooks(credentialsType identity.CredentialsType) []selfservice.HookRegistrationPostExecutor
	IdentityValidator() *identity.Validator
	SelfServiceRequestErrorHandler() *selfservice.ErrorHandler
	LoginRequestManager() selfservice.LoginRequestManager
	RegistrationRequestManager() selfservice.RegistrationRequestManager
}

func NewRegistry(c configuration.Provider) (Registry, error) {
	driver, err := dbal.GetDriverFor(c.DSN())
	if err != nil {
		return nil, err
	}

	registry, ok := driver.(Registry)
	if !ok {
		return nil, errors.Errorf("driver of type %T does not implement interface Registry", driver)
	}

	return registry, nil
}
