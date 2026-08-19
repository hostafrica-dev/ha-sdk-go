# VpsIpAddressDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | **string** | Primary IP address | 
**Address** | Pointer to **string** | IP address (may mirror ip) | [optional] 
**Subnet** | Pointer to **string** | Subnet mask | [optional] 
**Gateway** | Pointer to **string** | Gateway address | [optional] 
**Mac** | Pointer to **string** | MAC address when available | [optional] 

## Methods

### NewVpsIpAddressDetail

`func NewVpsIpAddressDetail(ip string, ) *VpsIpAddressDetail`

NewVpsIpAddressDetail instantiates a new VpsIpAddressDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpsIpAddressDetailWithDefaults

`func NewVpsIpAddressDetailWithDefaults() *VpsIpAddressDetail`

NewVpsIpAddressDetailWithDefaults instantiates a new VpsIpAddressDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *VpsIpAddressDetail) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *VpsIpAddressDetail) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *VpsIpAddressDetail) SetIp(v string)`

SetIp sets Ip field to given value.


### GetAddress

`func (o *VpsIpAddressDetail) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *VpsIpAddressDetail) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *VpsIpAddressDetail) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *VpsIpAddressDetail) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetSubnet

`func (o *VpsIpAddressDetail) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *VpsIpAddressDetail) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *VpsIpAddressDetail) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *VpsIpAddressDetail) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### GetGateway

`func (o *VpsIpAddressDetail) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *VpsIpAddressDetail) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *VpsIpAddressDetail) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *VpsIpAddressDetail) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetMac

`func (o *VpsIpAddressDetail) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *VpsIpAddressDetail) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *VpsIpAddressDetail) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *VpsIpAddressDetail) HasMac() bool`

HasMac returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


