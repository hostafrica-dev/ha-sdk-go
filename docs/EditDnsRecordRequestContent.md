# EditDnsRecordRequestContent

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

### NewEditDnsRecordRequestContent

`func NewEditDnsRecordRequestContent(record DnsRecordMutationRecord, ) *EditDnsRecordRequestContent`

NewEditDnsRecordRequestContent instantiates a new EditDnsRecordRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditDnsRecordRequestContentWithDefaults

`func NewEditDnsRecordRequestContentWithDefaults() *EditDnsRecordRequestContent`

NewEditDnsRecordRequestContentWithDefaults instantiates a new EditDnsRecordRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomainName

`func (o *EditDnsRecordRequestContent) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *EditDnsRecordRequestContent) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *EditDnsRecordRequestContent) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *EditDnsRecordRequestContent) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetZoneId

`func (o *EditDnsRecordRequestContent) GetZoneId() string`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *EditDnsRecordRequestContent) GetZoneIdOk() (*string, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *EditDnsRecordRequestContent) SetZoneId(v string)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *EditDnsRecordRequestContent) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

### GetDomainId

`func (o *EditDnsRecordRequestContent) GetDomainId() string`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *EditDnsRecordRequestContent) GetDomainIdOk() (*string, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *EditDnsRecordRequestContent) SetDomainId(v string)`

SetDomainId sets DomainId field to given value.

### HasDomainId

`func (o *EditDnsRecordRequestContent) HasDomainId() bool`

HasDomainId returns a boolean if a field has been set.

### GetServiceId

`func (o *EditDnsRecordRequestContent) GetServiceId() int32`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *EditDnsRecordRequestContent) GetServiceIdOk() (*int32, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *EditDnsRecordRequestContent) SetServiceId(v int32)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *EditDnsRecordRequestContent) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetBackend

`func (o *EditDnsRecordRequestContent) GetBackend() DnsBackend`

GetBackend returns the Backend field if non-nil, zero value otherwise.

### GetBackendOk

`func (o *EditDnsRecordRequestContent) GetBackendOk() (*DnsBackend, bool)`

GetBackendOk returns a tuple with the Backend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackend

`func (o *EditDnsRecordRequestContent) SetBackend(v DnsBackend)`

SetBackend sets Backend field to given value.

### HasBackend

`func (o *EditDnsRecordRequestContent) HasBackend() bool`

HasBackend returns a boolean if a field has been set.

### GetRecord

`func (o *EditDnsRecordRequestContent) GetRecord() DnsRecordMutationRecord`

GetRecord returns the Record field if non-nil, zero value otherwise.

### GetRecordOk

`func (o *EditDnsRecordRequestContent) GetRecordOk() (*DnsRecordMutationRecord, bool)`

GetRecordOk returns a tuple with the Record field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecord

`func (o *EditDnsRecordRequestContent) SetRecord(v DnsRecordMutationRecord)`

SetRecord sets Record field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


