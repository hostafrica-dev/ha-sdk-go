# VpsVmInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | VM status (e.g., running, stopped) | 
**Uptime** | Pointer to **string** | Uptime in human-readable format | [optional] 
**UptimeSeconds** | Pointer to **int64** | Uptime in seconds | [optional] 
**Hostname** | Pointer to **string** | Hostname of the VM | [optional] 
**BootDevices** | Pointer to **[]string** | Boot devices configuration (e.g., scsi0, scsi1) | [optional] 
**Vmid** | **string** | Proxmox VM ID | 
**Node** | **string** | Proxmox node name | 
**Virtualization** | **string** | Virtualization type (qemu or lxc) | 

## Methods

### NewVpsVmInfo

`func NewVpsVmInfo(status string, vmid string, node string, virtualization string, ) *VpsVmInfo`

NewVpsVmInfo instantiates a new VpsVmInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpsVmInfoWithDefaults

`func NewVpsVmInfoWithDefaults() *VpsVmInfo`

NewVpsVmInfoWithDefaults instantiates a new VpsVmInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *VpsVmInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VpsVmInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VpsVmInfo) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetUptime

`func (o *VpsVmInfo) GetUptime() string`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *VpsVmInfo) GetUptimeOk() (*string, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *VpsVmInfo) SetUptime(v string)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *VpsVmInfo) HasUptime() bool`

HasUptime returns a boolean if a field has been set.

### GetUptimeSeconds

`func (o *VpsVmInfo) GetUptimeSeconds() int64`

GetUptimeSeconds returns the UptimeSeconds field if non-nil, zero value otherwise.

### GetUptimeSecondsOk

`func (o *VpsVmInfo) GetUptimeSecondsOk() (*int64, bool)`

GetUptimeSecondsOk returns a tuple with the UptimeSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptimeSeconds

`func (o *VpsVmInfo) SetUptimeSeconds(v int64)`

SetUptimeSeconds sets UptimeSeconds field to given value.

### HasUptimeSeconds

`func (o *VpsVmInfo) HasUptimeSeconds() bool`

HasUptimeSeconds returns a boolean if a field has been set.

### GetHostname

`func (o *VpsVmInfo) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *VpsVmInfo) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *VpsVmInfo) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *VpsVmInfo) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetBootDevices

`func (o *VpsVmInfo) GetBootDevices() []string`

GetBootDevices returns the BootDevices field if non-nil, zero value otherwise.

### GetBootDevicesOk

`func (o *VpsVmInfo) GetBootDevicesOk() (*[]string, bool)`

GetBootDevicesOk returns a tuple with the BootDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootDevices

`func (o *VpsVmInfo) SetBootDevices(v []string)`

SetBootDevices sets BootDevices field to given value.

### HasBootDevices

`func (o *VpsVmInfo) HasBootDevices() bool`

HasBootDevices returns a boolean if a field has been set.

### GetVmid

`func (o *VpsVmInfo) GetVmid() string`

GetVmid returns the Vmid field if non-nil, zero value otherwise.

### GetVmidOk

`func (o *VpsVmInfo) GetVmidOk() (*string, bool)`

GetVmidOk returns a tuple with the Vmid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmid

`func (o *VpsVmInfo) SetVmid(v string)`

SetVmid sets Vmid field to given value.


### GetNode

`func (o *VpsVmInfo) GetNode() string`

GetNode returns the Node field if non-nil, zero value otherwise.

### GetNodeOk

`func (o *VpsVmInfo) GetNodeOk() (*string, bool)`

GetNodeOk returns a tuple with the Node field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNode

`func (o *VpsVmInfo) SetNode(v string)`

SetNode sets Node field to given value.


### GetVirtualization

`func (o *VpsVmInfo) GetVirtualization() string`

GetVirtualization returns the Virtualization field if non-nil, zero value otherwise.

### GetVirtualizationOk

`func (o *VpsVmInfo) GetVirtualizationOk() (*string, bool)`

GetVirtualizationOk returns a tuple with the Virtualization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualization

`func (o *VpsVmInfo) SetVirtualization(v string)`

SetVirtualization sets Virtualization field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


