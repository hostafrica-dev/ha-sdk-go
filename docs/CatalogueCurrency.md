# CatalogueCurrency

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | ISO currency code (e.g. KES, ZAR) | 
**Prefix** | **string** | Currency prefix symbol (e.g. KSh) | 
**Suffix** | **string** | Currency suffix symbol (empty string if none) | 

## Methods

### NewCatalogueCurrency

`func NewCatalogueCurrency(code string, prefix string, suffix string, ) *CatalogueCurrency`

NewCatalogueCurrency instantiates a new CatalogueCurrency object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogueCurrencyWithDefaults

`func NewCatalogueCurrencyWithDefaults() *CatalogueCurrency`

NewCatalogueCurrencyWithDefaults instantiates a new CatalogueCurrency object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *CatalogueCurrency) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CatalogueCurrency) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CatalogueCurrency) SetCode(v string)`

SetCode sets Code field to given value.


### GetPrefix

`func (o *CatalogueCurrency) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *CatalogueCurrency) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *CatalogueCurrency) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.


### GetSuffix

`func (o *CatalogueCurrency) GetSuffix() string`

GetSuffix returns the Suffix field if non-nil, zero value otherwise.

### GetSuffixOk

`func (o *CatalogueCurrency) GetSuffixOk() (*string, bool)`

GetSuffixOk returns a tuple with the Suffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuffix

`func (o *CatalogueCurrency) SetSuffix(v string)`

SetSuffix sets Suffix field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


