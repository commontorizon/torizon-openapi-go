# UpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PackageIds** | Pointer to **[]string** |  | [optional] 
**Custom** | Pointer to [**map[string]CustomUpdateData**](CustomUpdateData.md) |  | [optional] 
**Devices** | Pointer to **[]string** |  | [optional] 
**Fleets** | Pointer to **[]string** |  | [optional] 

## Methods

### NewUpdateRequest

`func NewUpdateRequest() *UpdateRequest`

NewUpdateRequest instantiates a new UpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRequestWithDefaults

`func NewUpdateRequestWithDefaults() *UpdateRequest`

NewUpdateRequestWithDefaults instantiates a new UpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPackageIds

`func (o *UpdateRequest) GetPackageIds() []string`

GetPackageIds returns the PackageIds field if non-nil, zero value otherwise.

### GetPackageIdsOk

`func (o *UpdateRequest) GetPackageIdsOk() (*[]string, bool)`

GetPackageIdsOk returns a tuple with the PackageIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageIds

`func (o *UpdateRequest) SetPackageIds(v []string)`

SetPackageIds sets PackageIds field to given value.

### HasPackageIds

`func (o *UpdateRequest) HasPackageIds() bool`

HasPackageIds returns a boolean if a field has been set.

### GetCustom

`func (o *UpdateRequest) GetCustom() map[string]CustomUpdateData`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *UpdateRequest) GetCustomOk() (*map[string]CustomUpdateData, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *UpdateRequest) SetCustom(v map[string]CustomUpdateData)`

SetCustom sets Custom field to given value.

### HasCustom

`func (o *UpdateRequest) HasCustom() bool`

HasCustom returns a boolean if a field has been set.

### GetDevices

`func (o *UpdateRequest) GetDevices() []string`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *UpdateRequest) GetDevicesOk() (*[]string, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *UpdateRequest) SetDevices(v []string)`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *UpdateRequest) HasDevices() bool`

HasDevices returns a boolean if a field has been set.

### GetFleets

`func (o *UpdateRequest) GetFleets() []string`

GetFleets returns the Fleets field if non-nil, zero value otherwise.

### GetFleetsOk

`func (o *UpdateRequest) GetFleetsOk() (*[]string, bool)`

GetFleetsOk returns a tuple with the Fleets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleets

`func (o *UpdateRequest) SetFleets(v []string)`

SetFleets sets Fleets field to given value.

### HasFleets

`func (o *UpdateRequest) HasFleets() bool`

HasFleets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


