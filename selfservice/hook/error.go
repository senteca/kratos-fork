package hook

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/tidwall/gjson"

	"github.com/ory/kratos/identity"
	"github.com/ory/kratos/selfservice/flow/login"
	"github.com/ory/kratos/selfservice/flow/registration"
	"github.com/ory/kratos/selfservice/flow/settings"
	"github.com/ory/kratos/session"
)

var (
	_ registration.PostHookPrePersistExecutor  = new(Error)
	_ registration.PostHookPostPersistExecutor = new(Error)
	_ registration.PreHookExecutor             = new(Error)

	_ login.PreHookExecutor  = new(Error)
	_ login.PostHookExecutor = new(Error)

	_ settings.PostHookPostPersistExecutor = new(Error)
	_ settings.PostHookPrePersistExecutor  = new(Error)
)

type Error struct {
	Config json.RawMessage
}

func (e Error) err(path string, abort error) error {
	switch gjson.GetBytes(e.Config, path).String() {
	case "err":
		return errors.New("err")
	case "abort":
		return abort
	}
	return nil
}

func (e Error) ExecuteSettingsPrePersistHook(w http.ResponseWriter, r *http.Request, a *settings.Request, s *identity.Identity) error {
	return e.err("ExecuteSettingsPrePersistHook", settings.ErrHookAbortRequest)
}

func (e Error) ExecuteSettingsPostPersistHook(w http.ResponseWriter, r *http.Request, a *settings.Request, s *identity.Identity) error {
	return e.err("ExecuteSettingsPostPersistHook", settings.ErrHookAbortRequest)
}

func (e Error) ExecuteLoginPostHook(w http.ResponseWriter, r *http.Request, a *login.Request, s *session.Session) error {
	return e.err("ExecuteLoginPostHook", login.ErrHookAbortRequest)
}

func (e Error) ExecuteLoginPreHook(w http.ResponseWriter, r *http.Request, a *login.Request) error {
	return e.err("ExecuteLoginPreHook", login.ErrHookAbortRequest)
}

func (e Error) ExecuteRegistrationPreHook(w http.ResponseWriter, r *http.Request, a *registration.Request) error {
	return e.err("ExecuteRegistrationPreHook", registration.ErrHookAbortRequest)
}

func (e Error) ExecutePostRegistrationPostPersistHook(w http.ResponseWriter, r *http.Request, a *registration.Request, s *session.Session) error {
	return e.err("ExecutePostRegistrationPostPersistHook", registration.ErrHookAbortRequest)
}

func (e Error) ExecutePostRegistrationPrePersistHook(w http.ResponseWriter, r *http.Request, a *registration.Request, i *identity.Identity) error {
	return e.err("ExecutePostRegistrationPrePersistHook", registration.ErrHookAbortRequest)
}
