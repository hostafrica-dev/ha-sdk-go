# InvoiceDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvoiceId** | **string** | Unique invoice identifier - must be sent as a string | 
**InvoiceNumber** | **string** | Human-readable invoice number when assigned | 
**IssuedTo** | [**InvoiceIssuedTo**](InvoiceIssuedTo.md) |  | 
**Date** | **string** | Invoice issue date (ISO 8601) | 
**DueDate** | **string** | Invoice due date (ISO 8601) | 
**DatePaid** | Pointer to **string** | Date the invoice was paid (ISO 8601), if paid | [optional] 
**Subtotal** | **string** | Invoice subtotal excluding tax as a decimal string | 
**Credit** | **string** | Credit applied as a decimal string | 
**Tax** | **string** | Primary tax amount as a decimal string | 
**Tax2** | Pointer to **string** | Secondary tax amount as a decimal string. Omitted when a second tax is not configured. | [optional] 
**Total** | **string** | Invoice total including tax as a decimal string | 
**TaxRate** | **string** | Primary tax rate as a decimal string | 
**TaxRate2** | Pointer to **string** | Secondary tax rate as a decimal string. Omitted when a second tax is not configured. | [optional] 
**Status** | **string** | Invoice status (e.g. Paid, Unpaid, Cancelled, Refunded) | 
**PaymentMethod** | **string** | Payment method identifier | 
**TotalDue** | **string** | Amount still due as a decimal string | 
**Notes** | **string** | Invoice notes | 
**Items** | [**[]InvoiceItem**](InvoiceItem.md) | Invoice line items | 
**Transactions** | [**[]InvoiceTransaction**](InvoiceTransaction.md) | Payment and refund transactions on the invoice | 

## Methods

### NewInvoiceDetails

`func NewInvoiceDetails(invoiceId string, invoiceNumber string, issuedTo InvoiceIssuedTo, date string, dueDate string, subtotal string, credit string, tax string, total string, taxRate string, status string, paymentMethod string, totalDue string, notes string, items []InvoiceItem, transactions []InvoiceTransaction, ) *InvoiceDetails`

NewInvoiceDetails instantiates a new InvoiceDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceDetailsWithDefaults

`func NewInvoiceDetailsWithDefaults() *InvoiceDetails`

NewInvoiceDetailsWithDefaults instantiates a new InvoiceDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvoiceId

`func (o *InvoiceDetails) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *InvoiceDetails) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *InvoiceDetails) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.


### GetInvoiceNumber

`func (o *InvoiceDetails) GetInvoiceNumber() string`

GetInvoiceNumber returns the InvoiceNumber field if non-nil, zero value otherwise.

### GetInvoiceNumberOk

`func (o *InvoiceDetails) GetInvoiceNumberOk() (*string, bool)`

GetInvoiceNumberOk returns a tuple with the InvoiceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceNumber

`func (o *InvoiceDetails) SetInvoiceNumber(v string)`

SetInvoiceNumber sets InvoiceNumber field to given value.


### GetIssuedTo

`func (o *InvoiceDetails) GetIssuedTo() InvoiceIssuedTo`

GetIssuedTo returns the IssuedTo field if non-nil, zero value otherwise.

### GetIssuedToOk

`func (o *InvoiceDetails) GetIssuedToOk() (*InvoiceIssuedTo, bool)`

GetIssuedToOk returns a tuple with the IssuedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedTo

`func (o *InvoiceDetails) SetIssuedTo(v InvoiceIssuedTo)`

SetIssuedTo sets IssuedTo field to given value.


### GetDate

`func (o *InvoiceDetails) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *InvoiceDetails) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *InvoiceDetails) SetDate(v string)`

SetDate sets Date field to given value.


### GetDueDate

`func (o *InvoiceDetails) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *InvoiceDetails) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *InvoiceDetails) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.


### GetDatePaid

`func (o *InvoiceDetails) GetDatePaid() string`

GetDatePaid returns the DatePaid field if non-nil, zero value otherwise.

### GetDatePaidOk

`func (o *InvoiceDetails) GetDatePaidOk() (*string, bool)`

GetDatePaidOk returns a tuple with the DatePaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatePaid

`func (o *InvoiceDetails) SetDatePaid(v string)`

SetDatePaid sets DatePaid field to given value.

### HasDatePaid

`func (o *InvoiceDetails) HasDatePaid() bool`

HasDatePaid returns a boolean if a field has been set.

### GetSubtotal

`func (o *InvoiceDetails) GetSubtotal() string`

GetSubtotal returns the Subtotal field if non-nil, zero value otherwise.

### GetSubtotalOk

`func (o *InvoiceDetails) GetSubtotalOk() (*string, bool)`

GetSubtotalOk returns a tuple with the Subtotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtotal

`func (o *InvoiceDetails) SetSubtotal(v string)`

SetSubtotal sets Subtotal field to given value.


### GetCredit

`func (o *InvoiceDetails) GetCredit() string`

GetCredit returns the Credit field if non-nil, zero value otherwise.

### GetCreditOk

`func (o *InvoiceDetails) GetCreditOk() (*string, bool)`

GetCreditOk returns a tuple with the Credit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredit

`func (o *InvoiceDetails) SetCredit(v string)`

SetCredit sets Credit field to given value.


### GetTax

`func (o *InvoiceDetails) GetTax() string`

GetTax returns the Tax field if non-nil, zero value otherwise.

### GetTaxOk

`func (o *InvoiceDetails) GetTaxOk() (*string, bool)`

GetTaxOk returns a tuple with the Tax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTax

`func (o *InvoiceDetails) SetTax(v string)`

SetTax sets Tax field to given value.


### GetTax2

`func (o *InvoiceDetails) GetTax2() string`

GetTax2 returns the Tax2 field if non-nil, zero value otherwise.

### GetTax2Ok

`func (o *InvoiceDetails) GetTax2Ok() (*string, bool)`

GetTax2Ok returns a tuple with the Tax2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTax2

`func (o *InvoiceDetails) SetTax2(v string)`

SetTax2 sets Tax2 field to given value.

### HasTax2

`func (o *InvoiceDetails) HasTax2() bool`

HasTax2 returns a boolean if a field has been set.

### GetTotal

`func (o *InvoiceDetails) GetTotal() string`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *InvoiceDetails) GetTotalOk() (*string, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *InvoiceDetails) SetTotal(v string)`

SetTotal sets Total field to given value.


### GetTaxRate

`func (o *InvoiceDetails) GetTaxRate() string`

GetTaxRate returns the TaxRate field if non-nil, zero value otherwise.

### GetTaxRateOk

`func (o *InvoiceDetails) GetTaxRateOk() (*string, bool)`

GetTaxRateOk returns a tuple with the TaxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRate

`func (o *InvoiceDetails) SetTaxRate(v string)`

SetTaxRate sets TaxRate field to given value.


### GetTaxRate2

`func (o *InvoiceDetails) GetTaxRate2() string`

GetTaxRate2 returns the TaxRate2 field if non-nil, zero value otherwise.

### GetTaxRate2Ok

`func (o *InvoiceDetails) GetTaxRate2Ok() (*string, bool)`

GetTaxRate2Ok returns a tuple with the TaxRate2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRate2

`func (o *InvoiceDetails) SetTaxRate2(v string)`

SetTaxRate2 sets TaxRate2 field to given value.

### HasTaxRate2

`func (o *InvoiceDetails) HasTaxRate2() bool`

HasTaxRate2 returns a boolean if a field has been set.

### GetStatus

`func (o *InvoiceDetails) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InvoiceDetails) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InvoiceDetails) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetPaymentMethod

`func (o *InvoiceDetails) GetPaymentMethod() string`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *InvoiceDetails) GetPaymentMethodOk() (*string, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *InvoiceDetails) SetPaymentMethod(v string)`

SetPaymentMethod sets PaymentMethod field to given value.


### GetTotalDue

`func (o *InvoiceDetails) GetTotalDue() string`

GetTotalDue returns the TotalDue field if non-nil, zero value otherwise.

### GetTotalDueOk

`func (o *InvoiceDetails) GetTotalDueOk() (*string, bool)`

GetTotalDueOk returns a tuple with the TotalDue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDue

`func (o *InvoiceDetails) SetTotalDue(v string)`

SetTotalDue sets TotalDue field to given value.


### GetNotes

`func (o *InvoiceDetails) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *InvoiceDetails) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *InvoiceDetails) SetNotes(v string)`

SetNotes sets Notes field to given value.


### GetItems

`func (o *InvoiceDetails) GetItems() []InvoiceItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *InvoiceDetails) GetItemsOk() (*[]InvoiceItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *InvoiceDetails) SetItems(v []InvoiceItem)`

SetItems sets Items field to given value.


### GetTransactions

`func (o *InvoiceDetails) GetTransactions() []InvoiceTransaction`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *InvoiceDetails) GetTransactionsOk() (*[]InvoiceTransaction, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *InvoiceDetails) SetTransactions(v []InvoiceTransaction)`

SetTransactions sets Transactions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


