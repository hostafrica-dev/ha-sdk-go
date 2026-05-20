# ListPowerTasksResponseContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**OperationStatus**](OperationStatus.md) |  | 
**Data** | [**PowerTaskListResponseData**](PowerTaskListResponseData.md) |  | 

## Methods

### NewListPowerTasksResponseContent

`func NewListPowerTasksResponseContent(status OperationStatus, data PowerTaskListResponseData, ) *ListPowerTasksResponseContent`

NewListPowerTasksResponseContent instantiates a new ListPowerTasksResponseContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListPowerTasksResponseContentWithDefaults

`func NewListPowerTasksResponseContentWithDefaults() *ListPowerTasksResponseContent`

NewListPowerTasksResponseContentWithDefaults instantiates a new ListPowerTasksResponseContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ListPowerTasksResponseContent) GetStatus() OperationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListPowerTasksResponseContent) GetStatusOk() (*OperationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListPowerTasksResponseContent) SetStatus(v OperationStatus)`

SetStatus sets Status field to given value.


### GetData

`func (o *ListPowerTasksResponseContent) GetData() PowerTaskListResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListPowerTasksResponseContent) GetDataOk() (*PowerTaskListResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListPowerTasksResponseContent) SetData(v PowerTaskListResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


