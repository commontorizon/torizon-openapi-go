# DeviceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceId** | **string** |  | 
**LastApiConnect** | **time.Time** |  | 
**LastSshSessionConnect** | [**LastSshSession**](LastSshSession.md) |  | 
**LastVersion** | **string** |  | 
**LastUserConnect** | [**UserInfo**](UserInfo.md) |  | 

## Methods

### NewDeviceInfo

`func NewDeviceInfo(deviceId string, lastApiConnect time.Time, lastSshSessionConnect LastSshSession, lastVersion string, lastUserConnect UserInfo, ) *DeviceInfo`

NewDeviceInfo instantiates a new DeviceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceInfoWithDefaults

`func NewDeviceInfoWithDefaults() *DeviceInfo`

NewDeviceInfoWithDefaults instantiates a new DeviceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceId

`func (o *DeviceInfo) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *DeviceInfo) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *DeviceInfo) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.


### GetLastApiConnect

`func (o *DeviceInfo) GetLastApiConnect() time.Time`

GetLastApiConnect returns the LastApiConnect field if non-nil, zero value otherwise.

### GetLastApiConnectOk

`func (o *DeviceInfo) GetLastApiConnectOk() (*time.Time, bool)`

GetLastApiConnectOk returns a tuple with the LastApiConnect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastApiConnect

`func (o *DeviceInfo) SetLastApiConnect(v time.Time)`

SetLastApiConnect sets LastApiConnect field to given value.


### GetLastSshSessionConnect

`func (o *DeviceInfo) GetLastSshSessionConnect() LastSshSession`

GetLastSshSessionConnect returns the LastSshSessionConnect field if non-nil, zero value otherwise.

### GetLastSshSessionConnectOk

`func (o *DeviceInfo) GetLastSshSessionConnectOk() (*LastSshSession, bool)`

GetLastSshSessionConnectOk returns a tuple with the LastSshSessionConnect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSshSessionConnect

`func (o *DeviceInfo) SetLastSshSessionConnect(v LastSshSession)`

SetLastSshSessionConnect sets LastSshSessionConnect field to given value.


### GetLastVersion

`func (o *DeviceInfo) GetLastVersion() string`

GetLastVersion returns the LastVersion field if non-nil, zero value otherwise.

### GetLastVersionOk

`func (o *DeviceInfo) GetLastVersionOk() (*string, bool)`

GetLastVersionOk returns a tuple with the LastVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastVersion

`func (o *DeviceInfo) SetLastVersion(v string)`

SetLastVersion sets LastVersion field to given value.


### GetLastUserConnect

`func (o *DeviceInfo) GetLastUserConnect() UserInfo`

GetLastUserConnect returns the LastUserConnect field if non-nil, zero value otherwise.

### GetLastUserConnectOk

`func (o *DeviceInfo) GetLastUserConnectOk() (*UserInfo, bool)`

GetLastUserConnectOk returns a tuple with the LastUserConnect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUserConnect

`func (o *DeviceInfo) SetLastUserConnect(v UserInfo)`

SetLastUserConnect sets LastUserConnect field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


