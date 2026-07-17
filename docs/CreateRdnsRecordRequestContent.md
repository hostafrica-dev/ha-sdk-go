# CreateRdnsRecordRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **int32** | Zone type: 0&#x3D;OTHER, 1&#x3D;DOMAIN, 2&#x3D;HOSTING, 3&#x3D;ADDON | 
**Relid** | **int32** | Related service id matching the zone type (hosting, addon, or domain) | 
**Ip** | **string** | IPv4 or IPv6 address the PTR record should point from | 
**Hostname** | **string** | PTR target hostname | 
**Ttl** | Pointer to **int32** | Time-to-live in seconds. Defaults to 14400; clamped to [30, 86400] | [optional] 

## Methods

### NewCreateRdnsRecordRequestContent

`func NewCreateRdnsRecordRequestContent(type_ int32, relid int32, ip string, hostname string, ) *CreateRdnsRecordRequestContent`

NewCreateRdnsRecordRequestContent instantiates a new CreateRdnsRecordRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRdnsRecordRequestContentWithDefaults

`func NewCreateRdnsRecordRequestContentWithDefaults() *CreateRdnsRecordRequestContent`

NewCreateRdnsRecordRequestContentWithDefaults instantiates a new CreateRdnsRecordRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CreateRdnsRecordRequestContent) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateRdnsRecordRequestContent) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateRdnsRecordRequestContent) SetType(v int32)`

SetType sets Type field to given value.


### GetRelid

`func (o *CreateRdnsRecordRequestContent) GetRelid() int32`

GetRelid returns the Relid field if non-nil, zero value otherwise.

### GetRelidOk

`func (o *CreateRdnsRecordRequestContent) GetRelidOk() (*int32, bool)`

GetRelidOk returns a tuple with the Relid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelid

`func (o *CreateRdnsRecordRequestContent) SetRelid(v int32)`

SetRelid sets Relid field to given value.


### GetIp

`func (o *CreateRdnsRecordRequestContent) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *CreateRdnsRecordRequestContent) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *CreateRdnsRecordRequestContent) SetIp(v string)`

SetIp sets Ip field to given value.


### GetHostname

`func (o *CreateRdnsRecordRequestContent) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *CreateRdnsRecordRequestContent) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *CreateRdnsRecordRequestContent) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetTtl

`func (o *CreateRdnsRecordRequestContent) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *CreateRdnsRecordRequestContent) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *CreateRdnsRecordRequestContent) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *CreateRdnsRecordRequestContent) HasTtl() bool`

HasTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


