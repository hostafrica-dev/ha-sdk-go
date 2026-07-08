# SnapshotJobDeleteResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Status message indicating the result | 
**Limits** | [**SnapshotJobLimits**](SnapshotJobLimits.md) |  | 

## Methods

### NewSnapshotJobDeleteResponseData

`func NewSnapshotJobDeleteResponseData(message string, limits SnapshotJobLimits, ) *SnapshotJobDeleteResponseData`

NewSnapshotJobDeleteResponseData instantiates a new SnapshotJobDeleteResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotJobDeleteResponseDataWithDefaults

`func NewSnapshotJobDeleteResponseDataWithDefaults() *SnapshotJobDeleteResponseData`

NewSnapshotJobDeleteResponseDataWithDefaults instantiates a new SnapshotJobDeleteResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *SnapshotJobDeleteResponseData) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SnapshotJobDeleteResponseData) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SnapshotJobDeleteResponseData) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetLimits

`func (o *SnapshotJobDeleteResponseData) GetLimits() SnapshotJobLimits`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *SnapshotJobDeleteResponseData) GetLimitsOk() (*SnapshotJobLimits, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *SnapshotJobDeleteResponseData) SetLimits(v SnapshotJobLimits)`

SetLimits sets Limits field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


