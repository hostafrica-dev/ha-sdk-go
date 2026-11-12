# RegisterPrivateNameserverRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainId** | **string** | Domain service id - must be sent as a string | 
**Nameserver** | **string** | Full child nameserver hostname (e.g. ns1.example.com) | 
**IpAddress** | **string** | Glue IPv4 or IPv6 address for the nameserver | 

## Methods

### NewRegisterPrivateNameserverRequestContent

`func NewRegisterPrivateNameserverRequestContent(domainId string, nameserver string, ipAddress string, ) *RegisterPrivateNameserverRequestContent`

NewRegisterPrivateNameserverRequestContent instantiates a new RegisterPrivateNameserverRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterPrivateNameserverRequestContentWithDefaults

`func NewRegisterPrivateNameserverRequestContentWithDefaults() *RegisterPrivateNameserverRequestContent`

NewRegisterPrivateNameserverRequestContentWithDefaults instantiates a new RegisterPrivateNameserverRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomainId

`func (o *RegisterPrivateNameserverRequestContent) GetDomainId() string`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *RegisterPrivateNameserverRequestContent) GetDomainIdOk() (*string, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *RegisterPrivateNameserverRequestContent) SetDomainId(v string)`

SetDomainId sets DomainId field to given value.


### GetNameserver

`func (o *RegisterPrivateNameserverRequestContent) GetNameserver() string`

GetNameserver returns the Nameserver field if non-nil, zero value otherwise.

### GetNameserverOk

`func (o *RegisterPrivateNameserverRequestContent) GetNameserverOk() (*string, bool)`

GetNameserverOk returns a tuple with the Nameserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameserver

`func (o *RegisterPrivateNameserverRequestContent) SetNameserver(v string)`

SetNameserver sets Nameserver field to given value.


### GetIpAddress

`func (o *RegisterPrivateNameserverRequestContent) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *RegisterPrivateNameserverRequestContent) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *RegisterPrivateNameserverRequestContent) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


