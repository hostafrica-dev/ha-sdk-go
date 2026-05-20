# MountIsoRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceId** | **string** | Service ID - must be sent as a string | 
**Iso** | **string** | Name of the ISO image to mount | 
**FormatMasterDisk** | Pointer to **bool** | Whether to format the master disk before mounting the ISO | [optional] 

## Methods

### NewMountIsoRequestContent

`func NewMountIsoRequestContent(serviceId string, iso string, ) *MountIsoRequestContent`

NewMountIsoRequestContent instantiates a new MountIsoRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMountIsoRequestContentWithDefaults

`func NewMountIsoRequestContentWithDefaults() *MountIsoRequestContent`

NewMountIsoRequestContentWithDefaults instantiates a new MountIsoRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceId

`func (o *MountIsoRequestContent) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *MountIsoRequestContent) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *MountIsoRequestContent) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetIso

`func (o *MountIsoRequestContent) GetIso() string`

GetIso returns the Iso field if non-nil, zero value otherwise.

### GetIsoOk

`func (o *MountIsoRequestContent) GetIsoOk() (*string, bool)`

GetIsoOk returns a tuple with the Iso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIso

`func (o *MountIsoRequestContent) SetIso(v string)`

SetIso sets Iso field to given value.


### GetFormatMasterDisk

`func (o *MountIsoRequestContent) GetFormatMasterDisk() bool`

GetFormatMasterDisk returns the FormatMasterDisk field if non-nil, zero value otherwise.

### GetFormatMasterDiskOk

`func (o *MountIsoRequestContent) GetFormatMasterDiskOk() (*bool, bool)`

GetFormatMasterDiskOk returns a tuple with the FormatMasterDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatMasterDisk

`func (o *MountIsoRequestContent) SetFormatMasterDisk(v bool)`

SetFormatMasterDisk sets FormatMasterDisk field to given value.

### HasFormatMasterDisk

`func (o *MountIsoRequestContent) HasFormatMasterDisk() bool`

HasFormatMasterDisk returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


