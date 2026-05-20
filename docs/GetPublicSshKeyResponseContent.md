# GetPublicSshKeyResponseContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**OperationStatus**](OperationStatus.md) |  | 
**Data** | [**PublicSshKeyResponseData**](PublicSshKeyResponseData.md) |  | 

## Methods

### NewGetPublicSshKeyResponseContent

`func NewGetPublicSshKeyResponseContent(status OperationStatus, data PublicSshKeyResponseData, ) *GetPublicSshKeyResponseContent`

NewGetPublicSshKeyResponseContent instantiates a new GetPublicSshKeyResponseContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPublicSshKeyResponseContentWithDefaults

`func NewGetPublicSshKeyResponseContentWithDefaults() *GetPublicSshKeyResponseContent`

NewGetPublicSshKeyResponseContentWithDefaults instantiates a new GetPublicSshKeyResponseContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetPublicSshKeyResponseContent) GetStatus() OperationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetPublicSshKeyResponseContent) GetStatusOk() (*OperationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetPublicSshKeyResponseContent) SetStatus(v OperationStatus)`

SetStatus sets Status field to given value.


### GetData

`func (o *GetPublicSshKeyResponseContent) GetData() PublicSshKeyResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetPublicSshKeyResponseContent) GetDataOk() (*PublicSshKeyResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetPublicSshKeyResponseContent) SetData(v PublicSshKeyResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


