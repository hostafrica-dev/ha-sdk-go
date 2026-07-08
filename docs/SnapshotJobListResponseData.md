# SnapshotJobListResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Status message indicating the result | 
**Jobs** | [**[]SnapshotJob**](SnapshotJob.md) | List of snapshot jobs | 
**Limits** | [**SnapshotJobLimits**](SnapshotJobLimits.md) |  | 
**AllowedPeriods** | Pointer to **[]string** | Allowed schedule periods for snapshot jobs | [optional] 

## Methods

### NewSnapshotJobListResponseData

`func NewSnapshotJobListResponseData(message string, jobs []SnapshotJob, limits SnapshotJobLimits, ) *SnapshotJobListResponseData`

NewSnapshotJobListResponseData instantiates a new SnapshotJobListResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotJobListResponseDataWithDefaults

`func NewSnapshotJobListResponseDataWithDefaults() *SnapshotJobListResponseData`

NewSnapshotJobListResponseDataWithDefaults instantiates a new SnapshotJobListResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *SnapshotJobListResponseData) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SnapshotJobListResponseData) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SnapshotJobListResponseData) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetJobs

`func (o *SnapshotJobListResponseData) GetJobs() []SnapshotJob`

GetJobs returns the Jobs field if non-nil, zero value otherwise.

### GetJobsOk

`func (o *SnapshotJobListResponseData) GetJobsOk() (*[]SnapshotJob, bool)`

GetJobsOk returns a tuple with the Jobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobs

`func (o *SnapshotJobListResponseData) SetJobs(v []SnapshotJob)`

SetJobs sets Jobs field to given value.


### GetLimits

`func (o *SnapshotJobListResponseData) GetLimits() SnapshotJobLimits`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *SnapshotJobListResponseData) GetLimitsOk() (*SnapshotJobLimits, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *SnapshotJobListResponseData) SetLimits(v SnapshotJobLimits)`

SetLimits sets Limits field to given value.


### GetAllowedPeriods

`func (o *SnapshotJobListResponseData) GetAllowedPeriods() []string`

GetAllowedPeriods returns the AllowedPeriods field if non-nil, zero value otherwise.

### GetAllowedPeriodsOk

`func (o *SnapshotJobListResponseData) GetAllowedPeriodsOk() (*[]string, bool)`

GetAllowedPeriodsOk returns a tuple with the AllowedPeriods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedPeriods

`func (o *SnapshotJobListResponseData) SetAllowedPeriods(v []string)`

SetAllowedPeriods sets AllowedPeriods field to given value.

### HasAllowedPeriods

`func (o *SnapshotJobListResponseData) HasAllowedPeriods() bool`

HasAllowedPeriods returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


