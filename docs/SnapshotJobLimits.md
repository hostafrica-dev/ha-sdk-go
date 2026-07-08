# SnapshotJobLimits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxJobs** | **int32** | Maximum number of snapshot jobs allowed | 
**JobCount** | **int32** | Current number of snapshot jobs | 
**CanAddMore** | **bool** | Whether more jobs can be added | 

## Methods

### NewSnapshotJobLimits

`func NewSnapshotJobLimits(maxJobs int32, jobCount int32, canAddMore bool, ) *SnapshotJobLimits`

NewSnapshotJobLimits instantiates a new SnapshotJobLimits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotJobLimitsWithDefaults

`func NewSnapshotJobLimitsWithDefaults() *SnapshotJobLimits`

NewSnapshotJobLimitsWithDefaults instantiates a new SnapshotJobLimits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxJobs

`func (o *SnapshotJobLimits) GetMaxJobs() int32`

GetMaxJobs returns the MaxJobs field if non-nil, zero value otherwise.

### GetMaxJobsOk

`func (o *SnapshotJobLimits) GetMaxJobsOk() (*int32, bool)`

GetMaxJobsOk returns a tuple with the MaxJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxJobs

`func (o *SnapshotJobLimits) SetMaxJobs(v int32)`

SetMaxJobs sets MaxJobs field to given value.


### GetJobCount

`func (o *SnapshotJobLimits) GetJobCount() int32`

GetJobCount returns the JobCount field if non-nil, zero value otherwise.

### GetJobCountOk

`func (o *SnapshotJobLimits) GetJobCountOk() (*int32, bool)`

GetJobCountOk returns a tuple with the JobCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobCount

`func (o *SnapshotJobLimits) SetJobCount(v int32)`

SetJobCount sets JobCount field to given value.


### GetCanAddMore

`func (o *SnapshotJobLimits) GetCanAddMore() bool`

GetCanAddMore returns the CanAddMore field if non-nil, zero value otherwise.

### GetCanAddMoreOk

`func (o *SnapshotJobLimits) GetCanAddMoreOk() (*bool, bool)`

GetCanAddMoreOk returns a tuple with the CanAddMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanAddMore

`func (o *SnapshotJobLimits) SetCanAddMore(v bool)`

SetCanAddMore sets CanAddMore field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


