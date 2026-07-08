# SnapshotJobUpdateResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Status message indicating the result | 
**Job** | [**SnapshotJob**](SnapshotJob.md) |  | 

## Methods

### NewSnapshotJobUpdateResponseData

`func NewSnapshotJobUpdateResponseData(message string, job SnapshotJob, ) *SnapshotJobUpdateResponseData`

NewSnapshotJobUpdateResponseData instantiates a new SnapshotJobUpdateResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotJobUpdateResponseDataWithDefaults

`func NewSnapshotJobUpdateResponseDataWithDefaults() *SnapshotJobUpdateResponseData`

NewSnapshotJobUpdateResponseDataWithDefaults instantiates a new SnapshotJobUpdateResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *SnapshotJobUpdateResponseData) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SnapshotJobUpdateResponseData) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SnapshotJobUpdateResponseData) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetJob

`func (o *SnapshotJobUpdateResponseData) GetJob() SnapshotJob`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *SnapshotJobUpdateResponseData) GetJobOk() (*SnapshotJob, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *SnapshotJobUpdateResponseData) SetJob(v SnapshotJob)`

SetJob sets Job field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


