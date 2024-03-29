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

// checks if the DeviceInfoExtended type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceInfoExtended{}

// DeviceInfoExtended struct for DeviceInfoExtended
type DeviceInfoExtended struct {
	DeviceUuid string `json:"deviceUuid"`
	DeviceName string `json:"deviceName"`
	DeviceId string `json:"deviceId"`
	LastSeen *time.Time `json:"lastSeen,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
	ActivatedAt *time.Time `json:"activatedAt,omitempty"`
	DeviceStatus DeviceStatus `json:"deviceStatus"`
	Notes *string `json:"notes,omitempty"`
	Hibernated bool `json:"hibernated"`
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
	DeviceFleets []Fleet `json:"deviceFleets,omitempty"`
	DevicePackages []DevicePackage `json:"devicePackages,omitempty"`
	DeviceTags []Tuple2DeviceTagIdDeviceTagValue `json:"deviceTags,omitempty"`
	NetworkInfo NetworkInfo `json:"networkInfo"`
}

type _DeviceInfoExtended DeviceInfoExtended

// NewDeviceInfoExtended instantiates a new DeviceInfoExtended object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceInfoExtended(deviceUuid string, deviceName string, deviceId string, createdAt time.Time, deviceStatus DeviceStatus, hibernated bool, networkInfo NetworkInfo) *DeviceInfoExtended {
	this := DeviceInfoExtended{}
	this.DeviceUuid = deviceUuid
	this.DeviceName = deviceName
	this.DeviceId = deviceId
	this.CreatedAt = createdAt
	this.DeviceStatus = deviceStatus
	this.Hibernated = hibernated
	this.NetworkInfo = networkInfo
	return &this
}

// NewDeviceInfoExtendedWithDefaults instantiates a new DeviceInfoExtended object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceInfoExtendedWithDefaults() *DeviceInfoExtended {
	this := DeviceInfoExtended{}
	return &this
}

// GetDeviceUuid returns the DeviceUuid field value
func (o *DeviceInfoExtended) GetDeviceUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeviceUuid
}

// GetDeviceUuidOk returns a tuple with the DeviceUuid field value
// and a boolean to check if the value has been set.
func (o *DeviceInfoExtended) GetDeviceUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeviceUuid, true
}

// SetDeviceUuid sets field value
func (o *DeviceInfoExtended) SetDeviceUuid(v string) {
	o.DeviceUuid = v
}

// GetDeviceName returns the DeviceName field value
func (o *DeviceInfoExtended) GetDeviceName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeviceName
}

// GetDeviceNameOk returns a tuple with the DeviceName field value
// and a boolean to check if the value has been set.
func (o *DeviceInfoExtended) GetDeviceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeviceName, true
}

// SetDeviceName sets field value
func (o *DeviceInfoExtended) SetDeviceName(v string) {
	o.DeviceName = v
}

// GetDeviceId returns the DeviceId field value
func (o *DeviceInfoExtended) GetDeviceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value
// and a boolean to check if the value has been set.
func (o *DeviceInfoExtended) GetDeviceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeviceId, true
}

// SetDeviceId sets field value
func (o *DeviceInfoExtended) SetDeviceId(v string) {
	o.DeviceId = v
}

// GetLastSeen returns the LastSeen field value if set, zero value otherwise.
func (o *DeviceInfoExtended) GetLastSeen() time.Time {
	if o == nil || IsNil(o.LastSeen) {
		var ret time.Time
		return ret
	}
	return *o.LastSeen
}

// GetLastSeenOk returns a tuple with the LastSeen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceInfoExtended) GetLastSeenOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastSeen) {
		return nil, false
	}
	return o.LastSeen, true
}

// HasLastSeen returns a boolean if a field has been set.
func (o *DeviceInfoExtended) HasLastSeen() bool {
	if o != nil && !IsNil(o.LastSeen) {
		return true
	}

	return false
}

// SetLastSeen gets a reference to the given time.Time and assigns it to the LastSeen field.
func (o *DeviceInfoExtended) SetLastSeen(v time.Time) {
	o.LastSeen = &v
}

// GetCreatedAt returns the CreatedAt field value
func (o *DeviceInfoExtended) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *DeviceInfoExtended) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *DeviceInfoExtended) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetActivatedAt returns the ActivatedAt field value if set, zero value otherwise.
func (o *DeviceInfoExtended) GetActivatedAt() time.Time {
	if o == nil || IsNil(o.ActivatedAt) {
		var ret time.Time
		return ret
	}
	return *o.ActivatedAt
}

// GetActivatedAtOk returns a tuple with the ActivatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceInfoExtended) GetActivatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ActivatedAt) {
		return nil, false
	}
	return o.ActivatedAt, true
}

// HasActivatedAt returns a boolean if a field has been set.
func (o *DeviceInfoExtended) HasActivatedAt() bool {
	if o != nil && !IsNil(o.ActivatedAt) {
		return true
	}

	return false
}

// SetActivatedAt gets a reference to the given time.Time and assigns it to the ActivatedAt field.
func (o *DeviceInfoExtended) SetActivatedAt(v time.Time) {
	o.ActivatedAt = &v
}

// GetDeviceStatus returns the DeviceStatus field value
func (o *DeviceInfoExtended) GetDeviceStatus() DeviceStatus {
	if o == nil {
		var ret DeviceStatus
		return ret
	}

	return o.DeviceStatus
}

// GetDeviceStatusOk returns a tuple with the DeviceStatus field value
// and a boolean to check if the value has been set.
func (o *DeviceInfoExtended) GetDeviceStatusOk() (*DeviceStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeviceStatus, true
}

// SetDeviceStatus sets field value
func (o *DeviceInfoExtended) SetDeviceStatus(v DeviceStatus) {
	o.DeviceStatus = v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *DeviceInfoExtended) GetNotes() string {
	if o == nil || IsNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceInfoExtended) GetNotesOk() (*string, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *DeviceInfoExtended) HasNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *DeviceInfoExtended) SetNotes(v string) {
	o.Notes = &v
}

// GetHibernated returns the Hibernated field value
func (o *DeviceInfoExtended) GetHibernated() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Hibernated
}

// GetHibernatedOk returns a tuple with the Hibernated field value
// and a boolean to check if the value has been set.
func (o *DeviceInfoExtended) GetHibernatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hibernated, true
}

// SetHibernated sets field value
func (o *DeviceInfoExtended) SetHibernated(v bool) {
	o.Hibernated = v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *DeviceInfoExtended) GetLastUpdated() time.Time {
	if o == nil || IsNil(o.LastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceInfoExtended) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *DeviceInfoExtended) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *DeviceInfoExtended) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetDeviceFleets returns the DeviceFleets field value if set, zero value otherwise.
func (o *DeviceInfoExtended) GetDeviceFleets() []Fleet {
	if o == nil || IsNil(o.DeviceFleets) {
		var ret []Fleet
		return ret
	}
	return o.DeviceFleets
}

// GetDeviceFleetsOk returns a tuple with the DeviceFleets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceInfoExtended) GetDeviceFleetsOk() ([]Fleet, bool) {
	if o == nil || IsNil(o.DeviceFleets) {
		return nil, false
	}
	return o.DeviceFleets, true
}

// HasDeviceFleets returns a boolean if a field has been set.
func (o *DeviceInfoExtended) HasDeviceFleets() bool {
	if o != nil && !IsNil(o.DeviceFleets) {
		return true
	}

	return false
}

// SetDeviceFleets gets a reference to the given []Fleet and assigns it to the DeviceFleets field.
func (o *DeviceInfoExtended) SetDeviceFleets(v []Fleet) {
	o.DeviceFleets = v
}

// GetDevicePackages returns the DevicePackages field value if set, zero value otherwise.
func (o *DeviceInfoExtended) GetDevicePackages() []DevicePackage {
	if o == nil || IsNil(o.DevicePackages) {
		var ret []DevicePackage
		return ret
	}
	return o.DevicePackages
}

// GetDevicePackagesOk returns a tuple with the DevicePackages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceInfoExtended) GetDevicePackagesOk() ([]DevicePackage, bool) {
	if o == nil || IsNil(o.DevicePackages) {
		return nil, false
	}
	return o.DevicePackages, true
}

// HasDevicePackages returns a boolean if a field has been set.
func (o *DeviceInfoExtended) HasDevicePackages() bool {
	if o != nil && !IsNil(o.DevicePackages) {
		return true
	}

	return false
}

// SetDevicePackages gets a reference to the given []DevicePackage and assigns it to the DevicePackages field.
func (o *DeviceInfoExtended) SetDevicePackages(v []DevicePackage) {
	o.DevicePackages = v
}

// GetDeviceTags returns the DeviceTags field value if set, zero value otherwise.
func (o *DeviceInfoExtended) GetDeviceTags() []Tuple2DeviceTagIdDeviceTagValue {
	if o == nil || IsNil(o.DeviceTags) {
		var ret []Tuple2DeviceTagIdDeviceTagValue
		return ret
	}
	return o.DeviceTags
}

// GetDeviceTagsOk returns a tuple with the DeviceTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceInfoExtended) GetDeviceTagsOk() ([]Tuple2DeviceTagIdDeviceTagValue, bool) {
	if o == nil || IsNil(o.DeviceTags) {
		return nil, false
	}
	return o.DeviceTags, true
}

// HasDeviceTags returns a boolean if a field has been set.
func (o *DeviceInfoExtended) HasDeviceTags() bool {
	if o != nil && !IsNil(o.DeviceTags) {
		return true
	}

	return false
}

// SetDeviceTags gets a reference to the given []Tuple2DeviceTagIdDeviceTagValue and assigns it to the DeviceTags field.
func (o *DeviceInfoExtended) SetDeviceTags(v []Tuple2DeviceTagIdDeviceTagValue) {
	o.DeviceTags = v
}

// GetNetworkInfo returns the NetworkInfo field value
func (o *DeviceInfoExtended) GetNetworkInfo() NetworkInfo {
	if o == nil {
		var ret NetworkInfo
		return ret
	}

	return o.NetworkInfo
}

// GetNetworkInfoOk returns a tuple with the NetworkInfo field value
// and a boolean to check if the value has been set.
func (o *DeviceInfoExtended) GetNetworkInfoOk() (*NetworkInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NetworkInfo, true
}

// SetNetworkInfo sets field value
func (o *DeviceInfoExtended) SetNetworkInfo(v NetworkInfo) {
	o.NetworkInfo = v
}

func (o DeviceInfoExtended) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceInfoExtended) ToMap() (map[string]interface{}, error) {
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
	toSerialize["hibernated"] = o.Hibernated
	if !IsNil(o.LastUpdated) {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if !IsNil(o.DeviceFleets) {
		toSerialize["deviceFleets"] = o.DeviceFleets
	}
	if !IsNil(o.DevicePackages) {
		toSerialize["devicePackages"] = o.DevicePackages
	}
	if !IsNil(o.DeviceTags) {
		toSerialize["deviceTags"] = o.DeviceTags
	}
	toSerialize["networkInfo"] = o.NetworkInfo
	return toSerialize, nil
}

func (o *DeviceInfoExtended) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"deviceUuid",
		"deviceName",
		"deviceId",
		"createdAt",
		"deviceStatus",
		"hibernated",
		"networkInfo",
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

	varDeviceInfoExtended := _DeviceInfoExtended{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDeviceInfoExtended)

	if err != nil {
		return err
	}

	*o = DeviceInfoExtended(varDeviceInfoExtended)

	return err
}

type NullableDeviceInfoExtended struct {
	value *DeviceInfoExtended
	isSet bool
}

func (v NullableDeviceInfoExtended) Get() *DeviceInfoExtended {
	return v.value
}

func (v *NullableDeviceInfoExtended) Set(val *DeviceInfoExtended) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceInfoExtended) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceInfoExtended) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceInfoExtended(val *DeviceInfoExtended) *NullableDeviceInfoExtended {
	return &NullableDeviceInfoExtended{value: val, isSet: true}
}

func (v NullableDeviceInfoExtended) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceInfoExtended) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


