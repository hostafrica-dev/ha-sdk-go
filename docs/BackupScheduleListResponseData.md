# BackupScheduleListResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Status message indicating the result | 
**Schedules** | [**[]BackupSchedule**](BackupSchedule.md) | List of backup schedules | 

## Methods

### NewBackupScheduleListResponseData

`func NewBackupScheduleListResponseData(message string, schedules []BackupSchedule, ) *BackupScheduleListResponseData`

NewBackupScheduleListResponseData instantiates a new BackupScheduleListResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupScheduleListResponseDataWithDefaults

`func NewBackupScheduleListResponseDataWithDefaults() *BackupScheduleListResponseData`

NewBackupScheduleListResponseDataWithDefaults instantiates a new BackupScheduleListResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *BackupScheduleListResponseData) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *BackupScheduleListResponseData) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *BackupScheduleListResponseData) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetSchedules

`func (o *BackupScheduleListResponseData) GetSchedules() []BackupSchedule`

GetSchedules returns the Schedules field if non-nil, zero value otherwise.

### GetSchedulesOk

`func (o *BackupScheduleListResponseData) GetSchedulesOk() (*[]BackupSchedule, bool)`

GetSchedulesOk returns a tuple with the Schedules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedules

`func (o *BackupScheduleListResponseData) SetSchedules(v []BackupSchedule)`

SetSchedules sets Schedules field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


