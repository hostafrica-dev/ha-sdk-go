# DeleteSnapshotJobRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceId** | **string** | Service ID - must be sent as a string | 
**JobId** | **string** | Snapshot job ID to delete | 

## Methods

### NewDeleteSnapshotJobRequestContent

`func NewDeleteSnapshotJobRequestContent(serviceId string, jobId string, ) *DeleteSnapshotJobRequestContent`

NewDeleteSnapshotJobRequestContent instantiates a new DeleteSnapshotJobRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteSnapshotJobRequestContentWithDefaults

`func NewDeleteSnapshotJobRequestContentWithDefaults() *DeleteSnapshotJobRequestContent`

NewDeleteSnapshotJobRequestContentWithDefaults instantiates a new DeleteSnapshotJobRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceId

`func (o *DeleteSnapshotJobRequestContent) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *DeleteSnapshotJobRequestContent) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *DeleteSnapshotJobRequestContent) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetJobId

`func (o *DeleteSnapshotJobRequestContent) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *DeleteSnapshotJobRequestContent) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *DeleteSnapshotJobRequestContent) SetJobId(v string)`

SetJobId sets JobId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


