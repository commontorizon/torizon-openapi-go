# PublicKeys

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Keys** | [**map[string]KeyData**](KeyData.md) |  | 

## Methods

### NewPublicKeys

`func NewPublicKeys(keys map[string]KeyData, ) *PublicKeys`

NewPublicKeys instantiates a new PublicKeys object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicKeysWithDefaults

`func NewPublicKeysWithDefaults() *PublicKeys`

NewPublicKeysWithDefaults instantiates a new PublicKeys object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeys

`func (o *PublicKeys) GetKeys() map[string]KeyData`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *PublicKeys) GetKeysOk() (*map[string]KeyData, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *PublicKeys) SetKeys(v map[string]KeyData)`

SetKeys sets Keys field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


