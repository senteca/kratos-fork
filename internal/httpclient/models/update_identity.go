// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UpdateIdentity UpdateIdentity UpdateIdentity UpdateIdentity UpdateIdentity UpdateIdentity UpdateIdentity UpdateIdentity UpdateIdentity UpdateIdentity UpdateIdentity UpdateIdentity UpdateIdentity UpdateIdentity UpdateIdentity update identity
//
// swagger:model UpdateIdentity
type UpdateIdentity struct {

	// SchemaID is the ID of the JSON Schema to be used for validating the identity's traits. If set
	// will update the Identity's SchemaID.
	SchemaID string `json:"schema_id,omitempty"`

	// Traits represent an identity's traits. The identity is able to create, modify, and delete traits
	// in a self-service manner. The input will always be validated against the JSON Schema defined
	// in `schema_id`.
	// Required: true
	Traits interface{} `json:"traits"`
}

// Validate validates this update identity
func (m *UpdateIdentity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTraits(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateIdentity) validateTraits(formats strfmt.Registry) error {

	if m.Traits == nil {
		return errors.Required("traits", "body", nil)
	}

	return nil
}

// ContextValidate validates this update identity based on context it is used
func (m *UpdateIdentity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpdateIdentity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateIdentity) UnmarshalBinary(b []byte) error {
	var res UpdateIdentity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
