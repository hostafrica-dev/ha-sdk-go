# ModifyPrivateNameserverRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainId** | **string** | Domain service id - must be sent as a string | 
**Nameserver** | **string** | Existing child nameserver hostname to modify | 
**NewNameserver** | Pointer to **string** | New hostname; unsupported for some registrars (COZA/COCCA) | [optional] 
**OldIpAddress** | Pointer to **string** | Current glue IP when the registrar requires it | [optional] 
**NewIpAddress** | Pointer to **string** | New glue IPv4 or IPv6 address | [optional] 

## Methods

### NewModifyPrivateNameserverRequestContent

`func NewModifyPrivateNameserverRequestContent(domainId string, nameserver string, ) *ModifyPrivateNameserverRequestContent`

NewModifyPrivateNameserverRequestContent instantiates a new ModifyPrivateNameserverRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyPrivateNameserverRequestContentWithDefaults

`func NewModifyPrivateNameserverRequestContentWithDefaults() *ModifyPrivateNameserverRequestContent`

NewModifyPrivateNameserverRequestContentWithDefaults instantiates a new ModifyPrivateNameserverRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomainId

`func (o *ModifyPrivateNameserverRequestContent) GetDomainId() string`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *ModifyPrivateNameserverRequestContent) GetDomainIdOk() (*string, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *ModifyPrivateNameserverRequestContent) SetDomainId(v string)`

SetDomainId sets DomainId field to given value.


### GetNameserver

`func (o *ModifyPrivateNameserverRequestContent) GetNameserver() string`

GetNameserver returns the Nameserver field if non-nil, zero value otherwise.

### GetNameserverOk

`func (o *ModifyPrivateNameserverRequestContent) GetNameserverOk() (*string, bool)`

GetNameserverOk returns a tuple with the Nameserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameserver

`func (o *ModifyPrivateNameserverRequestContent) SetNameserver(v string)`

SetNameserver sets Nameserver field to given value.


### GetNewNameserver

`func (o *ModifyPrivateNameserverRequestContent) GetNewNameserver() string`

GetNewNameserver returns the NewNameserver field if non-nil, zero value otherwise.

### GetNewNameserverOk

`func (o *ModifyPrivateNameserverRequestContent) GetNewNameserverOk() (*string, bool)`

GetNewNameserverOk returns a tuple with the NewNameserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewNameserver

`func (o *ModifyPrivateNameserverRequestContent) SetNewNameserver(v string)`

SetNewNameserver sets NewNameserver field to given value.

### HasNewNameserver

`func (o *ModifyPrivateNameserverRequestContent) HasNewNameserver() bool`

HasNewNameserver returns a boolean if a field has been set.

### GetOldIpAddress

`func (o *ModifyPrivateNameserverRequestContent) GetOldIpAddress() string`

GetOldIpAddress returns the OldIpAddress field if non-nil, zero value otherwise.

### GetOldIpAddressOk

`func (o *ModifyPrivateNameserverRequestContent) GetOldIpAddressOk() (*string, bool)`

GetOldIpAddressOk returns a tuple with the OldIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldIpAddress

`func (o *ModifyPrivateNameserverRequestContent) SetOldIpAddress(v string)`

SetOldIpAddress sets OldIpAddress field to given value.

### HasOldIpAddress

`func (o *ModifyPrivateNameserverRequestContent) HasOldIpAddress() bool`

HasOldIpAddress returns a boolean if a field has been set.

### GetNewIpAddress

`func (o *ModifyPrivateNameserverRequestContent) GetNewIpAddress() string`

GetNewIpAddress returns the NewIpAddress field if non-nil, zero value otherwise.

### GetNewIpAddressOk

`func (o *ModifyPrivateNameserverRequestContent) GetNewIpAddressOk() (*string, bool)`

GetNewIpAddressOk returns a tuple with the NewIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewIpAddress

`func (o *ModifyPrivateNameserverRequestContent) SetNewIpAddress(v string)`

SetNewIpAddress sets NewIpAddress field to given value.

### HasNewIpAddress

`func (o *ModifyPrivateNameserverRequestContent) HasNewIpAddress() bool`

HasNewIpAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


