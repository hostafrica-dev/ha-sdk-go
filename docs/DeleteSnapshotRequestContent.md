# DeleteSnapshotRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceId** | **string** | Service ID - must be sent as a string | 
**SnapshotName** | **string** | Snapshot name to delete | 

## Methods

### NewDeleteSnapshotRequestContent

`func NewDeleteSnapshotRequestContent(serviceId string, snapshotName string, ) *DeleteSnapshotRequestContent`

NewDeleteSnapshotRequestContent instantiates a new DeleteSnapshotRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteSnapshotRequestContentWithDefaults

`func NewDeleteSnapshotRequestContentWithDefaults() *DeleteSnapshotRequestContent`

NewDeleteSnapshotRequestContentWithDefaults instantiates a new DeleteSnapshotRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceId

`func (o *DeleteSnapshotRequestContent) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *DeleteSnapshotRequestContent) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *DeleteSnapshotRequestContent) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetSnapshotName

`func (o *DeleteSnapshotRequestContent) GetSnapshotName() string`

GetSnapshotName returns the SnapshotName field if non-nil, zero value otherwise.

### GetSnapshotNameOk

`func (o *DeleteSnapshotRequestContent) GetSnapshotNameOk() (*string, bool)`

GetSnapshotNameOk returns a tuple with the SnapshotName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotName

`func (o *DeleteSnapshotRequestContent) SetSnapshotName(v string)`

SetSnapshotName sets SnapshotName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


