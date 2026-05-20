# SnapshotCreateResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Status message indicating the result | 
**TaskId** | Pointer to **int32** | Task ID for the snapshot creation | [optional] 

## Methods

### NewSnapshotCreateResponseData

`func NewSnapshotCreateResponseData(message string, ) *SnapshotCreateResponseData`

NewSnapshotCreateResponseData instantiates a new SnapshotCreateResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotCreateResponseDataWithDefaults

`func NewSnapshotCreateResponseDataWithDefaults() *SnapshotCreateResponseData`

NewSnapshotCreateResponseDataWithDefaults instantiates a new SnapshotCreateResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *SnapshotCreateResponseData) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SnapshotCreateResponseData) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SnapshotCreateResponseData) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetTaskId

`func (o *SnapshotCreateResponseData) GetTaskId() int32`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *SnapshotCreateResponseData) GetTaskIdOk() (*int32, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *SnapshotCreateResponseData) SetTaskId(v int32)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *SnapshotCreateResponseData) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


