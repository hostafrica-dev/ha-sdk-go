# SnapshotJobMutationResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Status message indicating the result | 
**Job** | [**SnapshotJob**](SnapshotJob.md) |  | 
**Limits** | [**SnapshotJobLimits**](SnapshotJobLimits.md) |  | 

## Methods

### NewSnapshotJobMutationResponseData

`func NewSnapshotJobMutationResponseData(message string, job SnapshotJob, limits SnapshotJobLimits, ) *SnapshotJobMutationResponseData`

NewSnapshotJobMutationResponseData instantiates a new SnapshotJobMutationResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotJobMutationResponseDataWithDefaults

`func NewSnapshotJobMutationResponseDataWithDefaults() *SnapshotJobMutationResponseData`

NewSnapshotJobMutationResponseDataWithDefaults instantiates a new SnapshotJobMutationResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *SnapshotJobMutationResponseData) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SnapshotJobMutationResponseData) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SnapshotJobMutationResponseData) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetJob

`func (o *SnapshotJobMutationResponseData) GetJob() SnapshotJob`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *SnapshotJobMutationResponseData) GetJobOk() (*SnapshotJob, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *SnapshotJobMutationResponseData) SetJob(v SnapshotJob)`

SetJob sets Job field to given value.


### GetLimits

`func (o *SnapshotJobMutationResponseData) GetLimits() SnapshotJobLimits`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *SnapshotJobMutationResponseData) GetLimitsOk() (*SnapshotJobLimits, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *SnapshotJobMutationResponseData) SetLimits(v SnapshotJobLimits)`

SetLimits sets Limits field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


