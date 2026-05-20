# CatalogueData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | [**CatalogueCurrency**](CatalogueCurrency.md) |  | 
**Groups** | Pointer to [**[]CatalogueGroup**](CatalogueGroup.md) | Product groups with nested products, plans, and config options. Present when no product_id filter is applied. | [optional] 
**Product** | Pointer to [**CatalogueProduct**](CatalogueProduct.md) |  | [optional] 

## Methods

### NewCatalogueData

`func NewCatalogueData(currency CatalogueCurrency, ) *CatalogueData`

NewCatalogueData instantiates a new CatalogueData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogueDataWithDefaults

`func NewCatalogueDataWithDefaults() *CatalogueData`

NewCatalogueDataWithDefaults instantiates a new CatalogueData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *CatalogueData) GetCurrency() CatalogueCurrency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CatalogueData) GetCurrencyOk() (*CatalogueCurrency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CatalogueData) SetCurrency(v CatalogueCurrency)`

SetCurrency sets Currency field to given value.


### GetGroups

`func (o *CatalogueData) GetGroups() []CatalogueGroup`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *CatalogueData) GetGroupsOk() (*[]CatalogueGroup, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *CatalogueData) SetGroups(v []CatalogueGroup)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *CatalogueData) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetProduct

`func (o *CatalogueData) GetProduct() CatalogueProduct`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *CatalogueData) GetProductOk() (*CatalogueProduct, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *CatalogueData) SetProduct(v CatalogueProduct)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *CatalogueData) HasProduct() bool`

HasProduct returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


