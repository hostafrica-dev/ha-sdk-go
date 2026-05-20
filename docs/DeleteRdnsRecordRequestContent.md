# DeleteRdnsRecordRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordId** | **int32** | ID of the PTR record to delete | 

## Methods

### NewDeleteRdnsRecordRequestContent

`func NewDeleteRdnsRecordRequestContent(recordId int32, ) *DeleteRdnsRecordRequestContent`

NewDeleteRdnsRecordRequestContent instantiates a new DeleteRdnsRecordRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteRdnsRecordRequestContentWithDefaults

`func NewDeleteRdnsRecordRequestContentWithDefaults() *DeleteRdnsRecordRequestContent`

NewDeleteRdnsRecordRequestContentWithDefaults instantiates a new DeleteRdnsRecordRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordId

`func (o *DeleteRdnsRecordRequestContent) GetRecordId() int32`

GetRecordId returns the RecordId field if non-nil, zero value otherwise.

### GetRecordIdOk

`func (o *DeleteRdnsRecordRequestContent) GetRecordIdOk() (*int32, bool)`

GetRecordIdOk returns a tuple with the RecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordId

`func (o *DeleteRdnsRecordRequestContent) SetRecordId(v int32)`

SetRecordId sets RecordId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


