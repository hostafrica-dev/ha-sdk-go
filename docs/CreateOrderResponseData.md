# CreateOrderResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | **int32** | Order identifier | 
**InvoiceId** | **int32** | Invoice identifier for this order | 
**OrderNumber** | **string** | Human-readable order reference number | 
**Total** | [**CreateOrderTotal**](CreateOrderTotal.md) |  | 
**PromoApplied** | Pointer to **string** | Promotional code that was applied | [optional] 
**PaymentMethod** | **string** | Payment method used (e.g. stripe) | 
**PaymentStatus** | [**PaymentStatus**](PaymentStatus.md) |  | 
**Items** | [**CreateOrderItems**](CreateOrderItems.md) |  | 
**PaymentError** | Pointer to [**PaymentError**](PaymentError.md) |  | [optional] 
**Warnings** | Pointer to [**[]OrderWarning**](OrderWarning.md) | Warnings present when non-critical issues occurred (e.g. autorenew deferred) | [optional] 

## Methods

### NewCreateOrderResponseData

`func NewCreateOrderResponseData(orderId int32, invoiceId int32, orderNumber string, total CreateOrderTotal, paymentMethod string, paymentStatus PaymentStatus, items CreateOrderItems, ) *CreateOrderResponseData`

NewCreateOrderResponseData instantiates a new CreateOrderResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrderResponseDataWithDefaults

`func NewCreateOrderResponseDataWithDefaults() *CreateOrderResponseData`

NewCreateOrderResponseDataWithDefaults instantiates a new CreateOrderResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderId

`func (o *CreateOrderResponseData) GetOrderId() int32`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *CreateOrderResponseData) GetOrderIdOk() (*int32, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *CreateOrderResponseData) SetOrderId(v int32)`

SetOrderId sets OrderId field to given value.


### GetInvoiceId

`func (o *CreateOrderResponseData) GetInvoiceId() int32`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *CreateOrderResponseData) GetInvoiceIdOk() (*int32, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *CreateOrderResponseData) SetInvoiceId(v int32)`

SetInvoiceId sets InvoiceId field to given value.


### GetOrderNumber

`func (o *CreateOrderResponseData) GetOrderNumber() string`

GetOrderNumber returns the OrderNumber field if non-nil, zero value otherwise.

### GetOrderNumberOk

`func (o *CreateOrderResponseData) GetOrderNumberOk() (*string, bool)`

GetOrderNumberOk returns a tuple with the OrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderNumber

`func (o *CreateOrderResponseData) SetOrderNumber(v string)`

SetOrderNumber sets OrderNumber field to given value.


### GetTotal

`func (o *CreateOrderResponseData) GetTotal() CreateOrderTotal`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *CreateOrderResponseData) GetTotalOk() (*CreateOrderTotal, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *CreateOrderResponseData) SetTotal(v CreateOrderTotal)`

SetTotal sets Total field to given value.


### GetPromoApplied

`func (o *CreateOrderResponseData) GetPromoApplied() string`

GetPromoApplied returns the PromoApplied field if non-nil, zero value otherwise.

### GetPromoAppliedOk

`func (o *CreateOrderResponseData) GetPromoAppliedOk() (*string, bool)`

GetPromoAppliedOk returns a tuple with the PromoApplied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromoApplied

`func (o *CreateOrderResponseData) SetPromoApplied(v string)`

SetPromoApplied sets PromoApplied field to given value.

### HasPromoApplied

`func (o *CreateOrderResponseData) HasPromoApplied() bool`

HasPromoApplied returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *CreateOrderResponseData) GetPaymentMethod() string`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *CreateOrderResponseData) GetPaymentMethodOk() (*string, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *CreateOrderResponseData) SetPaymentMethod(v string)`

SetPaymentMethod sets PaymentMethod field to given value.


### GetPaymentStatus

`func (o *CreateOrderResponseData) GetPaymentStatus() PaymentStatus`

GetPaymentStatus returns the PaymentStatus field if non-nil, zero value otherwise.

### GetPaymentStatusOk

`func (o *CreateOrderResponseData) GetPaymentStatusOk() (*PaymentStatus, bool)`

GetPaymentStatusOk returns a tuple with the PaymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentStatus

`func (o *CreateOrderResponseData) SetPaymentStatus(v PaymentStatus)`

SetPaymentStatus sets PaymentStatus field to given value.


### GetItems

`func (o *CreateOrderResponseData) GetItems() CreateOrderItems`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CreateOrderResponseData) GetItemsOk() (*CreateOrderItems, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CreateOrderResponseData) SetItems(v CreateOrderItems)`

SetItems sets Items field to given value.


### GetPaymentError

`func (o *CreateOrderResponseData) GetPaymentError() PaymentError`

GetPaymentError returns the PaymentError field if non-nil, zero value otherwise.

### GetPaymentErrorOk

`func (o *CreateOrderResponseData) GetPaymentErrorOk() (*PaymentError, bool)`

GetPaymentErrorOk returns a tuple with the PaymentError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentError

`func (o *CreateOrderResponseData) SetPaymentError(v PaymentError)`

SetPaymentError sets PaymentError field to given value.

### HasPaymentError

`func (o *CreateOrderResponseData) HasPaymentError() bool`

HasPaymentError returns a boolean if a field has been set.

### GetWarnings

`func (o *CreateOrderResponseData) GetWarnings() []OrderWarning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *CreateOrderResponseData) GetWarningsOk() (*[]OrderWarning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *CreateOrderResponseData) SetWarnings(v []OrderWarning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *CreateOrderResponseData) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


