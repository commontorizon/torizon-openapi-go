/*
Torizon OTA

 This API is rate limited and will return the following headers for each API call.    - X-RateLimit-Limit - The total number of requests allowed within a time period   - X-RateLimit-Remaining - The total number of requests still allowed until the end of the rate limiting period   - X-RateLimit-Reset - The number of seconds until the limit is fully reset  In addition, if an API client is rate limited, it will receive a HTTP 420 response with the following header:     - Retry-After - The number of seconds to wait until this request is allowed  

API version: 2.0-Beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ClientSignature type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClientSignature{}

// ClientSignature struct for ClientSignature
type ClientSignature struct {
	Keyid string `json:"keyid"`
	Method SignatureMethod `json:"method"`
	Sig string `json:"sig"`
}

type _ClientSignature ClientSignature

// NewClientSignature instantiates a new ClientSignature object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientSignature(keyid string, method SignatureMethod, sig string) *ClientSignature {
	this := ClientSignature{}
	this.Keyid = keyid
	this.Method = method
	this.Sig = sig
	return &this
}

// NewClientSignatureWithDefaults instantiates a new ClientSignature object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientSignatureWithDefaults() *ClientSignature {
	this := ClientSignature{}
	return &this
}

// GetKeyid returns the Keyid field value
func (o *ClientSignature) GetKeyid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Keyid
}

// GetKeyidOk returns a tuple with the Keyid field value
// and a boolean to check if the value has been set.
func (o *ClientSignature) GetKeyidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Keyid, true
}

// SetKeyid sets field value
func (o *ClientSignature) SetKeyid(v string) {
	o.Keyid = v
}

// GetMethod returns the Method field value
func (o *ClientSignature) GetMethod() SignatureMethod {
	if o == nil {
		var ret SignatureMethod
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *ClientSignature) GetMethodOk() (*SignatureMethod, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *ClientSignature) SetMethod(v SignatureMethod) {
	o.Method = v
}

// GetSig returns the Sig field value
func (o *ClientSignature) GetSig() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sig
}

// GetSigOk returns a tuple with the Sig field value
// and a boolean to check if the value has been set.
func (o *ClientSignature) GetSigOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sig, true
}

// SetSig sets field value
func (o *ClientSignature) SetSig(v string) {
	o.Sig = v
}

func (o ClientSignature) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClientSignature) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["keyid"] = o.Keyid
	toSerialize["method"] = o.Method
	toSerialize["sig"] = o.Sig
	return toSerialize, nil
}

func (o *ClientSignature) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"keyid",
		"method",
		"sig",
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

	varClientSignature := _ClientSignature{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varClientSignature)

	if err != nil {
		return err
	}

	*o = ClientSignature(varClientSignature)

	return err
}

type NullableClientSignature struct {
	value *ClientSignature
	isSet bool
}

func (v NullableClientSignature) Get() *ClientSignature {
	return v.value
}

func (v *NullableClientSignature) Set(val *ClientSignature) {
	v.value = val
	v.isSet = true
}

func (v NullableClientSignature) IsSet() bool {
	return v.isSet
}

func (v *NullableClientSignature) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientSignature(val *ClientSignature) *NullableClientSignature {
	return &NullableClientSignature{value: val, isSet: true}
}

func (v NullableClientSignature) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientSignature) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


