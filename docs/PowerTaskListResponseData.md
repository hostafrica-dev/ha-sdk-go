# PowerTaskListResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Status message indicating the result | 
**Tasks** | [**[]PowerTask**](PowerTask.md) | List of power tasks | 
**AvailableActions** | **[]string** | Available power actions | 
**AvailableJobTypes** | **[]string** | Available job types | 
**DialogRules** | [**PowerTaskDialogRules**](PowerTaskDialogRules.md) |  | 

## Methods

### NewPowerTaskListResponseData

`func NewPowerTaskListResponseData(message string, tasks []PowerTask, availableActions []string, availableJobTypes []string, dialogRules PowerTaskDialogRules, ) *PowerTaskListResponseData`

NewPowerTaskListResponseData instantiates a new PowerTaskListResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPowerTaskListResponseDataWithDefaults

`func NewPowerTaskListResponseDataWithDefaults() *PowerTaskListResponseData`

NewPowerTaskListResponseDataWithDefaults instantiates a new PowerTaskListResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *PowerTaskListResponseData) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PowerTaskListResponseData) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PowerTaskListResponseData) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetTasks

`func (o *PowerTaskListResponseData) GetTasks() []PowerTask`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *PowerTaskListResponseData) GetTasksOk() (*[]PowerTask, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *PowerTaskListResponseData) SetTasks(v []PowerTask)`

SetTasks sets Tasks field to given value.


### GetAvailableActions

`func (o *PowerTaskListResponseData) GetAvailableActions() []string`

GetAvailableActions returns the AvailableActions field if non-nil, zero value otherwise.

### GetAvailableActionsOk

`func (o *PowerTaskListResponseData) GetAvailableActionsOk() (*[]string, bool)`

GetAvailableActionsOk returns a tuple with the AvailableActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableActions

`func (o *PowerTaskListResponseData) SetAvailableActions(v []string)`

SetAvailableActions sets AvailableActions field to given value.


### GetAvailableJobTypes

`func (o *PowerTaskListResponseData) GetAvailableJobTypes() []string`

GetAvailableJobTypes returns the AvailableJobTypes field if non-nil, zero value otherwise.

### GetAvailableJobTypesOk

`func (o *PowerTaskListResponseData) GetAvailableJobTypesOk() (*[]string, bool)`

GetAvailableJobTypesOk returns a tuple with the AvailableJobTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableJobTypes

`func (o *PowerTaskListResponseData) SetAvailableJobTypes(v []string)`

SetAvailableJobTypes sets AvailableJobTypes field to given value.


### GetDialogRules

`func (o *PowerTaskListResponseData) GetDialogRules() PowerTaskDialogRules`

GetDialogRules returns the DialogRules field if non-nil, zero value otherwise.

### GetDialogRulesOk

`func (o *PowerTaskListResponseData) GetDialogRulesOk() (*PowerTaskDialogRules, bool)`

GetDialogRulesOk returns a tuple with the DialogRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDialogRules

`func (o *PowerTaskListResponseData) SetDialogRules(v PowerTaskDialogRules)`

SetDialogRules sets DialogRules field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


