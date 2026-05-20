# VpsMemoryInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UsageBytes** | Pointer to **int64** | Memory usage in bytes | [optional] 
**TotalBytes** | Pointer to **int64** | Total memory in bytes | [optional] 
**Percent** | Pointer to **float64** | Memory usage percentage | [optional] 
**UsageHuman** | Pointer to **string** | Human-readable memory usage (e.g., &#39;2.5 GB&#39;) | [optional] 
**TotalHuman** | Pointer to **string** | Human-readable total memory (e.g., &#39;8 GB&#39;) | [optional] 

## Methods

### NewVpsMemoryInfo

`func NewVpsMemoryInfo() *VpsMemoryInfo`

NewVpsMemoryInfo instantiates a new VpsMemoryInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpsMemoryInfoWithDefaults

`func NewVpsMemoryInfoWithDefaults() *VpsMemoryInfo`

NewVpsMemoryInfoWithDefaults instantiates a new VpsMemoryInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsageBytes

`func (o *VpsMemoryInfo) GetUsageBytes() int64`

GetUsageBytes returns the UsageBytes field if non-nil, zero value otherwise.

### GetUsageBytesOk

`func (o *VpsMemoryInfo) GetUsageBytesOk() (*int64, bool)`

GetUsageBytesOk returns a tuple with the UsageBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageBytes

`func (o *VpsMemoryInfo) SetUsageBytes(v int64)`

SetUsageBytes sets UsageBytes field to given value.

### HasUsageBytes

`func (o *VpsMemoryInfo) HasUsageBytes() bool`

HasUsageBytes returns a boolean if a field has been set.

### GetTotalBytes

`func (o *VpsMemoryInfo) GetTotalBytes() int64`

GetTotalBytes returns the TotalBytes field if non-nil, zero value otherwise.

### GetTotalBytesOk

`func (o *VpsMemoryInfo) GetTotalBytesOk() (*int64, bool)`

GetTotalBytesOk returns a tuple with the TotalBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBytes

`func (o *VpsMemoryInfo) SetTotalBytes(v int64)`

SetTotalBytes sets TotalBytes field to given value.

### HasTotalBytes

`func (o *VpsMemoryInfo) HasTotalBytes() bool`

HasTotalBytes returns a boolean if a field has been set.

### GetPercent

`func (o *VpsMemoryInfo) GetPercent() float64`

GetPercent returns the Percent field if non-nil, zero value otherwise.

### GetPercentOk

`func (o *VpsMemoryInfo) GetPercentOk() (*float64, bool)`

GetPercentOk returns a tuple with the Percent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercent

`func (o *VpsMemoryInfo) SetPercent(v float64)`

SetPercent sets Percent field to given value.

### HasPercent

`func (o *VpsMemoryInfo) HasPercent() bool`

HasPercent returns a boolean if a field has been set.

### GetUsageHuman

`func (o *VpsMemoryInfo) GetUsageHuman() string`

GetUsageHuman returns the UsageHuman field if non-nil, zero value otherwise.

### GetUsageHumanOk

`func (o *VpsMemoryInfo) GetUsageHumanOk() (*string, bool)`

GetUsageHumanOk returns a tuple with the UsageHuman field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageHuman

`func (o *VpsMemoryInfo) SetUsageHuman(v string)`

SetUsageHuman sets UsageHuman field to given value.

### HasUsageHuman

`func (o *VpsMemoryInfo) HasUsageHuman() bool`

HasUsageHuman returns a boolean if a field has been set.

### GetTotalHuman

`func (o *VpsMemoryInfo) GetTotalHuman() string`

GetTotalHuman returns the TotalHuman field if non-nil, zero value otherwise.

### GetTotalHumanOk

`func (o *VpsMemoryInfo) GetTotalHumanOk() (*string, bool)`

GetTotalHumanOk returns a tuple with the TotalHuman field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalHuman

`func (o *VpsMemoryInfo) SetTotalHuman(v string)`

SetTotalHuman sets TotalHuman field to given value.

### HasTotalHuman

`func (o *VpsMemoryInfo) HasTotalHuman() bool`

HasTotalHuman returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


