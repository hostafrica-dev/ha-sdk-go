# EditDnsRecordRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainName** | Pointer to **string** | DNS zone domain name (FQDN); optional when zone_id is provided | [optional] 
**ZoneId** | **string** | DNS zone identifier from list-dns-zones or get-dns-zone-details | 
**Record** | [**DnsRecordMutationRecord**](DnsRecordMutationRecord.md) |  | 

## Methods

### NewEditDnsRecordRequestContent

`func NewEditDnsRecordRequestContent(zoneId string, record DnsRecordMutationRecord, ) *EditDnsRecordRequestContent`

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


