// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CompleteSelfServiceRecoveryFlowWithLinkMethod CompleteSelfServiceRecoveryFlowWithLinkMethod CompleteSelfServiceRecoveryFlowWithLinkMethod CompleteSelfServiceRecoveryFlowWithLinkMethod CompleteSelfServiceRecoveryFlowWithLinkMethod CompleteSelfServiceRecoveryFlowWithLinkMethod CompleteSelfServiceRecoveryFlowWithLinkMethod CompleteSelfServiceRecoveryFlowWithLinkMethod CompleteSelfServiceRecoveryFlowWithLinkMethod CompleteSelfServiceRecoveryFlowWithLinkMethod CompleteSelfServiceRecoveryFlowWithLinkMethod CompleteSelfServiceRecoveryFlowWithLinkMethod CompleteSelfServiceRecoveryFlowWithLinkMethod complete self service recovery flow with link method
//
// swagger:model completeSelfServiceRecoveryFlowWithLinkMethod
type CompleteSelfServiceRecoveryFlowWithLinkMethod struct {

	// Sending the anti-csrf token is only required for browser login flows.
	CsrfToken string `json:"csrf_token,omitempty"`

	// Email to Recover
	//
	// Needs to be set when initiating the flow. If the email is a registered
	// recovery email, a recovery link will be sent. If the email is not known,
	// a email with details on what happened will be sent instead.
	//
	// format: email
	// in: body
	Email string `json:"email,omitempty"`
}

// Validate validates this complete self service recovery flow with link method
func (m *CompleteSelfServiceRecoveryFlowWithLinkMethod) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CompleteSelfServiceRecoveryFlowWithLinkMethod) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CompleteSelfServiceRecoveryFlowWithLinkMethod) UnmarshalBinary(b []byte) error {
	var res CompleteSelfServiceRecoveryFlowWithLinkMethod
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
