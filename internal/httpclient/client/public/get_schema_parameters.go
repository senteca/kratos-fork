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

// NewGetSchemaParams creates a new GetSchemaParams object
// with the default values initialized.
func NewGetSchemaParams() *GetSchemaParams {
	var ()
	return &GetSchemaParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSchemaParamsWithTimeout creates a new GetSchemaParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSchemaParamsWithTimeout(timeout time.Duration) *GetSchemaParams {
	var ()
	return &GetSchemaParams{

		timeout: timeout,
	}
}

// NewGetSchemaParamsWithContext creates a new GetSchemaParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSchemaParamsWithContext(ctx context.Context) *GetSchemaParams {
	var ()
	return &GetSchemaParams{

		Context: ctx,
	}
}

// NewGetSchemaParamsWithHTTPClient creates a new GetSchemaParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSchemaParamsWithHTTPClient(client *http.Client) *GetSchemaParams {
	var ()
	return &GetSchemaParams{
		HTTPClient: client,
	}
}

/*GetSchemaParams contains all the parameters to send to the API endpoint
for the get schema operation typically these are written to a http.Request
*/
type GetSchemaParams struct {

	/*ID
	  ID must be set to the ID of schema you want to get

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get schema params
func (o *GetSchemaParams) WithTimeout(timeout time.Duration) *GetSchemaParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get schema params
func (o *GetSchemaParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get schema params
func (o *GetSchemaParams) WithContext(ctx context.Context) *GetSchemaParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get schema params
func (o *GetSchemaParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get schema params
func (o *GetSchemaParams) WithHTTPClient(client *http.Client) *GetSchemaParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get schema params
func (o *GetSchemaParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get schema params
func (o *GetSchemaParams) WithID(id string) *GetSchemaParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get schema params
func (o *GetSchemaParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetSchemaParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
