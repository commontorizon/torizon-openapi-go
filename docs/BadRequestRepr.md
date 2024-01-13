# BadRequestRepr

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Msg** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Code** | **string** |  | 
**Cause** | Pointer to **string** |  | [optional] 
**ErrorId** | Pointer to **string** |  | [optional] 

## Methods

### NewBadRequestRepr

`func NewBadRequestRepr(msg string, code string, ) *BadRequestRepr`

NewBadRequestRepr instantiates a new BadRequestRepr object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBadRequestReprWithDefaults

`func NewBadRequestReprWithDefaults() *BadRequestRepr`

NewBadRequestReprWithDefaults instantiates a new BadRequestRepr object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMsg

`func (o *BadRequestRepr) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *BadRequestRepr) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *BadRequestRepr) SetMsg(v string)`

SetMsg sets Msg field to given value.


### GetDescription

`func (o *BadRequestRepr) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BadRequestRepr) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BadRequestRepr) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BadRequestRepr) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCode

`func (o *BadRequestRepr) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *BadRequestRepr) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *BadRequestRepr) SetCode(v string)`

SetCode sets Code field to given value.


### GetCause

`func (o *BadRequestRepr) GetCause() string`

GetCause returns the Cause field if non-nil, zero value otherwise.

### GetCauseOk

`func (o *BadRequestRepr) GetCauseOk() (*string, bool)`

GetCauseOk returns a tuple with the Cause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCause

`func (o *BadRequestRepr) SetCause(v string)`

SetCause sets Cause field to given value.

### HasCause

`func (o *BadRequestRepr) HasCause() bool`

HasCause returns a boolean if a field has been set.

### GetErrorId

`func (o *BadRequestRepr) GetErrorId() string`

GetErrorId returns the ErrorId field if non-nil, zero value otherwise.

### GetErrorIdOk

`func (o *BadRequestRepr) GetErrorIdOk() (*string, bool)`

GetErrorIdOk returns a tuple with the ErrorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorId

`func (o *BadRequestRepr) SetErrorId(v string)`

SetErrorId sets ErrorId field to given value.

### HasErrorId

`func (o *BadRequestRepr) HasErrorId() bool`

HasErrorId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


