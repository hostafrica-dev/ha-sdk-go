# CatalogueConfigOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Config option identifier | 
**Name** | **string** | Config option display name | 
**Type** | **string** | Control type (e.g. selection, quantity) | 
**QtyMin** | Pointer to **int32** | Minimum quantity (null for selection types) | [optional] 
**QtyMax** | Pointer to **int32** | Maximum quantity (null for selection types) | [optional] 
**Suboptions** | [**[]CatalogueConfigSuboption**](CatalogueConfigSuboption.md) | Available suboptions for this config option | 

## Methods

### NewCatalogueConfigOption

`func NewCatalogueConfigOption(id int32, name string, type_ string, suboptions []CatalogueConfigSuboption, ) *CatalogueConfigOption`

NewCatalogueConfigOption instantiates a new CatalogueConfigOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogueConfigOptionWithDefaults

`func NewCatalogueConfigOptionWithDefaults() *CatalogueConfigOption`

NewCatalogueConfigOptionWithDefaults instantiates a new CatalogueConfigOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CatalogueConfigOption) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CatalogueConfigOption) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CatalogueConfigOption) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *CatalogueConfigOption) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CatalogueConfigOption) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CatalogueConfigOption) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *CatalogueConfigOption) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CatalogueConfigOption) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CatalogueConfigOption) SetType(v string)`

SetType sets Type field to given value.


### GetQtyMin

`func (o *CatalogueConfigOption) GetQtyMin() int32`

GetQtyMin returns the QtyMin field if non-nil, zero value otherwise.

### GetQtyMinOk

`func (o *CatalogueConfigOption) GetQtyMinOk() (*int32, bool)`

GetQtyMinOk returns a tuple with the QtyMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQtyMin

`func (o *CatalogueConfigOption) SetQtyMin(v int32)`

SetQtyMin sets QtyMin field to given value.

### HasQtyMin

`func (o *CatalogueConfigOption) HasQtyMin() bool`

HasQtyMin returns a boolean if a field has been set.

### GetQtyMax

`func (o *CatalogueConfigOption) GetQtyMax() int32`

GetQtyMax returns the QtyMax field if non-nil, zero value otherwise.

### GetQtyMaxOk

`func (o *CatalogueConfigOption) GetQtyMaxOk() (*int32, bool)`

GetQtyMaxOk returns a tuple with the QtyMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQtyMax

`func (o *CatalogueConfigOption) SetQtyMax(v int32)`

SetQtyMax sets QtyMax field to given value.

### HasQtyMax

`func (o *CatalogueConfigOption) HasQtyMax() bool`

HasQtyMax returns a boolean if a field has been set.

### GetSuboptions

`func (o *CatalogueConfigOption) GetSuboptions() []CatalogueConfigSuboption`

GetSuboptions returns the Suboptions field if non-nil, zero value otherwise.

### GetSuboptionsOk

`func (o *CatalogueConfigOption) GetSuboptionsOk() (*[]CatalogueConfigSuboption, bool)`

GetSuboptionsOk returns a tuple with the Suboptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuboptions

`func (o *CatalogueConfigOption) SetSuboptions(v []CatalogueConfigSuboption)`

SetSuboptions sets Suboptions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


