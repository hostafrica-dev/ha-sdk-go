# UpdateDomainLockStatusData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Status message indicating the result | 
**Domain** | Pointer to **string** | Domain name (FQDN) | [optional] 
**LockStatus** | Pointer to **string** | Applied lock state as reported after the update (e.g. locked, unlocked) | [optional] 
**RequestedStatus** | Pointer to **bool** | Boolean lock value that was requested | [optional] 

## Methods

### NewUpdateDomainLockStatusData

`func NewUpdateDomainLockStatusData(message string, ) *UpdateDomainLockStatusData`

NewUpdateDomainLockStatusData instantiates a new UpdateDomainLockStatusData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDomainLockStatusDataWithDefaults

`func NewUpdateDomainLockStatusDataWithDefaults() *UpdateDomainLockStatusData`

NewUpdateDomainLockStatusDataWithDefaults instantiates a new UpdateDomainLockStatusData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *UpdateDomainLockStatusData) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *UpdateDomainLockStatusData) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *UpdateDomainLockStatusData) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetDomain

`func (o *UpdateDomainLockStatusData) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *UpdateDomainLockStatusData) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *UpdateDomainLockStatusData) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *UpdateDomainLockStatusData) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetLockStatus

`func (o *UpdateDomainLockStatusData) GetLockStatus() string`

GetLockStatus returns the LockStatus field if non-nil, zero value otherwise.

### GetLockStatusOk

`func (o *UpdateDomainLockStatusData) GetLockStatusOk() (*string, bool)`

GetLockStatusOk returns a tuple with the LockStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockStatus

`func (o *UpdateDomainLockStatusData) SetLockStatus(v string)`

SetLockStatus sets LockStatus field to given value.

### HasLockStatus

`func (o *UpdateDomainLockStatusData) HasLockStatus() bool`

HasLockStatus returns a boolean if a field has been set.

### GetRequestedStatus

`func (o *UpdateDomainLockStatusData) GetRequestedStatus() bool`

GetRequestedStatus returns the RequestedStatus field if non-nil, zero value otherwise.

### GetRequestedStatusOk

`func (o *UpdateDomainLockStatusData) GetRequestedStatusOk() (*bool, bool)`

GetRequestedStatusOk returns a tuple with the RequestedStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedStatus

`func (o *UpdateDomainLockStatusData) SetRequestedStatus(v bool)`

SetRequestedStatus sets RequestedStatus field to given value.

### HasRequestedStatus

`func (o *UpdateDomainLockStatusData) HasRequestedStatus() bool`

HasRequestedStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


