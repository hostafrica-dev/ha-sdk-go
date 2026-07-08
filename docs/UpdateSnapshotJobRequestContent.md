# UpdateSnapshotJobRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceId** | **string** | Service ID - must be sent as a string | 
**JobId** | **string** | Snapshot job ID to update | 
**Name** | Pointer to **string** | New name for the snapshot job | [optional] 
**Description** | Pointer to **string** | New description for the snapshot job | [optional] 
**Vmstate** | Pointer to **bool** | Whether to include VM state in the snapshot | [optional] 
**Period** | Pointer to [**SnapshotJobPeriod**](SnapshotJobPeriod.md) |  | [optional] 
**RunEvery** | Pointer to **int32** | For hourly jobs: run every N hours (e.g. 6 &#x3D; every 6 hours) | [optional] 
**Days** | Pointer to [**[]DayOfWeek**](DayOfWeek.md) | For daily jobs: days of week when the job should run | [optional] 
**StartTime** | Pointer to **string** | For daily jobs: start time in HH:MM format (e.g. &#39;02:30&#39;) | [optional] 

## Methods

### NewUpdateSnapshotJobRequestContent

`func NewUpdateSnapshotJobRequestContent(serviceId string, jobId string, ) *UpdateSnapshotJobRequestContent`

NewUpdateSnapshotJobRequestContent instantiates a new UpdateSnapshotJobRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSnapshotJobRequestContentWithDefaults

`func NewUpdateSnapshotJobRequestContentWithDefaults() *UpdateSnapshotJobRequestContent`

NewUpdateSnapshotJobRequestContentWithDefaults instantiates a new UpdateSnapshotJobRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceId

`func (o *UpdateSnapshotJobRequestContent) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *UpdateSnapshotJobRequestContent) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *UpdateSnapshotJobRequestContent) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetJobId

`func (o *UpdateSnapshotJobRequestContent) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *UpdateSnapshotJobRequestContent) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *UpdateSnapshotJobRequestContent) SetJobId(v string)`

SetJobId sets JobId field to given value.


### GetName

`func (o *UpdateSnapshotJobRequestContent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateSnapshotJobRequestContent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateSnapshotJobRequestContent) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateSnapshotJobRequestContent) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateSnapshotJobRequestContent) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateSnapshotJobRequestContent) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateSnapshotJobRequestContent) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateSnapshotJobRequestContent) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVmstate

`func (o *UpdateSnapshotJobRequestContent) GetVmstate() bool`

GetVmstate returns the Vmstate field if non-nil, zero value otherwise.

### GetVmstateOk

`func (o *UpdateSnapshotJobRequestContent) GetVmstateOk() (*bool, bool)`

GetVmstateOk returns a tuple with the Vmstate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmstate

`func (o *UpdateSnapshotJobRequestContent) SetVmstate(v bool)`

SetVmstate sets Vmstate field to given value.

### HasVmstate

`func (o *UpdateSnapshotJobRequestContent) HasVmstate() bool`

HasVmstate returns a boolean if a field has been set.

### GetPeriod

`func (o *UpdateSnapshotJobRequestContent) GetPeriod() SnapshotJobPeriod`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *UpdateSnapshotJobRequestContent) GetPeriodOk() (*SnapshotJobPeriod, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *UpdateSnapshotJobRequestContent) SetPeriod(v SnapshotJobPeriod)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *UpdateSnapshotJobRequestContent) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetRunEvery

`func (o *UpdateSnapshotJobRequestContent) GetRunEvery() int32`

GetRunEvery returns the RunEvery field if non-nil, zero value otherwise.

### GetRunEveryOk

`func (o *UpdateSnapshotJobRequestContent) GetRunEveryOk() (*int32, bool)`

GetRunEveryOk returns a tuple with the RunEvery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunEvery

`func (o *UpdateSnapshotJobRequestContent) SetRunEvery(v int32)`

SetRunEvery sets RunEvery field to given value.

### HasRunEvery

`func (o *UpdateSnapshotJobRequestContent) HasRunEvery() bool`

HasRunEvery returns a boolean if a field has been set.

### GetDays

`func (o *UpdateSnapshotJobRequestContent) GetDays() []DayOfWeek`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *UpdateSnapshotJobRequestContent) GetDaysOk() (*[]DayOfWeek, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *UpdateSnapshotJobRequestContent) SetDays(v []DayOfWeek)`

SetDays sets Days field to given value.

### HasDays

`func (o *UpdateSnapshotJobRequestContent) HasDays() bool`

HasDays returns a boolean if a field has been set.

### GetStartTime

`func (o *UpdateSnapshotJobRequestContent) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *UpdateSnapshotJobRequestContent) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *UpdateSnapshotJobRequestContent) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *UpdateSnapshotJobRequestContent) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


