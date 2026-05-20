# UpdatePowerTaskRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceId** | **string** | Service ID - must be sent as a string | 
**TaskId** | **int32** | Power task ID to update | 
**PowerTaskAction** | Pointer to **string** | Power action to perform (start, stop, restart) | [optional] 
**StartDate** | Pointer to **string** | Start date in Y-m-d format (e.g., 2026-03-25) | [optional] 
**Description** | Pointer to **string** | Description of the power task | [optional] 
**StartTime** | Pointer to **string** | Start time in HH:MM or HH:MM:SS format | [optional] 
**EndDate** | Pointer to **string** | End date in Y-m-d format (e.g., 2026-12-31) | [optional] 
**EndTime** | Pointer to **string** | End time in HH:MM or HH:MM:SS format | [optional] 
**JobType** | Pointer to [**PowerTaskJobType**](PowerTaskJobType.md) |  | [optional] 
**JobTime** | Pointer to **string** | Job execution time in HH:MM or HH:MM:SS format | [optional] 
**JobHour** | Pointer to **int32** | Job hour (alternative to job_time) | [optional] 
**JobMinutes** | Pointer to **int32** | Job minutes (alternative to job_time) | [optional] 
**Days** | Pointer to [**[]DayOfWeek**](DayOfWeek.md) | Days of the week for weekly jobs | [optional] 

## Methods

### NewUpdatePowerTaskRequestContent

`func NewUpdatePowerTaskRequestContent(serviceId string, taskId int32, ) *UpdatePowerTaskRequestContent`

NewUpdatePowerTaskRequestContent instantiates a new UpdatePowerTaskRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePowerTaskRequestContentWithDefaults

`func NewUpdatePowerTaskRequestContentWithDefaults() *UpdatePowerTaskRequestContent`

NewUpdatePowerTaskRequestContentWithDefaults instantiates a new UpdatePowerTaskRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceId

`func (o *UpdatePowerTaskRequestContent) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *UpdatePowerTaskRequestContent) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *UpdatePowerTaskRequestContent) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetTaskId

`func (o *UpdatePowerTaskRequestContent) GetTaskId() int32`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *UpdatePowerTaskRequestContent) GetTaskIdOk() (*int32, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *UpdatePowerTaskRequestContent) SetTaskId(v int32)`

SetTaskId sets TaskId field to given value.


### GetPowerTaskAction

`func (o *UpdatePowerTaskRequestContent) GetPowerTaskAction() string`

GetPowerTaskAction returns the PowerTaskAction field if non-nil, zero value otherwise.

### GetPowerTaskActionOk

`func (o *UpdatePowerTaskRequestContent) GetPowerTaskActionOk() (*string, bool)`

GetPowerTaskActionOk returns a tuple with the PowerTaskAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerTaskAction

`func (o *UpdatePowerTaskRequestContent) SetPowerTaskAction(v string)`

SetPowerTaskAction sets PowerTaskAction field to given value.

### HasPowerTaskAction

`func (o *UpdatePowerTaskRequestContent) HasPowerTaskAction() bool`

HasPowerTaskAction returns a boolean if a field has been set.

### GetStartDate

`func (o *UpdatePowerTaskRequestContent) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *UpdatePowerTaskRequestContent) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *UpdatePowerTaskRequestContent) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *UpdatePowerTaskRequestContent) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetDescription

`func (o *UpdatePowerTaskRequestContent) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdatePowerTaskRequestContent) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdatePowerTaskRequestContent) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdatePowerTaskRequestContent) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStartTime

`func (o *UpdatePowerTaskRequestContent) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *UpdatePowerTaskRequestContent) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *UpdatePowerTaskRequestContent) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *UpdatePowerTaskRequestContent) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndDate

`func (o *UpdatePowerTaskRequestContent) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *UpdatePowerTaskRequestContent) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *UpdatePowerTaskRequestContent) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *UpdatePowerTaskRequestContent) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetEndTime

`func (o *UpdatePowerTaskRequestContent) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *UpdatePowerTaskRequestContent) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *UpdatePowerTaskRequestContent) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *UpdatePowerTaskRequestContent) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetJobType

`func (o *UpdatePowerTaskRequestContent) GetJobType() PowerTaskJobType`

GetJobType returns the JobType field if non-nil, zero value otherwise.

### GetJobTypeOk

`func (o *UpdatePowerTaskRequestContent) GetJobTypeOk() (*PowerTaskJobType, bool)`

GetJobTypeOk returns a tuple with the JobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobType

`func (o *UpdatePowerTaskRequestContent) SetJobType(v PowerTaskJobType)`

SetJobType sets JobType field to given value.

### HasJobType

`func (o *UpdatePowerTaskRequestContent) HasJobType() bool`

HasJobType returns a boolean if a field has been set.

### GetJobTime

`func (o *UpdatePowerTaskRequestContent) GetJobTime() string`

GetJobTime returns the JobTime field if non-nil, zero value otherwise.

### GetJobTimeOk

`func (o *UpdatePowerTaskRequestContent) GetJobTimeOk() (*string, bool)`

GetJobTimeOk returns a tuple with the JobTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTime

`func (o *UpdatePowerTaskRequestContent) SetJobTime(v string)`

SetJobTime sets JobTime field to given value.

### HasJobTime

`func (o *UpdatePowerTaskRequestContent) HasJobTime() bool`

HasJobTime returns a boolean if a field has been set.

### GetJobHour

`func (o *UpdatePowerTaskRequestContent) GetJobHour() int32`

GetJobHour returns the JobHour field if non-nil, zero value otherwise.

### GetJobHourOk

`func (o *UpdatePowerTaskRequestContent) GetJobHourOk() (*int32, bool)`

GetJobHourOk returns a tuple with the JobHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobHour

`func (o *UpdatePowerTaskRequestContent) SetJobHour(v int32)`

SetJobHour sets JobHour field to given value.

### HasJobHour

`func (o *UpdatePowerTaskRequestContent) HasJobHour() bool`

HasJobHour returns a boolean if a field has been set.

### GetJobMinutes

`func (o *UpdatePowerTaskRequestContent) GetJobMinutes() int32`

GetJobMinutes returns the JobMinutes field if non-nil, zero value otherwise.

### GetJobMinutesOk

`func (o *UpdatePowerTaskRequestContent) GetJobMinutesOk() (*int32, bool)`

GetJobMinutesOk returns a tuple with the JobMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobMinutes

`func (o *UpdatePowerTaskRequestContent) SetJobMinutes(v int32)`

SetJobMinutes sets JobMinutes field to given value.

### HasJobMinutes

`func (o *UpdatePowerTaskRequestContent) HasJobMinutes() bool`

HasJobMinutes returns a boolean if a field has been set.

### GetDays

`func (o *UpdatePowerTaskRequestContent) GetDays() []DayOfWeek`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *UpdatePowerTaskRequestContent) GetDaysOk() (*[]DayOfWeek, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *UpdatePowerTaskRequestContent) SetDays(v []DayOfWeek)`

SetDays sets Days field to given value.

### HasDays

`func (o *UpdatePowerTaskRequestContent) HasDays() bool`

HasDays returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


