# ListNotificationsResponseContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**OperationStatus**](OperationStatus.md) |  | 
**Data** | [**NotificationListResponseData**](NotificationListResponseData.md) |  | 

## Methods

### NewListNotificationsResponseContent

`func NewListNotificationsResponseContent(status OperationStatus, data NotificationListResponseData, ) *ListNotificationsResponseContent`

NewListNotificationsResponseContent instantiates a new ListNotificationsResponseContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListNotificationsResponseContentWithDefaults

`func NewListNotificationsResponseContentWithDefaults() *ListNotificationsResponseContent`

NewListNotificationsResponseContentWithDefaults instantiates a new ListNotificationsResponseContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ListNotificationsResponseContent) GetStatus() OperationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListNotificationsResponseContent) GetStatusOk() (*OperationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListNotificationsResponseContent) SetStatus(v OperationStatus)`

SetStatus sets Status field to given value.


### GetData

`func (o *ListNotificationsResponseContent) GetData() NotificationListResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListNotificationsResponseContent) GetDataOk() (*NotificationListResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListNotificationsResponseContent) SetData(v NotificationListResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


