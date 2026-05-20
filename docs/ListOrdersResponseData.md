# ListOrdersResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Orders** | [**[]Order**](Order.md) | List of orders for the authenticated user | 

## Methods

### NewListOrdersResponseData

`func NewListOrdersResponseData(orders []Order, ) *ListOrdersResponseData`

NewListOrdersResponseData instantiates a new ListOrdersResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListOrdersResponseDataWithDefaults

`func NewListOrdersResponseDataWithDefaults() *ListOrdersResponseData`

NewListOrdersResponseDataWithDefaults instantiates a new ListOrdersResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrders

`func (o *ListOrdersResponseData) GetOrders() []Order`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *ListOrdersResponseData) GetOrdersOk() (*[]Order, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *ListOrdersResponseData) SetOrders(v []Order)`

SetOrders sets Orders field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


