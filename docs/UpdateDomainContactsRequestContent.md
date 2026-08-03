# UpdateDomainContactsRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainId** | **string** | Domain service id - must be sent as a string | 
**Contacts** | [**DomainContactUpdates**](DomainContactUpdates.md) |  | 

## Methods

### NewUpdateDomainContactsRequestContent

`func NewUpdateDomainContactsRequestContent(domainId string, contacts DomainContactUpdates, ) *UpdateDomainContactsRequestContent`

NewUpdateDomainContactsRequestContent instantiates a new UpdateDomainContactsRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDomainContactsRequestContentWithDefaults

`func NewUpdateDomainContactsRequestContentWithDefaults() *UpdateDomainContactsRequestContent`

NewUpdateDomainContactsRequestContentWithDefaults instantiates a new UpdateDomainContactsRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomainId

`func (o *UpdateDomainContactsRequestContent) GetDomainId() string`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *UpdateDomainContactsRequestContent) GetDomainIdOk() (*string, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *UpdateDomainContactsRequestContent) SetDomainId(v string)`

SetDomainId sets DomainId field to given value.


### GetContacts

`func (o *UpdateDomainContactsRequestContent) GetContacts() DomainContactUpdates`

GetContacts returns the Contacts field if non-nil, zero value otherwise.

### GetContactsOk

`func (o *UpdateDomainContactsRequestContent) GetContactsOk() (*DomainContactUpdates, bool)`

GetContactsOk returns a tuple with the Contacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContacts

`func (o *UpdateDomainContactsRequestContent) SetContacts(v DomainContactUpdates)`

SetContacts sets Contacts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


