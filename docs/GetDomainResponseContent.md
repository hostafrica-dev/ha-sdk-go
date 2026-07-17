# GetDomainResponseContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**OperationStatus**](OperationStatus.md) |  | 
**Data** | [**GetDomainData**](GetDomainData.md) |  | 

## Methods

### NewGetDomainResponseContent

`func NewGetDomainResponseContent(status OperationStatus, data GetDomainData, ) *GetDomainResponseContent`

NewGetDomainResponseContent instantiates a new GetDomainResponseContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDomainResponseContentWithDefaults

`func NewGetDomainResponseContentWithDefaults() *GetDomainResponseContent`

NewGetDomainResponseContentWithDefaults instantiates a new GetDomainResponseContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetDomainResponseContent) GetStatus() OperationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetDomainResponseContent) GetStatusOk() (*OperationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetDomainResponseContent) SetStatus(v OperationStatus)`

SetStatus sets Status field to given value.


### GetData

`func (o *GetDomainResponseContent) GetData() GetDomainData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetDomainResponseContent) GetDataOk() (*GetDomainData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetDomainResponseContent) SetData(v GetDomainData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


