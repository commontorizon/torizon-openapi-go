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

// checks if the FileInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileInfo{}

// FileInfo struct for FileInfo
type FileInfo struct {
	Hashes string `json:"hashes"`
	Length int64 `json:"length"`
}

type _FileInfo FileInfo

// NewFileInfo instantiates a new FileInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileInfo(hashes string, length int64) *FileInfo {
	this := FileInfo{}
	this.Hashes = hashes
	this.Length = length
	return &this
}

// NewFileInfoWithDefaults instantiates a new FileInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileInfoWithDefaults() *FileInfo {
	this := FileInfo{}
	return &this
}

// GetHashes returns the Hashes field value
func (o *FileInfo) GetHashes() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hashes
}

// GetHashesOk returns a tuple with the Hashes field value
// and a boolean to check if the value has been set.
func (o *FileInfo) GetHashesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hashes, true
}

// SetHashes sets field value
func (o *FileInfo) SetHashes(v string) {
	o.Hashes = v
}

// GetLength returns the Length field value
func (o *FileInfo) GetLength() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Length
}

// GetLengthOk returns a tuple with the Length field value
// and a boolean to check if the value has been set.
func (o *FileInfo) GetLengthOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Length, true
}

// SetLength sets field value
func (o *FileInfo) SetLength(v int64) {
	o.Length = v
}

func (o FileInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["hashes"] = o.Hashes
	toSerialize["length"] = o.Length
	return toSerialize, nil
}

func (o *FileInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"hashes",
		"length",
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

	varFileInfo := _FileInfo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFileInfo)

	if err != nil {
		return err
	}

	*o = FileInfo(varFileInfo)

	return err
}

type NullableFileInfo struct {
	value *FileInfo
	isSet bool
}

func (v NullableFileInfo) Get() *FileInfo {
	return v.value
}

func (v *NullableFileInfo) Set(val *FileInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableFileInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableFileInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileInfo(val *FileInfo) *NullableFileInfo {
	return &NullableFileInfo{value: val, isSet: true}
}

func (v NullableFileInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


