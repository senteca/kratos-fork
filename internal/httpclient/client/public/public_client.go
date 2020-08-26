// Code generated by go-swagger; DO NOT EDIT.

package public

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new public API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for public API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CompleteSelfServiceBrowserRecoveryLinkStrategyFlow(params *CompleteSelfServiceBrowserRecoveryLinkStrategyFlowParams) error

	CompleteSelfServiceBrowserSettingsOIDCSettingsFlow(params *CompleteSelfServiceBrowserSettingsOIDCSettingsFlowParams) error

	CompleteSelfServiceBrowserVerificationFlow(params *CompleteSelfServiceBrowserVerificationFlowParams) error

	CompleteSelfServiceLoginFlowWithPasswordMethod(params *CompleteSelfServiceLoginFlowWithPasswordMethodParams) (*CompleteSelfServiceLoginFlowWithPasswordMethodOK, error)

	CompleteSelfServiceRegistrationFlowWithPasswordMethod(params *CompleteSelfServiceRegistrationFlowWithPasswordMethodParams) (*CompleteSelfServiceRegistrationFlowWithPasswordMethodOK, error)

	CompleteSelfServiceSettingsFlowWithPasswordMethod(params *CompleteSelfServiceSettingsFlowWithPasswordMethodParams) error

	CompleteSelfServiceSettingsFlowWithProfileMethod(params *CompleteSelfServiceSettingsFlowWithProfileMethodParams) (*CompleteSelfServiceSettingsFlowWithProfileMethodOK, error)

	InitializeSelfServiceBrowserLogoutFlow(params *InitializeSelfServiceBrowserLogoutFlowParams) error

	InitializeSelfServiceBrowserVerificationFlow(params *InitializeSelfServiceBrowserVerificationFlowParams) error

	InitializeSelfServiceLoginViaBrowserFlow(params *InitializeSelfServiceLoginViaBrowserFlowParams) error

	InitializeSelfServiceRecoveryFlow(params *InitializeSelfServiceRecoveryFlowParams) error

	InitializeSelfServiceRegistrationViaBrowserFlow(params *InitializeSelfServiceRegistrationViaBrowserFlowParams) error

	InitializeSelfServiceSettingsViaBrowserFlow(params *InitializeSelfServiceSettingsViaBrowserFlowParams) error

	RevokeSession(params *RevokeSessionParams) (*RevokeSessionNoContent, error)

	SelfServiceBrowserVerify(params *SelfServiceBrowserVerifyParams) error

	Whoami(params *WhoamiParams) (*WhoamiOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CompleteSelfServiceBrowserRecoveryLinkStrategyFlow completes the browser based recovery flow using a recovery link

  > This endpoint is NOT INTENDED for API clients and only works with browsers (Chrome, Firefox, ...) and HTML Forms.

More information can be found at [ORY Kratos Account Recovery Documentation](../self-service/flows/password-reset-account-recovery).
*/
func (a *Client) CompleteSelfServiceBrowserRecoveryLinkStrategyFlow(params *CompleteSelfServiceBrowserRecoveryLinkStrategyFlowParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCompleteSelfServiceBrowserRecoveryLinkStrategyFlowParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "completeSelfServiceBrowserRecoveryLinkStrategyFlow",
		Method:             "POST",
		PathPattern:        "/self-service/browser/flows/recovery/link",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CompleteSelfServiceBrowserRecoveryLinkStrategyFlowReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil
}

/*
  CompleteSelfServiceBrowserSettingsOIDCSettingsFlow completes the browser based settings flow for the open ID connect strategy

  This endpoint completes a browser-based settings flow. This is usually achieved by POSTing data to this
endpoint.

> This endpoint is NOT INTENDED for API clients and only works with browsers (Chrome, Firefox, ...) and HTML Forms.

More information can be found at [ORY Kratos User Settings & Profile Management Documentation](../self-service/flows/user-settings).
*/
func (a *Client) CompleteSelfServiceBrowserSettingsOIDCSettingsFlow(params *CompleteSelfServiceBrowserSettingsOIDCSettingsFlowParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCompleteSelfServiceBrowserSettingsOIDCSettingsFlowParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "completeSelfServiceBrowserSettingsOIDCSettingsFlow",
		Method:             "POST",
		PathPattern:        "/self-service/browser/flows/registration/strategies/oidc/settings/connections",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CompleteSelfServiceBrowserSettingsOIDCSettingsFlowReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil
}

/*
  CompleteSelfServiceBrowserVerificationFlow completes the browser based verification flows

  This endpoint completes a browser-based verification flow. This is usually achieved by POSTing data to this
endpoint.

If the provided data is valid against the Identity's Traits JSON Schema, the data will be updated and
the browser redirected to `url.settings_ui` for further steps.

> This endpoint is NOT INTENDED for API clients and only works with browsers (Chrome, Firefox, ...) and HTML Forms.

More information can be found at [ORY Kratos Email and Phone Verification Documentation](https://www.ory.sh/docs/kratos/selfservice/flows/verify-email-account-activation).
*/
func (a *Client) CompleteSelfServiceBrowserVerificationFlow(params *CompleteSelfServiceBrowserVerificationFlowParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCompleteSelfServiceBrowserVerificationFlowParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "completeSelfServiceBrowserVerificationFlow",
		Method:             "POST",
		PathPattern:        "/self-service/browser/flows/verification/{via}/complete",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CompleteSelfServiceBrowserVerificationFlowReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil
}

/*
  CompleteSelfServiceLoginFlowWithPasswordMethod completes login flow with username email password method

  Use this endpoint to complete a login flow by sending an identity's identifier and password. This endpoint
behaves differently for API and browser flows.

API flows expect `application/json` to be sent in the body and responds with
HTTP 200 and a application/json body with the session token on success;
HTTP 302 redirect to a fresh login flow if the original flow expired with the appropriate error messages set;
HTTP 400 on form validation errors.

Browser flows expect `application/x-www-form-urlencoded` to be sent in the body and responds with
a HTTP 302 redirect to the post/after login URL or the `return_to` value if it was set and if the login succeeded;
a HTTP 302 redirect to the login UI URL with the flow ID containing the validation errors otherwise.

More information can be found at [ORY Kratos User Login and User Registration Documentation](https://www.ory.sh/docs/next/kratos/self-service/flows/user-login-user-registration).
*/
func (a *Client) CompleteSelfServiceLoginFlowWithPasswordMethod(params *CompleteSelfServiceLoginFlowWithPasswordMethodParams) (*CompleteSelfServiceLoginFlowWithPasswordMethodOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCompleteSelfServiceLoginFlowWithPasswordMethodParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "completeSelfServiceLoginFlowWithPasswordMethod",
		Method:             "GET",
		PathPattern:        "/self-service/login/methods/password",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CompleteSelfServiceLoginFlowWithPasswordMethodReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CompleteSelfServiceLoginFlowWithPasswordMethodOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for completeSelfServiceLoginFlowWithPasswordMethod: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CompleteSelfServiceRegistrationFlowWithPasswordMethod completes registration flow with username email password method

  Use this endpoint to complete a registration flow by sending an identity's traits and password. This endpoint
behaves differently for API and browser flows.

API flows expect `application/json` to be sent in the body and responds with
HTTP 200 and a application/json body with the created identity success - if the session hook is configured the
`session` and `session_token` will also be included;
HTTP 302 redirect to a fresh registration flow if the original flow expired with the appropriate error messages set;
HTTP 400 on form validation errors.

Browser flows expect `application/x-www-form-urlencoded` to be sent in the body and responds with
a HTTP 302 redirect to the post/after registration URL or the `return_to` value if it was set and if the registration succeeded;
a HTTP 302 redirect to the registration UI URL with the flow ID containing the validation errors otherwise.

More information can be found at [ORY Kratos User Login and User Registration Documentation](https://www.ory.sh/docs/next/kratos/self-service/flows/user-login-user-registration).
*/
func (a *Client) CompleteSelfServiceRegistrationFlowWithPasswordMethod(params *CompleteSelfServiceRegistrationFlowWithPasswordMethodParams) (*CompleteSelfServiceRegistrationFlowWithPasswordMethodOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCompleteSelfServiceRegistrationFlowWithPasswordMethodParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "completeSelfServiceRegistrationFlowWithPasswordMethod",
		Method:             "POST",
		PathPattern:        "/self-service/registration/methods/password",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CompleteSelfServiceRegistrationFlowWithPasswordMethodReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CompleteSelfServiceRegistrationFlowWithPasswordMethodOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for completeSelfServiceRegistrationFlowWithPasswordMethod: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CompleteSelfServiceSettingsFlowWithPasswordMethod completes the browser based settings flow for the password strategy

  This endpoint completes a browser-based settings flow. This is usually achieved by POSTing data to this
endpoint.

> This endpoint is NOT INTENDED for API clients and only works with browsers (Chrome, Firefox, ...) and HTML Forms.

More information can be found at [ORY Kratos User Settings & Profile Management Documentation](../self-service/flows/user-settings).
*/
func (a *Client) CompleteSelfServiceSettingsFlowWithPasswordMethod(params *CompleteSelfServiceSettingsFlowWithPasswordMethodParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCompleteSelfServiceSettingsFlowWithPasswordMethodParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "completeSelfServiceSettingsFlowWithPasswordMethod",
		Method:             "POST",
		PathPattern:        "/self-service/settings/methods/password",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CompleteSelfServiceSettingsFlowWithPasswordMethodReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil
}

/*
  CompleteSelfServiceSettingsFlowWithProfileMethod completes the browser based settings flow for profile data

  This endpoint completes a browser-based settings flow. This is usually achieved by POSTing data to this
endpoint.

If the provided profile data is valid against the Identity's Traits JSON Schema, the data will be updated and
the browser redirected to `url.settings_ui` for further steps.

> This endpoint is NOT INTENDED for API clients and only works with browsers (Chrome, Firefox, ...) and HTML Forms.

More information can be found at [ORY Kratos User Settings & Profile Management Documentation](../self-service/flows/user-settings).
*/
func (a *Client) CompleteSelfServiceSettingsFlowWithProfileMethod(params *CompleteSelfServiceSettingsFlowWithProfileMethodParams) (*CompleteSelfServiceSettingsFlowWithProfileMethodOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCompleteSelfServiceSettingsFlowWithProfileMethodParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "completeSelfServiceSettingsFlowWithProfileMethod",
		Method:             "POST",
		PathPattern:        "/self-service/settings/methods/profile",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CompleteSelfServiceSettingsFlowWithProfileMethodReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CompleteSelfServiceSettingsFlowWithProfileMethodOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for completeSelfServiceSettingsFlowWithProfileMethod: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  InitializeSelfServiceBrowserLogoutFlow initializes browser based logout user flow

  This endpoint initializes a logout flow.

> This endpoint is NOT INTENDED for API clients and only works
with browsers (Chrome, Firefox, ...).

On successful logout, the browser will be redirected (HTTP 302 Found) to `urls.default_return_to`.

More information can be found at [ORY Kratos User Logout Documentation](https://www.ory.sh/docs/next/kratos/self-service/flows/user-logout).
*/
func (a *Client) InitializeSelfServiceBrowserLogoutFlow(params *InitializeSelfServiceBrowserLogoutFlowParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInitializeSelfServiceBrowserLogoutFlowParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "initializeSelfServiceBrowserLogoutFlow",
		Method:             "GET",
		PathPattern:        "/self-service/browser/flows/logout",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &InitializeSelfServiceBrowserLogoutFlowReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil
}

/*
  InitializeSelfServiceBrowserVerificationFlow initializes browser based verification flow

  This endpoint initializes a browser-based verification flow. Once initialized, the browser will be redirected to
`selfservice.flows.settings.ui_url` with the request ID set as a query parameter. If no valid user session exists, a login
flow will be initialized.

> This endpoint is NOT INTENDED for API clients and only works
with browsers (Chrome, Firefox, ...).

More information can be found at [ORY Kratos Email and Phone Verification Documentation](https://www.ory.sh/docs/kratos/selfservice/flows/verify-email-account-activation).
*/
func (a *Client) InitializeSelfServiceBrowserVerificationFlow(params *InitializeSelfServiceBrowserVerificationFlowParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInitializeSelfServiceBrowserVerificationFlowParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "initializeSelfServiceBrowserVerificationFlow",
		Method:             "GET",
		PathPattern:        "/self-service/browser/flows/verification/init/{via}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &InitializeSelfServiceBrowserVerificationFlowReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil
}

/*
  InitializeSelfServiceLoginViaBrowserFlow initializes login flow for browsers

  This endpoint initializes a browser-based user login flow. Once initialized, the browser will be redirected to
`selfservice.flows.login.ui_url` with the flow ID set as the query parameter `?flow=`. If a valid user session
exists already, the browser will be redirected to `urls.default_redirect_url` unless the query parameter
`?refresh=true` was set.

This endpoint is NOT INTENDED for API clients and only works with browsers (Chrome, Firefox, ...).

More information can be found at [ORY Kratos User Login and User Registration Documentation](https://www.ory.sh/docs/next/kratos/self-service/flows/user-login-user-registration).
*/
func (a *Client) InitializeSelfServiceLoginViaBrowserFlow(params *InitializeSelfServiceLoginViaBrowserFlowParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInitializeSelfServiceLoginViaBrowserFlowParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "initializeSelfServiceLoginViaBrowserFlow",
		Method:             "GET",
		PathPattern:        "/self-service/login/browser",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &InitializeSelfServiceLoginViaBrowserFlowReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil
}

/*
  InitializeSelfServiceRecoveryFlow initializes browser based account recovery flow

  This endpoint initializes a browser-based account recovery flow. Once initialized, the browser will be redirected to
`selfservice.flows.recovery.ui_url` with the request ID set as a query parameter. If a valid user session exists, the request
is aborted.

> This endpoint is NOT INTENDED for API clients and only works
with browsers (Chrome, Firefox, ...).

More information can be found at [ORY Kratos Account Recovery Documentation](../self-service/flows/password-reset-account-recovery).
*/
func (a *Client) InitializeSelfServiceRecoveryFlow(params *InitializeSelfServiceRecoveryFlowParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInitializeSelfServiceRecoveryFlowParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "initializeSelfServiceRecoveryFlow",
		Method:             "GET",
		PathPattern:        "/self-service/browser/flows/recovery",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &InitializeSelfServiceRecoveryFlowReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil
}

/*
  InitializeSelfServiceRegistrationViaBrowserFlow initializes registration flow for browsers

  This endpoint initializes a browser-based user registration flow. Once initialized, the browser will be redirected to
`selfservice.flows.registration.ui_url` with the flow ID set as the query parameter `?flow=`. If a valid user session
exists already, the browser will be redirected to `urls.default_redirect_url` unless the query parameter
`?refresh=true` was set.

:::note

This endpoint is NOT INTENDED for API clients and only works with browsers (Chrome, Firefox, ...).

:::

More information can be found at [ORY Kratos User Login and User Registration Documentation](https://www.ory.sh/docs/next/kratos/self-service/flows/user-login-user-registration).
*/
func (a *Client) InitializeSelfServiceRegistrationViaBrowserFlow(params *InitializeSelfServiceRegistrationViaBrowserFlowParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInitializeSelfServiceRegistrationViaBrowserFlowParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "initializeSelfServiceRegistrationViaBrowserFlow",
		Method:             "GET",
		PathPattern:        "/self-service/registration/browser",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &InitializeSelfServiceRegistrationViaBrowserFlowReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil
}

/*
  InitializeSelfServiceSettingsViaBrowserFlow initializes settings flow for browsers

  This endpoint initializes a browser-based user settings flow. Once initialized, the browser will be redirected to
`selfservice.flows.settings.ui_url` with the flow ID set as the query parameter `?flow=`. If no valid
ORY Kratos Session Cookie is included in the request, a login flow will be initialized.

:::note

This endpoint is NOT INTENDED for API clients and only works with browsers (Chrome, Firefox, ...).

:::

More information can be found at [ORY Kratos User Settings & Profile Management Documentation](../self-service/flows/user-settings).
*/
func (a *Client) InitializeSelfServiceSettingsViaBrowserFlow(params *InitializeSelfServiceSettingsViaBrowserFlowParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInitializeSelfServiceSettingsViaBrowserFlowParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "initializeSelfServiceSettingsViaBrowserFlow",
		Method:             "GET",
		PathPattern:        "/self-service/settings/browser/flows",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &InitializeSelfServiceSettingsViaBrowserFlowReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil
}

/*
  RevokeSession revokes and invalidate a session

  Use this endpoint to revoke a session using its token. This endpoint is particularly useful for API clients
such as mobile apps to log the user out of the system and invalidate the session.

This endpoint does not remove any HTTP Cookies - use the Self-Service Logout Flow instead.
*/
func (a *Client) RevokeSession(params *RevokeSessionParams) (*RevokeSessionNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRevokeSessionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "revokeSession",
		Method:             "DELETE",
		PathPattern:        "/sessions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RevokeSessionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RevokeSessionNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for revokeSession: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SelfServiceBrowserVerify completes the browser based verification flows

  This endpoint completes a browser-based verification flow.

> This endpoint is NOT INTENDED for API clients and only works with browsers (Chrome, Firefox, ...) and HTML Forms.

More information can be found at [ORY Kratos Email and Phone Verification Documentation](https://www.ory.sh/docs/kratos/selfservice/flows/verify-email-account-activation).
*/
func (a *Client) SelfServiceBrowserVerify(params *SelfServiceBrowserVerifyParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSelfServiceBrowserVerifyParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "selfServiceBrowserVerify",
		Method:             "GET",
		PathPattern:        "/self-service/browser/flows/verification/{via}/confirm/{code}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SelfServiceBrowserVerifyReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil
}

/*
  Whoami checks who the current HTTP session belongs to

  Uses the HTTP Headers in the GET request to determine (e.g. by using checking the cookies) who is authenticated.
Returns a session object in the body or 401 if the credentials are invalid or no credentials were sent.
Additionally when the request it successful it adds the user ID to the 'X-Kratos-Authenticated-Identity-Id' header in the response.

This endpoint is useful for reverse proxies and API Gateways.
*/
func (a *Client) Whoami(params *WhoamiParams) (*WhoamiOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWhoamiParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "whoami",
		Method:             "GET",
		PathPattern:        "/sessions/whoami",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &WhoamiReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*WhoamiOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for whoami: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
