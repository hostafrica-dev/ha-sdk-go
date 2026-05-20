# \ConsoleAccessAPI

All URIs are relative to *https://api.hostafrica.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNoVncConsole**](ConsoleAccessAPI.md#GetNoVncConsole) | **Post** /vps/novnc-console | 



## GetNoVncConsole

> GetNoVncConsoleResponseContent GetNoVncConsole(ctx).GetNoVncConsoleRequestContent(getNoVncConsoleRequestContent).Execute()





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
	getNoVncConsoleRequestContent := *openapiclient.NewGetNoVncConsoleRequestContent("ServiceId_example") // GetNoVncConsoleRequestContent | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsoleAccessAPI.GetNoVncConsole(context.Background()).GetNoVncConsoleRequestContent(getNoVncConsoleRequestContent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsoleAccessAPI.GetNoVncConsole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNoVncConsole`: GetNoVncConsoleResponseContent
	fmt.Fprintf(os.Stdout, "Response from `ConsoleAccessAPI.GetNoVncConsole`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNoVncConsoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getNoVncConsoleRequestContent** | [**GetNoVncConsoleRequestContent**](GetNoVncConsoleRequestContent.md) |  | 

### Return type

[**GetNoVncConsoleResponseContent**](GetNoVncConsoleResponseContent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

