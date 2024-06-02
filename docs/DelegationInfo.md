# DelegationInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastFetched** | Pointer to **NullableTime** |  | [optional] 
**RemoteUri** | Pointer to **NullableString** |  | [optional] 
**FriendlyName** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewDelegationInfo

`func NewDelegationInfo() *DelegationInfo`

NewDelegationInfo instantiates a new DelegationInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDelegationInfoWithDefaults

`func NewDelegationInfoWithDefaults() *DelegationInfo`

NewDelegationInfoWithDefaults instantiates a new DelegationInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastFetched

`func (o *DelegationInfo) GetLastFetched() time.Time`

GetLastFetched returns the LastFetched field if non-nil, zero value otherwise.

### GetLastFetchedOk

`func (o *DelegationInfo) GetLastFetchedOk() (*time.Time, bool)`

GetLastFetchedOk returns a tuple with the LastFetched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFetched

`func (o *DelegationInfo) SetLastFetched(v time.Time)`

SetLastFetched sets LastFetched field to given value.

### HasLastFetched

`func (o *DelegationInfo) HasLastFetched() bool`

HasLastFetched returns a boolean if a field has been set.

### SetLastFetchedNil

`func (o *DelegationInfo) SetLastFetchedNil(b bool)`

 SetLastFetchedNil sets the value for LastFetched to be an explicit nil

### UnsetLastFetched
`func (o *DelegationInfo) UnsetLastFetched()`

UnsetLastFetched ensures that no value is present for LastFetched, not even an explicit nil
### GetRemoteUri

`func (o *DelegationInfo) GetRemoteUri() string`

GetRemoteUri returns the RemoteUri field if non-nil, zero value otherwise.

### GetRemoteUriOk

`func (o *DelegationInfo) GetRemoteUriOk() (*string, bool)`

GetRemoteUriOk returns a tuple with the RemoteUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteUri

`func (o *DelegationInfo) SetRemoteUri(v string)`

SetRemoteUri sets RemoteUri field to given value.

### HasRemoteUri

`func (o *DelegationInfo) HasRemoteUri() bool`

HasRemoteUri returns a boolean if a field has been set.

### SetRemoteUriNil

`func (o *DelegationInfo) SetRemoteUriNil(b bool)`

 SetRemoteUriNil sets the value for RemoteUri to be an explicit nil

### UnsetRemoteUri
`func (o *DelegationInfo) UnsetRemoteUri()`

UnsetRemoteUri ensures that no value is present for RemoteUri, not even an explicit nil
### GetFriendlyName

`func (o *DelegationInfo) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *DelegationInfo) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *DelegationInfo) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *DelegationInfo) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### SetFriendlyNameNil

`func (o *DelegationInfo) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *DelegationInfo) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


