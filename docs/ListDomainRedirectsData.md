# ListDomainRedirectsData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**Redirects** | [**[]DomainRedirectInfo**](DomainRedirectInfo.md) |  | 
**Limit** | Pointer to **int32** | Pagination limit | [optional] 
**Offset** | Pointer to **int32** | Pagination offset | [optional] 

## Methods

### NewListDomainRedirectsData

`func NewListDomainRedirectsData(message string, redirects []DomainRedirectInfo, ) *ListDomainRedirectsData`

NewListDomainRedirectsData instantiates a new ListDomainRedirectsData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListDomainRedirectsDataWithDefaults

`func NewListDomainRedirectsDataWithDefaults() *ListDomainRedirectsData`

NewListDomainRedirectsDataWithDefaults instantiates a new ListDomainRedirectsData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ListDomainRedirectsData) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ListDomainRedirectsData) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ListDomainRedirectsData) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetRedirects

`func (o *ListDomainRedirectsData) GetRedirects() []DomainRedirectInfo`

GetRedirects returns the Redirects field if non-nil, zero value otherwise.

### GetRedirectsOk

`func (o *ListDomainRedirectsData) GetRedirectsOk() (*[]DomainRedirectInfo, bool)`

GetRedirectsOk returns a tuple with the Redirects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirects

`func (o *ListDomainRedirectsData) SetRedirects(v []DomainRedirectInfo)`

SetRedirects sets Redirects field to given value.


### GetLimit

`func (o *ListDomainRedirectsData) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ListDomainRedirectsData) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ListDomainRedirectsData) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ListDomainRedirectsData) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *ListDomainRedirectsData) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ListDomainRedirectsData) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ListDomainRedirectsData) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *ListDomainRedirectsData) HasOffset() bool`

HasOffset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


