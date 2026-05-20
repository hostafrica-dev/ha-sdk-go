# BackupSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Schedule ID | 
**Starttime** | **string** | Start time for backup (HH:MM format) | 
**Dow** | **string** | Days of week (comma-separated) | 
**Compress** | [**CompressionType**](CompressionType.md) |  | 
**Mode** | **string** | Backup mode | 
**Mailto** | Pointer to **string** | Email address for notifications (null if not set) | [optional] 

## Methods

### NewBackupSchedule

`func NewBackupSchedule(id string, starttime string, dow string, compress CompressionType, mode string, ) *BackupSchedule`

NewBackupSchedule instantiates a new BackupSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupScheduleWithDefaults

`func NewBackupScheduleWithDefaults() *BackupSchedule`

NewBackupScheduleWithDefaults instantiates a new BackupSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BackupSchedule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BackupSchedule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BackupSchedule) SetId(v string)`

SetId sets Id field to given value.


### GetStarttime

`func (o *BackupSchedule) GetStarttime() string`

GetStarttime returns the Starttime field if non-nil, zero value otherwise.

### GetStarttimeOk

`func (o *BackupSchedule) GetStarttimeOk() (*string, bool)`

GetStarttimeOk returns a tuple with the Starttime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarttime

`func (o *BackupSchedule) SetStarttime(v string)`

SetStarttime sets Starttime field to given value.


### GetDow

`func (o *BackupSchedule) GetDow() string`

GetDow returns the Dow field if non-nil, zero value otherwise.

### GetDowOk

`func (o *BackupSchedule) GetDowOk() (*string, bool)`

GetDowOk returns a tuple with the Dow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDow

`func (o *BackupSchedule) SetDow(v string)`

SetDow sets Dow field to given value.


### GetCompress

`func (o *BackupSchedule) GetCompress() CompressionType`

GetCompress returns the Compress field if non-nil, zero value otherwise.

### GetCompressOk

`func (o *BackupSchedule) GetCompressOk() (*CompressionType, bool)`

GetCompressOk returns a tuple with the Compress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompress

`func (o *BackupSchedule) SetCompress(v CompressionType)`

SetCompress sets Compress field to given value.


### GetMode

`func (o *BackupSchedule) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *BackupSchedule) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *BackupSchedule) SetMode(v string)`

SetMode sets Mode field to given value.


### GetMailto

`func (o *BackupSchedule) GetMailto() string`

GetMailto returns the Mailto field if non-nil, zero value otherwise.

### GetMailtoOk

`func (o *BackupSchedule) GetMailtoOk() (*string, bool)`

GetMailtoOk returns a tuple with the Mailto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailto

`func (o *BackupSchedule) SetMailto(v string)`

SetMailto sets Mailto field to given value.

### HasMailto

`func (o *BackupSchedule) HasMailto() bool`

HasMailto returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


