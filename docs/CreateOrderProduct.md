# CreateOrderProduct

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pid** | **int32** | Product ID | 
**BillingCycle** | [**BillingCycle**](BillingCycle.md) |  | 
**PlanId** | **int32** | Plan ID for the selected product configuration | 
**Hostname** | **string** | Hostname to assign to the service | 
**ConfigOptions** | **interface{}** | Configuration options as a free-form map | 
**AdditionalData** | Pointer to **interface{}** | Additional data as a free-form map | [optional] 

## Methods

### NewCreateOrderProduct

`func NewCreateOrderProduct(pid int32, billingCycle BillingCycle, planId int32, hostname string, configOptions interface{}, ) *CreateOrderProduct`

NewCreateOrderProduct instantiates a new CreateOrderProduct object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrderProductWithDefaults

`func NewCreateOrderProductWithDefaults() *CreateOrderProduct`

NewCreateOrderProductWithDefaults instantiates a new CreateOrderProduct object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPid

`func (o *CreateOrderProduct) GetPid() int32`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *CreateOrderProduct) GetPidOk() (*int32, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *CreateOrderProduct) SetPid(v int32)`

SetPid sets Pid field to given value.


### GetBillingCycle

`func (o *CreateOrderProduct) GetBillingCycle() BillingCycle`

GetBillingCycle returns the BillingCycle field if non-nil, zero value otherwise.

### GetBillingCycleOk

`func (o *CreateOrderProduct) GetBillingCycleOk() (*BillingCycle, bool)`

GetBillingCycleOk returns a tuple with the BillingCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCycle

`func (o *CreateOrderProduct) SetBillingCycle(v BillingCycle)`

SetBillingCycle sets BillingCycle field to given value.


### GetPlanId

`func (o *CreateOrderProduct) GetPlanId() int32`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *CreateOrderProduct) GetPlanIdOk() (*int32, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *CreateOrderProduct) SetPlanId(v int32)`

SetPlanId sets PlanId field to given value.


### GetHostname

`func (o *CreateOrderProduct) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *CreateOrderProduct) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *CreateOrderProduct) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetConfigOptions

`func (o *CreateOrderProduct) GetConfigOptions() interface{}`

GetConfigOptions returns the ConfigOptions field if non-nil, zero value otherwise.

### GetConfigOptionsOk

`func (o *CreateOrderProduct) GetConfigOptionsOk() (*interface{}, bool)`

GetConfigOptionsOk returns a tuple with the ConfigOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigOptions

`func (o *CreateOrderProduct) SetConfigOptions(v interface{})`

SetConfigOptions sets ConfigOptions field to given value.


### SetConfigOptionsNil

`func (o *CreateOrderProduct) SetConfigOptionsNil(b bool)`

 SetConfigOptionsNil sets the value for ConfigOptions to be an explicit nil

### UnsetConfigOptions
`func (o *CreateOrderProduct) UnsetConfigOptions()`

UnsetConfigOptions ensures that no value is present for ConfigOptions, not even an explicit nil
### GetAdditionalData

`func (o *CreateOrderProduct) GetAdditionalData() interface{}`

GetAdditionalData returns the AdditionalData field if non-nil, zero value otherwise.

### GetAdditionalDataOk

`func (o *CreateOrderProduct) GetAdditionalDataOk() (*interface{}, bool)`

GetAdditionalDataOk returns a tuple with the AdditionalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalData

`func (o *CreateOrderProduct) SetAdditionalData(v interface{})`

SetAdditionalData sets AdditionalData field to given value.

### HasAdditionalData

`func (o *CreateOrderProduct) HasAdditionalData() bool`

HasAdditionalData returns a boolean if a field has been set.

### SetAdditionalDataNil

`func (o *CreateOrderProduct) SetAdditionalDataNil(b bool)`

 SetAdditionalDataNil sets the value for AdditionalData to be an explicit nil

### UnsetAdditionalData
`func (o *CreateOrderProduct) UnsetAdditionalData()`

UnsetAdditionalData ensures that no value is present for AdditionalData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


