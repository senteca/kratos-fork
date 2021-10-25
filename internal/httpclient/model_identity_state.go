/*
 * Ory Kratos API
 *
 * Documentation for all public and administrative Ory Kratos APIs. Public and administrative APIs are exposed on different ports. Public APIs can face the public internet without any protection while administrative APIs should never be exposed without prior authorization. To protect the administative API port you should use something like Nginx, Ory Oathkeeper, or any other technology capable of authorizing incoming requests.
 *
 * API version: v0.8.0-alpha.1.pre.1
 * Contact: hi@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// IdentityState The state can either be `active` or `inactive`.
type IdentityState string

// List of identityState
const (
	IDENTITYSTATE_ACTIVE   IdentityState = "active"
	IDENTITYSTATE_INACTIVE IdentityState = "inactive"
)

func (v *IdentityState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := IdentityState(value)
	for _, existing := range []IdentityState{"active", "inactive"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid IdentityState", value)
}

// Ptr returns reference to identityState value
func (v IdentityState) Ptr() *IdentityState {
	return &v
}

type NullableIdentityState struct {
	value *IdentityState
	isSet bool
}

func (v NullableIdentityState) Get() *IdentityState {
	return v.value
}

func (v *NullableIdentityState) Set(val *IdentityState) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityState) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityState(val *IdentityState) *NullableIdentityState {
	return &NullableIdentityState{value: val, isSet: true}
}

func (v NullableIdentityState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
