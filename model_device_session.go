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

// checks if the DeviceSession type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceSession{}

// DeviceSession struct for DeviceSession
type DeviceSession struct {
	Ssh SshSession `json:"ssh"`
}

type _DeviceSession DeviceSession

// NewDeviceSession instantiates a new DeviceSession object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceSession(ssh SshSession) *DeviceSession {
	this := DeviceSession{}
	this.Ssh = ssh
	return &this
}

// NewDeviceSessionWithDefaults instantiates a new DeviceSession object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceSessionWithDefaults() *DeviceSession {
	this := DeviceSession{}
	return &this
}

// GetSsh returns the Ssh field value
func (o *DeviceSession) GetSsh() SshSession {
	if o == nil {
		var ret SshSession
		return ret
	}

	return o.Ssh
}

// GetSshOk returns a tuple with the Ssh field value
// and a boolean to check if the value has been set.
func (o *DeviceSession) GetSshOk() (*SshSession, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ssh, true
}

// SetSsh sets field value
func (o *DeviceSession) SetSsh(v SshSession) {
	o.Ssh = v
}

func (o DeviceSession) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceSession) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ssh"] = o.Ssh
	return toSerialize, nil
}

func (o *DeviceSession) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ssh",
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

	varDeviceSession := _DeviceSession{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDeviceSession)

	if err != nil {
		return err
	}

	*o = DeviceSession(varDeviceSession)

	return err
}

type NullableDeviceSession struct {
	value *DeviceSession
	isSet bool
}

func (v NullableDeviceSession) Get() *DeviceSession {
	return v.value
}

func (v *NullableDeviceSession) Set(val *DeviceSession) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceSession) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceSession) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceSession(val *DeviceSession) *NullableDeviceSession {
	return &NullableDeviceSession{value: val, isSet: true}
}

func (v NullableDeviceSession) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceSession) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


