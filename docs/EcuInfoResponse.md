# EcuInfoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**HardwareId** | **string** |  | 
**Primary** | **bool** |  | 
**Image** | [**EcuInfoImage**](EcuInfoImage.md) |  | 

## Methods

### NewEcuInfoResponse

`func NewEcuInfoResponse(id string, hardwareId string, primary bool, image EcuInfoImage, ) *EcuInfoResponse`

NewEcuInfoResponse instantiates a new EcuInfoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEcuInfoResponseWithDefaults

`func NewEcuInfoResponseWithDefaults() *EcuInfoResponse`

NewEcuInfoResponseWithDefaults instantiates a new EcuInfoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EcuInfoResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EcuInfoResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EcuInfoResponse) SetId(v string)`

SetId sets Id field to given value.


### GetHardwareId

`func (o *EcuInfoResponse) GetHardwareId() string`

GetHardwareId returns the HardwareId field if non-nil, zero value otherwise.

### GetHardwareIdOk

`func (o *EcuInfoResponse) GetHardwareIdOk() (*string, bool)`

GetHardwareIdOk returns a tuple with the HardwareId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareId

`func (o *EcuInfoResponse) SetHardwareId(v string)`

SetHardwareId sets HardwareId field to given value.


### GetPrimary

`func (o *EcuInfoResponse) GetPrimary() bool`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *EcuInfoResponse) GetPrimaryOk() (*bool, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *EcuInfoResponse) SetPrimary(v bool)`

SetPrimary sets Primary field to given value.


### GetImage

`func (o *EcuInfoResponse) GetImage() EcuInfoImage`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *EcuInfoResponse) GetImageOk() (*EcuInfoImage, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *EcuInfoResponse) SetImage(v EcuInfoImage)`

SetImage sets Image field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


