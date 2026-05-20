# ValidatePricingConfigItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OptionId** | Pointer to **int32** | Option ID (absent for plan-type items) | [optional] 
**Name** | **string** |  | 
**Type** | **string** |  | 
**Selected** | **interface{}** |  | 
**SelectedName** | Pointer to **string** |  | [optional] 
**PlanConfig** | Pointer to [**ValidatePricingPlanConfig**](ValidatePricingPlanConfig.md) |  | [optional] 
**Pricing** | Pointer to [**ValidatePricingPlanPricing**](ValidatePricingPlanPricing.md) |  | [optional] 
**Price** | Pointer to **string** |  | [optional] 
**Setup** | Pointer to **string** |  | [optional] 

## Methods

### NewValidatePricingConfigItem

`func NewValidatePricingConfigItem(name string, type_ string, selected interface{}, ) *ValidatePricingConfigItem`

NewValidatePricingConfigItem instantiates a new ValidatePricingConfigItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidatePricingConfigItemWithDefaults

`func NewValidatePricingConfigItemWithDefaults() *ValidatePricingConfigItem`

NewValidatePricingConfigItemWithDefaults instantiates a new ValidatePricingConfigItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOptionId

`func (o *ValidatePricingConfigItem) GetOptionId() int32`

GetOptionId returns the OptionId field if non-nil, zero value otherwise.

### GetOptionIdOk

`func (o *ValidatePricingConfigItem) GetOptionIdOk() (*int32, bool)`

GetOptionIdOk returns a tuple with the OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionId

`func (o *ValidatePricingConfigItem) SetOptionId(v int32)`

SetOptionId sets OptionId field to given value.

### HasOptionId

`func (o *ValidatePricingConfigItem) HasOptionId() bool`

HasOptionId returns a boolean if a field has been set.

### GetName

`func (o *ValidatePricingConfigItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ValidatePricingConfigItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ValidatePricingConfigItem) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *ValidatePricingConfigItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ValidatePricingConfigItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ValidatePricingConfigItem) SetType(v string)`

SetType sets Type field to given value.


### GetSelected

`func (o *ValidatePricingConfigItem) GetSelected() interface{}`

GetSelected returns the Selected field if non-nil, zero value otherwise.

### GetSelectedOk

`func (o *ValidatePricingConfigItem) GetSelectedOk() (*interface{}, bool)`

GetSelectedOk returns a tuple with the Selected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelected

`func (o *ValidatePricingConfigItem) SetSelected(v interface{})`

SetSelected sets Selected field to given value.


### SetSelectedNil

`func (o *ValidatePricingConfigItem) SetSelectedNil(b bool)`

 SetSelectedNil sets the value for Selected to be an explicit nil

### UnsetSelected
`func (o *ValidatePricingConfigItem) UnsetSelected()`

UnsetSelected ensures that no value is present for Selected, not even an explicit nil
### GetSelectedName

`func (o *ValidatePricingConfigItem) GetSelectedName() string`

GetSelectedName returns the SelectedName field if non-nil, zero value otherwise.

### GetSelectedNameOk

`func (o *ValidatePricingConfigItem) GetSelectedNameOk() (*string, bool)`

GetSelectedNameOk returns a tuple with the SelectedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedName

`func (o *ValidatePricingConfigItem) SetSelectedName(v string)`

SetSelectedName sets SelectedName field to given value.

### HasSelectedName

`func (o *ValidatePricingConfigItem) HasSelectedName() bool`

HasSelectedName returns a boolean if a field has been set.

### GetPlanConfig

`func (o *ValidatePricingConfigItem) GetPlanConfig() ValidatePricingPlanConfig`

GetPlanConfig returns the PlanConfig field if non-nil, zero value otherwise.

### GetPlanConfigOk

`func (o *ValidatePricingConfigItem) GetPlanConfigOk() (*ValidatePricingPlanConfig, bool)`

GetPlanConfigOk returns a tuple with the PlanConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanConfig

`func (o *ValidatePricingConfigItem) SetPlanConfig(v ValidatePricingPlanConfig)`

SetPlanConfig sets PlanConfig field to given value.

### HasPlanConfig

`func (o *ValidatePricingConfigItem) HasPlanConfig() bool`

HasPlanConfig returns a boolean if a field has been set.

### GetPricing

`func (o *ValidatePricingConfigItem) GetPricing() ValidatePricingPlanPricing`

GetPricing returns the Pricing field if non-nil, zero value otherwise.

### GetPricingOk

`func (o *ValidatePricingConfigItem) GetPricingOk() (*ValidatePricingPlanPricing, bool)`

GetPricingOk returns a tuple with the Pricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricing

`func (o *ValidatePricingConfigItem) SetPricing(v ValidatePricingPlanPricing)`

SetPricing sets Pricing field to given value.

### HasPricing

`func (o *ValidatePricingConfigItem) HasPricing() bool`

HasPricing returns a boolean if a field has been set.

### GetPrice

`func (o *ValidatePricingConfigItem) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ValidatePricingConfigItem) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ValidatePricingConfigItem) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *ValidatePricingConfigItem) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetSetup

`func (o *ValidatePricingConfigItem) GetSetup() string`

GetSetup returns the Setup field if non-nil, zero value otherwise.

### GetSetupOk

`func (o *ValidatePricingConfigItem) GetSetupOk() (*string, bool)`

GetSetupOk returns a tuple with the Setup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetup

`func (o *ValidatePricingConfigItem) SetSetup(v string)`

SetSetup sets Setup field to given value.

### HasSetup

`func (o *ValidatePricingConfigItem) HasSetup() bool`

HasSetup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


