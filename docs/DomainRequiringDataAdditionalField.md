# DomainRequiringDataAdditionalField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Field key to use when saving additional domain data | 
**Displayname** | **string** | Human-readable label for the field | 
**Type** | **string** | Input type (e.g. text) | 
**Required** | **bool** | Whether the field must be provided | 
**Value** | **string** | Current value; empty when not yet supplied | 

## Methods

### NewDomainRequiringDataAdditionalField

`func NewDomainRequiringDataAdditionalField(name string, displayname string, type_ string, required bool, value string, ) *DomainRequiringDataAdditionalField`

NewDomainRequiringDataAdditionalField instantiates a new DomainRequiringDataAdditionalField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainRequiringDataAdditionalFieldWithDefaults

`func NewDomainRequiringDataAdditionalFieldWithDefaults() *DomainRequiringDataAdditionalField`

NewDomainRequiringDataAdditionalFieldWithDefaults instantiates a new DomainRequiringDataAdditionalField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DomainRequiringDataAdditionalField) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DomainRequiringDataAdditionalField) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DomainRequiringDataAdditionalField) SetName(v string)`

SetName sets Name field to given value.


### GetDisplayname

`func (o *DomainRequiringDataAdditionalField) GetDisplayname() string`

GetDisplayname returns the Displayname field if non-nil, zero value otherwise.

### GetDisplaynameOk

`func (o *DomainRequiringDataAdditionalField) GetDisplaynameOk() (*string, bool)`

GetDisplaynameOk returns a tuple with the Displayname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayname

`func (o *DomainRequiringDataAdditionalField) SetDisplayname(v string)`

SetDisplayname sets Displayname field to given value.


### GetType

`func (o *DomainRequiringDataAdditionalField) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DomainRequiringDataAdditionalField) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DomainRequiringDataAdditionalField) SetType(v string)`

SetType sets Type field to given value.


### GetRequired

`func (o *DomainRequiringDataAdditionalField) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *DomainRequiringDataAdditionalField) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *DomainRequiringDataAdditionalField) SetRequired(v bool)`

SetRequired sets Required field to given value.


### GetValue

`func (o *DomainRequiringDataAdditionalField) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DomainRequiringDataAdditionalField) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DomainRequiringDataAdditionalField) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


