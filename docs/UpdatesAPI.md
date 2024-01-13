# \UpdatesAPI

All URIs are relative to *https://app.torizon.io/api/v2beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLockboxDetails**](UpdatesAPI.md#GetLockboxDetails) | **Get** /lockbox-details | List all existing lockboxes on the repository, and their detailed contents
[**GetLockboxes**](UpdatesAPI.md#GetLockboxes) | **Get** /lockboxes | List all existing lockboxes on the repository
[**GetLockboxesLockboxName**](UpdatesAPI.md#GetLockboxesLockboxName) | **Get** /lockboxes/{lockbox_name} | Get the raw Uptane metadata for a lockbox
[**PatchUpdates**](UpdatesAPI.md#PatchUpdates) | **Patch** /updates | Cancel a pending update for one or more devices
[**PostLockboxesLockboxName**](UpdatesAPI.md#PostLockboxesLockboxName) | **Post** /lockboxes/{lockbox_name} | Define a new lockbox, or update an existing one
[**PostUpdates**](UpdatesAPI.md#PostUpdates) | **Post** /updates | Launch an update to one or more devices or fleets



## GetLockboxDetails

> map[string]JsonSignedPayload GetLockboxDetails(ctx).Execute()

List all existing lockboxes on the repository, and their detailed contents



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
	resp, r, err := apiClient.UpdatesAPI.GetLockboxDetails(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UpdatesAPI.GetLockboxDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLockboxDetails`: map[string]JsonSignedPayload
	fmt.Fprintf(os.Stdout, "Response from `UpdatesAPI.GetLockboxDetails`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLockboxDetailsRequest struct via the builder pattern


### Return type

[**map[string]JsonSignedPayload**](JsonSignedPayload.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLockboxes

> []string GetLockboxes(ctx).Execute()

List all existing lockboxes on the repository



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
	resp, r, err := apiClient.UpdatesAPI.GetLockboxes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UpdatesAPI.GetLockboxes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLockboxes`: []string
	fmt.Fprintf(os.Stdout, "Response from `UpdatesAPI.GetLockboxes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLockboxesRequest struct via the builder pattern


### Return type

**[]string**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLockboxesLockboxName

> JsonSignedPayload GetLockboxesLockboxName(ctx, lockboxName).Version(version).Execute()

Get the raw Uptane metadata for a lockbox



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
	lockboxName := "lockboxName_example" // string | 
	version := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UpdatesAPI.GetLockboxesLockboxName(context.Background(), lockboxName).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UpdatesAPI.GetLockboxesLockboxName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLockboxesLockboxName`: JsonSignedPayload
	fmt.Fprintf(os.Stdout, "Response from `UpdatesAPI.GetLockboxesLockboxName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lockboxName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLockboxesLockboxNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **int32** |  | 

### Return type

[**JsonSignedPayload**](JsonSignedPayload.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchUpdates

> []string PatchUpdates(ctx).RequestBody(requestBody).Execute()

Cancel a pending update for one or more devices



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
	requestBody := []string{"Property_example"} // []string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UpdatesAPI.PatchUpdates(context.Background()).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UpdatesAPI.PatchUpdates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchUpdates`: []string
	fmt.Fprintf(os.Stdout, "Response from `UpdatesAPI.PatchUpdates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPatchUpdatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **[]string** |  | 

### Return type

**[]string**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostLockboxesLockboxName

> PostLockboxesLockboxName(ctx, lockboxName).CreateLockboxRequest(createLockboxRequest).Execute()

Define a new lockbox, or update an existing one



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
	lockboxName := "lockboxName_example" // string | 
	createLockboxRequest := *openapiclient.NewCreateLockboxRequest() // CreateLockboxRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UpdatesAPI.PostLockboxesLockboxName(context.Background(), lockboxName).CreateLockboxRequest(createLockboxRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UpdatesAPI.PostLockboxesLockboxName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lockboxName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostLockboxesLockboxNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createLockboxRequest** | [**CreateLockboxRequest**](CreateLockboxRequest.md) |  | 

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


## PostUpdates

> UpdateCreateResult PostUpdates(ctx).UpdateRequest(updateRequest).Execute()

Launch an update to one or more devices or fleets



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
	updateRequest := *openapiclient.NewUpdateRequest() // UpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UpdatesAPI.PostUpdates(context.Background()).UpdateRequest(updateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UpdatesAPI.PostUpdates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostUpdates`: UpdateCreateResult
	fmt.Fprintf(os.Stdout, "Response from `UpdatesAPI.PostUpdates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostUpdatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) |  | 

### Return type

[**UpdateCreateResult**](UpdateCreateResult.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

