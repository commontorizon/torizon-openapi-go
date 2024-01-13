# ClientSignature

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Keyid** | **string** |  | 
**Method** | [**SignatureMethod**](SignatureMethod.md) |  | 
**Sig** | **string** |  | 

## Methods

### NewClientSignature

`func NewClientSignature(keyid string, method SignatureMethod, sig string, ) *ClientSignature`

NewClientSignature instantiates a new ClientSignature object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientSignatureWithDefaults

`func NewClientSignatureWithDefaults() *ClientSignature`

NewClientSignatureWithDefaults instantiates a new ClientSignature object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyid

`func (o *ClientSignature) GetKeyid() string`

GetKeyid returns the Keyid field if non-nil, zero value otherwise.

### GetKeyidOk

`func (o *ClientSignature) GetKeyidOk() (*string, bool)`

GetKeyidOk returns a tuple with the Keyid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyid

`func (o *ClientSignature) SetKeyid(v string)`

SetKeyid sets Keyid field to given value.


### GetMethod

`func (o *ClientSignature) GetMethod() SignatureMethod`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *ClientSignature) GetMethodOk() (*SignatureMethod, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *ClientSignature) SetMethod(v SignatureMethod)`

SetMethod sets Method field to given value.


### GetSig

`func (o *ClientSignature) GetSig() string`

GetSig returns the Sig field if non-nil, zero value otherwise.

### GetSigOk

`func (o *ClientSignature) GetSigOk() (*string, bool)`

GetSigOk returns a tuple with the Sig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSig

`func (o *ClientSignature) SetSig(v string)`

SetSig sets Sig field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


