# LastSshSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectedAt** | **time.Time** |  | 
**IpAddress** | **string** |  | 

## Methods

### NewLastSshSession

`func NewLastSshSession(connectedAt time.Time, ipAddress string, ) *LastSshSession`

NewLastSshSession instantiates a new LastSshSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLastSshSessionWithDefaults

`func NewLastSshSessionWithDefaults() *LastSshSession`

NewLastSshSessionWithDefaults instantiates a new LastSshSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectedAt

`func (o *LastSshSession) GetConnectedAt() time.Time`

GetConnectedAt returns the ConnectedAt field if non-nil, zero value otherwise.

### GetConnectedAtOk

`func (o *LastSshSession) GetConnectedAtOk() (*time.Time, bool)`

GetConnectedAtOk returns a tuple with the ConnectedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedAt

`func (o *LastSshSession) SetConnectedAt(v time.Time)`

SetConnectedAt sets ConnectedAt field to given value.


### GetIpAddress

`func (o *LastSshSession) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *LastSshSession) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *LastSshSession) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


