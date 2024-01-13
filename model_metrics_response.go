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

// checks if the MetricsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetricsResponse{}

// MetricsResponse struct for MetricsResponse
type MetricsResponse struct {
	Series []Series `json:"series,omitempty"`
}

// NewMetricsResponse instantiates a new MetricsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricsResponse() *MetricsResponse {
	this := MetricsResponse{}
	return &this
}

// NewMetricsResponseWithDefaults instantiates a new MetricsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricsResponseWithDefaults() *MetricsResponse {
	this := MetricsResponse{}
	return &this
}

// GetSeries returns the Series field value if set, zero value otherwise.
func (o *MetricsResponse) GetSeries() []Series {
	if o == nil || IsNil(o.Series) {
		var ret []Series
		return ret
	}
	return o.Series
}

// GetSeriesOk returns a tuple with the Series field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsResponse) GetSeriesOk() ([]Series, bool) {
	if o == nil || IsNil(o.Series) {
		return nil, false
	}
	return o.Series, true
}

// HasSeries returns a boolean if a field has been set.
func (o *MetricsResponse) HasSeries() bool {
	if o != nil && !IsNil(o.Series) {
		return true
	}

	return false
}

// SetSeries gets a reference to the given []Series and assigns it to the Series field.
func (o *MetricsResponse) SetSeries(v []Series) {
	o.Series = v
}

func (o MetricsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetricsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Series) {
		toSerialize["series"] = o.Series
	}
	return toSerialize, nil
}

type NullableMetricsResponse struct {
	value *MetricsResponse
	isSet bool
}

func (v NullableMetricsResponse) Get() *MetricsResponse {
	return v.value
}

func (v *NullableMetricsResponse) Set(val *MetricsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricsResponse(val *MetricsResponse) *NullableMetricsResponse {
	return &NullableMetricsResponse{value: val, isSet: true}
}

func (v NullableMetricsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

