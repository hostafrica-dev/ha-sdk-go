# ServiceBackupsResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Status message indicating the result | 
**Backups** | [**[]BackupItem**](BackupItem.md) | List of backups for the service | 
**QuotaTotal** | Pointer to **float64** | Total backup quota (null if unlimited) | [optional] 
**QuotaUsed** | **float64** | Used backup quota | 
**QuotaUnit** | **string** | Unit for quota measurements (e.g., GiB) | 
**AvailableCompressMethods** | [**[]CompressionMethod**](CompressionMethod.md) | Available compression methods for creating backups | 
**AvailableModes** | [**[]BackupMode**](BackupMode.md) | Available backup modes | 
**BackupIsCreating** | **bool** | Whether a backup is currently being created | 
**BackupCreation** | Pointer to [**BackupCreationInfo**](BackupCreationInfo.md) |  | [optional] 

## Methods

### NewServiceBackupsResponseData

`func NewServiceBackupsResponseData(message string, backups []BackupItem, quotaUsed float64, quotaUnit string, availableCompressMethods []CompressionMethod, availableModes []BackupMode, backupIsCreating bool, ) *ServiceBackupsResponseData`

NewServiceBackupsResponseData instantiates a new ServiceBackupsResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceBackupsResponseDataWithDefaults

`func NewServiceBackupsResponseDataWithDefaults() *ServiceBackupsResponseData`

NewServiceBackupsResponseDataWithDefaults instantiates a new ServiceBackupsResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ServiceBackupsResponseData) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ServiceBackupsResponseData) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ServiceBackupsResponseData) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetBackups

`func (o *ServiceBackupsResponseData) GetBackups() []BackupItem`

GetBackups returns the Backups field if non-nil, zero value otherwise.

### GetBackupsOk

`func (o *ServiceBackupsResponseData) GetBackupsOk() (*[]BackupItem, bool)`

GetBackupsOk returns a tuple with the Backups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackups

`func (o *ServiceBackupsResponseData) SetBackups(v []BackupItem)`

SetBackups sets Backups field to given value.


### GetQuotaTotal

`func (o *ServiceBackupsResponseData) GetQuotaTotal() float64`

GetQuotaTotal returns the QuotaTotal field if non-nil, zero value otherwise.

### GetQuotaTotalOk

`func (o *ServiceBackupsResponseData) GetQuotaTotalOk() (*float64, bool)`

GetQuotaTotalOk returns a tuple with the QuotaTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaTotal

`func (o *ServiceBackupsResponseData) SetQuotaTotal(v float64)`

SetQuotaTotal sets QuotaTotal field to given value.

### HasQuotaTotal

`func (o *ServiceBackupsResponseData) HasQuotaTotal() bool`

HasQuotaTotal returns a boolean if a field has been set.

### GetQuotaUsed

`func (o *ServiceBackupsResponseData) GetQuotaUsed() float64`

GetQuotaUsed returns the QuotaUsed field if non-nil, zero value otherwise.

### GetQuotaUsedOk

`func (o *ServiceBackupsResponseData) GetQuotaUsedOk() (*float64, bool)`

GetQuotaUsedOk returns a tuple with the QuotaUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaUsed

`func (o *ServiceBackupsResponseData) SetQuotaUsed(v float64)`

SetQuotaUsed sets QuotaUsed field to given value.


### GetQuotaUnit

`func (o *ServiceBackupsResponseData) GetQuotaUnit() string`

GetQuotaUnit returns the QuotaUnit field if non-nil, zero value otherwise.

### GetQuotaUnitOk

`func (o *ServiceBackupsResponseData) GetQuotaUnitOk() (*string, bool)`

GetQuotaUnitOk returns a tuple with the QuotaUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaUnit

`func (o *ServiceBackupsResponseData) SetQuotaUnit(v string)`

SetQuotaUnit sets QuotaUnit field to given value.


### GetAvailableCompressMethods

`func (o *ServiceBackupsResponseData) GetAvailableCompressMethods() []CompressionMethod`

GetAvailableCompressMethods returns the AvailableCompressMethods field if non-nil, zero value otherwise.

### GetAvailableCompressMethodsOk

`func (o *ServiceBackupsResponseData) GetAvailableCompressMethodsOk() (*[]CompressionMethod, bool)`

GetAvailableCompressMethodsOk returns a tuple with the AvailableCompressMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableCompressMethods

`func (o *ServiceBackupsResponseData) SetAvailableCompressMethods(v []CompressionMethod)`

SetAvailableCompressMethods sets AvailableCompressMethods field to given value.


### GetAvailableModes

`func (o *ServiceBackupsResponseData) GetAvailableModes() []BackupMode`

GetAvailableModes returns the AvailableModes field if non-nil, zero value otherwise.

### GetAvailableModesOk

`func (o *ServiceBackupsResponseData) GetAvailableModesOk() (*[]BackupMode, bool)`

GetAvailableModesOk returns a tuple with the AvailableModes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableModes

`func (o *ServiceBackupsResponseData) SetAvailableModes(v []BackupMode)`

SetAvailableModes sets AvailableModes field to given value.


### GetBackupIsCreating

`func (o *ServiceBackupsResponseData) GetBackupIsCreating() bool`

GetBackupIsCreating returns the BackupIsCreating field if non-nil, zero value otherwise.

### GetBackupIsCreatingOk

`func (o *ServiceBackupsResponseData) GetBackupIsCreatingOk() (*bool, bool)`

GetBackupIsCreatingOk returns a tuple with the BackupIsCreating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupIsCreating

`func (o *ServiceBackupsResponseData) SetBackupIsCreating(v bool)`

SetBackupIsCreating sets BackupIsCreating field to given value.


### GetBackupCreation

`func (o *ServiceBackupsResponseData) GetBackupCreation() BackupCreationInfo`

GetBackupCreation returns the BackupCreation field if non-nil, zero value otherwise.

### GetBackupCreationOk

`func (o *ServiceBackupsResponseData) GetBackupCreationOk() (*BackupCreationInfo, bool)`

GetBackupCreationOk returns a tuple with the BackupCreation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupCreation

`func (o *ServiceBackupsResponseData) SetBackupCreation(v BackupCreationInfo)`

SetBackupCreation sets BackupCreation field to given value.

### HasBackupCreation

`func (o *ServiceBackupsResponseData) HasBackupCreation() bool`

HasBackupCreation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


