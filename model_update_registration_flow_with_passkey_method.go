/*
Ory Identities API

This is the API specification for Ory Identities with features such as registration, login, recovery, account verification, profile settings, password reset, identity management, session management, email and sms delivery, and more. 

API version: v1.2.0
Contact: office@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// checks if the UpdateRegistrationFlowWithPasskeyMethod type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateRegistrationFlowWithPasskeyMethod{}

// UpdateRegistrationFlowWithPasskeyMethod Update Registration Flow with Passkey Method
type UpdateRegistrationFlowWithPasskeyMethod struct {
	// CSRFToken is the anti-CSRF token
	CsrfToken *string `json:"csrf_token,omitempty"`
	// Method  Should be set to \"passkey\" when trying to add, update, or remove a Passkey.
	Method string `json:"method"`
	// Register a WebAuthn Security Key  It is expected that the JSON returned by the WebAuthn registration process is included here.
	PasskeyRegister *string `json:"passkey_register,omitempty"`
	// The identity's traits
	Traits map[string]interface{} `json:"traits"`
	// Transient data to pass along to any webhooks
	TransientPayload map[string]interface{} `json:"transient_payload,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateRegistrationFlowWithPasskeyMethod UpdateRegistrationFlowWithPasskeyMethod

// NewUpdateRegistrationFlowWithPasskeyMethod instantiates a new UpdateRegistrationFlowWithPasskeyMethod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateRegistrationFlowWithPasskeyMethod(method string, traits map[string]interface{}) *UpdateRegistrationFlowWithPasskeyMethod {
	this := UpdateRegistrationFlowWithPasskeyMethod{}
	this.Method = method
	this.Traits = traits
	return &this
}

// NewUpdateRegistrationFlowWithPasskeyMethodWithDefaults instantiates a new UpdateRegistrationFlowWithPasskeyMethod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateRegistrationFlowWithPasskeyMethodWithDefaults() *UpdateRegistrationFlowWithPasskeyMethod {
	this := UpdateRegistrationFlowWithPasskeyMethod{}
	return &this
}

// GetCsrfToken returns the CsrfToken field value if set, zero value otherwise.
func (o *UpdateRegistrationFlowWithPasskeyMethod) GetCsrfToken() string {
	if o == nil || IsNil(o.CsrfToken) {
		var ret string
		return ret
	}
	return *o.CsrfToken
}

// GetCsrfTokenOk returns a tuple with the CsrfToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRegistrationFlowWithPasskeyMethod) GetCsrfTokenOk() (*string, bool) {
	if o == nil || IsNil(o.CsrfToken) {
		return nil, false
	}
	return o.CsrfToken, true
}

// HasCsrfToken returns a boolean if a field has been set.
func (o *UpdateRegistrationFlowWithPasskeyMethod) HasCsrfToken() bool {
	if o != nil && !IsNil(o.CsrfToken) {
		return true
	}

	return false
}

// SetCsrfToken gets a reference to the given string and assigns it to the CsrfToken field.
func (o *UpdateRegistrationFlowWithPasskeyMethod) SetCsrfToken(v string) {
	o.CsrfToken = &v
}

// GetMethod returns the Method field value
func (o *UpdateRegistrationFlowWithPasskeyMethod) GetMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *UpdateRegistrationFlowWithPasskeyMethod) GetMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *UpdateRegistrationFlowWithPasskeyMethod) SetMethod(v string) {
	o.Method = v
}

// GetPasskeyRegister returns the PasskeyRegister field value if set, zero value otherwise.
func (o *UpdateRegistrationFlowWithPasskeyMethod) GetPasskeyRegister() string {
	if o == nil || IsNil(o.PasskeyRegister) {
		var ret string
		return ret
	}
	return *o.PasskeyRegister
}

// GetPasskeyRegisterOk returns a tuple with the PasskeyRegister field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRegistrationFlowWithPasskeyMethod) GetPasskeyRegisterOk() (*string, bool) {
	if o == nil || IsNil(o.PasskeyRegister) {
		return nil, false
	}
	return o.PasskeyRegister, true
}

// HasPasskeyRegister returns a boolean if a field has been set.
func (o *UpdateRegistrationFlowWithPasskeyMethod) HasPasskeyRegister() bool {
	if o != nil && !IsNil(o.PasskeyRegister) {
		return true
	}

	return false
}

// SetPasskeyRegister gets a reference to the given string and assigns it to the PasskeyRegister field.
func (o *UpdateRegistrationFlowWithPasskeyMethod) SetPasskeyRegister(v string) {
	o.PasskeyRegister = &v
}

// GetTraits returns the Traits field value
func (o *UpdateRegistrationFlowWithPasskeyMethod) GetTraits() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Traits
}

// GetTraitsOk returns a tuple with the Traits field value
// and a boolean to check if the value has been set.
func (o *UpdateRegistrationFlowWithPasskeyMethod) GetTraitsOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Traits, true
}

// SetTraits sets field value
func (o *UpdateRegistrationFlowWithPasskeyMethod) SetTraits(v map[string]interface{}) {
	o.Traits = v
}

// GetTransientPayload returns the TransientPayload field value if set, zero value otherwise.
func (o *UpdateRegistrationFlowWithPasskeyMethod) GetTransientPayload() map[string]interface{} {
	if o == nil || IsNil(o.TransientPayload) {
		var ret map[string]interface{}
		return ret
	}
	return o.TransientPayload
}

// GetTransientPayloadOk returns a tuple with the TransientPayload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRegistrationFlowWithPasskeyMethod) GetTransientPayloadOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.TransientPayload) {
		return map[string]interface{}{}, false
	}
	return o.TransientPayload, true
}

// HasTransientPayload returns a boolean if a field has been set.
func (o *UpdateRegistrationFlowWithPasskeyMethod) HasTransientPayload() bool {
	if o != nil && !IsNil(o.TransientPayload) {
		return true
	}

	return false
}

// SetTransientPayload gets a reference to the given map[string]interface{} and assigns it to the TransientPayload field.
func (o *UpdateRegistrationFlowWithPasskeyMethod) SetTransientPayload(v map[string]interface{}) {
	o.TransientPayload = v
}

func (o UpdateRegistrationFlowWithPasskeyMethod) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateRegistrationFlowWithPasskeyMethod) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CsrfToken) {
		toSerialize["csrf_token"] = o.CsrfToken
	}
	toSerialize["method"] = o.Method
	if !IsNil(o.PasskeyRegister) {
		toSerialize["passkey_register"] = o.PasskeyRegister
	}
	toSerialize["traits"] = o.Traits
	if !IsNil(o.TransientPayload) {
		toSerialize["transient_payload"] = o.TransientPayload
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateRegistrationFlowWithPasskeyMethod) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"method",
		"traits",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varUpdateRegistrationFlowWithPasskeyMethod := _UpdateRegistrationFlowWithPasskeyMethod{}

	err = json.Unmarshal(data, &varUpdateRegistrationFlowWithPasskeyMethod)

	if err != nil {
		return err
	}

	*o = UpdateRegistrationFlowWithPasskeyMethod(varUpdateRegistrationFlowWithPasskeyMethod)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "csrf_token")
		delete(additionalProperties, "method")
		delete(additionalProperties, "passkey_register")
		delete(additionalProperties, "traits")
		delete(additionalProperties, "transient_payload")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateRegistrationFlowWithPasskeyMethod struct {
	value *UpdateRegistrationFlowWithPasskeyMethod
	isSet bool
}

func (v NullableUpdateRegistrationFlowWithPasskeyMethod) Get() *UpdateRegistrationFlowWithPasskeyMethod {
	return v.value
}

func (v *NullableUpdateRegistrationFlowWithPasskeyMethod) Set(val *UpdateRegistrationFlowWithPasskeyMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateRegistrationFlowWithPasskeyMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateRegistrationFlowWithPasskeyMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateRegistrationFlowWithPasskeyMethod(val *UpdateRegistrationFlowWithPasskeyMethod) *NullableUpdateRegistrationFlowWithPasskeyMethod {
	return &NullableUpdateRegistrationFlowWithPasskeyMethod{value: val, isSet: true}
}

func (v NullableUpdateRegistrationFlowWithPasskeyMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateRegistrationFlowWithPasskeyMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


