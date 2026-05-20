# CreateRdnsRecordResponseContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**OperationStatus**](OperationStatus.md) |  | 
**Data** | [**CreateRdnsRecordResponseData**](CreateRdnsRecordResponseData.md) |  | 

## Methods

### NewCreateRdnsRecordResponseContent

`func NewCreateRdnsRecordResponseContent(status OperationStatus, data CreateRdnsRecordResponseData, ) *CreateRdnsRecordResponseContent`

NewCreateRdnsRecordResponseContent instantiates a new CreateRdnsRecordResponseContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRdnsRecordResponseContentWithDefaults

`func NewCreateRdnsRecordResponseContentWithDefaults() *CreateRdnsRecordResponseContent`

NewCreateRdnsRecordResponseContentWithDefaults instantiates a new CreateRdnsRecordResponseContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *CreateRdnsRecordResponseContent) GetStatus() OperationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateRdnsRecordResponseContent) GetStatusOk() (*OperationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateRdnsRecordResponseContent) SetStatus(v OperationStatus)`

SetStatus sets Status field to given value.


### GetData

`func (o *CreateRdnsRecordResponseContent) GetData() CreateRdnsRecordResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateRdnsRecordResponseContent) GetDataOk() (*CreateRdnsRecordResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateRdnsRecordResponseContent) SetData(v CreateRdnsRecordResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


