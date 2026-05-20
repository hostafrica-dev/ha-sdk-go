# ValidatePricingDiscount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Applied** | **bool** |  | 
**Code** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Amount** | **string** |  | 
**RecurringAmount** | **string** |  | 
**ApplyOnce** | **bool** |  | 

## Methods

### NewValidatePricingDiscount

`func NewValidatePricingDiscount(applied bool, amount string, recurringAmount string, applyOnce bool, ) *ValidatePricingDiscount`

NewValidatePricingDiscount instantiates a new ValidatePricingDiscount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidatePricingDiscountWithDefaults

`func NewValidatePricingDiscountWithDefaults() *ValidatePricingDiscount`

NewValidatePricingDiscountWithDefaults instantiates a new ValidatePricingDiscount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplied

`func (o *ValidatePricingDiscount) GetApplied() bool`

GetApplied returns the Applied field if non-nil, zero value otherwise.

### GetAppliedOk

`func (o *ValidatePricingDiscount) GetAppliedOk() (*bool, bool)`

GetAppliedOk returns a tuple with the Applied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplied

`func (o *ValidatePricingDiscount) SetApplied(v bool)`

SetApplied sets Applied field to given value.


### GetCode

`func (o *ValidatePricingDiscount) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ValidatePricingDiscount) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ValidatePricingDiscount) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ValidatePricingDiscount) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetType

`func (o *ValidatePricingDiscount) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ValidatePricingDiscount) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ValidatePricingDiscount) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ValidatePricingDiscount) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAmount

`func (o *ValidatePricingDiscount) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ValidatePricingDiscount) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ValidatePricingDiscount) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetRecurringAmount

`func (o *ValidatePricingDiscount) GetRecurringAmount() string`

GetRecurringAmount returns the RecurringAmount field if non-nil, zero value otherwise.

### GetRecurringAmountOk

`func (o *ValidatePricingDiscount) GetRecurringAmountOk() (*string, bool)`

GetRecurringAmountOk returns a tuple with the RecurringAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringAmount

`func (o *ValidatePricingDiscount) SetRecurringAmount(v string)`

SetRecurringAmount sets RecurringAmount field to given value.


### GetApplyOnce

`func (o *ValidatePricingDiscount) GetApplyOnce() bool`

GetApplyOnce returns the ApplyOnce field if non-nil, zero value otherwise.

### GetApplyOnceOk

`func (o *ValidatePricingDiscount) GetApplyOnceOk() (*bool, bool)`

GetApplyOnceOk returns a tuple with the ApplyOnce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyOnce

`func (o *ValidatePricingDiscount) SetApplyOnce(v bool)`

SetApplyOnce sets ApplyOnce field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


