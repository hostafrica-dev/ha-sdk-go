# ValidationErrorResponseContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**Data** | **map[string]interface{}** | Empty object for operations that don&#39;t return data | 

## Methods

### NewValidationErrorResponseContent

`func NewValidationErrorResponseContent(message string, data map[string]interface{}, ) *ValidationErrorResponseContent`

NewValidationErrorResponseContent instantiates a new ValidationErrorResponseContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidationErrorResponseContentWithDefaults

`func NewValidationErrorResponseContentWithDefaults() *ValidationErrorResponseContent`

NewValidationErrorResponseContentWithDefaults instantiates a new ValidationErrorResponseContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ValidationErrorResponseContent) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ValidationErrorResponseContent) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ValidationErrorResponseContent) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetData

`func (o *ValidationErrorResponseContent) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ValidationErrorResponseContent) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ValidationErrorResponseContent) SetData(v map[string]interface{})`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


