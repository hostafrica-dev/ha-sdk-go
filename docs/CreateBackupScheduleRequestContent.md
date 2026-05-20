# CreateBackupScheduleRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceId** | **string** | Service ID - must be sent as a string | 
**Starttime** | **string** | Time in HH:MM format (e.g., &#39;03:00&#39;). Hours must be between 00-23, minutes must be between 00-59 | 
**Dow** | [**[]DayOfWeek**](DayOfWeek.md) | Days of week when backup should run | 
**Compress** | [**CompressionType**](CompressionType.md) |  | 
**Mode** | [**BackupModeType**](BackupModeType.md) |  | 
**Mailto** | Pointer to **bool** | Email notification setting. Set to true to send notifications to client&#39;s email, false or omit to disable | [optional] 

## Methods

### NewCreateBackupScheduleRequestContent

`func NewCreateBackupScheduleRequestContent(serviceId string, starttime string, dow []DayOfWeek, compress CompressionType, mode BackupModeType, ) *CreateBackupScheduleRequestContent`

NewCreateBackupScheduleRequestContent instantiates a new CreateBackupScheduleRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBackupScheduleRequestContentWithDefaults

`func NewCreateBackupScheduleRequestContentWithDefaults() *CreateBackupScheduleRequestContent`

NewCreateBackupScheduleRequestContentWithDefaults instantiates a new CreateBackupScheduleRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceId

`func (o *CreateBackupScheduleRequestContent) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *CreateBackupScheduleRequestContent) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *CreateBackupScheduleRequestContent) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetStarttime

`func (o *CreateBackupScheduleRequestContent) GetStarttime() string`

GetStarttime returns the Starttime field if non-nil, zero value otherwise.

### GetStarttimeOk

`func (o *CreateBackupScheduleRequestContent) GetStarttimeOk() (*string, bool)`

GetStarttimeOk returns a tuple with the Starttime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarttime

`func (o *CreateBackupScheduleRequestContent) SetStarttime(v string)`

SetStarttime sets Starttime field to given value.


### GetDow

`func (o *CreateBackupScheduleRequestContent) GetDow() []DayOfWeek`

GetDow returns the Dow field if non-nil, zero value otherwise.

### GetDowOk

`func (o *CreateBackupScheduleRequestContent) GetDowOk() (*[]DayOfWeek, bool)`

GetDowOk returns a tuple with the Dow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDow

`func (o *CreateBackupScheduleRequestContent) SetDow(v []DayOfWeek)`

SetDow sets Dow field to given value.


### GetCompress

`func (o *CreateBackupScheduleRequestContent) GetCompress() CompressionType`

GetCompress returns the Compress field if non-nil, zero value otherwise.

### GetCompressOk

`func (o *CreateBackupScheduleRequestContent) GetCompressOk() (*CompressionType, bool)`

GetCompressOk returns a tuple with the Compress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompress

`func (o *CreateBackupScheduleRequestContent) SetCompress(v CompressionType)`

SetCompress sets Compress field to given value.


### GetMode

`func (o *CreateBackupScheduleRequestContent) GetMode() BackupModeType`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *CreateBackupScheduleRequestContent) GetModeOk() (*BackupModeType, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *CreateBackupScheduleRequestContent) SetMode(v BackupModeType)`

SetMode sets Mode field to given value.


### GetMailto

`func (o *CreateBackupScheduleRequestContent) GetMailto() bool`

GetMailto returns the Mailto field if non-nil, zero value otherwise.

### GetMailtoOk

`func (o *CreateBackupScheduleRequestContent) GetMailtoOk() (*bool, bool)`

GetMailtoOk returns a tuple with the Mailto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailto

`func (o *CreateBackupScheduleRequestContent) SetMailto(v bool)`

SetMailto sets Mailto field to given value.

### HasMailto

`func (o *CreateBackupScheduleRequestContent) HasMailto() bool`

HasMailto returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


