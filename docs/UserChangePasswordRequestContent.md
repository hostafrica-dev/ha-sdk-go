# UserChangePasswordRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OldPassword** | **string** | Current password | 
**Password** | **string** | New password | 
**ConfirmPassword** | **string** | Confirm new password (must match password) | 

## Methods

### NewUserChangePasswordRequestContent

`func NewUserChangePasswordRequestContent(oldPassword string, password string, confirmPassword string, ) *UserChangePasswordRequestContent`

NewUserChangePasswordRequestContent instantiates a new UserChangePasswordRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserChangePasswordRequestContentWithDefaults

`func NewUserChangePasswordRequestContentWithDefaults() *UserChangePasswordRequestContent`

NewUserChangePasswordRequestContentWithDefaults instantiates a new UserChangePasswordRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOldPassword

`func (o *UserChangePasswordRequestContent) GetOldPassword() string`

GetOldPassword returns the OldPassword field if non-nil, zero value otherwise.

### GetOldPasswordOk

`func (o *UserChangePasswordRequestContent) GetOldPasswordOk() (*string, bool)`

GetOldPasswordOk returns a tuple with the OldPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldPassword

`func (o *UserChangePasswordRequestContent) SetOldPassword(v string)`

SetOldPassword sets OldPassword field to given value.


### GetPassword

`func (o *UserChangePasswordRequestContent) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UserChangePasswordRequestContent) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UserChangePasswordRequestContent) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetConfirmPassword

`func (o *UserChangePasswordRequestContent) GetConfirmPassword() string`

GetConfirmPassword returns the ConfirmPassword field if non-nil, zero value otherwise.

### GetConfirmPasswordOk

`func (o *UserChangePasswordRequestContent) GetConfirmPasswordOk() (*string, bool)`

GetConfirmPasswordOk returns a tuple with the ConfirmPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmPassword

`func (o *UserChangePasswordRequestContent) SetConfirmPassword(v string)`

SetConfirmPassword sets ConfirmPassword field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


