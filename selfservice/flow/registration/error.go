package registration

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/pkg/errors"

	"github.com/ory/herodot"
	"github.com/ory/x/urlx"

	"github.com/ory/kratos/driver/configuration"
	"github.com/ory/kratos/identity"
	"github.com/ory/kratos/selfservice/errorx"
	"github.com/ory/kratos/x"
)

var (
	ErrHookAbortRequest = errors.New("abort hook")

	ErrRequestExpired = herodot.ErrBadRequest.
				WithError("registration request expired").
				WithReasonf(`The registration request has expired. Please restart the flow.`)
)

type (
	errorHandlerDependencies interface {
		errorx.ManagementProvider
		x.WriterProvider
		x.LoggingProvider

		RequestPersistenceProvider
	}

	ErrorHandlerProvider interface{ RegistrationRequestErrorHandler() *ErrorHandler }

	ErrorHandler struct {
		d  errorHandlerDependencies
		c  configuration.Provider
		bd *x.BodyDecoder
	}
)

func NewErrorHandler(d errorHandlerDependencies, c configuration.Provider) *ErrorHandler {
	return &ErrorHandler{
		d: d,
		c: c,
	}
}

func (s *ErrorHandler) HandleRegistrationError(
	w http.ResponseWriter,
	r *http.Request,
	ct identity.CredentialsType,
	rr *Request,
	err error,
) {
	s.d.Logger().WithError(err).
		WithField("details", fmt.Sprintf("%+v", err)).
		WithField("credentials_type", ct).
		WithField("login_request", rr).
		Warn("Encountered login error.")

	if rr == nil {
		s.d.ErrorManager().ForwardError(r.Context(), w, r, err)
		return
	} else if x.IsJSONRequest(r) {
		s.d.Writer().WriteError(w, r, err)
		return
	}

	method, ok := rr.Methods[ct]
	if !ok {
		s.d.Writer().WriteError(w, r, errors.WithStack(herodot.ErrInternalServerError.WithErrorf(`Expected method "%s" to exist in request. This is a bug in the code and should be reported on GitHub.`, ct)))
		return
	}

	if err := method.Config.ParseError(err); err != nil {
		s.d.ErrorManager().ForwardError(r.Context(), w, r, err)
		return
	}

	if err := s.d.RegistrationRequestPersister().UpdateRegistrationRequest(r.Context(), rr.ID, ct, method); err != nil {
		s.d.ErrorManager().ForwardError(r.Context(), w, r, err)
		return
	}

	http.Redirect(w, r,
		urlx.CopyWithQuery(s.c.RegisterURL(), url.Values{"request": {rr.ID}}).String(),
		http.StatusFound,
	)
}
