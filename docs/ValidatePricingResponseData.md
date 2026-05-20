# ValidatePricingResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | [**ValidatePricingCurrency**](ValidatePricingCurrency.md) |  | 
**Products** | [**[]ValidatePricingProductResult**](ValidatePricingProductResult.md) |  | 
**Summary** | [**ValidatePricingSummary**](ValidatePricingSummary.md) |  | 
**Errors** | **[]string** | List of strings | 

## Methods

### NewValidatePricingResponseData

`func NewValidatePricingResponseData(currency ValidatePricingCurrency, products []ValidatePricingProductResult, summary ValidatePricingSummary, errors []string, ) *ValidatePricingResponseData`

NewValidatePricingResponseData instantiates a new ValidatePricingResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidatePricingResponseDataWithDefaults

`func NewValidatePricingResponseDataWithDefaults() *ValidatePricingResponseData`

NewValidatePricingResponseDataWithDefaults instantiates a new ValidatePricingResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *ValidatePricingResponseData) GetCurrency() ValidatePricingCurrency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ValidatePricingResponseData) GetCurrencyOk() (*ValidatePricingCurrency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ValidatePricingResponseData) SetCurrency(v ValidatePricingCurrency)`

SetCurrency sets Currency field to given value.


### GetProducts

`func (o *ValidatePricingResponseData) GetProducts() []ValidatePricingProductResult`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *ValidatePricingResponseData) GetProductsOk() (*[]ValidatePricingProductResult, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *ValidatePricingResponseData) SetProducts(v []ValidatePricingProductResult)`

SetProducts sets Products field to given value.


### GetSummary

`func (o *ValidatePricingResponseData) GetSummary() ValidatePricingSummary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *ValidatePricingResponseData) GetSummaryOk() (*ValidatePricingSummary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *ValidatePricingResponseData) SetSummary(v ValidatePricingSummary)`

SetSummary sets Summary field to given value.


### GetErrors

`func (o *ValidatePricingResponseData) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ValidatePricingResponseData) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ValidatePricingResponseData) SetErrors(v []string)`

SetErrors sets Errors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


