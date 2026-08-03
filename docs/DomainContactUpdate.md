# DomainContactUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**DomainContactSourceType**](DomainContactSourceType.md) |  | 
**Id** | Pointer to **int32** | Saved WHMCS contact id; required when type is contact | [optional] 
**Fields** | Pointer to **map[string]string** | WHOIS contact field values for one role; field names vary by TLD/registrar | [optional] 

## Methods

### NewDomainContactUpdate

`func NewDomainContactUpdate(type_ DomainContactSourceType, ) *DomainContactUpdate`

NewDomainContactUpdate instantiates a new DomainContactUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainContactUpdateWithDefaults

`func NewDomainContactUpdateWithDefaults() *DomainContactUpdate`

NewDomainContactUpdateWithDefaults instantiates a new DomainContactUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DomainContactUpdate) GetType() DomainContactSourceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DomainContactUpdate) GetTypeOk() (*DomainContactSourceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DomainContactUpdate) SetType(v DomainContactSourceType)`

SetType sets Type field to given value.


### GetId

`func (o *DomainContactUpdate) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DomainContactUpdate) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DomainContactUpdate) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *DomainContactUpdate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFields

`func (o *DomainContactUpdate) GetFields() map[string]string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *DomainContactUpdate) GetFieldsOk() (*map[string]string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *DomainContactUpdate) SetFields(v map[string]string)`

SetFields sets Fields field to given value.

### HasFields

`func (o *DomainContactUpdate) HasFields() bool`

HasFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


