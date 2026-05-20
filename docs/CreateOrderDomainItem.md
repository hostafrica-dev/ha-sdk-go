# CreateOrderDomainItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LineId** | **int32** | Line item identifier | 
**DomainId** | **int32** | Domain identifier | 
**Domain** | **string** | Domain name | 
**Type** | **string** | Domain operation type (e.g. register, transfer) | 
**Period** | **int32** | Registration period in years | 
**DomainWarranty** | **bool** | Whether domain warranty is included | 
**Autorenew** | **bool** | Whether auto-renewal is enabled | 
**Amount** | **string** | Line item amount charged (decimal string, e.g. \&quot;2333.33\&quot;) | 

## Methods

### NewCreateOrderDomainItem

`func NewCreateOrderDomainItem(lineId int32, domainId int32, domain string, type_ string, period int32, domainWarranty bool, autorenew bool, amount string, ) *CreateOrderDomainItem`

NewCreateOrderDomainItem instantiates a new CreateOrderDomainItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrderDomainItemWithDefaults

`func NewCreateOrderDomainItemWithDefaults() *CreateOrderDomainItem`

NewCreateOrderDomainItemWithDefaults instantiates a new CreateOrderDomainItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLineId

`func (o *CreateOrderDomainItem) GetLineId() int32`

GetLineId returns the LineId field if non-nil, zero value otherwise.

### GetLineIdOk

`func (o *CreateOrderDomainItem) GetLineIdOk() (*int32, bool)`

GetLineIdOk returns a tuple with the LineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineId

`func (o *CreateOrderDomainItem) SetLineId(v int32)`

SetLineId sets LineId field to given value.


### GetDomainId

`func (o *CreateOrderDomainItem) GetDomainId() int32`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *CreateOrderDomainItem) GetDomainIdOk() (*int32, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *CreateOrderDomainItem) SetDomainId(v int32)`

SetDomainId sets DomainId field to given value.


### GetDomain

`func (o *CreateOrderDomainItem) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *CreateOrderDomainItem) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *CreateOrderDomainItem) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetType

`func (o *CreateOrderDomainItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateOrderDomainItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateOrderDomainItem) SetType(v string)`

SetType sets Type field to given value.


### GetPeriod

`func (o *CreateOrderDomainItem) GetPeriod() int32`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *CreateOrderDomainItem) GetPeriodOk() (*int32, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *CreateOrderDomainItem) SetPeriod(v int32)`

SetPeriod sets Period field to given value.


### GetDomainWarranty

`func (o *CreateOrderDomainItem) GetDomainWarranty() bool`

GetDomainWarranty returns the DomainWarranty field if non-nil, zero value otherwise.

### GetDomainWarrantyOk

`func (o *CreateOrderDomainItem) GetDomainWarrantyOk() (*bool, bool)`

GetDomainWarrantyOk returns a tuple with the DomainWarranty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainWarranty

`func (o *CreateOrderDomainItem) SetDomainWarranty(v bool)`

SetDomainWarranty sets DomainWarranty field to given value.


### GetAutorenew

`func (o *CreateOrderDomainItem) GetAutorenew() bool`

GetAutorenew returns the Autorenew field if non-nil, zero value otherwise.

### GetAutorenewOk

`func (o *CreateOrderDomainItem) GetAutorenewOk() (*bool, bool)`

GetAutorenewOk returns a tuple with the Autorenew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutorenew

`func (o *CreateOrderDomainItem) SetAutorenew(v bool)`

SetAutorenew sets Autorenew field to given value.


### GetAmount

`func (o *CreateOrderDomainItem) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CreateOrderDomainItem) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CreateOrderDomainItem) SetAmount(v string)`

SetAmount sets Amount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


