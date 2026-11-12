# AddDomainRedirectRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainId** | **string** | Domain service id - must be sent as a string | 
**Url** | **string** | Redirect target URL (absolute http(s) URL) | 
**Type** | **string** | Redirect type: 301, 302, permanent, or temporary | 
**Address** | Pointer to **string** | Optional redirector address override (passed through when present) | [optional] 
**Override** | Pointer to **bool** | When true, force-replace conflicting @/www A records that do not point at the redirector | [optional] 

## Methods

### NewAddDomainRedirectRequestContent

`func NewAddDomainRedirectRequestContent(domainId string, url string, type_ string, ) *AddDomainRedirectRequestContent`

NewAddDomainRedirectRequestContent instantiates a new AddDomainRedirectRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddDomainRedirectRequestContentWithDefaults

`func NewAddDomainRedirectRequestContentWithDefaults() *AddDomainRedirectRequestContent`

NewAddDomainRedirectRequestContentWithDefaults instantiates a new AddDomainRedirectRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomainId

`func (o *AddDomainRedirectRequestContent) GetDomainId() string`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *AddDomainRedirectRequestContent) GetDomainIdOk() (*string, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *AddDomainRedirectRequestContent) SetDomainId(v string)`

SetDomainId sets DomainId field to given value.


### GetUrl

`func (o *AddDomainRedirectRequestContent) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AddDomainRedirectRequestContent) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AddDomainRedirectRequestContent) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetType

`func (o *AddDomainRedirectRequestContent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddDomainRedirectRequestContent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddDomainRedirectRequestContent) SetType(v string)`

SetType sets Type field to given value.


### GetAddress

`func (o *AddDomainRedirectRequestContent) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *AddDomainRedirectRequestContent) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *AddDomainRedirectRequestContent) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *AddDomainRedirectRequestContent) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetOverride

`func (o *AddDomainRedirectRequestContent) GetOverride() bool`

GetOverride returns the Override field if non-nil, zero value otherwise.

### GetOverrideOk

`func (o *AddDomainRedirectRequestContent) GetOverrideOk() (*bool, bool)`

GetOverrideOk returns a tuple with the Override field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverride

`func (o *AddDomainRedirectRequestContent) SetOverride(v bool)`

SetOverride sets Override field to given value.

### HasOverride

`func (o *AddDomainRedirectRequestContent) HasOverride() bool`

HasOverride returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


