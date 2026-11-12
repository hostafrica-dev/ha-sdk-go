# DomainRedirectDnsStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ZoneMissing** | Pointer to **bool** | True when the domain has no DNS zone yet | [optional] 
**HasARecord** | Pointer to **bool** | True when @ and www A records point at the redirector | [optional] 
**OverrideRequired** | Pointer to **bool** | True when conflicting @/www A records require override&#x3D;true on add-redirect | [optional] 
**ConflictingRecords** | Pointer to [**[]DomainRedirectConflictingRecord**](DomainRedirectConflictingRecord.md) | Conflicting A records that would be replaced when override&#x3D;true | [optional] 

## Methods

### NewDomainRedirectDnsStatus

`func NewDomainRedirectDnsStatus() *DomainRedirectDnsStatus`

NewDomainRedirectDnsStatus instantiates a new DomainRedirectDnsStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainRedirectDnsStatusWithDefaults

`func NewDomainRedirectDnsStatusWithDefaults() *DomainRedirectDnsStatus`

NewDomainRedirectDnsStatusWithDefaults instantiates a new DomainRedirectDnsStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetZoneMissing

`func (o *DomainRedirectDnsStatus) GetZoneMissing() bool`

GetZoneMissing returns the ZoneMissing field if non-nil, zero value otherwise.

### GetZoneMissingOk

`func (o *DomainRedirectDnsStatus) GetZoneMissingOk() (*bool, bool)`

GetZoneMissingOk returns a tuple with the ZoneMissing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneMissing

`func (o *DomainRedirectDnsStatus) SetZoneMissing(v bool)`

SetZoneMissing sets ZoneMissing field to given value.

### HasZoneMissing

`func (o *DomainRedirectDnsStatus) HasZoneMissing() bool`

HasZoneMissing returns a boolean if a field has been set.

### GetHasARecord

`func (o *DomainRedirectDnsStatus) GetHasARecord() bool`

GetHasARecord returns the HasARecord field if non-nil, zero value otherwise.

### GetHasARecordOk

`func (o *DomainRedirectDnsStatus) GetHasARecordOk() (*bool, bool)`

GetHasARecordOk returns a tuple with the HasARecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasARecord

`func (o *DomainRedirectDnsStatus) SetHasARecord(v bool)`

SetHasARecord sets HasARecord field to given value.

### HasHasARecord

`func (o *DomainRedirectDnsStatus) HasHasARecord() bool`

HasHasARecord returns a boolean if a field has been set.

### GetOverrideRequired

`func (o *DomainRedirectDnsStatus) GetOverrideRequired() bool`

GetOverrideRequired returns the OverrideRequired field if non-nil, zero value otherwise.

### GetOverrideRequiredOk

`func (o *DomainRedirectDnsStatus) GetOverrideRequiredOk() (*bool, bool)`

GetOverrideRequiredOk returns a tuple with the OverrideRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideRequired

`func (o *DomainRedirectDnsStatus) SetOverrideRequired(v bool)`

SetOverrideRequired sets OverrideRequired field to given value.

### HasOverrideRequired

`func (o *DomainRedirectDnsStatus) HasOverrideRequired() bool`

HasOverrideRequired returns a boolean if a field has been set.

### GetConflictingRecords

`func (o *DomainRedirectDnsStatus) GetConflictingRecords() []DomainRedirectConflictingRecord`

GetConflictingRecords returns the ConflictingRecords field if non-nil, zero value otherwise.

### GetConflictingRecordsOk

`func (o *DomainRedirectDnsStatus) GetConflictingRecordsOk() (*[]DomainRedirectConflictingRecord, bool)`

GetConflictingRecordsOk returns a tuple with the ConflictingRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflictingRecords

`func (o *DomainRedirectDnsStatus) SetConflictingRecords(v []DomainRedirectConflictingRecord)`

SetConflictingRecords sets ConflictingRecords field to given value.

### HasConflictingRecords

`func (o *DomainRedirectDnsStatus) HasConflictingRecords() bool`

HasConflictingRecords returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


