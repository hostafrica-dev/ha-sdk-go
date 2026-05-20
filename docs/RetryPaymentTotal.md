# RetryPaymentTotal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **string** | Total amount charged | 
**Currency** | **string** | Currency code (e.g. USD) | 
**Prefix** | **string** | Currency prefix symbol (e.g. US$) | 

## Methods

### NewRetryPaymentTotal

`func NewRetryPaymentTotal(amount string, currency string, prefix string, ) *RetryPaymentTotal`

NewRetryPaymentTotal instantiates a new RetryPaymentTotal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRetryPaymentTotalWithDefaults

`func NewRetryPaymentTotalWithDefaults() *RetryPaymentTotal`

NewRetryPaymentTotalWithDefaults instantiates a new RetryPaymentTotal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *RetryPaymentTotal) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *RetryPaymentTotal) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *RetryPaymentTotal) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetCurrency

`func (o *RetryPaymentTotal) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *RetryPaymentTotal) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *RetryPaymentTotal) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetPrefix

`func (o *RetryPaymentTotal) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *RetryPaymentTotal) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *RetryPaymentTotal) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


