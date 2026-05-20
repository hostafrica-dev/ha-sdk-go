# ListReinstallOsResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Status message indicating the result | 
**Templates** | **[]string** | List of available OS templates | 

## Methods

### NewListReinstallOsResponseData

`func NewListReinstallOsResponseData(message string, templates []string, ) *ListReinstallOsResponseData`

NewListReinstallOsResponseData instantiates a new ListReinstallOsResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListReinstallOsResponseDataWithDefaults

`func NewListReinstallOsResponseDataWithDefaults() *ListReinstallOsResponseData`

NewListReinstallOsResponseDataWithDefaults instantiates a new ListReinstallOsResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ListReinstallOsResponseData) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ListReinstallOsResponseData) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ListReinstallOsResponseData) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetTemplates

`func (o *ListReinstallOsResponseData) GetTemplates() []string`

GetTemplates returns the Templates field if non-nil, zero value otherwise.

### GetTemplatesOk

`func (o *ListReinstallOsResponseData) GetTemplatesOk() (*[]string, bool)`

GetTemplatesOk returns a tuple with the Templates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplates

`func (o *ListReinstallOsResponseData) SetTemplates(v []string)`

SetTemplates sets Templates field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


