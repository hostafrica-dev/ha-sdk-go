# UserChangePasswordResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** | Whether the operation was successful | 
**Message** | **string** | Status message indicating the result | 
**Data** | [**UserChangePasswordDetails**](UserChangePasswordDetails.md) |  | 

## Methods

### NewUserChangePasswordResponseData

`func NewUserChangePasswordResponseData(success bool, message string, data UserChangePasswordDetails, ) *UserChangePasswordResponseData`

NewUserChangePasswordResponseData instantiates a new UserChangePasswordResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserChangePasswordResponseDataWithDefaults

`func NewUserChangePasswordResponseDataWithDefaults() *UserChangePasswordResponseData`

NewUserChangePasswordResponseDataWithDefaults instantiates a new UserChangePasswordResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *UserChangePasswordResponseData) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *UserChangePasswordResponseData) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *UserChangePasswordResponseData) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetMessage

`func (o *UserChangePasswordResponseData) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *UserChangePasswordResponseData) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *UserChangePasswordResponseData) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetData

`func (o *UserChangePasswordResponseData) GetData() UserChangePasswordDetails`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UserChangePasswordResponseData) GetDataOk() (*UserChangePasswordDetails, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UserChangePasswordResponseData) SetData(v UserChangePasswordDetails)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


