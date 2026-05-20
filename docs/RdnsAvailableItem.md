# RdnsAvailableItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **int32** | Zone type matching RdnsRecord.type | 
**Relid** | **int32** | Related service identifier | 
**Name** | **string** | Display label for the service | 
**Ips** | **[]string** | Concrete IP addresses available for this service | 
**Pools** | [**[]RdnsPool**](RdnsPool.md) | Subnet pools available when subnet_custom_ip_mode is on | 
**PtrLimit** | **int32** | Per-package PTR limit (-1 &#x3D; unlimited) | 
**ServerId** | **int32** |  | 
**AllowRdns** | **bool** | Server has ALLOW_RDNS&#x3D;on | 
**RdnsSupported** | **bool** | Server module supports rDNS | 

## Methods

### NewRdnsAvailableItem

`func NewRdnsAvailableItem(type_ int32, relid int32, name string, ips []string, pools []RdnsPool, ptrLimit int32, serverId int32, allowRdns bool, rdnsSupported bool, ) *RdnsAvailableItem`

NewRdnsAvailableItem instantiates a new RdnsAvailableItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRdnsAvailableItemWithDefaults

`func NewRdnsAvailableItemWithDefaults() *RdnsAvailableItem`

NewRdnsAvailableItemWithDefaults instantiates a new RdnsAvailableItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RdnsAvailableItem) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RdnsAvailableItem) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RdnsAvailableItem) SetType(v int32)`

SetType sets Type field to given value.


### GetRelid

`func (o *RdnsAvailableItem) GetRelid() int32`

GetRelid returns the Relid field if non-nil, zero value otherwise.

### GetRelidOk

`func (o *RdnsAvailableItem) GetRelidOk() (*int32, bool)`

GetRelidOk returns a tuple with the Relid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelid

`func (o *RdnsAvailableItem) SetRelid(v int32)`

SetRelid sets Relid field to given value.


### GetName

`func (o *RdnsAvailableItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RdnsAvailableItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RdnsAvailableItem) SetName(v string)`

SetName sets Name field to given value.


### GetIps

`func (o *RdnsAvailableItem) GetIps() []string`

GetIps returns the Ips field if non-nil, zero value otherwise.

### GetIpsOk

`func (o *RdnsAvailableItem) GetIpsOk() (*[]string, bool)`

GetIpsOk returns a tuple with the Ips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIps

`func (o *RdnsAvailableItem) SetIps(v []string)`

SetIps sets Ips field to given value.


### GetPools

`func (o *RdnsAvailableItem) GetPools() []RdnsPool`

GetPools returns the Pools field if non-nil, zero value otherwise.

### GetPoolsOk

`func (o *RdnsAvailableItem) GetPoolsOk() (*[]RdnsPool, bool)`

GetPoolsOk returns a tuple with the Pools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPools

`func (o *RdnsAvailableItem) SetPools(v []RdnsPool)`

SetPools sets Pools field to given value.


### GetPtrLimit

`func (o *RdnsAvailableItem) GetPtrLimit() int32`

GetPtrLimit returns the PtrLimit field if non-nil, zero value otherwise.

### GetPtrLimitOk

`func (o *RdnsAvailableItem) GetPtrLimitOk() (*int32, bool)`

GetPtrLimitOk returns a tuple with the PtrLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPtrLimit

`func (o *RdnsAvailableItem) SetPtrLimit(v int32)`

SetPtrLimit sets PtrLimit field to given value.


### GetServerId

`func (o *RdnsAvailableItem) GetServerId() int32`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *RdnsAvailableItem) GetServerIdOk() (*int32, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *RdnsAvailableItem) SetServerId(v int32)`

SetServerId sets ServerId field to given value.


### GetAllowRdns

`func (o *RdnsAvailableItem) GetAllowRdns() bool`

GetAllowRdns returns the AllowRdns field if non-nil, zero value otherwise.

### GetAllowRdnsOk

`func (o *RdnsAvailableItem) GetAllowRdnsOk() (*bool, bool)`

GetAllowRdnsOk returns a tuple with the AllowRdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRdns

`func (o *RdnsAvailableItem) SetAllowRdns(v bool)`

SetAllowRdns sets AllowRdns field to given value.


### GetRdnsSupported

`func (o *RdnsAvailableItem) GetRdnsSupported() bool`

GetRdnsSupported returns the RdnsSupported field if non-nil, zero value otherwise.

### GetRdnsSupportedOk

`func (o *RdnsAvailableItem) GetRdnsSupportedOk() (*bool, bool)`

GetRdnsSupportedOk returns a tuple with the RdnsSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdnsSupported

`func (o *RdnsAvailableItem) SetRdnsSupported(v bool)`

SetRdnsSupported sets RdnsSupported field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


