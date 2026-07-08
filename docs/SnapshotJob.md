# SnapshotJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Snapshot job ID | 
**HostingId** | Pointer to **int32** | Hosting account ID | [optional] 
**VmId** | Pointer to **int32** | VM ID (null if not yet assigned) | [optional] 
**Name** | **string** | Name of the snapshot job | 
**Description** | Pointer to **string** | Description of the snapshot job | [optional] 
**Vmstate** | Pointer to **bool** | Whether VM state is included in the snapshot | [optional] 
**Period** | [**SnapshotJobPeriod**](SnapshotJobPeriod.md) |  | 
**RunEvery** | Pointer to **int32** | For hourly jobs: run every N hours | [optional] 
**Days** | Pointer to [**[]DayOfWeek**](DayOfWeek.md) | For daily jobs: days of week | [optional] 
**StartTime** | Pointer to **string** | For daily jobs: start time in HH:MM format | [optional] 

## Methods

### NewSnapshotJob

`func NewSnapshotJob(id int32, name string, period SnapshotJobPeriod, ) *SnapshotJob`

NewSnapshotJob instantiates a new SnapshotJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotJobWithDefaults

`func NewSnapshotJobWithDefaults() *SnapshotJob`

NewSnapshotJobWithDefaults instantiates a new SnapshotJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SnapshotJob) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SnapshotJob) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SnapshotJob) SetId(v int32)`

SetId sets Id field to given value.


### GetHostingId

`func (o *SnapshotJob) GetHostingId() int32`

GetHostingId returns the HostingId field if non-nil, zero value otherwise.

### GetHostingIdOk

`func (o *SnapshotJob) GetHostingIdOk() (*int32, bool)`

GetHostingIdOk returns a tuple with the HostingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostingId

`func (o *SnapshotJob) SetHostingId(v int32)`

SetHostingId sets HostingId field to given value.

### HasHostingId

`func (o *SnapshotJob) HasHostingId() bool`

HasHostingId returns a boolean if a field has been set.

### GetVmId

`func (o *SnapshotJob) GetVmId() int32`

GetVmId returns the VmId field if non-nil, zero value otherwise.

### GetVmIdOk

`func (o *SnapshotJob) GetVmIdOk() (*int32, bool)`

GetVmIdOk returns a tuple with the VmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmId

`func (o *SnapshotJob) SetVmId(v int32)`

SetVmId sets VmId field to given value.

### HasVmId

`func (o *SnapshotJob) HasVmId() bool`

HasVmId returns a boolean if a field has been set.

### GetName

`func (o *SnapshotJob) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SnapshotJob) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SnapshotJob) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *SnapshotJob) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SnapshotJob) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SnapshotJob) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SnapshotJob) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVmstate

`func (o *SnapshotJob) GetVmstate() bool`

GetVmstate returns the Vmstate field if non-nil, zero value otherwise.

### GetVmstateOk

`func (o *SnapshotJob) GetVmstateOk() (*bool, bool)`

GetVmstateOk returns a tuple with the Vmstate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmstate

`func (o *SnapshotJob) SetVmstate(v bool)`

SetVmstate sets Vmstate field to given value.

### HasVmstate

`func (o *SnapshotJob) HasVmstate() bool`

HasVmstate returns a boolean if a field has been set.

### GetPeriod

`func (o *SnapshotJob) GetPeriod() SnapshotJobPeriod`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *SnapshotJob) GetPeriodOk() (*SnapshotJobPeriod, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *SnapshotJob) SetPeriod(v SnapshotJobPeriod)`

SetPeriod sets Period field to given value.


### GetRunEvery

`func (o *SnapshotJob) GetRunEvery() int32`

GetRunEvery returns the RunEvery field if non-nil, zero value otherwise.

### GetRunEveryOk

`func (o *SnapshotJob) GetRunEveryOk() (*int32, bool)`

GetRunEveryOk returns a tuple with the RunEvery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunEvery

`func (o *SnapshotJob) SetRunEvery(v int32)`

SetRunEvery sets RunEvery field to given value.

### HasRunEvery

`func (o *SnapshotJob) HasRunEvery() bool`

HasRunEvery returns a boolean if a field has been set.

### GetDays

`func (o *SnapshotJob) GetDays() []DayOfWeek`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *SnapshotJob) GetDaysOk() (*[]DayOfWeek, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *SnapshotJob) SetDays(v []DayOfWeek)`

SetDays sets Days field to given value.

### HasDays

`func (o *SnapshotJob) HasDays() bool`

HasDays returns a boolean if a field has been set.

### GetStartTime

`func (o *SnapshotJob) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *SnapshotJob) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *SnapshotJob) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *SnapshotJob) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


