package login

import (
	"database/sql/driver"
	"encoding/json"
	"time"

	"github.com/ory/x/sqlxx"

	"github.com/gofrs/uuid"

	"github.com/ory/kratos/identity"
	"github.com/ory/kratos/selfservice/form"
)

// swagger:model loginRequestMethod
type RequestMethod struct {
	// Method contains the request credentials type.
	//
	// required: true
	Method identity.CredentialsType `json:"method" db:"method"`

	// Config is the credential type's config.
	//
	// required: true
	Config *RequestMethodConfig `json:"config" db:"config"`

	// ID is a helper struct field for gobuffalo.pop.
	ID uuid.UUID `json:"-" db:"id"`

	// RequestID is a helper struct field for gobuffalo.pop.
	RequestID uuid.UUID `json:"-" db:"selfservice_login_request_id"`

	// Request is a helper struct field for gobuffalo.pop.
	Request *Request `json:"-" belongs_to:"selfservice_login_request" fk_id:"RequestID"`

	// CreatedAt is a helper struct field for gobuffalo.pop.
	CreatedAt time.Time `json:"-" db:"created_at"`

	// UpdatedAt is a helper struct field for gobuffalo.pop.
	UpdatedAt time.Time `json:"-" db:"updated_at"`
}

func (u RequestMethod) TableName() string {
	return "selfservice_login_request_methods"
}

type RequestMethodsRaw []RequestMethod // workaround for https://github.com/gobuffalo/pop/pull/478
type RequestMethods map[identity.CredentialsType]*RequestMethod

func (u RequestMethodsRaw) TableName() string {
	// This must be stay a value receiver, using a pointer receiver will cause issues with pop.
	return "selfservice_login_request_methods"
}

func (u RequestMethods) TableName() string {
	// This must be stay a value receiver, using a pointer receiver will cause issues with pop.
	return "selfservice_login_request_methods"
}

// swagger:ignore
type RequestMethodConfigurator interface {
	form.ErrorParser
	form.ValueSetter
	form.Resetter
	form.ErrorResetter
	form.CSRFSetter
	form.ErrorAdder
}

// swagger:model loginRequestMethodConfig
type RequestMethodConfig struct {
	// swagger:ignore
	RequestMethodConfigurator

	requestMethodConfigMock
}

// swagger:model loginRequestMethodConfigPayload
type requestMethodConfigMock struct {
	*form.HTMLForm

	// Providers is set for the "oidc" request method.
	Providers []form.Field `json:"providers"`
}

func (c *RequestMethodConfig) Scan(value interface{}) error {
	return sqlxx.JSONScan(c, value)
}

func (c *RequestMethodConfig) Value() (driver.Value, error) {
	return sqlxx.JSONValue(c)
}

func (c *RequestMethodConfig) UnmarshalJSON(data []byte) error {
	c.RequestMethodConfigurator = form.NewHTMLForm("")
	return json.Unmarshal(data, c.RequestMethodConfigurator)
}

func (c *RequestMethodConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.RequestMethodConfigurator)
}
