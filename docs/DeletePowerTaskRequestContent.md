# DeletePowerTaskRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceId** | **string** | Service ID - must be sent as a string | 
**TaskId** | **int32** | Power task ID to delete | 

## Methods

### NewDeletePowerTaskRequestContent

`func NewDeletePowerTaskRequestContent(serviceId string, taskId int32, ) *DeletePowerTaskRequestContent`

NewDeletePowerTaskRequestContent instantiates a new DeletePowerTaskRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeletePowerTaskRequestContentWithDefaults

`func NewDeletePowerTaskRequestContentWithDefaults() *DeletePowerTaskRequestContent`

NewDeletePowerTaskRequestContentWithDefaults instantiates a new DeletePowerTaskRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceId

`func (o *DeletePowerTaskRequestContent) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *DeletePowerTaskRequestContent) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *DeletePowerTaskRequestContent) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetTaskId

`func (o *DeletePowerTaskRequestContent) GetTaskId() int32`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *DeletePowerTaskRequestContent) GetTaskIdOk() (*int32, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *DeletePowerTaskRequestContent) SetTaskId(v int32)`

SetTaskId sets TaskId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


