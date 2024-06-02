# \RemoteAccessAPI

All URIs are relative to *https://app.torizon.io/api/v2beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteRemoteAccessDeviceDeviceuuidSessions**](RemoteAccessAPI.md#DeleteRemoteAccessDeviceDeviceuuidSessions) | **Delete** /remote-access/device/{deviceUuid}/sessions | Delete a remote access session for a device
[**GetRemoteAccessDeviceDeviceuuid**](RemoteAccessAPI.md#GetRemoteAccessDeviceDeviceuuid) | **Get** /remote-access/device/{deviceUuid} | Fetch remote access info for a device
[**GetRemoteAccessDeviceDeviceuuidSessions**](RemoteAccessAPI.md#GetRemoteAccessDeviceDeviceuuidSessions) | **Get** /remote-access/device/{deviceUuid}/sessions | get a remote access session for a device
[**GetRemoteAccessUserPublicKeys**](RemoteAccessAPI.md#GetRemoteAccessUserPublicKeys) | **Get** /remote-access/user/public-keys | Fetch all remote-access public keys registered to a user
[**GetRemoteAccessUserSessions**](RemoteAccessAPI.md#GetRemoteAccessUserSessions) | **Get** /remote-access/user/sessions | Fetch all sessions (and their related deviceId) for a user
[**PostRemoteAccessDeviceDeviceuuidSessions**](RemoteAccessAPI.md#PostRemoteAccessDeviceDeviceuuidSessions) | **Post** /remote-access/device/{deviceUuid}/sessions | Create new remote access session for device



## DeleteRemoteAccessDeviceDeviceuuidSessions

> DeleteRemoteAccessDeviceDeviceuuidSessions(ctx, deviceUuid).Execute()

Delete a remote access session for a device



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/commontorizon/torizon-openapi-go"
)

func main() {
	deviceUuid := "deviceUuid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RemoteAccessAPI.DeleteRemoteAccessDeviceDeviceuuidSessions(context.Background(), deviceUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemoteAccessAPI.DeleteRemoteAccessDeviceDeviceuuidSessions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRemoteAccessDeviceDeviceuuidSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRemoteAccessDeviceDeviceuuid

> DeviceInfo GetRemoteAccessDeviceDeviceuuid(ctx, deviceUuid).Execute()

Fetch remote access info for a device



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/commontorizon/torizon-openapi-go"
)

func main() {
	deviceUuid := "deviceUuid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemoteAccessAPI.GetRemoteAccessDeviceDeviceuuid(context.Background(), deviceUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemoteAccessAPI.GetRemoteAccessDeviceDeviceuuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRemoteAccessDeviceDeviceuuid`: DeviceInfo
	fmt.Fprintf(os.Stdout, "Response from `RemoteAccessAPI.GetRemoteAccessDeviceDeviceuuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRemoteAccessDeviceDeviceuuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeviceInfo**](DeviceInfo.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRemoteAccessDeviceDeviceuuidSessions

> DeviceSession GetRemoteAccessDeviceDeviceuuidSessions(ctx, deviceUuid).Execute()

get a remote access session for a device



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/commontorizon/torizon-openapi-go"
)

func main() {
	deviceUuid := "deviceUuid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemoteAccessAPI.GetRemoteAccessDeviceDeviceuuidSessions(context.Background(), deviceUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemoteAccessAPI.GetRemoteAccessDeviceDeviceuuidSessions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRemoteAccessDeviceDeviceuuidSessions`: DeviceSession
	fmt.Fprintf(os.Stdout, "Response from `RemoteAccessAPI.GetRemoteAccessDeviceDeviceuuidSessions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRemoteAccessDeviceDeviceuuidSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeviceSession**](DeviceSession.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRemoteAccessUserPublicKeys

> PublicKeys GetRemoteAccessUserPublicKeys(ctx).Execute()

Fetch all remote-access public keys registered to a user



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/commontorizon/torizon-openapi-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemoteAccessAPI.GetRemoteAccessUserPublicKeys(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemoteAccessAPI.GetRemoteAccessUserPublicKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRemoteAccessUserPublicKeys`: PublicKeys
	fmt.Fprintf(os.Stdout, "Response from `RemoteAccessAPI.GetRemoteAccessUserPublicKeys`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRemoteAccessUserPublicKeysRequest struct via the builder pattern


### Return type

[**PublicKeys**](PublicKeys.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRemoteAccessUserSessions

> []UserSession GetRemoteAccessUserSessions(ctx).Execute()

Fetch all sessions (and their related deviceId) for a user



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/commontorizon/torizon-openapi-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemoteAccessAPI.GetRemoteAccessUserSessions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemoteAccessAPI.GetRemoteAccessUserSessions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRemoteAccessUserSessions`: []UserSession
	fmt.Fprintf(os.Stdout, "Response from `RemoteAccessAPI.GetRemoteAccessUserSessions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRemoteAccessUserSessionsRequest struct via the builder pattern


### Return type

[**[]UserSession**](UserSession.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostRemoteAccessDeviceDeviceuuidSessions

> PostRemoteAccessDeviceDeviceuuidSessions(ctx, deviceUuid).CreateSessionRequest(createSessionRequest).Execute()

Create new remote access session for device



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/commontorizon/torizon-openapi-go"
)

func main() {
	deviceUuid := "deviceUuid_example" // string | 
	createSessionRequest := *openapiclient.NewCreateSessionRequest("SessionDuration_example") // CreateSessionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RemoteAccessAPI.PostRemoteAccessDeviceDeviceuuidSessions(context.Background(), deviceUuid).CreateSessionRequest(createSessionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemoteAccessAPI.PostRemoteAccessDeviceDeviceuuidSessions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostRemoteAccessDeviceDeviceuuidSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createSessionRequest** | [**CreateSessionRequest**](CreateSessionRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

