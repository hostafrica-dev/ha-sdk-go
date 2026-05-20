# UserChangePasswordDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PasswordChanged** | **bool** | Whether the password was changed successfully | 
**SessionsRevoked** | **int32** | Number of active sessions that were revoked | 
**RequiresRelogin** | **bool** | Whether the user needs to login again | 

## Methods

### NewUserChangePasswordDetails

`func NewUserChangePasswordDetails(passwordChanged bool, sessionsRevoked int32, requiresRelogin bool, ) *UserChangePasswordDetails`

NewUserChangePasswordDetails instantiates a new UserChangePasswordDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserChangePasswordDetailsWithDefaults

`func NewUserChangePasswordDetailsWithDefaults() *UserChangePasswordDetails`

NewUserChangePasswordDetailsWithDefaults instantiates a new UserChangePasswordDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPasswordChanged

`func (o *UserChangePasswordDetails) GetPasswordChanged() bool`

GetPasswordChanged returns the PasswordChanged field if non-nil, zero value otherwise.

### GetPasswordChangedOk

`func (o *UserChangePasswordDetails) GetPasswordChangedOk() (*bool, bool)`

GetPasswordChangedOk returns a tuple with the PasswordChanged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordChanged

`func (o *UserChangePasswordDetails) SetPasswordChanged(v bool)`

SetPasswordChanged sets PasswordChanged field to given value.


### GetSessionsRevoked

`func (o *UserChangePasswordDetails) GetSessionsRevoked() int32`

GetSessionsRevoked returns the SessionsRevoked field if non-nil, zero value otherwise.

### GetSessionsRevokedOk

`func (o *UserChangePasswordDetails) GetSessionsRevokedOk() (*int32, bool)`

GetSessionsRevokedOk returns a tuple with the SessionsRevoked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionsRevoked

`func (o *UserChangePasswordDetails) SetSessionsRevoked(v int32)`

SetSessionsRevoked sets SessionsRevoked field to given value.


### GetRequiresRelogin

`func (o *UserChangePasswordDetails) GetRequiresRelogin() bool`

GetRequiresRelogin returns the RequiresRelogin field if non-nil, zero value otherwise.

### GetRequiresReloginOk

`func (o *UserChangePasswordDetails) GetRequiresReloginOk() (*bool, bool)`

GetRequiresReloginOk returns a tuple with the RequiresRelogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresRelogin

`func (o *UserChangePasswordDetails) SetRequiresRelogin(v bool)`

SetRequiresRelogin sets RequiresRelogin field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


