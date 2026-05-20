# VpsBandwidthInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UsageMb** | Pointer to **float64** | Bandwidth usage in MB | [optional] 
**LimitMb** | Pointer to **float64** | Bandwidth limit in MB | [optional] 
**Unlimited** | **bool** | Whether bandwidth is unlimited | 
**Percent** | Pointer to **float64** | Bandwidth usage percentage | [optional] 

## Methods

### NewVpsBandwidthInfo

`func NewVpsBandwidthInfo(unlimited bool, ) *VpsBandwidthInfo`

NewVpsBandwidthInfo instantiates a new VpsBandwidthInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpsBandwidthInfoWithDefaults

`func NewVpsBandwidthInfoWithDefaults() *VpsBandwidthInfo`

NewVpsBandwidthInfoWithDefaults instantiates a new VpsBandwidthInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsageMb

`func (o *VpsBandwidthInfo) GetUsageMb() float64`

GetUsageMb returns the UsageMb field if non-nil, zero value otherwise.

### GetUsageMbOk

`func (o *VpsBandwidthInfo) GetUsageMbOk() (*float64, bool)`

GetUsageMbOk returns a tuple with the UsageMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageMb

`func (o *VpsBandwidthInfo) SetUsageMb(v float64)`

SetUsageMb sets UsageMb field to given value.

### HasUsageMb

`func (o *VpsBandwidthInfo) HasUsageMb() bool`

HasUsageMb returns a boolean if a field has been set.

### GetLimitMb

`func (o *VpsBandwidthInfo) GetLimitMb() float64`

GetLimitMb returns the LimitMb field if non-nil, zero value otherwise.

### GetLimitMbOk

`func (o *VpsBandwidthInfo) GetLimitMbOk() (*float64, bool)`

GetLimitMbOk returns a tuple with the LimitMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitMb

`func (o *VpsBandwidthInfo) SetLimitMb(v float64)`

SetLimitMb sets LimitMb field to given value.

### HasLimitMb

`func (o *VpsBandwidthInfo) HasLimitMb() bool`

HasLimitMb returns a boolean if a field has been set.

### GetUnlimited

`func (o *VpsBandwidthInfo) GetUnlimited() bool`

GetUnlimited returns the Unlimited field if non-nil, zero value otherwise.

### GetUnlimitedOk

`func (o *VpsBandwidthInfo) GetUnlimitedOk() (*bool, bool)`

GetUnlimitedOk returns a tuple with the Unlimited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnlimited

`func (o *VpsBandwidthInfo) SetUnlimited(v bool)`

SetUnlimited sets Unlimited field to given value.


### GetPercent

`func (o *VpsBandwidthInfo) GetPercent() float64`

GetPercent returns the Percent field if non-nil, zero value otherwise.

### GetPercentOk

`func (o *VpsBandwidthInfo) GetPercentOk() (*float64, bool)`

GetPercentOk returns a tuple with the Percent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercent

`func (o *VpsBandwidthInfo) SetPercent(v float64)`

SetPercent sets Percent field to given value.

### HasPercent

`func (o *VpsBandwidthInfo) HasPercent() bool`

HasPercent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


