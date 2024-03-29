/*
Torizon OTA

 This API is rate limited and will return the following headers for each API call.    - X-RateLimit-Limit - The total number of requests allowed within a time period   - X-RateLimit-Remaining - The total number of requests still allowed until the end of the rate limiting period   - X-RateLimit-Reset - The number of seconds until the limit is fully reset  In addition, if an API client is rate limited, it will receive a HTTP 420 response with the following header:     - Retry-After - The number of seconds to wait until this request is allowed  

API version: 2.0-Beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the EditPackage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EditPackage{}

// EditPackage struct for EditPackage
type EditPackage struct {
	HardwareIds []string `json:"hardwareIds,omitempty"`
	Uri *string `json:"uri,omitempty"`
	ProprietaryMeta *string `json:"proprietaryMeta,omitempty"`
	Comment *string `json:"comment,omitempty"`
}

// NewEditPackage instantiates a new EditPackage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEditPackage() *EditPackage {
	this := EditPackage{}
	return &this
}

// NewEditPackageWithDefaults instantiates a new EditPackage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEditPackageWithDefaults() *EditPackage {
	this := EditPackage{}
	return &this
}

// GetHardwareIds returns the HardwareIds field value if set, zero value otherwise.
func (o *EditPackage) GetHardwareIds() []string {
	if o == nil || IsNil(o.HardwareIds) {
		var ret []string
		return ret
	}
	return o.HardwareIds
}

// GetHardwareIdsOk returns a tuple with the HardwareIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditPackage) GetHardwareIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.HardwareIds) {
		return nil, false
	}
	return o.HardwareIds, true
}

// HasHardwareIds returns a boolean if a field has been set.
func (o *EditPackage) HasHardwareIds() bool {
	if o != nil && !IsNil(o.HardwareIds) {
		return true
	}

	return false
}

// SetHardwareIds gets a reference to the given []string and assigns it to the HardwareIds field.
func (o *EditPackage) SetHardwareIds(v []string) {
	o.HardwareIds = v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *EditPackage) GetUri() string {
	if o == nil || IsNil(o.Uri) {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditPackage) GetUriOk() (*string, bool) {
	if o == nil || IsNil(o.Uri) {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *EditPackage) HasUri() bool {
	if o != nil && !IsNil(o.Uri) {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *EditPackage) SetUri(v string) {
	o.Uri = &v
}

// GetProprietaryMeta returns the ProprietaryMeta field value if set, zero value otherwise.
func (o *EditPackage) GetProprietaryMeta() string {
	if o == nil || IsNil(o.ProprietaryMeta) {
		var ret string
		return ret
	}
	return *o.ProprietaryMeta
}

// GetProprietaryMetaOk returns a tuple with the ProprietaryMeta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditPackage) GetProprietaryMetaOk() (*string, bool) {
	if o == nil || IsNil(o.ProprietaryMeta) {
		return nil, false
	}
	return o.ProprietaryMeta, true
}

// HasProprietaryMeta returns a boolean if a field has been set.
func (o *EditPackage) HasProprietaryMeta() bool {
	if o != nil && !IsNil(o.ProprietaryMeta) {
		return true
	}

	return false
}

// SetProprietaryMeta gets a reference to the given string and assigns it to the ProprietaryMeta field.
func (o *EditPackage) SetProprietaryMeta(v string) {
	o.ProprietaryMeta = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *EditPackage) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditPackage) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *EditPackage) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *EditPackage) SetComment(v string) {
	o.Comment = &v
}

func (o EditPackage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EditPackage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HardwareIds) {
		toSerialize["hardwareIds"] = o.HardwareIds
	}
	if !IsNil(o.Uri) {
		toSerialize["uri"] = o.Uri
	}
	if !IsNil(o.ProprietaryMeta) {
		toSerialize["proprietaryMeta"] = o.ProprietaryMeta
	}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	return toSerialize, nil
}

type NullableEditPackage struct {
	value *EditPackage
	isSet bool
}

func (v NullableEditPackage) Get() *EditPackage {
	return v.value
}

func (v *NullableEditPackage) Set(val *EditPackage) {
	v.value = val
	v.isSet = true
}

func (v NullableEditPackage) IsSet() bool {
	return v.isSet
}

func (v *NullableEditPackage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditPackage(val *EditPackage) *NullableEditPackage {
	return &NullableEditPackage{value: val, isSet: true}
}

func (v NullableEditPackage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditPackage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


