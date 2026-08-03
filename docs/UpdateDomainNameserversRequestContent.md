# UpdateDomainNameserversRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainId** | **string** | Domain service id - must be sent as a string | 
**Ns1** | **string** | Primary nameserver hostname | 
**Ns2** | **string** | Secondary nameserver hostname | 
**Ns3** | Pointer to **string** | Optional tertiary nameserver hostname | [optional] 
**Ns4** | Pointer to **string** | Optional quaternary nameserver hostname | [optional] 
**Ns5** | Pointer to **string** | Optional fifth nameserver hostname | [optional] 

## Methods

### NewUpdateDomainNameserversRequestContent

`func NewUpdateDomainNameserversRequestContent(domainId string, ns1 string, ns2 string, ) *UpdateDomainNameserversRequestContent`

NewUpdateDomainNameserversRequestContent instantiates a new UpdateDomainNameserversRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDomainNameserversRequestContentWithDefaults

`func NewUpdateDomainNameserversRequestContentWithDefaults() *UpdateDomainNameserversRequestContent`

NewUpdateDomainNameserversRequestContentWithDefaults instantiates a new UpdateDomainNameserversRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomainId

`func (o *UpdateDomainNameserversRequestContent) GetDomainId() string`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *UpdateDomainNameserversRequestContent) GetDomainIdOk() (*string, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *UpdateDomainNameserversRequestContent) SetDomainId(v string)`

SetDomainId sets DomainId field to given value.


### GetNs1

`func (o *UpdateDomainNameserversRequestContent) GetNs1() string`

GetNs1 returns the Ns1 field if non-nil, zero value otherwise.

### GetNs1Ok

`func (o *UpdateDomainNameserversRequestContent) GetNs1Ok() (*string, bool)`

GetNs1Ok returns a tuple with the Ns1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNs1

`func (o *UpdateDomainNameserversRequestContent) SetNs1(v string)`

SetNs1 sets Ns1 field to given value.


### GetNs2

`func (o *UpdateDomainNameserversRequestContent) GetNs2() string`

GetNs2 returns the Ns2 field if non-nil, zero value otherwise.

### GetNs2Ok

`func (o *UpdateDomainNameserversRequestContent) GetNs2Ok() (*string, bool)`

GetNs2Ok returns a tuple with the Ns2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNs2

`func (o *UpdateDomainNameserversRequestContent) SetNs2(v string)`

SetNs2 sets Ns2 field to given value.


### GetNs3

`func (o *UpdateDomainNameserversRequestContent) GetNs3() string`

GetNs3 returns the Ns3 field if non-nil, zero value otherwise.

### GetNs3Ok

`func (o *UpdateDomainNameserversRequestContent) GetNs3Ok() (*string, bool)`

GetNs3Ok returns a tuple with the Ns3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNs3

`func (o *UpdateDomainNameserversRequestContent) SetNs3(v string)`

SetNs3 sets Ns3 field to given value.

### HasNs3

`func (o *UpdateDomainNameserversRequestContent) HasNs3() bool`

HasNs3 returns a boolean if a field has been set.

### GetNs4

`func (o *UpdateDomainNameserversRequestContent) GetNs4() string`

GetNs4 returns the Ns4 field if non-nil, zero value otherwise.

### GetNs4Ok

`func (o *UpdateDomainNameserversRequestContent) GetNs4Ok() (*string, bool)`

GetNs4Ok returns a tuple with the Ns4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNs4

`func (o *UpdateDomainNameserversRequestContent) SetNs4(v string)`

SetNs4 sets Ns4 field to given value.

### HasNs4

`func (o *UpdateDomainNameserversRequestContent) HasNs4() bool`

HasNs4 returns a boolean if a field has been set.

### GetNs5

`func (o *UpdateDomainNameserversRequestContent) GetNs5() string`

GetNs5 returns the Ns5 field if non-nil, zero value otherwise.

### GetNs5Ok

`func (o *UpdateDomainNameserversRequestContent) GetNs5Ok() (*string, bool)`

GetNs5Ok returns a tuple with the Ns5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNs5

`func (o *UpdateDomainNameserversRequestContent) SetNs5(v string)`

SetNs5 sets Ns5 field to given value.

### HasNs5

`func (o *UpdateDomainNameserversRequestContent) HasNs5() bool`

HasNs5 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


