# ValidatePricingRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Promo** | Pointer to **string** | Optional promotional code to apply | [optional] 
**Products** | [**[]ValidatePricingProduct**](ValidatePricingProduct.md) | List of products to validate pricing for | 

## Methods

### NewValidatePricingRequestContent

`func NewValidatePricingRequestContent(products []ValidatePricingProduct, ) *ValidatePricingRequestContent`

NewValidatePricingRequestContent instantiates a new ValidatePricingRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidatePricingRequestContentWithDefaults

`func NewValidatePricingRequestContentWithDefaults() *ValidatePricingRequestContent`

NewValidatePricingRequestContentWithDefaults instantiates a new ValidatePricingRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPromo

`func (o *ValidatePricingRequestContent) GetPromo() string`

GetPromo returns the Promo field if non-nil, zero value otherwise.

### GetPromoOk

`func (o *ValidatePricingRequestContent) GetPromoOk() (*string, bool)`

GetPromoOk returns a tuple with the Promo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromo

`func (o *ValidatePricingRequestContent) SetPromo(v string)`

SetPromo sets Promo field to given value.

### HasPromo

`func (o *ValidatePricingRequestContent) HasPromo() bool`

HasPromo returns a boolean if a field has been set.

### GetProducts

`func (o *ValidatePricingRequestContent) GetProducts() []ValidatePricingProduct`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *ValidatePricingRequestContent) GetProductsOk() (*[]ValidatePricingProduct, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *ValidatePricingRequestContent) SetProducts(v []ValidatePricingProduct)`

SetProducts sets Products field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


