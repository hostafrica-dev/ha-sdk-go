# CataloguePlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Plan identifier | 
**Name** | **string** | Plan display name (e.g. C1, C2) | 
**Cpu** | **int32** | Number of vCPUs | 
**Ram** | **int32** | RAM in GB | 
**Disk** | **int32** | Disk size in GB | 
**Bandwidth** | **int32** | Bandwidth allowance in GB | 
**Backups** | **int32** | Number of included backup slots | 
**Snapshots** | **int32** | Number of included snapshot slots | 
**Pricing** | **interface{}** | Per-cycle pricing for this plan | 

## Methods

### NewCataloguePlan

`func NewCataloguePlan(id int32, name string, cpu int32, ram int32, disk int32, bandwidth int32, backups int32, snapshots int32, pricing interface{}, ) *CataloguePlan`

NewCataloguePlan instantiates a new CataloguePlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCataloguePlanWithDefaults

`func NewCataloguePlanWithDefaults() *CataloguePlan`

NewCataloguePlanWithDefaults instantiates a new CataloguePlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CataloguePlan) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CataloguePlan) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CataloguePlan) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *CataloguePlan) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CataloguePlan) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CataloguePlan) SetName(v string)`

SetName sets Name field to given value.


### GetCpu

`func (o *CataloguePlan) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *CataloguePlan) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *CataloguePlan) SetCpu(v int32)`

SetCpu sets Cpu field to given value.


### GetRam

`func (o *CataloguePlan) GetRam() int32`

GetRam returns the Ram field if non-nil, zero value otherwise.

### GetRamOk

`func (o *CataloguePlan) GetRamOk() (*int32, bool)`

GetRamOk returns a tuple with the Ram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRam

`func (o *CataloguePlan) SetRam(v int32)`

SetRam sets Ram field to given value.


### GetDisk

`func (o *CataloguePlan) GetDisk() int32`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *CataloguePlan) GetDiskOk() (*int32, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *CataloguePlan) SetDisk(v int32)`

SetDisk sets Disk field to given value.


### GetBandwidth

`func (o *CataloguePlan) GetBandwidth() int32`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *CataloguePlan) GetBandwidthOk() (*int32, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *CataloguePlan) SetBandwidth(v int32)`

SetBandwidth sets Bandwidth field to given value.


### GetBackups

`func (o *CataloguePlan) GetBackups() int32`

GetBackups returns the Backups field if non-nil, zero value otherwise.

### GetBackupsOk

`func (o *CataloguePlan) GetBackupsOk() (*int32, bool)`

GetBackupsOk returns a tuple with the Backups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackups

`func (o *CataloguePlan) SetBackups(v int32)`

SetBackups sets Backups field to given value.


### GetSnapshots

`func (o *CataloguePlan) GetSnapshots() int32`

GetSnapshots returns the Snapshots field if non-nil, zero value otherwise.

### GetSnapshotsOk

`func (o *CataloguePlan) GetSnapshotsOk() (*int32, bool)`

GetSnapshotsOk returns a tuple with the Snapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshots

`func (o *CataloguePlan) SetSnapshots(v int32)`

SetSnapshots sets Snapshots field to given value.


### GetPricing

`func (o *CataloguePlan) GetPricing() interface{}`

GetPricing returns the Pricing field if non-nil, zero value otherwise.

### GetPricingOk

`func (o *CataloguePlan) GetPricingOk() (*interface{}, bool)`

GetPricingOk returns a tuple with the Pricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricing

`func (o *CataloguePlan) SetPricing(v interface{})`

SetPricing sets Pricing field to given value.


### SetPricingNil

`func (o *CataloguePlan) SetPricingNil(b bool)`

 SetPricingNil sets the value for Pricing to be an explicit nil

### UnsetPricing
`func (o *CataloguePlan) UnsetPricing()`

UnsetPricing ensures that no value is present for Pricing, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


