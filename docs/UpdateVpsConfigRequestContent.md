# UpdateVpsConfigRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceId** | **string** | Service ID - must be sent as a string | 
**Name** | Pointer to **string** | New name for the VPS | [optional] 
**Hostname** | Pointer to **string** | New hostname for the VPS | [optional] 
**AutoStart** | Pointer to **int32** | Auto-start on boot (0 or 1) | [optional] 
**Boot** | Pointer to **string** | Boot order configuration | [optional] 
**Ide2** | Pointer to **string** | IDE2 device configuration | [optional] 
**Cdrom** | Pointer to **string** | CD-ROM configuration | [optional] 

## Methods

### NewUpdateVpsConfigRequestContent

`func NewUpdateVpsConfigRequestContent(serviceId string, ) *UpdateVpsConfigRequestContent`

NewUpdateVpsConfigRequestContent instantiates a new UpdateVpsConfigRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateVpsConfigRequestContentWithDefaults

`func NewUpdateVpsConfigRequestContentWithDefaults() *UpdateVpsConfigRequestContent`

NewUpdateVpsConfigRequestContentWithDefaults instantiates a new UpdateVpsConfigRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceId

`func (o *UpdateVpsConfigRequestContent) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *UpdateVpsConfigRequestContent) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *UpdateVpsConfigRequestContent) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetName

`func (o *UpdateVpsConfigRequestContent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateVpsConfigRequestContent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateVpsConfigRequestContent) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateVpsConfigRequestContent) HasName() bool`

HasName returns a boolean if a field has been set.

### GetHostname

`func (o *UpdateVpsConfigRequestContent) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *UpdateVpsConfigRequestContent) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *UpdateVpsConfigRequestContent) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *UpdateVpsConfigRequestContent) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetAutoStart

`func (o *UpdateVpsConfigRequestContent) GetAutoStart() int32`

GetAutoStart returns the AutoStart field if non-nil, zero value otherwise.

### GetAutoStartOk

`func (o *UpdateVpsConfigRequestContent) GetAutoStartOk() (*int32, bool)`

GetAutoStartOk returns a tuple with the AutoStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoStart

`func (o *UpdateVpsConfigRequestContent) SetAutoStart(v int32)`

SetAutoStart sets AutoStart field to given value.

### HasAutoStart

`func (o *UpdateVpsConfigRequestContent) HasAutoStart() bool`

HasAutoStart returns a boolean if a field has been set.

### GetBoot

`func (o *UpdateVpsConfigRequestContent) GetBoot() string`

GetBoot returns the Boot field if non-nil, zero value otherwise.

### GetBootOk

`func (o *UpdateVpsConfigRequestContent) GetBootOk() (*string, bool)`

GetBootOk returns a tuple with the Boot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoot

`func (o *UpdateVpsConfigRequestContent) SetBoot(v string)`

SetBoot sets Boot field to given value.

### HasBoot

`func (o *UpdateVpsConfigRequestContent) HasBoot() bool`

HasBoot returns a boolean if a field has been set.

### GetIde2

`func (o *UpdateVpsConfigRequestContent) GetIde2() string`

GetIde2 returns the Ide2 field if non-nil, zero value otherwise.

### GetIde2Ok

`func (o *UpdateVpsConfigRequestContent) GetIde2Ok() (*string, bool)`

GetIde2Ok returns a tuple with the Ide2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIde2

`func (o *UpdateVpsConfigRequestContent) SetIde2(v string)`

SetIde2 sets Ide2 field to given value.

### HasIde2

`func (o *UpdateVpsConfigRequestContent) HasIde2() bool`

HasIde2 returns a boolean if a field has been set.

### GetCdrom

`func (o *UpdateVpsConfigRequestContent) GetCdrom() string`

GetCdrom returns the Cdrom field if non-nil, zero value otherwise.

### GetCdromOk

`func (o *UpdateVpsConfigRequestContent) GetCdromOk() (*string, bool)`

GetCdromOk returns a tuple with the Cdrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdrom

`func (o *UpdateVpsConfigRequestContent) SetCdrom(v string)`

SetCdrom sets Cdrom field to given value.

### HasCdrom

`func (o *UpdateVpsConfigRequestContent) HasCdrom() bool`

HasCdrom returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


