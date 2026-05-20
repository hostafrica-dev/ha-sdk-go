# CreatePowerTaskRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceId** | **string** | Service ID - must be sent as a string | 
**PowerTaskAction** | [**PowerTaskAction**](PowerTaskAction.md) |  | 
**StartDate** | **string** | Start date in Y-m-d format (e.g., 2026-03-25) | 
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

### NewCreatePowerTaskRequestContent

`func NewCreatePowerTaskRequestContent(serviceId string, powerTaskAction PowerTaskAction, startDate string, ) *CreatePowerTaskRequestContent`

NewCreatePowerTaskRequestContent instantiates a new CreatePowerTaskRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePowerTaskRequestContentWithDefaults

`func NewCreatePowerTaskRequestContentWithDefaults() *CreatePowerTaskRequestContent`

NewCreatePowerTaskRequestContentWithDefaults instantiates a new CreatePowerTaskRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceId

`func (o *CreatePowerTaskRequestContent) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *CreatePowerTaskRequestContent) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *CreatePowerTaskRequestContent) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetPowerTaskAction

`func (o *CreatePowerTaskRequestContent) GetPowerTaskAction() PowerTaskAction`

GetPowerTaskAction returns the PowerTaskAction field if non-nil, zero value otherwise.

### GetPowerTaskActionOk

`func (o *CreatePowerTaskRequestContent) GetPowerTaskActionOk() (*PowerTaskAction, bool)`

GetPowerTaskActionOk returns a tuple with the PowerTaskAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerTaskAction

`func (o *CreatePowerTaskRequestContent) SetPowerTaskAction(v PowerTaskAction)`

SetPowerTaskAction sets PowerTaskAction field to given value.


### GetStartDate

`func (o *CreatePowerTaskRequestContent) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *CreatePowerTaskRequestContent) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *CreatePowerTaskRequestContent) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.


### GetDescription

`func (o *CreatePowerTaskRequestContent) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreatePowerTaskRequestContent) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreatePowerTaskRequestContent) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreatePowerTaskRequestContent) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStartTime

`func (o *CreatePowerTaskRequestContent) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *CreatePowerTaskRequestContent) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *CreatePowerTaskRequestContent) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *CreatePowerTaskRequestContent) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndDate

`func (o *CreatePowerTaskRequestContent) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *CreatePowerTaskRequestContent) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *CreatePowerTaskRequestContent) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *CreatePowerTaskRequestContent) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetEndTime

`func (o *CreatePowerTaskRequestContent) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *CreatePowerTaskRequestContent) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *CreatePowerTaskRequestContent) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *CreatePowerTaskRequestContent) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetJobType

`func (o *CreatePowerTaskRequestContent) GetJobType() PowerTaskJobType`

GetJobType returns the JobType field if non-nil, zero value otherwise.

### GetJobTypeOk

`func (o *CreatePowerTaskRequestContent) GetJobTypeOk() (*PowerTaskJobType, bool)`

GetJobTypeOk returns a tuple with the JobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobType

`func (o *CreatePowerTaskRequestContent) SetJobType(v PowerTaskJobType)`

SetJobType sets JobType field to given value.

### HasJobType

`func (o *CreatePowerTaskRequestContent) HasJobType() bool`

HasJobType returns a boolean if a field has been set.

### GetJobTime

`func (o *CreatePowerTaskRequestContent) GetJobTime() string`

GetJobTime returns the JobTime field if non-nil, zero value otherwise.

### GetJobTimeOk

`func (o *CreatePowerTaskRequestContent) GetJobTimeOk() (*string, bool)`

GetJobTimeOk returns a tuple with the JobTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTime

`func (o *CreatePowerTaskRequestContent) SetJobTime(v string)`

SetJobTime sets JobTime field to given value.

### HasJobTime

`func (o *CreatePowerTaskRequestContent) HasJobTime() bool`

HasJobTime returns a boolean if a field has been set.

### GetJobHour

`func (o *CreatePowerTaskRequestContent) GetJobHour() int32`

GetJobHour returns the JobHour field if non-nil, zero value otherwise.

### GetJobHourOk

`func (o *CreatePowerTaskRequestContent) GetJobHourOk() (*int32, bool)`

GetJobHourOk returns a tuple with the JobHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobHour

`func (o *CreatePowerTaskRequestContent) SetJobHour(v int32)`

SetJobHour sets JobHour field to given value.

### HasJobHour

`func (o *CreatePowerTaskRequestContent) HasJobHour() bool`

HasJobHour returns a boolean if a field has been set.

### GetJobMinutes

`func (o *CreatePowerTaskRequestContent) GetJobMinutes() int32`

GetJobMinutes returns the JobMinutes field if non-nil, zero value otherwise.

### GetJobMinutesOk

`func (o *CreatePowerTaskRequestContent) GetJobMinutesOk() (*int32, bool)`

GetJobMinutesOk returns a tuple with the JobMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobMinutes

`func (o *CreatePowerTaskRequestContent) SetJobMinutes(v int32)`

SetJobMinutes sets JobMinutes field to given value.

### HasJobMinutes

`func (o *CreatePowerTaskRequestContent) HasJobMinutes() bool`

HasJobMinutes returns a boolean if a field has been set.

### GetDays

`func (o *CreatePowerTaskRequestContent) GetDays() []DayOfWeek`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *CreatePowerTaskRequestContent) GetDaysOk() (*[]DayOfWeek, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *CreatePowerTaskRequestContent) SetDays(v []DayOfWeek)`

SetDays sets Days field to given value.

### HasDays

`func (o *CreatePowerTaskRequestContent) HasDays() bool`

HasDays returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


