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
)

// checks if the CreateLockboxRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateLockboxRequest{}

// CreateLockboxRequest struct for CreateLockboxRequest
type CreateLockboxRequest struct {
	PackageIds []string `json:"packageIds,omitempty"`
	Custom *map[string]CustomUpdateData `json:"custom,omitempty"`
	ExpiresAt *time.Time `json:"expiresAt,omitempty"`
}

// NewCreateLockboxRequest instantiates a new CreateLockboxRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateLockboxRequest() *CreateLockboxRequest {
	this := CreateLockboxRequest{}
	return &this
}

// NewCreateLockboxRequestWithDefaults instantiates a new CreateLockboxRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateLockboxRequestWithDefaults() *CreateLockboxRequest {
	this := CreateLockboxRequest{}
	return &this
}

// GetPackageIds returns the PackageIds field value if set, zero value otherwise.
func (o *CreateLockboxRequest) GetPackageIds() []string {
	if o == nil || IsNil(o.PackageIds) {
		var ret []string
		return ret
	}
	return o.PackageIds
}

// GetPackageIdsOk returns a tuple with the PackageIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLockboxRequest) GetPackageIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.PackageIds) {
		return nil, false
	}
	return o.PackageIds, true
}

// HasPackageIds returns a boolean if a field has been set.
func (o *CreateLockboxRequest) HasPackageIds() bool {
	if o != nil && !IsNil(o.PackageIds) {
		return true
	}

	return false
}

// SetPackageIds gets a reference to the given []string and assigns it to the PackageIds field.
func (o *CreateLockboxRequest) SetPackageIds(v []string) {
	o.PackageIds = v
}

// GetCustom returns the Custom field value if set, zero value otherwise.
func (o *CreateLockboxRequest) GetCustom() map[string]CustomUpdateData {
	if o == nil || IsNil(o.Custom) {
		var ret map[string]CustomUpdateData
		return ret
	}
	return *o.Custom
}

// GetCustomOk returns a tuple with the Custom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLockboxRequest) GetCustomOk() (*map[string]CustomUpdateData, bool) {
	if o == nil || IsNil(o.Custom) {
		return nil, false
	}
	return o.Custom, true
}

// HasCustom returns a boolean if a field has been set.
func (o *CreateLockboxRequest) HasCustom() bool {
	if o != nil && !IsNil(o.Custom) {
		return true
	}

	return false
}

// SetCustom gets a reference to the given map[string]CustomUpdateData and assigns it to the Custom field.
func (o *CreateLockboxRequest) SetCustom(v map[string]CustomUpdateData) {
	o.Custom = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *CreateLockboxRequest) GetExpiresAt() time.Time {
	if o == nil || IsNil(o.ExpiresAt) {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLockboxRequest) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpiresAt) {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *CreateLockboxRequest) HasExpiresAt() bool {
	if o != nil && !IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given time.Time and assigns it to the ExpiresAt field.
func (o *CreateLockboxRequest) SetExpiresAt(v time.Time) {
	o.ExpiresAt = &v
}

func (o CreateLockboxRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateLockboxRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PackageIds) {
		toSerialize["packageIds"] = o.PackageIds
	}
	if !IsNil(o.Custom) {
		toSerialize["custom"] = o.Custom
	}
	if !IsNil(o.ExpiresAt) {
		toSerialize["expiresAt"] = o.ExpiresAt
	}
	return toSerialize, nil
}

type NullableCreateLockboxRequest struct {
	value *CreateLockboxRequest
	isSet bool
}

func (v NullableCreateLockboxRequest) Get() *CreateLockboxRequest {
	return v.value
}

func (v *NullableCreateLockboxRequest) Set(val *CreateLockboxRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateLockboxRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateLockboxRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateLockboxRequest(val *CreateLockboxRequest) *NullableCreateLockboxRequest {
	return &NullableCreateLockboxRequest{value: val, isSet: true}
}

func (v NullableCreateLockboxRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateLockboxRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

