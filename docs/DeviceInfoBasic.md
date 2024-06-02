# DeviceInfoBasic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceUuid** | **string** |  | 
**DeviceName** | **string** |  | 
**DeviceId** | **string** |  | 
**LastSeen** | Pointer to **NullableTime** |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**ActivatedAt** | Pointer to **NullableTime** |  | [optional] 
**DeviceStatus** | [**DeviceStatus**](DeviceStatus.md) |  | 
**Notes** | Pointer to **NullableString** |  | [optional] 
**Hibernated** | **bool** |  | 

## Methods

### NewDeviceInfoBasic

`func NewDeviceInfoBasic(deviceUuid string, deviceName string, deviceId string, createdAt time.Time, deviceStatus DeviceStatus, hibernated bool, ) *DeviceInfoBasic`

NewDeviceInfoBasic instantiates a new DeviceInfoBasic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceInfoBasicWithDefaults

`func NewDeviceInfoBasicWithDefaults() *DeviceInfoBasic`

NewDeviceInfoBasicWithDefaults instantiates a new DeviceInfoBasic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceUuid

`func (o *DeviceInfoBasic) GetDeviceUuid() string`

GetDeviceUuid returns the DeviceUuid field if non-nil, zero value otherwise.

### GetDeviceUuidOk

`func (o *DeviceInfoBasic) GetDeviceUuidOk() (*string, bool)`

GetDeviceUuidOk returns a tuple with the DeviceUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceUuid

`func (o *DeviceInfoBasic) SetDeviceUuid(v string)`

SetDeviceUuid sets DeviceUuid field to given value.


### GetDeviceName

`func (o *DeviceInfoBasic) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *DeviceInfoBasic) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *DeviceInfoBasic) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.


### GetDeviceId

`func (o *DeviceInfoBasic) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *DeviceInfoBasic) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *DeviceInfoBasic) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.


### GetLastSeen

`func (o *DeviceInfoBasic) GetLastSeen() time.Time`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *DeviceInfoBasic) GetLastSeenOk() (*time.Time, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *DeviceInfoBasic) SetLastSeen(v time.Time)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *DeviceInfoBasic) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### SetLastSeenNil

`func (o *DeviceInfoBasic) SetLastSeenNil(b bool)`

 SetLastSeenNil sets the value for LastSeen to be an explicit nil

### UnsetLastSeen
`func (o *DeviceInfoBasic) UnsetLastSeen()`

UnsetLastSeen ensures that no value is present for LastSeen, not even an explicit nil
### GetCreatedAt

`func (o *DeviceInfoBasic) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DeviceInfoBasic) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DeviceInfoBasic) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetActivatedAt

`func (o *DeviceInfoBasic) GetActivatedAt() time.Time`

GetActivatedAt returns the ActivatedAt field if non-nil, zero value otherwise.

### GetActivatedAtOk

`func (o *DeviceInfoBasic) GetActivatedAtOk() (*time.Time, bool)`

GetActivatedAtOk returns a tuple with the ActivatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivatedAt

`func (o *DeviceInfoBasic) SetActivatedAt(v time.Time)`

SetActivatedAt sets ActivatedAt field to given value.

### HasActivatedAt

`func (o *DeviceInfoBasic) HasActivatedAt() bool`

HasActivatedAt returns a boolean if a field has been set.

### SetActivatedAtNil

`func (o *DeviceInfoBasic) SetActivatedAtNil(b bool)`

 SetActivatedAtNil sets the value for ActivatedAt to be an explicit nil

### UnsetActivatedAt
`func (o *DeviceInfoBasic) UnsetActivatedAt()`

UnsetActivatedAt ensures that no value is present for ActivatedAt, not even an explicit nil
### GetDeviceStatus

`func (o *DeviceInfoBasic) GetDeviceStatus() DeviceStatus`

GetDeviceStatus returns the DeviceStatus field if non-nil, zero value otherwise.

### GetDeviceStatusOk

`func (o *DeviceInfoBasic) GetDeviceStatusOk() (*DeviceStatus, bool)`

GetDeviceStatusOk returns a tuple with the DeviceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceStatus

`func (o *DeviceInfoBasic) SetDeviceStatus(v DeviceStatus)`

SetDeviceStatus sets DeviceStatus field to given value.


### GetNotes

`func (o *DeviceInfoBasic) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *DeviceInfoBasic) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *DeviceInfoBasic) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *DeviceInfoBasic) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *DeviceInfoBasic) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *DeviceInfoBasic) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetHibernated

`func (o *DeviceInfoBasic) GetHibernated() bool`

GetHibernated returns the Hibernated field if non-nil, zero value otherwise.

### GetHibernatedOk

`func (o *DeviceInfoBasic) GetHibernatedOk() (*bool, bool)`

GetHibernatedOk returns a tuple with the Hibernated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHibernated

`func (o *DeviceInfoBasic) SetHibernated(v bool)`

SetHibernated sets Hibernated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


