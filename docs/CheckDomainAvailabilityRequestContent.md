# CheckDomainAvailabilityRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | Pointer to **string** | A single domain name to look up, e.g. \&quot;example.co.za\&quot;. Required unless domains is provided. | [optional] 
**Domains** | Pointer to **string** | Comma-separated list of domain names for a batch check. Required unless domain is provided. | [optional] 
**Currency** | Pointer to **string** | Currency for pricing. Accepts ISO code, symbol, country name, or numeric ID. Defaults to ZAR. | [optional] 

## Methods

### NewCheckDomainAvailabilityRequestContent

`func NewCheckDomainAvailabilityRequestContent() *CheckDomainAvailabilityRequestContent`

NewCheckDomainAvailabilityRequestContent instantiates a new CheckDomainAvailabilityRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckDomainAvailabilityRequestContentWithDefaults

`func NewCheckDomainAvailabilityRequestContentWithDefaults() *CheckDomainAvailabilityRequestContent`

NewCheckDomainAvailabilityRequestContentWithDefaults instantiates a new CheckDomainAvailabilityRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *CheckDomainAvailabilityRequestContent) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *CheckDomainAvailabilityRequestContent) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *CheckDomainAvailabilityRequestContent) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *CheckDomainAvailabilityRequestContent) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetDomains

`func (o *CheckDomainAvailabilityRequestContent) GetDomains() string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *CheckDomainAvailabilityRequestContent) GetDomainsOk() (*string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *CheckDomainAvailabilityRequestContent) SetDomains(v string)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *CheckDomainAvailabilityRequestContent) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### GetCurrency

`func (o *CheckDomainAvailabilityRequestContent) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CheckDomainAvailabilityRequestContent) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CheckDomainAvailabilityRequestContent) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CheckDomainAvailabilityRequestContent) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


