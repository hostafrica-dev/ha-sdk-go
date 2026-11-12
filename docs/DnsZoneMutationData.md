# DnsZoneMutationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Status message indicating the result | 
**ZoneId** | Pointer to **string** | DNS zone identifier when returned by upstream | [optional] 
**DomainId** | Pointer to **string** | Domain service id when returned by upstream | [optional] 
**DomainName** | Pointer to **string** | Zone domain name (FQDN) when returned by upstream | [optional] 
**Backend** | Pointer to [**DnsBackend**](DnsBackend.md) |  | [optional] 

## Methods

### NewDnsZoneMutationData

`func NewDnsZoneMutationData(message string, ) *DnsZoneMutationData`

NewDnsZoneMutationData instantiates a new DnsZoneMutationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsZoneMutationDataWithDefaults

`func NewDnsZoneMutationDataWithDefaults() *DnsZoneMutationData`

NewDnsZoneMutationDataWithDefaults instantiates a new DnsZoneMutationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *DnsZoneMutationData) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DnsZoneMutationData) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DnsZoneMutationData) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetZoneId

`func (o *DnsZoneMutationData) GetZoneId() string`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *DnsZoneMutationData) GetZoneIdOk() (*string, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *DnsZoneMutationData) SetZoneId(v string)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *DnsZoneMutationData) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

### GetDomainId

`func (o *DnsZoneMutationData) GetDomainId() string`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *DnsZoneMutationData) GetDomainIdOk() (*string, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *DnsZoneMutationData) SetDomainId(v string)`

SetDomainId sets DomainId field to given value.

### HasDomainId

`func (o *DnsZoneMutationData) HasDomainId() bool`

HasDomainId returns a boolean if a field has been set.

### GetDomainName

`func (o *DnsZoneMutationData) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *DnsZoneMutationData) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *DnsZoneMutationData) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *DnsZoneMutationData) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetBackend

`func (o *DnsZoneMutationData) GetBackend() DnsBackend`

GetBackend returns the Backend field if non-nil, zero value otherwise.

### GetBackendOk

`func (o *DnsZoneMutationData) GetBackendOk() (*DnsBackend, bool)`

GetBackendOk returns a tuple with the Backend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackend

`func (o *DnsZoneMutationData) SetBackend(v DnsBackend)`

SetBackend sets Backend field to given value.

### HasBackend

`func (o *DnsZoneMutationData) HasBackend() bool`

HasBackend returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


