# Image

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filepath** | **string** |  | 
**Fileinfo** | [**FileInfo**](FileInfo.md) |  | 

## Methods

### NewImage

`func NewImage(filepath string, fileinfo FileInfo, ) *Image`

NewImage instantiates a new Image object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageWithDefaults

`func NewImageWithDefaults() *Image`

NewImageWithDefaults instantiates a new Image object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilepath

`func (o *Image) GetFilepath() string`

GetFilepath returns the Filepath field if non-nil, zero value otherwise.

### GetFilepathOk

`func (o *Image) GetFilepathOk() (*string, bool)`

GetFilepathOk returns a tuple with the Filepath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilepath

`func (o *Image) SetFilepath(v string)`

SetFilepath sets Filepath field to given value.


### GetFileinfo

`func (o *Image) GetFileinfo() FileInfo`

GetFileinfo returns the Fileinfo field if non-nil, zero value otherwise.

### GetFileinfoOk

`func (o *Image) GetFileinfoOk() (*FileInfo, bool)`

GetFileinfoOk returns a tuple with the Fileinfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileinfo

`func (o *Image) SetFileinfo(v FileInfo)`

SetFileinfo sets Fileinfo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


