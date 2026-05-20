# SnapshotItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Snapshot name | 
**Description** | **string** | Snapshot description | 
**Snaptime** | Pointer to **string** | Snapshot creation date | [optional] 
**Size** | Pointer to **int64** | Snapshot size in bytes | [optional] 
**Vmstate** | Pointer to **bool** | Whether VM state is included | [optional] 

## Methods

### NewSnapshotItem

`func NewSnapshotItem(name string, description string, ) *SnapshotItem`

NewSnapshotItem instantiates a new SnapshotItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotItemWithDefaults

`func NewSnapshotItemWithDefaults() *SnapshotItem`

NewSnapshotItemWithDefaults instantiates a new SnapshotItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SnapshotItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SnapshotItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SnapshotItem) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *SnapshotItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SnapshotItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SnapshotItem) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetSnaptime

`func (o *SnapshotItem) GetSnaptime() string`

GetSnaptime returns the Snaptime field if non-nil, zero value otherwise.

### GetSnaptimeOk

`func (o *SnapshotItem) GetSnaptimeOk() (*string, bool)`

GetSnaptimeOk returns a tuple with the Snaptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnaptime

`func (o *SnapshotItem) SetSnaptime(v string)`

SetSnaptime sets Snaptime field to given value.

### HasSnaptime

`func (o *SnapshotItem) HasSnaptime() bool`

HasSnaptime returns a boolean if a field has been set.

### GetSize

`func (o *SnapshotItem) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *SnapshotItem) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *SnapshotItem) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *SnapshotItem) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetVmstate

`func (o *SnapshotItem) GetVmstate() bool`

GetVmstate returns the Vmstate field if non-nil, zero value otherwise.

### GetVmstateOk

`func (o *SnapshotItem) GetVmstateOk() (*bool, bool)`

GetVmstateOk returns a tuple with the Vmstate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmstate

`func (o *SnapshotItem) SetVmstate(v bool)`

SetVmstate sets Vmstate field to given value.

### HasVmstate

`func (o *SnapshotItem) HasVmstate() bool`

HasVmstate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


