// Code generated by go-swagger; DO NOT EDIT.

package common

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/ory/kratos/internal/httpclient/models"
)

// GetSelfServiceBrowserLoginRequestReader is a Reader for the GetSelfServiceBrowserLoginRequest structure.
type GetSelfServiceBrowserLoginRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSelfServiceBrowserLoginRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSelfServiceBrowserLoginRequestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetSelfServiceBrowserLoginRequestForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetSelfServiceBrowserLoginRequestNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewGetSelfServiceBrowserLoginRequestGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSelfServiceBrowserLoginRequestInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSelfServiceBrowserLoginRequestOK creates a GetSelfServiceBrowserLoginRequestOK with default headers values
func NewGetSelfServiceBrowserLoginRequestOK() *GetSelfServiceBrowserLoginRequestOK {
	return &GetSelfServiceBrowserLoginRequestOK{}
}

/*GetSelfServiceBrowserLoginRequestOK handles this case with default header values.

loginRequest
*/
type GetSelfServiceBrowserLoginRequestOK struct {
	Payload *models.LoginRequest
}

func (o *GetSelfServiceBrowserLoginRequestOK) Error() string {
	return fmt.Sprintf("[GET /self-service/browser/flows/requests/login][%d] getSelfServiceBrowserLoginRequestOK  %+v", 200, o.Payload)
}

func (o *GetSelfServiceBrowserLoginRequestOK) GetPayload() *models.LoginRequest {
	return o.Payload
}

func (o *GetSelfServiceBrowserLoginRequestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LoginRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSelfServiceBrowserLoginRequestForbidden creates a GetSelfServiceBrowserLoginRequestForbidden with default headers values
func NewGetSelfServiceBrowserLoginRequestForbidden() *GetSelfServiceBrowserLoginRequestForbidden {
	return &GetSelfServiceBrowserLoginRequestForbidden{}
}

/*GetSelfServiceBrowserLoginRequestForbidden handles this case with default header values.

genericError
*/
type GetSelfServiceBrowserLoginRequestForbidden struct {
	Payload *models.GenericError
}

func (o *GetSelfServiceBrowserLoginRequestForbidden) Error() string {
	return fmt.Sprintf("[GET /self-service/browser/flows/requests/login][%d] getSelfServiceBrowserLoginRequestForbidden  %+v", 403, o.Payload)
}

func (o *GetSelfServiceBrowserLoginRequestForbidden) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *GetSelfServiceBrowserLoginRequestForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSelfServiceBrowserLoginRequestNotFound creates a GetSelfServiceBrowserLoginRequestNotFound with default headers values
func NewGetSelfServiceBrowserLoginRequestNotFound() *GetSelfServiceBrowserLoginRequestNotFound {
	return &GetSelfServiceBrowserLoginRequestNotFound{}
}

/*GetSelfServiceBrowserLoginRequestNotFound handles this case with default header values.

genericError
*/
type GetSelfServiceBrowserLoginRequestNotFound struct {
	Payload *models.GenericError
}

func (o *GetSelfServiceBrowserLoginRequestNotFound) Error() string {
	return fmt.Sprintf("[GET /self-service/browser/flows/requests/login][%d] getSelfServiceBrowserLoginRequestNotFound  %+v", 404, o.Payload)
}

func (o *GetSelfServiceBrowserLoginRequestNotFound) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *GetSelfServiceBrowserLoginRequestNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSelfServiceBrowserLoginRequestGone creates a GetSelfServiceBrowserLoginRequestGone with default headers values
func NewGetSelfServiceBrowserLoginRequestGone() *GetSelfServiceBrowserLoginRequestGone {
	return &GetSelfServiceBrowserLoginRequestGone{}
}

/*GetSelfServiceBrowserLoginRequestGone handles this case with default header values.

genericError
*/
type GetSelfServiceBrowserLoginRequestGone struct {
	Payload *models.GenericError
}

func (o *GetSelfServiceBrowserLoginRequestGone) Error() string {
	return fmt.Sprintf("[GET /self-service/browser/flows/requests/login][%d] getSelfServiceBrowserLoginRequestGone  %+v", 410, o.Payload)
}

func (o *GetSelfServiceBrowserLoginRequestGone) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *GetSelfServiceBrowserLoginRequestGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSelfServiceBrowserLoginRequestInternalServerError creates a GetSelfServiceBrowserLoginRequestInternalServerError with default headers values
func NewGetSelfServiceBrowserLoginRequestInternalServerError() *GetSelfServiceBrowserLoginRequestInternalServerError {
	return &GetSelfServiceBrowserLoginRequestInternalServerError{}
}

/*GetSelfServiceBrowserLoginRequestInternalServerError handles this case with default header values.

genericError
*/
type GetSelfServiceBrowserLoginRequestInternalServerError struct {
	Payload *models.GenericError
}

func (o *GetSelfServiceBrowserLoginRequestInternalServerError) Error() string {
	return fmt.Sprintf("[GET /self-service/browser/flows/requests/login][%d] getSelfServiceBrowserLoginRequestInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSelfServiceBrowserLoginRequestInternalServerError) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *GetSelfServiceBrowserLoginRequestInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
