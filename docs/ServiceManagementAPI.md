# \ServiceManagementAPI

All URIs are relative to *https://api.hostafrica.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelVps**](ServiceManagementAPI.md#CancelVps) | **Post** /vps/cancel | 
[**CreateOrder**](ServiceManagementAPI.md#CreateOrder) | **Post** /vps/create-order | 
[**GetCatalogue**](ServiceManagementAPI.md#GetCatalogue) | **Post** /vps/get-catalogue | 
[**ListOrders**](ServiceManagementAPI.md#ListOrders) | **Post** /vps/list-orders | 
[**RetryPayment**](ServiceManagementAPI.md#RetryPayment) | **Post** /vps/retry-payment | 
[**ValidatePricing**](ServiceManagementAPI.md#ValidatePricing) | **Post** /vps/validate-pricing | 



## CancelVps

> CancelVpsResponseContent CancelVps(ctx).CancelVpsRequestContent(cancelVpsRequestContent).Execute()





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
	cancelVpsRequestContent := *openapiclient.NewCancelVpsRequestContent("ServiceId_example") // CancelVpsRequestContent | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceManagementAPI.CancelVps(context.Background()).CancelVpsRequestContent(cancelVpsRequestContent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceManagementAPI.CancelVps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelVps`: CancelVpsResponseContent
	fmt.Fprintf(os.Stdout, "Response from `ServiceManagementAPI.CancelVps`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCancelVpsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cancelVpsRequestContent** | [**CancelVpsRequestContent**](CancelVpsRequestContent.md) |  | 

### Return type

[**CancelVpsResponseContent**](CancelVpsResponseContent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrder

> CreateOrderResponseContent CreateOrder(ctx).CreateOrderRequestContent(createOrderRequestContent).Execute()





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
	createOrderRequestContent := *openapiclient.NewCreateOrderRequestContent([]openapiclient.CreateOrderProduct{*openapiclient.NewCreateOrderProduct(int32(123), openapiclient.BillingCycle("monthly"), int32(123), "Hostname_example", interface{}(123))}) // CreateOrderRequestContent | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceManagementAPI.CreateOrder(context.Background()).CreateOrderRequestContent(createOrderRequestContent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceManagementAPI.CreateOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrder`: CreateOrderResponseContent
	fmt.Fprintf(os.Stdout, "Response from `ServiceManagementAPI.CreateOrder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createOrderRequestContent** | [**CreateOrderRequestContent**](CreateOrderRequestContent.md) |  | 

### Return type

[**CreateOrderResponseContent**](CreateOrderResponseContent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCatalogue

> GetCatalogueResponseContent GetCatalogue(ctx).GetCatalogueRequestContent(getCatalogueRequestContent).Execute()





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
	getCatalogueRequestContent := *openapiclient.NewGetCatalogueRequestContent() // GetCatalogueRequestContent |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceManagementAPI.GetCatalogue(context.Background()).GetCatalogueRequestContent(getCatalogueRequestContent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceManagementAPI.GetCatalogue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCatalogue`: GetCatalogueResponseContent
	fmt.Fprintf(os.Stdout, "Response from `ServiceManagementAPI.GetCatalogue`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCatalogueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getCatalogueRequestContent** | [**GetCatalogueRequestContent**](GetCatalogueRequestContent.md) |  | 

### Return type

[**GetCatalogueResponseContent**](GetCatalogueResponseContent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrders

> ListOrdersResponseContent ListOrders(ctx).Execute()





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
	resp, r, err := apiClient.ServiceManagementAPI.ListOrders(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceManagementAPI.ListOrders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrders`: ListOrdersResponseContent
	fmt.Fprintf(os.Stdout, "Response from `ServiceManagementAPI.ListOrders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListOrdersRequest struct via the builder pattern


### Return type

[**ListOrdersResponseContent**](ListOrdersResponseContent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetryPayment

> RetryPaymentResponseContent RetryPayment(ctx).RetryPaymentRequestContent(retryPaymentRequestContent).Execute()





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
	retryPaymentRequestContent := *openapiclient.NewRetryPaymentRequestContent("ServiceId_example", int32(123)) // RetryPaymentRequestContent | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceManagementAPI.RetryPayment(context.Background()).RetryPaymentRequestContent(retryPaymentRequestContent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceManagementAPI.RetryPayment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetryPayment`: RetryPaymentResponseContent
	fmt.Fprintf(os.Stdout, "Response from `ServiceManagementAPI.RetryPayment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetryPaymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **retryPaymentRequestContent** | [**RetryPaymentRequestContent**](RetryPaymentRequestContent.md) |  | 

### Return type

[**RetryPaymentResponseContent**](RetryPaymentResponseContent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidatePricing

> ValidatePricingResponseContent ValidatePricing(ctx).ValidatePricingRequestContent(validatePricingRequestContent).Execute()





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
	validatePricingRequestContent := *openapiclient.NewValidatePricingRequestContent([]openapiclient.ValidatePricingProduct{*openapiclient.NewValidatePricingProduct(int32(123), openapiclient.BillingCycle("monthly"), int32(123), interface{}(123))}) // ValidatePricingRequestContent | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceManagementAPI.ValidatePricing(context.Background()).ValidatePricingRequestContent(validatePricingRequestContent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceManagementAPI.ValidatePricing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidatePricing`: ValidatePricingResponseContent
	fmt.Fprintf(os.Stdout, "Response from `ServiceManagementAPI.ValidatePricing`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidatePricingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **validatePricingRequestContent** | [**ValidatePricingRequestContent**](ValidatePricingRequestContent.md) |  | 

### Return type

[**ValidatePricingResponseContent**](ValidatePricingResponseContent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

