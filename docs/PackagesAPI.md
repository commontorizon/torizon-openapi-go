# \PackagesAPI

All URIs are relative to *https://app.torizon.io/api/v2beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePackagesPackageid**](PackagesAPI.md#DeletePackagesPackageid) | **Delete** /packages/{packageId} | Delete a package
[**GetPackages**](PackagesAPI.md#GetPackages) | **Get** /packages | Retrieve metadata about packages in your repository
[**GetPackagesExternal**](PackagesAPI.md#GetPackagesExternal) | **Get** /packages_external | Retrieve metadata about packages in your repository from other sources
[**GetPackagesExternalInfo**](PackagesAPI.md#GetPackagesExternalInfo) | **Get** /packages_external/info | Fetch information about external package sources
[**GetPackagesExternalRefreshSourceFileName**](PackagesAPI.md#GetPackagesExternalRefreshSourceFileName) | **Get** /packages_external/refresh/{source_file_name} | Refresh metadata from an external package source
[**PatchPackagesPackageid**](PackagesAPI.md#PatchPackagesPackageid) | **Patch** /packages/{packageId} | Edit metadata about a package
[**PostPackages**](PackagesAPI.md#PostPackages) | **Post** /packages | Upload a new package



## DeletePackagesPackageid

> DeletePackagesPackageid(ctx, packageId).Execute()

Delete a package



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
	packageId := "packageId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PackagesAPI.DeletePackagesPackageid(context.Background(), packageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PackagesAPI.DeletePackagesPackageid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**packageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePackagesPackageidRequest struct via the builder pattern


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


## GetPackages

> PaginationResultPackage GetPackages(ctx).Offset(offset).Limit(limit).IdContains(idContains).Execute()

Retrieve metadata about packages in your repository



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
	idContains := "idContains_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PackagesAPI.GetPackages(context.Background()).Offset(offset).Limit(limit).IdContains(idContains).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PackagesAPI.GetPackages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPackages`: PaginationResultPackage
	fmt.Fprintf(os.Stdout, "Response from `PackagesAPI.GetPackages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPackagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int64** |  | 
 **limit** | **int64** |  | 
 **idContains** | **string** |  | 

### Return type

[**PaginationResultPackage**](PaginationResultPackage.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPackagesExternal

> PaginationResultExternalPackage GetPackagesExternal(ctx).Offset(offset).Limit(limit).IdContains(idContains).Execute()

Retrieve metadata about packages in your repository from other sources



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
	idContains := "idContains_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PackagesAPI.GetPackagesExternal(context.Background()).Offset(offset).Limit(limit).IdContains(idContains).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PackagesAPI.GetPackagesExternal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPackagesExternal`: PaginationResultExternalPackage
	fmt.Fprintf(os.Stdout, "Response from `PackagesAPI.GetPackagesExternal`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPackagesExternalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int64** |  | 
 **limit** | **int64** |  | 
 **idContains** | **string** |  | 

### Return type

[**PaginationResultExternalPackage**](PaginationResultExternalPackage.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPackagesExternalInfo

> map[string]DelegationInfo GetPackagesExternalInfo(ctx).Execute()

Fetch information about external package sources



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
	resp, r, err := apiClient.PackagesAPI.GetPackagesExternalInfo(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PackagesAPI.GetPackagesExternalInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPackagesExternalInfo`: map[string]DelegationInfo
	fmt.Fprintf(os.Stdout, "Response from `PackagesAPI.GetPackagesExternalInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPackagesExternalInfoRequest struct via the builder pattern


### Return type

[**map[string]DelegationInfo**](DelegationInfo.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPackagesExternalRefreshSourceFileName

> GetPackagesExternalRefreshSourceFileName(ctx, sourceFileName).Execute()

Refresh metadata from an external package source



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
	sourceFileName := "sourceFileName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PackagesAPI.GetPackagesExternalRefreshSourceFileName(context.Background(), sourceFileName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PackagesAPI.GetPackagesExternalRefreshSourceFileName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceFileName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPackagesExternalRefreshSourceFileNameRequest struct via the builder pattern


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


## PatchPackagesPackageid

> Package PatchPackagesPackageid(ctx, packageId).EditPackage(editPackage).Execute()

Edit metadata about a package



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
	packageId := "packageId_example" // string | 
	editPackage := *openapiclient.NewEditPackage() // EditPackage | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PackagesAPI.PatchPackagesPackageid(context.Background(), packageId).EditPackage(editPackage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PackagesAPI.PatchPackagesPackageid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchPackagesPackageid`: Package
	fmt.Fprintf(os.Stdout, "Response from `PackagesAPI.PatchPackagesPackageid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**packageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchPackagesPackageidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **editPackage** | [**EditPackage**](EditPackage.md) |  | 

### Return type

[**Package**](Package.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPackages

> Package PostPackages(ctx).Name(name).Version(version).TargetFormat(targetFormat).ContentLength(contentLength).Body(body).HardwareId(hardwareId).Execute()

Upload a new package



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
	name := "name_example" // string | 
	version := "version_example" // string | 
	targetFormat := "targetFormat_example" // string | 
	contentLength := int32(56) // int32 | 
	body := os.NewFile(1234, "some_file") // *os.File | 
	hardwareId := []string{"Inner_example"} // []string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PackagesAPI.PostPackages(context.Background()).Name(name).Version(version).TargetFormat(targetFormat).ContentLength(contentLength).Body(body).HardwareId(hardwareId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PackagesAPI.PostPackages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPackages`: Package
	fmt.Fprintf(os.Stdout, "Response from `PackagesAPI.PostPackages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostPackagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** |  | 
 **version** | **string** |  | 
 **targetFormat** | **string** |  | 
 **contentLength** | **int32** |  | 
 **body** | ***os.File** |  | 
 **hardwareId** | **[]string** |  | 

### Return type

[**Package**](Package.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

