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

// GetSelfServiceRecoveryFlowReader is a Reader for the GetSelfServiceRecoveryFlow structure.
type GetSelfServiceRecoveryFlowReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSelfServiceRecoveryFlowReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSelfServiceRecoveryFlowOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetSelfServiceRecoveryFlowNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewGetSelfServiceRecoveryFlowGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSelfServiceRecoveryFlowInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSelfServiceRecoveryFlowOK creates a GetSelfServiceRecoveryFlowOK with default headers values
func NewGetSelfServiceRecoveryFlowOK() *GetSelfServiceRecoveryFlowOK {
	return &GetSelfServiceRecoveryFlowOK{}
}

/*GetSelfServiceRecoveryFlowOK handles this case with default header values.

recoveryFlow
*/
type GetSelfServiceRecoveryFlowOK struct {
	Payload *models.RecoveryFlow
}

func (o *GetSelfServiceRecoveryFlowOK) Error() string {
	return fmt.Sprintf("[GET /self-service/recovery/flows][%d] getSelfServiceRecoveryFlowOK  %+v", 200, o.Payload)
}

func (o *GetSelfServiceRecoveryFlowOK) GetPayload() *models.RecoveryFlow {
	return o.Payload
}

func (o *GetSelfServiceRecoveryFlowOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RecoveryFlow)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSelfServiceRecoveryFlowNotFound creates a GetSelfServiceRecoveryFlowNotFound with default headers values
func NewGetSelfServiceRecoveryFlowNotFound() *GetSelfServiceRecoveryFlowNotFound {
	return &GetSelfServiceRecoveryFlowNotFound{}
}

/*GetSelfServiceRecoveryFlowNotFound handles this case with default header values.

genericError
*/
type GetSelfServiceRecoveryFlowNotFound struct {
	Payload *models.GenericError
}

func (o *GetSelfServiceRecoveryFlowNotFound) Error() string {
	return fmt.Sprintf("[GET /self-service/recovery/flows][%d] getSelfServiceRecoveryFlowNotFound  %+v", 404, o.Payload)
}

func (o *GetSelfServiceRecoveryFlowNotFound) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *GetSelfServiceRecoveryFlowNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSelfServiceRecoveryFlowGone creates a GetSelfServiceRecoveryFlowGone with default headers values
func NewGetSelfServiceRecoveryFlowGone() *GetSelfServiceRecoveryFlowGone {
	return &GetSelfServiceRecoveryFlowGone{}
}

/*GetSelfServiceRecoveryFlowGone handles this case with default header values.

genericError
*/
type GetSelfServiceRecoveryFlowGone struct {
	Payload *models.GenericError
}

func (o *GetSelfServiceRecoveryFlowGone) Error() string {
	return fmt.Sprintf("[GET /self-service/recovery/flows][%d] getSelfServiceRecoveryFlowGone  %+v", 410, o.Payload)
}

func (o *GetSelfServiceRecoveryFlowGone) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *GetSelfServiceRecoveryFlowGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSelfServiceRecoveryFlowInternalServerError creates a GetSelfServiceRecoveryFlowInternalServerError with default headers values
func NewGetSelfServiceRecoveryFlowInternalServerError() *GetSelfServiceRecoveryFlowInternalServerError {
	return &GetSelfServiceRecoveryFlowInternalServerError{}
}

/*GetSelfServiceRecoveryFlowInternalServerError handles this case with default header values.

genericError
*/
type GetSelfServiceRecoveryFlowInternalServerError struct {
	Payload *models.GenericError
}

func (o *GetSelfServiceRecoveryFlowInternalServerError) Error() string {
	return fmt.Sprintf("[GET /self-service/recovery/flows][%d] getSelfServiceRecoveryFlowInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSelfServiceRecoveryFlowInternalServerError) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *GetSelfServiceRecoveryFlowInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
