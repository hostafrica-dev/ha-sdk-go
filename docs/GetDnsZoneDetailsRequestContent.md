# GetDnsZoneDetailsRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainId** | **string** | Domain service id - must be sent as a string | 
**Backend** | Pointer to [**DnsBackend**](DnsBackend.md) |  | [optional] 

## Methods

### NewGetDnsZoneDetailsRequestContent

`func NewGetDnsZoneDetailsRequestContent(domainId string, ) *GetDnsZoneDetailsRequestContent`

NewGetDnsZoneDetailsRequestContent instantiates a new GetDnsZoneDetailsRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDnsZoneDetailsRequestContentWithDefaults

`func NewGetDnsZoneDetailsRequestContentWithDefaults() *GetDnsZoneDetailsRequestContent`

NewGetDnsZoneDetailsRequestContentWithDefaults instantiates a new GetDnsZoneDetailsRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomainId

`func (o *GetDnsZoneDetailsRequestContent) GetDomainId() string`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *GetDnsZoneDetailsRequestContent) GetDomainIdOk() (*string, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *GetDnsZoneDetailsRequestContent) SetDomainId(v string)`

SetDomainId sets DomainId field to given value.


### GetBackend

`func (o *GetDnsZoneDetailsRequestContent) GetBackend() DnsBackend`

GetBackend returns the Backend field if non-nil, zero value otherwise.

### GetBackendOk

`func (o *GetDnsZoneDetailsRequestContent) GetBackendOk() (*DnsBackend, bool)`

GetBackendOk returns a tuple with the Backend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackend

`func (o *GetDnsZoneDetailsRequestContent) SetBackend(v DnsBackend)`

SetBackend sets Backend field to given value.

### HasBackend

`func (o *GetDnsZoneDetailsRequestContent) HasBackend() bool`

HasBackend returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


