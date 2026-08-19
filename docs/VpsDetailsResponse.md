# VpsDetailsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Status message | 
**VmInfo** | [**VpsVmInfo**](VpsVmInfo.md) |  | 
**Cpu** | [**VpsCpuInfo**](VpsCpuInfo.md) |  | 
**Memory** | [**VpsMemoryInfo**](VpsMemoryInfo.md) |  | 
**Disk** | [**VpsDiskInfo**](VpsDiskInfo.md) |  | 
**NetworkRate** | Pointer to [**VpsNetworkRate**](VpsNetworkRate.md) |  | [optional] 
**IpAddresses** | [**[]VpsIpAddressDetail**](VpsIpAddressDetail.md) | List of IP addresses assigned to the VPS, including subnet, gateway, and MAC | 
**Credentials** | [**VpsCredentials**](VpsCredentials.md) |  | 
**AvailableFeatures** | [**VpsAvailableFeatures**](VpsAvailableFeatures.md) |  | 
**OsInfo** | Pointer to [**VpsOsInfo**](VpsOsInfo.md) |  | [optional] 

## Methods

### NewVpsDetailsResponse

`func NewVpsDetailsResponse(message string, vmInfo VpsVmInfo, cpu VpsCpuInfo, memory VpsMemoryInfo, disk VpsDiskInfo, ipAddresses []VpsIpAddressDetail, credentials VpsCredentials, availableFeatures VpsAvailableFeatures, ) *VpsDetailsResponse`

NewVpsDetailsResponse instantiates a new VpsDetailsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpsDetailsResponseWithDefaults

`func NewVpsDetailsResponseWithDefaults() *VpsDetailsResponse`

NewVpsDetailsResponseWithDefaults instantiates a new VpsDetailsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *VpsDetailsResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *VpsDetailsResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *VpsDetailsResponse) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetVmInfo

`func (o *VpsDetailsResponse) GetVmInfo() VpsVmInfo`

GetVmInfo returns the VmInfo field if non-nil, zero value otherwise.

### GetVmInfoOk

`func (o *VpsDetailsResponse) GetVmInfoOk() (*VpsVmInfo, bool)`

GetVmInfoOk returns a tuple with the VmInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmInfo

`func (o *VpsDetailsResponse) SetVmInfo(v VpsVmInfo)`

SetVmInfo sets VmInfo field to given value.


### GetCpu

`func (o *VpsDetailsResponse) GetCpu() VpsCpuInfo`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *VpsDetailsResponse) GetCpuOk() (*VpsCpuInfo, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *VpsDetailsResponse) SetCpu(v VpsCpuInfo)`

SetCpu sets Cpu field to given value.


### GetMemory

`func (o *VpsDetailsResponse) GetMemory() VpsMemoryInfo`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *VpsDetailsResponse) GetMemoryOk() (*VpsMemoryInfo, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *VpsDetailsResponse) SetMemory(v VpsMemoryInfo)`

SetMemory sets Memory field to given value.


### GetDisk

`func (o *VpsDetailsResponse) GetDisk() VpsDiskInfo`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *VpsDetailsResponse) GetDiskOk() (*VpsDiskInfo, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *VpsDetailsResponse) SetDisk(v VpsDiskInfo)`

SetDisk sets Disk field to given value.


### GetNetworkRate

`func (o *VpsDetailsResponse) GetNetworkRate() VpsNetworkRate`

GetNetworkRate returns the NetworkRate field if non-nil, zero value otherwise.

### GetNetworkRateOk

`func (o *VpsDetailsResponse) GetNetworkRateOk() (*VpsNetworkRate, bool)`

GetNetworkRateOk returns a tuple with the NetworkRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkRate

`func (o *VpsDetailsResponse) SetNetworkRate(v VpsNetworkRate)`

SetNetworkRate sets NetworkRate field to given value.

### HasNetworkRate

`func (o *VpsDetailsResponse) HasNetworkRate() bool`

HasNetworkRate returns a boolean if a field has been set.

### GetIpAddresses

`func (o *VpsDetailsResponse) GetIpAddresses() []VpsIpAddressDetail`

GetIpAddresses returns the IpAddresses field if non-nil, zero value otherwise.

### GetIpAddressesOk

`func (o *VpsDetailsResponse) GetIpAddressesOk() (*[]VpsIpAddressDetail, bool)`

GetIpAddressesOk returns a tuple with the IpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddresses

`func (o *VpsDetailsResponse) SetIpAddresses(v []VpsIpAddressDetail)`

SetIpAddresses sets IpAddresses field to given value.


### GetCredentials

`func (o *VpsDetailsResponse) GetCredentials() VpsCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *VpsDetailsResponse) GetCredentialsOk() (*VpsCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *VpsDetailsResponse) SetCredentials(v VpsCredentials)`

SetCredentials sets Credentials field to given value.


### GetAvailableFeatures

`func (o *VpsDetailsResponse) GetAvailableFeatures() VpsAvailableFeatures`

GetAvailableFeatures returns the AvailableFeatures field if non-nil, zero value otherwise.

### GetAvailableFeaturesOk

`func (o *VpsDetailsResponse) GetAvailableFeaturesOk() (*VpsAvailableFeatures, bool)`

GetAvailableFeaturesOk returns a tuple with the AvailableFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableFeatures

`func (o *VpsDetailsResponse) SetAvailableFeatures(v VpsAvailableFeatures)`

SetAvailableFeatures sets AvailableFeatures field to given value.


### GetOsInfo

`func (o *VpsDetailsResponse) GetOsInfo() VpsOsInfo`

GetOsInfo returns the OsInfo field if non-nil, zero value otherwise.

### GetOsInfoOk

`func (o *VpsDetailsResponse) GetOsInfoOk() (*VpsOsInfo, bool)`

GetOsInfoOk returns a tuple with the OsInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsInfo

`func (o *VpsDetailsResponse) SetOsInfo(v VpsOsInfo)`

SetOsInfo sets OsInfo field to given value.

### HasOsInfo

`func (o *VpsDetailsResponse) HasOsInfo() bool`

HasOsInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


