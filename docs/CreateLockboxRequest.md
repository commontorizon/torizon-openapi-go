# CreateLockboxRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PackageIds** | Pointer to **[]string** |  | [optional] 
**Custom** | Pointer to [**map[string]CustomUpdateData**](CustomUpdateData.md) |  | [optional] 
**ExpiresAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewCreateLockboxRequest

`func NewCreateLockboxRequest() *CreateLockboxRequest`

NewCreateLockboxRequest instantiates a new CreateLockboxRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLockboxRequestWithDefaults

`func NewCreateLockboxRequestWithDefaults() *CreateLockboxRequest`

NewCreateLockboxRequestWithDefaults instantiates a new CreateLockboxRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPackageIds

`func (o *CreateLockboxRequest) GetPackageIds() []string`

GetPackageIds returns the PackageIds field if non-nil, zero value otherwise.

### GetPackageIdsOk

`func (o *CreateLockboxRequest) GetPackageIdsOk() (*[]string, bool)`

GetPackageIdsOk returns a tuple with the PackageIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageIds

`func (o *CreateLockboxRequest) SetPackageIds(v []string)`

SetPackageIds sets PackageIds field to given value.

### HasPackageIds

`func (o *CreateLockboxRequest) HasPackageIds() bool`

HasPackageIds returns a boolean if a field has been set.

### GetCustom

`func (o *CreateLockboxRequest) GetCustom() map[string]CustomUpdateData`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *CreateLockboxRequest) GetCustomOk() (*map[string]CustomUpdateData, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *CreateLockboxRequest) SetCustom(v map[string]CustomUpdateData)`

SetCustom sets Custom field to given value.

### HasCustom

`func (o *CreateLockboxRequest) HasCustom() bool`

HasCustom returns a boolean if a field has been set.

### GetExpiresAt

`func (o *CreateLockboxRequest) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *CreateLockboxRequest) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *CreateLockboxRequest) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *CreateLockboxRequest) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


