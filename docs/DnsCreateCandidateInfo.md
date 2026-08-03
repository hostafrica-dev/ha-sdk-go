# DnsCreateCandidateInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainName** | Pointer to **string** | Domain name when the candidate is a domain | [optional] 
**DomainId** | Pointer to **string** | Domain service id when the candidate is a domain | [optional] 
**Relid** | Pointer to **int32** | Related service id when the candidate is a hosting or addon service | [optional] 
**Type** | Pointer to **int32** | Zone type: 0&#x3D;OTHER, 1&#x3D;DOMAIN, 2&#x3D;HOSTING, 3&#x3D;ADDON | [optional] 
**Name** | Pointer to **string** | Display name for the candidate | [optional] 

## Methods

### NewDnsCreateCandidateInfo

`func NewDnsCreateCandidateInfo() *DnsCreateCandidateInfo`

NewDnsCreateCandidateInfo instantiates a new DnsCreateCandidateInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsCreateCandidateInfoWithDefaults

`func NewDnsCreateCandidateInfoWithDefaults() *DnsCreateCandidateInfo`

NewDnsCreateCandidateInfoWithDefaults instantiates a new DnsCreateCandidateInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomainName

`func (o *DnsCreateCandidateInfo) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *DnsCreateCandidateInfo) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *DnsCreateCandidateInfo) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *DnsCreateCandidateInfo) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetDomainId

`func (o *DnsCreateCandidateInfo) GetDomainId() string`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *DnsCreateCandidateInfo) GetDomainIdOk() (*string, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *DnsCreateCandidateInfo) SetDomainId(v string)`

SetDomainId sets DomainId field to given value.

### HasDomainId

`func (o *DnsCreateCandidateInfo) HasDomainId() bool`

HasDomainId returns a boolean if a field has been set.

### GetRelid

`func (o *DnsCreateCandidateInfo) GetRelid() int32`

GetRelid returns the Relid field if non-nil, zero value otherwise.

### GetRelidOk

`func (o *DnsCreateCandidateInfo) GetRelidOk() (*int32, bool)`

GetRelidOk returns a tuple with the Relid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelid

`func (o *DnsCreateCandidateInfo) SetRelid(v int32)`

SetRelid sets Relid field to given value.

### HasRelid

`func (o *DnsCreateCandidateInfo) HasRelid() bool`

HasRelid returns a boolean if a field has been set.

### GetType

`func (o *DnsCreateCandidateInfo) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DnsCreateCandidateInfo) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DnsCreateCandidateInfo) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *DnsCreateCandidateInfo) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *DnsCreateCandidateInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DnsCreateCandidateInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DnsCreateCandidateInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DnsCreateCandidateInfo) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


