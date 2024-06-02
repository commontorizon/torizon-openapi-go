# SshSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorizedPubKeys** | Pointer to **[]string** |  | [optional] 
**ReversePort** | **int32** |  | 
**RaServerUrl** | **string** |  | 
**RaServerSshPubKey** | **string** |  | 
**ExpiresAt** | **time.Time** |  | 

## Methods

### NewSshSession

`func NewSshSession(reversePort int32, raServerUrl string, raServerSshPubKey string, expiresAt time.Time, ) *SshSession`

NewSshSession instantiates a new SshSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSshSessionWithDefaults

`func NewSshSessionWithDefaults() *SshSession`

NewSshSessionWithDefaults instantiates a new SshSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorizedPubKeys

`func (o *SshSession) GetAuthorizedPubKeys() []string`

GetAuthorizedPubKeys returns the AuthorizedPubKeys field if non-nil, zero value otherwise.

### GetAuthorizedPubKeysOk

`func (o *SshSession) GetAuthorizedPubKeysOk() (*[]string, bool)`

GetAuthorizedPubKeysOk returns a tuple with the AuthorizedPubKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedPubKeys

`func (o *SshSession) SetAuthorizedPubKeys(v []string)`

SetAuthorizedPubKeys sets AuthorizedPubKeys field to given value.

### HasAuthorizedPubKeys

`func (o *SshSession) HasAuthorizedPubKeys() bool`

HasAuthorizedPubKeys returns a boolean if a field has been set.

### GetReversePort

`func (o *SshSession) GetReversePort() int32`

GetReversePort returns the ReversePort field if non-nil, zero value otherwise.

### GetReversePortOk

`func (o *SshSession) GetReversePortOk() (*int32, bool)`

GetReversePortOk returns a tuple with the ReversePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReversePort

`func (o *SshSession) SetReversePort(v int32)`

SetReversePort sets ReversePort field to given value.


### GetRaServerUrl

`func (o *SshSession) GetRaServerUrl() string`

GetRaServerUrl returns the RaServerUrl field if non-nil, zero value otherwise.

### GetRaServerUrlOk

`func (o *SshSession) GetRaServerUrlOk() (*string, bool)`

GetRaServerUrlOk returns a tuple with the RaServerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaServerUrl

`func (o *SshSession) SetRaServerUrl(v string)`

SetRaServerUrl sets RaServerUrl field to given value.


### GetRaServerSshPubKey

`func (o *SshSession) GetRaServerSshPubKey() string`

GetRaServerSshPubKey returns the RaServerSshPubKey field if non-nil, zero value otherwise.

### GetRaServerSshPubKeyOk

`func (o *SshSession) GetRaServerSshPubKeyOk() (*string, bool)`

GetRaServerSshPubKeyOk returns a tuple with the RaServerSshPubKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaServerSshPubKey

`func (o *SshSession) SetRaServerSshPubKey(v string)`

SetRaServerSshPubKey sets RaServerSshPubKey field to given value.


### GetExpiresAt

`func (o *SshSession) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *SshSession) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *SshSession) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


