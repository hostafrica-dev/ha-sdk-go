# ListOrdersResponseContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**OperationStatus**](OperationStatus.md) |  | 
**Data** | [**ListOrdersResponseData**](ListOrdersResponseData.md) |  | 

## Methods

### NewListOrdersResponseContent

`func NewListOrdersResponseContent(status OperationStatus, data ListOrdersResponseData, ) *ListOrdersResponseContent`

NewListOrdersResponseContent instantiates a new ListOrdersResponseContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListOrdersResponseContentWithDefaults

`func NewListOrdersResponseContentWithDefaults() *ListOrdersResponseContent`

NewListOrdersResponseContentWithDefaults instantiates a new ListOrdersResponseContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ListOrdersResponseContent) GetStatus() OperationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListOrdersResponseContent) GetStatusOk() (*OperationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListOrdersResponseContent) SetStatus(v OperationStatus)`

SetStatus sets Status field to given value.


### GetData

`func (o *ListOrdersResponseContent) GetData() ListOrdersResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListOrdersResponseContent) GetDataOk() (*ListOrdersResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListOrdersResponseContent) SetData(v ListOrdersResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


