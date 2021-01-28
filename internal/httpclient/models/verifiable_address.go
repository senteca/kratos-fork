// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// VerifiableAddress VerifiableAddress VerifiableAddress VerifiableAddress VerifiableAddress VerifiableAddress VerifiableAddress VerifiableAddress VerifiableAddress VerifiableAddress VerifiableAddress verifiable address
//
// swagger:model VerifiableAddress
type VerifiableAddress struct {

	// id
	// Required: true
	// Format: uuid4
	ID *UUID `json:"id"`

	// status
	// Required: true
	Status *VerifiableAddressStatus `json:"status"`

	// value
	// Required: true
	Value *string `json:"value"`

	// verified
	// Required: true
	Verified *bool `json:"verified"`

	// verified at
	// Format: date-time
	VerifiedAt NullTime `json:"verified_at,omitempty"`

	// via
	// Required: true
	Via *VerifiableAddressType `json:"via"`
}

// Validate validates this verifiable address
func (m *VerifiableAddress) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVerified(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVerifiedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVia(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VerifiableAddress) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if m.ID != nil {
		if err := m.ID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("id")
			}
			return err
		}
	}

	return nil
}

func (m *VerifiableAddress) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *VerifiableAddress) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

func (m *VerifiableAddress) validateVerified(formats strfmt.Registry) error {

	if err := validate.Required("verified", "body", m.Verified); err != nil {
		return err
	}

	return nil
}

func (m *VerifiableAddress) validateVerifiedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.VerifiedAt) { // not required
		return nil
	}

	if err := m.VerifiedAt.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("verified_at")
		}
		return err
	}

	return nil
}

func (m *VerifiableAddress) validateVia(formats strfmt.Registry) error {

	if err := validate.Required("via", "body", m.Via); err != nil {
		return err
	}

	if err := validate.Required("via", "body", m.Via); err != nil {
		return err
	}

	if m.Via != nil {
		if err := m.Via.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("via")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this verifiable address based on the context it is used
func (m *VerifiableAddress) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVerifiedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVia(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VerifiableAddress) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if m.ID != nil {
		if err := m.ID.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("id")
			}
			return err
		}
	}

	return nil
}

func (m *VerifiableAddress) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.Status != nil {
		if err := m.Status.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *VerifiableAddress) contextValidateVerifiedAt(ctx context.Context, formats strfmt.Registry) error {

	if err := m.VerifiedAt.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("verified_at")
		}
		return err
	}

	return nil
}

func (m *VerifiableAddress) contextValidateVia(ctx context.Context, formats strfmt.Registry) error {

	if m.Via != nil {
		if err := m.Via.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("via")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VerifiableAddress) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VerifiableAddress) UnmarshalBinary(b []byte) error {
	var res VerifiableAddress
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
