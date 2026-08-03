# DomainAddonFeature

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** |  | 
**Price** | **float64** |  | 
**Title** | **string** |  | 
**Description** | **string** |  | 
**ProductDesc** | Pointer to **string** | Optional HTML product description from upstream | [optional] 

## Methods

### NewDomainAddonFeature

`func NewDomainAddonFeature(enabled bool, price float64, title string, description string, ) *DomainAddonFeature`

NewDomainAddonFeature instantiates a new DomainAddonFeature object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainAddonFeatureWithDefaults

`func NewDomainAddonFeatureWithDefaults() *DomainAddonFeature`

NewDomainAddonFeatureWithDefaults instantiates a new DomainAddonFeature object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *DomainAddonFeature) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DomainAddonFeature) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DomainAddonFeature) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetPrice

`func (o *DomainAddonFeature) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *DomainAddonFeature) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *DomainAddonFeature) SetPrice(v float64)`

SetPrice sets Price field to given value.


### GetTitle

`func (o *DomainAddonFeature) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *DomainAddonFeature) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *DomainAddonFeature) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *DomainAddonFeature) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DomainAddonFeature) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DomainAddonFeature) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetProductDesc

`func (o *DomainAddonFeature) GetProductDesc() string`

GetProductDesc returns the ProductDesc field if non-nil, zero value otherwise.

### GetProductDescOk

`func (o *DomainAddonFeature) GetProductDescOk() (*string, bool)`

GetProductDescOk returns a tuple with the ProductDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductDesc

`func (o *DomainAddonFeature) SetProductDesc(v string)`

SetProductDesc sets ProductDesc field to given value.

### HasProductDesc

`func (o *DomainAddonFeature) HasProductDesc() bool`

HasProductDesc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


