# DeleteBackupScheduleRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceId** | **string** | Service ID - must be sent as a string | 
**ScheduleId** | **string** | Backup schedule ID to delete | 

## Methods

### NewDeleteBackupScheduleRequestContent

`func NewDeleteBackupScheduleRequestContent(serviceId string, scheduleId string, ) *DeleteBackupScheduleRequestContent`

NewDeleteBackupScheduleRequestContent instantiates a new DeleteBackupScheduleRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteBackupScheduleRequestContentWithDefaults

`func NewDeleteBackupScheduleRequestContentWithDefaults() *DeleteBackupScheduleRequestContent`

NewDeleteBackupScheduleRequestContentWithDefaults instantiates a new DeleteBackupScheduleRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceId

`func (o *DeleteBackupScheduleRequestContent) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *DeleteBackupScheduleRequestContent) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *DeleteBackupScheduleRequestContent) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetScheduleId

`func (o *DeleteBackupScheduleRequestContent) GetScheduleId() string`

GetScheduleId returns the ScheduleId field if non-nil, zero value otherwise.

### GetScheduleIdOk

`func (o *DeleteBackupScheduleRequestContent) GetScheduleIdOk() (*string, bool)`

GetScheduleIdOk returns a tuple with the ScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleId

`func (o *DeleteBackupScheduleRequestContent) SetScheduleId(v string)`

SetScheduleId sets ScheduleId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


