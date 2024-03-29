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

// checks if the PaginationResultDevicePackages type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginationResultDevicePackages{}

// PaginationResultDevicePackages struct for PaginationResultDevicePackages
type PaginationResultDevicePackages struct {
	Values []DevicePackages `json:"values,omitempty"`
	Total int64 `json:"total"`
	Offset int64 `json:"offset"`
	Limit int64 `json:"limit"`
}

type _PaginationResultDevicePackages PaginationResultDevicePackages

// NewPaginationResultDevicePackages instantiates a new PaginationResultDevicePackages object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginationResultDevicePackages(total int64, offset int64, limit int64) *PaginationResultDevicePackages {
	this := PaginationResultDevicePackages{}
	this.Total = total
	this.Offset = offset
	this.Limit = limit
	return &this
}

// NewPaginationResultDevicePackagesWithDefaults instantiates a new PaginationResultDevicePackages object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginationResultDevicePackagesWithDefaults() *PaginationResultDevicePackages {
	this := PaginationResultDevicePackages{}
	return &this
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *PaginationResultDevicePackages) GetValues() []DevicePackages {
	if o == nil || IsNil(o.Values) {
		var ret []DevicePackages
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginationResultDevicePackages) GetValuesOk() ([]DevicePackages, bool) {
	if o == nil || IsNil(o.Values) {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *PaginationResultDevicePackages) HasValues() bool {
	if o != nil && !IsNil(o.Values) {
		return true
	}

	return false
}

// SetValues gets a reference to the given []DevicePackages and assigns it to the Values field.
func (o *PaginationResultDevicePackages) SetValues(v []DevicePackages) {
	o.Values = v
}

// GetTotal returns the Total field value
func (o *PaginationResultDevicePackages) GetTotal() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *PaginationResultDevicePackages) GetTotalOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *PaginationResultDevicePackages) SetTotal(v int64) {
	o.Total = v
}

// GetOffset returns the Offset field value
func (o *PaginationResultDevicePackages) GetOffset() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value
// and a boolean to check if the value has been set.
func (o *PaginationResultDevicePackages) GetOffsetOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Offset, true
}

// SetOffset sets field value
func (o *PaginationResultDevicePackages) SetOffset(v int64) {
	o.Offset = v
}

// GetLimit returns the Limit field value
func (o *PaginationResultDevicePackages) GetLimit() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *PaginationResultDevicePackages) GetLimitOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *PaginationResultDevicePackages) SetLimit(v int64) {
	o.Limit = v
}

func (o PaginationResultDevicePackages) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginationResultDevicePackages) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Values) {
		toSerialize["values"] = o.Values
	}
	toSerialize["total"] = o.Total
	toSerialize["offset"] = o.Offset
	toSerialize["limit"] = o.Limit
	return toSerialize, nil
}

func (o *PaginationResultDevicePackages) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"total",
		"offset",
		"limit",
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

	varPaginationResultDevicePackages := _PaginationResultDevicePackages{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPaginationResultDevicePackages)

	if err != nil {
		return err
	}

	*o = PaginationResultDevicePackages(varPaginationResultDevicePackages)

	return err
}

type NullablePaginationResultDevicePackages struct {
	value *PaginationResultDevicePackages
	isSet bool
}

func (v NullablePaginationResultDevicePackages) Get() *PaginationResultDevicePackages {
	return v.value
}

func (v *NullablePaginationResultDevicePackages) Set(val *PaginationResultDevicePackages) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginationResultDevicePackages) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginationResultDevicePackages) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginationResultDevicePackages(val *PaginationResultDevicePackages) *NullablePaginationResultDevicePackages {
	return &NullablePaginationResultDevicePackages{value: val, isSet: true}
}

func (v NullablePaginationResultDevicePackages) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginationResultDevicePackages) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


