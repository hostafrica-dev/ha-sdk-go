# DomainAvailabilityResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | **string** | Fully qualified domain name | 
**Premium** | Pointer to **bool** | Whether the domain is a premium domain | [optional] 
**Status** | **string** | Availability status: \&quot;available\&quot; or \&quot;unavailable\&quot; | 
**Addons** | Pointer to [**DomainAddons**](DomainAddons.md) |  | [optional] 
**Pricing** | Pointer to [**DomainPricing**](DomainPricing.md) |  | [optional] 
**RequiresAdditionalInfo** | Pointer to **bool** | Whether the domain requires additional registration information | [optional] 
**AdditionalInfo** | Pointer to **interface{}** | Opaque additional registration data from upstream | [optional] 
**Group** | Pointer to **string** | Product group from upstream | [optional] 
**RegisterUrl** | Pointer to **string** | Checkout URL for available domains, enriched server-side | [optional] 

## Methods

### NewDomainAvailabilityResult

`func NewDomainAvailabilityResult(domain string, status string, ) *DomainAvailabilityResult`

NewDomainAvailabilityResult instantiates a new DomainAvailabilityResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainAvailabilityResultWithDefaults

`func NewDomainAvailabilityResultWithDefaults() *DomainAvailabilityResult`

NewDomainAvailabilityResultWithDefaults instantiates a new DomainAvailabilityResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *DomainAvailabilityResult) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *DomainAvailabilityResult) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *DomainAvailabilityResult) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetPremium

`func (o *DomainAvailabilityResult) GetPremium() bool`

GetPremium returns the Premium field if non-nil, zero value otherwise.

### GetPremiumOk

`func (o *DomainAvailabilityResult) GetPremiumOk() (*bool, bool)`

GetPremiumOk returns a tuple with the Premium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPremium

`func (o *DomainAvailabilityResult) SetPremium(v bool)`

SetPremium sets Premium field to given value.

### HasPremium

`func (o *DomainAvailabilityResult) HasPremium() bool`

HasPremium returns a boolean if a field has been set.

### GetStatus

`func (o *DomainAvailabilityResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DomainAvailabilityResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DomainAvailabilityResult) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetAddons

`func (o *DomainAvailabilityResult) GetAddons() DomainAddons`

GetAddons returns the Addons field if non-nil, zero value otherwise.

### GetAddonsOk

`func (o *DomainAvailabilityResult) GetAddonsOk() (*DomainAddons, bool)`

GetAddonsOk returns a tuple with the Addons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddons

`func (o *DomainAvailabilityResult) SetAddons(v DomainAddons)`

SetAddons sets Addons field to given value.

### HasAddons

`func (o *DomainAvailabilityResult) HasAddons() bool`

HasAddons returns a boolean if a field has been set.

### GetPricing

`func (o *DomainAvailabilityResult) GetPricing() DomainPricing`

GetPricing returns the Pricing field if non-nil, zero value otherwise.

### GetPricingOk

`func (o *DomainAvailabilityResult) GetPricingOk() (*DomainPricing, bool)`

GetPricingOk returns a tuple with the Pricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricing

`func (o *DomainAvailabilityResult) SetPricing(v DomainPricing)`

SetPricing sets Pricing field to given value.

### HasPricing

`func (o *DomainAvailabilityResult) HasPricing() bool`

HasPricing returns a boolean if a field has been set.

### GetRequiresAdditionalInfo

`func (o *DomainAvailabilityResult) GetRequiresAdditionalInfo() bool`

GetRequiresAdditionalInfo returns the RequiresAdditionalInfo field if non-nil, zero value otherwise.

### GetRequiresAdditionalInfoOk

`func (o *DomainAvailabilityResult) GetRequiresAdditionalInfoOk() (*bool, bool)`

GetRequiresAdditionalInfoOk returns a tuple with the RequiresAdditionalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresAdditionalInfo

`func (o *DomainAvailabilityResult) SetRequiresAdditionalInfo(v bool)`

SetRequiresAdditionalInfo sets RequiresAdditionalInfo field to given value.

### HasRequiresAdditionalInfo

`func (o *DomainAvailabilityResult) HasRequiresAdditionalInfo() bool`

HasRequiresAdditionalInfo returns a boolean if a field has been set.

### GetAdditionalInfo

`func (o *DomainAvailabilityResult) GetAdditionalInfo() interface{}`

GetAdditionalInfo returns the AdditionalInfo field if non-nil, zero value otherwise.

### GetAdditionalInfoOk

`func (o *DomainAvailabilityResult) GetAdditionalInfoOk() (*interface{}, bool)`

GetAdditionalInfoOk returns a tuple with the AdditionalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInfo

`func (o *DomainAvailabilityResult) SetAdditionalInfo(v interface{})`

SetAdditionalInfo sets AdditionalInfo field to given value.

### HasAdditionalInfo

`func (o *DomainAvailabilityResult) HasAdditionalInfo() bool`

HasAdditionalInfo returns a boolean if a field has been set.

### SetAdditionalInfoNil

`func (o *DomainAvailabilityResult) SetAdditionalInfoNil(b bool)`

 SetAdditionalInfoNil sets the value for AdditionalInfo to be an explicit nil

### UnsetAdditionalInfo
`func (o *DomainAvailabilityResult) UnsetAdditionalInfo()`

UnsetAdditionalInfo ensures that no value is present for AdditionalInfo, not even an explicit nil
### GetGroup

`func (o *DomainAvailabilityResult) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *DomainAvailabilityResult) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *DomainAvailabilityResult) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *DomainAvailabilityResult) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetRegisterUrl

`func (o *DomainAvailabilityResult) GetRegisterUrl() string`

GetRegisterUrl returns the RegisterUrl field if non-nil, zero value otherwise.

### GetRegisterUrlOk

`func (o *DomainAvailabilityResult) GetRegisterUrlOk() (*string, bool)`

GetRegisterUrlOk returns a tuple with the RegisterUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisterUrl

`func (o *DomainAvailabilityResult) SetRegisterUrl(v string)`

SetRegisterUrl sets RegisterUrl field to given value.

### HasRegisterUrl

`func (o *DomainAvailabilityResult) HasRegisterUrl() bool`

HasRegisterUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


