# SaveDomainRequiredDataRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainId** | **string** | Domain service id - must be sent as a string | 
**Fields** | **map[string]string** | Registrar field values keyed by additionalFields[].name | 

## Methods

### NewSaveDomainRequiredDataRequestContent

`func NewSaveDomainRequiredDataRequestContent(domainId string, fields map[string]string, ) *SaveDomainRequiredDataRequestContent`

NewSaveDomainRequiredDataRequestContent instantiates a new SaveDomainRequiredDataRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSaveDomainRequiredDataRequestContentWithDefaults

`func NewSaveDomainRequiredDataRequestContentWithDefaults() *SaveDomainRequiredDataRequestContent`

NewSaveDomainRequiredDataRequestContentWithDefaults instantiates a new SaveDomainRequiredDataRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomainId

`func (o *SaveDomainRequiredDataRequestContent) GetDomainId() string`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *SaveDomainRequiredDataRequestContent) GetDomainIdOk() (*string, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *SaveDomainRequiredDataRequestContent) SetDomainId(v string)`

SetDomainId sets DomainId field to given value.


### GetFields

`func (o *SaveDomainRequiredDataRequestContent) GetFields() map[string]string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *SaveDomainRequiredDataRequestContent) GetFieldsOk() (*map[string]string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *SaveDomainRequiredDataRequestContent) SetFields(v map[string]string)`

SetFields sets Fields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


