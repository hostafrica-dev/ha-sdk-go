# DomainRedirectConflictingRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Line** | Pointer to **int32** | Zone line / record id when present | [optional] 
**Name** | Pointer to **string** | Record name (e.g. @, www, or FQDN) | [optional] 
**Type** | Pointer to **string** | Record type (typically A) | [optional] 
**Value** | Pointer to **string** | Record IP value when present | [optional] 

## Methods

### NewDomainRedirectConflictingRecord

`func NewDomainRedirectConflictingRecord() *DomainRedirectConflictingRecord`

NewDomainRedirectConflictingRecord instantiates a new DomainRedirectConflictingRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainRedirectConflictingRecordWithDefaults

`func NewDomainRedirectConflictingRecordWithDefaults() *DomainRedirectConflictingRecord`

NewDomainRedirectConflictingRecordWithDefaults instantiates a new DomainRedirectConflictingRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLine

`func (o *DomainRedirectConflictingRecord) GetLine() int32`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *DomainRedirectConflictingRecord) GetLineOk() (*int32, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *DomainRedirectConflictingRecord) SetLine(v int32)`

SetLine sets Line field to given value.

### HasLine

`func (o *DomainRedirectConflictingRecord) HasLine() bool`

HasLine returns a boolean if a field has been set.

### GetName

`func (o *DomainRedirectConflictingRecord) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DomainRedirectConflictingRecord) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DomainRedirectConflictingRecord) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DomainRedirectConflictingRecord) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *DomainRedirectConflictingRecord) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DomainRedirectConflictingRecord) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DomainRedirectConflictingRecord) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DomainRedirectConflictingRecord) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *DomainRedirectConflictingRecord) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DomainRedirectConflictingRecord) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DomainRedirectConflictingRecord) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *DomainRedirectConflictingRecord) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


