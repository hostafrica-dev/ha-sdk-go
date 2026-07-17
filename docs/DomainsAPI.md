# \DomainsAPI

All URIs are relative to *https://api.hostafrica.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckDomainAvailability**](DomainsAPI.md#CheckDomainAvailability) | **Post** /domain/check-availability | 
[**GetDomain**](DomainsAPI.md#GetDomain) | **Post** /domain/get-domain | 
[**GetDomainContacts**](DomainsAPI.md#GetDomainContacts) | **Post** /domain/get-domain-contacts | 
[**ListDomains**](DomainsAPI.md#ListDomains) | **Post** /domain/list-domains | 
[**ListDomainsRequiringData**](DomainsAPI.md#ListDomainsRequiringData) | **Post** /domain/list-domains-requiring-data | 
[**SaveDomainRequiredData**](DomainsAPI.md#SaveDomainRequiredData) | **Post** /domain/save-domain-required-data | 
[**SuggestDomains**](DomainsAPI.md#SuggestDomains) | **Post** /domain/suggest | 
[**UpdateDomainSettings**](DomainsAPI.md#UpdateDomainSettings) | **Post** /domain/update-domain-settings | 



## CheckDomainAvailability

> CheckDomainAvailabilityResponseContent CheckDomainAvailability(ctx).CheckDomainAvailabilityRequestContent(checkDomainAvailabilityRequestContent).Execute()





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
	checkDomainAvailabilityRequestContent := *openapiclient.NewCheckDomainAvailabilityRequestContent() // CheckDomainAvailabilityRequestContent |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainsAPI.CheckDomainAvailability(context.Background()).CheckDomainAvailabilityRequestContent(checkDomainAvailabilityRequestContent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.CheckDomainAvailability``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckDomainAvailability`: CheckDomainAvailabilityResponseContent
	fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.CheckDomainAvailability`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckDomainAvailabilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **checkDomainAvailabilityRequestContent** | [**CheckDomainAvailabilityRequestContent**](CheckDomainAvailabilityRequestContent.md) |  | 

### Return type

[**CheckDomainAvailabilityResponseContent**](CheckDomainAvailabilityResponseContent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDomain

> GetDomainResponseContent GetDomain(ctx).GetDomainRequestContent(getDomainRequestContent).Execute()





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
	getDomainRequestContent := *openapiclient.NewGetDomainRequestContent("DomainId_example") // GetDomainRequestContent | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainsAPI.GetDomain(context.Background()).GetDomainRequestContent(getDomainRequestContent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.GetDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDomain`: GetDomainResponseContent
	fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.GetDomain`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getDomainRequestContent** | [**GetDomainRequestContent**](GetDomainRequestContent.md) |  | 

### Return type

[**GetDomainResponseContent**](GetDomainResponseContent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDomainContacts

> GetDomainContactsResponseContent GetDomainContacts(ctx).GetDomainContactsRequestContent(getDomainContactsRequestContent).Execute()





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
	getDomainContactsRequestContent := *openapiclient.NewGetDomainContactsRequestContent("DomainId_example") // GetDomainContactsRequestContent | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainsAPI.GetDomainContacts(context.Background()).GetDomainContactsRequestContent(getDomainContactsRequestContent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.GetDomainContacts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDomainContacts`: GetDomainContactsResponseContent
	fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.GetDomainContacts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDomainContactsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getDomainContactsRequestContent** | [**GetDomainContactsRequestContent**](GetDomainContactsRequestContent.md) |  | 

### Return type

[**GetDomainContactsResponseContent**](GetDomainContactsResponseContent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDomains

> ListDomainsResponseContent ListDomains(ctx).Execute()





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
	resp, r, err := apiClient.DomainsAPI.ListDomains(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.ListDomains``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDomains`: ListDomainsResponseContent
	fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.ListDomains`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListDomainsRequest struct via the builder pattern


### Return type

[**ListDomainsResponseContent**](ListDomainsResponseContent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDomainsRequiringData

> ListDomainsRequiringDataResponseContent ListDomainsRequiringData(ctx).Execute()





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
	resp, r, err := apiClient.DomainsAPI.ListDomainsRequiringData(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.ListDomainsRequiringData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDomainsRequiringData`: ListDomainsRequiringDataResponseContent
	fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.ListDomainsRequiringData`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListDomainsRequiringDataRequest struct via the builder pattern


### Return type

[**ListDomainsRequiringDataResponseContent**](ListDomainsRequiringDataResponseContent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SaveDomainRequiredData

> SaveDomainRequiredDataResponseContent SaveDomainRequiredData(ctx).SaveDomainRequiredDataRequestContent(saveDomainRequiredDataRequestContent).Execute()





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
	saveDomainRequiredDataRequestContent := *openapiclient.NewSaveDomainRequiredDataRequestContent("DomainId_example", map[string]string{"key": "Inner_example"}) // SaveDomainRequiredDataRequestContent | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainsAPI.SaveDomainRequiredData(context.Background()).SaveDomainRequiredDataRequestContent(saveDomainRequiredDataRequestContent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.SaveDomainRequiredData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SaveDomainRequiredData`: SaveDomainRequiredDataResponseContent
	fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.SaveDomainRequiredData`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSaveDomainRequiredDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **saveDomainRequiredDataRequestContent** | [**SaveDomainRequiredDataRequestContent**](SaveDomainRequiredDataRequestContent.md) |  | 

### Return type

[**SaveDomainRequiredDataResponseContent**](SaveDomainRequiredDataResponseContent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SuggestDomains

> SuggestDomainsResponseContent SuggestDomains(ctx).SuggestDomainsRequestContent(suggestDomainsRequestContent).Execute()





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
	suggestDomainsRequestContent := *openapiclient.NewSuggestDomainsRequestContent("Prompt_example") // SuggestDomainsRequestContent | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainsAPI.SuggestDomains(context.Background()).SuggestDomainsRequestContent(suggestDomainsRequestContent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.SuggestDomains``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SuggestDomains`: SuggestDomainsResponseContent
	fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.SuggestDomains`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSuggestDomainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **suggestDomainsRequestContent** | [**SuggestDomainsRequestContent**](SuggestDomainsRequestContent.md) |  | 

### Return type

[**SuggestDomainsResponseContent**](SuggestDomainsResponseContent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDomainSettings

> UpdateDomainSettingsResponseContent UpdateDomainSettings(ctx).UpdateDomainSettingsRequestContent(updateDomainSettingsRequestContent).Execute()





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
	updateDomainSettingsRequestContent := *openapiclient.NewUpdateDomainSettingsRequestContent("DomainId_example", openapiclient.DomainSettingKey("donotrenew"), false) // UpdateDomainSettingsRequestContent | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainsAPI.UpdateDomainSettings(context.Background()).UpdateDomainSettingsRequestContent(updateDomainSettingsRequestContent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.UpdateDomainSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDomainSettings`: UpdateDomainSettingsResponseContent
	fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.UpdateDomainSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDomainSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateDomainSettingsRequestContent** | [**UpdateDomainSettingsRequestContent**](UpdateDomainSettingsRequestContent.md) |  | 

### Return type

[**UpdateDomainSettingsResponseContent**](UpdateDomainSettingsResponseContent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

