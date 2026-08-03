# \DNSAPI

All URIs are relative to *https://api.hostafrica.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddDnsRecord**](DNSAPI.md#AddDnsRecord) | **Post** /dns/add-record | 
[**CreateRdnsRecord**](DNSAPI.md#CreateRdnsRecord) | **Post** /dns/create-rdns-record | 
[**DeleteDnsRecord**](DNSAPI.md#DeleteDnsRecord) | **Post** /dns/delete-record | 
[**DeleteRdnsRecord**](DNSAPI.md#DeleteRdnsRecord) | **Post** /dns/delete-rdns-record | 
[**EditDnsRecord**](DNSAPI.md#EditDnsRecord) | **Post** /dns/edit-record | 
[**GetDnsZoneDetails**](DNSAPI.md#GetDnsZoneDetails) | **Post** /dns/get-zone | 
[**ListDnsCreateCandidates**](DNSAPI.md#ListDnsCreateCandidates) | **Post** /dns/list-create-candidates | 
[**ListDnsZones**](DNSAPI.md#ListDnsZones) | **Post** /dns/list-zones | 
[**ListRdnsRecords**](DNSAPI.md#ListRdnsRecords) | **Post** /dns/list-rdns-records | 



## AddDnsRecord

> AddDnsRecordResponseContent AddDnsRecord(ctx).AddDnsRecordRequestContent(addDnsRecordRequestContent).Execute()





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
	addDnsRecordRequestContent := *openapiclient.NewAddDnsRecordRequestContent("ZoneId_example", *openapiclient.NewDnsRecordMutationRecord()) // AddDnsRecordRequestContent | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DNSAPI.AddDnsRecord(context.Background()).AddDnsRecordRequestContent(addDnsRecordRequestContent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DNSAPI.AddDnsRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddDnsRecord`: AddDnsRecordResponseContent
	fmt.Fprintf(os.Stdout, "Response from `DNSAPI.AddDnsRecord`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddDnsRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addDnsRecordRequestContent** | [**AddDnsRecordRequestContent**](AddDnsRecordRequestContent.md) |  | 

### Return type

[**AddDnsRecordResponseContent**](AddDnsRecordResponseContent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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


## DeleteDnsRecord

> DeleteDnsRecordResponseContent DeleteDnsRecord(ctx).DeleteDnsRecordRequestContent(deleteDnsRecordRequestContent).Execute()





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
	deleteDnsRecordRequestContent := *openapiclient.NewDeleteDnsRecordRequestContent("ZoneId_example", *openapiclient.NewDnsRecordMutationRecord()) // DeleteDnsRecordRequestContent | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DNSAPI.DeleteDnsRecord(context.Background()).DeleteDnsRecordRequestContent(deleteDnsRecordRequestContent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DNSAPI.DeleteDnsRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDnsRecord`: DeleteDnsRecordResponseContent
	fmt.Fprintf(os.Stdout, "Response from `DNSAPI.DeleteDnsRecord`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDnsRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteDnsRecordRequestContent** | [**DeleteDnsRecordRequestContent**](DeleteDnsRecordRequestContent.md) |  | 

### Return type

[**DeleteDnsRecordResponseContent**](DeleteDnsRecordResponseContent.md)

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


## EditDnsRecord

> EditDnsRecordResponseContent EditDnsRecord(ctx).EditDnsRecordRequestContent(editDnsRecordRequestContent).Execute()





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
	editDnsRecordRequestContent := *openapiclient.NewEditDnsRecordRequestContent("ZoneId_example", *openapiclient.NewDnsRecordMutationRecord()) // EditDnsRecordRequestContent | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DNSAPI.EditDnsRecord(context.Background()).EditDnsRecordRequestContent(editDnsRecordRequestContent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DNSAPI.EditDnsRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditDnsRecord`: EditDnsRecordResponseContent
	fmt.Fprintf(os.Stdout, "Response from `DNSAPI.EditDnsRecord`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEditDnsRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **editDnsRecordRequestContent** | [**EditDnsRecordRequestContent**](EditDnsRecordRequestContent.md) |  | 

### Return type

[**EditDnsRecordResponseContent**](EditDnsRecordResponseContent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDnsZoneDetails

> GetDnsZoneDetailsResponseContent GetDnsZoneDetails(ctx).GetDnsZoneDetailsRequestContent(getDnsZoneDetailsRequestContent).Execute()





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
	getDnsZoneDetailsRequestContent := *openapiclient.NewGetDnsZoneDetailsRequestContent("DomainId_example") // GetDnsZoneDetailsRequestContent | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DNSAPI.GetDnsZoneDetails(context.Background()).GetDnsZoneDetailsRequestContent(getDnsZoneDetailsRequestContent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DNSAPI.GetDnsZoneDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDnsZoneDetails`: GetDnsZoneDetailsResponseContent
	fmt.Fprintf(os.Stdout, "Response from `DNSAPI.GetDnsZoneDetails`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDnsZoneDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getDnsZoneDetailsRequestContent** | [**GetDnsZoneDetailsRequestContent**](GetDnsZoneDetailsRequestContent.md) |  | 

### Return type

[**GetDnsZoneDetailsResponseContent**](GetDnsZoneDetailsResponseContent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDnsCreateCandidates

> ListDnsCreateCandidatesResponseContent ListDnsCreateCandidates(ctx).Execute()





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
	resp, r, err := apiClient.DNSAPI.ListDnsCreateCandidates(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DNSAPI.ListDnsCreateCandidates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDnsCreateCandidates`: ListDnsCreateCandidatesResponseContent
	fmt.Fprintf(os.Stdout, "Response from `DNSAPI.ListDnsCreateCandidates`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListDnsCreateCandidatesRequest struct via the builder pattern


### Return type

[**ListDnsCreateCandidatesResponseContent**](ListDnsCreateCandidatesResponseContent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDnsZones

> ListDnsZonesResponseContent ListDnsZones(ctx).Execute()





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
	resp, r, err := apiClient.DNSAPI.ListDnsZones(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DNSAPI.ListDnsZones``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDnsZones`: ListDnsZonesResponseContent
	fmt.Fprintf(os.Stdout, "Response from `DNSAPI.ListDnsZones`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListDnsZonesRequest struct via the builder pattern


### Return type

[**ListDnsZonesResponseContent**](ListDnsZonesResponseContent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
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

