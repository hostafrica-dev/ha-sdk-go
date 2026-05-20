# PublicSshKeyResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Status message indicating the result | 
**Data** | [**SshKeyDetails**](SshKeyDetails.md) |  | 

## Methods

### NewPublicSshKeyResponseData

`func NewPublicSshKeyResponseData(message string, data SshKeyDetails, ) *PublicSshKeyResponseData`

NewPublicSshKeyResponseData instantiates a new PublicSshKeyResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicSshKeyResponseDataWithDefaults

`func NewPublicSshKeyResponseDataWithDefaults() *PublicSshKeyResponseData`

NewPublicSshKeyResponseDataWithDefaults instantiates a new PublicSshKeyResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *PublicSshKeyResponseData) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PublicSshKeyResponseData) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PublicSshKeyResponseData) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetData

`func (o *PublicSshKeyResponseData) GetData() SshKeyDetails`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PublicSshKeyResponseData) GetDataOk() (*SshKeyDetails, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PublicSshKeyResponseData) SetData(v SshKeyDetails)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


