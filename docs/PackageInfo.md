# PackageInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PackageName** | **string** |  | 
**PackageVersion** | **string** |  | 
**Checksum** | **string** |  | 

## Methods

### NewPackageInfo

`func NewPackageInfo(packageName string, packageVersion string, checksum string, ) *PackageInfo`

NewPackageInfo instantiates a new PackageInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackageInfoWithDefaults

`func NewPackageInfoWithDefaults() *PackageInfo`

NewPackageInfoWithDefaults instantiates a new PackageInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPackageName

`func (o *PackageInfo) GetPackageName() string`

GetPackageName returns the PackageName field if non-nil, zero value otherwise.

### GetPackageNameOk

`func (o *PackageInfo) GetPackageNameOk() (*string, bool)`

GetPackageNameOk returns a tuple with the PackageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageName

`func (o *PackageInfo) SetPackageName(v string)`

SetPackageName sets PackageName field to given value.


### GetPackageVersion

`func (o *PackageInfo) GetPackageVersion() string`

GetPackageVersion returns the PackageVersion field if non-nil, zero value otherwise.

### GetPackageVersionOk

`func (o *PackageInfo) GetPackageVersionOk() (*string, bool)`

GetPackageVersionOk returns a tuple with the PackageVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageVersion

`func (o *PackageInfo) SetPackageVersion(v string)`

SetPackageVersion sets PackageVersion field to given value.


### GetChecksum

`func (o *PackageInfo) GetChecksum() string`

GetChecksum returns the Checksum field if non-nil, zero value otherwise.

### GetChecksumOk

`func (o *PackageInfo) GetChecksumOk() (*string, bool)`

GetChecksumOk returns a tuple with the Checksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksum

`func (o *PackageInfo) SetChecksum(v string)`

SetChecksum sets Checksum field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


