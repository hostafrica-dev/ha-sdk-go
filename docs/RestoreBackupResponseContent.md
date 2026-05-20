# RestoreBackupResponseContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**OperationStatus**](OperationStatus.md) |  | 
**Data** | [**VpsSimpleActionResponseData**](VpsSimpleActionResponseData.md) |  | 

## Methods

### NewRestoreBackupResponseContent

`func NewRestoreBackupResponseContent(status OperationStatus, data VpsSimpleActionResponseData, ) *RestoreBackupResponseContent`

NewRestoreBackupResponseContent instantiates a new RestoreBackupResponseContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreBackupResponseContentWithDefaults

`func NewRestoreBackupResponseContentWithDefaults() *RestoreBackupResponseContent`

NewRestoreBackupResponseContentWithDefaults instantiates a new RestoreBackupResponseContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *RestoreBackupResponseContent) GetStatus() OperationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RestoreBackupResponseContent) GetStatusOk() (*OperationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RestoreBackupResponseContent) SetStatus(v OperationStatus)`

SetStatus sets Status field to given value.


### GetData

`func (o *RestoreBackupResponseContent) GetData() VpsSimpleActionResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RestoreBackupResponseContent) GetDataOk() (*VpsSimpleActionResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RestoreBackupResponseContent) SetData(v VpsSimpleActionResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


