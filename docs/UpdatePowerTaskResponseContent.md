# UpdatePowerTaskResponseContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**OperationStatus**](OperationStatus.md) |  | 
**Data** | [**PowerTaskCreateResponseData**](PowerTaskCreateResponseData.md) |  | 

## Methods

### NewUpdatePowerTaskResponseContent

`func NewUpdatePowerTaskResponseContent(status OperationStatus, data PowerTaskCreateResponseData, ) *UpdatePowerTaskResponseContent`

NewUpdatePowerTaskResponseContent instantiates a new UpdatePowerTaskResponseContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePowerTaskResponseContentWithDefaults

`func NewUpdatePowerTaskResponseContentWithDefaults() *UpdatePowerTaskResponseContent`

NewUpdatePowerTaskResponseContentWithDefaults instantiates a new UpdatePowerTaskResponseContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *UpdatePowerTaskResponseContent) GetStatus() OperationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdatePowerTaskResponseContent) GetStatusOk() (*OperationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdatePowerTaskResponseContent) SetStatus(v OperationStatus)`

SetStatus sets Status field to given value.


### GetData

`func (o *UpdatePowerTaskResponseContent) GetData() PowerTaskCreateResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UpdatePowerTaskResponseContent) GetDataOk() (*PowerTaskCreateResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UpdatePowerTaskResponseContent) SetData(v PowerTaskCreateResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


