# CreateOrderRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Promo** | Pointer to **string** | Promotional code to apply to the order | [optional] 
**Products** | [**[]CreateOrderProduct**](CreateOrderProduct.md) | List of products to order | 

## Methods

### NewCreateOrderRequestContent

`func NewCreateOrderRequestContent(products []CreateOrderProduct, ) *CreateOrderRequestContent`

NewCreateOrderRequestContent instantiates a new CreateOrderRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrderRequestContentWithDefaults

`func NewCreateOrderRequestContentWithDefaults() *CreateOrderRequestContent`

NewCreateOrderRequestContentWithDefaults instantiates a new CreateOrderRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPromo

`func (o *CreateOrderRequestContent) GetPromo() string`

GetPromo returns the Promo field if non-nil, zero value otherwise.

### GetPromoOk

`func (o *CreateOrderRequestContent) GetPromoOk() (*string, bool)`

GetPromoOk returns a tuple with the Promo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromo

`func (o *CreateOrderRequestContent) SetPromo(v string)`

SetPromo sets Promo field to given value.

### HasPromo

`func (o *CreateOrderRequestContent) HasPromo() bool`

HasPromo returns a boolean if a field has been set.

### GetProducts

`func (o *CreateOrderRequestContent) GetProducts() []CreateOrderProduct`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *CreateOrderRequestContent) GetProductsOk() (*[]CreateOrderProduct, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *CreateOrderRequestContent) SetProducts(v []CreateOrderProduct)`

SetProducts sets Products field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


