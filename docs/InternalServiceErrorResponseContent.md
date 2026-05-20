# InternalServiceErrorResponseContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**Data** | **map[string]interface{}** | Empty object for operations that don&#39;t return data | 

## Methods

### NewInternalServiceErrorResponseContent

`func NewInternalServiceErrorResponseContent(message string, data map[string]interface{}, ) *InternalServiceErrorResponseContent`

NewInternalServiceErrorResponseContent instantiates a new InternalServiceErrorResponseContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalServiceErrorResponseContentWithDefaults

`func NewInternalServiceErrorResponseContentWithDefaults() *InternalServiceErrorResponseContent`

NewInternalServiceErrorResponseContentWithDefaults instantiates a new InternalServiceErrorResponseContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *InternalServiceErrorResponseContent) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *InternalServiceErrorResponseContent) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *InternalServiceErrorResponseContent) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetData

`func (o *InternalServiceErrorResponseContent) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *InternalServiceErrorResponseContent) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *InternalServiceErrorResponseContent) SetData(v map[string]interface{})`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


