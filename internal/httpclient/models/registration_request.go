// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RegistrationRequest RegistrationRequest registration request
// swagger:model registrationRequest
type RegistrationRequest struct {

	// active
	Active CredentialsType `json:"active,omitempty"`

	// ExpiresAt is the time (UTC) when the request expires. If the user still wishes to log in,
	// a new request has to be initiated.
	// Format: date-time
	// Format: date-time
	ExpiresAt strfmt.DateTime `json:"expires_at,omitempty"`

	// id
	// Format: uuid4
	ID UUID `json:"id,omitempty"`

	// IssuedAt is the time (UTC) when the request occurred.
	// Format: date-time
	// Format: date-time
	IssuedAt strfmt.DateTime `json:"issued_at,omitempty"`

	// Methods contains context for all enabled registration methods. If a registration request has been
	// processed, but for example the password is incorrect, this will contain error messages.
	Methods map[string]RegistrationRequestMethod `json:"methods,omitempty"`

	// RequestURL is the initial URL that was requested from ORY Kratos. It can be used
	// to forward information contained in the URL's path or query for example.
	RequestURL string `json:"request_url,omitempty"`
}

// Validate validates this registration request
func (m *RegistrationRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActive(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpiresAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIssuedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMethods(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RegistrationRequest) validateActive(formats strfmt.Registry) error {

	if swag.IsZero(m.Active) { // not required
		return nil
	}

	if err := m.Active.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("active")
		}
		return err
	}

	return nil
}

func (m *RegistrationRequest) validateExpiresAt(formats strfmt.Registry) error {

	if swag.IsZero(m.ExpiresAt) { // not required
		return nil
	}

	if err := validate.FormatOf("expires_at", "body", "date-time", m.ExpiresAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RegistrationRequest) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := m.ID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("id")
		}
		return err
	}

	return nil
}

func (m *RegistrationRequest) validateIssuedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.IssuedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("issued_at", "body", "date-time", m.IssuedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RegistrationRequest) validateMethods(formats strfmt.Registry) error {

	if swag.IsZero(m.Methods) { // not required
		return nil
	}

	for k := range m.Methods {

		if err := validate.Required("methods"+"."+k, "body", m.Methods[k]); err != nil {
			return err
		}
		if val, ok := m.Methods[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RegistrationRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RegistrationRequest) UnmarshalBinary(b []byte) error {
	var res RegistrationRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
