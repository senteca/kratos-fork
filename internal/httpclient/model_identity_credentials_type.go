/*
 * Ory Kratos API
 *
 * Documentation for all public and administrative Ory Kratos APIs. Public and administrative APIs are exposed on different ports. Public APIs can face the public internet without any protection while administrative APIs should never be exposed without prior authorization. To protect the administative API port you should use something like Nginx, Ory Oathkeeper, or any other technology capable of authorizing incoming requests.
 *
 * API version: v0.8.0-alpha.2.pre.1
 * Contact: hi@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// IdentityCredentialsType and so on.
type IdentityCredentialsType string

// List of identityCredentialsType
const (
	IDENTITYCREDENTIALSTYPE_PASSWORD IdentityCredentialsType = "password"
	IDENTITYCREDENTIALSTYPE_TOTP     IdentityCredentialsType = "totp"
	IDENTITYCREDENTIALSTYPE_OIDC     IdentityCredentialsType = "oidc"
)

func (v *IdentityCredentialsType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := IdentityCredentialsType(value)
	for _, existing := range []IdentityCredentialsType{"password", "totp", "oidc"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid IdentityCredentialsType", value)
}

// Ptr returns reference to identityCredentialsType value
func (v IdentityCredentialsType) Ptr() *IdentityCredentialsType {
	return &v
}

type NullableIdentityCredentialsType struct {
	value *IdentityCredentialsType
	isSet bool
}

func (v NullableIdentityCredentialsType) Get() *IdentityCredentialsType {
	return v.value
}

func (v *NullableIdentityCredentialsType) Set(val *IdentityCredentialsType) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityCredentialsType) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityCredentialsType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityCredentialsType(val *IdentityCredentialsType) *NullableIdentityCredentialsType {
	return &NullableIdentityCredentialsType{value: val, isSet: true}
}

func (v NullableIdentityCredentialsType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityCredentialsType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
