# GetDomainContactsData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Status message indicating the result | 
**DomainId** | **string** | Domain service id | 
**Domain** | **string** | Fully qualified domain name | 
**Contacts** | **interface{}** | Contact roles keyed by Registrant, Admin, Tech, and Billing, or an array of contact records. Inner field names and values vary by TLD/registrar. | 

## Methods

### NewGetDomainContactsData

`func NewGetDomainContactsData(message string, domainId string, domain string, contacts interface{}, ) *GetDomainContactsData`

NewGetDomainContactsData instantiates a new GetDomainContactsData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDomainContactsDataWithDefaults

`func NewGetDomainContactsDataWithDefaults() *GetDomainContactsData`

NewGetDomainContactsDataWithDefaults instantiates a new GetDomainContactsData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *GetDomainContactsData) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetDomainContactsData) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetDomainContactsData) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetDomainId

`func (o *GetDomainContactsData) GetDomainId() string`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *GetDomainContactsData) GetDomainIdOk() (*string, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *GetDomainContactsData) SetDomainId(v string)`

SetDomainId sets DomainId field to given value.


### GetDomain

`func (o *GetDomainContactsData) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *GetDomainContactsData) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *GetDomainContactsData) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetContacts

`func (o *GetDomainContactsData) GetContacts() interface{}`

GetContacts returns the Contacts field if non-nil, zero value otherwise.

### GetContactsOk

`func (o *GetDomainContactsData) GetContactsOk() (*interface{}, bool)`

GetContactsOk returns a tuple with the Contacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContacts

`func (o *GetDomainContactsData) SetContacts(v interface{})`

SetContacts sets Contacts field to given value.


### SetContactsNil

`func (o *GetDomainContactsData) SetContactsNil(b bool)`

 SetContactsNil sets the value for Contacts to be an explicit nil

### UnsetContacts
`func (o *GetDomainContactsData) UnsetContacts()`

UnsetContacts ensures that no value is present for Contacts, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


