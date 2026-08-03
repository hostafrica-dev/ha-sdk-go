# GetDomainData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Status message indicating the result | 
**Domain** | [**DomainDetail**](DomainDetail.md) |  | 

## Methods

### NewGetDomainData

`func NewGetDomainData(message string, domain DomainDetail, ) *GetDomainData`

NewGetDomainData instantiates a new GetDomainData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDomainDataWithDefaults

`func NewGetDomainDataWithDefaults() *GetDomainData`

NewGetDomainDataWithDefaults instantiates a new GetDomainData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *GetDomainData) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetDomainData) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetDomainData) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetDomain

`func (o *GetDomainData) GetDomain() DomainDetail`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *GetDomainData) GetDomainOk() (*DomainDetail, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *GetDomainData) SetDomain(v DomainDetail)`

SetDomain sets Domain field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


