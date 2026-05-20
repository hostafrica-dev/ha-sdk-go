# \UserManagementAPI

All URIs are relative to *https://api.hostafrica.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserChangePassword**](UserManagementAPI.md#UserChangePassword) | **Post** /user/change-password | 



## UserChangePassword

> UserChangePasswordResponseContent UserChangePassword(ctx).UserChangePasswordRequestContent(userChangePasswordRequestContent).Execute()





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
	userChangePasswordRequestContent := *openapiclient.NewUserChangePasswordRequestContent("OldPassword_example", "Password_example", "ConfirmPassword_example") // UserChangePasswordRequestContent | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserManagementAPI.UserChangePassword(context.Background()).UserChangePasswordRequestContent(userChangePasswordRequestContent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserManagementAPI.UserChangePassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserChangePassword`: UserChangePasswordResponseContent
	fmt.Fprintf(os.Stdout, "Response from `UserManagementAPI.UserChangePassword`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserChangePasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userChangePasswordRequestContent** | [**UserChangePasswordRequestContent**](UserChangePasswordRequestContent.md) |  | 

### Return type

[**UserChangePasswordResponseContent**](UserChangePasswordResponseContent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

