# DnsRecordMutationRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Record line from get-dns-zone-details (positive integer string); required for edit and delete | [optional] 
**Name** | Pointer to **string** | Record host/name (e.g. @, www, mail); required for add | [optional] 
**Type** | Pointer to **string** | Record type (e.g. A, AAAA, CNAME, MX, TXT, NS, SRV); required for add | [optional] 
**Content** | Pointer to **string** | Primary record value; required for add and edit. Meaning depends on type — see DnsRecordInfo | [optional] 
**Ttl** | Pointer to **int32** | Time-to-live in seconds; defaults to 3600 on add when omitted | [optional] 
**Priority** | Pointer to **int32** | MX preference when type is MX; SRV priority when type is SRV | [optional] 
**Weight** | Pointer to **int32** | SRV weight when type is SRV | [optional] 
**Port** | Pointer to **int32** | SRV port when type is SRV; required for add and edit when type is SRV | [optional] 

## Methods

### NewDnsRecordMutationRecord

`func NewDnsRecordMutationRecord() *DnsRecordMutationRecord`

NewDnsRecordMutationRecord instantiates a new DnsRecordMutationRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsRecordMutationRecordWithDefaults

`func NewDnsRecordMutationRecordWithDefaults() *DnsRecordMutationRecord`

NewDnsRecordMutationRecordWithDefaults instantiates a new DnsRecordMutationRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DnsRecordMutationRecord) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DnsRecordMutationRecord) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DnsRecordMutationRecord) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DnsRecordMutationRecord) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DnsRecordMutationRecord) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DnsRecordMutationRecord) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DnsRecordMutationRecord) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DnsRecordMutationRecord) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *DnsRecordMutationRecord) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DnsRecordMutationRecord) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DnsRecordMutationRecord) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DnsRecordMutationRecord) HasType() bool`

HasType returns a boolean if a field has been set.

### GetContent

`func (o *DnsRecordMutationRecord) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *DnsRecordMutationRecord) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *DnsRecordMutationRecord) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *DnsRecordMutationRecord) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetTtl

`func (o *DnsRecordMutationRecord) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *DnsRecordMutationRecord) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *DnsRecordMutationRecord) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *DnsRecordMutationRecord) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetPriority

`func (o *DnsRecordMutationRecord) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *DnsRecordMutationRecord) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *DnsRecordMutationRecord) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *DnsRecordMutationRecord) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetWeight

`func (o *DnsRecordMutationRecord) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *DnsRecordMutationRecord) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *DnsRecordMutationRecord) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *DnsRecordMutationRecord) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetPort

`func (o *DnsRecordMutationRecord) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *DnsRecordMutationRecord) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *DnsRecordMutationRecord) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *DnsRecordMutationRecord) HasPort() bool`

HasPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


