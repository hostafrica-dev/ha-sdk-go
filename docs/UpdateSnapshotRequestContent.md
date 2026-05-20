# UpdateSnapshotRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceId** | **string** | Service ID - must be sent as a string | 
**SnapshotName** | **string** | Snapshot name to update | 
**Description** | Pointer to **string** | New description for the snapshot | [optional] 

## Methods

### NewUpdateSnapshotRequestContent

`func NewUpdateSnapshotRequestContent(serviceId string, snapshotName string, ) *UpdateSnapshotRequestContent`

NewUpdateSnapshotRequestContent instantiates a new UpdateSnapshotRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSnapshotRequestContentWithDefaults

`func NewUpdateSnapshotRequestContentWithDefaults() *UpdateSnapshotRequestContent`

NewUpdateSnapshotRequestContentWithDefaults instantiates a new UpdateSnapshotRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceId

`func (o *UpdateSnapshotRequestContent) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *UpdateSnapshotRequestContent) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *UpdateSnapshotRequestContent) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetSnapshotName

`func (o *UpdateSnapshotRequestContent) GetSnapshotName() string`

GetSnapshotName returns the SnapshotName field if non-nil, zero value otherwise.

### GetSnapshotNameOk

`func (o *UpdateSnapshotRequestContent) GetSnapshotNameOk() (*string, bool)`

GetSnapshotNameOk returns a tuple with the SnapshotName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotName

`func (o *UpdateSnapshotRequestContent) SetSnapshotName(v string)`

SetSnapshotName sets SnapshotName field to given value.


### GetDescription

`func (o *UpdateSnapshotRequestContent) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateSnapshotRequestContent) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateSnapshotRequestContent) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateSnapshotRequestContent) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


