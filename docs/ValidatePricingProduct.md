# ValidatePricingProduct

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pid** | **int32** | WHMCS product ID | 
**BillingCycle** | [**BillingCycle**](BillingCycle.md) |  | 
**PlanId** | **int32** | Plan ID for the selected product configuration | 
**Hostname** | Pointer to **string** | Hostname to assign to the service | [optional] 
**ConfigOptions** | **interface{}** | Configuration options as a free-form map (option_id -&gt; value) | 

## Methods

### NewValidatePricingProduct

`func NewValidatePricingProduct(pid int32, billingCycle BillingCycle, planId int32, configOptions interface{}, ) *ValidatePricingProduct`

NewValidatePricingProduct instantiates a new ValidatePricingProduct object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidatePricingProductWithDefaults

`func NewValidatePricingProductWithDefaults() *ValidatePricingProduct`

NewValidatePricingProductWithDefaults instantiates a new ValidatePricingProduct object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPid

`func (o *ValidatePricingProduct) GetPid() int32`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *ValidatePricingProduct) GetPidOk() (*int32, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *ValidatePricingProduct) SetPid(v int32)`

SetPid sets Pid field to given value.


### GetBillingCycle

`func (o *ValidatePricingProduct) GetBillingCycle() BillingCycle`

GetBillingCycle returns the BillingCycle field if non-nil, zero value otherwise.

### GetBillingCycleOk

`func (o *ValidatePricingProduct) GetBillingCycleOk() (*BillingCycle, bool)`

GetBillingCycleOk returns a tuple with the BillingCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCycle

`func (o *ValidatePricingProduct) SetBillingCycle(v BillingCycle)`

SetBillingCycle sets BillingCycle field to given value.


### GetPlanId

`func (o *ValidatePricingProduct) GetPlanId() int32`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *ValidatePricingProduct) GetPlanIdOk() (*int32, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *ValidatePricingProduct) SetPlanId(v int32)`

SetPlanId sets PlanId field to given value.


### GetHostname

`func (o *ValidatePricingProduct) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *ValidatePricingProduct) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *ValidatePricingProduct) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *ValidatePricingProduct) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetConfigOptions

`func (o *ValidatePricingProduct) GetConfigOptions() interface{}`

GetConfigOptions returns the ConfigOptions field if non-nil, zero value otherwise.

### GetConfigOptionsOk

`func (o *ValidatePricingProduct) GetConfigOptionsOk() (*interface{}, bool)`

GetConfigOptionsOk returns a tuple with the ConfigOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigOptions

`func (o *ValidatePricingProduct) SetConfigOptions(v interface{})`

SetConfigOptions sets ConfigOptions field to given value.


### SetConfigOptionsNil

`func (o *ValidatePricingProduct) SetConfigOptionsNil(b bool)`

 SetConfigOptionsNil sets the value for ConfigOptions to be an explicit nil

### UnsetConfigOptions
`func (o *ValidatePricingProduct) UnsetConfigOptions()`

UnsetConfigOptions ensures that no value is present for ConfigOptions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


