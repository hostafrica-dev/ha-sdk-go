# NotificationListResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Status message indicating the result | 
**Notifications** | [**[]Notification**](Notification.md) | List of notifications | 
**DialogRules** | [**NotificationDialogRules**](NotificationDialogRules.md) |  | 

## Methods

### NewNotificationListResponseData

`func NewNotificationListResponseData(message string, notifications []Notification, dialogRules NotificationDialogRules, ) *NotificationListResponseData`

NewNotificationListResponseData instantiates a new NotificationListResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationListResponseDataWithDefaults

`func NewNotificationListResponseDataWithDefaults() *NotificationListResponseData`

NewNotificationListResponseDataWithDefaults instantiates a new NotificationListResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *NotificationListResponseData) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *NotificationListResponseData) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *NotificationListResponseData) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetNotifications

`func (o *NotificationListResponseData) GetNotifications() []Notification`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *NotificationListResponseData) GetNotificationsOk() (*[]Notification, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *NotificationListResponseData) SetNotifications(v []Notification)`

SetNotifications sets Notifications field to given value.


### GetDialogRules

`func (o *NotificationListResponseData) GetDialogRules() NotificationDialogRules`

GetDialogRules returns the DialogRules field if non-nil, zero value otherwise.

### GetDialogRulesOk

`func (o *NotificationListResponseData) GetDialogRulesOk() (*NotificationDialogRules, bool)`

GetDialogRulesOk returns a tuple with the DialogRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDialogRules

`func (o *NotificationListResponseData) SetDialogRules(v NotificationDialogRules)`

SetDialogRules sets DialogRules field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


