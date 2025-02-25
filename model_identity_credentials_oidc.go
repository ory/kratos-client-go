/*
Ory Identities API

This is the API specification for Ory Identities with features such as registration, login, recovery, account verification, profile settings, password reset, identity management, session management, email and sms delivery, and more. 

API version: v1.3.5
Contact: office@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the IdentityCredentialsOidc type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdentityCredentialsOidc{}

// IdentityCredentialsOidc struct for IdentityCredentialsOidc
type IdentityCredentialsOidc struct {
	Providers []IdentityCredentialsOidcProvider `json:"providers,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IdentityCredentialsOidc IdentityCredentialsOidc

// NewIdentityCredentialsOidc instantiates a new IdentityCredentialsOidc object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityCredentialsOidc() *IdentityCredentialsOidc {
	this := IdentityCredentialsOidc{}
	return &this
}

// NewIdentityCredentialsOidcWithDefaults instantiates a new IdentityCredentialsOidc object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityCredentialsOidcWithDefaults() *IdentityCredentialsOidc {
	this := IdentityCredentialsOidc{}
	return &this
}

// GetProviders returns the Providers field value if set, zero value otherwise.
func (o *IdentityCredentialsOidc) GetProviders() []IdentityCredentialsOidcProvider {
	if o == nil || IsNil(o.Providers) {
		var ret []IdentityCredentialsOidcProvider
		return ret
	}
	return o.Providers
}

// GetProvidersOk returns a tuple with the Providers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityCredentialsOidc) GetProvidersOk() ([]IdentityCredentialsOidcProvider, bool) {
	if o == nil || IsNil(o.Providers) {
		return nil, false
	}
	return o.Providers, true
}

// HasProviders returns a boolean if a field has been set.
func (o *IdentityCredentialsOidc) HasProviders() bool {
	if o != nil && !IsNil(o.Providers) {
		return true
	}

	return false
}

// SetProviders gets a reference to the given []IdentityCredentialsOidcProvider and assigns it to the Providers field.
func (o *IdentityCredentialsOidc) SetProviders(v []IdentityCredentialsOidcProvider) {
	o.Providers = v
}

func (o IdentityCredentialsOidc) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdentityCredentialsOidc) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Providers) {
		toSerialize["providers"] = o.Providers
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IdentityCredentialsOidc) UnmarshalJSON(data []byte) (err error) {
	varIdentityCredentialsOidc := _IdentityCredentialsOidc{}

	err = json.Unmarshal(data, &varIdentityCredentialsOidc)

	if err != nil {
		return err
	}

	*o = IdentityCredentialsOidc(varIdentityCredentialsOidc)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "providers")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIdentityCredentialsOidc struct {
	value *IdentityCredentialsOidc
	isSet bool
}

func (v NullableIdentityCredentialsOidc) Get() *IdentityCredentialsOidc {
	return v.value
}

func (v *NullableIdentityCredentialsOidc) Set(val *IdentityCredentialsOidc) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityCredentialsOidc) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityCredentialsOidc) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityCredentialsOidc(val *IdentityCredentialsOidc) *NullableIdentityCredentialsOidc {
	return &NullableIdentityCredentialsOidc{value: val, isSet: true}
}

func (v NullableIdentityCredentialsOidc) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityCredentialsOidc) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


