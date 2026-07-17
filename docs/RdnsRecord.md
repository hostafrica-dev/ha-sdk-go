# RdnsRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Record identifier — use as record_id in PUT/DELETE | 
**Ip** | **string** | The IPv4/IPv6 address the PTR points from | 
**Hostname** | **string** | Final concatenated PTR target (sub.from) | 
**Ttl** | Pointer to **int32** | Time-to-live in seconds | [optional] 
**Type** | Pointer to **int32** | Zone type: 0&#x3D;OTHER, 1&#x3D;DOMAIN, 2&#x3D;HOSTING, 3&#x3D;ADDON | [optional] 
**Relid** | Pointer to **int32** | Related service id for the zone type (hosting, addon, or domain) | [optional] 
**Serverid** | **int32** |  | 
**Clientid** | **int32** |  | 
**Packageid** | **int32** |  | 
**From** | Pointer to **string** | Reverse-DNS zone name (e.g. 10.113.0.203.in-addr.arpa) | [optional] 
**Sub** | Pointer to **string** | Sub-label component of the PTR | [optional] 
**Name** | Pointer to **string** | Full PTR name (e.g. 10.113.0.203.in-addr.arpa) | [optional] 

## Methods

### NewRdnsRecord

`func NewRdnsRecord(id int32, ip string, hostname string, serverid int32, clientid int32, packageid int32, ) *RdnsRecord`

NewRdnsRecord instantiates a new RdnsRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRdnsRecordWithDefaults

`func NewRdnsRecordWithDefaults() *RdnsRecord`

NewRdnsRecordWithDefaults instantiates a new RdnsRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RdnsRecord) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RdnsRecord) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RdnsRecord) SetId(v int32)`

SetId sets Id field to given value.


### GetIp

`func (o *RdnsRecord) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *RdnsRecord) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *RdnsRecord) SetIp(v string)`

SetIp sets Ip field to given value.


### GetHostname

`func (o *RdnsRecord) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *RdnsRecord) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *RdnsRecord) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetTtl

`func (o *RdnsRecord) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *RdnsRecord) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *RdnsRecord) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *RdnsRecord) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetType

`func (o *RdnsRecord) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RdnsRecord) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RdnsRecord) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *RdnsRecord) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRelid

`func (o *RdnsRecord) GetRelid() int32`

GetRelid returns the Relid field if non-nil, zero value otherwise.

### GetRelidOk

`func (o *RdnsRecord) GetRelidOk() (*int32, bool)`

GetRelidOk returns a tuple with the Relid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelid

`func (o *RdnsRecord) SetRelid(v int32)`

SetRelid sets Relid field to given value.

### HasRelid

`func (o *RdnsRecord) HasRelid() bool`

HasRelid returns a boolean if a field has been set.

### GetServerid

`func (o *RdnsRecord) GetServerid() int32`

GetServerid returns the Serverid field if non-nil, zero value otherwise.

### GetServeridOk

`func (o *RdnsRecord) GetServeridOk() (*int32, bool)`

GetServeridOk returns a tuple with the Serverid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerid

`func (o *RdnsRecord) SetServerid(v int32)`

SetServerid sets Serverid field to given value.


### GetClientid

`func (o *RdnsRecord) GetClientid() int32`

GetClientid returns the Clientid field if non-nil, zero value otherwise.

### GetClientidOk

`func (o *RdnsRecord) GetClientidOk() (*int32, bool)`

GetClientidOk returns a tuple with the Clientid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientid

`func (o *RdnsRecord) SetClientid(v int32)`

SetClientid sets Clientid field to given value.


### GetPackageid

`func (o *RdnsRecord) GetPackageid() int32`

GetPackageid returns the Packageid field if non-nil, zero value otherwise.

### GetPackageidOk

`func (o *RdnsRecord) GetPackageidOk() (*int32, bool)`

GetPackageidOk returns a tuple with the Packageid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageid

`func (o *RdnsRecord) SetPackageid(v int32)`

SetPackageid sets Packageid field to given value.


### GetFrom

`func (o *RdnsRecord) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *RdnsRecord) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *RdnsRecord) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *RdnsRecord) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetSub

`func (o *RdnsRecord) GetSub() string`

GetSub returns the Sub field if non-nil, zero value otherwise.

### GetSubOk

`func (o *RdnsRecord) GetSubOk() (*string, bool)`

GetSubOk returns a tuple with the Sub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSub

`func (o *RdnsRecord) SetSub(v string)`

SetSub sets Sub field to given value.

### HasSub

`func (o *RdnsRecord) HasSub() bool`

HasSub returns a boolean if a field has been set.

### GetName

`func (o *RdnsRecord) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RdnsRecord) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RdnsRecord) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RdnsRecord) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


