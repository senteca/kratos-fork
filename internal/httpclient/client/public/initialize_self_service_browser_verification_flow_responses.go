// Code generated by go-swagger; DO NOT EDIT.

package public

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ory/kratos/internal/httpclient/models"
)

// InitializeSelfServiceBrowserVerificationFlowReader is a Reader for the InitializeSelfServiceBrowserVerificationFlow structure.
type InitializeSelfServiceBrowserVerificationFlowReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InitializeSelfServiceBrowserVerificationFlowReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 302:
		result := NewInitializeSelfServiceBrowserVerificationFlowFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewInitializeSelfServiceBrowserVerificationFlowInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewInitializeSelfServiceBrowserVerificationFlowFound creates a InitializeSelfServiceBrowserVerificationFlowFound with default headers values
func NewInitializeSelfServiceBrowserVerificationFlowFound() *InitializeSelfServiceBrowserVerificationFlowFound {
	return &InitializeSelfServiceBrowserVerificationFlowFound{}
}

/*InitializeSelfServiceBrowserVerificationFlowFound handles this case with default header values.

Empty responses are sent when, for example, resources are deleted. The HTTP status code for empty responses is
typically 201.
*/
type InitializeSelfServiceBrowserVerificationFlowFound struct {
}

func (o *InitializeSelfServiceBrowserVerificationFlowFound) Error() string {
	return fmt.Sprintf("[GET /self-service/browser/flows/verification/init/{via}][%d] initializeSelfServiceBrowserVerificationFlowFound ", 302)
}

func (o *InitializeSelfServiceBrowserVerificationFlowFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewInitializeSelfServiceBrowserVerificationFlowInternalServerError creates a InitializeSelfServiceBrowserVerificationFlowInternalServerError with default headers values
func NewInitializeSelfServiceBrowserVerificationFlowInternalServerError() *InitializeSelfServiceBrowserVerificationFlowInternalServerError {
	return &InitializeSelfServiceBrowserVerificationFlowInternalServerError{}
}

/*InitializeSelfServiceBrowserVerificationFlowInternalServerError handles this case with default header values.

genericError
*/
type InitializeSelfServiceBrowserVerificationFlowInternalServerError struct {
	Payload *models.GenericError
}

func (o *InitializeSelfServiceBrowserVerificationFlowInternalServerError) Error() string {
	return fmt.Sprintf("[GET /self-service/browser/flows/verification/init/{via}][%d] initializeSelfServiceBrowserVerificationFlowInternalServerError  %+v", 500, o.Payload)
}

func (o *InitializeSelfServiceBrowserVerificationFlowInternalServerError) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *InitializeSelfServiceBrowserVerificationFlowInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
