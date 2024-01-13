# InstalledPackage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Component** | **string** |  | 
**Installed** | [**PackageInfo**](PackageInfo.md) |  | 

## Methods

### NewInstalledPackage

`func NewInstalledPackage(component string, installed PackageInfo, ) *InstalledPackage`

NewInstalledPackage instantiates a new InstalledPackage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstalledPackageWithDefaults

`func NewInstalledPackageWithDefaults() *InstalledPackage`

NewInstalledPackageWithDefaults instantiates a new InstalledPackage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponent

`func (o *InstalledPackage) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *InstalledPackage) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *InstalledPackage) SetComponent(v string)`

SetComponent sets Component field to given value.


### GetInstalled

`func (o *InstalledPackage) GetInstalled() PackageInfo`

GetInstalled returns the Installed field if non-nil, zero value otherwise.

### GetInstalledOk

`func (o *InstalledPackage) GetInstalledOk() (*PackageInfo, bool)`

GetInstalledOk returns a tuple with the Installed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalled

`func (o *InstalledPackage) SetInstalled(v PackageInfo)`

SetInstalled sets Installed field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


