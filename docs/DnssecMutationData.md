# DnssecMutationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**Records** | Pointer to [**[]DnssecRecordInfo**](DnssecRecordInfo.md) | DNSSEC DS records after mutation when returned by upstream | [optional] 

## Methods

### NewDnssecMutationData

`func NewDnssecMutationData(message string, ) *DnssecMutationData`

NewDnssecMutationData instantiates a new DnssecMutationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnssecMutationDataWithDefaults

`func NewDnssecMutationDataWithDefaults() *DnssecMutationData`

NewDnssecMutationDataWithDefaults instantiates a new DnssecMutationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *DnssecMutationData) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DnssecMutationData) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DnssecMutationData) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetRecords

`func (o *DnssecMutationData) GetRecords() []DnssecRecordInfo`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *DnssecMutationData) GetRecordsOk() (*[]DnssecRecordInfo, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *DnssecMutationData) SetRecords(v []DnssecRecordInfo)`

SetRecords sets Records field to given value.

### HasRecords

`func (o *DnssecMutationData) HasRecords() bool`

HasRecords returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


