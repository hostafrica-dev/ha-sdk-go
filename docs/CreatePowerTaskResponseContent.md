# CreatePowerTaskResponseContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**OperationStatus**](OperationStatus.md) |  | 
**Data** | [**PowerTaskCreateResponseData**](PowerTaskCreateResponseData.md) |  | 

## Methods

### NewCreatePowerTaskResponseContent

`func NewCreatePowerTaskResponseContent(status OperationStatus, data PowerTaskCreateResponseData, ) *CreatePowerTaskResponseContent`

NewCreatePowerTaskResponseContent instantiates a new CreatePowerTaskResponseContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePowerTaskResponseContentWithDefaults

`func NewCreatePowerTaskResponseContentWithDefaults() *CreatePowerTaskResponseContent`

NewCreatePowerTaskResponseContentWithDefaults instantiates a new CreatePowerTaskResponseContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *CreatePowerTaskResponseContent) GetStatus() OperationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreatePowerTaskResponseContent) GetStatusOk() (*OperationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreatePowerTaskResponseContent) SetStatus(v OperationStatus)`

SetStatus sets Status field to given value.


### GetData

`func (o *CreatePowerTaskResponseContent) GetData() PowerTaskCreateResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreatePowerTaskResponseContent) GetDataOk() (*PowerTaskCreateResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreatePowerTaskResponseContent) SetData(v PowerTaskCreateResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


