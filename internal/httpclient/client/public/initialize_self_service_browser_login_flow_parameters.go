// Code generated by go-swagger; DO NOT EDIT.

package public

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewInitializeSelfServiceBrowserLoginFlowParams creates a new InitializeSelfServiceBrowserLoginFlowParams object
// with the default values initialized.
func NewInitializeSelfServiceBrowserLoginFlowParams() *InitializeSelfServiceBrowserLoginFlowParams {
	var ()
	return &InitializeSelfServiceBrowserLoginFlowParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewInitializeSelfServiceBrowserLoginFlowParamsWithTimeout creates a new InitializeSelfServiceBrowserLoginFlowParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewInitializeSelfServiceBrowserLoginFlowParamsWithTimeout(timeout time.Duration) *InitializeSelfServiceBrowserLoginFlowParams {
	var ()
	return &InitializeSelfServiceBrowserLoginFlowParams{

		timeout: timeout,
	}
}

// NewInitializeSelfServiceBrowserLoginFlowParamsWithContext creates a new InitializeSelfServiceBrowserLoginFlowParams object
// with the default values initialized, and the ability to set a context for a request
func NewInitializeSelfServiceBrowserLoginFlowParamsWithContext(ctx context.Context) *InitializeSelfServiceBrowserLoginFlowParams {
	var ()
	return &InitializeSelfServiceBrowserLoginFlowParams{

		Context: ctx,
	}
}

// NewInitializeSelfServiceBrowserLoginFlowParamsWithHTTPClient creates a new InitializeSelfServiceBrowserLoginFlowParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewInitializeSelfServiceBrowserLoginFlowParamsWithHTTPClient(client *http.Client) *InitializeSelfServiceBrowserLoginFlowParams {
	var ()
	return &InitializeSelfServiceBrowserLoginFlowParams{
		HTTPClient: client,
	}
}

/*InitializeSelfServiceBrowserLoginFlowParams contains all the parameters to send to the API endpoint
for the initialize self service browser login flow operation typically these are written to a http.Request
*/
type InitializeSelfServiceBrowserLoginFlowParams struct {

	/*Refresh
	  Refresh a login session

	If set to true, this will refresh an existing login session by
	asking the user to sign in again. This will reset the
	authenticated_at time of the session.

	*/
	Refresh *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the initialize self service browser login flow params
func (o *InitializeSelfServiceBrowserLoginFlowParams) WithTimeout(timeout time.Duration) *InitializeSelfServiceBrowserLoginFlowParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the initialize self service browser login flow params
func (o *InitializeSelfServiceBrowserLoginFlowParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the initialize self service browser login flow params
func (o *InitializeSelfServiceBrowserLoginFlowParams) WithContext(ctx context.Context) *InitializeSelfServiceBrowserLoginFlowParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the initialize self service browser login flow params
func (o *InitializeSelfServiceBrowserLoginFlowParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the initialize self service browser login flow params
func (o *InitializeSelfServiceBrowserLoginFlowParams) WithHTTPClient(client *http.Client) *InitializeSelfServiceBrowserLoginFlowParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the initialize self service browser login flow params
func (o *InitializeSelfServiceBrowserLoginFlowParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRefresh adds the refresh to the initialize self service browser login flow params
func (o *InitializeSelfServiceBrowserLoginFlowParams) WithRefresh(refresh *bool) *InitializeSelfServiceBrowserLoginFlowParams {
	o.SetRefresh(refresh)
	return o
}

// SetRefresh adds the refresh to the initialize self service browser login flow params
func (o *InitializeSelfServiceBrowserLoginFlowParams) SetRefresh(refresh *bool) {
	o.Refresh = refresh
}

// WriteToRequest writes these params to a swagger request
func (o *InitializeSelfServiceBrowserLoginFlowParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Refresh != nil {

		// query param refresh
		var qrRefresh bool
		if o.Refresh != nil {
			qrRefresh = *o.Refresh
		}
		qRefresh := swag.FormatBool(qrRefresh)
		if qRefresh != "" {
			if err := r.SetQueryParam("refresh", qRefresh); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
