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

// checks if the CreateFleet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateFleet{}

// CreateFleet struct for CreateFleet
type CreateFleet struct {
	Name string `json:"name"`
	FleetType FleetType `json:"fleetType"`
	Expression *string `json:"expression,omitempty"`
}

type _CreateFleet CreateFleet

// NewCreateFleet instantiates a new CreateFleet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateFleet(name string, fleetType FleetType) *CreateFleet {
	this := CreateFleet{}
	this.Name = name
	this.FleetType = fleetType
	return &this
}

// NewCreateFleetWithDefaults instantiates a new CreateFleet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateFleetWithDefaults() *CreateFleet {
	this := CreateFleet{}
	return &this
}

// GetName returns the Name field value
func (o *CreateFleet) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateFleet) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateFleet) SetName(v string) {
	o.Name = v
}

// GetFleetType returns the FleetType field value
func (o *CreateFleet) GetFleetType() FleetType {
	if o == nil {
		var ret FleetType
		return ret
	}

	return o.FleetType
}

// GetFleetTypeOk returns a tuple with the FleetType field value
// and a boolean to check if the value has been set.
func (o *CreateFleet) GetFleetTypeOk() (*FleetType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FleetType, true
}

// SetFleetType sets field value
func (o *CreateFleet) SetFleetType(v FleetType) {
	o.FleetType = v
}

// GetExpression returns the Expression field value if set, zero value otherwise.
func (o *CreateFleet) GetExpression() string {
	if o == nil || IsNil(o.Expression) {
		var ret string
		return ret
	}
	return *o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFleet) GetExpressionOk() (*string, bool) {
	if o == nil || IsNil(o.Expression) {
		return nil, false
	}
	return o.Expression, true
}

// HasExpression returns a boolean if a field has been set.
func (o *CreateFleet) HasExpression() bool {
	if o != nil && !IsNil(o.Expression) {
		return true
	}

	return false
}

// SetExpression gets a reference to the given string and assigns it to the Expression field.
func (o *CreateFleet) SetExpression(v string) {
	o.Expression = &v
}

func (o CreateFleet) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateFleet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["fleetType"] = o.FleetType
	if !IsNil(o.Expression) {
		toSerialize["expression"] = o.Expression
	}
	return toSerialize, nil
}

func (o *CreateFleet) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"fleetType",
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

	varCreateFleet := _CreateFleet{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateFleet)

	if err != nil {
		return err
	}

	*o = CreateFleet(varCreateFleet)

	return err
}

type NullableCreateFleet struct {
	value *CreateFleet
	isSet bool
}

func (v NullableCreateFleet) Get() *CreateFleet {
	return v.value
}

func (v *NullableCreateFleet) Set(val *CreateFleet) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateFleet) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateFleet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateFleet(val *CreateFleet) *NullableCreateFleet {
	return &NullableCreateFleet{value: val, isSet: true}
}

func (v NullableCreateFleet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateFleet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


