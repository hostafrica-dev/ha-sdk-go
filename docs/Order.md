# Order

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | **int32** | Unique order identifier | 
**OrderNumber** | **string** | Human-readable order reference number | 
**InvoiceId** | **int32** | Associated invoice identifier | 
**PlacedAt** | **string** | Timestamp when the order was placed (ISO 8601) | 
**Currency** | **string** | ISO currency code (e.g. ZAR) | 
**Total** | **string** | Total order amount as a decimal string | 
**BalanceDue** | **string** | Remaining balance due as a decimal string | 
**InvoiceStatus** | **string** | Invoice status (e.g. Paid, Unpaid) | 
**LastAttempt** | Pointer to [**OrderLastAttempt**](OrderLastAttempt.md) |  | [optional] 
**PaymentStatus** | **string** | Payment status (e.g. paid, pending, failed) | 

## Methods

### NewOrder

`func NewOrder(orderId int32, orderNumber string, invoiceId int32, placedAt string, currency string, total string, balanceDue string, invoiceStatus string, paymentStatus string, ) *Order`

NewOrder instantiates a new Order object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderWithDefaults

`func NewOrderWithDefaults() *Order`

NewOrderWithDefaults instantiates a new Order object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderId

`func (o *Order) GetOrderId() int32`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *Order) GetOrderIdOk() (*int32, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *Order) SetOrderId(v int32)`

SetOrderId sets OrderId field to given value.


### GetOrderNumber

`func (o *Order) GetOrderNumber() string`

GetOrderNumber returns the OrderNumber field if non-nil, zero value otherwise.

### GetOrderNumberOk

`func (o *Order) GetOrderNumberOk() (*string, bool)`

GetOrderNumberOk returns a tuple with the OrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderNumber

`func (o *Order) SetOrderNumber(v string)`

SetOrderNumber sets OrderNumber field to given value.


### GetInvoiceId

`func (o *Order) GetInvoiceId() int32`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *Order) GetInvoiceIdOk() (*int32, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *Order) SetInvoiceId(v int32)`

SetInvoiceId sets InvoiceId field to given value.


### GetPlacedAt

`func (o *Order) GetPlacedAt() string`

GetPlacedAt returns the PlacedAt field if non-nil, zero value otherwise.

### GetPlacedAtOk

`func (o *Order) GetPlacedAtOk() (*string, bool)`

GetPlacedAtOk returns a tuple with the PlacedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacedAt

`func (o *Order) SetPlacedAt(v string)`

SetPlacedAt sets PlacedAt field to given value.


### GetCurrency

`func (o *Order) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Order) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Order) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetTotal

`func (o *Order) GetTotal() string`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *Order) GetTotalOk() (*string, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *Order) SetTotal(v string)`

SetTotal sets Total field to given value.


### GetBalanceDue

`func (o *Order) GetBalanceDue() string`

GetBalanceDue returns the BalanceDue field if non-nil, zero value otherwise.

### GetBalanceDueOk

`func (o *Order) GetBalanceDueOk() (*string, bool)`

GetBalanceDueOk returns a tuple with the BalanceDue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceDue

`func (o *Order) SetBalanceDue(v string)`

SetBalanceDue sets BalanceDue field to given value.


### GetInvoiceStatus

`func (o *Order) GetInvoiceStatus() string`

GetInvoiceStatus returns the InvoiceStatus field if non-nil, zero value otherwise.

### GetInvoiceStatusOk

`func (o *Order) GetInvoiceStatusOk() (*string, bool)`

GetInvoiceStatusOk returns a tuple with the InvoiceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceStatus

`func (o *Order) SetInvoiceStatus(v string)`

SetInvoiceStatus sets InvoiceStatus field to given value.


### GetLastAttempt

`func (o *Order) GetLastAttempt() OrderLastAttempt`

GetLastAttempt returns the LastAttempt field if non-nil, zero value otherwise.

### GetLastAttemptOk

`func (o *Order) GetLastAttemptOk() (*OrderLastAttempt, bool)`

GetLastAttemptOk returns a tuple with the LastAttempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAttempt

`func (o *Order) SetLastAttempt(v OrderLastAttempt)`

SetLastAttempt sets LastAttempt field to given value.

### HasLastAttempt

`func (o *Order) HasLastAttempt() bool`

HasLastAttempt returns a boolean if a field has been set.

### GetPaymentStatus

`func (o *Order) GetPaymentStatus() string`

GetPaymentStatus returns the PaymentStatus field if non-nil, zero value otherwise.

### GetPaymentStatusOk

`func (o *Order) GetPaymentStatusOk() (*string, bool)`

GetPaymentStatusOk returns a tuple with the PaymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentStatus

`func (o *Order) SetPaymentStatus(v string)`

SetPaymentStatus sets PaymentStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


