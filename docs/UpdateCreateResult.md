# UpdateCreateResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Affected** | Pointer to [**[]SimpleDeviceInfo**](SimpleDeviceInfo.md) |  | [optional] 
**NotAffected** | Pointer to [**[]SimpleDeviceNotAffectedInfo**](SimpleDeviceNotAffectedInfo.md) |  | [optional] 

## Methods

### NewUpdateCreateResult

`func NewUpdateCreateResult() *UpdateCreateResult`

NewUpdateCreateResult instantiates a new UpdateCreateResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCreateResultWithDefaults

`func NewUpdateCreateResultWithDefaults() *UpdateCreateResult`

NewUpdateCreateResultWithDefaults instantiates a new UpdateCreateResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffected

`func (o *UpdateCreateResult) GetAffected() []SimpleDeviceInfo`

GetAffected returns the Affected field if non-nil, zero value otherwise.

### GetAffectedOk

`func (o *UpdateCreateResult) GetAffectedOk() (*[]SimpleDeviceInfo, bool)`

GetAffectedOk returns a tuple with the Affected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffected

`func (o *UpdateCreateResult) SetAffected(v []SimpleDeviceInfo)`

SetAffected sets Affected field to given value.

### HasAffected

`func (o *UpdateCreateResult) HasAffected() bool`

HasAffected returns a boolean if a field has been set.

### GetNotAffected

`func (o *UpdateCreateResult) GetNotAffected() []SimpleDeviceNotAffectedInfo`

GetNotAffected returns the NotAffected field if non-nil, zero value otherwise.

### GetNotAffectedOk

`func (o *UpdateCreateResult) GetNotAffectedOk() (*[]SimpleDeviceNotAffectedInfo, bool)`

GetNotAffectedOk returns a tuple with the NotAffected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAffected

`func (o *UpdateCreateResult) SetNotAffected(v []SimpleDeviceNotAffectedInfo)`

SetNotAffected sets NotAffected field to given value.

### HasNotAffected

`func (o *UpdateCreateResult) HasNotAffected() bool`

HasNotAffected returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


