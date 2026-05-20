# PowerTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Task ID | 
**HostingId** | **int32** | Hosting/service ID | 
**Description** | Pointer to **string** | Task description | [optional] 
**Action** | **string** | Power action (start, stop, restart) | 
**Start** | **string** | Task start date/time (ISO 8601 format) | 
**End** | Pointer to **string** | Task end date/time (ISO 8601 format, null if no end) | [optional] 
**JobType** | **string** | Job type (oneTime, daily, weekly) | 
**JobTime** | **string** | Job execution time (HH:MM:SS format) | 
**Days** | **[]string** | Days of week for weekly jobs (empty array for oneTime/daily) | 
**LastRun** | Pointer to **string** | Last run timestamp (null if never run) | [optional] 

## Methods

### NewPowerTask

`func NewPowerTask(id int32, hostingId int32, action string, start string, jobType string, jobTime string, days []string, ) *PowerTask`

NewPowerTask instantiates a new PowerTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPowerTaskWithDefaults

`func NewPowerTaskWithDefaults() *PowerTask`

NewPowerTaskWithDefaults instantiates a new PowerTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PowerTask) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PowerTask) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PowerTask) SetId(v int32)`

SetId sets Id field to given value.


### GetHostingId

`func (o *PowerTask) GetHostingId() int32`

GetHostingId returns the HostingId field if non-nil, zero value otherwise.

### GetHostingIdOk

`func (o *PowerTask) GetHostingIdOk() (*int32, bool)`

GetHostingIdOk returns a tuple with the HostingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostingId

`func (o *PowerTask) SetHostingId(v int32)`

SetHostingId sets HostingId field to given value.


### GetDescription

`func (o *PowerTask) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PowerTask) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PowerTask) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PowerTask) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAction

`func (o *PowerTask) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *PowerTask) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *PowerTask) SetAction(v string)`

SetAction sets Action field to given value.


### GetStart

`func (o *PowerTask) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *PowerTask) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *PowerTask) SetStart(v string)`

SetStart sets Start field to given value.


### GetEnd

`func (o *PowerTask) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *PowerTask) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *PowerTask) SetEnd(v string)`

SetEnd sets End field to given value.

### HasEnd

`func (o *PowerTask) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetJobType

`func (o *PowerTask) GetJobType() string`

GetJobType returns the JobType field if non-nil, zero value otherwise.

### GetJobTypeOk

`func (o *PowerTask) GetJobTypeOk() (*string, bool)`

GetJobTypeOk returns a tuple with the JobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobType

`func (o *PowerTask) SetJobType(v string)`

SetJobType sets JobType field to given value.


### GetJobTime

`func (o *PowerTask) GetJobTime() string`

GetJobTime returns the JobTime field if non-nil, zero value otherwise.

### GetJobTimeOk

`func (o *PowerTask) GetJobTimeOk() (*string, bool)`

GetJobTimeOk returns a tuple with the JobTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTime

`func (o *PowerTask) SetJobTime(v string)`

SetJobTime sets JobTime field to given value.


### GetDays

`func (o *PowerTask) GetDays() []string`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *PowerTask) GetDaysOk() (*[]string, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *PowerTask) SetDays(v []string)`

SetDays sets Days field to given value.


### GetLastRun

`func (o *PowerTask) GetLastRun() string`

GetLastRun returns the LastRun field if non-nil, zero value otherwise.

### GetLastRunOk

`func (o *PowerTask) GetLastRunOk() (*string, bool)`

GetLastRunOk returns a tuple with the LastRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRun

`func (o *PowerTask) SetLastRun(v string)`

SetLastRun sets LastRun field to given value.

### HasLastRun

`func (o *PowerTask) HasLastRun() bool`

HasLastRun returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


