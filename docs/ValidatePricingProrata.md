# ValidatePricingProrata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Applied** | **bool** |  | 
**Amount** | Pointer to **string** |  | [optional] 
**Date** | Pointer to **string** |  | [optional] 
**InvoiceDate** | Pointer to **string** |  | [optional] 
**Days** | Pointer to **int32** |  | [optional] 

## Methods

### NewValidatePricingProrata

`func NewValidatePricingProrata(applied bool, ) *ValidatePricingProrata`

NewValidatePricingProrata instantiates a new ValidatePricingProrata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidatePricingProrataWithDefaults

`func NewValidatePricingProrataWithDefaults() *ValidatePricingProrata`

NewValidatePricingProrataWithDefaults instantiates a new ValidatePricingProrata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplied

`func (o *ValidatePricingProrata) GetApplied() bool`

GetApplied returns the Applied field if non-nil, zero value otherwise.

### GetAppliedOk

`func (o *ValidatePricingProrata) GetAppliedOk() (*bool, bool)`

GetAppliedOk returns a tuple with the Applied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplied

`func (o *ValidatePricingProrata) SetApplied(v bool)`

SetApplied sets Applied field to given value.


### GetAmount

`func (o *ValidatePricingProrata) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ValidatePricingProrata) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ValidatePricingProrata) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *ValidatePricingProrata) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetDate

`func (o *ValidatePricingProrata) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *ValidatePricingProrata) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *ValidatePricingProrata) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *ValidatePricingProrata) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetInvoiceDate

`func (o *ValidatePricingProrata) GetInvoiceDate() string`

GetInvoiceDate returns the InvoiceDate field if non-nil, zero value otherwise.

### GetInvoiceDateOk

`func (o *ValidatePricingProrata) GetInvoiceDateOk() (*string, bool)`

GetInvoiceDateOk returns a tuple with the InvoiceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceDate

`func (o *ValidatePricingProrata) SetInvoiceDate(v string)`

SetInvoiceDate sets InvoiceDate field to given value.

### HasInvoiceDate

`func (o *ValidatePricingProrata) HasInvoiceDate() bool`

HasInvoiceDate returns a boolean if a field has been set.

### GetDays

`func (o *ValidatePricingProrata) GetDays() int32`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *ValidatePricingProrata) GetDaysOk() (*int32, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *ValidatePricingProrata) SetDays(v int32)`

SetDays sets Days field to given value.

### HasDays

`func (o *ValidatePricingProrata) HasDays() bool`

HasDays returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


