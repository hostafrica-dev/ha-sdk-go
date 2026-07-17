# DomainInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainId** | **string** | Domain service id | 
**Type** | **string** | Domain operation type (e.g. Register, Transfer) | 
**Domain** | **string** | Fully qualified domain name | 
**Status** | **string** | Domain status (e.g. Active, Expired, Cancelled) | 
**Period** | **int32** | Registration or billing period in years | 
**Donotrenew** | **int32** | Renewal flag: 0 &#x3D; renew, 1 &#x3D; do not renew | 
**IdProtection** | **int32** | WHOIS privacy enabled: 0 &#x3D; off, 1 &#x3D; on | 
**IdProtectionSupported** | **bool** | Whether the TLD supports ID protection | 
**Recurringamount** | **string** | Recurring amount as a decimal string (e.g. \&quot;149.99\&quot;) | 
**Expirydate** | Pointer to **string** | Domain expiry date (YYYY-MM-DD) | [optional] 
**Nextinvoicedate** | Pointer to **string** | Next invoice date (YYYY-MM-DD) | [optional] 
**Nextduedate** | Pointer to **string** | Next due date (YYYY-MM-DD) | [optional] 
**HasHosting** | Pointer to [**DomainHostingLink**](DomainHostingLink.md) |  | [optional] 
**HasDnsManagerZone** | **bool** | Whether a DNS Manager zone exists for this domain name | 
**Evaluation** | Pointer to **interface{}** | Domain evaluator result when enabled; null when unavailable | [optional] 

## Methods

### NewDomainInfo

`func NewDomainInfo(domainId string, type_ string, domain string, status string, period int32, donotrenew int32, idProtection int32, idProtectionSupported bool, recurringamount string, hasDnsManagerZone bool, ) *DomainInfo`

NewDomainInfo instantiates a new DomainInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainInfoWithDefaults

`func NewDomainInfoWithDefaults() *DomainInfo`

NewDomainInfoWithDefaults instantiates a new DomainInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomainId

`func (o *DomainInfo) GetDomainId() string`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *DomainInfo) GetDomainIdOk() (*string, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *DomainInfo) SetDomainId(v string)`

SetDomainId sets DomainId field to given value.


### GetType

`func (o *DomainInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DomainInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DomainInfo) SetType(v string)`

SetType sets Type field to given value.


### GetDomain

`func (o *DomainInfo) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *DomainInfo) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *DomainInfo) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetStatus

`func (o *DomainInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DomainInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DomainInfo) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetPeriod

`func (o *DomainInfo) GetPeriod() int32`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *DomainInfo) GetPeriodOk() (*int32, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *DomainInfo) SetPeriod(v int32)`

SetPeriod sets Period field to given value.


### GetDonotrenew

`func (o *DomainInfo) GetDonotrenew() int32`

GetDonotrenew returns the Donotrenew field if non-nil, zero value otherwise.

### GetDonotrenewOk

`func (o *DomainInfo) GetDonotrenewOk() (*int32, bool)`

GetDonotrenewOk returns a tuple with the Donotrenew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDonotrenew

`func (o *DomainInfo) SetDonotrenew(v int32)`

SetDonotrenew sets Donotrenew field to given value.


### GetIdProtection

`func (o *DomainInfo) GetIdProtection() int32`

GetIdProtection returns the IdProtection field if non-nil, zero value otherwise.

### GetIdProtectionOk

`func (o *DomainInfo) GetIdProtectionOk() (*int32, bool)`

GetIdProtectionOk returns a tuple with the IdProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdProtection

`func (o *DomainInfo) SetIdProtection(v int32)`

SetIdProtection sets IdProtection field to given value.


### GetIdProtectionSupported

`func (o *DomainInfo) GetIdProtectionSupported() bool`

GetIdProtectionSupported returns the IdProtectionSupported field if non-nil, zero value otherwise.

### GetIdProtectionSupportedOk

`func (o *DomainInfo) GetIdProtectionSupportedOk() (*bool, bool)`

GetIdProtectionSupportedOk returns a tuple with the IdProtectionSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdProtectionSupported

`func (o *DomainInfo) SetIdProtectionSupported(v bool)`

SetIdProtectionSupported sets IdProtectionSupported field to given value.


### GetRecurringamount

`func (o *DomainInfo) GetRecurringamount() string`

GetRecurringamount returns the Recurringamount field if non-nil, zero value otherwise.

### GetRecurringamountOk

`func (o *DomainInfo) GetRecurringamountOk() (*string, bool)`

GetRecurringamountOk returns a tuple with the Recurringamount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringamount

`func (o *DomainInfo) SetRecurringamount(v string)`

SetRecurringamount sets Recurringamount field to given value.


### GetExpirydate

`func (o *DomainInfo) GetExpirydate() string`

GetExpirydate returns the Expirydate field if non-nil, zero value otherwise.

### GetExpirydateOk

`func (o *DomainInfo) GetExpirydateOk() (*string, bool)`

GetExpirydateOk returns a tuple with the Expirydate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirydate

`func (o *DomainInfo) SetExpirydate(v string)`

SetExpirydate sets Expirydate field to given value.

### HasExpirydate

`func (o *DomainInfo) HasExpirydate() bool`

HasExpirydate returns a boolean if a field has been set.

### GetNextinvoicedate

`func (o *DomainInfo) GetNextinvoicedate() string`

GetNextinvoicedate returns the Nextinvoicedate field if non-nil, zero value otherwise.

### GetNextinvoicedateOk

`func (o *DomainInfo) GetNextinvoicedateOk() (*string, bool)`

GetNextinvoicedateOk returns a tuple with the Nextinvoicedate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextinvoicedate

`func (o *DomainInfo) SetNextinvoicedate(v string)`

SetNextinvoicedate sets Nextinvoicedate field to given value.

### HasNextinvoicedate

`func (o *DomainInfo) HasNextinvoicedate() bool`

HasNextinvoicedate returns a boolean if a field has been set.

### GetNextduedate

`func (o *DomainInfo) GetNextduedate() string`

GetNextduedate returns the Nextduedate field if non-nil, zero value otherwise.

### GetNextduedateOk

`func (o *DomainInfo) GetNextduedateOk() (*string, bool)`

GetNextduedateOk returns a tuple with the Nextduedate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextduedate

`func (o *DomainInfo) SetNextduedate(v string)`

SetNextduedate sets Nextduedate field to given value.

### HasNextduedate

`func (o *DomainInfo) HasNextduedate() bool`

HasNextduedate returns a boolean if a field has been set.

### GetHasHosting

`func (o *DomainInfo) GetHasHosting() DomainHostingLink`

GetHasHosting returns the HasHosting field if non-nil, zero value otherwise.

### GetHasHostingOk

`func (o *DomainInfo) GetHasHostingOk() (*DomainHostingLink, bool)`

GetHasHostingOk returns a tuple with the HasHosting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasHosting

`func (o *DomainInfo) SetHasHosting(v DomainHostingLink)`

SetHasHosting sets HasHosting field to given value.

### HasHasHosting

`func (o *DomainInfo) HasHasHosting() bool`

HasHasHosting returns a boolean if a field has been set.

### GetHasDnsManagerZone

`func (o *DomainInfo) GetHasDnsManagerZone() bool`

GetHasDnsManagerZone returns the HasDnsManagerZone field if non-nil, zero value otherwise.

### GetHasDnsManagerZoneOk

`func (o *DomainInfo) GetHasDnsManagerZoneOk() (*bool, bool)`

GetHasDnsManagerZoneOk returns a tuple with the HasDnsManagerZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDnsManagerZone

`func (o *DomainInfo) SetHasDnsManagerZone(v bool)`

SetHasDnsManagerZone sets HasDnsManagerZone field to given value.


### GetEvaluation

`func (o *DomainInfo) GetEvaluation() interface{}`

GetEvaluation returns the Evaluation field if non-nil, zero value otherwise.

### GetEvaluationOk

`func (o *DomainInfo) GetEvaluationOk() (*interface{}, bool)`

GetEvaluationOk returns a tuple with the Evaluation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluation

`func (o *DomainInfo) SetEvaluation(v interface{})`

SetEvaluation sets Evaluation field to given value.

### HasEvaluation

`func (o *DomainInfo) HasEvaluation() bool`

HasEvaluation returns a boolean if a field has been set.

### SetEvaluationNil

`func (o *DomainInfo) SetEvaluationNil(b bool)`

 SetEvaluationNil sets the value for Evaluation to be an explicit nil

### UnsetEvaluation
`func (o *DomainInfo) UnsetEvaluation()`

UnsetEvaluation ensures that no value is present for Evaluation, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


