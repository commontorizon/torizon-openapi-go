# ErrorRepresentation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** |  | 
**Description** | **string** |  | 
**Cause** | Pointer to **NullableString** |  | [optional] 
**ErrorId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewErrorRepresentation

`func NewErrorRepresentation(code string, description string, ) *ErrorRepresentation`

NewErrorRepresentation instantiates a new ErrorRepresentation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorRepresentationWithDefaults

`func NewErrorRepresentationWithDefaults() *ErrorRepresentation`

NewErrorRepresentationWithDefaults instantiates a new ErrorRepresentation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ErrorRepresentation) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ErrorRepresentation) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ErrorRepresentation) SetCode(v string)`

SetCode sets Code field to given value.


### GetDescription

`func (o *ErrorRepresentation) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ErrorRepresentation) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ErrorRepresentation) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetCause

`func (o *ErrorRepresentation) GetCause() string`

GetCause returns the Cause field if non-nil, zero value otherwise.

### GetCauseOk

`func (o *ErrorRepresentation) GetCauseOk() (*string, bool)`

GetCauseOk returns a tuple with the Cause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCause

`func (o *ErrorRepresentation) SetCause(v string)`

SetCause sets Cause field to given value.

### HasCause

`func (o *ErrorRepresentation) HasCause() bool`

HasCause returns a boolean if a field has been set.

### SetCauseNil

`func (o *ErrorRepresentation) SetCauseNil(b bool)`

 SetCauseNil sets the value for Cause to be an explicit nil

### UnsetCause
`func (o *ErrorRepresentation) UnsetCause()`

UnsetCause ensures that no value is present for Cause, not even an explicit nil
### GetErrorId

`func (o *ErrorRepresentation) GetErrorId() string`

GetErrorId returns the ErrorId field if non-nil, zero value otherwise.

### GetErrorIdOk

`func (o *ErrorRepresentation) GetErrorIdOk() (*string, bool)`

GetErrorIdOk returns a tuple with the ErrorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorId

`func (o *ErrorRepresentation) SetErrorId(v string)`

SetErrorId sets ErrorId field to given value.

### HasErrorId

`func (o *ErrorRepresentation) HasErrorId() bool`

HasErrorId returns a boolean if a field has been set.

### SetErrorIdNil

`func (o *ErrorRepresentation) SetErrorIdNil(b bool)`

 SetErrorIdNil sets the value for ErrorId to be an explicit nil

### UnsetErrorId
`func (o *ErrorRepresentation) UnsetErrorId()`

UnsetErrorId ensures that no value is present for ErrorId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


