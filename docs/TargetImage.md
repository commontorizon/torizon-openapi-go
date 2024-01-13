# TargetImage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Image** | [**Image**](Image.md) |  | 
**Uri** | Pointer to **string** |  | [optional] 
**CreatedAt** | **time.Time** |  | 

## Methods

### NewTargetImage

`func NewTargetImage(image Image, createdAt time.Time, ) *TargetImage`

NewTargetImage instantiates a new TargetImage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetImageWithDefaults

`func NewTargetImageWithDefaults() *TargetImage`

NewTargetImageWithDefaults instantiates a new TargetImage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImage

`func (o *TargetImage) GetImage() Image`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *TargetImage) GetImageOk() (*Image, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *TargetImage) SetImage(v Image)`

SetImage sets Image field to given value.


### GetUri

`func (o *TargetImage) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *TargetImage) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *TargetImage) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *TargetImage) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TargetImage) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TargetImage) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TargetImage) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


