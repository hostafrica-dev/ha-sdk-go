# CreateDnsZoneRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainId** | **string** | Domain service id - must be sent as a string | 
**ZoneIp** | **string** | Apex A record IPv4 for the new zone. Required by DNS Manager create. | 
**DomainName** | Pointer to **string** | Zone domain name (FQDN). Optional; upstream resolves from domain_id when omitted. | [optional] 
**ServiceId** | Pointer to **int32** | WHMCS hosting service id — pass list-dns-zones / create-candidates hosting_id (or relid for HOSTING) as service_id. Helps DirectAdmin create routing when hosting is linked. | [optional] 
**Type** | Pointer to **int32** | Zone type from list-dns-create-candidates: 0&#x3D;OTHER, 1&#x3D;DOMAIN, 2&#x3D;HOSTING, 3&#x3D;ADDON. Optional; upstream defaults from eligible packages. | [optional] 
**Relid** | Pointer to **int32** | Related service id matching the zone type (from list-dns-create-candidates relid). Optional with type; for HOSTING/ADDON often the same value sent as service_id. | [optional] 
**TemplateId** | Pointer to **int32** | Optional DNS template id to apply after zone creation | [optional] 

## Methods

### NewCreateDnsZoneRequestContent

`func NewCreateDnsZoneRequestContent(domainId string, zoneIp string, ) *CreateDnsZoneRequestContent`

NewCreateDnsZoneRequestContent instantiates a new CreateDnsZoneRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDnsZoneRequestContentWithDefaults

`func NewCreateDnsZoneRequestContentWithDefaults() *CreateDnsZoneRequestContent`

NewCreateDnsZoneRequestContentWithDefaults instantiates a new CreateDnsZoneRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomainId

`func (o *CreateDnsZoneRequestContent) GetDomainId() string`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *CreateDnsZoneRequestContent) GetDomainIdOk() (*string, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *CreateDnsZoneRequestContent) SetDomainId(v string)`

SetDomainId sets DomainId field to given value.


### GetZoneIp

`func (o *CreateDnsZoneRequestContent) GetZoneIp() string`

GetZoneIp returns the ZoneIp field if non-nil, zero value otherwise.

### GetZoneIpOk

`func (o *CreateDnsZoneRequestContent) GetZoneIpOk() (*string, bool)`

GetZoneIpOk returns a tuple with the ZoneIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneIp

`func (o *CreateDnsZoneRequestContent) SetZoneIp(v string)`

SetZoneIp sets ZoneIp field to given value.


### GetDomainName

`func (o *CreateDnsZoneRequestContent) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *CreateDnsZoneRequestContent) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *CreateDnsZoneRequestContent) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *CreateDnsZoneRequestContent) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetServiceId

`func (o *CreateDnsZoneRequestContent) GetServiceId() int32`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *CreateDnsZoneRequestContent) GetServiceIdOk() (*int32, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *CreateDnsZoneRequestContent) SetServiceId(v int32)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *CreateDnsZoneRequestContent) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetType

`func (o *CreateDnsZoneRequestContent) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateDnsZoneRequestContent) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateDnsZoneRequestContent) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *CreateDnsZoneRequestContent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRelid

`func (o *CreateDnsZoneRequestContent) GetRelid() int32`

GetRelid returns the Relid field if non-nil, zero value otherwise.

### GetRelidOk

`func (o *CreateDnsZoneRequestContent) GetRelidOk() (*int32, bool)`

GetRelidOk returns a tuple with the Relid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelid

`func (o *CreateDnsZoneRequestContent) SetRelid(v int32)`

SetRelid sets Relid field to given value.

### HasRelid

`func (o *CreateDnsZoneRequestContent) HasRelid() bool`

HasRelid returns a boolean if a field has been set.

### GetTemplateId

`func (o *CreateDnsZoneRequestContent) GetTemplateId() int32`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *CreateDnsZoneRequestContent) GetTemplateIdOk() (*int32, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *CreateDnsZoneRequestContent) SetTemplateId(v int32)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *CreateDnsZoneRequestContent) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


