package driver

import (
	"github.com/ory/kratos/identity"
	"github.com/ory/kratos/selfservice/flow/registration"
)

func (m *RegistryDefault) PostRegistrationHooks(credentialsType identity.CredentialsType) []registration.PostHookExecutor {
	a := m.hooksPost(credentialsType, m.c.SelfServiceRegistrationAfterHooks(string(credentialsType)))
	b := make([]registration.PostHookExecutor, len(a))
	for k, v := range a {
		b[k] = v
	}
	return b
}

func (m *RegistryDefault) PreRegistrationHooks() []registration.PreHookExecutor {
	return []registration.PreHookExecutor{}
}
func (m *RegistryDefault) RegistrationExecutor() *registration.HookExecutor {
	if m.selfserviceRegistrationExecutor == nil {
		m.selfserviceRegistrationExecutor = registration.NewHookExecutor(m, m.c)
	}
	return m.selfserviceRegistrationExecutor
}

func (m *RegistryDefault) RegistrationHookExecutor() *registration.HookExecutor {
	if m.selfserviceRegistrationExecutor == nil {
		m.selfserviceRegistrationExecutor = registration.NewHookExecutor(m, m.c)
	}
	return m.selfserviceRegistrationExecutor
}

func (m *RegistryDefault) RegistrationErrorHandler() *registration.ErrorHandler {
	if m.seflserviceRegistrationErrorHandler == nil {
		m.seflserviceRegistrationErrorHandler = registration.NewErrorHandler(m, m.c)
	}
	return m.seflserviceRegistrationErrorHandler
}

func (m *RegistryDefault) RegistrationHandler() *registration.Handler {
	if m.selfserviceRegistrationHandler == nil {
		m.selfserviceRegistrationHandler = registration.NewHandler(m, m.c)
	}

	return m.selfserviceRegistrationHandler
}

func (m *RegistryDefault) RegistrationRequestErrorHandler() *registration.ErrorHandler {
	if m.selfserviceRegistrationRequestErrorHandler == nil {
		m.selfserviceRegistrationRequestErrorHandler = registration.NewErrorHandler(m, m.c)
	}

	return m.selfserviceRegistrationRequestErrorHandler
}
