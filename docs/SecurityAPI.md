# \SecurityAPI

All URIs are relative to *https://api.hostafrica.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChangePassword**](SecurityAPI.md#ChangePassword) | **Post** /vps/change-password | 
[**GetPrivateSshKey**](SecurityAPI.md#GetPrivateSshKey) | **Post** /vps/get-private-ssh-keys | 
[**GetPublicSshKey**](SecurityAPI.md#GetPublicSshKey) | **Post** /vps/get-public-ssh-keys | 
[**UpdateSshKeys**](SecurityAPI.md#UpdateSshKeys) | **Post** /vps/update-ssh-keys | 



## ChangePassword

> ChangePasswordResponseContent ChangePassword(ctx).ChangePasswordRequestContent(changePasswordRequestContent).Execute()





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
	changePasswordRequestContent := *openapiclient.NewChangePasswordRequestContent("ServiceId_example", "Password_example") // ChangePasswordRequestContent | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.ChangePassword(context.Background()).ChangePasswordRequestContent(changePasswordRequestContent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.ChangePassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChangePassword`: ChangePasswordResponseContent
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.ChangePassword`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChangePasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **changePasswordRequestContent** | [**ChangePasswordRequestContent**](ChangePasswordRequestContent.md) |  | 

### Return type

[**ChangePasswordResponseContent**](ChangePasswordResponseContent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPrivateSshKey

> GetPrivateSshKeyResponseContent GetPrivateSshKey(ctx).GetPrivateSshKeyRequestContent(getPrivateSshKeyRequestContent).Execute()





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
	getPrivateSshKeyRequestContent := *openapiclient.NewGetPrivateSshKeyRequestContent("ServiceId_example") // GetPrivateSshKeyRequestContent | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.GetPrivateSshKey(context.Background()).GetPrivateSshKeyRequestContent(getPrivateSshKeyRequestContent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.GetPrivateSshKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPrivateSshKey`: GetPrivateSshKeyResponseContent
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.GetPrivateSshKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPrivateSshKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getPrivateSshKeyRequestContent** | [**GetPrivateSshKeyRequestContent**](GetPrivateSshKeyRequestContent.md) |  | 

### Return type

[**GetPrivateSshKeyResponseContent**](GetPrivateSshKeyResponseContent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPublicSshKey

> GetPublicSshKeyResponseContent GetPublicSshKey(ctx).GetPublicSshKeyRequestContent(getPublicSshKeyRequestContent).Execute()





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
	getPublicSshKeyRequestContent := *openapiclient.NewGetPublicSshKeyRequestContent("ServiceId_example") // GetPublicSshKeyRequestContent | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.GetPublicSshKey(context.Background()).GetPublicSshKeyRequestContent(getPublicSshKeyRequestContent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.GetPublicSshKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPublicSshKey`: GetPublicSshKeyResponseContent
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.GetPublicSshKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPublicSshKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getPublicSshKeyRequestContent** | [**GetPublicSshKeyRequestContent**](GetPublicSshKeyRequestContent.md) |  | 

### Return type

[**GetPublicSshKeyResponseContent**](GetPublicSshKeyResponseContent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSshKeys

> UpdateSshKeysResponseContent UpdateSshKeys(ctx).UpdateSshKeysRequestContent(updateSshKeysRequestContent).Execute()





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
	updateSshKeysRequestContent := *openapiclient.NewUpdateSshKeysRequestContent("ServiceId_example", "SshKeys_example") // UpdateSshKeysRequestContent | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.UpdateSshKeys(context.Background()).UpdateSshKeysRequestContent(updateSshKeysRequestContent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.UpdateSshKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSshKeys`: UpdateSshKeysResponseContent
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.UpdateSshKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSshKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateSshKeysRequestContent** | [**UpdateSshKeysRequestContent**](UpdateSshKeysRequestContent.md) |  | 

### Return type

[**UpdateSshKeysResponseContent**](UpdateSshKeysResponseContent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

