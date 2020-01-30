// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// RegistrationRequestMethod RegistrationRequestMethod RegistrationRequestMethod registration request method
// swagger:model registrationRequestMethod
type RegistrationRequestMethod struct {

	// config
	Config *RegistrationRequestMethodConfig `json:"config,omitempty"`

	// method
	Method CredentialsType `json:"method,omitempty"`
}

// Validate validates this registration request method
func (m *RegistrationRequestMethod) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMethod(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RegistrationRequestMethod) validateMethod(formats strfmt.Registry) error {

	if swag.IsZero(m.Method) { // not required
		return nil
	}

	if err := m.Method.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("method")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RegistrationRequestMethod) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RegistrationRequestMethod) UnmarshalBinary(b []byte) error {
	var res RegistrationRequestMethod
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
