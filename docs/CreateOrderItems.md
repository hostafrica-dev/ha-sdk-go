# CreateOrderItems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Products** | Pointer to [**[]CreateOrderProductItem**](CreateOrderProductItem.md) | Product line items in the order | [optional] 
**Domains** | Pointer to [**[]CreateOrderDomainItem**](CreateOrderDomainItem.md) | Domain line items in the order | [optional] 

## Methods

### NewCreateOrderItems

`func NewCreateOrderItems() *CreateOrderItems`

NewCreateOrderItems instantiates a new CreateOrderItems object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrderItemsWithDefaults

`func NewCreateOrderItemsWithDefaults() *CreateOrderItems`

NewCreateOrderItemsWithDefaults instantiates a new CreateOrderItems object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProducts

`func (o *CreateOrderItems) GetProducts() []CreateOrderProductItem`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *CreateOrderItems) GetProductsOk() (*[]CreateOrderProductItem, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *CreateOrderItems) SetProducts(v []CreateOrderProductItem)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *CreateOrderItems) HasProducts() bool`

HasProducts returns a boolean if a field has been set.

### GetDomains

`func (o *CreateOrderItems) GetDomains() []CreateOrderDomainItem`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *CreateOrderItems) GetDomainsOk() (*[]CreateOrderDomainItem, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *CreateOrderItems) SetDomains(v []CreateOrderDomainItem)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *CreateOrderItems) HasDomains() bool`

HasDomains returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


