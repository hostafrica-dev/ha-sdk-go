# DeleteDnssecRecordResponseContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**OperationStatus**](OperationStatus.md) |  | 
**Data** | [**DnssecMutationData**](DnssecMutationData.md) |  | 

## Methods

### NewDeleteDnssecRecordResponseContent

`func NewDeleteDnssecRecordResponseContent(status OperationStatus, data DnssecMutationData, ) *DeleteDnssecRecordResponseContent`

NewDeleteDnssecRecordResponseContent instantiates a new DeleteDnssecRecordResponseContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteDnssecRecordResponseContentWithDefaults

`func NewDeleteDnssecRecordResponseContentWithDefaults() *DeleteDnssecRecordResponseContent`

NewDeleteDnssecRecordResponseContentWithDefaults instantiates a new DeleteDnssecRecordResponseContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *DeleteDnssecRecordResponseContent) GetStatus() OperationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeleteDnssecRecordResponseContent) GetStatusOk() (*OperationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeleteDnssecRecordResponseContent) SetStatus(v OperationStatus)`

SetStatus sets Status field to given value.


### GetData

`func (o *DeleteDnssecRecordResponseContent) GetData() DnssecMutationData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DeleteDnssecRecordResponseContent) GetDataOk() (*DnssecMutationData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DeleteDnssecRecordResponseContent) SetData(v DnssecMutationData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


