# SimpleDeviceNotAffectedInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceUuid** | **string** |  | 
**DeviceId** | **string** |  | 
**Name** | **string** |  | 
**EcuErrors** | [**map[string]ErrorRepresentation**](ErrorRepresentation.md) |  | 

## Methods

### NewSimpleDeviceNotAffectedInfo

`func NewSimpleDeviceNotAffectedInfo(deviceUuid string, deviceId string, name string, ecuErrors map[string]ErrorRepresentation, ) *SimpleDeviceNotAffectedInfo`

NewSimpleDeviceNotAffectedInfo instantiates a new SimpleDeviceNotAffectedInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimpleDeviceNotAffectedInfoWithDefaults

`func NewSimpleDeviceNotAffectedInfoWithDefaults() *SimpleDeviceNotAffectedInfo`

NewSimpleDeviceNotAffectedInfoWithDefaults instantiates a new SimpleDeviceNotAffectedInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceUuid

`func (o *SimpleDeviceNotAffectedInfo) GetDeviceUuid() string`

GetDeviceUuid returns the DeviceUuid field if non-nil, zero value otherwise.

### GetDeviceUuidOk

`func (o *SimpleDeviceNotAffectedInfo) GetDeviceUuidOk() (*string, bool)`

GetDeviceUuidOk returns a tuple with the DeviceUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceUuid

`func (o *SimpleDeviceNotAffectedInfo) SetDeviceUuid(v string)`

SetDeviceUuid sets DeviceUuid field to given value.


### GetDeviceId

`func (o *SimpleDeviceNotAffectedInfo) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *SimpleDeviceNotAffectedInfo) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *SimpleDeviceNotAffectedInfo) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.


### GetName

`func (o *SimpleDeviceNotAffectedInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SimpleDeviceNotAffectedInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SimpleDeviceNotAffectedInfo) SetName(v string)`

SetName sets Name field to given value.


### GetEcuErrors

`func (o *SimpleDeviceNotAffectedInfo) GetEcuErrors() map[string]ErrorRepresentation`

GetEcuErrors returns the EcuErrors field if non-nil, zero value otherwise.

### GetEcuErrorsOk

`func (o *SimpleDeviceNotAffectedInfo) GetEcuErrorsOk() (*map[string]ErrorRepresentation, bool)`

GetEcuErrorsOk returns a tuple with the EcuErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcuErrors

`func (o *SimpleDeviceNotAffectedInfo) SetEcuErrors(v map[string]ErrorRepresentation)`

SetEcuErrors sets EcuErrors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


