# \DeviceMetricsAPI

All URIs are relative to *https://app.torizon.io/api/v2beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDeviceDataDevicesDeviceuuidMetrics**](DeviceMetricsAPI.md#GetDeviceDataDevicesDeviceuuidMetrics) | **Get** /device-data/devices/{deviceUuid}/metrics | Get metrics data from a single device
[**GetDeviceDataFleetsFleetidMetrics**](DeviceMetricsAPI.md#GetDeviceDataFleetsFleetidMetrics) | **Get** /device-data/fleets/{fleetId}/metrics | Get aggregated metrics data from a fleet of devices
[**GetDeviceDataMetricNames**](DeviceMetricsAPI.md#GetDeviceDataMetricNames) | **Get** /device-data/metric-names | Get the list of metrics available in your repository



## GetDeviceDataDevicesDeviceuuidMetrics

> MetricsResponse GetDeviceDataDevicesDeviceuuidMetrics(ctx, deviceUuid).From(from).To(to).Metric(metric).Resolution(resolution).Execute()

Get metrics data from a single device



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
	from := int64(789) // int64 | 
	to := int64(789) // int64 | 
	metric := []string{"Inner_example"} // []string |  (optional)
	resolution := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceMetricsAPI.GetDeviceDataDevicesDeviceuuidMetrics(context.Background(), deviceUuid).From(from).To(to).Metric(metric).Resolution(resolution).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceMetricsAPI.GetDeviceDataDevicesDeviceuuidMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceDataDevicesDeviceuuidMetrics`: MetricsResponse
	fmt.Fprintf(os.Stdout, "Response from `DeviceMetricsAPI.GetDeviceDataDevicesDeviceuuidMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceDataDevicesDeviceuuidMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **int64** |  | 
 **to** | **int64** |  | 
 **metric** | **[]string** |  | 
 **resolution** | **int32** |  | 

### Return type

[**MetricsResponse**](MetricsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceDataFleetsFleetidMetrics

> MetricsResponse GetDeviceDataFleetsFleetidMetrics(ctx, fleetId).From(from).To(to).Metric(metric).Resolution(resolution).Execute()

Get aggregated metrics data from a fleet of devices



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
	from := int64(789) // int64 | 
	to := int64(789) // int64 | 
	metric := []string{"Inner_example"} // []string |  (optional)
	resolution := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceMetricsAPI.GetDeviceDataFleetsFleetidMetrics(context.Background(), fleetId).From(from).To(to).Metric(metric).Resolution(resolution).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceMetricsAPI.GetDeviceDataFleetsFleetidMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceDataFleetsFleetidMetrics`: MetricsResponse
	fmt.Fprintf(os.Stdout, "Response from `DeviceMetricsAPI.GetDeviceDataFleetsFleetidMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fleetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceDataFleetsFleetidMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **int64** |  | 
 **to** | **int64** |  | 
 **metric** | **[]string** |  | 
 **resolution** | **int32** |  | 

### Return type

[**MetricsResponse**](MetricsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceDataMetricNames

> PaginationResultString GetDeviceDataMetricNames(ctx).From(from).To(to).Execute()

Get the list of metrics available in your repository



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
	from := int64(789) // int64 |  (optional)
	to := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceMetricsAPI.GetDeviceDataMetricNames(context.Background()).From(from).To(to).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceMetricsAPI.GetDeviceDataMetricNames``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceDataMetricNames`: PaginationResultString
	fmt.Fprintf(os.Stdout, "Response from `DeviceMetricsAPI.GetDeviceDataMetricNames`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceDataMetricNamesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **int64** |  | 
 **to** | **int64** |  | 

### Return type

[**PaginationResultString**](PaginationResultString.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

