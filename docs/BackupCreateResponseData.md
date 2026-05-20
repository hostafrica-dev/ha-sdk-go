# BackupCreateResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Status message indicating the result | 
**TaskId** | **int32** | Task ID for the backup creation process | 
**BackupIsCreating** | **bool** | Whether a backup is currently being created | 
**BackupCreation** | [**BackupCreationInfo**](BackupCreationInfo.md) |  | 

## Methods

### NewBackupCreateResponseData

`func NewBackupCreateResponseData(message string, taskId int32, backupIsCreating bool, backupCreation BackupCreationInfo, ) *BackupCreateResponseData`

NewBackupCreateResponseData instantiates a new BackupCreateResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupCreateResponseDataWithDefaults

`func NewBackupCreateResponseDataWithDefaults() *BackupCreateResponseData`

NewBackupCreateResponseDataWithDefaults instantiates a new BackupCreateResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *BackupCreateResponseData) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *BackupCreateResponseData) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *BackupCreateResponseData) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetTaskId

`func (o *BackupCreateResponseData) GetTaskId() int32`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *BackupCreateResponseData) GetTaskIdOk() (*int32, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *BackupCreateResponseData) SetTaskId(v int32)`

SetTaskId sets TaskId field to given value.


### GetBackupIsCreating

`func (o *BackupCreateResponseData) GetBackupIsCreating() bool`

GetBackupIsCreating returns the BackupIsCreating field if non-nil, zero value otherwise.

### GetBackupIsCreatingOk

`func (o *BackupCreateResponseData) GetBackupIsCreatingOk() (*bool, bool)`

GetBackupIsCreatingOk returns a tuple with the BackupIsCreating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupIsCreating

`func (o *BackupCreateResponseData) SetBackupIsCreating(v bool)`

SetBackupIsCreating sets BackupIsCreating field to given value.


### GetBackupCreation

`func (o *BackupCreateResponseData) GetBackupCreation() BackupCreationInfo`

GetBackupCreation returns the BackupCreation field if non-nil, zero value otherwise.

### GetBackupCreationOk

`func (o *BackupCreateResponseData) GetBackupCreationOk() (*BackupCreationInfo, bool)`

GetBackupCreationOk returns a tuple with the BackupCreation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupCreation

`func (o *BackupCreateResponseData) SetBackupCreation(v BackupCreationInfo)`

SetBackupCreation sets BackupCreation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


