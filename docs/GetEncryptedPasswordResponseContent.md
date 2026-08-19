# GetEncryptedPasswordResponseContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**OperationStatus**](OperationStatus.md) |  | 
**Data** | [**EncryptedPasswordResponseData**](EncryptedPasswordResponseData.md) |  | 

## Methods

### NewGetEncryptedPasswordResponseContent

`func NewGetEncryptedPasswordResponseContent(status OperationStatus, data EncryptedPasswordResponseData, ) *GetEncryptedPasswordResponseContent`

NewGetEncryptedPasswordResponseContent instantiates a new GetEncryptedPasswordResponseContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEncryptedPasswordResponseContentWithDefaults

`func NewGetEncryptedPasswordResponseContentWithDefaults() *GetEncryptedPasswordResponseContent`

NewGetEncryptedPasswordResponseContentWithDefaults instantiates a new GetEncryptedPasswordResponseContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetEncryptedPasswordResponseContent) GetStatus() OperationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetEncryptedPasswordResponseContent) GetStatusOk() (*OperationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetEncryptedPasswordResponseContent) SetStatus(v OperationStatus)`

SetStatus sets Status field to given value.


### GetData

`func (o *GetEncryptedPasswordResponseContent) GetData() EncryptedPasswordResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetEncryptedPasswordResponseContent) GetDataOk() (*EncryptedPasswordResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetEncryptedPasswordResponseContent) SetData(v EncryptedPasswordResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


