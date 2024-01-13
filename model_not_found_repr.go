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

// checks if the NotFoundRepr type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotFoundRepr{}

// NotFoundRepr struct for NotFoundRepr
type NotFoundRepr struct {
	Msg string `json:"msg"`
	Description *string `json:"description,omitempty"`
	Code string `json:"code"`
	Cause *string `json:"cause,omitempty"`
	ErrorId *string `json:"errorId,omitempty"`
}

type _NotFoundRepr NotFoundRepr

// NewNotFoundRepr instantiates a new NotFoundRepr object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotFoundRepr(msg string, code string) *NotFoundRepr {
	this := NotFoundRepr{}
	this.Msg = msg
	this.Code = code
	return &this
}

// NewNotFoundReprWithDefaults instantiates a new NotFoundRepr object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotFoundReprWithDefaults() *NotFoundRepr {
	this := NotFoundRepr{}
	return &this
}

// GetMsg returns the Msg field value
func (o *NotFoundRepr) GetMsg() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Msg
}

// GetMsgOk returns a tuple with the Msg field value
// and a boolean to check if the value has been set.
func (o *NotFoundRepr) GetMsgOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Msg, true
}

// SetMsg sets field value
func (o *NotFoundRepr) SetMsg(v string) {
	o.Msg = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *NotFoundRepr) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotFoundRepr) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *NotFoundRepr) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *NotFoundRepr) SetDescription(v string) {
	o.Description = &v
}

// GetCode returns the Code field value
func (o *NotFoundRepr) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *NotFoundRepr) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *NotFoundRepr) SetCode(v string) {
	o.Code = v
}

// GetCause returns the Cause field value if set, zero value otherwise.
func (o *NotFoundRepr) GetCause() string {
	if o == nil || IsNil(o.Cause) {
		var ret string
		return ret
	}
	return *o.Cause
}

// GetCauseOk returns a tuple with the Cause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotFoundRepr) GetCauseOk() (*string, bool) {
	if o == nil || IsNil(o.Cause) {
		return nil, false
	}
	return o.Cause, true
}

// HasCause returns a boolean if a field has been set.
func (o *NotFoundRepr) HasCause() bool {
	if o != nil && !IsNil(o.Cause) {
		return true
	}

	return false
}

// SetCause gets a reference to the given string and assigns it to the Cause field.
func (o *NotFoundRepr) SetCause(v string) {
	o.Cause = &v
}

// GetErrorId returns the ErrorId field value if set, zero value otherwise.
func (o *NotFoundRepr) GetErrorId() string {
	if o == nil || IsNil(o.ErrorId) {
		var ret string
		return ret
	}
	return *o.ErrorId
}

// GetErrorIdOk returns a tuple with the ErrorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotFoundRepr) GetErrorIdOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorId) {
		return nil, false
	}
	return o.ErrorId, true
}

// HasErrorId returns a boolean if a field has been set.
func (o *NotFoundRepr) HasErrorId() bool {
	if o != nil && !IsNil(o.ErrorId) {
		return true
	}

	return false
}

// SetErrorId gets a reference to the given string and assigns it to the ErrorId field.
func (o *NotFoundRepr) SetErrorId(v string) {
	o.ErrorId = &v
}

func (o NotFoundRepr) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotFoundRepr) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["msg"] = o.Msg
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["code"] = o.Code
	if !IsNil(o.Cause) {
		toSerialize["cause"] = o.Cause
	}
	if !IsNil(o.ErrorId) {
		toSerialize["errorId"] = o.ErrorId
	}
	return toSerialize, nil
}

func (o *NotFoundRepr) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"msg",
		"code",
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

	varNotFoundRepr := _NotFoundRepr{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNotFoundRepr)

	if err != nil {
		return err
	}

	*o = NotFoundRepr(varNotFoundRepr)

	return err
}

type NullableNotFoundRepr struct {
	value *NotFoundRepr
	isSet bool
}

func (v NullableNotFoundRepr) Get() *NotFoundRepr {
	return v.value
}

func (v *NullableNotFoundRepr) Set(val *NotFoundRepr) {
	v.value = val
	v.isSet = true
}

func (v NullableNotFoundRepr) IsSet() bool {
	return v.isSet
}

func (v *NullableNotFoundRepr) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotFoundRepr(val *NotFoundRepr) *NullableNotFoundRepr {
	return &NullableNotFoundRepr{value: val, isSet: true}
}

func (v NullableNotFoundRepr) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotFoundRepr) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

