/*
Torizon OTA

 This API is rate limited and will return the following headers for each API call.    - X-RateLimit-Limit - The total number of requests allowed within a time period   - X-RateLimit-Remaining - The total number of requests still allowed until the end of the rate limiting period   - X-RateLimit-Reset - The number of seconds until the limit is fully reset  In addition, if an API client is rate limited, it will receive a HTTP 420 response with the following header:     - Retry-After - The number of seconds to wait until this request is allowed  

API version: 2.0-Beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the DeviceInfoBasic type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceInfoBasic{}

// DeviceInfoBasic struct for DeviceInfoBasic
type DeviceInfoBasic struct {
	DeviceUuid string `json:"deviceUuid"`
	DeviceName string `json:"deviceName"`
	DeviceId string `json:"deviceId"`
	LastSeen *time.Time `json:"lastSeen,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
	ActivatedAt *time.Time `json:"activatedAt,omitempty"`
	DeviceStatus DeviceStatus `json:"deviceStatus"`
	Notes *string `json:"notes,omitempty"`
}

type _DeviceInfoBasic DeviceInfoBasic

// NewDeviceInfoBasic instantiates a new DeviceInfoBasic object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceInfoBasic(deviceUuid string, deviceName string, deviceId string, createdAt time.Time, deviceStatus DeviceStatus) *DeviceInfoBasic {
	this := DeviceInfoBasic{}
	this.DeviceUuid = deviceUuid
	this.DeviceName = deviceName
	this.DeviceId = deviceId
	this.CreatedAt = createdAt
	this.DeviceStatus = deviceStatus
	return &this
}

// NewDeviceInfoBasicWithDefaults instantiates a new DeviceInfoBasic object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceInfoBasicWithDefaults() *DeviceInfoBasic {
	this := DeviceInfoBasic{}
	return &this
}

// GetDeviceUuid returns the DeviceUuid field value
func (o *DeviceInfoBasic) GetDeviceUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeviceUuid
}

// GetDeviceUuidOk returns a tuple with the DeviceUuid field value
// and a boolean to check if the value has been set.
func (o *DeviceInfoBasic) GetDeviceUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeviceUuid, true
}

// SetDeviceUuid sets field value
func (o *DeviceInfoBasic) SetDeviceUuid(v string) {
	o.DeviceUuid = v
}

// GetDeviceName returns the DeviceName field value
func (o *DeviceInfoBasic) GetDeviceName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeviceName
}

// GetDeviceNameOk returns a tuple with the DeviceName field value
// and a boolean to check if the value has been set.
func (o *DeviceInfoBasic) GetDeviceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeviceName, true
}

// SetDeviceName sets field value
func (o *DeviceInfoBasic) SetDeviceName(v string) {
	o.DeviceName = v
}

// GetDeviceId returns the DeviceId field value
func (o *DeviceInfoBasic) GetDeviceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value
// and a boolean to check if the value has been set.
func (o *DeviceInfoBasic) GetDeviceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeviceId, true
}

// SetDeviceId sets field value
func (o *DeviceInfoBasic) SetDeviceId(v string) {
	o.DeviceId = v
}

// GetLastSeen returns the LastSeen field value if set, zero value otherwise.
func (o *DeviceInfoBasic) GetLastSeen() time.Time {
	if o == nil || IsNil(o.LastSeen) {
		var ret time.Time
		return ret
	}
	return *o.LastSeen
}

// GetLastSeenOk returns a tuple with the LastSeen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceInfoBasic) GetLastSeenOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastSeen) {
		return nil, false
	}
	return o.LastSeen, true
}

// HasLastSeen returns a boolean if a field has been set.
func (o *DeviceInfoBasic) HasLastSeen() bool {
	if o != nil && !IsNil(o.LastSeen) {
		return true
	}

	return false
}

// SetLastSeen gets a reference to the given time.Time and assigns it to the LastSeen field.
func (o *DeviceInfoBasic) SetLastSeen(v time.Time) {
	o.LastSeen = &v
}

// GetCreatedAt returns the CreatedAt field value
func (o *DeviceInfoBasic) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *DeviceInfoBasic) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *DeviceInfoBasic) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetActivatedAt returns the ActivatedAt field value if set, zero value otherwise.
func (o *DeviceInfoBasic) GetActivatedAt() time.Time {
	if o == nil || IsNil(o.ActivatedAt) {
		var ret time.Time
		return ret
	}
	return *o.ActivatedAt
}

// GetActivatedAtOk returns a tuple with the ActivatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceInfoBasic) GetActivatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ActivatedAt) {
		return nil, false
	}
	return o.ActivatedAt, true
}

// HasActivatedAt returns a boolean if a field has been set.
func (o *DeviceInfoBasic) HasActivatedAt() bool {
	if o != nil && !IsNil(o.ActivatedAt) {
		return true
	}

	return false
}

// SetActivatedAt gets a reference to the given time.Time and assigns it to the ActivatedAt field.
func (o *DeviceInfoBasic) SetActivatedAt(v time.Time) {
	o.ActivatedAt = &v
}

// GetDeviceStatus returns the DeviceStatus field value
func (o *DeviceInfoBasic) GetDeviceStatus() DeviceStatus {
	if o == nil {
		var ret DeviceStatus
		return ret
	}

	return o.DeviceStatus
}

// GetDeviceStatusOk returns a tuple with the DeviceStatus field value
// and a boolean to check if the value has been set.
func (o *DeviceInfoBasic) GetDeviceStatusOk() (*DeviceStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeviceStatus, true
}

// SetDeviceStatus sets field value
func (o *DeviceInfoBasic) SetDeviceStatus(v DeviceStatus) {
	o.DeviceStatus = v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *DeviceInfoBasic) GetNotes() string {
	if o == nil || IsNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceInfoBasic) GetNotesOk() (*string, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *DeviceInfoBasic) HasNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *DeviceInfoBasic) SetNotes(v string) {
	o.Notes = &v
}

func (o DeviceInfoBasic) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceInfoBasic) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["deviceUuid"] = o.DeviceUuid
	toSerialize["deviceName"] = o.DeviceName
	toSerialize["deviceId"] = o.DeviceId
	if !IsNil(o.LastSeen) {
		toSerialize["lastSeen"] = o.LastSeen
	}
	toSerialize["createdAt"] = o.CreatedAt
	if !IsNil(o.ActivatedAt) {
		toSerialize["activatedAt"] = o.ActivatedAt
	}
	toSerialize["deviceStatus"] = o.DeviceStatus
	if !IsNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	return toSerialize, nil
}

func (o *DeviceInfoBasic) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"deviceUuid",
		"deviceName",
		"deviceId",
		"createdAt",
		"deviceStatus",
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

	varDeviceInfoBasic := _DeviceInfoBasic{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDeviceInfoBasic)

	if err != nil {
		return err
	}

	*o = DeviceInfoBasic(varDeviceInfoBasic)

	return err
}

type NullableDeviceInfoBasic struct {
	value *DeviceInfoBasic
	isSet bool
}

func (v NullableDeviceInfoBasic) Get() *DeviceInfoBasic {
	return v.value
}

func (v *NullableDeviceInfoBasic) Set(val *DeviceInfoBasic) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceInfoBasic) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceInfoBasic) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceInfoBasic(val *DeviceInfoBasic) *NullableDeviceInfoBasic {
	return &NullableDeviceInfoBasic{value: val, isSet: true}
}

func (v NullableDeviceInfoBasic) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceInfoBasic) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

