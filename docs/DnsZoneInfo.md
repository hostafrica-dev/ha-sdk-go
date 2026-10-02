# DnsZoneInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ZoneId** | Pointer to **string** | DNS zone identifier | [optional] 
**DomainId** | Pointer to **string** | Domain service id when the zone is tied to a domain | [optional] 
**DomainName** | Pointer to **string** | Zone domain name (FQDN) | [optional] 
**HostingId** | Pointer to **int32** | Linked hosting service id; omitted when no hosting is linked | [optional] 
**Type** | Pointer to **int32** | Zone type: 0&#x3D;OTHER, 1&#x3D;DOMAIN, 2&#x3D;HOSTING, 3&#x3D;ADDON | [optional] 
**TypeKey** | Pointer to **string** | Zone type key (e.g. other, domain, hosting, addon) | [optional] 
**PackageName** | Pointer to **string** | Product or package name for the zone | [optional] 
**HasHosting** | Pointer to [**DomainHostingLink**](DomainHostingLink.md) |  | [optional] 
**HasDnsManagerZone** | **bool** | Whether a DNS Manager zone exists for this domain name | 
**Backend** | Pointer to [**DnsBackend**](DnsBackend.md) |  | [optional] 

## Methods

### NewDnsZoneInfo

`func NewDnsZoneInfo(hasDnsManagerZone bool, ) *DnsZoneInfo`

NewDnsZoneInfo instantiates a new DnsZoneInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsZoneInfoWithDefaults

`func NewDnsZoneInfoWithDefaults() *DnsZoneInfo`

NewDnsZoneInfoWithDefaults instantiates a new DnsZoneInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetZoneId

`func (o *DnsZoneInfo) GetZoneId() string`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *DnsZoneInfo) GetZoneIdOk() (*string, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *DnsZoneInfo) SetZoneId(v string)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *DnsZoneInfo) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

### GetDomainId

`func (o *DnsZoneInfo) GetDomainId() string`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *DnsZoneInfo) GetDomainIdOk() (*string, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *DnsZoneInfo) SetDomainId(v string)`

SetDomainId sets DomainId field to given value.

### HasDomainId

`func (o *DnsZoneInfo) HasDomainId() bool`

HasDomainId returns a boolean if a field has been set.

### GetDomainName

`func (o *DnsZoneInfo) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *DnsZoneInfo) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *DnsZoneInfo) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *DnsZoneInfo) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetHostingId

`func (o *DnsZoneInfo) GetHostingId() int32`

GetHostingId returns the HostingId field if non-nil, zero value otherwise.

### GetHostingIdOk

`func (o *DnsZoneInfo) GetHostingIdOk() (*int32, bool)`

GetHostingIdOk returns a tuple with the HostingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostingId

`func (o *DnsZoneInfo) SetHostingId(v int32)`

SetHostingId sets HostingId field to given value.

### HasHostingId

`func (o *DnsZoneInfo) HasHostingId() bool`

HasHostingId returns a boolean if a field has been set.

### GetType

`func (o *DnsZoneInfo) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DnsZoneInfo) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DnsZoneInfo) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *DnsZoneInfo) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTypeKey

`func (o *DnsZoneInfo) GetTypeKey() string`

GetTypeKey returns the TypeKey field if non-nil, zero value otherwise.

### GetTypeKeyOk

`func (o *DnsZoneInfo) GetTypeKeyOk() (*string, bool)`

GetTypeKeyOk returns a tuple with the TypeKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeKey

`func (o *DnsZoneInfo) SetTypeKey(v string)`

SetTypeKey sets TypeKey field to given value.

### HasTypeKey

`func (o *DnsZoneInfo) HasTypeKey() bool`

HasTypeKey returns a boolean if a field has been set.

### GetPackageName

`func (o *DnsZoneInfo) GetPackageName() string`

GetPackageName returns the PackageName field if non-nil, zero value otherwise.

### GetPackageNameOk

`func (o *DnsZoneInfo) GetPackageNameOk() (*string, bool)`

GetPackageNameOk returns a tuple with the PackageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageName

`func (o *DnsZoneInfo) SetPackageName(v string)`

SetPackageName sets PackageName field to given value.

### HasPackageName

`func (o *DnsZoneInfo) HasPackageName() bool`

HasPackageName returns a boolean if a field has been set.

### GetHasHosting

`func (o *DnsZoneInfo) GetHasHosting() DomainHostingLink`

GetHasHosting returns the HasHosting field if non-nil, zero value otherwise.

### GetHasHostingOk

`func (o *DnsZoneInfo) GetHasHostingOk() (*DomainHostingLink, bool)`

GetHasHostingOk returns a tuple with the HasHosting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasHosting

`func (o *DnsZoneInfo) SetHasHosting(v DomainHostingLink)`

SetHasHosting sets HasHosting field to given value.

### HasHasHosting

`func (o *DnsZoneInfo) HasHasHosting() bool`

HasHasHosting returns a boolean if a field has been set.

### GetHasDnsManagerZone

`func (o *DnsZoneInfo) GetHasDnsManagerZone() bool`

GetHasDnsManagerZone returns the HasDnsManagerZone field if non-nil, zero value otherwise.

### GetHasDnsManagerZoneOk

`func (o *DnsZoneInfo) GetHasDnsManagerZoneOk() (*bool, bool)`

GetHasDnsManagerZoneOk returns a tuple with the HasDnsManagerZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDnsManagerZone

`func (o *DnsZoneInfo) SetHasDnsManagerZone(v bool)`

SetHasDnsManagerZone sets HasDnsManagerZone field to given value.


### GetBackend

`func (o *DnsZoneInfo) GetBackend() DnsBackend`

GetBackend returns the Backend field if non-nil, zero value otherwise.

### GetBackendOk

`func (o *DnsZoneInfo) GetBackendOk() (*DnsBackend, bool)`

GetBackendOk returns a tuple with the Backend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackend

`func (o *DnsZoneInfo) SetBackend(v DnsBackend)`

SetBackend sets Backend field to given value.

### HasBackend

`func (o *DnsZoneInfo) HasBackend() bool`

HasBackend returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


