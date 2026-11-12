# DeleteDomainRedirectData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**Deleted** | Pointer to **bool** | True when upstream confirmed deletion | [optional] 
**Domain** | Pointer to **string** | Domain name when returned by upstream | [optional] 

## Methods

### NewDeleteDomainRedirectData

`func NewDeleteDomainRedirectData(message string, ) *DeleteDomainRedirectData`

NewDeleteDomainRedirectData instantiates a new DeleteDomainRedirectData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteDomainRedirectDataWithDefaults

`func NewDeleteDomainRedirectDataWithDefaults() *DeleteDomainRedirectData`

NewDeleteDomainRedirectDataWithDefaults instantiates a new DeleteDomainRedirectData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *DeleteDomainRedirectData) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DeleteDomainRedirectData) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DeleteDomainRedirectData) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetDeleted

`func (o *DeleteDomainRedirectData) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *DeleteDomainRedirectData) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *DeleteDomainRedirectData) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *DeleteDomainRedirectData) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetDomain

`func (o *DeleteDomainRedirectData) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *DeleteDomainRedirectData) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *DeleteDomainRedirectData) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *DeleteDomainRedirectData) HasDomain() bool`

HasDomain returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


