# RetryPaymentRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceId** | **string** | Service ID - must be sent as a string | 
**OrderId** | **int32** | The order identifier to retry payment for | 

## Methods

### NewRetryPaymentRequestContent

`func NewRetryPaymentRequestContent(serviceId string, orderId int32, ) *RetryPaymentRequestContent`

NewRetryPaymentRequestContent instantiates a new RetryPaymentRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRetryPaymentRequestContentWithDefaults

`func NewRetryPaymentRequestContentWithDefaults() *RetryPaymentRequestContent`

NewRetryPaymentRequestContentWithDefaults instantiates a new RetryPaymentRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceId

`func (o *RetryPaymentRequestContent) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *RetryPaymentRequestContent) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *RetryPaymentRequestContent) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetOrderId

`func (o *RetryPaymentRequestContent) GetOrderId() int32`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *RetryPaymentRequestContent) GetOrderIdOk() (*int32, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *RetryPaymentRequestContent) SetOrderId(v int32)`

SetOrderId sets OrderId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


