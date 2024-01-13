# UpstreamEndpointErrorRepr

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Msg** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Code** | **string** |  | 
**Cause** | Pointer to **string** |  | [optional] 
**ErrorId** | Pointer to **string** |  | [optional] 

## Methods

### NewUpstreamEndpointErrorRepr

`func NewUpstreamEndpointErrorRepr(msg string, code string, ) *UpstreamEndpointErrorRepr`

NewUpstreamEndpointErrorRepr instantiates a new UpstreamEndpointErrorRepr object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpstreamEndpointErrorReprWithDefaults

`func NewUpstreamEndpointErrorReprWithDefaults() *UpstreamEndpointErrorRepr`

NewUpstreamEndpointErrorReprWithDefaults instantiates a new UpstreamEndpointErrorRepr object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMsg

`func (o *UpstreamEndpointErrorRepr) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *UpstreamEndpointErrorRepr) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *UpstreamEndpointErrorRepr) SetMsg(v string)`

SetMsg sets Msg field to given value.


### GetDescription

`func (o *UpstreamEndpointErrorRepr) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpstreamEndpointErrorRepr) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpstreamEndpointErrorRepr) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpstreamEndpointErrorRepr) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCode

`func (o *UpstreamEndpointErrorRepr) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *UpstreamEndpointErrorRepr) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *UpstreamEndpointErrorRepr) SetCode(v string)`

SetCode sets Code field to given value.


### GetCause

`func (o *UpstreamEndpointErrorRepr) GetCause() string`

GetCause returns the Cause field if non-nil, zero value otherwise.

### GetCauseOk

`func (o *UpstreamEndpointErrorRepr) GetCauseOk() (*string, bool)`

GetCauseOk returns a tuple with the Cause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCause

`func (o *UpstreamEndpointErrorRepr) SetCause(v string)`

SetCause sets Cause field to given value.

### HasCause

`func (o *UpstreamEndpointErrorRepr) HasCause() bool`

HasCause returns a boolean if a field has been set.

### GetErrorId

`func (o *UpstreamEndpointErrorRepr) GetErrorId() string`

GetErrorId returns the ErrorId field if non-nil, zero value otherwise.

### GetErrorIdOk

`func (o *UpstreamEndpointErrorRepr) GetErrorIdOk() (*string, bool)`

GetErrorIdOk returns a tuple with the ErrorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorId

`func (o *UpstreamEndpointErrorRepr) SetErrorId(v string)`

SetErrorId sets ErrorId field to given value.

### HasErrorId

`func (o *UpstreamEndpointErrorRepr) HasErrorId() bool`

HasErrorId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


