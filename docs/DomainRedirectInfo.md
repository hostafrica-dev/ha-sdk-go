# DomainRedirectInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainId** | Pointer to **string** | Domain service id | [optional] 
**Domain** | Pointer to **string** | Domain name | [optional] 
**Url** | Pointer to **string** | Redirect target URL | [optional] 
**Type** | Pointer to **string** | Redirect type (e.g. 301, 302, permanent, temporary) | [optional] 
**Address** | Pointer to **string** | Optional redirector address override | [optional] 

## Methods

### NewDomainRedirectInfo

`func NewDomainRedirectInfo() *DomainRedirectInfo`

NewDomainRedirectInfo instantiates a new DomainRedirectInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainRedirectInfoWithDefaults

`func NewDomainRedirectInfoWithDefaults() *DomainRedirectInfo`

NewDomainRedirectInfoWithDefaults instantiates a new DomainRedirectInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomainId

`func (o *DomainRedirectInfo) GetDomainId() string`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *DomainRedirectInfo) GetDomainIdOk() (*string, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *DomainRedirectInfo) SetDomainId(v string)`

SetDomainId sets DomainId field to given value.

### HasDomainId

`func (o *DomainRedirectInfo) HasDomainId() bool`

HasDomainId returns a boolean if a field has been set.

### GetDomain

`func (o *DomainRedirectInfo) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *DomainRedirectInfo) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *DomainRedirectInfo) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *DomainRedirectInfo) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetUrl

`func (o *DomainRedirectInfo) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *DomainRedirectInfo) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *DomainRedirectInfo) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *DomainRedirectInfo) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetType

`func (o *DomainRedirectInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DomainRedirectInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DomainRedirectInfo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DomainRedirectInfo) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAddress

`func (o *DomainRedirectInfo) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *DomainRedirectInfo) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *DomainRedirectInfo) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *DomainRedirectInfo) HasAddress() bool`

HasAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


