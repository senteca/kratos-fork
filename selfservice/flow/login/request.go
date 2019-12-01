package login

import (
	"net/http"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"

	"github.com/ory/herodot"
	"github.com/ory/x/urlx"

	"github.com/ory/kratos/identity"
	"github.com/ory/kratos/x"
)

// swagger:model loginRequest
type Request struct {
	// ID represents the request's unique ID. When performing the login flow, this
	// represents the id in the login ui's query parameter: http://<urls.login_ui>/?request=<id>
	ID uuid.UUID `json:"id" faker:"uuid" rw:"r" db:"id"`

	// ExpiresAt is the time (UTC) when the request expires. If the user still wishes to log in,
	// a new request has to be initiated.
	ExpiresAt time.Time `json:"expires_at" faker:"time_type" db:"expires_at"`

	// IssuedAt is the time (UTC) when the request occurred.
	IssuedAt time.Time `json:"issued_at" faker:"time_type" db:"issued_at"`

	// RequestURL is the initial URL that was requested from ORY Kratos. It can be used
	// to forward information contained in the URL's path or query for example.
	RequestURL string `json:"request_url" db:"request_url"`

	// Active, if set, contains the login method that is being used. It is initially
	// not set.
	Active identity.CredentialsType `json:"active,omitempty" db:"active_method"`

	// Methods contains context for all enabled login methods. If a login request has been
	// processed, but for example the password is incorrect, this will contain error messages.
	Methods map[identity.CredentialsType]*RequestMethod `json:"methods" faker:"login_request_methods" db:"-"`

	// MethodsRaw is a helper struct field for gobuffalo.pop.
	MethodsRaw RequestMethodsRaw `json:"-" faker:"-" has_many:"selfservice_login_request_methods" fk_id:"selfservice_login_request_id"`

	// CreatedAt is a helper struct field for gobuffalo.pop.
	CreatedAt time.Time `json:"-" db:"created_at"`

	// UpdatedAt is a helper struct field for gobuffalo.pop.
	UpdatedAt time.Time `json:"-" db:"updated_at"`
}

func NewLoginRequest(exp time.Duration, r *http.Request) *Request {
	source := urlx.Copy(r.URL)
	source.Host = r.Host

	if len(source.Scheme) == 0 {
		source.Scheme = "http"
		if r.TLS != nil {
			source.Scheme = "https"
		}
	}

	return &Request{
		ID:         x.NewUUID(),
		ExpiresAt:  time.Now().UTC().Add(exp),
		IssuedAt:   time.Now().UTC(),
		RequestURL: source.String(),
		Methods:    map[identity.CredentialsType]*RequestMethod{},
	}
}

func (r *Request) BeforeSave(_ *pop.Connection) error {
	r.MethodsRaw = make([]RequestMethod, 0, len(r.Methods))
	for _, m := range r.Methods {
		r.MethodsRaw = append(r.MethodsRaw, *m)
	}
	r.Methods = nil
	return nil
}

func (r *Request) AfterCreate(c *pop.Connection) error {
	return r.AfterFind(c)
}

func (r *Request) AfterUpdate(c *pop.Connection) error {
	return r.AfterFind(c)
}

func (r *Request) AfterFind(_ *pop.Connection) error {
	r.Methods = make(RequestMethods)
	for key := range r.MethodsRaw {
		m := r.MethodsRaw[key] // required for pointer dereference
		r.Methods[m.Method] = &m
	}
	r.MethodsRaw = nil
	return nil
}

func (r Request) TableName() string {
	// This must be stay a value receiver, using a pointer receiver will cause issues with pop.
	return "selfservice_login_requests"
}

func (r *Request) Valid() error {
	if r.ExpiresAt.Before(time.Now()) {
		return errors.WithStack(ErrRequestExpired.WithReasonf("The login request expired %.2f minutes ago, please try again.", time.Since(r.ExpiresAt).Minutes()))
	}

	if r.IssuedAt.After(time.Now()) {
		return errors.WithStack(herodot.ErrBadRequest.WithReason("The login request was issued in the future."))
	}
	return nil
}

func (r *Request) GetID() uuid.UUID {
	return r.ID
}
