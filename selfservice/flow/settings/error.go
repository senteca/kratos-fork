package settings

import (
	"net/http"
	"net/url"

	"github.com/pkg/errors"

	"github.com/ory/x/sqlxx"

	"github.com/ory/herodot"
	"github.com/ory/x/urlx"

	"github.com/ory/kratos/driver/configuration"
	"github.com/ory/kratos/selfservice/errorx"
	"github.com/ory/kratos/selfservice/flow/login"
	"github.com/ory/kratos/x"
)

var (
	ErrRequestExpired = herodot.ErrBadRequest.
				WithError("settings request expired").
				WithReasonf(`The settings request has expired. Please restart the flow.`)
	ErrHookAbortRequest             = errors.New("aborted settings hook execution")
	ErrRequestNeedsReAuthentication = herodot.ErrForbidden.WithReasonf("The login session is too old and thus not allowed to update these fields. Please re-authenticate.")
)

type (
	errorHandlerDependencies interface {
		errorx.ManagementProvider
		x.WriterProvider
		x.LoggingProvider

		RequestPersistenceProvider
	}

	ErrorHandlerProvider interface{ SettingsRequestErrorHandler() *ErrorHandler }

	ErrorHandler struct {
		d errorHandlerDependencies
		c configuration.Provider
	}
)

func NewErrorHandler(d errorHandlerDependencies, c configuration.Provider) *ErrorHandler {
	return &ErrorHandler{
		d: d,
		c: c,
	}
}

func (s *ErrorHandler) reauthenticate(
	w http.ResponseWriter,
	r *http.Request,
	rr *Request) {
	if err := s.d.SettingsRequestPersister().UpdateSettingsRequest(r.Context(), rr); err != nil {
		s.d.SelfServiceErrorManager().Forward(r.Context(), w, r, err)
		return
	}

	returnTo := urlx.CopyWithQuery(urlx.AppendPaths(s.c.SelfPublicURL(), r.URL.Path), r.URL.Query())
	s.c.SelfPublicURL()
	u := urlx.AppendPaths(
		urlx.CopyWithQuery(s.c.SelfPublicURL(), url.Values{
			"prompt":    {"login"},
			"return_to": {returnTo.String()},
		}), login.BrowserLoginPath)

	http.Redirect(w, r, u.String(), http.StatusFound)
}

func (s *ErrorHandler) HandleSettingsError(
	w http.ResponseWriter,
	r *http.Request,
	rr *Request,
	err error,
	method string,
) {
	s.d.Audit().
		WithError(err).
		WithRequest(r).
		WithField("settings_request", rr).
		Info("Encountered self-service settings error.")

	if rr == nil {
		s.d.SelfServiceErrorManager().Forward(r.Context(), w, r, err)
		return
	} else if x.IsJSONRequest(r) {
		s.d.Writer().WriteError(w, r, err)
		return
	}

	if errors.Is(err, ErrRequestNeedsReAuthentication) {
		s.reauthenticate(w, r, rr)
		return
	}

	if _, ok := rr.Methods[method]; !ok {
		s.d.SelfServiceErrorManager().Forward(r.Context(), w, r, errors.WithStack(herodot.ErrInternalServerError.WithReasonf("Expected settings method %s to exist.", method)))
		return
	}

	rr.Active = sqlxx.NullString(method)
	if err := rr.Methods[method].Config.ParseError(err); err != nil {
		s.d.SelfServiceErrorManager().Forward(r.Context(), w, r, err)
		return
	}

	if err := s.d.SettingsRequestPersister().UpdateSettingsRequest(r.Context(), rr); err != nil {
		s.d.SelfServiceErrorManager().Forward(r.Context(), w, r, err)
		return
	}

	http.Redirect(w, r,
		urlx.CopyWithQuery(s.c.SettingsURL(), url.Values{"request": {rr.ID.String()}}).String(),
		http.StatusFound,
	)
}
