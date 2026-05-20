# BadRequestErrorResponseContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**Data** | **map[string]interface{}** | Empty object for operations that don&#39;t return data | 

## Methods

### NewBadRequestErrorResponseContent

`func NewBadRequestErrorResponseContent(message string, data map[string]interface{}, ) *BadRequestErrorResponseContent`

NewBadRequestErrorResponseContent instantiates a new BadRequestErrorResponseContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBadRequestErrorResponseContentWithDefaults

`func NewBadRequestErrorResponseContentWithDefaults() *BadRequestErrorResponseContent`

NewBadRequestErrorResponseContentWithDefaults instantiates a new BadRequestErrorResponseContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *BadRequestErrorResponseContent) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *BadRequestErrorResponseContent) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *BadRequestErrorResponseContent) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetData

`func (o *BadRequestErrorResponseContent) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BadRequestErrorResponseContent) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BadRequestErrorResponseContent) SetData(v map[string]interface{})`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


