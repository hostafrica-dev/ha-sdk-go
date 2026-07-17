# CreateOrderProductItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LineId** | **int32** | Line item identifier | 
**ServiceId** | **int32** | Provisioned service identifier | 
**Pid** | **int32** | Product ID | 
**Name** | **string** | Product name | 
**BillingCycle** | **string** | Billing cycle (e.g. monthly, annually) | 
**Domain** | Pointer to **string** | Domain associated with this service | [optional] 
**Hostname** | Pointer to **string** | Hostname assigned to this service | [optional] 
**Amount** | **string** | Line item amount charged (decimal string, e.g. \&quot;2333.33\&quot;) | 

## Methods

### NewCreateOrderProductItem

`func NewCreateOrderProductItem(lineId int32, serviceId int32, pid int32, name string, billingCycle string, amount string, ) *CreateOrderProductItem`

NewCreateOrderProductItem instantiates a new CreateOrderProductItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrderProductItemWithDefaults

`func NewCreateOrderProductItemWithDefaults() *CreateOrderProductItem`

NewCreateOrderProductItemWithDefaults instantiates a new CreateOrderProductItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLineId

`func (o *CreateOrderProductItem) GetLineId() int32`

GetLineId returns the LineId field if non-nil, zero value otherwise.

### GetLineIdOk

`func (o *CreateOrderProductItem) GetLineIdOk() (*int32, bool)`

GetLineIdOk returns a tuple with the LineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineId

`func (o *CreateOrderProductItem) SetLineId(v int32)`

SetLineId sets LineId field to given value.


### GetServiceId

`func (o *CreateOrderProductItem) GetServiceId() int32`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *CreateOrderProductItem) GetServiceIdOk() (*int32, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *CreateOrderProductItem) SetServiceId(v int32)`

SetServiceId sets ServiceId field to given value.


### GetPid

`func (o *CreateOrderProductItem) GetPid() int32`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *CreateOrderProductItem) GetPidOk() (*int32, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *CreateOrderProductItem) SetPid(v int32)`

SetPid sets Pid field to given value.


### GetName

`func (o *CreateOrderProductItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateOrderProductItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateOrderProductItem) SetName(v string)`

SetName sets Name field to given value.


### GetBillingCycle

`func (o *CreateOrderProductItem) GetBillingCycle() string`

GetBillingCycle returns the BillingCycle field if non-nil, zero value otherwise.

### GetBillingCycleOk

`func (o *CreateOrderProductItem) GetBillingCycleOk() (*string, bool)`

GetBillingCycleOk returns a tuple with the BillingCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCycle

`func (o *CreateOrderProductItem) SetBillingCycle(v string)`

SetBillingCycle sets BillingCycle field to given value.


### GetDomain

`func (o *CreateOrderProductItem) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *CreateOrderProductItem) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *CreateOrderProductItem) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *CreateOrderProductItem) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetHostname

`func (o *CreateOrderProductItem) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *CreateOrderProductItem) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *CreateOrderProductItem) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *CreateOrderProductItem) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetAmount

`func (o *CreateOrderProductItem) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CreateOrderProductItem) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CreateOrderProductItem) SetAmount(v string)`

SetAmount sets Amount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


