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

// checks if the Image type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Image{}

// Image struct for Image
type Image struct {
	Filepath string `json:"filepath"`
	Fileinfo FileInfo `json:"fileinfo"`
}

type _Image Image

// NewImage instantiates a new Image object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImage(filepath string, fileinfo FileInfo) *Image {
	this := Image{}
	this.Filepath = filepath
	this.Fileinfo = fileinfo
	return &this
}

// NewImageWithDefaults instantiates a new Image object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageWithDefaults() *Image {
	this := Image{}
	return &this
}

// GetFilepath returns the Filepath field value
func (o *Image) GetFilepath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Filepath
}

// GetFilepathOk returns a tuple with the Filepath field value
// and a boolean to check if the value has been set.
func (o *Image) GetFilepathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Filepath, true
}

// SetFilepath sets field value
func (o *Image) SetFilepath(v string) {
	o.Filepath = v
}

// GetFileinfo returns the Fileinfo field value
func (o *Image) GetFileinfo() FileInfo {
	if o == nil {
		var ret FileInfo
		return ret
	}

	return o.Fileinfo
}

// GetFileinfoOk returns a tuple with the Fileinfo field value
// and a boolean to check if the value has been set.
func (o *Image) GetFileinfoOk() (*FileInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fileinfo, true
}

// SetFileinfo sets field value
func (o *Image) SetFileinfo(v FileInfo) {
	o.Fileinfo = v
}

func (o Image) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Image) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["filepath"] = o.Filepath
	toSerialize["fileinfo"] = o.Fileinfo
	return toSerialize, nil
}

func (o *Image) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"filepath",
		"fileinfo",
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

	varImage := _Image{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varImage)

	if err != nil {
		return err
	}

	*o = Image(varImage)

	return err
}

type NullableImage struct {
	value *Image
	isSet bool
}

func (v NullableImage) Get() *Image {
	return v.value
}

func (v *NullableImage) Set(val *Image) {
	v.value = val
	v.isSet = true
}

func (v NullableImage) IsSet() bool {
	return v.isSet
}

func (v *NullableImage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImage(val *Image) *NullableImage {
	return &NullableImage{value: val, isSet: true}
}

func (v NullableImage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


