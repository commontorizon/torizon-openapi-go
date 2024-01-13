# DeviceCreateReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceName** | Pointer to **string** |  | [optional] 
**DeviceId** | **string** |  | 

## Methods

### NewDeviceCreateReq

`func NewDeviceCreateReq(deviceId string, ) *DeviceCreateReq`

NewDeviceCreateReq instantiates a new DeviceCreateReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceCreateReqWithDefaults

`func NewDeviceCreateReqWithDefaults() *DeviceCreateReq`

NewDeviceCreateReqWithDefaults instantiates a new DeviceCreateReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceName

`func (o *DeviceCreateReq) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *DeviceCreateReq) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *DeviceCreateReq) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *DeviceCreateReq) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetDeviceId

`func (o *DeviceCreateReq) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *DeviceCreateReq) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *DeviceCreateReq) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


