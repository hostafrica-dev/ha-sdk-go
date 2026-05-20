# RollbackSnapshotRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceId** | **string** | Service ID - must be sent as a string | 
**SnapshotName** | **string** | Snapshot name to rollback to | 

## Methods

### NewRollbackSnapshotRequestContent

`func NewRollbackSnapshotRequestContent(serviceId string, snapshotName string, ) *RollbackSnapshotRequestContent`

NewRollbackSnapshotRequestContent instantiates a new RollbackSnapshotRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRollbackSnapshotRequestContentWithDefaults

`func NewRollbackSnapshotRequestContentWithDefaults() *RollbackSnapshotRequestContent`

NewRollbackSnapshotRequestContentWithDefaults instantiates a new RollbackSnapshotRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceId

`func (o *RollbackSnapshotRequestContent) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *RollbackSnapshotRequestContent) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *RollbackSnapshotRequestContent) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetSnapshotName

`func (o *RollbackSnapshotRequestContent) GetSnapshotName() string`

GetSnapshotName returns the SnapshotName field if non-nil, zero value otherwise.

### GetSnapshotNameOk

`func (o *RollbackSnapshotRequestContent) GetSnapshotNameOk() (*string, bool)`

GetSnapshotNameOk returns a tuple with the SnapshotName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotName

`func (o *RollbackSnapshotRequestContent) SetSnapshotName(v string)`

SetSnapshotName sets SnapshotName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


