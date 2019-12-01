package oidc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	"github.com/justinas/nosurf"
	"github.com/pkg/errors"

	"github.com/ory/x/jsonx"

	"github.com/ory/gojsonschema"
	"github.com/ory/herodot"
	"github.com/ory/x/stringsx"
	"github.com/ory/x/urlx"

	"github.com/ory/kratos/selfservice/flow/login"
	"github.com/ory/kratos/selfservice/flow/registration"
	"github.com/ory/kratos/selfservice/form"

	"github.com/ory/kratos/driver/configuration"
	"github.com/ory/kratos/identity"
	"github.com/ory/kratos/schema"
	"github.com/ory/kratos/selfservice/errorx"
	"github.com/ory/kratos/session"
	"github.com/ory/kratos/x"
)

const (
	BasePath     = "/auth/browser/methods/oidc"
	AuthPath     = BasePath + "/auth/:request"
	CallbackPath = BasePath + "/callback/:provider"

	registrationFormPayloadSchema = `{
  "$id": "https://schemas.ory.sh/kratos/selfservice/oidc/registration/config.schema.json",
  "$schema": "http://json-schema.org/draft-07/schema#",
  "type": "object",
  "properties": {
    "traits": {}
  }
}`
)

var _ login.Strategy = new(Strategy)
var _ registration.Strategy = new(Strategy)

type dependencies interface {
	errorx.ManagementProvider

	x.LoggingProvider
	x.CookieProvider

	identity.ValidationProvider
	identity.PoolProvider

	session.ManagementProvider
	session.HandlerProvider

	login.HookExecutorProvider
	login.RequestPersistenceProvider
	login.HooksProvider
	login.StrategyProvider
	login.HandlerProvider
	login.ErrorHandlerProvider

	registration.HookExecutorProvider
	registration.RequestPersistenceProvider
	registration.HooksProvider
	registration.StrategyProvider
	registration.HandlerProvider
	registration.ErrorHandlerProvider
}

// Strategy implements selfservice.LoginStrategy, selfservice.RegistrationStrategy. It supports both login
// and registration via OpenID Providers.
type Strategy struct {
	c configuration.Provider
	d dependencies

	dec       *x.BodyDecoder
	validator *schema.Validator
	cg        form.CSRFGenerator
}

func (s *Strategy) WithTokenGenerator(g form.CSRFGenerator) {
	s.cg = g
}

func (s *Strategy) RegisterLoginRoutes(r *x.RouterPublic) {
	s.setRoutes(r)
}

func (s *Strategy) RegisterRegistrationRoutes(r *x.RouterPublic) {
	s.setRoutes(r)
}

func (s *Strategy) setRoutes(r *x.RouterPublic) {
	if handle, _, _ := r.Lookup("GET", CallbackPath); handle == nil {
		r.GET(CallbackPath, s.d.SessionHandler().IsNotAuthenticated(s.handleCallback, session.RedirectOnAuthenticated(s.c)))
	}

	if handle, _, _ := r.Lookup("POST", AuthPath); handle == nil {
		r.POST(AuthPath, s.d.SessionHandler().IsNotAuthenticated(s.handleAuth, session.RedirectOnAuthenticated(s.c)))
	}

	if handle, _, _ := r.Lookup("GET", AuthPath); handle == nil {
		r.GET(AuthPath, s.d.SessionHandler().IsNotAuthenticated(s.handleAuth, session.RedirectOnAuthenticated(s.c)))
	}
}

func NewStrategy(
	d dependencies,
	c configuration.Provider,
) *Strategy {
	return &Strategy{
		c:         c,
		d:         d,
		cg:        nosurf.Token,
		validator: schema.NewValidator(),
		dec:       x.NewBodyDecoder(),
	}
}

func (s *Strategy) ID() identity.CredentialsType {
	return identity.CredentialsTypeOIDC
}

func (s *Strategy) RegistrationStrategyID() identity.CredentialsType {
	return s.ID()
}

func (s *Strategy) LoginStrategyID() identity.CredentialsType {
	return s.ID()
}

func (s *Strategy) handleAuth(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var rid = ps.ByName("request")

	if err := r.ParseForm(); err != nil {
		s.handleError(w, r, rid, nil, errors.WithStack(herodot.ErrBadRequest.WithDebug(err.Error()).WithReasonf("Unable to parse HTTP form request: %s", err.Error())))
		return
	}

	var (
		pid = r.Form.Get("provider") // this can come from both url query and post body
	)

	if pid == "" {
		s.handleError(w, r, rid, nil, errors.WithStack(herodot.ErrBadRequest.WithReasonf(`The HTTP request did not contain the required "provider" form field`)))
		return
	}

	provider, err := s.provider(pid)
	if err != nil {
		s.handleError(w, r, rid, nil, err)
		return
	}

	config, err := provider.OAuth2(r.Context())
	if err != nil {
		s.handleError(w, r, rid, nil, err)
		return
	}

	if _, err := s.validateRequest(r.Context(), rid); err != nil {
		s.handleError(w, r, rid, nil, err)
		return
	}

	state := uuid.New().String()
	// Any data that is posted to this endpoint will be used to fill out missing data from the oidc provider.
	if err := x.SessionPersistValues(w, r, s.d.CookieManager(), sessionName, map[string]interface{}{
		sessionKeyState:  state,
		sessionRequestID: rid,
		sessionFormState: r.PostForm.Encode(),
	}); err != nil {
		s.handleError(w, r, rid, nil, err)
		return
	}

	http.Redirect(w, r, config.AuthCodeURL(state), http.StatusFound)
}

func (s *Strategy) validateRequest(ctx context.Context, rid string) (request, error) {
	if rid == "" {
		return nil, errors.WithStack(herodot.ErrBadRequest.WithReason("The session cookie contains invalid values and the request could not be executed. Please try again."))
	}

	if ar, err := s.d.RegistrationRequestPersister().GetRegistrationRequest(ctx, rid); err == nil {
		if err := ar.Valid(); err != nil {
			return nil, err
		}
		return ar, nil
	}

	ar, err := s.d.LoginRequestPersister().GetLoginRequest(ctx, rid)
	if err != nil {
		return nil, err
	}

	if err := ar.Valid(); err != nil {
		return nil, err
	}

	return ar, nil
}

func (s *Strategy) validateCallback(r *http.Request) (request, error) {
	var (
		code = r.URL.Query().Get("code")
	)
	if state := r.URL.Query().Get("state"); state == "" {
		return nil, errors.WithStack(herodot.ErrBadRequest.WithReasonf(`Unable to complete OpenID Connect flow because the OpenID Provider did not return the state query parameter.`))
	} else if state != x.SessionGetStringOr(r, s.d.CookieManager(), sessionName, sessionKeyState, "") {
		return nil, errors.WithStack(herodot.ErrBadRequest.WithReasonf(`Unable to complete OpenID Connect flow because the query state parameter does not match the state parameter from the session cookie.`))
	}

	ar, err := s.validateRequest(r.Context(), x.SessionGetStringOr(r, s.d.CookieManager(), sessionName, sessionRequestID, ""))
	if err != nil {
		return nil, err
	}

	if r.URL.Query().Get("error") != "" {
		return ar, errors.WithStack(herodot.ErrBadRequest.WithReasonf(`Unable to complete OpenID Connect flow because the OpenID Provider returned error "%s": %s`, r.URL.Query().Get("error"), r.URL.Query().Get("error_description")))
	}

	if code == "" {
		return ar, errors.WithStack(herodot.ErrBadRequest.WithReasonf(`Unable to complete OpenID Connect flow because the OpenID Provider did not return the code query parameter.`))
	}

	return ar, nil
}

func (s *Strategy) handleCallback(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var (
		code = r.URL.Query().Get("code")
		pid  = ps.ByName("provider")
	)

	ar, err := s.validateCallback(r)
	if err != nil {
		if ar != nil {
			s.handleError(w, r, ar.GetID(), nil, err)
		} else {
			s.handleError(w, r, "", nil, err)
		}
		return
	}

	provider, err := s.provider(pid)
	if err != nil {
		s.handleError(w, r, ar.GetID(), nil, err)
		return
	}

	config, err := provider.OAuth2(context.Background())
	if err != nil {
		s.handleError(w, r, ar.GetID(), nil, err)
		return
	}

	token, err := config.Exchange(r.Context(), code)
	if err != nil {
		s.handleError(w, r, ar.GetID(), nil, err)
		return
	}

	claims, err := provider.Claims(r.Context(), token)
	if err != nil {
		s.handleError(w, r, ar.GetID(), nil, err)
		return
	}

	switch a := ar.(type) {
	case *login.Request:
		s.processLogin(w, r, a, claims, provider)
		return
	case *registration.Request:
		s.processRegistration(w, r, a, claims, provider)
		return
	default:
		panic(fmt.Sprintf("unexpected type: %T", a))
	}
}

func uid(provider, subject string) string {
	return fmt.Sprintf("%s:%s", provider, subject)
}

func (s *Strategy) authURL(request, provider string) string {
	u := urlx.AppendPaths(
		urlx.Copy(s.c.SelfPublicURL()),
		strings.Replace(
			AuthPath, ":request", request, 1,
		),
	)

	if provider != "" {
		return urlx.CopyWithQuery(u, url.Values{"provider": {provider}}).String()
	}

	return u.String()
}

func (s *Strategy) processLogin(w http.ResponseWriter, r *http.Request, a *login.Request, claims *Claims, provider Provider) {
	i, c, err := s.d.IdentityPool().FindByCredentialsIdentifier(r.Context(), identity.CredentialsTypeOIDC, uid(provider.Config().ID, claims.Subject))
	if err != nil {
		if errors.Cause(err).Error() == herodot.ErrNotFound.Error() {
			// If no account was found we're "manually" creating a new registration request and redirecting the browser
			// to that endpoint.

			// That will execute the "pre registration" hook which allows to e.g. disallow this request. The registration
			// ui however will NOT be shown, instead the user is directly redirected to the auth path. That should then
			// do a silent re-request. While this might be a bit excessive from a network perspective it should usually
			// happen without any downsides to user experience as the request has already been authorized and should
			// not need additional consent/login.

			// This is kinda hacky but the only way to ensure seamless login/registration flows when using OIDC.

			s.d.Logger().WithField("provider", provider.Config().ID).WithField("subject", claims.Subject).Debug("Received successful OpenID Connect callback but user is not registered. Re-initializing registration flow now.")
			if err := s.d.RegistrationHandler().NewRegistrationRequest(w, r, func(aa *registration.Request) string {
				return s.authURL(aa.ID, provider.Config().ID)
			}); err != nil {
				s.handleError(w, r, a.GetID(), nil, err)
				return
			}
			return
		}
		s.handleError(w, r, a.GetID(), nil, err)
		return
	}

	var o CredentialsConfig
	if err := json.NewDecoder(bytes.NewBuffer(c.Config)).Decode(&o); err != nil {
		s.handleError(w, r, a.GetID(), nil, errors.WithStack(herodot.ErrInternalServerError.WithReason("The password credentials could not be decoded properly").WithDebug(err.Error())))
		return
	}

	if o.Subject != claims.Subject {
		s.handleError(w, r, a.GetID(), nil, errors.WithStack(herodot.ErrInternalServerError.WithReason("The subjects do not match").WithDebugf("Expected credential subject to match subject from RequestID Token but values are not equal: %s != %s", o.Subject, claims.Subject)))
		return
	} else if o.Provider != provider.Config().ID {
		s.handleError(w, r, a.GetID(), nil, errors.WithStack(herodot.ErrInternalServerError.WithReason("The providers do not match").WithDebugf("Expected credential provider to match provider from path but values are not equal: %s != %s", o.Subject, provider.Config().ID)))
		return
	}

	if err = s.d.LoginHookExecutor().PostLoginHook(w, r, s.d.PostLoginHooks(identity.CredentialsTypeOIDC), a, i); err != nil {
		s.handleError(w, r, a.GetID(), nil, err)
		return
	}
}

func (s *Strategy) processRegistration(w http.ResponseWriter, r *http.Request, a *registration.Request, claims *Claims, provider Provider) {
	if _, _, err := s.d.IdentityPool().FindByCredentialsIdentifier(r.Context(), identity.CredentialsTypeOIDC, uid(provider.Config().ID, claims.Subject)); err == nil {
		// If the identity already exists, we should perform the login flow instead.

		// That will execute the "pre login" hook which allows to e.g. disallow this request. The login
		// ui however will NOT be shown, instead the user is directly redirected to the auth path. That should then
		// do a silent re-request. While this might be a bit excessive from a network perspective it should usually
		// happen without any downsides to user experience as the request has already been authorized and should
		// not need additional consent/login.

		// This is kinda hacky but the only way to ensure seamless login/registration flows when using OIDC.
		s.d.Logger().WithField("provider", provider.Config().ID).WithField("subject", claims.Subject).Debug("Received successful OpenID Connect callback but user is already registered. Re-initializing login flow now.")
		if err := s.d.LoginHandler().NewLoginRequest(w, r, func(aa *login.Request) string {
			return s.authURL(aa.ID, provider.Config().ID)
		}); err != nil {
			s.handleError(w, r, a.GetID(), nil, err)
			return
		}
		return
	}

	i := identity.NewIdentity(s.c.DefaultIdentityTraitsSchemaURL().String())
	extension := NewValidationExtension()
	extension.WithIdentity(i)

	// Validate the claims first (which will also copy the values around based on the schema)
	if err := s.validator.Validate(
		stringsx.Coalesce(
			provider.Config().SchemaURL,
		),
		gojsonschema.NewGoLoader(claims),
		extension,
	); err != nil {
		s.d.Logger().
			WithField("provider", provider.Config().ID).
			WithField("schema_url", provider.Config().SchemaURL).
			WithField("claims", fmt.Sprintf("%+v", claims)).
			Error("Unable to validate claims against provider schema. Your schema should work regardless of these values.")
		// Force a system error because this can not be resolved by the user.
		s.handleError(w, r, a.GetID(), nil, errors.WithStack(herodot.ErrInternalServerError.WithTrace(err).WithReasonf("%s", err)))
		return
	}

	option, err := decoderRegistration(s.c.DefaultIdentityTraitsSchemaURL().String())
	if err != nil {
		s.handleError(w, r, a.GetID(), nil, err)
		return
	}

	traits, err := merge(
		x.SessionGetStringOr(r, s.d.CookieManager(), sessionName, sessionFormState, ""),
		i.Traits, option,
	)
	if err != nil {
		s.handleError(w, r, a.GetID(), nil, err)
		return
	}

	i.Traits = traits

	// Validate the identity itself
	if err := s.d.IdentityValidator().Validate(i); err != nil {
		s.handleError(w, r, a.GetID(), traits, err)
		return
	}

	var b bytes.Buffer
	if err := json.NewEncoder(&b).Encode(&CredentialsConfig{
		Subject:  claims.Subject,
		Provider: provider.Config().ID,
	}); err != nil {
		s.handleError(w, r, a.GetID(), traits, errors.WithStack(herodot.ErrInternalServerError.WithReasonf("Unable to encode password options to JSON: %s", err)))
		return
	}

	i.SetCredentials(s.RegistrationStrategyID(), identity.Credentials{
		ID:          s.RegistrationStrategyID(),
		Identifiers: []string{uid(provider.Config().ID, claims.Subject)},
		Config:      b.Bytes(),
	})

	if err := s.d.RegistrationExecutor().PostRegistrationHook(w, r, s.d.PostRegistrationHooks(identity.CredentialsTypeOIDC), a, i); err != nil {
		s.handleError(w, r, a.GetID(), traits, err)
		return
	}
}

// func (s *Strategy) verifyIdentity(i *identity.Identity, c identity.Credentials, token oidc.IDToken, pid string) error {
// 	var o CredentialsConfig
//
// 	if err := json.NewDecoder(bytes.NewBuffer(c.Config)).Decode(&o); err != nil {
// 		return errors.WithStack(herodot.ErrInternalServerError.WithReason("The password credentials could not be decoded properly").WithDebug(err.Error()))
// 	}
//
// 	if o.Subject != token.Subject {
// 		return errors.WithStack(herodot.ErrInternalServerError.WithReason("The subjects do not match").WithDebugf("Expected credential subject to match subject from RequestID Token but values are not equal: %s != %s", o.Subject, token.Subject))
// 	} else if o.Provider != pid {
// 		return errors.WithStack(herodot.ErrInternalServerError.WithReason("The providers do not match").WithDebugf("Expected credential provider to match provider from path but values are not equal: %s != %s", o.Subject, pid))
// 	}
//
// 	return nil
// }

func (s *Strategy) populateMethod(r *http.Request, request string) (*RequestMethod, error) {
	conf, err := s.Config()
	if err != nil {
		return nil, err
	}

	f := form.NewHTMLForm(s.authURL(request, ""))
	f.SetCSRF(s.cg(r))

	return NewRequestMethodConfig(f).AddProviders(conf.Providers), nil
}

func (s *Strategy) PopulateLoginMethod(r *http.Request, sr *login.Request) error {
	config, err := s.populateMethod(r, sr.ID)
	if err != nil {
		return err
	}
	sr.Methods[identity.CredentialsTypeOIDC] = &login.RequestMethod{
		Method: identity.CredentialsTypeOIDC,
		Config: config,
	}
	return nil
}

func (s *Strategy) PopulateRegistrationMethod(r *http.Request, sr *registration.Request) error {
	config, err := s.populateMethod(r, sr.ID)
	if err != nil {
		return err
	}
	sr.Methods[identity.CredentialsTypeOIDC] = &registration.RequestMethod{
		Method: identity.CredentialsTypeOIDC,
		Config: config,
	}
	return nil
}

func (s *Strategy) Config() (*ConfigurationCollection, error) {
	var c ConfigurationCollection

	if err := jsonx.
		NewStrictDecoder(
			bytes.NewBuffer(s.c.SelfServiceStrategy(string(identity.CredentialsTypeOIDC)).Config),
		).
		Decode(&c); err != nil {
		return nil, errors.WithStack(herodot.ErrInternalServerError.WithReasonf("Unable to decode OpenID Connect Provider configuration: %s", err))
	}

	return &c, nil
}

func (s *Strategy) provider(id string) (Provider, error) {
	if c, err := s.Config(); err != nil {
		return nil, err
	} else if provider, err := c.Provider(id, s.c.SelfPublicURL()); err != nil {
		return nil, err
	} else {
		return provider, nil
	}
}

func (s *Strategy) handleError(w http.ResponseWriter, r *http.Request, rid string, traits json.RawMessage, err error) {
	if rid == "" {
		s.d.ErrorManager().ForwardError(r.Context(), w, r, err)
		return
	}

	if lr, rerr := s.d.LoginRequestPersister().GetLoginRequest(r.Context(), rid); rerr == nil {
		s.d.LoginRequestErrorHandler().HandleLoginError(w, r, identity.CredentialsTypeOIDC, lr, err)
		return
	} else if rr, rerr := s.d.RegistrationRequestPersister().GetRegistrationRequest(r.Context(), rid); rerr == nil {
		if method, ok := rr.Methods[s.ID()]; ok {
			method.Config.Reset()

			if traits != nil {
				for name, field := range form.NewHTMLFormFromJSON("", traits, "traits").Fields {
					method.Config.SetField(name, field)
				}
			}

			method.Config.SetField("request", form.Field{
				Name:     "request",
				Type:     "hidden",
				Required: true,
				Value:    r.PostForm.Get("request"),
			})
			method.Config.SetCSRF(s.cg(r))

			rr.Methods[s.ID()] = method
		}

		s.d.RegistrationRequestErrorHandler().HandleRegistrationError(w, r, identity.CredentialsTypeOIDC, rr, err)
		return
	}

	s.d.ErrorManager().ForwardError(r.Context(), w, r, err)
}
