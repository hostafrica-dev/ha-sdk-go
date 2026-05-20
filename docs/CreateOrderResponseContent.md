# CreateOrderResponseContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**OperationStatus**](OperationStatus.md) |  | 
**Data** | [**CreateOrderResponseData**](CreateOrderResponseData.md) |  | 

## Methods

### NewCreateOrderResponseContent

`func NewCreateOrderResponseContent(status OperationStatus, data CreateOrderResponseData, ) *CreateOrderResponseContent`

NewCreateOrderResponseContent instantiates a new CreateOrderResponseContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrderResponseContentWithDefaults

`func NewCreateOrderResponseContentWithDefaults() *CreateOrderResponseContent`

NewCreateOrderResponseContentWithDefaults instantiates a new CreateOrderResponseContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *CreateOrderResponseContent) GetStatus() OperationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateOrderResponseContent) GetStatusOk() (*OperationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateOrderResponseContent) SetStatus(v OperationStatus)`

SetStatus sets Status field to given value.


### GetData

`func (o *CreateOrderResponseContent) GetData() CreateOrderResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateOrderResponseContent) GetDataOk() (*CreateOrderResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateOrderResponseContent) SetData(v CreateOrderResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


