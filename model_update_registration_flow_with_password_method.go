/*
Ory Identities API

This is the API specification for Ory Identities with features such as registration, login, recovery, account verification, profile settings, password reset, identity management, session management, email and sms delivery, and more. 

API version: v1.3.4
Contact: office@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// checks if the UpdateRegistrationFlowWithPasswordMethod type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateRegistrationFlowWithPasswordMethod{}

// UpdateRegistrationFlowWithPasswordMethod Update Registration Flow with Password Method
type UpdateRegistrationFlowWithPasswordMethod struct {
	// The CSRF Token
	CsrfToken *string `json:"csrf_token,omitempty"`
	// Method to use  This field must be set to `password` when using the password method.
	Method string `json:"method"`
	// Password to sign the user up with
	Password string `json:"password"`
	// The identity's traits
	Traits map[string]interface{} `json:"traits"`
	// Transient data to pass along to any webhooks
	TransientPayload map[string]interface{} `json:"transient_payload,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateRegistrationFlowWithPasswordMethod UpdateRegistrationFlowWithPasswordMethod

// NewUpdateRegistrationFlowWithPasswordMethod instantiates a new UpdateRegistrationFlowWithPasswordMethod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateRegistrationFlowWithPasswordMethod(method string, password string, traits map[string]interface{}) *UpdateRegistrationFlowWithPasswordMethod {
	this := UpdateRegistrationFlowWithPasswordMethod{}
	this.Method = method
	this.Password = password
	this.Traits = traits
	return &this
}

// NewUpdateRegistrationFlowWithPasswordMethodWithDefaults instantiates a new UpdateRegistrationFlowWithPasswordMethod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateRegistrationFlowWithPasswordMethodWithDefaults() *UpdateRegistrationFlowWithPasswordMethod {
	this := UpdateRegistrationFlowWithPasswordMethod{}
	return &this
}

// GetCsrfToken returns the CsrfToken field value if set, zero value otherwise.
func (o *UpdateRegistrationFlowWithPasswordMethod) GetCsrfToken() string {
	if o == nil || IsNil(o.CsrfToken) {
		var ret string
		return ret
	}
	return *o.CsrfToken
}

// GetCsrfTokenOk returns a tuple with the CsrfToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRegistrationFlowWithPasswordMethod) GetCsrfTokenOk() (*string, bool) {
	if o == nil || IsNil(o.CsrfToken) {
		return nil, false
	}
	return o.CsrfToken, true
}

// HasCsrfToken returns a boolean if a field has been set.
func (o *UpdateRegistrationFlowWithPasswordMethod) HasCsrfToken() bool {
	if o != nil && !IsNil(o.CsrfToken) {
		return true
	}

	return false
}

// SetCsrfToken gets a reference to the given string and assigns it to the CsrfToken field.
func (o *UpdateRegistrationFlowWithPasswordMethod) SetCsrfToken(v string) {
	o.CsrfToken = &v
}

// GetMethod returns the Method field value
func (o *UpdateRegistrationFlowWithPasswordMethod) GetMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *UpdateRegistrationFlowWithPasswordMethod) GetMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *UpdateRegistrationFlowWithPasswordMethod) SetMethod(v string) {
	o.Method = v
}

// GetPassword returns the Password field value
func (o *UpdateRegistrationFlowWithPasswordMethod) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *UpdateRegistrationFlowWithPasswordMethod) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *UpdateRegistrationFlowWithPasswordMethod) SetPassword(v string) {
	o.Password = v
}

// GetTraits returns the Traits field value
func (o *UpdateRegistrationFlowWithPasswordMethod) GetTraits() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Traits
}

// GetTraitsOk returns a tuple with the Traits field value
// and a boolean to check if the value has been set.
func (o *UpdateRegistrationFlowWithPasswordMethod) GetTraitsOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Traits, true
}

// SetTraits sets field value
func (o *UpdateRegistrationFlowWithPasswordMethod) SetTraits(v map[string]interface{}) {
	o.Traits = v
}

// GetTransientPayload returns the TransientPayload field value if set, zero value otherwise.
func (o *UpdateRegistrationFlowWithPasswordMethod) GetTransientPayload() map[string]interface{} {
	if o == nil || IsNil(o.TransientPayload) {
		var ret map[string]interface{}
		return ret
	}
	return o.TransientPayload
}

// GetTransientPayloadOk returns a tuple with the TransientPayload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRegistrationFlowWithPasswordMethod) GetTransientPayloadOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.TransientPayload) {
		return map[string]interface{}{}, false
	}
	return o.TransientPayload, true
}

// HasTransientPayload returns a boolean if a field has been set.
func (o *UpdateRegistrationFlowWithPasswordMethod) HasTransientPayload() bool {
	if o != nil && !IsNil(o.TransientPayload) {
		return true
	}

	return false
}

// SetTransientPayload gets a reference to the given map[string]interface{} and assigns it to the TransientPayload field.
func (o *UpdateRegistrationFlowWithPasswordMethod) SetTransientPayload(v map[string]interface{}) {
	o.TransientPayload = v
}

func (o UpdateRegistrationFlowWithPasswordMethod) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateRegistrationFlowWithPasswordMethod) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CsrfToken) {
		toSerialize["csrf_token"] = o.CsrfToken
	}
	toSerialize["method"] = o.Method
	toSerialize["password"] = o.Password
	toSerialize["traits"] = o.Traits
	if !IsNil(o.TransientPayload) {
		toSerialize["transient_payload"] = o.TransientPayload
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateRegistrationFlowWithPasswordMethod) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"method",
		"password",
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

	varUpdateRegistrationFlowWithPasswordMethod := _UpdateRegistrationFlowWithPasswordMethod{}

	err = json.Unmarshal(data, &varUpdateRegistrationFlowWithPasswordMethod)

	if err != nil {
		return err
	}

	*o = UpdateRegistrationFlowWithPasswordMethod(varUpdateRegistrationFlowWithPasswordMethod)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "csrf_token")
		delete(additionalProperties, "method")
		delete(additionalProperties, "password")
		delete(additionalProperties, "traits")
		delete(additionalProperties, "transient_payload")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateRegistrationFlowWithPasswordMethod struct {
	value *UpdateRegistrationFlowWithPasswordMethod
	isSet bool
}

func (v NullableUpdateRegistrationFlowWithPasswordMethod) Get() *UpdateRegistrationFlowWithPasswordMethod {
	return v.value
}

func (v *NullableUpdateRegistrationFlowWithPasswordMethod) Set(val *UpdateRegistrationFlowWithPasswordMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateRegistrationFlowWithPasswordMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateRegistrationFlowWithPasswordMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateRegistrationFlowWithPasswordMethod(val *UpdateRegistrationFlowWithPasswordMethod) *NullableUpdateRegistrationFlowWithPasswordMethod {
	return &NullableUpdateRegistrationFlowWithPasswordMethod{value: val, isSet: true}
}

func (v NullableUpdateRegistrationFlowWithPasswordMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateRegistrationFlowWithPasswordMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


