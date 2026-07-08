# CreateSnapshotRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceId** | **string** | Service ID - must be sent as a string | 
**Name** | **string** | Name for the snapshot | 
**Description** | Pointer to **string** | Description for the snapshot | [optional] 
**IncludeRam** | Pointer to **bool** | Whether to include RAM state in the snapshot. Defaults to false when omitted. | [optional] 

## Methods

### NewCreateSnapshotRequestContent

`func NewCreateSnapshotRequestContent(serviceId string, name string, ) *CreateSnapshotRequestContent`

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


### GetName

`func (o *CreateSnapshotRequestContent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateSnapshotRequestContent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateSnapshotRequestContent) SetName(v string)`

SetName sets Name field to given value.


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

### GetIncludeRam

`func (o *CreateSnapshotRequestContent) GetIncludeRam() bool`

GetIncludeRam returns the IncludeRam field if non-nil, zero value otherwise.

### GetIncludeRamOk

`func (o *CreateSnapshotRequestContent) GetIncludeRamOk() (*bool, bool)`

GetIncludeRamOk returns a tuple with the IncludeRam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeRam

`func (o *CreateSnapshotRequestContent) SetIncludeRam(v bool)`

SetIncludeRam sets IncludeRam field to given value.

### HasIncludeRam

`func (o *CreateSnapshotRequestContent) HasIncludeRam() bool`

HasIncludeRam returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


