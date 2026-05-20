# \VPSManagementAPI

All URIs are relative to *https://api.hostafrica.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetVpsConfig**](VPSManagementAPI.md#GetVpsConfig) | **Post** /vps/get-config | 
[**GetVpsDetails**](VPSManagementAPI.md#GetVpsDetails) | **Post** /vps/get-details | 
[**ListIsos**](VPSManagementAPI.md#ListIsos) | **Post** /vps/list-isos | 
[**ListReinstallOs**](VPSManagementAPI.md#ListReinstallOs) | **Post** /vps/list-reinstall-images | 
[**ListVpsServices**](VPSManagementAPI.md#ListVpsServices) | **Post** /vps/list-vps-services | 
[**MountIso**](VPSManagementAPI.md#MountIso) | **Post** /vps/mount-iso | 
[**TriggerReinstall**](VPSManagementAPI.md#TriggerReinstall) | **Post** /vps/trigger-reinstall | 
[**UpdateVpsConfig**](VPSManagementAPI.md#UpdateVpsConfig) | **Post** /vps/update-config | 



## GetVpsConfig

> GetVpsConfigResponseContent GetVpsConfig(ctx).GetVpsConfigRequestContent(getVpsConfigRequestContent).Execute()





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
	getVpsConfigRequestContent := *openapiclient.NewGetVpsConfigRequestContent("ServiceId_example") // GetVpsConfigRequestContent | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPSManagementAPI.GetVpsConfig(context.Background()).GetVpsConfigRequestContent(getVpsConfigRequestContent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPSManagementAPI.GetVpsConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVpsConfig`: GetVpsConfigResponseContent
	fmt.Fprintf(os.Stdout, "Response from `VPSManagementAPI.GetVpsConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVpsConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getVpsConfigRequestContent** | [**GetVpsConfigRequestContent**](GetVpsConfigRequestContent.md) |  | 

### Return type

[**GetVpsConfigResponseContent**](GetVpsConfigResponseContent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVpsDetails

> GetVpsDetailsResponseContent GetVpsDetails(ctx).GetVpsDetailsRequestContent(getVpsDetailsRequestContent).Execute()





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
	getVpsDetailsRequestContent := *openapiclient.NewGetVpsDetailsRequestContent("ServiceId_example") // GetVpsDetailsRequestContent | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPSManagementAPI.GetVpsDetails(context.Background()).GetVpsDetailsRequestContent(getVpsDetailsRequestContent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPSManagementAPI.GetVpsDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVpsDetails`: GetVpsDetailsResponseContent
	fmt.Fprintf(os.Stdout, "Response from `VPSManagementAPI.GetVpsDetails`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVpsDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getVpsDetailsRequestContent** | [**GetVpsDetailsRequestContent**](GetVpsDetailsRequestContent.md) |  | 

### Return type

[**GetVpsDetailsResponseContent**](GetVpsDetailsResponseContent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIsos

> ListIsosResponseContent ListIsos(ctx).ListIsosRequestContent(listIsosRequestContent).Execute()





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
	listIsosRequestContent := *openapiclient.NewListIsosRequestContent("ServiceId_example") // ListIsosRequestContent | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPSManagementAPI.ListIsos(context.Background()).ListIsosRequestContent(listIsosRequestContent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPSManagementAPI.ListIsos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListIsos`: ListIsosResponseContent
	fmt.Fprintf(os.Stdout, "Response from `VPSManagementAPI.ListIsos`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListIsosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listIsosRequestContent** | [**ListIsosRequestContent**](ListIsosRequestContent.md) |  | 

### Return type

[**ListIsosResponseContent**](ListIsosResponseContent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListReinstallOs

> ListReinstallOsResponseContent ListReinstallOs(ctx).ListReinstallOsRequestContent(listReinstallOsRequestContent).Execute()





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
	listReinstallOsRequestContent := *openapiclient.NewListReinstallOsRequestContent("ServiceId_example") // ListReinstallOsRequestContent | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPSManagementAPI.ListReinstallOs(context.Background()).ListReinstallOsRequestContent(listReinstallOsRequestContent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPSManagementAPI.ListReinstallOs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListReinstallOs`: ListReinstallOsResponseContent
	fmt.Fprintf(os.Stdout, "Response from `VPSManagementAPI.ListReinstallOs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListReinstallOsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listReinstallOsRequestContent** | [**ListReinstallOsRequestContent**](ListReinstallOsRequestContent.md) |  | 

### Return type

[**ListReinstallOsResponseContent**](ListReinstallOsResponseContent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVpsServices

> ListVpsServicesResponseContent ListVpsServices(ctx).Execute()





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
	resp, r, err := apiClient.VPSManagementAPI.ListVpsServices(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPSManagementAPI.ListVpsServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVpsServices`: ListVpsServicesResponseContent
	fmt.Fprintf(os.Stdout, "Response from `VPSManagementAPI.ListVpsServices`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVpsServicesRequest struct via the builder pattern


### Return type

[**ListVpsServicesResponseContent**](ListVpsServicesResponseContent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MountIso

> MountIsoResponseContent MountIso(ctx).MountIsoRequestContent(mountIsoRequestContent).Execute()





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
	mountIsoRequestContent := *openapiclient.NewMountIsoRequestContent("ServiceId_example", "Iso_example") // MountIsoRequestContent | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPSManagementAPI.MountIso(context.Background()).MountIsoRequestContent(mountIsoRequestContent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPSManagementAPI.MountIso``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MountIso`: MountIsoResponseContent
	fmt.Fprintf(os.Stdout, "Response from `VPSManagementAPI.MountIso`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMountIsoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mountIsoRequestContent** | [**MountIsoRequestContent**](MountIsoRequestContent.md) |  | 

### Return type

[**MountIsoResponseContent**](MountIsoResponseContent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TriggerReinstall

> TriggerReinstallResponseContent TriggerReinstall(ctx).TriggerReinstallRequestContent(triggerReinstallRequestContent).Execute()





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
	triggerReinstallRequestContent := *openapiclient.NewTriggerReinstallRequestContent("ServiceId_example", "Template_example") // TriggerReinstallRequestContent | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPSManagementAPI.TriggerReinstall(context.Background()).TriggerReinstallRequestContent(triggerReinstallRequestContent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPSManagementAPI.TriggerReinstall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TriggerReinstall`: TriggerReinstallResponseContent
	fmt.Fprintf(os.Stdout, "Response from `VPSManagementAPI.TriggerReinstall`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTriggerReinstallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **triggerReinstallRequestContent** | [**TriggerReinstallRequestContent**](TriggerReinstallRequestContent.md) |  | 

### Return type

[**TriggerReinstallResponseContent**](TriggerReinstallResponseContent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVpsConfig

> UpdateVpsConfigResponseContent UpdateVpsConfig(ctx).UpdateVpsConfigRequestContent(updateVpsConfigRequestContent).Execute()





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
	updateVpsConfigRequestContent := *openapiclient.NewUpdateVpsConfigRequestContent("ServiceId_example") // UpdateVpsConfigRequestContent | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPSManagementAPI.UpdateVpsConfig(context.Background()).UpdateVpsConfigRequestContent(updateVpsConfigRequestContent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPSManagementAPI.UpdateVpsConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateVpsConfig`: UpdateVpsConfigResponseContent
	fmt.Fprintf(os.Stdout, "Response from `VPSManagementAPI.UpdateVpsConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVpsConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateVpsConfigRequestContent** | [**UpdateVpsConfigRequestContent**](UpdateVpsConfigRequestContent.md) |  | 

### Return type

[**UpdateVpsConfigResponseContent**](UpdateVpsConfigResponseContent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

