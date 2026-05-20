# ServiceSnapshotsResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Status message indicating the result | 
**Snapshots** | Pointer to [**[]SnapshotItem**](SnapshotItem.md) | List of snapshots | [optional] 
**MaxSnapshots** | Pointer to **int32** | Maximum number of snapshots allowed | [optional] 
**SnapshotsCount** | Pointer to **int32** | Number of snapshots currently stored | [optional] 

## Methods

### NewServiceSnapshotsResponseData

`func NewServiceSnapshotsResponseData(message string, ) *ServiceSnapshotsResponseData`

NewServiceSnapshotsResponseData instantiates a new ServiceSnapshotsResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceSnapshotsResponseDataWithDefaults

`func NewServiceSnapshotsResponseDataWithDefaults() *ServiceSnapshotsResponseData`

NewServiceSnapshotsResponseDataWithDefaults instantiates a new ServiceSnapshotsResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ServiceSnapshotsResponseData) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ServiceSnapshotsResponseData) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ServiceSnapshotsResponseData) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetSnapshots

`func (o *ServiceSnapshotsResponseData) GetSnapshots() []SnapshotItem`

GetSnapshots returns the Snapshots field if non-nil, zero value otherwise.

### GetSnapshotsOk

`func (o *ServiceSnapshotsResponseData) GetSnapshotsOk() (*[]SnapshotItem, bool)`

GetSnapshotsOk returns a tuple with the Snapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshots

`func (o *ServiceSnapshotsResponseData) SetSnapshots(v []SnapshotItem)`

SetSnapshots sets Snapshots field to given value.

### HasSnapshots

`func (o *ServiceSnapshotsResponseData) HasSnapshots() bool`

HasSnapshots returns a boolean if a field has been set.

### GetMaxSnapshots

`func (o *ServiceSnapshotsResponseData) GetMaxSnapshots() int32`

GetMaxSnapshots returns the MaxSnapshots field if non-nil, zero value otherwise.

### GetMaxSnapshotsOk

`func (o *ServiceSnapshotsResponseData) GetMaxSnapshotsOk() (*int32, bool)`

GetMaxSnapshotsOk returns a tuple with the MaxSnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSnapshots

`func (o *ServiceSnapshotsResponseData) SetMaxSnapshots(v int32)`

SetMaxSnapshots sets MaxSnapshots field to given value.

### HasMaxSnapshots

`func (o *ServiceSnapshotsResponseData) HasMaxSnapshots() bool`

HasMaxSnapshots returns a boolean if a field has been set.

### GetSnapshotsCount

`func (o *ServiceSnapshotsResponseData) GetSnapshotsCount() int32`

GetSnapshotsCount returns the SnapshotsCount field if non-nil, zero value otherwise.

### GetSnapshotsCountOk

`func (o *ServiceSnapshotsResponseData) GetSnapshotsCountOk() (*int32, bool)`

GetSnapshotsCountOk returns a tuple with the SnapshotsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotsCount

`func (o *ServiceSnapshotsResponseData) SetSnapshotsCount(v int32)`

SetSnapshotsCount sets SnapshotsCount field to given value.

### HasSnapshotsCount

`func (o *ServiceSnapshotsResponseData) HasSnapshotsCount() bool`

HasSnapshotsCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


