# AddDomainRedirectData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**Domain** | Pointer to **string** | Domain name when returned by upstream | [optional] 
**Url** | Pointer to **string** | Redirect target URL when returned by upstream | [optional] 
**Type** | Pointer to **string** | Redirect type when returned by upstream (e.g. 301, 302) | [optional] 

## Methods

### NewAddDomainRedirectData

`func NewAddDomainRedirectData(message string, ) *AddDomainRedirectData`

NewAddDomainRedirectData instantiates a new AddDomainRedirectData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddDomainRedirectDataWithDefaults

`func NewAddDomainRedirectDataWithDefaults() *AddDomainRedirectData`

NewAddDomainRedirectDataWithDefaults instantiates a new AddDomainRedirectData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *AddDomainRedirectData) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AddDomainRedirectData) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AddDomainRedirectData) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetDomain

`func (o *AddDomainRedirectData) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *AddDomainRedirectData) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *AddDomainRedirectData) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *AddDomainRedirectData) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetUrl

`func (o *AddDomainRedirectData) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AddDomainRedirectData) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AddDomainRedirectData) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AddDomainRedirectData) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetType

`func (o *AddDomainRedirectData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddDomainRedirectData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddDomainRedirectData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AddDomainRedirectData) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


