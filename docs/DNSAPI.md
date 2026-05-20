# \DNSAPI

All URIs are relative to *https://api.hostafrica.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRdnsRecord**](DNSAPI.md#CreateRdnsRecord) | **Post** /dns/create-rdns-record | 
[**DeleteRdnsRecord**](DNSAPI.md#DeleteRdnsRecord) | **Post** /dns/delete-rdns-record | 
[**ListRdnsRecords**](DNSAPI.md#ListRdnsRecords) | **Post** /dns/list-rdns-records | 



## CreateRdnsRecord

> CreateRdnsRecordResponseContent CreateRdnsRecord(ctx).CreateRdnsRecordRequestContent(createRdnsRecordRequestContent).Execute()





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
	createRdnsRecordRequestContent := *openapiclient.NewCreateRdnsRecordRequestContent(int32(123), int32(123), "Ip_example", "Hostname_example") // CreateRdnsRecordRequestContent | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DNSAPI.CreateRdnsRecord(context.Background()).CreateRdnsRecordRequestContent(createRdnsRecordRequestContent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DNSAPI.CreateRdnsRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRdnsRecord`: CreateRdnsRecordResponseContent
	fmt.Fprintf(os.Stdout, "Response from `DNSAPI.CreateRdnsRecord`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRdnsRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createRdnsRecordRequestContent** | [**CreateRdnsRecordRequestContent**](CreateRdnsRecordRequestContent.md) |  | 

### Return type

[**CreateRdnsRecordResponseContent**](CreateRdnsRecordResponseContent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRdnsRecord

> DeleteRdnsRecordResponseContent DeleteRdnsRecord(ctx).DeleteRdnsRecordRequestContent(deleteRdnsRecordRequestContent).Execute()





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
	deleteRdnsRecordRequestContent := *openapiclient.NewDeleteRdnsRecordRequestContent(int32(123)) // DeleteRdnsRecordRequestContent | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DNSAPI.DeleteRdnsRecord(context.Background()).DeleteRdnsRecordRequestContent(deleteRdnsRecordRequestContent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DNSAPI.DeleteRdnsRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteRdnsRecord`: DeleteRdnsRecordResponseContent
	fmt.Fprintf(os.Stdout, "Response from `DNSAPI.DeleteRdnsRecord`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRdnsRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteRdnsRecordRequestContent** | [**DeleteRdnsRecordRequestContent**](DeleteRdnsRecordRequestContent.md) |  | 

### Return type

[**DeleteRdnsRecordResponseContent**](DeleteRdnsRecordResponseContent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRdnsRecords

> ListRdnsRecordsResponseContent ListRdnsRecords(ctx).Execute()





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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DNSAPI.ListRdnsRecords(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DNSAPI.ListRdnsRecords``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRdnsRecords`: ListRdnsRecordsResponseContent
	fmt.Fprintf(os.Stdout, "Response from `DNSAPI.ListRdnsRecords`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListRdnsRecordsRequest struct via the builder pattern


### Return type

[**ListRdnsRecordsResponseContent**](ListRdnsRecordsResponseContent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

