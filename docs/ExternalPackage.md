# ExternalPackage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**DelegationOrigin** | **string** |  | 
**Version** | **string** |  | 
**PackageId** | **string** |  | 
**Size** | **int64** |  | 
**Hashes** | **map[string]string** |  | 
**PkgType** | Pointer to **NullableString** |  | [optional] 
**HardwareIds** | Pointer to **[]string** |  | [optional] 
**CreatedAt** | Pointer to **NullableTime** |  | [optional] 
**Uri** | Pointer to **NullableString** |  | [optional] 
**ProprietaryMeta** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewExternalPackage

`func NewExternalPackage(name string, delegationOrigin string, version string, packageId string, size int64, hashes map[string]string, ) *ExternalPackage`

NewExternalPackage instantiates a new ExternalPackage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalPackageWithDefaults

`func NewExternalPackageWithDefaults() *ExternalPackage`

NewExternalPackageWithDefaults instantiates a new ExternalPackage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ExternalPackage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExternalPackage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExternalPackage) SetName(v string)`

SetName sets Name field to given value.


### GetDelegationOrigin

`func (o *ExternalPackage) GetDelegationOrigin() string`

GetDelegationOrigin returns the DelegationOrigin field if non-nil, zero value otherwise.

### GetDelegationOriginOk

`func (o *ExternalPackage) GetDelegationOriginOk() (*string, bool)`

GetDelegationOriginOk returns a tuple with the DelegationOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegationOrigin

`func (o *ExternalPackage) SetDelegationOrigin(v string)`

SetDelegationOrigin sets DelegationOrigin field to given value.


### GetVersion

`func (o *ExternalPackage) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ExternalPackage) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ExternalPackage) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetPackageId

`func (o *ExternalPackage) GetPackageId() string`

GetPackageId returns the PackageId field if non-nil, zero value otherwise.

### GetPackageIdOk

`func (o *ExternalPackage) GetPackageIdOk() (*string, bool)`

GetPackageIdOk returns a tuple with the PackageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageId

`func (o *ExternalPackage) SetPackageId(v string)`

SetPackageId sets PackageId field to given value.


### GetSize

`func (o *ExternalPackage) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ExternalPackage) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ExternalPackage) SetSize(v int64)`

SetSize sets Size field to given value.


### GetHashes

`func (o *ExternalPackage) GetHashes() map[string]string`

GetHashes returns the Hashes field if non-nil, zero value otherwise.

### GetHashesOk

`func (o *ExternalPackage) GetHashesOk() (*map[string]string, bool)`

GetHashesOk returns a tuple with the Hashes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashes

`func (o *ExternalPackage) SetHashes(v map[string]string)`

SetHashes sets Hashes field to given value.


### GetPkgType

`func (o *ExternalPackage) GetPkgType() string`

GetPkgType returns the PkgType field if non-nil, zero value otherwise.

### GetPkgTypeOk

`func (o *ExternalPackage) GetPkgTypeOk() (*string, bool)`

GetPkgTypeOk returns a tuple with the PkgType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkgType

`func (o *ExternalPackage) SetPkgType(v string)`

SetPkgType sets PkgType field to given value.

### HasPkgType

`func (o *ExternalPackage) HasPkgType() bool`

HasPkgType returns a boolean if a field has been set.

### SetPkgTypeNil

`func (o *ExternalPackage) SetPkgTypeNil(b bool)`

 SetPkgTypeNil sets the value for PkgType to be an explicit nil

### UnsetPkgType
`func (o *ExternalPackage) UnsetPkgType()`

UnsetPkgType ensures that no value is present for PkgType, not even an explicit nil
### GetHardwareIds

`func (o *ExternalPackage) GetHardwareIds() []string`

GetHardwareIds returns the HardwareIds field if non-nil, zero value otherwise.

### GetHardwareIdsOk

`func (o *ExternalPackage) GetHardwareIdsOk() (*[]string, bool)`

GetHardwareIdsOk returns a tuple with the HardwareIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareIds

`func (o *ExternalPackage) SetHardwareIds(v []string)`

SetHardwareIds sets HardwareIds field to given value.

### HasHardwareIds

`func (o *ExternalPackage) HasHardwareIds() bool`

HasHardwareIds returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ExternalPackage) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ExternalPackage) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ExternalPackage) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ExternalPackage) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *ExternalPackage) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *ExternalPackage) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUri

`func (o *ExternalPackage) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ExternalPackage) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ExternalPackage) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *ExternalPackage) HasUri() bool`

HasUri returns a boolean if a field has been set.

### SetUriNil

`func (o *ExternalPackage) SetUriNil(b bool)`

 SetUriNil sets the value for Uri to be an explicit nil

### UnsetUri
`func (o *ExternalPackage) UnsetUri()`

UnsetUri ensures that no value is present for Uri, not even an explicit nil
### GetProprietaryMeta

`func (o *ExternalPackage) GetProprietaryMeta() string`

GetProprietaryMeta returns the ProprietaryMeta field if non-nil, zero value otherwise.

### GetProprietaryMetaOk

`func (o *ExternalPackage) GetProprietaryMetaOk() (*string, bool)`

GetProprietaryMetaOk returns a tuple with the ProprietaryMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProprietaryMeta

`func (o *ExternalPackage) SetProprietaryMeta(v string)`

SetProprietaryMeta sets ProprietaryMeta field to given value.

### HasProprietaryMeta

`func (o *ExternalPackage) HasProprietaryMeta() bool`

HasProprietaryMeta returns a boolean if a field has been set.

### SetProprietaryMetaNil

`func (o *ExternalPackage) SetProprietaryMetaNil(b bool)`

 SetProprietaryMetaNil sets the value for ProprietaryMeta to be an explicit nil

### UnsetProprietaryMeta
`func (o *ExternalPackage) UnsetProprietaryMeta()`

UnsetProprietaryMeta ensures that no value is present for ProprietaryMeta, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


