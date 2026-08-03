# DnsRecordInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Record identifier | [optional] 
**Name** | Pointer to **string** | Record host/name (e.g. @, www, mail) | [optional] 
**Type** | Pointer to **string** | Record type (e.g. A, AAAA, CNAME, MX, TXT, NS, SRV) | [optional] 
**Content** | Pointer to **string** | Primary record value: IP address (A/AAAA), hostname (CNAME/NS/MX/SRV), or text (TXT) | [optional] 
**Ttl** | Pointer to **int32** | Time-to-live in seconds | [optional] 
**Priority** | Pointer to **int32** | MX preference when type is MX; SRV priority when type is SRV | [optional] 
**Weight** | Pointer to **int32** | SRV weight when type is SRV | [optional] 
**Port** | Pointer to **int32** | SRV port when type is SRV | [optional] 

## Methods

### NewDnsRecordInfo

`func NewDnsRecordInfo() *DnsRecordInfo`

NewDnsRecordInfo instantiates a new DnsRecordInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsRecordInfoWithDefaults

`func NewDnsRecordInfoWithDefaults() *DnsRecordInfo`

NewDnsRecordInfoWithDefaults instantiates a new DnsRecordInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DnsRecordInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DnsRecordInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DnsRecordInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DnsRecordInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DnsRecordInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DnsRecordInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DnsRecordInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DnsRecordInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *DnsRecordInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DnsRecordInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DnsRecordInfo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DnsRecordInfo) HasType() bool`

HasType returns a boolean if a field has been set.

### GetContent

`func (o *DnsRecordInfo) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *DnsRecordInfo) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *DnsRecordInfo) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *DnsRecordInfo) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetTtl

`func (o *DnsRecordInfo) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *DnsRecordInfo) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *DnsRecordInfo) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *DnsRecordInfo) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetPriority

`func (o *DnsRecordInfo) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *DnsRecordInfo) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *DnsRecordInfo) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *DnsRecordInfo) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetWeight

`func (o *DnsRecordInfo) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *DnsRecordInfo) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *DnsRecordInfo) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *DnsRecordInfo) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetPort

`func (o *DnsRecordInfo) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *DnsRecordInfo) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *DnsRecordInfo) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *DnsRecordInfo) HasPort() bool`

HasPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


