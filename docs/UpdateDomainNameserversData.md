# UpdateDomainNameserversData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Status message indicating the result | 
**Domain** | Pointer to **string** | Domain name (FQDN) | [optional] 
**Nameservers** | Pointer to [**DomainNameservers**](DomainNameservers.md) |  | [optional] 

## Methods

### NewUpdateDomainNameserversData

`func NewUpdateDomainNameserversData(message string, ) *UpdateDomainNameserversData`

NewUpdateDomainNameserversData instantiates a new UpdateDomainNameserversData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDomainNameserversDataWithDefaults

`func NewUpdateDomainNameserversDataWithDefaults() *UpdateDomainNameserversData`

NewUpdateDomainNameserversDataWithDefaults instantiates a new UpdateDomainNameserversData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *UpdateDomainNameserversData) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *UpdateDomainNameserversData) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *UpdateDomainNameserversData) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetDomain

`func (o *UpdateDomainNameserversData) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *UpdateDomainNameserversData) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *UpdateDomainNameserversData) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *UpdateDomainNameserversData) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetNameservers

`func (o *UpdateDomainNameserversData) GetNameservers() DomainNameservers`

GetNameservers returns the Nameservers field if non-nil, zero value otherwise.

### GetNameserversOk

`func (o *UpdateDomainNameserversData) GetNameserversOk() (*DomainNameservers, bool)`

GetNameserversOk returns a tuple with the Nameservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameservers

`func (o *UpdateDomainNameserversData) SetNameservers(v DomainNameservers)`

SetNameservers sets Nameservers field to given value.

### HasNameservers

`func (o *UpdateDomainNameserversData) HasNameservers() bool`

HasNameservers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


