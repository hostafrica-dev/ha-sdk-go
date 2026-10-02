# AddDnsRecordRequestContent

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

### NewAddDnsRecordRequestContent

`func NewAddDnsRecordRequestContent(record DnsRecordMutationRecord, ) *AddDnsRecordRequestContent`

NewAddDnsRecordRequestContent instantiates a new AddDnsRecordRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddDnsRecordRequestContentWithDefaults

`func NewAddDnsRecordRequestContentWithDefaults() *AddDnsRecordRequestContent`

NewAddDnsRecordRequestContentWithDefaults instantiates a new AddDnsRecordRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomainName

`func (o *AddDnsRecordRequestContent) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *AddDnsRecordRequestContent) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *AddDnsRecordRequestContent) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *AddDnsRecordRequestContent) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetZoneId

`func (o *AddDnsRecordRequestContent) GetZoneId() string`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *AddDnsRecordRequestContent) GetZoneIdOk() (*string, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *AddDnsRecordRequestContent) SetZoneId(v string)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *AddDnsRecordRequestContent) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

### GetDomainId

`func (o *AddDnsRecordRequestContent) GetDomainId() string`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *AddDnsRecordRequestContent) GetDomainIdOk() (*string, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *AddDnsRecordRequestContent) SetDomainId(v string)`

SetDomainId sets DomainId field to given value.

### HasDomainId

`func (o *AddDnsRecordRequestContent) HasDomainId() bool`

HasDomainId returns a boolean if a field has been set.

### GetServiceId

`func (o *AddDnsRecordRequestContent) GetServiceId() int32`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *AddDnsRecordRequestContent) GetServiceIdOk() (*int32, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *AddDnsRecordRequestContent) SetServiceId(v int32)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *AddDnsRecordRequestContent) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetBackend

`func (o *AddDnsRecordRequestContent) GetBackend() DnsBackend`

GetBackend returns the Backend field if non-nil, zero value otherwise.

### GetBackendOk

`func (o *AddDnsRecordRequestContent) GetBackendOk() (*DnsBackend, bool)`

GetBackendOk returns a tuple with the Backend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackend

`func (o *AddDnsRecordRequestContent) SetBackend(v DnsBackend)`

SetBackend sets Backend field to given value.

### HasBackend

`func (o *AddDnsRecordRequestContent) HasBackend() bool`

HasBackend returns a boolean if a field has been set.

### GetRecord

`func (o *AddDnsRecordRequestContent) GetRecord() DnsRecordMutationRecord`

GetRecord returns the Record field if non-nil, zero value otherwise.

### GetRecordOk

`func (o *AddDnsRecordRequestContent) GetRecordOk() (*DnsRecordMutationRecord, bool)`

GetRecordOk returns a tuple with the Record field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecord

`func (o *AddDnsRecordRequestContent) SetRecord(v DnsRecordMutationRecord)`

SetRecord sets Record field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


