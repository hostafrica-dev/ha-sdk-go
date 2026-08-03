# DomainContacts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Registrant** | Pointer to **map[string]string** | WHOIS contact field values for one role; field names vary by TLD/registrar | [optional] 
**Admin** | Pointer to **map[string]string** | WHOIS contact field values for one role; field names vary by TLD/registrar | [optional] 
**Tech** | Pointer to **map[string]string** | WHOIS contact field values for one role; field names vary by TLD/registrar | [optional] 
**Billing** | Pointer to **map[string]string** | WHOIS contact field values for one role; field names vary by TLD/registrar | [optional] 

## Methods

### NewDomainContacts

`func NewDomainContacts() *DomainContacts`

NewDomainContacts instantiates a new DomainContacts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainContactsWithDefaults

`func NewDomainContactsWithDefaults() *DomainContacts`

NewDomainContactsWithDefaults instantiates a new DomainContacts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegistrant

`func (o *DomainContacts) GetRegistrant() map[string]string`

GetRegistrant returns the Registrant field if non-nil, zero value otherwise.

### GetRegistrantOk

`func (o *DomainContacts) GetRegistrantOk() (*map[string]string, bool)`

GetRegistrantOk returns a tuple with the Registrant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrant

`func (o *DomainContacts) SetRegistrant(v map[string]string)`

SetRegistrant sets Registrant field to given value.

### HasRegistrant

`func (o *DomainContacts) HasRegistrant() bool`

HasRegistrant returns a boolean if a field has been set.

### GetAdmin

`func (o *DomainContacts) GetAdmin() map[string]string`

GetAdmin returns the Admin field if non-nil, zero value otherwise.

### GetAdminOk

`func (o *DomainContacts) GetAdminOk() (*map[string]string, bool)`

GetAdminOk returns a tuple with the Admin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmin

`func (o *DomainContacts) SetAdmin(v map[string]string)`

SetAdmin sets Admin field to given value.

### HasAdmin

`func (o *DomainContacts) HasAdmin() bool`

HasAdmin returns a boolean if a field has been set.

### GetTech

`func (o *DomainContacts) GetTech() map[string]string`

GetTech returns the Tech field if non-nil, zero value otherwise.

### GetTechOk

`func (o *DomainContacts) GetTechOk() (*map[string]string, bool)`

GetTechOk returns a tuple with the Tech field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTech

`func (o *DomainContacts) SetTech(v map[string]string)`

SetTech sets Tech field to given value.

### HasTech

`func (o *DomainContacts) HasTech() bool`

HasTech returns a boolean if a field has been set.

### GetBilling

`func (o *DomainContacts) GetBilling() map[string]string`

GetBilling returns the Billing field if non-nil, zero value otherwise.

### GetBillingOk

`func (o *DomainContacts) GetBillingOk() (*map[string]string, bool)`

GetBillingOk returns a tuple with the Billing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBilling

`func (o *DomainContacts) SetBilling(v map[string]string)`

SetBilling sets Billing field to given value.

### HasBilling

`func (o *DomainContacts) HasBilling() bool`

HasBilling returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


