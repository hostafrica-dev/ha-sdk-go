# ListDomainRedirectsResponseContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**OperationStatus**](OperationStatus.md) |  | 
**Data** | [**ListDomainRedirectsData**](ListDomainRedirectsData.md) |  | 

## Methods

### NewListDomainRedirectsResponseContent

`func NewListDomainRedirectsResponseContent(status OperationStatus, data ListDomainRedirectsData, ) *ListDomainRedirectsResponseContent`

NewListDomainRedirectsResponseContent instantiates a new ListDomainRedirectsResponseContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListDomainRedirectsResponseContentWithDefaults

`func NewListDomainRedirectsResponseContentWithDefaults() *ListDomainRedirectsResponseContent`

NewListDomainRedirectsResponseContentWithDefaults instantiates a new ListDomainRedirectsResponseContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ListDomainRedirectsResponseContent) GetStatus() OperationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListDomainRedirectsResponseContent) GetStatusOk() (*OperationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListDomainRedirectsResponseContent) SetStatus(v OperationStatus)`

SetStatus sets Status field to given value.


### GetData

`func (o *ListDomainRedirectsResponseContent) GetData() ListDomainRedirectsData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListDomainRedirectsResponseContent) GetDataOk() (*ListDomainRedirectsData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListDomainRedirectsResponseContent) SetData(v ListDomainRedirectsData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


