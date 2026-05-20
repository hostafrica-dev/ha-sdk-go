# \SnapshotsAPI

All URIs are relative to *https://api.hostafrica.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSnapshot**](SnapshotsAPI.md#CreateSnapshot) | **Post** /vps/create-snapshot | 
[**DeleteSnapshot**](SnapshotsAPI.md#DeleteSnapshot) | **Post** /vps/delete-snapshot | 
[**ListSnapshots**](SnapshotsAPI.md#ListSnapshots) | **Post** /vps/list-snapshots | 
[**RollbackSnapshot**](SnapshotsAPI.md#RollbackSnapshot) | **Post** /vps/rollback-snapshot | 
[**UpdateSnapshot**](SnapshotsAPI.md#UpdateSnapshot) | **Post** /vps/update-snapshot | 



## CreateSnapshot

> CreateSnapshotResponseContent CreateSnapshot(ctx).CreateSnapshotRequestContent(createSnapshotRequestContent).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hostafrica/ha-sdk-go"
)

func main() {
	createSnapshotRequestContent := *openapiclient.NewCreateSnapshotRequestContent("ServiceId_example") // CreateSnapshotRequestContent | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SnapshotsAPI.CreateSnapshot(context.Background()).CreateSnapshotRequestContent(createSnapshotRequestContent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsAPI.CreateSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSnapshot`: CreateSnapshotResponseContent
	fmt.Fprintf(os.Stdout, "Response from `SnapshotsAPI.CreateSnapshot`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createSnapshotRequestContent** | [**CreateSnapshotRequestContent**](CreateSnapshotRequestContent.md) |  | 

### Return type

[**CreateSnapshotResponseContent**](CreateSnapshotResponseContent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSnapshot

> DeleteSnapshotResponseContent DeleteSnapshot(ctx).DeleteSnapshotRequestContent(deleteSnapshotRequestContent).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hostafrica/ha-sdk-go"
)

func main() {
	deleteSnapshotRequestContent := *openapiclient.NewDeleteSnapshotRequestContent("ServiceId_example", "SnapshotName_example") // DeleteSnapshotRequestContent | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SnapshotsAPI.DeleteSnapshot(context.Background()).DeleteSnapshotRequestContent(deleteSnapshotRequestContent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsAPI.DeleteSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSnapshot`: DeleteSnapshotResponseContent
	fmt.Fprintf(os.Stdout, "Response from `SnapshotsAPI.DeleteSnapshot`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteSnapshotRequestContent** | [**DeleteSnapshotRequestContent**](DeleteSnapshotRequestContent.md) |  | 

### Return type

[**DeleteSnapshotResponseContent**](DeleteSnapshotResponseContent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSnapshots

> ListSnapshotsResponseContent ListSnapshots(ctx).ListSnapshotsRequestContent(listSnapshotsRequestContent).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hostafrica/ha-sdk-go"
)

func main() {
	listSnapshotsRequestContent := *openapiclient.NewListSnapshotsRequestContent("ServiceId_example") // ListSnapshotsRequestContent | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SnapshotsAPI.ListSnapshots(context.Background()).ListSnapshotsRequestContent(listSnapshotsRequestContent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsAPI.ListSnapshots``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSnapshots`: ListSnapshotsResponseContent
	fmt.Fprintf(os.Stdout, "Response from `SnapshotsAPI.ListSnapshots`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSnapshotsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listSnapshotsRequestContent** | [**ListSnapshotsRequestContent**](ListSnapshotsRequestContent.md) |  | 

### Return type

[**ListSnapshotsResponseContent**](ListSnapshotsResponseContent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RollbackSnapshot

> RollbackSnapshotResponseContent RollbackSnapshot(ctx).RollbackSnapshotRequestContent(rollbackSnapshotRequestContent).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hostafrica/ha-sdk-go"
)

func main() {
	rollbackSnapshotRequestContent := *openapiclient.NewRollbackSnapshotRequestContent("ServiceId_example", "SnapshotName_example") // RollbackSnapshotRequestContent | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SnapshotsAPI.RollbackSnapshot(context.Background()).RollbackSnapshotRequestContent(rollbackSnapshotRequestContent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsAPI.RollbackSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RollbackSnapshot`: RollbackSnapshotResponseContent
	fmt.Fprintf(os.Stdout, "Response from `SnapshotsAPI.RollbackSnapshot`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRollbackSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rollbackSnapshotRequestContent** | [**RollbackSnapshotRequestContent**](RollbackSnapshotRequestContent.md) |  | 

### Return type

[**RollbackSnapshotResponseContent**](RollbackSnapshotResponseContent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSnapshot

> UpdateSnapshotResponseContent UpdateSnapshot(ctx).UpdateSnapshotRequestContent(updateSnapshotRequestContent).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hostafrica/ha-sdk-go"
)

func main() {
	updateSnapshotRequestContent := *openapiclient.NewUpdateSnapshotRequestContent("ServiceId_example", "SnapshotName_example") // UpdateSnapshotRequestContent | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SnapshotsAPI.UpdateSnapshot(context.Background()).UpdateSnapshotRequestContent(updateSnapshotRequestContent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsAPI.UpdateSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSnapshot`: UpdateSnapshotResponseContent
	fmt.Fprintf(os.Stdout, "Response from `SnapshotsAPI.UpdateSnapshot`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateSnapshotRequestContent** | [**UpdateSnapshotRequestContent**](UpdateSnapshotRequestContent.md) |  | 

### Return type

[**UpdateSnapshotResponseContent**](UpdateSnapshotResponseContent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

