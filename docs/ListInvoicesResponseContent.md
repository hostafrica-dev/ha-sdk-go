# ListInvoicesResponseContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**OperationStatus**](OperationStatus.md) |  | 
**Data** | [**ListInvoicesResponseData**](ListInvoicesResponseData.md) |  | 

## Methods

### NewListInvoicesResponseContent

`func NewListInvoicesResponseContent(status OperationStatus, data ListInvoicesResponseData, ) *ListInvoicesResponseContent`

NewListInvoicesResponseContent instantiates a new ListInvoicesResponseContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListInvoicesResponseContentWithDefaults

`func NewListInvoicesResponseContentWithDefaults() *ListInvoicesResponseContent`

NewListInvoicesResponseContentWithDefaults instantiates a new ListInvoicesResponseContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ListInvoicesResponseContent) GetStatus() OperationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListInvoicesResponseContent) GetStatusOk() (*OperationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListInvoicesResponseContent) SetStatus(v OperationStatus)`

SetStatus sets Status field to given value.


### GetData

`func (o *ListInvoicesResponseContent) GetData() ListInvoicesResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListInvoicesResponseContent) GetDataOk() (*ListInvoicesResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListInvoicesResponseContent) SetData(v ListInvoicesResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


