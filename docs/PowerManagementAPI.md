# \PowerManagementAPI

All URIs are relative to *https://api.hostafrica.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePowerTask**](PowerManagementAPI.md#CreatePowerTask) | **Post** /vps/create-power-task | 
[**DeletePowerTask**](PowerManagementAPI.md#DeletePowerTask) | **Post** /vps/delete-power-task | 
[**ListPowerTasks**](PowerManagementAPI.md#ListPowerTasks) | **Post** /vps/list-power-tasks | 
[**RebootVps**](PowerManagementAPI.md#RebootVps) | **Post** /vps/reboot | 
[**ShutdownVps**](PowerManagementAPI.md#ShutdownVps) | **Post** /vps/shutdown | 
[**StartVps**](PowerManagementAPI.md#StartVps) | **Post** /vps/start | 
[**StopVps**](PowerManagementAPI.md#StopVps) | **Post** /vps/stop | 
[**UpdatePowerTask**](PowerManagementAPI.md#UpdatePowerTask) | **Post** /vps/update-power-task | 



## CreatePowerTask

> CreatePowerTaskResponseContent CreatePowerTask(ctx).CreatePowerTaskRequestContent(createPowerTaskRequestContent).Execute()





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
	createPowerTaskRequestContent := *openapiclient.NewCreatePowerTaskRequestContent("ServiceId_example", openapiclient.PowerTaskAction("start"), "StartDate_example") // CreatePowerTaskRequestContent | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PowerManagementAPI.CreatePowerTask(context.Background()).CreatePowerTaskRequestContent(createPowerTaskRequestContent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PowerManagementAPI.CreatePowerTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePowerTask`: CreatePowerTaskResponseContent
	fmt.Fprintf(os.Stdout, "Response from `PowerManagementAPI.CreatePowerTask`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePowerTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createPowerTaskRequestContent** | [**CreatePowerTaskRequestContent**](CreatePowerTaskRequestContent.md) |  | 

### Return type

[**CreatePowerTaskResponseContent**](CreatePowerTaskResponseContent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePowerTask

> DeletePowerTaskResponseContent DeletePowerTask(ctx).DeletePowerTaskRequestContent(deletePowerTaskRequestContent).Execute()





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
	deletePowerTaskRequestContent := *openapiclient.NewDeletePowerTaskRequestContent("ServiceId_example", int32(123)) // DeletePowerTaskRequestContent | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PowerManagementAPI.DeletePowerTask(context.Background()).DeletePowerTaskRequestContent(deletePowerTaskRequestContent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PowerManagementAPI.DeletePowerTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeletePowerTask`: DeletePowerTaskResponseContent
	fmt.Fprintf(os.Stdout, "Response from `PowerManagementAPI.DeletePowerTask`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeletePowerTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deletePowerTaskRequestContent** | [**DeletePowerTaskRequestContent**](DeletePowerTaskRequestContent.md) |  | 

### Return type

[**DeletePowerTaskResponseContent**](DeletePowerTaskResponseContent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPowerTasks

> ListPowerTasksResponseContent ListPowerTasks(ctx).ListPowerTasksRequestContent(listPowerTasksRequestContent).Execute()





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
	listPowerTasksRequestContent := *openapiclient.NewListPowerTasksRequestContent("ServiceId_example") // ListPowerTasksRequestContent | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PowerManagementAPI.ListPowerTasks(context.Background()).ListPowerTasksRequestContent(listPowerTasksRequestContent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PowerManagementAPI.ListPowerTasks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPowerTasks`: ListPowerTasksResponseContent
	fmt.Fprintf(os.Stdout, "Response from `PowerManagementAPI.ListPowerTasks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPowerTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listPowerTasksRequestContent** | [**ListPowerTasksRequestContent**](ListPowerTasksRequestContent.md) |  | 

### Return type

[**ListPowerTasksResponseContent**](ListPowerTasksResponseContent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RebootVps

> RebootVpsResponseContent RebootVps(ctx).RebootVpsRequestContent(rebootVpsRequestContent).Execute()





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
	rebootVpsRequestContent := *openapiclient.NewRebootVpsRequestContent("ServiceId_example") // RebootVpsRequestContent | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PowerManagementAPI.RebootVps(context.Background()).RebootVpsRequestContent(rebootVpsRequestContent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PowerManagementAPI.RebootVps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RebootVps`: RebootVpsResponseContent
	fmt.Fprintf(os.Stdout, "Response from `PowerManagementAPI.RebootVps`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRebootVpsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rebootVpsRequestContent** | [**RebootVpsRequestContent**](RebootVpsRequestContent.md) |  | 

### Return type

[**RebootVpsResponseContent**](RebootVpsResponseContent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShutdownVps

> ShutdownVpsResponseContent ShutdownVps(ctx).ShutdownVpsRequestContent(shutdownVpsRequestContent).Execute()





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
	shutdownVpsRequestContent := *openapiclient.NewShutdownVpsRequestContent("ServiceId_example") // ShutdownVpsRequestContent | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PowerManagementAPI.ShutdownVps(context.Background()).ShutdownVpsRequestContent(shutdownVpsRequestContent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PowerManagementAPI.ShutdownVps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShutdownVps`: ShutdownVpsResponseContent
	fmt.Fprintf(os.Stdout, "Response from `PowerManagementAPI.ShutdownVps`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiShutdownVpsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **shutdownVpsRequestContent** | [**ShutdownVpsRequestContent**](ShutdownVpsRequestContent.md) |  | 

### Return type

[**ShutdownVpsResponseContent**](ShutdownVpsResponseContent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartVps

> StartVpsResponseContent StartVps(ctx).StartVpsRequestContent(startVpsRequestContent).Execute()





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
	startVpsRequestContent := *openapiclient.NewStartVpsRequestContent("ServiceId_example") // StartVpsRequestContent | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PowerManagementAPI.StartVps(context.Background()).StartVpsRequestContent(startVpsRequestContent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PowerManagementAPI.StartVps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartVps`: StartVpsResponseContent
	fmt.Fprintf(os.Stdout, "Response from `PowerManagementAPI.StartVps`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStartVpsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startVpsRequestContent** | [**StartVpsRequestContent**](StartVpsRequestContent.md) |  | 

### Return type

[**StartVpsResponseContent**](StartVpsResponseContent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopVps

> StopVpsResponseContent StopVps(ctx).StopVpsRequestContent(stopVpsRequestContent).Execute()





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
	stopVpsRequestContent := *openapiclient.NewStopVpsRequestContent("ServiceId_example") // StopVpsRequestContent | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PowerManagementAPI.StopVps(context.Background()).StopVpsRequestContent(stopVpsRequestContent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PowerManagementAPI.StopVps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StopVps`: StopVpsResponseContent
	fmt.Fprintf(os.Stdout, "Response from `PowerManagementAPI.StopVps`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStopVpsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stopVpsRequestContent** | [**StopVpsRequestContent**](StopVpsRequestContent.md) |  | 

### Return type

[**StopVpsResponseContent**](StopVpsResponseContent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePowerTask

> UpdatePowerTaskResponseContent UpdatePowerTask(ctx).UpdatePowerTaskRequestContent(updatePowerTaskRequestContent).Execute()





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
	updatePowerTaskRequestContent := *openapiclient.NewUpdatePowerTaskRequestContent("ServiceId_example", int32(123)) // UpdatePowerTaskRequestContent | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PowerManagementAPI.UpdatePowerTask(context.Background()).UpdatePowerTaskRequestContent(updatePowerTaskRequestContent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PowerManagementAPI.UpdatePowerTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePowerTask`: UpdatePowerTaskResponseContent
	fmt.Fprintf(os.Stdout, "Response from `PowerManagementAPI.UpdatePowerTask`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePowerTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updatePowerTaskRequestContent** | [**UpdatePowerTaskRequestContent**](UpdatePowerTaskRequestContent.md) |  | 

### Return type

[**UpdatePowerTaskResponseContent**](UpdatePowerTaskResponseContent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

