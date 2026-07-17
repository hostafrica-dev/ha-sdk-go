# CatalogueProduct

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Product identifier | 
**Name** | **string** | Product display name | 
**Type** | **string** | Product type (e.g. server) | 
**RequiresHostname** | **bool** | Whether the product requires a hostname at order time | 
**BillingCycles** | **[]string** | Billing cycles available for this product | 
**Pricing** | **interface{}** | Base per-cycle pricing for the product (before plan selection) | 
**UsePlans** | **bool** | Whether this product uses plan-based sizing | 
**Plans** | [**[]CataloguePlan**](CataloguePlan.md) | Available plans (size tiers) for this product | 
**ConfigOptions** | [**[]CatalogueConfigOption**](CatalogueConfigOption.md) | Configurable add-on options for this product | 
**AdditionalInformation** | **interface{}** | Additional product information fields | 

## Methods

### NewCatalogueProduct

`func NewCatalogueProduct(id int32, name string, type_ string, requiresHostname bool, billingCycles []string, pricing interface{}, usePlans bool, plans []CataloguePlan, configOptions []CatalogueConfigOption, additionalInformation interface{}, ) *CatalogueProduct`

NewCatalogueProduct instantiates a new CatalogueProduct object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogueProductWithDefaults

`func NewCatalogueProductWithDefaults() *CatalogueProduct`

NewCatalogueProductWithDefaults instantiates a new CatalogueProduct object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CatalogueProduct) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CatalogueProduct) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CatalogueProduct) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *CatalogueProduct) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CatalogueProduct) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CatalogueProduct) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *CatalogueProduct) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CatalogueProduct) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CatalogueProduct) SetType(v string)`

SetType sets Type field to given value.


### GetRequiresHostname

`func (o *CatalogueProduct) GetRequiresHostname() bool`

GetRequiresHostname returns the RequiresHostname field if non-nil, zero value otherwise.

### GetRequiresHostnameOk

`func (o *CatalogueProduct) GetRequiresHostnameOk() (*bool, bool)`

GetRequiresHostnameOk returns a tuple with the RequiresHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresHostname

`func (o *CatalogueProduct) SetRequiresHostname(v bool)`

SetRequiresHostname sets RequiresHostname field to given value.


### GetBillingCycles

`func (o *CatalogueProduct) GetBillingCycles() []string`

GetBillingCycles returns the BillingCycles field if non-nil, zero value otherwise.

### GetBillingCyclesOk

`func (o *CatalogueProduct) GetBillingCyclesOk() (*[]string, bool)`

GetBillingCyclesOk returns a tuple with the BillingCycles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCycles

`func (o *CatalogueProduct) SetBillingCycles(v []string)`

SetBillingCycles sets BillingCycles field to given value.


### GetPricing

`func (o *CatalogueProduct) GetPricing() interface{}`

GetPricing returns the Pricing field if non-nil, zero value otherwise.

### GetPricingOk

`func (o *CatalogueProduct) GetPricingOk() (*interface{}, bool)`

GetPricingOk returns a tuple with the Pricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricing

`func (o *CatalogueProduct) SetPricing(v interface{})`

SetPricing sets Pricing field to given value.


### SetPricingNil

`func (o *CatalogueProduct) SetPricingNil(b bool)`

 SetPricingNil sets the value for Pricing to be an explicit nil

### UnsetPricing
`func (o *CatalogueProduct) UnsetPricing()`

UnsetPricing ensures that no value is present for Pricing, not even an explicit nil
### GetUsePlans

`func (o *CatalogueProduct) GetUsePlans() bool`

GetUsePlans returns the UsePlans field if non-nil, zero value otherwise.

### GetUsePlansOk

`func (o *CatalogueProduct) GetUsePlansOk() (*bool, bool)`

GetUsePlansOk returns a tuple with the UsePlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsePlans

`func (o *CatalogueProduct) SetUsePlans(v bool)`

SetUsePlans sets UsePlans field to given value.


### GetPlans

`func (o *CatalogueProduct) GetPlans() []CataloguePlan`

GetPlans returns the Plans field if non-nil, zero value otherwise.

### GetPlansOk

`func (o *CatalogueProduct) GetPlansOk() (*[]CataloguePlan, bool)`

GetPlansOk returns a tuple with the Plans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlans

`func (o *CatalogueProduct) SetPlans(v []CataloguePlan)`

SetPlans sets Plans field to given value.


### GetConfigOptions

`func (o *CatalogueProduct) GetConfigOptions() []CatalogueConfigOption`

GetConfigOptions returns the ConfigOptions field if non-nil, zero value otherwise.

### GetConfigOptionsOk

`func (o *CatalogueProduct) GetConfigOptionsOk() (*[]CatalogueConfigOption, bool)`

GetConfigOptionsOk returns a tuple with the ConfigOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigOptions

`func (o *CatalogueProduct) SetConfigOptions(v []CatalogueConfigOption)`

SetConfigOptions sets ConfigOptions field to given value.


### GetAdditionalInformation

`func (o *CatalogueProduct) GetAdditionalInformation() interface{}`

GetAdditionalInformation returns the AdditionalInformation field if non-nil, zero value otherwise.

### GetAdditionalInformationOk

`func (o *CatalogueProduct) GetAdditionalInformationOk() (*interface{}, bool)`

GetAdditionalInformationOk returns a tuple with the AdditionalInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInformation

`func (o *CatalogueProduct) SetAdditionalInformation(v interface{})`

SetAdditionalInformation sets AdditionalInformation field to given value.


### SetAdditionalInformationNil

`func (o *CatalogueProduct) SetAdditionalInformationNil(b bool)`

 SetAdditionalInformationNil sets the value for AdditionalInformation to be an explicit nil

### UnsetAdditionalInformation
`func (o *CatalogueProduct) UnsetAdditionalInformation()`

UnsetAdditionalInformation ensures that no value is present for AdditionalInformation, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


