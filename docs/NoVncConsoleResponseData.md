# NoVncConsoleResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Status message indicating the result | 
**Console** | [**NoVncConsoleDetails**](NoVncConsoleDetails.md) |  | 

## Methods

### NewNoVncConsoleResponseData

`func NewNoVncConsoleResponseData(message string, console NoVncConsoleDetails, ) *NoVncConsoleResponseData`

NewNoVncConsoleResponseData instantiates a new NoVncConsoleResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNoVncConsoleResponseDataWithDefaults

`func NewNoVncConsoleResponseDataWithDefaults() *NoVncConsoleResponseData`

NewNoVncConsoleResponseDataWithDefaults instantiates a new NoVncConsoleResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *NoVncConsoleResponseData) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *NoVncConsoleResponseData) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *NoVncConsoleResponseData) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetConsole

`func (o *NoVncConsoleResponseData) GetConsole() NoVncConsoleDetails`

GetConsole returns the Console field if non-nil, zero value otherwise.

### GetConsoleOk

`func (o *NoVncConsoleResponseData) GetConsoleOk() (*NoVncConsoleDetails, bool)`

GetConsoleOk returns a tuple with the Console field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsole

`func (o *NoVncConsoleResponseData) SetConsole(v NoVncConsoleDetails)`

SetConsole sets Console field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


