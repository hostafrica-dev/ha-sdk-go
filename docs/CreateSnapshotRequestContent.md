# CreateSnapshotRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceId** | **string** | Service ID - must be sent as a string | 
**Snapname** | Pointer to **string** | Name for the snapshot | [optional] 
**Description** | Pointer to **string** | Description for the snapshot | [optional] 

## Methods

### NewCreateSnapshotRequestContent

`func NewCreateSnapshotRequestContent(serviceId string, ) *CreateSnapshotRequestContent`

NewCreateSnapshotRequestContent instantiates a new CreateSnapshotRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSnapshotRequestContentWithDefaults

`func NewCreateSnapshotRequestContentWithDefaults() *CreateSnapshotRequestContent`

NewCreateSnapshotRequestContentWithDefaults instantiates a new CreateSnapshotRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceId

`func (o *CreateSnapshotRequestContent) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *CreateSnapshotRequestContent) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *CreateSnapshotRequestContent) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetSnapname

`func (o *CreateSnapshotRequestContent) GetSnapname() string`

GetSnapname returns the Snapname field if non-nil, zero value otherwise.

### GetSnapnameOk

`func (o *CreateSnapshotRequestContent) GetSnapnameOk() (*string, bool)`

GetSnapnameOk returns a tuple with the Snapname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapname

`func (o *CreateSnapshotRequestContent) SetSnapname(v string)`

SetSnapname sets Snapname field to given value.

### HasSnapname

`func (o *CreateSnapshotRequestContent) HasSnapname() bool`

HasSnapname returns a boolean if a field has been set.

### GetDescription

`func (o *CreateSnapshotRequestContent) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateSnapshotRequestContent) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateSnapshotRequestContent) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateSnapshotRequestContent) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


