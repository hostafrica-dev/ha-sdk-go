# CatalogueConfigSuboption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Suboption identifier | 
**Name** | **string** | Suboption display name | 
**Prices** | **interface{}** | Per-cycle prices for this suboption | 

## Methods

### NewCatalogueConfigSuboption

`func NewCatalogueConfigSuboption(id int32, name string, prices interface{}, ) *CatalogueConfigSuboption`

NewCatalogueConfigSuboption instantiates a new CatalogueConfigSuboption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogueConfigSuboptionWithDefaults

`func NewCatalogueConfigSuboptionWithDefaults() *CatalogueConfigSuboption`

NewCatalogueConfigSuboptionWithDefaults instantiates a new CatalogueConfigSuboption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CatalogueConfigSuboption) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CatalogueConfigSuboption) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CatalogueConfigSuboption) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *CatalogueConfigSuboption) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CatalogueConfigSuboption) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CatalogueConfigSuboption) SetName(v string)`

SetName sets Name field to given value.


### GetPrices

`func (o *CatalogueConfigSuboption) GetPrices() interface{}`

GetPrices returns the Prices field if non-nil, zero value otherwise.

### GetPricesOk

`func (o *CatalogueConfigSuboption) GetPricesOk() (*interface{}, bool)`

GetPricesOk returns a tuple with the Prices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrices

`func (o *CatalogueConfigSuboption) SetPrices(v interface{})`

SetPrices sets Prices field to given value.


### SetPricesNil

`func (o *CatalogueConfigSuboption) SetPricesNil(b bool)`

 SetPricesNil sets the value for Prices to be an explicit nil

### UnsetPrices
`func (o *CatalogueConfigSuboption) UnsetPrices()`

UnsetPrices ensures that no value is present for Prices, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


