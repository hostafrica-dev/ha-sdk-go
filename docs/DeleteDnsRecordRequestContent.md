# DeleteDnsRecordRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainName** | Pointer to **string** | DNS zone domain name (FQDN); optional when zone_id is provided | [optional] 
**ZoneId** | **string** | DNS zone identifier from list-dns-zones or get-dns-zone-details | 
**Record** | [**DnsRecordMutationRecord**](DnsRecordMutationRecord.md) |  | 

## Methods

### NewDeleteDnsRecordRequestContent

`func NewDeleteDnsRecordRequestContent(zoneId string, record DnsRecordMutationRecord, ) *DeleteDnsRecordRequestContent`

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


