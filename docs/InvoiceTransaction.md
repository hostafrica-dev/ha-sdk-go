# InvoiceTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Transaction identifier | 
**Gateway** | Pointer to **string** | Payment gateway identifier | [optional] 
**Date** | Pointer to **string** | Transaction date (ISO 8601 or upstream date string) | [optional] 
**Description** | Pointer to **string** | Transaction description | [optional] 
**AmountIn** | Pointer to **string** | Amount received as a decimal string | [optional] 
**Fees** | Pointer to **string** | Gateway fees as a decimal string | [optional] 
**AmountOut** | Pointer to **string** | Amount paid out (e.g. refund) as a decimal string | [optional] 
**TransId** | Pointer to **string** | Gateway transaction reference | [optional] 

## Methods

### NewInvoiceTransaction

`func NewInvoiceTransaction(id int32, ) *InvoiceTransaction`

NewInvoiceTransaction instantiates a new InvoiceTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceTransactionWithDefaults

`func NewInvoiceTransactionWithDefaults() *InvoiceTransaction`

NewInvoiceTransactionWithDefaults instantiates a new InvoiceTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InvoiceTransaction) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InvoiceTransaction) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InvoiceTransaction) SetId(v int32)`

SetId sets Id field to given value.


### GetGateway

`func (o *InvoiceTransaction) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *InvoiceTransaction) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *InvoiceTransaction) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *InvoiceTransaction) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetDate

`func (o *InvoiceTransaction) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *InvoiceTransaction) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *InvoiceTransaction) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *InvoiceTransaction) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetDescription

`func (o *InvoiceTransaction) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InvoiceTransaction) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InvoiceTransaction) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InvoiceTransaction) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAmountIn

`func (o *InvoiceTransaction) GetAmountIn() string`

GetAmountIn returns the AmountIn field if non-nil, zero value otherwise.

### GetAmountInOk

`func (o *InvoiceTransaction) GetAmountInOk() (*string, bool)`

GetAmountInOk returns a tuple with the AmountIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountIn

`func (o *InvoiceTransaction) SetAmountIn(v string)`

SetAmountIn sets AmountIn field to given value.

### HasAmountIn

`func (o *InvoiceTransaction) HasAmountIn() bool`

HasAmountIn returns a boolean if a field has been set.

### GetFees

`func (o *InvoiceTransaction) GetFees() string`

GetFees returns the Fees field if non-nil, zero value otherwise.

### GetFeesOk

`func (o *InvoiceTransaction) GetFeesOk() (*string, bool)`

GetFeesOk returns a tuple with the Fees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFees

`func (o *InvoiceTransaction) SetFees(v string)`

SetFees sets Fees field to given value.

### HasFees

`func (o *InvoiceTransaction) HasFees() bool`

HasFees returns a boolean if a field has been set.

### GetAmountOut

`func (o *InvoiceTransaction) GetAmountOut() string`

GetAmountOut returns the AmountOut field if non-nil, zero value otherwise.

### GetAmountOutOk

`func (o *InvoiceTransaction) GetAmountOutOk() (*string, bool)`

GetAmountOutOk returns a tuple with the AmountOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountOut

`func (o *InvoiceTransaction) SetAmountOut(v string)`

SetAmountOut sets AmountOut field to given value.

### HasAmountOut

`func (o *InvoiceTransaction) HasAmountOut() bool`

HasAmountOut returns a boolean if a field has been set.

### GetTransId

`func (o *InvoiceTransaction) GetTransId() string`

GetTransId returns the TransId field if non-nil, zero value otherwise.

### GetTransIdOk

`func (o *InvoiceTransaction) GetTransIdOk() (*string, bool)`

GetTransIdOk returns a tuple with the TransId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransId

`func (o *InvoiceTransaction) SetTransId(v string)`

SetTransId sets TransId field to given value.

### HasTransId

`func (o *InvoiceTransaction) HasTransId() bool`

HasTransId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


