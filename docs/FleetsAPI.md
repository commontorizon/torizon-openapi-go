# \FleetsAPI

All URIs are relative to *https://app.torizon.io/api/v2beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteFleetsFleetid**](FleetsAPI.md#DeleteFleetsFleetid) | **Delete** /fleets/{fleetId} | Delete a fleet
[**DeleteFleetsFleetidDevices**](FleetsAPI.md#DeleteFleetsFleetidDevices) | **Delete** /fleets/{fleetId}/devices | Remove devices from a fleet
[**GetFleets**](FleetsAPI.md#GetFleets) | **Get** /fleets | Get information about all fleets in your repository
[**GetFleetsFleetidDevices**](FleetsAPI.md#GetFleetsFleetidDevices) | **Get** /fleets/{fleetId}/devices | Get information about the devices in a single fleet
[**PostFleets**](FleetsAPI.md#PostFleets) | **Post** /fleets | Create a new fleet
[**PostFleetsFleetidDevices**](FleetsAPI.md#PostFleetsFleetidDevices) | **Post** /fleets/{fleetId}/devices | Add devices to a fleet



## DeleteFleetsFleetid

> DeleteFleetsFleetid(ctx, fleetId).Execute()

Delete a fleet



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
	fleetId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FleetsAPI.DeleteFleetsFleetid(context.Background(), fleetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetsAPI.DeleteFleetsFleetid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fleetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFleetsFleetidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFleetsFleetidDevices

> DeleteFleetsFleetidDevices(ctx, fleetId).RequestBody(requestBody).Execute()

Remove devices from a fleet



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
	fleetId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	requestBody := []string{"Property_example"} // []string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FleetsAPI.DeleteFleetsFleetidDevices(context.Background(), fleetId).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetsAPI.DeleteFleetsFleetidDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fleetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFleetsFleetidDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **[]string** |  | 

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


## GetFleets

> PaginationResultFleet GetFleets(ctx).Offset(offset).Limit(limit).Execute()

Get information about all fleets in your repository



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FleetsAPI.GetFleets(context.Background()).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetsAPI.GetFleets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFleets`: PaginationResultFleet
	fmt.Fprintf(os.Stdout, "Response from `FleetsAPI.GetFleets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFleetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int64** |  | 
 **limit** | **int64** |  | 

### Return type

[**PaginationResultFleet**](PaginationResultFleet.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFleetsFleetidDevices

> PaginationResultDeviceInfoBasic GetFleetsFleetidDevices(ctx, fleetId).Offset(offset).Limit(limit).Execute()

Get information about the devices in a single fleet



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
	fleetId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	offset := int64(789) // int64 |  (optional)
	limit := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FleetsAPI.GetFleetsFleetidDevices(context.Background(), fleetId).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetsAPI.GetFleetsFleetidDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFleetsFleetidDevices`: PaginationResultDeviceInfoBasic
	fmt.Fprintf(os.Stdout, "Response from `FleetsAPI.GetFleetsFleetidDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fleetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFleetsFleetidDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **int64** |  | 
 **limit** | **int64** |  | 

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


## PostFleets

> string PostFleets(ctx).CreateFleet(createFleet).Execute()

Create a new fleet



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
	createFleet := *openapiclient.NewCreateFleet("Name_example", openapiclient.FleetType("static")) // CreateFleet | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FleetsAPI.PostFleets(context.Background()).CreateFleet(createFleet).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetsAPI.PostFleets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostFleets`: string
	fmt.Fprintf(os.Stdout, "Response from `FleetsAPI.PostFleets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostFleetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createFleet** | [**CreateFleet**](CreateFleet.md) |  | 

### Return type

**string**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostFleetsFleetidDevices

> PostFleetsFleetidDevices(ctx, fleetId).RequestBody(requestBody).Execute()

Add devices to a fleet



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
	fleetId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	requestBody := []string{"Property_example"} // []string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FleetsAPI.PostFleetsFleetidDevices(context.Background(), fleetId).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetsAPI.PostFleetsFleetidDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fleetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostFleetsFleetidDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **[]string** |  | 

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

