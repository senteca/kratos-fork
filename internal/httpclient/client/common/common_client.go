// Code generated by go-swagger; DO NOT EDIT.

package common

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new common API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for common API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetSchema(params *GetSchemaParams) (*GetSchemaOK, error)

	GetSelfServiceError(params *GetSelfServiceErrorParams) (*GetSelfServiceErrorOK, error)

	GetSelfServiceLoginFlow(params *GetSelfServiceLoginFlowParams) (*GetSelfServiceLoginFlowOK, error)

	GetSelfServiceRecoveryFlow(params *GetSelfServiceRecoveryFlowParams) (*GetSelfServiceRecoveryFlowOK, error)

	GetSelfServiceRegistrationFlow(params *GetSelfServiceRegistrationFlowParams) (*GetSelfServiceRegistrationFlowOK, error)

	GetSelfServiceSettingsFlow(params *GetSelfServiceSettingsFlowParams) (*GetSelfServiceSettingsFlowOK, error)

	GetSelfServiceVerificationFlow(params *GetSelfServiceVerificationFlowParams) (*GetSelfServiceVerificationFlowOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetSchema Get a Traits Schema Definition
*/
func (a *Client) GetSchema(params *GetSchemaParams) (*GetSchemaOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSchemaParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSchema",
		Method:             "GET",
		PathPattern:        "/schemas/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetSchemaReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetSchemaOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getSchema: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetSelfServiceError gets user facing self service errors

  This endpoint returns the error associated with a user-facing self service errors.

When accessing this endpoint through ORY Kratos' Public API, ensure that cookies are set as they are required for CSRF to work. To prevent
token scanning attacks, the public endpoint does not return 404 status codes.

This endpoint supports stub values to help you implement the error UI:

`?error=stub:500` - returns a stub 500 (Internal Server Error) error.

More information can be found at [ORY Kratos User User Facing Error Documentation](https://www.ory.sh/docs/kratos/self-service/flows/user-facing-errors).
*/
func (a *Client) GetSelfServiceError(params *GetSelfServiceErrorParams) (*GetSelfServiceErrorOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSelfServiceErrorParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSelfServiceError",
		Method:             "GET",
		PathPattern:        "/self-service/errors",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetSelfServiceErrorReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetSelfServiceErrorOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getSelfServiceError: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetSelfServiceLoginFlow gets login flow

  This endpoint returns a login flow's context with, for example, error details and other information.

More information can be found at [ORY Kratos User Login and User Registration Documentation](https://www.ory.sh/docs/next/kratos/self-service/flows/user-login-user-registration).
*/
func (a *Client) GetSelfServiceLoginFlow(params *GetSelfServiceLoginFlowParams) (*GetSelfServiceLoginFlowOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSelfServiceLoginFlowParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSelfServiceLoginFlow",
		Method:             "GET",
		PathPattern:        "/self-service/login/flows",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetSelfServiceLoginFlowReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetSelfServiceLoginFlowOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getSelfServiceLoginFlow: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetSelfServiceRecoveryFlow gets information about a recovery flow

  This endpoint returns a recovery flow's context with, for example, error details and other information.

More information can be found at [ORY Kratos Account Recovery Documentation](../self-service/flows/password-reset-account-recovery).
*/
func (a *Client) GetSelfServiceRecoveryFlow(params *GetSelfServiceRecoveryFlowParams) (*GetSelfServiceRecoveryFlowOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSelfServiceRecoveryFlowParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSelfServiceRecoveryFlow",
		Method:             "GET",
		PathPattern:        "/self-service/recovery/flows",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetSelfServiceRecoveryFlowReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetSelfServiceRecoveryFlowOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getSelfServiceRecoveryFlow: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetSelfServiceRegistrationFlow gets registration flow

  This endpoint returns a registration flow's context with, for example, error details and other information.

More information can be found at [ORY Kratos User Login and User Registration Documentation](https://www.ory.sh/docs/next/kratos/self-service/flows/user-login-user-registration).
*/
func (a *Client) GetSelfServiceRegistrationFlow(params *GetSelfServiceRegistrationFlowParams) (*GetSelfServiceRegistrationFlowOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSelfServiceRegistrationFlowParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSelfServiceRegistrationFlow",
		Method:             "GET",
		PathPattern:        "/self-service/registration/flows",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetSelfServiceRegistrationFlowReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetSelfServiceRegistrationFlowOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getSelfServiceRegistrationFlow: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetSelfServiceSettingsFlow gets settings flow

  When accessing this endpoint through ORY Kratos' Public API you must ensure that either the ORY Kratos Session Cookie
or the ORY Kratos Session Token are set. The public endpoint does not return 404 status codes
but instead 403 or 500 to improve data privacy.

You can access this endpoint without credentials when using ORY Kratos' Admin API.

More information can be found at [ORY Kratos User Settings & Profile Management Documentation](../self-service/flows/user-settings).
*/
func (a *Client) GetSelfServiceSettingsFlow(params *GetSelfServiceSettingsFlowParams) (*GetSelfServiceSettingsFlowOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSelfServiceSettingsFlowParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSelfServiceSettingsFlow",
		Method:             "GET",
		PathPattern:        "/self-service/settings/flows",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetSelfServiceSettingsFlowReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetSelfServiceSettingsFlowOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getSelfServiceSettingsFlow: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetSelfServiceVerificationFlow gets verification flow

  This endpoint returns a verification flow's context with, for example, error details and other information.

More information can be found at [ORY Kratos Email and Phone Verification Documentation](https://www.ory.sh/docs/kratos/selfservice/flows/verify-email-account-activation).
*/
func (a *Client) GetSelfServiceVerificationFlow(params *GetSelfServiceVerificationFlowParams) (*GetSelfServiceVerificationFlowOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSelfServiceVerificationFlowParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSelfServiceVerificationFlow",
		Method:             "GET",
		PathPattern:        "/self-service/verification/flows",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetSelfServiceVerificationFlowReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetSelfServiceVerificationFlowOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getSelfServiceVerificationFlow: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
