# VpsConfigResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Status message indicating the result | 
**Hostname** | Pointer to **string** | Hostname of the VPS | [optional] 
**IsKvm** | Pointer to **bool** | Whether the VPS uses KVM virtualisation | [optional] 
**SshkeysEnabled** | Pointer to **bool** | Whether SSH key management is enabled for this VPS | [optional] 
**AvailableBootOrder** | Pointer to **interface{}** | Available boot order options keyed by device identifier | [optional] 
**BootOrder0** | Pointer to **string** | Primary boot order slot (device identifier, or null) | [optional] 
**BootOrder1** | Pointer to **string** | Secondary boot order slot (device identifier, or null) | [optional] 
**BootOrder2** | Pointer to **string** | Tertiary boot order slot (device identifier, or null) | [optional] 
**Sshkeys** | Pointer to **string** | Current SSH key(s) configured on the VPS (cloud-init) | [optional] 

## Methods

### NewVpsConfigResponseData

`func NewVpsConfigResponseData(message string, ) *VpsConfigResponseData`

NewVpsConfigResponseData instantiates a new VpsConfigResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpsConfigResponseDataWithDefaults

`func NewVpsConfigResponseDataWithDefaults() *VpsConfigResponseData`

NewVpsConfigResponseDataWithDefaults instantiates a new VpsConfigResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *VpsConfigResponseData) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *VpsConfigResponseData) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *VpsConfigResponseData) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetHostname

`func (o *VpsConfigResponseData) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *VpsConfigResponseData) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *VpsConfigResponseData) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *VpsConfigResponseData) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetIsKvm

`func (o *VpsConfigResponseData) GetIsKvm() bool`

GetIsKvm returns the IsKvm field if non-nil, zero value otherwise.

### GetIsKvmOk

`func (o *VpsConfigResponseData) GetIsKvmOk() (*bool, bool)`

GetIsKvmOk returns a tuple with the IsKvm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsKvm

`func (o *VpsConfigResponseData) SetIsKvm(v bool)`

SetIsKvm sets IsKvm field to given value.

### HasIsKvm

`func (o *VpsConfigResponseData) HasIsKvm() bool`

HasIsKvm returns a boolean if a field has been set.

### GetSshkeysEnabled

`func (o *VpsConfigResponseData) GetSshkeysEnabled() bool`

GetSshkeysEnabled returns the SshkeysEnabled field if non-nil, zero value otherwise.

### GetSshkeysEnabledOk

`func (o *VpsConfigResponseData) GetSshkeysEnabledOk() (*bool, bool)`

GetSshkeysEnabledOk returns a tuple with the SshkeysEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshkeysEnabled

`func (o *VpsConfigResponseData) SetSshkeysEnabled(v bool)`

SetSshkeysEnabled sets SshkeysEnabled field to given value.

### HasSshkeysEnabled

`func (o *VpsConfigResponseData) HasSshkeysEnabled() bool`

HasSshkeysEnabled returns a boolean if a field has been set.

### GetAvailableBootOrder

`func (o *VpsConfigResponseData) GetAvailableBootOrder() interface{}`

GetAvailableBootOrder returns the AvailableBootOrder field if non-nil, zero value otherwise.

### GetAvailableBootOrderOk

`func (o *VpsConfigResponseData) GetAvailableBootOrderOk() (*interface{}, bool)`

GetAvailableBootOrderOk returns a tuple with the AvailableBootOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableBootOrder

`func (o *VpsConfigResponseData) SetAvailableBootOrder(v interface{})`

SetAvailableBootOrder sets AvailableBootOrder field to given value.

### HasAvailableBootOrder

`func (o *VpsConfigResponseData) HasAvailableBootOrder() bool`

HasAvailableBootOrder returns a boolean if a field has been set.

### SetAvailableBootOrderNil

`func (o *VpsConfigResponseData) SetAvailableBootOrderNil(b bool)`

 SetAvailableBootOrderNil sets the value for AvailableBootOrder to be an explicit nil

### UnsetAvailableBootOrder
`func (o *VpsConfigResponseData) UnsetAvailableBootOrder()`

UnsetAvailableBootOrder ensures that no value is present for AvailableBootOrder, not even an explicit nil
### GetBootOrder0

`func (o *VpsConfigResponseData) GetBootOrder0() string`

GetBootOrder0 returns the BootOrder0 field if non-nil, zero value otherwise.

### GetBootOrder0Ok

`func (o *VpsConfigResponseData) GetBootOrder0Ok() (*string, bool)`

GetBootOrder0Ok returns a tuple with the BootOrder0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootOrder0

`func (o *VpsConfigResponseData) SetBootOrder0(v string)`

SetBootOrder0 sets BootOrder0 field to given value.

### HasBootOrder0

`func (o *VpsConfigResponseData) HasBootOrder0() bool`

HasBootOrder0 returns a boolean if a field has been set.

### GetBootOrder1

`func (o *VpsConfigResponseData) GetBootOrder1() string`

GetBootOrder1 returns the BootOrder1 field if non-nil, zero value otherwise.

### GetBootOrder1Ok

`func (o *VpsConfigResponseData) GetBootOrder1Ok() (*string, bool)`

GetBootOrder1Ok returns a tuple with the BootOrder1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootOrder1

`func (o *VpsConfigResponseData) SetBootOrder1(v string)`

SetBootOrder1 sets BootOrder1 field to given value.

### HasBootOrder1

`func (o *VpsConfigResponseData) HasBootOrder1() bool`

HasBootOrder1 returns a boolean if a field has been set.

### GetBootOrder2

`func (o *VpsConfigResponseData) GetBootOrder2() string`

GetBootOrder2 returns the BootOrder2 field if non-nil, zero value otherwise.

### GetBootOrder2Ok

`func (o *VpsConfigResponseData) GetBootOrder2Ok() (*string, bool)`

GetBootOrder2Ok returns a tuple with the BootOrder2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootOrder2

`func (o *VpsConfigResponseData) SetBootOrder2(v string)`

SetBootOrder2 sets BootOrder2 field to given value.

### HasBootOrder2

`func (o *VpsConfigResponseData) HasBootOrder2() bool`

HasBootOrder2 returns a boolean if a field has been set.

### GetSshkeys

`func (o *VpsConfigResponseData) GetSshkeys() string`

GetSshkeys returns the Sshkeys field if non-nil, zero value otherwise.

### GetSshkeysOk

`func (o *VpsConfigResponseData) GetSshkeysOk() (*string, bool)`

GetSshkeysOk returns a tuple with the Sshkeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshkeys

`func (o *VpsConfigResponseData) SetSshkeys(v string)`

SetSshkeys sets Sshkeys field to given value.

### HasSshkeys

`func (o *VpsConfigResponseData) HasSshkeys() bool`

HasSshkeys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


