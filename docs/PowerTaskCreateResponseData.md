# PowerTaskCreateResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Status message indicating the result | 
**Task** | [**PowerTask**](PowerTask.md) |  | 

## Methods

### NewPowerTaskCreateResponseData

`func NewPowerTaskCreateResponseData(message string, task PowerTask, ) *PowerTaskCreateResponseData`

NewPowerTaskCreateResponseData instantiates a new PowerTaskCreateResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPowerTaskCreateResponseDataWithDefaults

`func NewPowerTaskCreateResponseDataWithDefaults() *PowerTaskCreateResponseData`

NewPowerTaskCreateResponseDataWithDefaults instantiates a new PowerTaskCreateResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *PowerTaskCreateResponseData) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PowerTaskCreateResponseData) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PowerTaskCreateResponseData) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetTask

`func (o *PowerTaskCreateResponseData) GetTask() PowerTask`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *PowerTaskCreateResponseData) GetTaskOk() (*PowerTask, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *PowerTaskCreateResponseData) SetTask(v PowerTask)`

SetTask sets Task field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


