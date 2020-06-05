package driver

import "github.com/ory/kratos/selfservice/flow/settings"

func (m *RegistryDefault) PostSettingsPrePersistHooks(settingsType string) (b []settings.PostHookPrePersistExecutor) {
	for _, v := range m.getHooks(settingsType, m.c.SelfServiceSettingsAfterHooks(settingsType)) {
		if hook, ok := v.(settings.PostHookPrePersistExecutor); ok {
			b = append(b, hook)
		}
	}
	return
}

func (m *RegistryDefault) PostSettingsPostPersistHooks(settingsType string) (b []settings.PostHookPostPersistExecutor) {
	for _, v := range m.getHooks(settingsType, m.c.SelfServiceSettingsAfterHooks(settingsType)) {
		if hook, ok := v.(settings.PostHookPostPersistExecutor); ok {
			b = append(b, hook)
		}
	}
	return
}

func (m *RegistryDefault) SettingsHookExecutor() *settings.HookExecutor {
	if m.selfserviceSettingsExecutor == nil {
		m.selfserviceSettingsExecutor = settings.NewHookExecutor(m, m.c)
	}
	return m.selfserviceSettingsExecutor
}

func (m *RegistryDefault) SettingsHandler() *settings.Handler {
	if m.selfserviceSettingsHandler == nil {
		m.selfserviceSettingsHandler = settings.NewHandler(m, m.c)
	}
	return m.selfserviceSettingsHandler
}

func (m *RegistryDefault) SettingsRequestErrorHandler() *settings.ErrorHandler {
	if m.selfserviceSettingsErrorHandler == nil {
		m.selfserviceSettingsErrorHandler = settings.NewErrorHandler(m, m.c)
	}
	return m.selfserviceSettingsErrorHandler
}

func (m *RegistryDefault) SettingsStrategies() settings.Strategies {
	if len(m.profileStrategies) == 0 {
		for _, strategy := range m.selfServiceStrategies() {
			if s, ok := strategy.(settings.Strategy); ok {
				m.profileStrategies = append(m.profileStrategies, s)
			}
		}
	}
	return m.profileStrategies
}
