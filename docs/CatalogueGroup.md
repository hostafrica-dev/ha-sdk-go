# CatalogueGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Group identifier | 
**Name** | **string** | Group display name (e.g. Linux Cloud Servers) | 
**Products** | [**[]CatalogueProduct**](CatalogueProduct.md) | Products within this group | 

## Methods

### NewCatalogueGroup

`func NewCatalogueGroup(id int32, name string, products []CatalogueProduct, ) *CatalogueGroup`

NewCatalogueGroup instantiates a new CatalogueGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogueGroupWithDefaults

`func NewCatalogueGroupWithDefaults() *CatalogueGroup`

NewCatalogueGroupWithDefaults instantiates a new CatalogueGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CatalogueGroup) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CatalogueGroup) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CatalogueGroup) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *CatalogueGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CatalogueGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CatalogueGroup) SetName(v string)`

SetName sets Name field to given value.


### GetProducts

`func (o *CatalogueGroup) GetProducts() []CatalogueProduct`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *CatalogueGroup) GetProductsOk() (*[]CatalogueProduct, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *CatalogueGroup) SetProducts(v []CatalogueProduct)`

SetProducts sets Products field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


