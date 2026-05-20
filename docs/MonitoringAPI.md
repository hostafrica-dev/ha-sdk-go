# \MonitoringAPI

All URIs are relative to *https://api.hostafrica.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNotification**](MonitoringAPI.md#CreateNotification) | **Post** /vps/create-notification | 
[**DeleteNotification**](MonitoringAPI.md#DeleteNotification) | **Post** /vps/delete-notification | 
[**ListNotifications**](MonitoringAPI.md#ListNotifications) | **Post** /vps/list-notifications | 
[**UpdateNotification**](MonitoringAPI.md#UpdateNotification) | **Post** /vps/update-notification | 



## CreateNotification

> CreateNotificationResponseContent CreateNotification(ctx).CreateNotificationRequestContent(createNotificationRequestContent).Execute()





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
	createNotificationRequestContent := *openapiclient.NewCreateNotificationRequestContent("ServiceId_example", "Name_example") // CreateNotificationRequestContent | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MonitoringAPI.CreateNotification(context.Background()).CreateNotificationRequestContent(createNotificationRequestContent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MonitoringAPI.CreateNotification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNotification`: CreateNotificationResponseContent
	fmt.Fprintf(os.Stdout, "Response from `MonitoringAPI.CreateNotification`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNotificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createNotificationRequestContent** | [**CreateNotificationRequestContent**](CreateNotificationRequestContent.md) |  | 

### Return type

[**CreateNotificationResponseContent**](CreateNotificationResponseContent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNotification

> DeleteNotificationResponseContent DeleteNotification(ctx).DeleteNotificationRequestContent(deleteNotificationRequestContent).Execute()





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
	deleteNotificationRequestContent := *openapiclient.NewDeleteNotificationRequestContent("ServiceId_example", int32(123)) // DeleteNotificationRequestContent | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MonitoringAPI.DeleteNotification(context.Background()).DeleteNotificationRequestContent(deleteNotificationRequestContent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MonitoringAPI.DeleteNotification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteNotification`: DeleteNotificationResponseContent
	fmt.Fprintf(os.Stdout, "Response from `MonitoringAPI.DeleteNotification`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNotificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteNotificationRequestContent** | [**DeleteNotificationRequestContent**](DeleteNotificationRequestContent.md) |  | 

### Return type

[**DeleteNotificationResponseContent**](DeleteNotificationResponseContent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNotifications

> ListNotificationsResponseContent ListNotifications(ctx).ListNotificationsRequestContent(listNotificationsRequestContent).Execute()





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
	listNotificationsRequestContent := *openapiclient.NewListNotificationsRequestContent("ServiceId_example") // ListNotificationsRequestContent | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MonitoringAPI.ListNotifications(context.Background()).ListNotificationsRequestContent(listNotificationsRequestContent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MonitoringAPI.ListNotifications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNotifications`: ListNotificationsResponseContent
	fmt.Fprintf(os.Stdout, "Response from `MonitoringAPI.ListNotifications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNotificationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listNotificationsRequestContent** | [**ListNotificationsRequestContent**](ListNotificationsRequestContent.md) |  | 

### Return type

[**ListNotificationsResponseContent**](ListNotificationsResponseContent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNotification

> UpdateNotificationResponseContent UpdateNotification(ctx).UpdateNotificationRequestContent(updateNotificationRequestContent).Execute()





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
	updateNotificationRequestContent := *openapiclient.NewUpdateNotificationRequestContent("ServiceId_example", int32(123)) // UpdateNotificationRequestContent | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MonitoringAPI.UpdateNotification(context.Background()).UpdateNotificationRequestContent(updateNotificationRequestContent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MonitoringAPI.UpdateNotification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNotification`: UpdateNotificationResponseContent
	fmt.Fprintf(os.Stdout, "Response from `MonitoringAPI.UpdateNotification`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNotificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateNotificationRequestContent** | [**UpdateNotificationRequestContent**](UpdateNotificationRequestContent.md) |  | 

### Return type

[**UpdateNotificationResponseContent**](UpdateNotificationResponseContent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

