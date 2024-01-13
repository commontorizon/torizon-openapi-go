# Package

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Version** | **string** |  | 
**PackageId** | **string** |  | 
**Size** | **int64** |  | 
**Hashes** | **map[string]string** |  | 
**PkgType** | Pointer to **string** |  | [optional] 
**HardwareIds** | Pointer to **[]string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Uri** | Pointer to **string** |  | [optional] 
**ProprietaryMeta** | Pointer to **string** |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 

## Methods

### NewPackage

`func NewPackage(name string, version string, packageId string, size int64, hashes map[string]string, ) *Package`

NewPackage instantiates a new Package object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackageWithDefaults

`func NewPackageWithDefaults() *Package`

NewPackageWithDefaults instantiates a new Package object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Package) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Package) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Package) SetName(v string)`

SetName sets Name field to given value.


### GetVersion

`func (o *Package) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Package) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Package) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetPackageId

`func (o *Package) GetPackageId() string`

GetPackageId returns the PackageId field if non-nil, zero value otherwise.

### GetPackageIdOk

`func (o *Package) GetPackageIdOk() (*string, bool)`

GetPackageIdOk returns a tuple with the PackageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageId

`func (o *Package) SetPackageId(v string)`

SetPackageId sets PackageId field to given value.


### GetSize

`func (o *Package) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Package) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Package) SetSize(v int64)`

SetSize sets Size field to given value.


### GetHashes

`func (o *Package) GetHashes() map[string]string`

GetHashes returns the Hashes field if non-nil, zero value otherwise.

### GetHashesOk

`func (o *Package) GetHashesOk() (*map[string]string, bool)`

GetHashesOk returns a tuple with the Hashes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashes

`func (o *Package) SetHashes(v map[string]string)`

SetHashes sets Hashes field to given value.


### GetPkgType

`func (o *Package) GetPkgType() string`

GetPkgType returns the PkgType field if non-nil, zero value otherwise.

### GetPkgTypeOk

`func (o *Package) GetPkgTypeOk() (*string, bool)`

GetPkgTypeOk returns a tuple with the PkgType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkgType

`func (o *Package) SetPkgType(v string)`

SetPkgType sets PkgType field to given value.

### HasPkgType

`func (o *Package) HasPkgType() bool`

HasPkgType returns a boolean if a field has been set.

### GetHardwareIds

`func (o *Package) GetHardwareIds() []string`

GetHardwareIds returns the HardwareIds field if non-nil, zero value otherwise.

### GetHardwareIdsOk

`func (o *Package) GetHardwareIdsOk() (*[]string, bool)`

GetHardwareIdsOk returns a tuple with the HardwareIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareIds

`func (o *Package) SetHardwareIds(v []string)`

SetHardwareIds sets HardwareIds field to given value.

### HasHardwareIds

`func (o *Package) HasHardwareIds() bool`

HasHardwareIds returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Package) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Package) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Package) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Package) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUri

`func (o *Package) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *Package) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *Package) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *Package) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetProprietaryMeta

`func (o *Package) GetProprietaryMeta() string`

GetProprietaryMeta returns the ProprietaryMeta field if non-nil, zero value otherwise.

### GetProprietaryMetaOk

`func (o *Package) GetProprietaryMetaOk() (*string, bool)`

GetProprietaryMetaOk returns a tuple with the ProprietaryMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProprietaryMeta

`func (o *Package) SetProprietaryMeta(v string)`

SetProprietaryMeta sets ProprietaryMeta field to given value.

### HasProprietaryMeta

`func (o *Package) HasProprietaryMeta() bool`

HasProprietaryMeta returns a boolean if a field has been set.

### GetComment

`func (o *Package) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Package) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Package) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Package) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


