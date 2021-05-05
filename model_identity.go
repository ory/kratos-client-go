/*
 * Ory Kratos API
 *
 * Documentation for all public and administrative Ory Kratos APIs. Public and administrative APIs are exposed on different ports. Public APIs can face the public internet without any protection while administrative APIs should never be exposed without prior authorization. To protect the administative API port you should use something like Nginx, Ory Oathkeeper, or any other technology capable of authorizing incoming requests. 
 *
 * API version: v0.6.0-alpha.1
 * Contact: hi@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// Identity struct for Identity
type Identity struct {
	Id string `json:"id"`
	// RecoveryAddresses contains all the addresses that can be used to recover an identity.
	RecoveryAddresses []RecoveryAddress `json:"recovery_addresses,omitempty"`
	// SchemaID is the ID of the JSON Schema to be used for validating the identity's traits.
	SchemaId string `json:"schema_id"`
	// SchemaURL is the URL of the endpoint where the identity's traits schema can be fetched from.  format: url
	SchemaUrl string `json:"schema_url"`
	Traits map[string]interface{} `json:"traits"`
	// VerifiableAddresses contains all the addresses that can be verified by the user.
	VerifiableAddresses []VerifiableAddress `json:"verifiable_addresses,omitempty"`
}

// NewIdentity instantiates a new Identity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentity(id string, schemaId string, schemaUrl string, traits map[string]interface{}) *Identity {
	this := Identity{}
	this.Id = id
	this.SchemaId = schemaId
	this.SchemaUrl = schemaUrl
	this.Traits = traits
	return &this
}

// NewIdentityWithDefaults instantiates a new Identity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityWithDefaults() *Identity {
	this := Identity{}
	return &this
}

// GetId returns the Id field value
func (o *Identity) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Identity) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Identity) SetId(v string) {
	o.Id = v
}

// GetRecoveryAddresses returns the RecoveryAddresses field value if set, zero value otherwise.
func (o *Identity) GetRecoveryAddresses() []RecoveryAddress {
	if o == nil || o.RecoveryAddresses == nil {
		var ret []RecoveryAddress
		return ret
	}
	return o.RecoveryAddresses
}

// GetRecoveryAddressesOk returns a tuple with the RecoveryAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Identity) GetRecoveryAddressesOk() ([]RecoveryAddress, bool) {
	if o == nil || o.RecoveryAddresses == nil {
		return nil, false
	}
	return o.RecoveryAddresses, true
}

// HasRecoveryAddresses returns a boolean if a field has been set.
func (o *Identity) HasRecoveryAddresses() bool {
	if o != nil && o.RecoveryAddresses != nil {
		return true
	}

	return false
}

// SetRecoveryAddresses gets a reference to the given []RecoveryAddress and assigns it to the RecoveryAddresses field.
func (o *Identity) SetRecoveryAddresses(v []RecoveryAddress) {
	o.RecoveryAddresses = v
}

// GetSchemaId returns the SchemaId field value
func (o *Identity) GetSchemaId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SchemaId
}

// GetSchemaIdOk returns a tuple with the SchemaId field value
// and a boolean to check if the value has been set.
func (o *Identity) GetSchemaIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SchemaId, true
}

// SetSchemaId sets field value
func (o *Identity) SetSchemaId(v string) {
	o.SchemaId = v
}

// GetSchemaUrl returns the SchemaUrl field value
func (o *Identity) GetSchemaUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SchemaUrl
}

// GetSchemaUrlOk returns a tuple with the SchemaUrl field value
// and a boolean to check if the value has been set.
func (o *Identity) GetSchemaUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SchemaUrl, true
}

// SetSchemaUrl sets field value
func (o *Identity) SetSchemaUrl(v string) {
	o.SchemaUrl = v
}

// GetTraits returns the Traits field value
func (o *Identity) GetTraits() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Traits
}

// GetTraitsOk returns a tuple with the Traits field value
// and a boolean to check if the value has been set.
func (o *Identity) GetTraitsOk() (map[string]interface{}, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Traits, true
}

// SetTraits sets field value
func (o *Identity) SetTraits(v map[string]interface{}) {
	o.Traits = v
}

// GetVerifiableAddresses returns the VerifiableAddresses field value if set, zero value otherwise.
func (o *Identity) GetVerifiableAddresses() []VerifiableAddress {
	if o == nil || o.VerifiableAddresses == nil {
		var ret []VerifiableAddress
		return ret
	}
	return o.VerifiableAddresses
}

// GetVerifiableAddressesOk returns a tuple with the VerifiableAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Identity) GetVerifiableAddressesOk() ([]VerifiableAddress, bool) {
	if o == nil || o.VerifiableAddresses == nil {
		return nil, false
	}
	return o.VerifiableAddresses, true
}

// HasVerifiableAddresses returns a boolean if a field has been set.
func (o *Identity) HasVerifiableAddresses() bool {
	if o != nil && o.VerifiableAddresses != nil {
		return true
	}

	return false
}

// SetVerifiableAddresses gets a reference to the given []VerifiableAddress and assigns it to the VerifiableAddresses field.
func (o *Identity) SetVerifiableAddresses(v []VerifiableAddress) {
	o.VerifiableAddresses = v
}

func (o Identity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.RecoveryAddresses != nil {
		toSerialize["recovery_addresses"] = o.RecoveryAddresses
	}
	if true {
		toSerialize["schema_id"] = o.SchemaId
	}
	if true {
		toSerialize["schema_url"] = o.SchemaUrl
	}
	if true {
		toSerialize["traits"] = o.Traits
	}
	if o.VerifiableAddresses != nil {
		toSerialize["verifiable_addresses"] = o.VerifiableAddresses
	}
	return json.Marshal(toSerialize)
}

type NullableIdentity struct {
	value *Identity
	isSet bool
}

func (v NullableIdentity) Get() *Identity {
	return v.value
}

func (v *NullableIdentity) Set(val *Identity) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentity) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentity(val *Identity) *NullableIdentity {
	return &NullableIdentity{value: val, isSet: true}
}

func (v NullableIdentity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


