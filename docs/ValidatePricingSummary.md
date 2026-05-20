# ValidatePricingSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subtotal** | **string** |  | 
**DiscountTotal** | **string** |  | 
**TotalDue** | **string** |  | 
**Recurring** | Pointer to [**ValidatePricingSummaryRecurring**](ValidatePricingSummaryRecurring.md) |  | [optional] 
**ProrataTotal** | Pointer to **string** |  | [optional] 
**PromoApplied** | Pointer to **string** |  | [optional] 

## Methods

### NewValidatePricingSummary

`func NewValidatePricingSummary(subtotal string, discountTotal string, totalDue string, ) *ValidatePricingSummary`

NewValidatePricingSummary instantiates a new ValidatePricingSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidatePricingSummaryWithDefaults

`func NewValidatePricingSummaryWithDefaults() *ValidatePricingSummary`

NewValidatePricingSummaryWithDefaults instantiates a new ValidatePricingSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubtotal

`func (o *ValidatePricingSummary) GetSubtotal() string`

GetSubtotal returns the Subtotal field if non-nil, zero value otherwise.

### GetSubtotalOk

`func (o *ValidatePricingSummary) GetSubtotalOk() (*string, bool)`

GetSubtotalOk returns a tuple with the Subtotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtotal

`func (o *ValidatePricingSummary) SetSubtotal(v string)`

SetSubtotal sets Subtotal field to given value.


### GetDiscountTotal

`func (o *ValidatePricingSummary) GetDiscountTotal() string`

GetDiscountTotal returns the DiscountTotal field if non-nil, zero value otherwise.

### GetDiscountTotalOk

`func (o *ValidatePricingSummary) GetDiscountTotalOk() (*string, bool)`

GetDiscountTotalOk returns a tuple with the DiscountTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountTotal

`func (o *ValidatePricingSummary) SetDiscountTotal(v string)`

SetDiscountTotal sets DiscountTotal field to given value.


### GetTotalDue

`func (o *ValidatePricingSummary) GetTotalDue() string`

GetTotalDue returns the TotalDue field if non-nil, zero value otherwise.

### GetTotalDueOk

`func (o *ValidatePricingSummary) GetTotalDueOk() (*string, bool)`

GetTotalDueOk returns a tuple with the TotalDue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDue

`func (o *ValidatePricingSummary) SetTotalDue(v string)`

SetTotalDue sets TotalDue field to given value.


### GetRecurring

`func (o *ValidatePricingSummary) GetRecurring() ValidatePricingSummaryRecurring`

GetRecurring returns the Recurring field if non-nil, zero value otherwise.

### GetRecurringOk

`func (o *ValidatePricingSummary) GetRecurringOk() (*ValidatePricingSummaryRecurring, bool)`

GetRecurringOk returns a tuple with the Recurring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurring

`func (o *ValidatePricingSummary) SetRecurring(v ValidatePricingSummaryRecurring)`

SetRecurring sets Recurring field to given value.

### HasRecurring

`func (o *ValidatePricingSummary) HasRecurring() bool`

HasRecurring returns a boolean if a field has been set.

### GetProrataTotal

`func (o *ValidatePricingSummary) GetProrataTotal() string`

GetProrataTotal returns the ProrataTotal field if non-nil, zero value otherwise.

### GetProrataTotalOk

`func (o *ValidatePricingSummary) GetProrataTotalOk() (*string, bool)`

GetProrataTotalOk returns a tuple with the ProrataTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProrataTotal

`func (o *ValidatePricingSummary) SetProrataTotal(v string)`

SetProrataTotal sets ProrataTotal field to given value.

### HasProrataTotal

`func (o *ValidatePricingSummary) HasProrataTotal() bool`

HasProrataTotal returns a boolean if a field has been set.

### GetPromoApplied

`func (o *ValidatePricingSummary) GetPromoApplied() string`

GetPromoApplied returns the PromoApplied field if non-nil, zero value otherwise.

### GetPromoAppliedOk

`func (o *ValidatePricingSummary) GetPromoAppliedOk() (*string, bool)`

GetPromoAppliedOk returns a tuple with the PromoApplied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromoApplied

`func (o *ValidatePricingSummary) SetPromoApplied(v string)`

SetPromoApplied sets PromoApplied field to given value.

### HasPromoApplied

`func (o *ValidatePricingSummary) HasPromoApplied() bool`

HasPromoApplied returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


