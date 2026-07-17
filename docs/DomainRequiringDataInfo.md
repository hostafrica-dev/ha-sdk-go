# DomainRequiringDataInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainId** | **string** | Domain service id | 
**Domain** | **string** | Fully qualified domain name | 
**Status** | **string** | Domain status (e.g. Pending) | 
**DomainType** | **string** | Domain operation type (e.g. Register) | 
**Tld** | **string** | Top-level domain suffix (e.g. .co.za) | 
**AdditionalFields** | [**[]DomainRequiringDataAdditionalField**](DomainRequiringDataAdditionalField.md) | Additional registrar fields required for this TLD | 

## Methods

### NewDomainRequiringDataInfo

`func NewDomainRequiringDataInfo(domainId string, domain string, status string, domainType string, tld string, additionalFields []DomainRequiringDataAdditionalField, ) *DomainRequiringDataInfo`

NewDomainRequiringDataInfo instantiates a new DomainRequiringDataInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainRequiringDataInfoWithDefaults

`func NewDomainRequiringDataInfoWithDefaults() *DomainRequiringDataInfo`

NewDomainRequiringDataInfoWithDefaults instantiates a new DomainRequiringDataInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomainId

`func (o *DomainRequiringDataInfo) GetDomainId() string`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *DomainRequiringDataInfo) GetDomainIdOk() (*string, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *DomainRequiringDataInfo) SetDomainId(v string)`

SetDomainId sets DomainId field to given value.


### GetDomain

`func (o *DomainRequiringDataInfo) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *DomainRequiringDataInfo) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *DomainRequiringDataInfo) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetStatus

`func (o *DomainRequiringDataInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DomainRequiringDataInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DomainRequiringDataInfo) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetDomainType

`func (o *DomainRequiringDataInfo) GetDomainType() string`

GetDomainType returns the DomainType field if non-nil, zero value otherwise.

### GetDomainTypeOk

`func (o *DomainRequiringDataInfo) GetDomainTypeOk() (*string, bool)`

GetDomainTypeOk returns a tuple with the DomainType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainType

`func (o *DomainRequiringDataInfo) SetDomainType(v string)`

SetDomainType sets DomainType field to given value.


### GetTld

`func (o *DomainRequiringDataInfo) GetTld() string`

GetTld returns the Tld field if non-nil, zero value otherwise.

### GetTldOk

`func (o *DomainRequiringDataInfo) GetTldOk() (*string, bool)`

GetTldOk returns a tuple with the Tld field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTld

`func (o *DomainRequiringDataInfo) SetTld(v string)`

SetTld sets Tld field to given value.


### GetAdditionalFields

`func (o *DomainRequiringDataInfo) GetAdditionalFields() []DomainRequiringDataAdditionalField`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *DomainRequiringDataInfo) GetAdditionalFieldsOk() (*[]DomainRequiringDataAdditionalField, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *DomainRequiringDataInfo) SetAdditionalFields(v []DomainRequiringDataAdditionalField)`

SetAdditionalFields sets AdditionalFields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


