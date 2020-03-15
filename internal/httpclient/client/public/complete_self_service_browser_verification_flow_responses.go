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

// CompleteSelfServiceBrowserVerificationFlowReader is a Reader for the CompleteSelfServiceBrowserVerificationFlow structure.
type CompleteSelfServiceBrowserVerificationFlowReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CompleteSelfServiceBrowserVerificationFlowReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 302:
		result := NewCompleteSelfServiceBrowserVerificationFlowFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCompleteSelfServiceBrowserVerificationFlowInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCompleteSelfServiceBrowserVerificationFlowFound creates a CompleteSelfServiceBrowserVerificationFlowFound with default headers values
func NewCompleteSelfServiceBrowserVerificationFlowFound() *CompleteSelfServiceBrowserVerificationFlowFound {
	return &CompleteSelfServiceBrowserVerificationFlowFound{}
}

/*CompleteSelfServiceBrowserVerificationFlowFound handles this case with default header values.

Empty responses are sent when, for example, resources are deleted. The HTTP status code for empty responses is
typically 201.
*/
type CompleteSelfServiceBrowserVerificationFlowFound struct {
}

func (o *CompleteSelfServiceBrowserVerificationFlowFound) Error() string {
	return fmt.Sprintf("[POST /self-service/browser/flows/verification/{via}/complete][%d] completeSelfServiceBrowserVerificationFlowFound ", 302)
}

func (o *CompleteSelfServiceBrowserVerificationFlowFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCompleteSelfServiceBrowserVerificationFlowInternalServerError creates a CompleteSelfServiceBrowserVerificationFlowInternalServerError with default headers values
func NewCompleteSelfServiceBrowserVerificationFlowInternalServerError() *CompleteSelfServiceBrowserVerificationFlowInternalServerError {
	return &CompleteSelfServiceBrowserVerificationFlowInternalServerError{}
}

/*CompleteSelfServiceBrowserVerificationFlowInternalServerError handles this case with default header values.

genericError
*/
type CompleteSelfServiceBrowserVerificationFlowInternalServerError struct {
	Payload *models.GenericError
}

func (o *CompleteSelfServiceBrowserVerificationFlowInternalServerError) Error() string {
	return fmt.Sprintf("[POST /self-service/browser/flows/verification/{via}/complete][%d] completeSelfServiceBrowserVerificationFlowInternalServerError  %+v", 500, o.Payload)
}

func (o *CompleteSelfServiceBrowserVerificationFlowInternalServerError) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *CompleteSelfServiceBrowserVerificationFlowInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
