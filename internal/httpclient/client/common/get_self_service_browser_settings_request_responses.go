// Code generated by go-swagger; DO NOT EDIT.

package common

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ory/kratos/internal/httpclient/models"
)

// GetSelfServiceBrowserSettingsRequestReader is a Reader for the GetSelfServiceBrowserSettingsRequest structure.
type GetSelfServiceBrowserSettingsRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSelfServiceBrowserSettingsRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSelfServiceBrowserSettingsRequestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetSelfServiceBrowserSettingsRequestForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetSelfServiceBrowserSettingsRequestNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewGetSelfServiceBrowserSettingsRequestGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSelfServiceBrowserSettingsRequestInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSelfServiceBrowserSettingsRequestOK creates a GetSelfServiceBrowserSettingsRequestOK with default headers values
func NewGetSelfServiceBrowserSettingsRequestOK() *GetSelfServiceBrowserSettingsRequestOK {
	return &GetSelfServiceBrowserSettingsRequestOK{}
}

/*GetSelfServiceBrowserSettingsRequestOK handles this case with default header values.

settingsRequest
*/
type GetSelfServiceBrowserSettingsRequestOK struct {
	Payload *models.SettingsRequest
}

func (o *GetSelfServiceBrowserSettingsRequestOK) Error() string {
	return fmt.Sprintf("[GET /self-service/browser/flows/requests/settings][%d] getSelfServiceBrowserSettingsRequestOK  %+v", 200, o.Payload)
}

func (o *GetSelfServiceBrowserSettingsRequestOK) GetPayload() *models.SettingsRequest {
	return o.Payload
}

func (o *GetSelfServiceBrowserSettingsRequestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SettingsRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSelfServiceBrowserSettingsRequestForbidden creates a GetSelfServiceBrowserSettingsRequestForbidden with default headers values
func NewGetSelfServiceBrowserSettingsRequestForbidden() *GetSelfServiceBrowserSettingsRequestForbidden {
	return &GetSelfServiceBrowserSettingsRequestForbidden{}
}

/*GetSelfServiceBrowserSettingsRequestForbidden handles this case with default header values.

genericError
*/
type GetSelfServiceBrowserSettingsRequestForbidden struct {
	Payload *models.GenericError
}

func (o *GetSelfServiceBrowserSettingsRequestForbidden) Error() string {
	return fmt.Sprintf("[GET /self-service/browser/flows/requests/settings][%d] getSelfServiceBrowserSettingsRequestForbidden  %+v", 403, o.Payload)
}

func (o *GetSelfServiceBrowserSettingsRequestForbidden) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *GetSelfServiceBrowserSettingsRequestForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSelfServiceBrowserSettingsRequestNotFound creates a GetSelfServiceBrowserSettingsRequestNotFound with default headers values
func NewGetSelfServiceBrowserSettingsRequestNotFound() *GetSelfServiceBrowserSettingsRequestNotFound {
	return &GetSelfServiceBrowserSettingsRequestNotFound{}
}

/*GetSelfServiceBrowserSettingsRequestNotFound handles this case with default header values.

genericError
*/
type GetSelfServiceBrowserSettingsRequestNotFound struct {
	Payload *models.GenericError
}

func (o *GetSelfServiceBrowserSettingsRequestNotFound) Error() string {
	return fmt.Sprintf("[GET /self-service/browser/flows/requests/settings][%d] getSelfServiceBrowserSettingsRequestNotFound  %+v", 404, o.Payload)
}

func (o *GetSelfServiceBrowserSettingsRequestNotFound) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *GetSelfServiceBrowserSettingsRequestNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSelfServiceBrowserSettingsRequestGone creates a GetSelfServiceBrowserSettingsRequestGone with default headers values
func NewGetSelfServiceBrowserSettingsRequestGone() *GetSelfServiceBrowserSettingsRequestGone {
	return &GetSelfServiceBrowserSettingsRequestGone{}
}

/*GetSelfServiceBrowserSettingsRequestGone handles this case with default header values.

genericError
*/
type GetSelfServiceBrowserSettingsRequestGone struct {
	Payload *models.GenericError
}

func (o *GetSelfServiceBrowserSettingsRequestGone) Error() string {
	return fmt.Sprintf("[GET /self-service/browser/flows/requests/settings][%d] getSelfServiceBrowserSettingsRequestGone  %+v", 410, o.Payload)
}

func (o *GetSelfServiceBrowserSettingsRequestGone) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *GetSelfServiceBrowserSettingsRequestGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSelfServiceBrowserSettingsRequestInternalServerError creates a GetSelfServiceBrowserSettingsRequestInternalServerError with default headers values
func NewGetSelfServiceBrowserSettingsRequestInternalServerError() *GetSelfServiceBrowserSettingsRequestInternalServerError {
	return &GetSelfServiceBrowserSettingsRequestInternalServerError{}
}

/*GetSelfServiceBrowserSettingsRequestInternalServerError handles this case with default header values.

genericError
*/
type GetSelfServiceBrowserSettingsRequestInternalServerError struct {
	Payload *models.GenericError
}

func (o *GetSelfServiceBrowserSettingsRequestInternalServerError) Error() string {
	return fmt.Sprintf("[GET /self-service/browser/flows/requests/settings][%d] getSelfServiceBrowserSettingsRequestInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSelfServiceBrowserSettingsRequestInternalServerError) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *GetSelfServiceBrowserSettingsRequestInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
