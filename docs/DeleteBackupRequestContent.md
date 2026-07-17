# DeleteBackupRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceId** | **string** | Service ID - must be sent as a string | 
**BackupId** | **string** | Backup ID to delete | 

## Methods

### NewDeleteBackupRequestContent

`func NewDeleteBackupRequestContent(serviceId string, backupId string, ) *DeleteBackupRequestContent`

NewDeleteBackupRequestContent instantiates a new DeleteBackupRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteBackupRequestContentWithDefaults

`func NewDeleteBackupRequestContentWithDefaults() *DeleteBackupRequestContent`

NewDeleteBackupRequestContentWithDefaults instantiates a new DeleteBackupRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceId

`func (o *DeleteBackupRequestContent) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *DeleteBackupRequestContent) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *DeleteBackupRequestContent) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetBackupId

`func (o *DeleteBackupRequestContent) GetBackupId() string`

GetBackupId returns the BackupId field if non-nil, zero value otherwise.

### GetBackupIdOk

`func (o *DeleteBackupRequestContent) GetBackupIdOk() (*string, bool)`

GetBackupIdOk returns a tuple with the BackupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupId

`func (o *DeleteBackupRequestContent) SetBackupId(v string)`

SetBackupId sets BackupId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


