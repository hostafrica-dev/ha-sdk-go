# DeleteNotificationRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceId** | **string** | Service ID - must be sent as a string | 
**NotificationId** | **int32** | Notification ID to delete | 

## Methods

### NewDeleteNotificationRequestContent

`func NewDeleteNotificationRequestContent(serviceId string, notificationId int32, ) *DeleteNotificationRequestContent`

NewDeleteNotificationRequestContent instantiates a new DeleteNotificationRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteNotificationRequestContentWithDefaults

`func NewDeleteNotificationRequestContentWithDefaults() *DeleteNotificationRequestContent`

NewDeleteNotificationRequestContentWithDefaults instantiates a new DeleteNotificationRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceId

`func (o *DeleteNotificationRequestContent) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *DeleteNotificationRequestContent) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *DeleteNotificationRequestContent) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetNotificationId

`func (o *DeleteNotificationRequestContent) GetNotificationId() int32`

GetNotificationId returns the NotificationId field if non-nil, zero value otherwise.

### GetNotificationIdOk

`func (o *DeleteNotificationRequestContent) GetNotificationIdOk() (*int32, bool)`

GetNotificationIdOk returns a tuple with the NotificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationId

`func (o *DeleteNotificationRequestContent) SetNotificationId(v int32)`

SetNotificationId sets NotificationId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


