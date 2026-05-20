# UpdateNotificationResponseContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**OperationStatus**](OperationStatus.md) |  | 
**Data** | [**NotificationCreateResponseData**](NotificationCreateResponseData.md) |  | 

## Methods

### NewUpdateNotificationResponseContent

`func NewUpdateNotificationResponseContent(status OperationStatus, data NotificationCreateResponseData, ) *UpdateNotificationResponseContent`

NewUpdateNotificationResponseContent instantiates a new UpdateNotificationResponseContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNotificationResponseContentWithDefaults

`func NewUpdateNotificationResponseContentWithDefaults() *UpdateNotificationResponseContent`

NewUpdateNotificationResponseContentWithDefaults instantiates a new UpdateNotificationResponseContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *UpdateNotificationResponseContent) GetStatus() OperationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateNotificationResponseContent) GetStatusOk() (*OperationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateNotificationResponseContent) SetStatus(v OperationStatus)`

SetStatus sets Status field to given value.


### GetData

`func (o *UpdateNotificationResponseContent) GetData() NotificationCreateResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UpdateNotificationResponseContent) GetDataOk() (*NotificationCreateResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UpdateNotificationResponseContent) SetData(v NotificationCreateResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


