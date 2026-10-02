# DeleteDnsRecordRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainName** | Pointer to **string** | DNS zone domain name (FQDN); optional for dns_manager when zone_id is provided. Not forwarded on DirectAdmin mutations. | [optional] 
**ZoneId** | Pointer to **string** | DNS zone identifier from list-dns-zones or get-dns-zone-details; required for dns_manager / legacy callers | [optional] 
**DomainId** | Pointer to **string** | WHMCS domain id from list-dns-zones; required when backend is directadmin | [optional] 
**ServiceId** | Pointer to **int32** | Optional WHMCS hosting service id from list-dns-zones hosting_id. Not forwarded on DirectAdmin mutations. | [optional] 
**Backend** | Pointer to [**DnsBackend**](DnsBackend.md) |  | [optional] 
**Record** | [**DnsRecordMutationRecord**](DnsRecordMutationRecord.md) |  | 

## Methods

### NewDeleteDnsRecordRequestContent

`func NewDeleteDnsRecordRequestContent(record DnsRecordMutationRecord, ) *DeleteDnsRecordRequestContent`

NewDeleteDnsRecordRequestContent instantiates a new DeleteDnsRecordRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteDnsRecordRequestContentWithDefaults

`func NewDeleteDnsRecordRequestContentWithDefaults() *DeleteDnsRecordRequestContent`

NewDeleteDnsRecordRequestContentWithDefaults instantiates a new DeleteDnsRecordRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomainName

`func (o *DeleteDnsRecordRequestContent) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *DeleteDnsRecordRequestContent) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *DeleteDnsRecordRequestContent) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *DeleteDnsRecordRequestContent) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetZoneId

`func (o *DeleteDnsRecordRequestContent) GetZoneId() string`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *DeleteDnsRecordRequestContent) GetZoneIdOk() (*string, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *DeleteDnsRecordRequestContent) SetZoneId(v string)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *DeleteDnsRecordRequestContent) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

### GetDomainId

`func (o *DeleteDnsRecordRequestContent) GetDomainId() string`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *DeleteDnsRecordRequestContent) GetDomainIdOk() (*string, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *DeleteDnsRecordRequestContent) SetDomainId(v string)`

SetDomainId sets DomainId field to given value.

### HasDomainId

`func (o *DeleteDnsRecordRequestContent) HasDomainId() bool`

HasDomainId returns a boolean if a field has been set.

### GetServiceId

`func (o *DeleteDnsRecordRequestContent) GetServiceId() int32`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *DeleteDnsRecordRequestContent) GetServiceIdOk() (*int32, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *DeleteDnsRecordRequestContent) SetServiceId(v int32)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *DeleteDnsRecordRequestContent) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetBackend

`func (o *DeleteDnsRecordRequestContent) GetBackend() DnsBackend`

GetBackend returns the Backend field if non-nil, zero value otherwise.

### GetBackendOk

`func (o *DeleteDnsRecordRequestContent) GetBackendOk() (*DnsBackend, bool)`

GetBackendOk returns a tuple with the Backend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackend

`func (o *DeleteDnsRecordRequestContent) SetBackend(v DnsBackend)`

SetBackend sets Backend field to given value.

### HasBackend

`func (o *DeleteDnsRecordRequestContent) HasBackend() bool`

HasBackend returns a boolean if a field has been set.

### GetRecord

`func (o *DeleteDnsRecordRequestContent) GetRecord() DnsRecordMutationRecord`

GetRecord returns the Record field if non-nil, zero value otherwise.

### GetRecordOk

`func (o *DeleteDnsRecordRequestContent) GetRecordOk() (*DnsRecordMutationRecord, bool)`

GetRecordOk returns a tuple with the Record field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecord

`func (o *DeleteDnsRecordRequestContent) SetRecord(v DnsRecordMutationRecord)`

SetRecord sets Record field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


