# DeviceSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ssh** | [**SshSession**](SshSession.md) |  | 

## Methods

### NewDeviceSession

`func NewDeviceSession(ssh SshSession, ) *DeviceSession`

NewDeviceSession instantiates a new DeviceSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceSessionWithDefaults

`func NewDeviceSessionWithDefaults() *DeviceSession`

NewDeviceSessionWithDefaults instantiates a new DeviceSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSsh

`func (o *DeviceSession) GetSsh() SshSession`

GetSsh returns the Ssh field if non-nil, zero value otherwise.

### GetSshOk

`func (o *DeviceSession) GetSshOk() (*SshSession, bool)`

GetSshOk returns a tuple with the Ssh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsh

`func (o *DeviceSession) SetSsh(v SshSession)`

SetSsh sets Ssh field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


