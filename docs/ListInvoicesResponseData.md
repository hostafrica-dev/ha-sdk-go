# ListInvoicesResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Invoices** | [**[]InvoiceSummary**](InvoiceSummary.md) | List of invoices for the authenticated user | 
**TotalCount** | **int32** | Total number of invoices returned | 

## Methods

### NewListInvoicesResponseData

`func NewListInvoicesResponseData(invoices []InvoiceSummary, totalCount int32, ) *ListInvoicesResponseData`

NewListInvoicesResponseData instantiates a new ListInvoicesResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListInvoicesResponseDataWithDefaults

`func NewListInvoicesResponseDataWithDefaults() *ListInvoicesResponseData`

NewListInvoicesResponseDataWithDefaults instantiates a new ListInvoicesResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvoices

`func (o *ListInvoicesResponseData) GetInvoices() []InvoiceSummary`

GetInvoices returns the Invoices field if non-nil, zero value otherwise.

### GetInvoicesOk

`func (o *ListInvoicesResponseData) GetInvoicesOk() (*[]InvoiceSummary, bool)`

GetInvoicesOk returns a tuple with the Invoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoices

`func (o *ListInvoicesResponseData) SetInvoices(v []InvoiceSummary)`

SetInvoices sets Invoices field to given value.


### GetTotalCount

`func (o *ListInvoicesResponseData) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ListInvoicesResponseData) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ListInvoicesResponseData) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


