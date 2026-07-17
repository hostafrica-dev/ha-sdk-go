# DomainPricing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domainregister** | Pointer to [**[]DomainPricingEntry**](DomainPricingEntry.md) |  | [optional] 
**Domainrenew** | Pointer to [**[]DomainPricingEntry**](DomainPricingEntry.md) |  | [optional] 
**Domaintransfer** | Pointer to [**[]DomainPricingEntry**](DomainPricingEntry.md) |  | [optional] 

## Methods

### NewDomainPricing

`func NewDomainPricing() *DomainPricing`

NewDomainPricing instantiates a new DomainPricing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainPricingWithDefaults

`func NewDomainPricingWithDefaults() *DomainPricing`

NewDomainPricingWithDefaults instantiates a new DomainPricing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomainregister

`func (o *DomainPricing) GetDomainregister() []DomainPricingEntry`

GetDomainregister returns the Domainregister field if non-nil, zero value otherwise.

### GetDomainregisterOk

`func (o *DomainPricing) GetDomainregisterOk() (*[]DomainPricingEntry, bool)`

GetDomainregisterOk returns a tuple with the Domainregister field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainregister

`func (o *DomainPricing) SetDomainregister(v []DomainPricingEntry)`

SetDomainregister sets Domainregister field to given value.

### HasDomainregister

`func (o *DomainPricing) HasDomainregister() bool`

HasDomainregister returns a boolean if a field has been set.

### GetDomainrenew

`func (o *DomainPricing) GetDomainrenew() []DomainPricingEntry`

GetDomainrenew returns the Domainrenew field if non-nil, zero value otherwise.

### GetDomainrenewOk

`func (o *DomainPricing) GetDomainrenewOk() (*[]DomainPricingEntry, bool)`

GetDomainrenewOk returns a tuple with the Domainrenew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainrenew

`func (o *DomainPricing) SetDomainrenew(v []DomainPricingEntry)`

SetDomainrenew sets Domainrenew field to given value.

### HasDomainrenew

`func (o *DomainPricing) HasDomainrenew() bool`

HasDomainrenew returns a boolean if a field has been set.

### GetDomaintransfer

`func (o *DomainPricing) GetDomaintransfer() []DomainPricingEntry`

GetDomaintransfer returns the Domaintransfer field if non-nil, zero value otherwise.

### GetDomaintransferOk

`func (o *DomainPricing) GetDomaintransferOk() (*[]DomainPricingEntry, bool)`

GetDomaintransferOk returns a tuple with the Domaintransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomaintransfer

`func (o *DomainPricing) SetDomaintransfer(v []DomainPricingEntry)`

SetDomaintransfer sets Domaintransfer field to given value.

### HasDomaintransfer

`func (o *DomainPricing) HasDomaintransfer() bool`

HasDomaintransfer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


