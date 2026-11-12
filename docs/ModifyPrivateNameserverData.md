# ModifyPrivateNameserverData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Status message indicating the result | 
**Nameserver** | Pointer to **string** | Existing nameserver hostname that was modified | [optional] 
**NewNameserver** | Pointer to **string** | New hostname when a rename was requested | [optional] 
**OldIpAddress** | Pointer to **string** | Previous glue IP when supplied | [optional] 
**NewIpAddress** | Pointer to **string** | New glue IP when supplied | [optional] 

## Methods

### NewModifyPrivateNameserverData

`func NewModifyPrivateNameserverData(message string, ) *ModifyPrivateNameserverData`

NewModifyPrivateNameserverData instantiates a new ModifyPrivateNameserverData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyPrivateNameserverDataWithDefaults

`func NewModifyPrivateNameserverDataWithDefaults() *ModifyPrivateNameserverData`

NewModifyPrivateNameserverDataWithDefaults instantiates a new ModifyPrivateNameserverData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ModifyPrivateNameserverData) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ModifyPrivateNameserverData) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ModifyPrivateNameserverData) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetNameserver

`func (o *ModifyPrivateNameserverData) GetNameserver() string`

GetNameserver returns the Nameserver field if non-nil, zero value otherwise.

### GetNameserverOk

`func (o *ModifyPrivateNameserverData) GetNameserverOk() (*string, bool)`

GetNameserverOk returns a tuple with the Nameserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameserver

`func (o *ModifyPrivateNameserverData) SetNameserver(v string)`

SetNameserver sets Nameserver field to given value.

### HasNameserver

`func (o *ModifyPrivateNameserverData) HasNameserver() bool`

HasNameserver returns a boolean if a field has been set.

### GetNewNameserver

`func (o *ModifyPrivateNameserverData) GetNewNameserver() string`

GetNewNameserver returns the NewNameserver field if non-nil, zero value otherwise.

### GetNewNameserverOk

`func (o *ModifyPrivateNameserverData) GetNewNameserverOk() (*string, bool)`

GetNewNameserverOk returns a tuple with the NewNameserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewNameserver

`func (o *ModifyPrivateNameserverData) SetNewNameserver(v string)`

SetNewNameserver sets NewNameserver field to given value.

### HasNewNameserver

`func (o *ModifyPrivateNameserverData) HasNewNameserver() bool`

HasNewNameserver returns a boolean if a field has been set.

### GetOldIpAddress

`func (o *ModifyPrivateNameserverData) GetOldIpAddress() string`

GetOldIpAddress returns the OldIpAddress field if non-nil, zero value otherwise.

### GetOldIpAddressOk

`func (o *ModifyPrivateNameserverData) GetOldIpAddressOk() (*string, bool)`

GetOldIpAddressOk returns a tuple with the OldIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldIpAddress

`func (o *ModifyPrivateNameserverData) SetOldIpAddress(v string)`

SetOldIpAddress sets OldIpAddress field to given value.

### HasOldIpAddress

`func (o *ModifyPrivateNameserverData) HasOldIpAddress() bool`

HasOldIpAddress returns a boolean if a field has been set.

### GetNewIpAddress

`func (o *ModifyPrivateNameserverData) GetNewIpAddress() string`

GetNewIpAddress returns the NewIpAddress field if non-nil, zero value otherwise.

### GetNewIpAddressOk

`func (o *ModifyPrivateNameserverData) GetNewIpAddressOk() (*string, bool)`

GetNewIpAddressOk returns a tuple with the NewIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewIpAddress

`func (o *ModifyPrivateNameserverData) SetNewIpAddress(v string)`

SetNewIpAddress sets NewIpAddress field to given value.

### HasNewIpAddress

`func (o *ModifyPrivateNameserverData) HasNewIpAddress() bool`

HasNewIpAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


