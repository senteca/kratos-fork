// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ory/kratos/internal/httpclient/models"
)

// ListIdentitiesReader is a Reader for the ListIdentities structure.
type ListIdentitiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListIdentitiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListIdentitiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewListIdentitiesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListIdentitiesOK creates a ListIdentitiesOK with default headers values
func NewListIdentitiesOK() *ListIdentitiesOK {
	return &ListIdentitiesOK{}
}

/*ListIdentitiesOK handles this case with default header values.

A list of identities.

nolint:deadcode,unused
*/
type ListIdentitiesOK struct {
	Payload []*models.Identity
}

func (o *ListIdentitiesOK) Error() string {
	return fmt.Sprintf("[GET /identities][%d] listIdentitiesOK  %+v", 200, o.Payload)
}

func (o *ListIdentitiesOK) GetPayload() []*models.Identity {
	return o.Payload
}

func (o *ListIdentitiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListIdentitiesInternalServerError creates a ListIdentitiesInternalServerError with default headers values
func NewListIdentitiesInternalServerError() *ListIdentitiesInternalServerError {
	return &ListIdentitiesInternalServerError{}
}

/*ListIdentitiesInternalServerError handles this case with default header values.

genericError
*/
type ListIdentitiesInternalServerError struct {
	Payload *models.GenericError
}

func (o *ListIdentitiesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /identities][%d] listIdentitiesInternalServerError  %+v", 500, o.Payload)
}

func (o *ListIdentitiesInternalServerError) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *ListIdentitiesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
