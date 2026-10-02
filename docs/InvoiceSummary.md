# InvoiceSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvoiceId** | **string** | Unique invoice identifier - must be sent as a string | 
**InvoiceNumber** | **string** | Human-readable invoice number when assigned | 
**Date** | **string** | Invoice issue date (ISO 8601) | 
**DueDate** | **string** | Invoice due date (ISO 8601) | 
**Status** | **string** | Invoice status (e.g. Paid, Unpaid, Cancelled, Refunded) | 
**Tax** | **string** | Primary tax amount as a decimal string | 
**Tax2** | Pointer to **string** | Secondary tax amount as a decimal string. Omitted when a second tax is not configured. | [optional] 
**TaxRate** | **string** | Primary tax rate as a decimal string | 
**TaxRate2** | Pointer to **string** | Secondary tax rate as a decimal string. Omitted when a second tax is not configured. | [optional] 
**Total** | **string** | Invoice total including tax as a decimal string | 
**Subtotal** | **string** | Invoice subtotal excluding tax as a decimal string | 
**GrandTotal** | **string** | Grand total as a decimal string | 

## Methods

### NewInvoiceSummary

`func NewInvoiceSummary(invoiceId string, invoiceNumber string, date string, dueDate string, status string, tax string, taxRate string, total string, subtotal string, grandTotal string, ) *InvoiceSummary`

NewInvoiceSummary instantiates a new InvoiceSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceSummaryWithDefaults

`func NewInvoiceSummaryWithDefaults() *InvoiceSummary`

NewInvoiceSummaryWithDefaults instantiates a new InvoiceSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvoiceId

`func (o *InvoiceSummary) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *InvoiceSummary) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *InvoiceSummary) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.


### GetInvoiceNumber

`func (o *InvoiceSummary) GetInvoiceNumber() string`

GetInvoiceNumber returns the InvoiceNumber field if non-nil, zero value otherwise.

### GetInvoiceNumberOk

`func (o *InvoiceSummary) GetInvoiceNumberOk() (*string, bool)`

GetInvoiceNumberOk returns a tuple with the InvoiceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceNumber

`func (o *InvoiceSummary) SetInvoiceNumber(v string)`

SetInvoiceNumber sets InvoiceNumber field to given value.


### GetDate

`func (o *InvoiceSummary) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *InvoiceSummary) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *InvoiceSummary) SetDate(v string)`

SetDate sets Date field to given value.


### GetDueDate

`func (o *InvoiceSummary) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *InvoiceSummary) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *InvoiceSummary) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.


### GetStatus

`func (o *InvoiceSummary) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InvoiceSummary) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InvoiceSummary) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTax

`func (o *InvoiceSummary) GetTax() string`

GetTax returns the Tax field if non-nil, zero value otherwise.

### GetTaxOk

`func (o *InvoiceSummary) GetTaxOk() (*string, bool)`

GetTaxOk returns a tuple with the Tax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTax

`func (o *InvoiceSummary) SetTax(v string)`

SetTax sets Tax field to given value.


### GetTax2

`func (o *InvoiceSummary) GetTax2() string`

GetTax2 returns the Tax2 field if non-nil, zero value otherwise.

### GetTax2Ok

`func (o *InvoiceSummary) GetTax2Ok() (*string, bool)`

GetTax2Ok returns a tuple with the Tax2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTax2

`func (o *InvoiceSummary) SetTax2(v string)`

SetTax2 sets Tax2 field to given value.

### HasTax2

`func (o *InvoiceSummary) HasTax2() bool`

HasTax2 returns a boolean if a field has been set.

### GetTaxRate

`func (o *InvoiceSummary) GetTaxRate() string`

GetTaxRate returns the TaxRate field if non-nil, zero value otherwise.

### GetTaxRateOk

`func (o *InvoiceSummary) GetTaxRateOk() (*string, bool)`

GetTaxRateOk returns a tuple with the TaxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRate

`func (o *InvoiceSummary) SetTaxRate(v string)`

SetTaxRate sets TaxRate field to given value.


### GetTaxRate2

`func (o *InvoiceSummary) GetTaxRate2() string`

GetTaxRate2 returns the TaxRate2 field if non-nil, zero value otherwise.

### GetTaxRate2Ok

`func (o *InvoiceSummary) GetTaxRate2Ok() (*string, bool)`

GetTaxRate2Ok returns a tuple with the TaxRate2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRate2

`func (o *InvoiceSummary) SetTaxRate2(v string)`

SetTaxRate2 sets TaxRate2 field to given value.

### HasTaxRate2

`func (o *InvoiceSummary) HasTaxRate2() bool`

HasTaxRate2 returns a boolean if a field has been set.

### GetTotal

`func (o *InvoiceSummary) GetTotal() string`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *InvoiceSummary) GetTotalOk() (*string, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *InvoiceSummary) SetTotal(v string)`

SetTotal sets Total field to given value.


### GetSubtotal

`func (o *InvoiceSummary) GetSubtotal() string`

GetSubtotal returns the Subtotal field if non-nil, zero value otherwise.

### GetSubtotalOk

`func (o *InvoiceSummary) GetSubtotalOk() (*string, bool)`

GetSubtotalOk returns a tuple with the Subtotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtotal

`func (o *InvoiceSummary) SetSubtotal(v string)`

SetSubtotal sets Subtotal field to given value.


### GetGrandTotal

`func (o *InvoiceSummary) GetGrandTotal() string`

GetGrandTotal returns the GrandTotal field if non-nil, zero value otherwise.

### GetGrandTotalOk

`func (o *InvoiceSummary) GetGrandTotalOk() (*string, bool)`

GetGrandTotalOk returns a tuple with the GrandTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrandTotal

`func (o *InvoiceSummary) SetGrandTotal(v string)`

SetGrandTotal sets GrandTotal field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


