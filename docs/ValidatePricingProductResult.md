# ValidatePricingProductResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pid** | **int32** |  | 
**Name** | **string** |  | 
**BillingCycle** | **string** |  | 
**PlanId** | Pointer to **int32** |  | [optional] 
**Prorata** | Pointer to [**ValidatePricingProrata**](ValidatePricingProrata.md) |  | [optional] 
**RecurringPrice** | Pointer to [**ValidatePricingPriceRange**](ValidatePricingPriceRange.md) |  | [optional] 
**Discount** | Pointer to [**ValidatePricingDiscount**](ValidatePricingDiscount.md) |  | [optional] 
**LineTotalBeforeDiscount** | **string** |  | 
**LineTotal** | **string** |  | 
**Breakdown** | Pointer to [**ValidatePricingBreakdown**](ValidatePricingBreakdown.md) |  | [optional] 

## Methods

### NewValidatePricingProductResult

`func NewValidatePricingProductResult(pid int32, name string, billingCycle string, lineTotalBeforeDiscount string, lineTotal string, ) *ValidatePricingProductResult`

NewValidatePricingProductResult instantiates a new ValidatePricingProductResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidatePricingProductResultWithDefaults

`func NewValidatePricingProductResultWithDefaults() *ValidatePricingProductResult`

NewValidatePricingProductResultWithDefaults instantiates a new ValidatePricingProductResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPid

`func (o *ValidatePricingProductResult) GetPid() int32`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *ValidatePricingProductResult) GetPidOk() (*int32, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *ValidatePricingProductResult) SetPid(v int32)`

SetPid sets Pid field to given value.


### GetName

`func (o *ValidatePricingProductResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ValidatePricingProductResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ValidatePricingProductResult) SetName(v string)`

SetName sets Name field to given value.


### GetBillingCycle

`func (o *ValidatePricingProductResult) GetBillingCycle() string`

GetBillingCycle returns the BillingCycle field if non-nil, zero value otherwise.

### GetBillingCycleOk

`func (o *ValidatePricingProductResult) GetBillingCycleOk() (*string, bool)`

GetBillingCycleOk returns a tuple with the BillingCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCycle

`func (o *ValidatePricingProductResult) SetBillingCycle(v string)`

SetBillingCycle sets BillingCycle field to given value.


### GetPlanId

`func (o *ValidatePricingProductResult) GetPlanId() int32`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *ValidatePricingProductResult) GetPlanIdOk() (*int32, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *ValidatePricingProductResult) SetPlanId(v int32)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *ValidatePricingProductResult) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetProrata

`func (o *ValidatePricingProductResult) GetProrata() ValidatePricingProrata`

GetProrata returns the Prorata field if non-nil, zero value otherwise.

### GetProrataOk

`func (o *ValidatePricingProductResult) GetProrataOk() (*ValidatePricingProrata, bool)`

GetProrataOk returns a tuple with the Prorata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProrata

`func (o *ValidatePricingProductResult) SetProrata(v ValidatePricingProrata)`

SetProrata sets Prorata field to given value.

### HasProrata

`func (o *ValidatePricingProductResult) HasProrata() bool`

HasProrata returns a boolean if a field has been set.

### GetRecurringPrice

`func (o *ValidatePricingProductResult) GetRecurringPrice() ValidatePricingPriceRange`

GetRecurringPrice returns the RecurringPrice field if non-nil, zero value otherwise.

### GetRecurringPriceOk

`func (o *ValidatePricingProductResult) GetRecurringPriceOk() (*ValidatePricingPriceRange, bool)`

GetRecurringPriceOk returns a tuple with the RecurringPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringPrice

`func (o *ValidatePricingProductResult) SetRecurringPrice(v ValidatePricingPriceRange)`

SetRecurringPrice sets RecurringPrice field to given value.

### HasRecurringPrice

`func (o *ValidatePricingProductResult) HasRecurringPrice() bool`

HasRecurringPrice returns a boolean if a field has been set.

### GetDiscount

`func (o *ValidatePricingProductResult) GetDiscount() ValidatePricingDiscount`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *ValidatePricingProductResult) GetDiscountOk() (*ValidatePricingDiscount, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *ValidatePricingProductResult) SetDiscount(v ValidatePricingDiscount)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *ValidatePricingProductResult) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### GetLineTotalBeforeDiscount

`func (o *ValidatePricingProductResult) GetLineTotalBeforeDiscount() string`

GetLineTotalBeforeDiscount returns the LineTotalBeforeDiscount field if non-nil, zero value otherwise.

### GetLineTotalBeforeDiscountOk

`func (o *ValidatePricingProductResult) GetLineTotalBeforeDiscountOk() (*string, bool)`

GetLineTotalBeforeDiscountOk returns a tuple with the LineTotalBeforeDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineTotalBeforeDiscount

`func (o *ValidatePricingProductResult) SetLineTotalBeforeDiscount(v string)`

SetLineTotalBeforeDiscount sets LineTotalBeforeDiscount field to given value.


### GetLineTotal

`func (o *ValidatePricingProductResult) GetLineTotal() string`

GetLineTotal returns the LineTotal field if non-nil, zero value otherwise.

### GetLineTotalOk

`func (o *ValidatePricingProductResult) GetLineTotalOk() (*string, bool)`

GetLineTotalOk returns a tuple with the LineTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineTotal

`func (o *ValidatePricingProductResult) SetLineTotal(v string)`

SetLineTotal sets LineTotal field to given value.


### GetBreakdown

`func (o *ValidatePricingProductResult) GetBreakdown() ValidatePricingBreakdown`

GetBreakdown returns the Breakdown field if non-nil, zero value otherwise.

### GetBreakdownOk

`func (o *ValidatePricingProductResult) GetBreakdownOk() (*ValidatePricingBreakdown, bool)`

GetBreakdownOk returns a tuple with the Breakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdown

`func (o *ValidatePricingProductResult) SetBreakdown(v ValidatePricingBreakdown)`

SetBreakdown sets Breakdown field to given value.

### HasBreakdown

`func (o *ValidatePricingProductResult) HasBreakdown() bool`

HasBreakdown returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


