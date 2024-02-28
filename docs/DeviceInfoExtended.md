# DeviceInfoExtended

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceUuid** | **string** |  | 
**DeviceName** | **string** |  | 
**DeviceId** | **string** |  | 
**LastSeen** | Pointer to **time.Time** |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**ActivatedAt** | Pointer to **time.Time** |  | [optional] 
**DeviceStatus** | [**DeviceStatus**](DeviceStatus.md) |  | 
**Notes** | Pointer to **string** |  | [optional] 
**Hibernated** | **bool** |  | 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**DeviceFleets** | Pointer to [**[]Fleet**](Fleet.md) |  | [optional] 
**DevicePackages** | Pointer to [**[]DevicePackage**](DevicePackage.md) |  | [optional] 
**DeviceTags** | Pointer to [**[]Tuple2DeviceTagIdDeviceTagValue**](Tuple2DeviceTagIdDeviceTagValue.md) |  | [optional] 
**NetworkInfo** | [**NetworkInfo**](NetworkInfo.md) |  | 

## Methods

### NewDeviceInfoExtended

`func NewDeviceInfoExtended(deviceUuid string, deviceName string, deviceId string, createdAt time.Time, deviceStatus DeviceStatus, hibernated bool, networkInfo NetworkInfo, ) *DeviceInfoExtended`

NewDeviceInfoExtended instantiates a new DeviceInfoExtended object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceInfoExtendedWithDefaults

`func NewDeviceInfoExtendedWithDefaults() *DeviceInfoExtended`

NewDeviceInfoExtendedWithDefaults instantiates a new DeviceInfoExtended object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceUuid

`func (o *DeviceInfoExtended) GetDeviceUuid() string`

GetDeviceUuid returns the DeviceUuid field if non-nil, zero value otherwise.

### GetDeviceUuidOk

`func (o *DeviceInfoExtended) GetDeviceUuidOk() (*string, bool)`

GetDeviceUuidOk returns a tuple with the DeviceUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceUuid

`func (o *DeviceInfoExtended) SetDeviceUuid(v string)`

SetDeviceUuid sets DeviceUuid field to given value.


### GetDeviceName

`func (o *DeviceInfoExtended) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *DeviceInfoExtended) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *DeviceInfoExtended) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.


### GetDeviceId

`func (o *DeviceInfoExtended) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *DeviceInfoExtended) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *DeviceInfoExtended) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.


### GetLastSeen

`func (o *DeviceInfoExtended) GetLastSeen() time.Time`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *DeviceInfoExtended) GetLastSeenOk() (*time.Time, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *DeviceInfoExtended) SetLastSeen(v time.Time)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *DeviceInfoExtended) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DeviceInfoExtended) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DeviceInfoExtended) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DeviceInfoExtended) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetActivatedAt

`func (o *DeviceInfoExtended) GetActivatedAt() time.Time`

GetActivatedAt returns the ActivatedAt field if non-nil, zero value otherwise.

### GetActivatedAtOk

`func (o *DeviceInfoExtended) GetActivatedAtOk() (*time.Time, bool)`

GetActivatedAtOk returns a tuple with the ActivatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivatedAt

`func (o *DeviceInfoExtended) SetActivatedAt(v time.Time)`

SetActivatedAt sets ActivatedAt field to given value.

### HasActivatedAt

`func (o *DeviceInfoExtended) HasActivatedAt() bool`

HasActivatedAt returns a boolean if a field has been set.

### GetDeviceStatus

`func (o *DeviceInfoExtended) GetDeviceStatus() DeviceStatus`

GetDeviceStatus returns the DeviceStatus field if non-nil, zero value otherwise.

### GetDeviceStatusOk

`func (o *DeviceInfoExtended) GetDeviceStatusOk() (*DeviceStatus, bool)`

GetDeviceStatusOk returns a tuple with the DeviceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceStatus

`func (o *DeviceInfoExtended) SetDeviceStatus(v DeviceStatus)`

SetDeviceStatus sets DeviceStatus field to given value.


### GetNotes

`func (o *DeviceInfoExtended) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *DeviceInfoExtended) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *DeviceInfoExtended) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *DeviceInfoExtended) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetHibernated

`func (o *DeviceInfoExtended) GetHibernated() bool`

GetHibernated returns the Hibernated field if non-nil, zero value otherwise.

### GetHibernatedOk

`func (o *DeviceInfoExtended) GetHibernatedOk() (*bool, bool)`

GetHibernatedOk returns a tuple with the Hibernated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHibernated

`func (o *DeviceInfoExtended) SetHibernated(v bool)`

SetHibernated sets Hibernated field to given value.


### GetLastUpdated

`func (o *DeviceInfoExtended) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *DeviceInfoExtended) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *DeviceInfoExtended) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *DeviceInfoExtended) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetDeviceFleets

`func (o *DeviceInfoExtended) GetDeviceFleets() []Fleet`

GetDeviceFleets returns the DeviceFleets field if non-nil, zero value otherwise.

### GetDeviceFleetsOk

`func (o *DeviceInfoExtended) GetDeviceFleetsOk() (*[]Fleet, bool)`

GetDeviceFleetsOk returns a tuple with the DeviceFleets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceFleets

`func (o *DeviceInfoExtended) SetDeviceFleets(v []Fleet)`

SetDeviceFleets sets DeviceFleets field to given value.

### HasDeviceFleets

`func (o *DeviceInfoExtended) HasDeviceFleets() bool`

HasDeviceFleets returns a boolean if a field has been set.

### GetDevicePackages

`func (o *DeviceInfoExtended) GetDevicePackages() []DevicePackage`

GetDevicePackages returns the DevicePackages field if non-nil, zero value otherwise.

### GetDevicePackagesOk

`func (o *DeviceInfoExtended) GetDevicePackagesOk() (*[]DevicePackage, bool)`

GetDevicePackagesOk returns a tuple with the DevicePackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicePackages

`func (o *DeviceInfoExtended) SetDevicePackages(v []DevicePackage)`

SetDevicePackages sets DevicePackages field to given value.

### HasDevicePackages

`func (o *DeviceInfoExtended) HasDevicePackages() bool`

HasDevicePackages returns a boolean if a field has been set.

### GetDeviceTags

`func (o *DeviceInfoExtended) GetDeviceTags() []Tuple2DeviceTagIdDeviceTagValue`

GetDeviceTags returns the DeviceTags field if non-nil, zero value otherwise.

### GetDeviceTagsOk

`func (o *DeviceInfoExtended) GetDeviceTagsOk() (*[]Tuple2DeviceTagIdDeviceTagValue, bool)`

GetDeviceTagsOk returns a tuple with the DeviceTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceTags

`func (o *DeviceInfoExtended) SetDeviceTags(v []Tuple2DeviceTagIdDeviceTagValue)`

SetDeviceTags sets DeviceTags field to given value.

### HasDeviceTags

`func (o *DeviceInfoExtended) HasDeviceTags() bool`

HasDeviceTags returns a boolean if a field has been set.

### GetNetworkInfo

`func (o *DeviceInfoExtended) GetNetworkInfo() NetworkInfo`

GetNetworkInfo returns the NetworkInfo field if non-nil, zero value otherwise.

### GetNetworkInfoOk

`func (o *DeviceInfoExtended) GetNetworkInfoOk() (*NetworkInfo, bool)`

GetNetworkInfoOk returns a tuple with the NetworkInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInfo

`func (o *DeviceInfoExtended) SetNetworkInfo(v NetworkInfo)`

SetNetworkInfo sets NetworkInfo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


