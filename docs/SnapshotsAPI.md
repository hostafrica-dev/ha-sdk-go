# \SnapshotsAPI

All URIs are relative to *https://api.hostafrica.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSnapshot**](SnapshotsAPI.md#CreateSnapshot) | **Post** /vps/create-snapshot | 
[**CreateSnapshotJob**](SnapshotsAPI.md#CreateSnapshotJob) | **Post** /vps/create-snapshot-job | 
[**DeleteSnapshot**](SnapshotsAPI.md#DeleteSnapshot) | **Post** /vps/delete-snapshot | 
[**DeleteSnapshotJob**](SnapshotsAPI.md#DeleteSnapshotJob) | **Post** /vps/delete-snapshot-job | 
[**ListSnapshotJobs**](SnapshotsAPI.md#ListSnapshotJobs) | **Post** /vps/list-snapshot-jobs | 
[**ListSnapshots**](SnapshotsAPI.md#ListSnapshots) | **Post** /vps/list-snapshots | 
[**RollbackSnapshot**](SnapshotsAPI.md#RollbackSnapshot) | **Post** /vps/rollback-snapshot | 
[**UpdateSnapshot**](SnapshotsAPI.md#UpdateSnapshot) | **Post** /vps/update-snapshot | 
[**UpdateSnapshotJob**](SnapshotsAPI.md#UpdateSnapshotJob) | **Post** /vps/update-snapshot-job | 



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
	createSnapshotRequestContent := *openapiclient.NewCreateSnapshotRequestContent("ServiceId_example", "Name_example") // CreateSnapshotRequestContent | 

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


## CreateSnapshotJob

> CreateSnapshotJobResponseContent CreateSnapshotJob(ctx).CreateSnapshotJobRequestContent(createSnapshotJobRequestContent).Execute()





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
	createSnapshotJobRequestContent := *openapiclient.NewCreateSnapshotJobRequestContent("ServiceId_example", "Name_example", openapiclient.SnapshotJobPeriod("hourly")) // CreateSnapshotJobRequestContent | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SnapshotsAPI.CreateSnapshotJob(context.Background()).CreateSnapshotJobRequestContent(createSnapshotJobRequestContent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsAPI.CreateSnapshotJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSnapshotJob`: CreateSnapshotJobResponseContent
	fmt.Fprintf(os.Stdout, "Response from `SnapshotsAPI.CreateSnapshotJob`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSnapshotJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createSnapshotJobRequestContent** | [**CreateSnapshotJobRequestContent**](CreateSnapshotJobRequestContent.md) |  | 

### Return type

[**CreateSnapshotJobResponseContent**](CreateSnapshotJobResponseContent.md)

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


## DeleteSnapshotJob

> DeleteSnapshotJobResponseContent DeleteSnapshotJob(ctx).DeleteSnapshotJobRequestContent(deleteSnapshotJobRequestContent).Execute()





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
	deleteSnapshotJobRequestContent := *openapiclient.NewDeleteSnapshotJobRequestContent("ServiceId_example", "JobId_example") // DeleteSnapshotJobRequestContent | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SnapshotsAPI.DeleteSnapshotJob(context.Background()).DeleteSnapshotJobRequestContent(deleteSnapshotJobRequestContent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsAPI.DeleteSnapshotJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSnapshotJob`: DeleteSnapshotJobResponseContent
	fmt.Fprintf(os.Stdout, "Response from `SnapshotsAPI.DeleteSnapshotJob`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSnapshotJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteSnapshotJobRequestContent** | [**DeleteSnapshotJobRequestContent**](DeleteSnapshotJobRequestContent.md) |  | 

### Return type

[**DeleteSnapshotJobResponseContent**](DeleteSnapshotJobResponseContent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSnapshotJobs

> ListSnapshotJobsResponseContent ListSnapshotJobs(ctx).ListSnapshotJobsRequestContent(listSnapshotJobsRequestContent).Execute()





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
	listSnapshotJobsRequestContent := *openapiclient.NewListSnapshotJobsRequestContent("ServiceId_example") // ListSnapshotJobsRequestContent | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SnapshotsAPI.ListSnapshotJobs(context.Background()).ListSnapshotJobsRequestContent(listSnapshotJobsRequestContent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsAPI.ListSnapshotJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSnapshotJobs`: ListSnapshotJobsResponseContent
	fmt.Fprintf(os.Stdout, "Response from `SnapshotsAPI.ListSnapshotJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSnapshotJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listSnapshotJobsRequestContent** | [**ListSnapshotJobsRequestContent**](ListSnapshotJobsRequestContent.md) |  | 

### Return type

[**ListSnapshotJobsResponseContent**](ListSnapshotJobsResponseContent.md)

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


## UpdateSnapshotJob

> UpdateSnapshotJobResponseContent UpdateSnapshotJob(ctx).UpdateSnapshotJobRequestContent(updateSnapshotJobRequestContent).Execute()





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
	updateSnapshotJobRequestContent := *openapiclient.NewUpdateSnapshotJobRequestContent("ServiceId_example", "JobId_example") // UpdateSnapshotJobRequestContent | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SnapshotsAPI.UpdateSnapshotJob(context.Background()).UpdateSnapshotJobRequestContent(updateSnapshotJobRequestContent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsAPI.UpdateSnapshotJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSnapshotJob`: UpdateSnapshotJobResponseContent
	fmt.Fprintf(os.Stdout, "Response from `SnapshotsAPI.UpdateSnapshotJob`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSnapshotJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateSnapshotJobRequestContent** | [**UpdateSnapshotJobRequestContent**](UpdateSnapshotJobRequestContent.md) |  | 

### Return type

[**UpdateSnapshotJobResponseContent**](UpdateSnapshotJobResponseContent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

