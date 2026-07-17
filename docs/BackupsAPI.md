# \BackupsAPI

All URIs are relative to *https://api.hostafrica.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBackup**](BackupsAPI.md#CreateBackup) | **Post** /vps/create-backup | 
[**CreateBackupSchedule**](BackupsAPI.md#CreateBackupSchedule) | **Post** /vps/create-backup-schedule | 
[**DeleteBackup**](BackupsAPI.md#DeleteBackup) | **Post** /vps/delete-backup | 
[**DeleteBackupSchedule**](BackupsAPI.md#DeleteBackupSchedule) | **Post** /vps/delete-backup-schedule | 
[**EditBackupSchedule**](BackupsAPI.md#EditBackupSchedule) | **Post** /vps/edit-backup-schedule | 
[**ListBackupSchedules**](BackupsAPI.md#ListBackupSchedules) | **Post** /vps/list-backup-schedules | 
[**ListBackups**](BackupsAPI.md#ListBackups) | **Post** /vps/list-backups | 
[**RestoreBackup**](BackupsAPI.md#RestoreBackup) | **Post** /vps/restore-backup | 



## CreateBackup

> CreateBackupResponseContent CreateBackup(ctx).CreateBackupRequestContent(createBackupRequestContent).Execute()





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
	createBackupRequestContent := *openapiclient.NewCreateBackupRequestContent("ServiceId_example") // CreateBackupRequestContent | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupsAPI.CreateBackup(context.Background()).CreateBackupRequestContent(createBackupRequestContent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupsAPI.CreateBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBackup`: CreateBackupResponseContent
	fmt.Fprintf(os.Stdout, "Response from `BackupsAPI.CreateBackup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createBackupRequestContent** | [**CreateBackupRequestContent**](CreateBackupRequestContent.md) |  | 

### Return type

[**CreateBackupResponseContent**](CreateBackupResponseContent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBackupSchedule

> CreateBackupScheduleResponseContent CreateBackupSchedule(ctx).CreateBackupScheduleRequestContent(createBackupScheduleRequestContent).Execute()





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
	createBackupScheduleRequestContent := *openapiclient.NewCreateBackupScheduleRequestContent("ServiceId_example", "Starttime_example", []openapiclient.DayOfWeek{openapiclient.DayOfWeek("mon")}, openapiclient.CompressionType("0"), openapiclient.BackupModeType("snapshot")) // CreateBackupScheduleRequestContent | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupsAPI.CreateBackupSchedule(context.Background()).CreateBackupScheduleRequestContent(createBackupScheduleRequestContent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupsAPI.CreateBackupSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBackupSchedule`: CreateBackupScheduleResponseContent
	fmt.Fprintf(os.Stdout, "Response from `BackupsAPI.CreateBackupSchedule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBackupScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createBackupScheduleRequestContent** | [**CreateBackupScheduleRequestContent**](CreateBackupScheduleRequestContent.md) |  | 

### Return type

[**CreateBackupScheduleResponseContent**](CreateBackupScheduleResponseContent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBackup

> DeleteBackupResponseContent DeleteBackup(ctx).DeleteBackupRequestContent(deleteBackupRequestContent).Execute()





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
	deleteBackupRequestContent := *openapiclient.NewDeleteBackupRequestContent("ServiceId_example", "BackupId_example") // DeleteBackupRequestContent | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupsAPI.DeleteBackup(context.Background()).DeleteBackupRequestContent(deleteBackupRequestContent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupsAPI.DeleteBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteBackup`: DeleteBackupResponseContent
	fmt.Fprintf(os.Stdout, "Response from `BackupsAPI.DeleteBackup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteBackupRequestContent** | [**DeleteBackupRequestContent**](DeleteBackupRequestContent.md) |  | 

### Return type

[**DeleteBackupResponseContent**](DeleteBackupResponseContent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBackupSchedule

> DeleteBackupScheduleResponseContent DeleteBackupSchedule(ctx).DeleteBackupScheduleRequestContent(deleteBackupScheduleRequestContent).Execute()





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
	deleteBackupScheduleRequestContent := *openapiclient.NewDeleteBackupScheduleRequestContent("ServiceId_example", "ScheduleId_example") // DeleteBackupScheduleRequestContent | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupsAPI.DeleteBackupSchedule(context.Background()).DeleteBackupScheduleRequestContent(deleteBackupScheduleRequestContent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupsAPI.DeleteBackupSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteBackupSchedule`: DeleteBackupScheduleResponseContent
	fmt.Fprintf(os.Stdout, "Response from `BackupsAPI.DeleteBackupSchedule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBackupScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteBackupScheduleRequestContent** | [**DeleteBackupScheduleRequestContent**](DeleteBackupScheduleRequestContent.md) |  | 

### Return type

[**DeleteBackupScheduleResponseContent**](DeleteBackupScheduleResponseContent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditBackupSchedule

> EditBackupScheduleResponseContent EditBackupSchedule(ctx).EditBackupScheduleRequestContent(editBackupScheduleRequestContent).Execute()





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
	editBackupScheduleRequestContent := *openapiclient.NewEditBackupScheduleRequestContent("ServiceId_example", "ScheduleId_example", "Starttime_example", []openapiclient.DayOfWeek{openapiclient.DayOfWeek("mon")}, openapiclient.CompressionType("0"), openapiclient.BackupModeType("snapshot")) // EditBackupScheduleRequestContent | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupsAPI.EditBackupSchedule(context.Background()).EditBackupScheduleRequestContent(editBackupScheduleRequestContent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupsAPI.EditBackupSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditBackupSchedule`: EditBackupScheduleResponseContent
	fmt.Fprintf(os.Stdout, "Response from `BackupsAPI.EditBackupSchedule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEditBackupScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **editBackupScheduleRequestContent** | [**EditBackupScheduleRequestContent**](EditBackupScheduleRequestContent.md) |  | 

### Return type

[**EditBackupScheduleResponseContent**](EditBackupScheduleResponseContent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBackupSchedules

> ListBackupSchedulesResponseContent ListBackupSchedules(ctx).ListBackupSchedulesRequestContent(listBackupSchedulesRequestContent).Execute()





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
	listBackupSchedulesRequestContent := *openapiclient.NewListBackupSchedulesRequestContent("ServiceId_example") // ListBackupSchedulesRequestContent | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupsAPI.ListBackupSchedules(context.Background()).ListBackupSchedulesRequestContent(listBackupSchedulesRequestContent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupsAPI.ListBackupSchedules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListBackupSchedules`: ListBackupSchedulesResponseContent
	fmt.Fprintf(os.Stdout, "Response from `BackupsAPI.ListBackupSchedules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBackupSchedulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listBackupSchedulesRequestContent** | [**ListBackupSchedulesRequestContent**](ListBackupSchedulesRequestContent.md) |  | 

### Return type

[**ListBackupSchedulesResponseContent**](ListBackupSchedulesResponseContent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBackups

> ListBackupsResponseContent ListBackups(ctx).ListBackupsRequestContent(listBackupsRequestContent).Execute()





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
	listBackupsRequestContent := *openapiclient.NewListBackupsRequestContent("ServiceId_example") // ListBackupsRequestContent | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupsAPI.ListBackups(context.Background()).ListBackupsRequestContent(listBackupsRequestContent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupsAPI.ListBackups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListBackups`: ListBackupsResponseContent
	fmt.Fprintf(os.Stdout, "Response from `BackupsAPI.ListBackups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBackupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listBackupsRequestContent** | [**ListBackupsRequestContent**](ListBackupsRequestContent.md) |  | 

### Return type

[**ListBackupsResponseContent**](ListBackupsResponseContent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestoreBackup

> RestoreBackupResponseContent RestoreBackup(ctx).RestoreBackupRequestContent(restoreBackupRequestContent).Execute()





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
	restoreBackupRequestContent := *openapiclient.NewRestoreBackupRequestContent("ServiceId_example", "BackupId_example") // RestoreBackupRequestContent | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupsAPI.RestoreBackup(context.Background()).RestoreBackupRequestContent(restoreBackupRequestContent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupsAPI.RestoreBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RestoreBackup`: RestoreBackupResponseContent
	fmt.Fprintf(os.Stdout, "Response from `BackupsAPI.RestoreBackup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRestoreBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **restoreBackupRequestContent** | [**RestoreBackupRequestContent**](RestoreBackupRequestContent.md) |  | 

### Return type

[**RestoreBackupResponseContent**](RestoreBackupResponseContent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

