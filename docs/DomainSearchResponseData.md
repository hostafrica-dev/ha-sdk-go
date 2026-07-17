# DomainSearchResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domains** | [**[]DomainAvailabilityResult**](DomainAvailabilityResult.md) | Domain availability results | 
**Suggestions** | Pointer to **[]string** | Alternative domain suggestions (single-domain lookup only) | [optional] 
**Currency** | **int32** | Upstream currency identifier | 
**CurrencyCode** | **string** | ISO currency code for the quoted prices | 
**CurrencyNote** | **string** | Note that prices and register_url are valid only for currency_code | 
**Warnings** | Pointer to **[]string** | Non-fatal warnings the client should be aware of | [optional] 

## Methods

### NewDomainSearchResponseData

`func NewDomainSearchResponseData(domains []DomainAvailabilityResult, currency int32, currencyCode string, currencyNote string, ) *DomainSearchResponseData`

NewDomainSearchResponseData instantiates a new DomainSearchResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainSearchResponseDataWithDefaults

`func NewDomainSearchResponseDataWithDefaults() *DomainSearchResponseData`

NewDomainSearchResponseDataWithDefaults instantiates a new DomainSearchResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomains

`func (o *DomainSearchResponseData) GetDomains() []DomainAvailabilityResult`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *DomainSearchResponseData) GetDomainsOk() (*[]DomainAvailabilityResult, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *DomainSearchResponseData) SetDomains(v []DomainAvailabilityResult)`

SetDomains sets Domains field to given value.


### GetSuggestions

`func (o *DomainSearchResponseData) GetSuggestions() []string`

GetSuggestions returns the Suggestions field if non-nil, zero value otherwise.

### GetSuggestionsOk

`func (o *DomainSearchResponseData) GetSuggestionsOk() (*[]string, bool)`

GetSuggestionsOk returns a tuple with the Suggestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestions

`func (o *DomainSearchResponseData) SetSuggestions(v []string)`

SetSuggestions sets Suggestions field to given value.

### HasSuggestions

`func (o *DomainSearchResponseData) HasSuggestions() bool`

HasSuggestions returns a boolean if a field has been set.

### GetCurrency

`func (o *DomainSearchResponseData) GetCurrency() int32`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *DomainSearchResponseData) GetCurrencyOk() (*int32, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *DomainSearchResponseData) SetCurrency(v int32)`

SetCurrency sets Currency field to given value.


### GetCurrencyCode

`func (o *DomainSearchResponseData) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *DomainSearchResponseData) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *DomainSearchResponseData) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.


### GetCurrencyNote

`func (o *DomainSearchResponseData) GetCurrencyNote() string`

GetCurrencyNote returns the CurrencyNote field if non-nil, zero value otherwise.

### GetCurrencyNoteOk

`func (o *DomainSearchResponseData) GetCurrencyNoteOk() (*string, bool)`

GetCurrencyNoteOk returns a tuple with the CurrencyNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyNote

`func (o *DomainSearchResponseData) SetCurrencyNote(v string)`

SetCurrencyNote sets CurrencyNote field to given value.


### GetWarnings

`func (o *DomainSearchResponseData) GetWarnings() []string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *DomainSearchResponseData) GetWarningsOk() (*[]string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *DomainSearchResponseData) SetWarnings(v []string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *DomainSearchResponseData) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


