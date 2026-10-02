# InvoiceItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Line item identifier | 
**Type** | **string** | Line item type (e.g. Hosting) | 
**Description** | **string** | Line item description | 
**Amount** | **string** | Line item amount as a decimal string | 
**Taxed** | **int32** | Whether the line item is taxed (1 &#x3D; taxed, 0 &#x3D; not taxed) | 

## Methods

### NewInvoiceItem

`func NewInvoiceItem(id int32, type_ string, description string, amount string, taxed int32, ) *InvoiceItem`

NewInvoiceItem instantiates a new InvoiceItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceItemWithDefaults

`func NewInvoiceItemWithDefaults() *InvoiceItem`

NewInvoiceItemWithDefaults instantiates a new InvoiceItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InvoiceItem) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InvoiceItem) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InvoiceItem) SetId(v int32)`

SetId sets Id field to given value.


### GetType

`func (o *InvoiceItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InvoiceItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InvoiceItem) SetType(v string)`

SetType sets Type field to given value.


### GetDescription

`func (o *InvoiceItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InvoiceItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InvoiceItem) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetAmount

`func (o *InvoiceItem) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *InvoiceItem) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *InvoiceItem) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetTaxed

`func (o *InvoiceItem) GetTaxed() int32`

GetTaxed returns the Taxed field if non-nil, zero value otherwise.

### GetTaxedOk

`func (o *InvoiceItem) GetTaxedOk() (*int32, bool)`

GetTaxedOk returns a tuple with the Taxed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxed

`func (o *InvoiceItem) SetTaxed(v int32)`

SetTaxed sets Taxed field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


