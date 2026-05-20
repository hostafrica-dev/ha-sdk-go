# DeletePowerTaskResponseContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**OperationStatus**](OperationStatus.md) |  | 
**Data** | [**VpsSimpleActionResponseData**](VpsSimpleActionResponseData.md) |  | 

## Methods

### NewDeletePowerTaskResponseContent

`func NewDeletePowerTaskResponseContent(status OperationStatus, data VpsSimpleActionResponseData, ) *DeletePowerTaskResponseContent`

NewDeletePowerTaskResponseContent instantiates a new DeletePowerTaskResponseContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeletePowerTaskResponseContentWithDefaults

`func NewDeletePowerTaskResponseContentWithDefaults() *DeletePowerTaskResponseContent`

NewDeletePowerTaskResponseContentWithDefaults instantiates a new DeletePowerTaskResponseContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *DeletePowerTaskResponseContent) GetStatus() OperationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeletePowerTaskResponseContent) GetStatusOk() (*OperationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeletePowerTaskResponseContent) SetStatus(v OperationStatus)`

SetStatus sets Status field to given value.


### GetData

`func (o *DeletePowerTaskResponseContent) GetData() VpsSimpleActionResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DeletePowerTaskResponseContent) GetDataOk() (*VpsSimpleActionResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DeletePowerTaskResponseContent) SetData(v VpsSimpleActionResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


