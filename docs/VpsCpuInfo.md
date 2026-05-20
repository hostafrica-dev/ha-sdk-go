# VpsCpuInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UsageRatio** | Pointer to **float64** | CPU usage ratio (0.0 to 1.0) | [optional] 
**Percent** | Pointer to **float64** | CPU usage percentage | [optional] 
**Cores** | Pointer to **int32** | Number of CPU cores | [optional] 

## Methods

### NewVpsCpuInfo

`func NewVpsCpuInfo() *VpsCpuInfo`

NewVpsCpuInfo instantiates a new VpsCpuInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpsCpuInfoWithDefaults

`func NewVpsCpuInfoWithDefaults() *VpsCpuInfo`

NewVpsCpuInfoWithDefaults instantiates a new VpsCpuInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsageRatio

`func (o *VpsCpuInfo) GetUsageRatio() float64`

GetUsageRatio returns the UsageRatio field if non-nil, zero value otherwise.

### GetUsageRatioOk

`func (o *VpsCpuInfo) GetUsageRatioOk() (*float64, bool)`

GetUsageRatioOk returns a tuple with the UsageRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageRatio

`func (o *VpsCpuInfo) SetUsageRatio(v float64)`

SetUsageRatio sets UsageRatio field to given value.

### HasUsageRatio

`func (o *VpsCpuInfo) HasUsageRatio() bool`

HasUsageRatio returns a boolean if a field has been set.

### GetPercent

`func (o *VpsCpuInfo) GetPercent() float64`

GetPercent returns the Percent field if non-nil, zero value otherwise.

### GetPercentOk

`func (o *VpsCpuInfo) GetPercentOk() (*float64, bool)`

GetPercentOk returns a tuple with the Percent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercent

`func (o *VpsCpuInfo) SetPercent(v float64)`

SetPercent sets Percent field to given value.

### HasPercent

`func (o *VpsCpuInfo) HasPercent() bool`

HasPercent returns a boolean if a field has been set.

### GetCores

`func (o *VpsCpuInfo) GetCores() int32`

GetCores returns the Cores field if non-nil, zero value otherwise.

### GetCoresOk

`func (o *VpsCpuInfo) GetCoresOk() (*int32, bool)`

GetCoresOk returns a tuple with the Cores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCores

`func (o *VpsCpuInfo) SetCores(v int32)`

SetCores sets Cores field to given value.

### HasCores

`func (o *VpsCpuInfo) HasCores() bool`

HasCores returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


