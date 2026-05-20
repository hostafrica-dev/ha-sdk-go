# BackupItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Backup ID | 
**Size** | **string** | Backup size in human-readable format | 
**CreateDate** | **string** | Backup creation date and time | 
**Format** | **string** | Backup format (e.g., vma, vma.lzo) | 
**Protected** | **bool** | Whether the backup is protected from deletion | 

## Methods

### NewBackupItem

`func NewBackupItem(id int32, size string, createDate string, format string, protected bool, ) *BackupItem`

NewBackupItem instantiates a new BackupItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupItemWithDefaults

`func NewBackupItemWithDefaults() *BackupItem`

NewBackupItemWithDefaults instantiates a new BackupItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BackupItem) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BackupItem) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BackupItem) SetId(v int32)`

SetId sets Id field to given value.


### GetSize

`func (o *BackupItem) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *BackupItem) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *BackupItem) SetSize(v string)`

SetSize sets Size field to given value.


### GetCreateDate

`func (o *BackupItem) GetCreateDate() string`

GetCreateDate returns the CreateDate field if non-nil, zero value otherwise.

### GetCreateDateOk

`func (o *BackupItem) GetCreateDateOk() (*string, bool)`

GetCreateDateOk returns a tuple with the CreateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDate

`func (o *BackupItem) SetCreateDate(v string)`

SetCreateDate sets CreateDate field to given value.


### GetFormat

`func (o *BackupItem) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *BackupItem) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *BackupItem) SetFormat(v string)`

SetFormat sets Format field to given value.


### GetProtected

`func (o *BackupItem) GetProtected() bool`

GetProtected returns the Protected field if non-nil, zero value otherwise.

### GetProtectedOk

`func (o *BackupItem) GetProtectedOk() (*bool, bool)`

GetProtectedOk returns a tuple with the Protected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtected

`func (o *BackupItem) SetProtected(v bool)`

SetProtected sets Protected field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


