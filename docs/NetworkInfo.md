# NetworkInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceUuid** | **string** |  | 
**LocalIpV4** | Pointer to **string** |  | [optional] 
**Hostname** | Pointer to **string** |  | [optional] 
**MacAddress** | Pointer to **string** |  | [optional] 

## Methods

### NewNetworkInfo

`func NewNetworkInfo(deviceUuid string, ) *NetworkInfo`

NewNetworkInfo instantiates a new NetworkInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkInfoWithDefaults

`func NewNetworkInfoWithDefaults() *NetworkInfo`

NewNetworkInfoWithDefaults instantiates a new NetworkInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceUuid

`func (o *NetworkInfo) GetDeviceUuid() string`

GetDeviceUuid returns the DeviceUuid field if non-nil, zero value otherwise.

### GetDeviceUuidOk

`func (o *NetworkInfo) GetDeviceUuidOk() (*string, bool)`

GetDeviceUuidOk returns a tuple with the DeviceUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceUuid

`func (o *NetworkInfo) SetDeviceUuid(v string)`

SetDeviceUuid sets DeviceUuid field to given value.


### GetLocalIpV4

`func (o *NetworkInfo) GetLocalIpV4() string`

GetLocalIpV4 returns the LocalIpV4 field if non-nil, zero value otherwise.

### GetLocalIpV4Ok

`func (o *NetworkInfo) GetLocalIpV4Ok() (*string, bool)`

GetLocalIpV4Ok returns a tuple with the LocalIpV4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalIpV4

`func (o *NetworkInfo) SetLocalIpV4(v string)`

SetLocalIpV4 sets LocalIpV4 field to given value.

### HasLocalIpV4

`func (o *NetworkInfo) HasLocalIpV4() bool`

HasLocalIpV4 returns a boolean if a field has been set.

### GetHostname

`func (o *NetworkInfo) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *NetworkInfo) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *NetworkInfo) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *NetworkInfo) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetMacAddress

`func (o *NetworkInfo) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *NetworkInfo) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *NetworkInfo) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *NetworkInfo) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


