# NotificationCreateResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Status message indicating the result | 
**Notification** | [**Notification**](Notification.md) |  | 

## Methods

### NewNotificationCreateResponseData

`func NewNotificationCreateResponseData(message string, notification Notification, ) *NotificationCreateResponseData`

NewNotificationCreateResponseData instantiates a new NotificationCreateResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationCreateResponseDataWithDefaults

`func NewNotificationCreateResponseDataWithDefaults() *NotificationCreateResponseData`

NewNotificationCreateResponseDataWithDefaults instantiates a new NotificationCreateResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *NotificationCreateResponseData) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *NotificationCreateResponseData) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *NotificationCreateResponseData) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetNotification

`func (o *NotificationCreateResponseData) GetNotification() Notification`

GetNotification returns the Notification field if non-nil, zero value otherwise.

### GetNotificationOk

`func (o *NotificationCreateResponseData) GetNotificationOk() (*Notification, bool)`

GetNotificationOk returns a tuple with the Notification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotification

`func (o *NotificationCreateResponseData) SetNotification(v Notification)`

SetNotification sets Notification field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


