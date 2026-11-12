# GetDomainRedirectDetailsData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**DomainId** | Pointer to **string** | Domain service id | [optional] 
**Domain** | Pointer to **string** | Domain name | [optional] 
**Url** | Pointer to **string** | Redirect target URL when configured | [optional] 
**Type** | Pointer to **string** | Redirect type when configured | [optional] 
**Address** | Pointer to **string** | Optional redirector address override | [optional] 
**Dns** | Pointer to [**DomainRedirectDnsStatus**](DomainRedirectDnsStatus.md) |  | [optional] 

## Methods

### NewGetDomainRedirectDetailsData

`func NewGetDomainRedirectDetailsData(message string, ) *GetDomainRedirectDetailsData`

NewGetDomainRedirectDetailsData instantiates a new GetDomainRedirectDetailsData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDomainRedirectDetailsDataWithDefaults

`func NewGetDomainRedirectDetailsDataWithDefaults() *GetDomainRedirectDetailsData`

NewGetDomainRedirectDetailsDataWithDefaults instantiates a new GetDomainRedirectDetailsData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *GetDomainRedirectDetailsData) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetDomainRedirectDetailsData) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetDomainRedirectDetailsData) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetDomainId

`func (o *GetDomainRedirectDetailsData) GetDomainId() string`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *GetDomainRedirectDetailsData) GetDomainIdOk() (*string, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *GetDomainRedirectDetailsData) SetDomainId(v string)`

SetDomainId sets DomainId field to given value.

### HasDomainId

`func (o *GetDomainRedirectDetailsData) HasDomainId() bool`

HasDomainId returns a boolean if a field has been set.

### GetDomain

`func (o *GetDomainRedirectDetailsData) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *GetDomainRedirectDetailsData) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *GetDomainRedirectDetailsData) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *GetDomainRedirectDetailsData) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetUrl

`func (o *GetDomainRedirectDetailsData) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *GetDomainRedirectDetailsData) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *GetDomainRedirectDetailsData) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *GetDomainRedirectDetailsData) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetType

`func (o *GetDomainRedirectDetailsData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetDomainRedirectDetailsData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetDomainRedirectDetailsData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetDomainRedirectDetailsData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAddress

`func (o *GetDomainRedirectDetailsData) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *GetDomainRedirectDetailsData) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *GetDomainRedirectDetailsData) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *GetDomainRedirectDetailsData) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetDns

`func (o *GetDomainRedirectDetailsData) GetDns() DomainRedirectDnsStatus`

GetDns returns the Dns field if non-nil, zero value otherwise.

### GetDnsOk

`func (o *GetDomainRedirectDetailsData) GetDnsOk() (*DomainRedirectDnsStatus, bool)`

GetDnsOk returns a tuple with the Dns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns

`func (o *GetDomainRedirectDetailsData) SetDns(v DomainRedirectDnsStatus)`

SetDns sets Dns field to given value.

### HasDns

`func (o *GetDomainRedirectDetailsData) HasDns() bool`

HasDns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


