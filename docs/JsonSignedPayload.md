# JsonSignedPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Signatures** | Pointer to [**[]ClientSignature**](ClientSignature.md) |  | [optional] 
**Signed** | **interface{}** |  | 

## Methods

### NewJsonSignedPayload

`func NewJsonSignedPayload(signed interface{}, ) *JsonSignedPayload`

NewJsonSignedPayload instantiates a new JsonSignedPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJsonSignedPayloadWithDefaults

`func NewJsonSignedPayloadWithDefaults() *JsonSignedPayload`

NewJsonSignedPayloadWithDefaults instantiates a new JsonSignedPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSignatures

`func (o *JsonSignedPayload) GetSignatures() []ClientSignature`

GetSignatures returns the Signatures field if non-nil, zero value otherwise.

### GetSignaturesOk

`func (o *JsonSignedPayload) GetSignaturesOk() (*[]ClientSignature, bool)`

GetSignaturesOk returns a tuple with the Signatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatures

`func (o *JsonSignedPayload) SetSignatures(v []ClientSignature)`

SetSignatures sets Signatures field to given value.

### HasSignatures

`func (o *JsonSignedPayload) HasSignatures() bool`

HasSignatures returns a boolean if a field has been set.

### GetSigned

`func (o *JsonSignedPayload) GetSigned() interface{}`

GetSigned returns the Signed field if non-nil, zero value otherwise.

### GetSignedOk

`func (o *JsonSignedPayload) GetSignedOk() (*interface{}, bool)`

GetSignedOk returns a tuple with the Signed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigned

`func (o *JsonSignedPayload) SetSigned(v interface{})`

SetSigned sets Signed field to given value.


### SetSignedNil

`func (o *JsonSignedPayload) SetSignedNil(b bool)`

 SetSignedNil sets the value for Signed to be an explicit nil

### UnsetSigned
`func (o *JsonSignedPayload) UnsetSigned()`

UnsetSigned ensures that no value is present for Signed, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


