# UpdateDomainLockStatusRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainId** | **string** | Domain service id - must be sent as a string | 
**LockStatus** | **bool** | True to lock (prevent transfer), false to unlock | 
**TwoFaCode** | Pointer to **string** | OTP from email/TOTP when upstream returns 403 two-factor required; forwarded as 2fa_code | [optional] 
**TwoFaBackup** | Pointer to **string** | Backup 2FA code when OTP is unavailable; forwarded as 2fa_backup | [optional] 

## Methods

### NewUpdateDomainLockStatusRequestContent

`func NewUpdateDomainLockStatusRequestContent(domainId string, lockStatus bool, ) *UpdateDomainLockStatusRequestContent`

NewUpdateDomainLockStatusRequestContent instantiates a new UpdateDomainLockStatusRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDomainLockStatusRequestContentWithDefaults

`func NewUpdateDomainLockStatusRequestContentWithDefaults() *UpdateDomainLockStatusRequestContent`

NewUpdateDomainLockStatusRequestContentWithDefaults instantiates a new UpdateDomainLockStatusRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomainId

`func (o *UpdateDomainLockStatusRequestContent) GetDomainId() string`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *UpdateDomainLockStatusRequestContent) GetDomainIdOk() (*string, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *UpdateDomainLockStatusRequestContent) SetDomainId(v string)`

SetDomainId sets DomainId field to given value.


### GetLockStatus

`func (o *UpdateDomainLockStatusRequestContent) GetLockStatus() bool`

GetLockStatus returns the LockStatus field if non-nil, zero value otherwise.

### GetLockStatusOk

`func (o *UpdateDomainLockStatusRequestContent) GetLockStatusOk() (*bool, bool)`

GetLockStatusOk returns a tuple with the LockStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockStatus

`func (o *UpdateDomainLockStatusRequestContent) SetLockStatus(v bool)`

SetLockStatus sets LockStatus field to given value.


### GetTwoFaCode

`func (o *UpdateDomainLockStatusRequestContent) GetTwoFaCode() string`

GetTwoFaCode returns the TwoFaCode field if non-nil, zero value otherwise.

### GetTwoFaCodeOk

`func (o *UpdateDomainLockStatusRequestContent) GetTwoFaCodeOk() (*string, bool)`

GetTwoFaCodeOk returns a tuple with the TwoFaCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoFaCode

`func (o *UpdateDomainLockStatusRequestContent) SetTwoFaCode(v string)`

SetTwoFaCode sets TwoFaCode field to given value.

### HasTwoFaCode

`func (o *UpdateDomainLockStatusRequestContent) HasTwoFaCode() bool`

HasTwoFaCode returns a boolean if a field has been set.

### GetTwoFaBackup

`func (o *UpdateDomainLockStatusRequestContent) GetTwoFaBackup() string`

GetTwoFaBackup returns the TwoFaBackup field if non-nil, zero value otherwise.

### GetTwoFaBackupOk

`func (o *UpdateDomainLockStatusRequestContent) GetTwoFaBackupOk() (*string, bool)`

GetTwoFaBackupOk returns a tuple with the TwoFaBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoFaBackup

`func (o *UpdateDomainLockStatusRequestContent) SetTwoFaBackup(v string)`

SetTwoFaBackup sets TwoFaBackup field to given value.

### HasTwoFaBackup

`func (o *UpdateDomainLockStatusRequestContent) HasTwoFaBackup() bool`

HasTwoFaBackup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


