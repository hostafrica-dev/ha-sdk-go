# StopVpsResponseContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**OperationStatus**](OperationStatus.md) |  | 
**Data** | [**VpsSimpleActionResponseData**](VpsSimpleActionResponseData.md) |  | 

## Methods

### NewStopVpsResponseContent

`func NewStopVpsResponseContent(status OperationStatus, data VpsSimpleActionResponseData, ) *StopVpsResponseContent`

NewStopVpsResponseContent instantiates a new StopVpsResponseContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStopVpsResponseContentWithDefaults

`func NewStopVpsResponseContentWithDefaults() *StopVpsResponseContent`

NewStopVpsResponseContentWithDefaults instantiates a new StopVpsResponseContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *StopVpsResponseContent) GetStatus() OperationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StopVpsResponseContent) GetStatusOk() (*OperationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StopVpsResponseContent) SetStatus(v OperationStatus)`

SetStatus sets Status field to given value.


### GetData

`func (o *StopVpsResponseContent) GetData() VpsSimpleActionResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *StopVpsResponseContent) GetDataOk() (*VpsSimpleActionResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *StopVpsResponseContent) SetData(v VpsSimpleActionResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


