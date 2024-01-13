# \DevicesAPI

All URIs are relative to *https://app.torizon.io/api/v2beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteDevicesDeviceuuid**](DevicesAPI.md#DeleteDevicesDeviceuuid) | **Delete** /devices/{deviceUuid} | Delete a single device
[**GetDevices**](DevicesAPI.md#GetDevices) | **Get** /devices | Query device information
[**GetDevicesDeviceuuid**](DevicesAPI.md#GetDevicesDeviceuuid) | **Get** /devices/{deviceUuid} | Get detailed information about a single device
[**GetDevicesNameDeviceuuid**](DevicesAPI.md#GetDevicesNameDeviceuuid) | **Get** /devices/name/{deviceUuid} | Get the display name of a single device
[**GetDevicesNetwork**](DevicesAPI.md#GetDevicesNetwork) | **Get** /devices/network | Get network information for many devices
[**GetDevicesNetworkDeviceuuid**](DevicesAPI.md#GetDevicesNetworkDeviceuuid) | **Get** /devices/network/{deviceUuid} | Get network information for a single device
[**GetDevicesNotesDeviceuuid**](DevicesAPI.md#GetDevicesNotesDeviceuuid) | **Get** /devices/notes/{deviceUuid} | Get the device notes for a specific device
[**GetDevicesPackages**](DevicesAPI.md#GetDevicesPackages) | **Get** /devices/packages | Get information about the installed packages for many devices
[**GetDevicesPackagesDeviceuuid**](DevicesAPI.md#GetDevicesPackagesDeviceuuid) | **Get** /devices/packages/{deviceUuid} | Get information about the installed packages for a single device
[**GetDevicesToken**](DevicesAPI.md#GetDevicesToken) | **Get** /devices/token | Retrieve device provisioning token
[**GetDevicesUptaneDeviceuuidAssignment**](DevicesAPI.md#GetDevicesUptaneDeviceuuidAssignment) | **Get** /devices/uptane/{deviceUuid}/assignment | Show detailed information about the currently-assigned update for a single device
[**GetDevicesUptaneDeviceuuidComponents**](DevicesAPI.md#GetDevicesUptaneDeviceuuidComponents) | **Get** /devices/uptane/{deviceUuid}/components | Get a list of the software components reported by a single device
[**PostDevices**](DevicesAPI.md#PostDevices) | **Post** /devices | Manually create a new device
[**PutDevicesNameDeviceuuid**](DevicesAPI.md#PutDevicesNameDeviceuuid) | **Put** /devices/name/{deviceUuid} | Set the display name of a single device
[**PutDevicesNotesDeviceuuid**](DevicesAPI.md#PutDevicesNotesDeviceuuid) | **Put** /devices/notes/{deviceUuid} | Set the device notes for a specific device



## DeleteDevicesDeviceuuid

> DeleteDevicesDeviceuuid(ctx, deviceUuid).Execute()

Delete a single device



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
	r, err := apiClient.DevicesAPI.DeleteDevicesDeviceuuid(context.Background(), deviceUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.DeleteDevicesDeviceuuid``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteDevicesDeviceuuidRequest struct via the builder pattern


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


## GetDevices

> PaginationResultDeviceInfoBasic GetDevices(ctx).Offset(offset).Limit(limit).DeviceUuid(deviceUuid).NameContains(nameContains).SortBy(sortBy).SortDirection(sortDirection).Execute()

Query device information



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
	offset := int64(789) // int64 |  (optional)
	limit := int64(789) // int64 |  (optional)
	deviceUuid := []string{"Inner_example"} // []string |  (optional)
	nameContains := "nameContains_example" // string |  (optional)
	sortBy := openapiclient.DeviceSort("Name") // DeviceSort |  (optional)
	sortDirection := openapiclient.DeviceSortDirection("Asc") // DeviceSortDirection |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.GetDevices(context.Background()).Offset(offset).Limit(limit).DeviceUuid(deviceUuid).NameContains(nameContains).SortBy(sortBy).SortDirection(sortDirection).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.GetDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDevices`: PaginationResultDeviceInfoBasic
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.GetDevices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int64** |  | 
 **limit** | **int64** |  | 
 **deviceUuid** | **[]string** |  | 
 **nameContains** | **string** |  | 
 **sortBy** | [**DeviceSort**](DeviceSort.md) |  | 
 **sortDirection** | [**DeviceSortDirection**](DeviceSortDirection.md) |  | 

### Return type

[**PaginationResultDeviceInfoBasic**](PaginationResultDeviceInfoBasic.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDevicesDeviceuuid

> DeviceInfoExtended GetDevicesDeviceuuid(ctx, deviceUuid).Execute()

Get detailed information about a single device



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
	resp, r, err := apiClient.DevicesAPI.GetDevicesDeviceuuid(context.Background(), deviceUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.GetDevicesDeviceuuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDevicesDeviceuuid`: DeviceInfoExtended
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.GetDevicesDeviceuuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDevicesDeviceuuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeviceInfoExtended**](DeviceInfoExtended.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDevicesNameDeviceuuid

> string GetDevicesNameDeviceuuid(ctx, deviceUuid).Execute()

Get the display name of a single device



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
	resp, r, err := apiClient.DevicesAPI.GetDevicesNameDeviceuuid(context.Background(), deviceUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.GetDevicesNameDeviceuuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDevicesNameDeviceuuid`: string
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.GetDevicesNameDeviceuuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDevicesNameDeviceuuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDevicesNetwork

> PaginationResultNetworkInfo GetDevicesNetwork(ctx).Offset(offset).Limit(limit).DeviceUuid(deviceUuid).Execute()

Get network information for many devices



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
	offset := int64(789) // int64 |  (optional)
	limit := int64(789) // int64 |  (optional)
	deviceUuid := []string{"Inner_example"} // []string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.GetDevicesNetwork(context.Background()).Offset(offset).Limit(limit).DeviceUuid(deviceUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.GetDevicesNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDevicesNetwork`: PaginationResultNetworkInfo
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.GetDevicesNetwork`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDevicesNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int64** |  | 
 **limit** | **int64** |  | 
 **deviceUuid** | **[]string** |  | 

### Return type

[**PaginationResultNetworkInfo**](PaginationResultNetworkInfo.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDevicesNetworkDeviceuuid

> NetworkInfo GetDevicesNetworkDeviceuuid(ctx, deviceUuid).Execute()

Get network information for a single device



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
	resp, r, err := apiClient.DevicesAPI.GetDevicesNetworkDeviceuuid(context.Background(), deviceUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.GetDevicesNetworkDeviceuuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDevicesNetworkDeviceuuid`: NetworkInfo
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.GetDevicesNetworkDeviceuuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDevicesNetworkDeviceuuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NetworkInfo**](NetworkInfo.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDevicesNotesDeviceuuid

> string GetDevicesNotesDeviceuuid(ctx, deviceUuid).Execute()

Get the device notes for a specific device



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
	resp, r, err := apiClient.DevicesAPI.GetDevicesNotesDeviceuuid(context.Background(), deviceUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.GetDevicesNotesDeviceuuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDevicesNotesDeviceuuid`: string
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.GetDevicesNotesDeviceuuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDevicesNotesDeviceuuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDevicesPackages

> PaginationResultDevicePackages GetDevicesPackages(ctx).Offset(offset).Limit(limit).DeviceUuid(deviceUuid).NameContains(nameContains).Execute()

Get information about the installed packages for many devices



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
	offset := int64(789) // int64 |  (optional)
	limit := int64(789) // int64 |  (optional)
	deviceUuid := []string{"Inner_example"} // []string |  (optional)
	nameContains := "nameContains_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.GetDevicesPackages(context.Background()).Offset(offset).Limit(limit).DeviceUuid(deviceUuid).NameContains(nameContains).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.GetDevicesPackages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDevicesPackages`: PaginationResultDevicePackages
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.GetDevicesPackages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDevicesPackagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int64** |  | 
 **limit** | **int64** |  | 
 **deviceUuid** | **[]string** |  | 
 **nameContains** | **string** |  | 

### Return type

[**PaginationResultDevicePackages**](PaginationResultDevicePackages.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDevicesPackagesDeviceuuid

> DevicePackages GetDevicesPackagesDeviceuuid(ctx, deviceUuid).Execute()

Get information about the installed packages for a single device



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
	resp, r, err := apiClient.DevicesAPI.GetDevicesPackagesDeviceuuid(context.Background(), deviceUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.GetDevicesPackagesDeviceuuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDevicesPackagesDeviceuuid`: DevicePackages
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.GetDevicesPackagesDeviceuuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDevicesPackagesDeviceuuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DevicePackages**](DevicePackages.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDevicesToken

> ProvisionInfo GetDevicesToken(ctx).Execute()

Retrieve device provisioning token



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
	resp, r, err := apiClient.DevicesAPI.GetDevicesToken(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.GetDevicesToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDevicesToken`: ProvisionInfo
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.GetDevicesToken`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDevicesTokenRequest struct via the builder pattern


### Return type

[**ProvisionInfo**](ProvisionInfo.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDevicesUptaneDeviceuuidAssignment

> []QueueResponse GetDevicesUptaneDeviceuuidAssignment(ctx, deviceUuid).Execute()

Show detailed information about the currently-assigned update for a single device



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
	resp, r, err := apiClient.DevicesAPI.GetDevicesUptaneDeviceuuidAssignment(context.Background(), deviceUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.GetDevicesUptaneDeviceuuidAssignment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDevicesUptaneDeviceuuidAssignment`: []QueueResponse
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.GetDevicesUptaneDeviceuuidAssignment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDevicesUptaneDeviceuuidAssignmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]QueueResponse**](QueueResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDevicesUptaneDeviceuuidComponents

> []EcuInfoResponse GetDevicesUptaneDeviceuuidComponents(ctx, deviceUuid).Execute()

Get a list of the software components reported by a single device



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
	resp, r, err := apiClient.DevicesAPI.GetDevicesUptaneDeviceuuidComponents(context.Background(), deviceUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.GetDevicesUptaneDeviceuuidComponents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDevicesUptaneDeviceuuidComponents`: []EcuInfoResponse
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.GetDevicesUptaneDeviceuuidComponents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDevicesUptaneDeviceuuidComponentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]EcuInfoResponse**](EcuInfoResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostDevices

> *os.File PostDevices(ctx).DeviceCreateReq(deviceCreateReq).Execute()

Manually create a new device



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
	deviceCreateReq := *openapiclient.NewDeviceCreateReq("DeviceId_example") // DeviceCreateReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.PostDevices(context.Background()).DeviceCreateReq(deviceCreateReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.PostDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostDevices`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.PostDevices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceCreateReq** | [**DeviceCreateReq**](DeviceCreateReq.md) |  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/octet-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutDevicesNameDeviceuuid

> PutDevicesNameDeviceuuid(ctx, deviceUuid).Body(body).Execute()

Set the display name of a single device



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
	body := "body_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DevicesAPI.PutDevicesNameDeviceuuid(context.Background(), deviceUuid).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.PutDevicesNameDeviceuuid``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPutDevicesNameDeviceuuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** |  | 

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


## PutDevicesNotesDeviceuuid

> PutDevicesNotesDeviceuuid(ctx, deviceUuid).Body(body).Execute()

Set the device notes for a specific device



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
	body := "body_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DevicesAPI.PutDevicesNotesDeviceuuid(context.Background(), deviceUuid).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.PutDevicesNotesDeviceuuid``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPutDevicesNotesDeviceuuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

