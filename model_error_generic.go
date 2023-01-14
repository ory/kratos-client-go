/*
Ory Identities API

This is the API specification for Ory Identities with features such as registration, login, recovery, account verification, profile settings, password reset, identity management, session management, email and sms delivery, and more. 

API version: v0.11.1
Contact: office@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// ErrorGeneric The standard Ory JSON API error format.
type ErrorGeneric struct {
	Error GenericError `json:"error"`
}

// NewErrorGeneric instantiates a new ErrorGeneric object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorGeneric(error_ GenericError) *ErrorGeneric {
	this := ErrorGeneric{}
	this.Error = error_
	return &this
}

// NewErrorGenericWithDefaults instantiates a new ErrorGeneric object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorGenericWithDefaults() *ErrorGeneric {
	this := ErrorGeneric{}
	return &this
}

// GetError returns the Error field value
func (o *ErrorGeneric) GetError() GenericError {
	if o == nil {
		var ret GenericError
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *ErrorGeneric) GetErrorOk() (*GenericError, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value
func (o *ErrorGeneric) SetError(v GenericError) {
	o.Error = v
}

func (o ErrorGeneric) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableErrorGeneric struct {
	value *ErrorGeneric
	isSet bool
}

func (v NullableErrorGeneric) Get() *ErrorGeneric {
	return v.value
}

func (v *NullableErrorGeneric) Set(val *ErrorGeneric) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorGeneric) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorGeneric) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorGeneric(val *ErrorGeneric) *NullableErrorGeneric {
	return &NullableErrorGeneric{value: val, isSet: true}
}

func (v NullableErrorGeneric) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorGeneric) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


