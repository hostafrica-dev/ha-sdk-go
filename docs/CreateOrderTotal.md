# CreateOrderTotal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **string** | Total amount charged (decimal string, e.g. \&quot;2333.33\&quot;) | 
**Currency** | **string** | Currency code (e.g. USD) | 
**Prefix** | **string** | Currency prefix symbol (e.g. $) | 

## Methods

### NewCreateOrderTotal

`func NewCreateOrderTotal(amount string, currency string, prefix string, ) *CreateOrderTotal`

NewCreateOrderTotal instantiates a new CreateOrderTotal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrderTotalWithDefaults

`func NewCreateOrderTotalWithDefaults() *CreateOrderTotal`

NewCreateOrderTotalWithDefaults instantiates a new CreateOrderTotal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *CreateOrderTotal) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CreateOrderTotal) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CreateOrderTotal) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetCurrency

`func (o *CreateOrderTotal) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CreateOrderTotal) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CreateOrderTotal) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetPrefix

`func (o *CreateOrderTotal) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *CreateOrderTotal) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *CreateOrderTotal) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


