# EditRdnsRecordRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordId** | **int32** | ID of the PTR record to update | 
**Hostname** | **string** | Updated PTR target hostname | 
**Ttl** | Pointer to **int32** | Time-to-live in seconds. When omitted, upstream keeps the existing TTL; clamped to [30, 86400] when set | [optional] 

## Methods

### NewEditRdnsRecordRequestContent

`func NewEditRdnsRecordRequestContent(recordId int32, hostname string, ) *EditRdnsRecordRequestContent`

NewEditRdnsRecordRequestContent instantiates a new EditRdnsRecordRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditRdnsRecordRequestContentWithDefaults

`func NewEditRdnsRecordRequestContentWithDefaults() *EditRdnsRecordRequestContent`

NewEditRdnsRecordRequestContentWithDefaults instantiates a new EditRdnsRecordRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordId

`func (o *EditRdnsRecordRequestContent) GetRecordId() int32`

GetRecordId returns the RecordId field if non-nil, zero value otherwise.

### GetRecordIdOk

`func (o *EditRdnsRecordRequestContent) GetRecordIdOk() (*int32, bool)`

GetRecordIdOk returns a tuple with the RecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordId

`func (o *EditRdnsRecordRequestContent) SetRecordId(v int32)`

SetRecordId sets RecordId field to given value.


### GetHostname

`func (o *EditRdnsRecordRequestContent) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *EditRdnsRecordRequestContent) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *EditRdnsRecordRequestContent) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetTtl

`func (o *EditRdnsRecordRequestContent) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *EditRdnsRecordRequestContent) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *EditRdnsRecordRequestContent) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *EditRdnsRecordRequestContent) HasTtl() bool`

HasTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


