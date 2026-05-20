# ListIsosResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Status message indicating the result | 
**Isos** | **[]string** | List of available ISO images | 

## Methods

### NewListIsosResponseData

`func NewListIsosResponseData(message string, isos []string, ) *ListIsosResponseData`

NewListIsosResponseData instantiates a new ListIsosResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListIsosResponseDataWithDefaults

`func NewListIsosResponseDataWithDefaults() *ListIsosResponseData`

NewListIsosResponseDataWithDefaults instantiates a new ListIsosResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ListIsosResponseData) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ListIsosResponseData) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ListIsosResponseData) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetIsos

`func (o *ListIsosResponseData) GetIsos() []string`

GetIsos returns the Isos field if non-nil, zero value otherwise.

### GetIsosOk

`func (o *ListIsosResponseData) GetIsosOk() (*[]string, bool)`

GetIsosOk returns a tuple with the Isos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsos

`func (o *ListIsosResponseData) SetIsos(v []string)`

SetIsos sets Isos field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


