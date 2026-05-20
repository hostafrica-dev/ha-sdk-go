# \FirewallAPI

All URIs are relative to *https://api.hostafrica.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFirewallRule**](FirewallAPI.md#CreateFirewallRule) | **Post** /vps/create-firewall-rule | 
[**DeleteFirewallRule**](FirewallAPI.md#DeleteFirewallRule) | **Post** /vps/delete-firewall-rule | 
[**ListFirewallRules**](FirewallAPI.md#ListFirewallRules) | **Post** /vps/list-firewall-rules | 
[**MoveFirewallRule**](FirewallAPI.md#MoveFirewallRule) | **Post** /vps/move-firewall-rule | 
[**UpdateFirewallRule**](FirewallAPI.md#UpdateFirewallRule) | **Post** /vps/update-firewall-rule | 



## CreateFirewallRule

> CreateFirewallRuleResponseContent CreateFirewallRule(ctx).CreateFirewallRuleRequestContent(createFirewallRuleRequestContent).Execute()





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
	createFirewallRuleRequestContent := *openapiclient.NewCreateFirewallRuleRequestContent("ServiceId_example", openapiclient.FirewallRuleType("in"), openapiclient.FirewallRuleAction("ACCEPT")) // CreateFirewallRuleRequestContent | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirewallAPI.CreateFirewallRule(context.Background()).CreateFirewallRuleRequestContent(createFirewallRuleRequestContent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallAPI.CreateFirewallRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFirewallRule`: CreateFirewallRuleResponseContent
	fmt.Fprintf(os.Stdout, "Response from `FirewallAPI.CreateFirewallRule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFirewallRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createFirewallRuleRequestContent** | [**CreateFirewallRuleRequestContent**](CreateFirewallRuleRequestContent.md) |  | 

### Return type

[**CreateFirewallRuleResponseContent**](CreateFirewallRuleResponseContent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFirewallRule

> DeleteFirewallRuleResponseContent DeleteFirewallRule(ctx).DeleteFirewallRuleRequestContent(deleteFirewallRuleRequestContent).Execute()





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
	deleteFirewallRuleRequestContent := *openapiclient.NewDeleteFirewallRuleRequestContent("ServiceId_example", int32(123)) // DeleteFirewallRuleRequestContent | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirewallAPI.DeleteFirewallRule(context.Background()).DeleteFirewallRuleRequestContent(deleteFirewallRuleRequestContent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallAPI.DeleteFirewallRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteFirewallRule`: DeleteFirewallRuleResponseContent
	fmt.Fprintf(os.Stdout, "Response from `FirewallAPI.DeleteFirewallRule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFirewallRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteFirewallRuleRequestContent** | [**DeleteFirewallRuleRequestContent**](DeleteFirewallRuleRequestContent.md) |  | 

### Return type

[**DeleteFirewallRuleResponseContent**](DeleteFirewallRuleResponseContent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFirewallRules

> ListFirewallRulesResponseContent ListFirewallRules(ctx).ListFirewallRulesRequestContent(listFirewallRulesRequestContent).Execute()





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
	listFirewallRulesRequestContent := *openapiclient.NewListFirewallRulesRequestContent("ServiceId_example") // ListFirewallRulesRequestContent | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirewallAPI.ListFirewallRules(context.Background()).ListFirewallRulesRequestContent(listFirewallRulesRequestContent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallAPI.ListFirewallRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFirewallRules`: ListFirewallRulesResponseContent
	fmt.Fprintf(os.Stdout, "Response from `FirewallAPI.ListFirewallRules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFirewallRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listFirewallRulesRequestContent** | [**ListFirewallRulesRequestContent**](ListFirewallRulesRequestContent.md) |  | 

### Return type

[**ListFirewallRulesResponseContent**](ListFirewallRulesResponseContent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MoveFirewallRule

> MoveFirewallRuleResponseContent MoveFirewallRule(ctx).MoveFirewallRuleRequestContent(moveFirewallRuleRequestContent).Execute()





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
	moveFirewallRuleRequestContent := *openapiclient.NewMoveFirewallRuleRequestContent("ServiceId_example", int32(123)) // MoveFirewallRuleRequestContent | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirewallAPI.MoveFirewallRule(context.Background()).MoveFirewallRuleRequestContent(moveFirewallRuleRequestContent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallAPI.MoveFirewallRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MoveFirewallRule`: MoveFirewallRuleResponseContent
	fmt.Fprintf(os.Stdout, "Response from `FirewallAPI.MoveFirewallRule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMoveFirewallRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **moveFirewallRuleRequestContent** | [**MoveFirewallRuleRequestContent**](MoveFirewallRuleRequestContent.md) |  | 

### Return type

[**MoveFirewallRuleResponseContent**](MoveFirewallRuleResponseContent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFirewallRule

> UpdateFirewallRuleResponseContent UpdateFirewallRule(ctx).UpdateFirewallRuleRequestContent(updateFirewallRuleRequestContent).Execute()





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
	updateFirewallRuleRequestContent := *openapiclient.NewUpdateFirewallRuleRequestContent("ServiceId_example", int32(123)) // UpdateFirewallRuleRequestContent | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirewallAPI.UpdateFirewallRule(context.Background()).UpdateFirewallRuleRequestContent(updateFirewallRuleRequestContent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallAPI.UpdateFirewallRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateFirewallRule`: UpdateFirewallRuleResponseContent
	fmt.Fprintf(os.Stdout, "Response from `FirewallAPI.UpdateFirewallRule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFirewallRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateFirewallRuleRequestContent** | [**UpdateFirewallRuleRequestContent**](UpdateFirewallRuleRequestContent.md) |  | 

### Return type

[**UpdateFirewallRuleResponseContent**](UpdateFirewallRuleResponseContent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

