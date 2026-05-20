# RetryPaymentResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | **int32** | Order identifier | 
**InvoiceId** | **int32** | Invoice identifier for this order | 
**Total** | [**RetryPaymentTotal**](RetryPaymentTotal.md) |  | 
**PaymentStatus** | [**PaymentStatus**](PaymentStatus.md) |  | 

## Methods

### NewRetryPaymentResponseData

`func NewRetryPaymentResponseData(orderId int32, invoiceId int32, total RetryPaymentTotal, paymentStatus PaymentStatus, ) *RetryPaymentResponseData`

NewRetryPaymentResponseData instantiates a new RetryPaymentResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRetryPaymentResponseDataWithDefaults

`func NewRetryPaymentResponseDataWithDefaults() *RetryPaymentResponseData`

NewRetryPaymentResponseDataWithDefaults instantiates a new RetryPaymentResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderId

`func (o *RetryPaymentResponseData) GetOrderId() int32`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *RetryPaymentResponseData) GetOrderIdOk() (*int32, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *RetryPaymentResponseData) SetOrderId(v int32)`

SetOrderId sets OrderId field to given value.


### GetInvoiceId

`func (o *RetryPaymentResponseData) GetInvoiceId() int32`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *RetryPaymentResponseData) GetInvoiceIdOk() (*int32, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *RetryPaymentResponseData) SetInvoiceId(v int32)`

SetInvoiceId sets InvoiceId field to given value.


### GetTotal

`func (o *RetryPaymentResponseData) GetTotal() RetryPaymentTotal`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *RetryPaymentResponseData) GetTotalOk() (*RetryPaymentTotal, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *RetryPaymentResponseData) SetTotal(v RetryPaymentTotal)`

SetTotal sets Total field to given value.


### GetPaymentStatus

`func (o *RetryPaymentResponseData) GetPaymentStatus() PaymentStatus`

GetPaymentStatus returns the PaymentStatus field if non-nil, zero value otherwise.

### GetPaymentStatusOk

`func (o *RetryPaymentResponseData) GetPaymentStatusOk() (*PaymentStatus, bool)`

GetPaymentStatusOk returns a tuple with the PaymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentStatus

`func (o *RetryPaymentResponseData) SetPaymentStatus(v PaymentStatus)`

SetPaymentStatus sets PaymentStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


