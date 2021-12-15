/*
 * Ory Kratos API
 *
 * Documentation for all public and administrative Ory Kratos APIs. Public and administrative APIs are exposed on different ports. Public APIs can face the public internet without any protection while administrative APIs should never be exposed without prior authorization. To protect the administative API port you should use something like Nginx, Ory Oathkeeper, or any other technology capable of authorizing incoming requests.
 *
 * API version: v0.8.0-alpha.2.pre.2
 * Contact: hi@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// SubmitSelfServiceVerificationFlowBody - nolint:deadcode,unused
type SubmitSelfServiceVerificationFlowBody struct {
	SubmitSelfServiceVerificationFlowWithLinkMethodBody *SubmitSelfServiceVerificationFlowWithLinkMethodBody
}

// SubmitSelfServiceVerificationFlowWithLinkMethodBodyAsSubmitSelfServiceVerificationFlowBody is a convenience function that returns SubmitSelfServiceVerificationFlowWithLinkMethodBody wrapped in SubmitSelfServiceVerificationFlowBody
func SubmitSelfServiceVerificationFlowWithLinkMethodBodyAsSubmitSelfServiceVerificationFlowBody(v *SubmitSelfServiceVerificationFlowWithLinkMethodBody) SubmitSelfServiceVerificationFlowBody {
	return SubmitSelfServiceVerificationFlowBody{
		SubmitSelfServiceVerificationFlowWithLinkMethodBody: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *SubmitSelfServiceVerificationFlowBody) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into SubmitSelfServiceVerificationFlowWithLinkMethodBody
	err = newStrictDecoder(data).Decode(&dst.SubmitSelfServiceVerificationFlowWithLinkMethodBody)
	if err == nil {
		jsonSubmitSelfServiceVerificationFlowWithLinkMethodBody, _ := json.Marshal(dst.SubmitSelfServiceVerificationFlowWithLinkMethodBody)
		if string(jsonSubmitSelfServiceVerificationFlowWithLinkMethodBody) == "{}" { // empty struct
			dst.SubmitSelfServiceVerificationFlowWithLinkMethodBody = nil
		} else {
			match++
		}
	} else {
		dst.SubmitSelfServiceVerificationFlowWithLinkMethodBody = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.SubmitSelfServiceVerificationFlowWithLinkMethodBody = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(SubmitSelfServiceVerificationFlowBody)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(SubmitSelfServiceVerificationFlowBody)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SubmitSelfServiceVerificationFlowBody) MarshalJSON() ([]byte, error) {
	if src.SubmitSelfServiceVerificationFlowWithLinkMethodBody != nil {
		return json.Marshal(&src.SubmitSelfServiceVerificationFlowWithLinkMethodBody)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *SubmitSelfServiceVerificationFlowBody) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.SubmitSelfServiceVerificationFlowWithLinkMethodBody != nil {
		return obj.SubmitSelfServiceVerificationFlowWithLinkMethodBody
	}

	// all schemas are nil
	return nil
}

type NullableSubmitSelfServiceVerificationFlowBody struct {
	value *SubmitSelfServiceVerificationFlowBody
	isSet bool
}

func (v NullableSubmitSelfServiceVerificationFlowBody) Get() *SubmitSelfServiceVerificationFlowBody {
	return v.value
}

func (v *NullableSubmitSelfServiceVerificationFlowBody) Set(val *SubmitSelfServiceVerificationFlowBody) {
	v.value = val
	v.isSet = true
}

func (v NullableSubmitSelfServiceVerificationFlowBody) IsSet() bool {
	return v.isSet
}

func (v *NullableSubmitSelfServiceVerificationFlowBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubmitSelfServiceVerificationFlowBody(val *SubmitSelfServiceVerificationFlowBody) *NullableSubmitSelfServiceVerificationFlowBody {
	return &NullableSubmitSelfServiceVerificationFlowBody{value: val, isSet: true}
}

func (v NullableSubmitSelfServiceVerificationFlowBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubmitSelfServiceVerificationFlowBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
