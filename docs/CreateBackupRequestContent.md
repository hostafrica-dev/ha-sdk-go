# CreateBackupRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceId** | **string** | Service ID - must be sent as a string | 
**Compress** | Pointer to [**CompressionType**](CompressionType.md) |  | [optional] 
**Mode** | Pointer to [**BackupModeType**](BackupModeType.md) |  | [optional] 

## Methods

### NewCreateBackupRequestContent

`func NewCreateBackupRequestContent(serviceId string, ) *CreateBackupRequestContent`

NewCreateBackupRequestContent instantiates a new CreateBackupRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBackupRequestContentWithDefaults

`func NewCreateBackupRequestContentWithDefaults() *CreateBackupRequestContent`

NewCreateBackupRequestContentWithDefaults instantiates a new CreateBackupRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceId

`func (o *CreateBackupRequestContent) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *CreateBackupRequestContent) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *CreateBackupRequestContent) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetCompress

`func (o *CreateBackupRequestContent) GetCompress() CompressionType`

GetCompress returns the Compress field if non-nil, zero value otherwise.

### GetCompressOk

`func (o *CreateBackupRequestContent) GetCompressOk() (*CompressionType, bool)`

GetCompressOk returns a tuple with the Compress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompress

`func (o *CreateBackupRequestContent) SetCompress(v CompressionType)`

SetCompress sets Compress field to given value.

### HasCompress

`func (o *CreateBackupRequestContent) HasCompress() bool`

HasCompress returns a boolean if a field has been set.

### GetMode

`func (o *CreateBackupRequestContent) GetMode() BackupModeType`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *CreateBackupRequestContent) GetModeOk() (*BackupModeType, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *CreateBackupRequestContent) SetMode(v BackupModeType)`

SetMode sets Mode field to given value.

### HasMode

`func (o *CreateBackupRequestContent) HasMode() bool`

HasMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


