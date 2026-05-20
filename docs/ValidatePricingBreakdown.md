# ValidatePricingBreakdown

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Base** | [**ValidatePricingBase**](ValidatePricingBase.md) |  | 
**Configuration** | [**[]ValidatePricingConfigItem**](ValidatePricingConfigItem.md) |  | 
**SetupTotal** | **string** |  | 

## Methods

### NewValidatePricingBreakdown

`func NewValidatePricingBreakdown(base ValidatePricingBase, configuration []ValidatePricingConfigItem, setupTotal string, ) *ValidatePricingBreakdown`

NewValidatePricingBreakdown instantiates a new ValidatePricingBreakdown object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidatePricingBreakdownWithDefaults

`func NewValidatePricingBreakdownWithDefaults() *ValidatePricingBreakdown`

NewValidatePricingBreakdownWithDefaults instantiates a new ValidatePricingBreakdown object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBase

`func (o *ValidatePricingBreakdown) GetBase() ValidatePricingBase`

GetBase returns the Base field if non-nil, zero value otherwise.

### GetBaseOk

`func (o *ValidatePricingBreakdown) GetBaseOk() (*ValidatePricingBase, bool)`

GetBaseOk returns a tuple with the Base field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase

`func (o *ValidatePricingBreakdown) SetBase(v ValidatePricingBase)`

SetBase sets Base field to given value.


### GetConfiguration

`func (o *ValidatePricingBreakdown) GetConfiguration() []ValidatePricingConfigItem`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *ValidatePricingBreakdown) GetConfigurationOk() (*[]ValidatePricingConfigItem, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *ValidatePricingBreakdown) SetConfiguration(v []ValidatePricingConfigItem)`

SetConfiguration sets Configuration field to given value.


### GetSetupTotal

`func (o *ValidatePricingBreakdown) GetSetupTotal() string`

GetSetupTotal returns the SetupTotal field if non-nil, zero value otherwise.

### GetSetupTotalOk

`func (o *ValidatePricingBreakdown) GetSetupTotalOk() (*string, bool)`

GetSetupTotalOk returns a tuple with the SetupTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetupTotal

`func (o *ValidatePricingBreakdown) SetSetupTotal(v string)`

SetSetupTotal sets SetupTotal field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


