/*
 * Ory Kratos API
 *
 * Documentation for all public and administrative Ory Kratos APIs. Public and administrative APIs are exposed on different ports. Public APIs can face the public internet without any protection while administrative APIs should never be exposed without prior authorization. To protect the administative API port you should use something like Nginx, Ory Oathkeeper, or any other technology capable of authorizing incoming requests.
 *
 * API version: v0.8.0-alpha.2.pre.6
 * Contact: hi@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// SelfServiceRecoveryFlowState The state represents the state of the recovery flow.  choose_method: ask the user to choose a method (e.g. recover account via email) sent_email: the email has been sent to the user passed_challenge: the request was successful and the recovery challenge was passed.
type SelfServiceRecoveryFlowState string

// List of selfServiceRecoveryFlowState
const (
	SELFSERVICERECOVERYFLOWSTATE_CHOOSE_METHOD    SelfServiceRecoveryFlowState = "choose_method"
	SELFSERVICERECOVERYFLOWSTATE_SENT_EMAIL       SelfServiceRecoveryFlowState = "sent_email"
	SELFSERVICERECOVERYFLOWSTATE_PASSED_CHALLENGE SelfServiceRecoveryFlowState = "passed_challenge"
)

func (v *SelfServiceRecoveryFlowState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SelfServiceRecoveryFlowState(value)
	for _, existing := range []SelfServiceRecoveryFlowState{"choose_method", "sent_email", "passed_challenge"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SelfServiceRecoveryFlowState", value)
}

// Ptr returns reference to selfServiceRecoveryFlowState value
func (v SelfServiceRecoveryFlowState) Ptr() *SelfServiceRecoveryFlowState {
	return &v
}

type NullableSelfServiceRecoveryFlowState struct {
	value *SelfServiceRecoveryFlowState
	isSet bool
}

func (v NullableSelfServiceRecoveryFlowState) Get() *SelfServiceRecoveryFlowState {
	return v.value
}

func (v *NullableSelfServiceRecoveryFlowState) Set(val *SelfServiceRecoveryFlowState) {
	v.value = val
	v.isSet = true
}

func (v NullableSelfServiceRecoveryFlowState) IsSet() bool {
	return v.isSet
}

func (v *NullableSelfServiceRecoveryFlowState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSelfServiceRecoveryFlowState(val *SelfServiceRecoveryFlowState) *NullableSelfServiceRecoveryFlowState {
	return &NullableSelfServiceRecoveryFlowState{value: val, isSet: true}
}

func (v NullableSelfServiceRecoveryFlowState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSelfServiceRecoveryFlowState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
