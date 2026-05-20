# ListRdnsRecordsResponseContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**OperationStatus**](OperationStatus.md) |  | 
**Data** | [**ListRdnsResponseData**](ListRdnsResponseData.md) |  | 

## Methods

### NewListRdnsRecordsResponseContent

`func NewListRdnsRecordsResponseContent(status OperationStatus, data ListRdnsResponseData, ) *ListRdnsRecordsResponseContent`

NewListRdnsRecordsResponseContent instantiates a new ListRdnsRecordsResponseContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListRdnsRecordsResponseContentWithDefaults

`func NewListRdnsRecordsResponseContentWithDefaults() *ListRdnsRecordsResponseContent`

NewListRdnsRecordsResponseContentWithDefaults instantiates a new ListRdnsRecordsResponseContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ListRdnsRecordsResponseContent) GetStatus() OperationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListRdnsRecordsResponseContent) GetStatusOk() (*OperationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListRdnsRecordsResponseContent) SetStatus(v OperationStatus)`

SetStatus sets Status field to given value.


### GetData

`func (o *ListRdnsRecordsResponseContent) GetData() ListRdnsResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListRdnsRecordsResponseContent) GetDataOk() (*ListRdnsResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListRdnsRecordsResponseContent) SetData(v ListRdnsResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


