# UserSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceId** | **string** |  | 
**Session** | [**DeviceSession**](DeviceSession.md) |  | 

## Methods

### NewUserSession

`func NewUserSession(deviceId string, session DeviceSession, ) *UserSession`

NewUserSession instantiates a new UserSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSessionWithDefaults

`func NewUserSessionWithDefaults() *UserSession`

NewUserSessionWithDefaults instantiates a new UserSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceId

`func (o *UserSession) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *UserSession) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *UserSession) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.


### GetSession

`func (o *UserSession) GetSession() DeviceSession`

GetSession returns the Session field if non-nil, zero value otherwise.

### GetSessionOk

`func (o *UserSession) GetSessionOk() (*DeviceSession, bool)`

GetSessionOk returns a tuple with the Session field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSession

`func (o *UserSession) SetSession(v DeviceSession)`

SetSession sets Session field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


