# DomainPricingEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Period** | **int32** |  | 
**Price** | **string** | Price as a decimal string with two decimal places | 

## Methods

### NewDomainPricingEntry

`func NewDomainPricingEntry(period int32, price string, ) *DomainPricingEntry`

NewDomainPricingEntry instantiates a new DomainPricingEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainPricingEntryWithDefaults

`func NewDomainPricingEntryWithDefaults() *DomainPricingEntry`

NewDomainPricingEntryWithDefaults instantiates a new DomainPricingEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPeriod

`func (o *DomainPricingEntry) GetPeriod() int32`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *DomainPricingEntry) GetPeriodOk() (*int32, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *DomainPricingEntry) SetPeriod(v int32)`

SetPeriod sets Period field to given value.


### GetPrice

`func (o *DomainPricingEntry) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *DomainPricingEntry) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *DomainPricingEntry) SetPrice(v string)`

SetPrice sets Price field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


