# ListDnssecRecordsData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**Records** | [**[]DnssecRecordInfo**](DnssecRecordInfo.md) | DNSSEC DS records configured for the domain | 

## Methods

### NewListDnssecRecordsData

`func NewListDnssecRecordsData(message string, records []DnssecRecordInfo, ) *ListDnssecRecordsData`

NewListDnssecRecordsData instantiates a new ListDnssecRecordsData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListDnssecRecordsDataWithDefaults

`func NewListDnssecRecordsDataWithDefaults() *ListDnssecRecordsData`

NewListDnssecRecordsDataWithDefaults instantiates a new ListDnssecRecordsData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ListDnssecRecordsData) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ListDnssecRecordsData) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ListDnssecRecordsData) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetRecords

`func (o *ListDnssecRecordsData) GetRecords() []DnssecRecordInfo`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *ListDnssecRecordsData) GetRecordsOk() (*[]DnssecRecordInfo, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *ListDnssecRecordsData) SetRecords(v []DnssecRecordInfo)`

SetRecords sets Records field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


