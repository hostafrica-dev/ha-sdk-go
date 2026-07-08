# CreateSnapshotJobRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceId** | **string** | Service ID - must be sent as a string | 
**Name** | **string** | Name for the snapshot job (e.g. &#39;auto_hourly&#39;) | 
**Description** | Pointer to **string** | Description for the snapshot job | [optional] 
**Vmstate** | Pointer to **bool** | Whether to include VM state in the snapshot | [optional] 
**Period** | [**SnapshotJobPeriod**](SnapshotJobPeriod.md) |  | 
**RunEvery** | Pointer to **int32** | For hourly jobs: run every N hours (e.g. 6 &#x3D; every 6 hours) | [optional] 
**Days** | Pointer to [**[]DayOfWeek**](DayOfWeek.md) | For daily jobs: days of week when the job should run | [optional] 
**StartTime** | Pointer to **string** | For daily jobs: start time in HH:MM format (e.g. &#39;02:30&#39;) | [optional] 

## Methods

### NewCreateSnapshotJobRequestContent

`func NewCreateSnapshotJobRequestContent(serviceId string, name string, period SnapshotJobPeriod, ) *CreateSnapshotJobRequestContent`

NewCreateSnapshotJobRequestContent instantiates a new CreateSnapshotJobRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSnapshotJobRequestContentWithDefaults

`func NewCreateSnapshotJobRequestContentWithDefaults() *CreateSnapshotJobRequestContent`

NewCreateSnapshotJobRequestContentWithDefaults instantiates a new CreateSnapshotJobRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceId

`func (o *CreateSnapshotJobRequestContent) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *CreateSnapshotJobRequestContent) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *CreateSnapshotJobRequestContent) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetName

`func (o *CreateSnapshotJobRequestContent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateSnapshotJobRequestContent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateSnapshotJobRequestContent) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateSnapshotJobRequestContent) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateSnapshotJobRequestContent) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateSnapshotJobRequestContent) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateSnapshotJobRequestContent) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVmstate

`func (o *CreateSnapshotJobRequestContent) GetVmstate() bool`

GetVmstate returns the Vmstate field if non-nil, zero value otherwise.

### GetVmstateOk

`func (o *CreateSnapshotJobRequestContent) GetVmstateOk() (*bool, bool)`

GetVmstateOk returns a tuple with the Vmstate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmstate

`func (o *CreateSnapshotJobRequestContent) SetVmstate(v bool)`

SetVmstate sets Vmstate field to given value.

### HasVmstate

`func (o *CreateSnapshotJobRequestContent) HasVmstate() bool`

HasVmstate returns a boolean if a field has been set.

### GetPeriod

`func (o *CreateSnapshotJobRequestContent) GetPeriod() SnapshotJobPeriod`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *CreateSnapshotJobRequestContent) GetPeriodOk() (*SnapshotJobPeriod, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *CreateSnapshotJobRequestContent) SetPeriod(v SnapshotJobPeriod)`

SetPeriod sets Period field to given value.


### GetRunEvery

`func (o *CreateSnapshotJobRequestContent) GetRunEvery() int32`

GetRunEvery returns the RunEvery field if non-nil, zero value otherwise.

### GetRunEveryOk

`func (o *CreateSnapshotJobRequestContent) GetRunEveryOk() (*int32, bool)`

GetRunEveryOk returns a tuple with the RunEvery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunEvery

`func (o *CreateSnapshotJobRequestContent) SetRunEvery(v int32)`

SetRunEvery sets RunEvery field to given value.

### HasRunEvery

`func (o *CreateSnapshotJobRequestContent) HasRunEvery() bool`

HasRunEvery returns a boolean if a field has been set.

### GetDays

`func (o *CreateSnapshotJobRequestContent) GetDays() []DayOfWeek`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *CreateSnapshotJobRequestContent) GetDaysOk() (*[]DayOfWeek, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *CreateSnapshotJobRequestContent) SetDays(v []DayOfWeek)`

SetDays sets Days field to given value.

### HasDays

`func (o *CreateSnapshotJobRequestContent) HasDays() bool`

HasDays returns a boolean if a field has been set.

### GetStartTime

`func (o *CreateSnapshotJobRequestContent) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *CreateSnapshotJobRequestContent) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *CreateSnapshotJobRequestContent) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *CreateSnapshotJobRequestContent) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


