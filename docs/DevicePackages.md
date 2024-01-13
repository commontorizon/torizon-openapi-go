# DevicePackages

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceUuid** | **string** |  | 
**InstalledPackages** | Pointer to [**[]InstalledPackage**](InstalledPackage.md) |  | [optional] 

## Methods

### NewDevicePackages

`func NewDevicePackages(deviceUuid string, ) *DevicePackages`

NewDevicePackages instantiates a new DevicePackages object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDevicePackagesWithDefaults

`func NewDevicePackagesWithDefaults() *DevicePackages`

NewDevicePackagesWithDefaults instantiates a new DevicePackages object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceUuid

`func (o *DevicePackages) GetDeviceUuid() string`

GetDeviceUuid returns the DeviceUuid field if non-nil, zero value otherwise.

### GetDeviceUuidOk

`func (o *DevicePackages) GetDeviceUuidOk() (*string, bool)`

GetDeviceUuidOk returns a tuple with the DeviceUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceUuid

`func (o *DevicePackages) SetDeviceUuid(v string)`

SetDeviceUuid sets DeviceUuid field to given value.


### GetInstalledPackages

`func (o *DevicePackages) GetInstalledPackages() []InstalledPackage`

GetInstalledPackages returns the InstalledPackages field if non-nil, zero value otherwise.

### GetInstalledPackagesOk

`func (o *DevicePackages) GetInstalledPackagesOk() (*[]InstalledPackage, bool)`

GetInstalledPackagesOk returns a tuple with the InstalledPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalledPackages

`func (o *DevicePackages) SetInstalledPackages(v []InstalledPackage)`

SetInstalledPackages sets InstalledPackages field to given value.

### HasInstalledPackages

`func (o *DevicePackages) HasInstalledPackages() bool`

HasInstalledPackages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


