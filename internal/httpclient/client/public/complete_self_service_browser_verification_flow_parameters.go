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
)

// NewCompleteSelfServiceBrowserVerificationFlowParams creates a new CompleteSelfServiceBrowserVerificationFlowParams object
// with the default values initialized.
func NewCompleteSelfServiceBrowserVerificationFlowParams() *CompleteSelfServiceBrowserVerificationFlowParams {
	var ()
	return &CompleteSelfServiceBrowserVerificationFlowParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCompleteSelfServiceBrowserVerificationFlowParamsWithTimeout creates a new CompleteSelfServiceBrowserVerificationFlowParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCompleteSelfServiceBrowserVerificationFlowParamsWithTimeout(timeout time.Duration) *CompleteSelfServiceBrowserVerificationFlowParams {
	var ()
	return &CompleteSelfServiceBrowserVerificationFlowParams{

		timeout: timeout,
	}
}

// NewCompleteSelfServiceBrowserVerificationFlowParamsWithContext creates a new CompleteSelfServiceBrowserVerificationFlowParams object
// with the default values initialized, and the ability to set a context for a request
func NewCompleteSelfServiceBrowserVerificationFlowParamsWithContext(ctx context.Context) *CompleteSelfServiceBrowserVerificationFlowParams {
	var ()
	return &CompleteSelfServiceBrowserVerificationFlowParams{

		Context: ctx,
	}
}

// NewCompleteSelfServiceBrowserVerificationFlowParamsWithHTTPClient creates a new CompleteSelfServiceBrowserVerificationFlowParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCompleteSelfServiceBrowserVerificationFlowParamsWithHTTPClient(client *http.Client) *CompleteSelfServiceBrowserVerificationFlowParams {
	var ()
	return &CompleteSelfServiceBrowserVerificationFlowParams{
		HTTPClient: client,
	}
}

/*CompleteSelfServiceBrowserVerificationFlowParams contains all the parameters to send to the API endpoint
for the complete self service browser verification flow operation typically these are written to a http.Request
*/
type CompleteSelfServiceBrowserVerificationFlowParams struct {

	/*Request
	  Request is the Request ID

	The value for this parameter comes from `request` URL Query parameter sent to your
	application (e.g. `/verify?request=abcde`).

	*/
	Request string
	/*Via
	  What to verify

	Currently only "email" is supported.

	*/
	Via string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the complete self service browser verification flow params
func (o *CompleteSelfServiceBrowserVerificationFlowParams) WithTimeout(timeout time.Duration) *CompleteSelfServiceBrowserVerificationFlowParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the complete self service browser verification flow params
func (o *CompleteSelfServiceBrowserVerificationFlowParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the complete self service browser verification flow params
func (o *CompleteSelfServiceBrowserVerificationFlowParams) WithContext(ctx context.Context) *CompleteSelfServiceBrowserVerificationFlowParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the complete self service browser verification flow params
func (o *CompleteSelfServiceBrowserVerificationFlowParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the complete self service browser verification flow params
func (o *CompleteSelfServiceBrowserVerificationFlowParams) WithHTTPClient(client *http.Client) *CompleteSelfServiceBrowserVerificationFlowParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the complete self service browser verification flow params
func (o *CompleteSelfServiceBrowserVerificationFlowParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the complete self service browser verification flow params
func (o *CompleteSelfServiceBrowserVerificationFlowParams) WithRequest(request string) *CompleteSelfServiceBrowserVerificationFlowParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the complete self service browser verification flow params
func (o *CompleteSelfServiceBrowserVerificationFlowParams) SetRequest(request string) {
	o.Request = request
}

// WithVia adds the via to the complete self service browser verification flow params
func (o *CompleteSelfServiceBrowserVerificationFlowParams) WithVia(via string) *CompleteSelfServiceBrowserVerificationFlowParams {
	o.SetVia(via)
	return o
}

// SetVia adds the via to the complete self service browser verification flow params
func (o *CompleteSelfServiceBrowserVerificationFlowParams) SetVia(via string) {
	o.Via = via
}

// WriteToRequest writes these params to a swagger request
func (o *CompleteSelfServiceBrowserVerificationFlowParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param request
	qrRequest := o.Request
	qRequest := qrRequest
	if qRequest != "" {
		if err := r.SetQueryParam("request", qRequest); err != nil {
			return err
		}
	}

	// path param via
	if err := r.SetPathParam("via", o.Via); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
