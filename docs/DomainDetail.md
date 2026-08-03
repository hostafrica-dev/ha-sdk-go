# DomainDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainId** | **string** | Domain service id | 
**Type** | **string** | Domain operation type (e.g. Register, Transfer) | 
**Domain** | **string** | Fully qualified domain name | 
**Sld** | Pointer to **string** | Second-level domain label (e.g. dandelyn) | [optional] 
**Tld** | Pointer to **string** | Top-level domain including dot (e.g. .co.za) | [optional] 
**Status** | **string** | Domain status (e.g. Active, Expired, Cancelled) | 
**Period** | **int32** | Registration or billing period in years | 
**Registrationdate** | Pointer to **string** | Registration timestamp from upstream (ISO 8601) | [optional] 
**Donotrenew** | **int32** | Renewal flag: 0 &#x3D; renew, 1 &#x3D; do not renew | 
**IdProtection** | **int32** | WHOIS privacy enabled: 0 &#x3D; off, 1 &#x3D; on | 
**IdProtectionSupported** | **bool** | Whether the TLD supports ID protection | 
**Firstpaymentamount** | Pointer to **string** | First payment amount as a decimal string | [optional] 
**Recurringamount** | **string** | Recurring amount as a decimal string (e.g. \&quot;149.99\&quot;) | 
**Dnsmanagement** | Pointer to **bool** | Whether DNS management addon is currently enabled | [optional] 
**Emailforwarding** | Pointer to **bool** | Whether email forwarding addon is currently enabled | [optional] 
**IsPremium** | Pointer to **bool** | Whether the domain is premium | [optional] 
**LockStatus** | Pointer to **string** | Registrar lock status (e.g. locked, unlocked, unavailable) | [optional] 
**GracePeriod** | Pointer to **int32** | Grace period length in days | [optional] 
**RedemptionPeriod** | Pointer to **int32** | Redemption period length in days | [optional] 
**GracePeriodFee** | Pointer to **int32** | Grace period renewal fee | [optional] 
**RedemptionPeriodFee** | Pointer to **int32** | Redemption period renewal fee | [optional] 
**InGrace** | Pointer to **bool** | Whether the domain is currently in grace period | [optional] 
**InRedemption** | Pointer to **bool** | Whether the domain is currently in redemption period | [optional] 
**Expirydate** | Pointer to **string** | Domain expiry date (YYYY-MM-DD or ISO 8601 from upstream) | [optional] 
**Nextinvoicedate** | Pointer to **string** | Next invoice date (YYYY-MM-DD or ISO 8601 from upstream) | [optional] 
**Nextduedate** | Pointer to **string** | Next due date (YYYY-MM-DD or ISO 8601 from upstream) | [optional] 
**AvailableFeatures** | Pointer to [**DomainAvailableFeatures**](DomainAvailableFeatures.md) |  | [optional] 
**DomainNameservers** | Pointer to **[]string** | Currently applied nameserver hostnames | [optional] 
**DefaultNameservers** | Pointer to **[]string** | Default nameserver hostnames for this domain/product | [optional] 
**NsChanging** | Pointer to **string** | Nameserver change state (e.g. own, pending) | [optional] 
**HasHosting** | Pointer to [**DomainHostingLink**](DomainHostingLink.md) |  | [optional] 
**HasDnsManagerZone** | **bool** | Whether a DNS Manager zone exists for this domain name | 
**ExpiryCountdown** | Pointer to [**DomainExpiryCountdown**](DomainExpiryCountdown.md) |  | [optional] 
**Evaluation** | Pointer to **interface{}** | Domain evaluator result when enabled; null when unavailable | [optional] 
**DomainEvaluationAvailable** | Pointer to **bool** | Whether domain evaluation is available for this domain/account | [optional] 
**HasRedirect** | Pointer to **interface{}** | Configured redirect details when present; null when none | [optional] 
**NoEpp** | Pointer to **bool** | True when EPP/auth code retrieval is disabled for this domain | [optional] 

## Methods

### NewDomainDetail

`func NewDomainDetail(domainId string, type_ string, domain string, status string, period int32, donotrenew int32, idProtection int32, idProtectionSupported bool, recurringamount string, hasDnsManagerZone bool, ) *DomainDetail`

NewDomainDetail instantiates a new DomainDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainDetailWithDefaults

`func NewDomainDetailWithDefaults() *DomainDetail`

NewDomainDetailWithDefaults instantiates a new DomainDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomainId

`func (o *DomainDetail) GetDomainId() string`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *DomainDetail) GetDomainIdOk() (*string, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *DomainDetail) SetDomainId(v string)`

SetDomainId sets DomainId field to given value.


### GetType

`func (o *DomainDetail) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DomainDetail) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DomainDetail) SetType(v string)`

SetType sets Type field to given value.


### GetDomain

`func (o *DomainDetail) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *DomainDetail) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *DomainDetail) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetSld

`func (o *DomainDetail) GetSld() string`

GetSld returns the Sld field if non-nil, zero value otherwise.

### GetSldOk

`func (o *DomainDetail) GetSldOk() (*string, bool)`

GetSldOk returns a tuple with the Sld field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSld

`func (o *DomainDetail) SetSld(v string)`

SetSld sets Sld field to given value.

### HasSld

`func (o *DomainDetail) HasSld() bool`

HasSld returns a boolean if a field has been set.

### GetTld

`func (o *DomainDetail) GetTld() string`

GetTld returns the Tld field if non-nil, zero value otherwise.

### GetTldOk

`func (o *DomainDetail) GetTldOk() (*string, bool)`

GetTldOk returns a tuple with the Tld field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTld

`func (o *DomainDetail) SetTld(v string)`

SetTld sets Tld field to given value.

### HasTld

`func (o *DomainDetail) HasTld() bool`

HasTld returns a boolean if a field has been set.

### GetStatus

`func (o *DomainDetail) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DomainDetail) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DomainDetail) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetPeriod

`func (o *DomainDetail) GetPeriod() int32`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *DomainDetail) GetPeriodOk() (*int32, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *DomainDetail) SetPeriod(v int32)`

SetPeriod sets Period field to given value.


### GetRegistrationdate

`func (o *DomainDetail) GetRegistrationdate() string`

GetRegistrationdate returns the Registrationdate field if non-nil, zero value otherwise.

### GetRegistrationdateOk

`func (o *DomainDetail) GetRegistrationdateOk() (*string, bool)`

GetRegistrationdateOk returns a tuple with the Registrationdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationdate

`func (o *DomainDetail) SetRegistrationdate(v string)`

SetRegistrationdate sets Registrationdate field to given value.

### HasRegistrationdate

`func (o *DomainDetail) HasRegistrationdate() bool`

HasRegistrationdate returns a boolean if a field has been set.

### GetDonotrenew

`func (o *DomainDetail) GetDonotrenew() int32`

GetDonotrenew returns the Donotrenew field if non-nil, zero value otherwise.

### GetDonotrenewOk

`func (o *DomainDetail) GetDonotrenewOk() (*int32, bool)`

GetDonotrenewOk returns a tuple with the Donotrenew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDonotrenew

`func (o *DomainDetail) SetDonotrenew(v int32)`

SetDonotrenew sets Donotrenew field to given value.


### GetIdProtection

`func (o *DomainDetail) GetIdProtection() int32`

GetIdProtection returns the IdProtection field if non-nil, zero value otherwise.

### GetIdProtectionOk

`func (o *DomainDetail) GetIdProtectionOk() (*int32, bool)`

GetIdProtectionOk returns a tuple with the IdProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdProtection

`func (o *DomainDetail) SetIdProtection(v int32)`

SetIdProtection sets IdProtection field to given value.


### GetIdProtectionSupported

`func (o *DomainDetail) GetIdProtectionSupported() bool`

GetIdProtectionSupported returns the IdProtectionSupported field if non-nil, zero value otherwise.

### GetIdProtectionSupportedOk

`func (o *DomainDetail) GetIdProtectionSupportedOk() (*bool, bool)`

GetIdProtectionSupportedOk returns a tuple with the IdProtectionSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdProtectionSupported

`func (o *DomainDetail) SetIdProtectionSupported(v bool)`

SetIdProtectionSupported sets IdProtectionSupported field to given value.


### GetFirstpaymentamount

`func (o *DomainDetail) GetFirstpaymentamount() string`

GetFirstpaymentamount returns the Firstpaymentamount field if non-nil, zero value otherwise.

### GetFirstpaymentamountOk

`func (o *DomainDetail) GetFirstpaymentamountOk() (*string, bool)`

GetFirstpaymentamountOk returns a tuple with the Firstpaymentamount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstpaymentamount

`func (o *DomainDetail) SetFirstpaymentamount(v string)`

SetFirstpaymentamount sets Firstpaymentamount field to given value.

### HasFirstpaymentamount

`func (o *DomainDetail) HasFirstpaymentamount() bool`

HasFirstpaymentamount returns a boolean if a field has been set.

### GetRecurringamount

`func (o *DomainDetail) GetRecurringamount() string`

GetRecurringamount returns the Recurringamount field if non-nil, zero value otherwise.

### GetRecurringamountOk

`func (o *DomainDetail) GetRecurringamountOk() (*string, bool)`

GetRecurringamountOk returns a tuple with the Recurringamount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringamount

`func (o *DomainDetail) SetRecurringamount(v string)`

SetRecurringamount sets Recurringamount field to given value.


### GetDnsmanagement

`func (o *DomainDetail) GetDnsmanagement() bool`

GetDnsmanagement returns the Dnsmanagement field if non-nil, zero value otherwise.

### GetDnsmanagementOk

`func (o *DomainDetail) GetDnsmanagementOk() (*bool, bool)`

GetDnsmanagementOk returns a tuple with the Dnsmanagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsmanagement

`func (o *DomainDetail) SetDnsmanagement(v bool)`

SetDnsmanagement sets Dnsmanagement field to given value.

### HasDnsmanagement

`func (o *DomainDetail) HasDnsmanagement() bool`

HasDnsmanagement returns a boolean if a field has been set.

### GetEmailforwarding

`func (o *DomainDetail) GetEmailforwarding() bool`

GetEmailforwarding returns the Emailforwarding field if non-nil, zero value otherwise.

### GetEmailforwardingOk

`func (o *DomainDetail) GetEmailforwardingOk() (*bool, bool)`

GetEmailforwardingOk returns a tuple with the Emailforwarding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailforwarding

`func (o *DomainDetail) SetEmailforwarding(v bool)`

SetEmailforwarding sets Emailforwarding field to given value.

### HasEmailforwarding

`func (o *DomainDetail) HasEmailforwarding() bool`

HasEmailforwarding returns a boolean if a field has been set.

### GetIsPremium

`func (o *DomainDetail) GetIsPremium() bool`

GetIsPremium returns the IsPremium field if non-nil, zero value otherwise.

### GetIsPremiumOk

`func (o *DomainDetail) GetIsPremiumOk() (*bool, bool)`

GetIsPremiumOk returns a tuple with the IsPremium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPremium

`func (o *DomainDetail) SetIsPremium(v bool)`

SetIsPremium sets IsPremium field to given value.

### HasIsPremium

`func (o *DomainDetail) HasIsPremium() bool`

HasIsPremium returns a boolean if a field has been set.

### GetLockStatus

`func (o *DomainDetail) GetLockStatus() string`

GetLockStatus returns the LockStatus field if non-nil, zero value otherwise.

### GetLockStatusOk

`func (o *DomainDetail) GetLockStatusOk() (*string, bool)`

GetLockStatusOk returns a tuple with the LockStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockStatus

`func (o *DomainDetail) SetLockStatus(v string)`

SetLockStatus sets LockStatus field to given value.

### HasLockStatus

`func (o *DomainDetail) HasLockStatus() bool`

HasLockStatus returns a boolean if a field has been set.

### GetGracePeriod

`func (o *DomainDetail) GetGracePeriod() int32`

GetGracePeriod returns the GracePeriod field if non-nil, zero value otherwise.

### GetGracePeriodOk

`func (o *DomainDetail) GetGracePeriodOk() (*int32, bool)`

GetGracePeriodOk returns a tuple with the GracePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGracePeriod

`func (o *DomainDetail) SetGracePeriod(v int32)`

SetGracePeriod sets GracePeriod field to given value.

### HasGracePeriod

`func (o *DomainDetail) HasGracePeriod() bool`

HasGracePeriod returns a boolean if a field has been set.

### GetRedemptionPeriod

`func (o *DomainDetail) GetRedemptionPeriod() int32`

GetRedemptionPeriod returns the RedemptionPeriod field if non-nil, zero value otherwise.

### GetRedemptionPeriodOk

`func (o *DomainDetail) GetRedemptionPeriodOk() (*int32, bool)`

GetRedemptionPeriodOk returns a tuple with the RedemptionPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedemptionPeriod

`func (o *DomainDetail) SetRedemptionPeriod(v int32)`

SetRedemptionPeriod sets RedemptionPeriod field to given value.

### HasRedemptionPeriod

`func (o *DomainDetail) HasRedemptionPeriod() bool`

HasRedemptionPeriod returns a boolean if a field has been set.

### GetGracePeriodFee

`func (o *DomainDetail) GetGracePeriodFee() int32`

GetGracePeriodFee returns the GracePeriodFee field if non-nil, zero value otherwise.

### GetGracePeriodFeeOk

`func (o *DomainDetail) GetGracePeriodFeeOk() (*int32, bool)`

GetGracePeriodFeeOk returns a tuple with the GracePeriodFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGracePeriodFee

`func (o *DomainDetail) SetGracePeriodFee(v int32)`

SetGracePeriodFee sets GracePeriodFee field to given value.

### HasGracePeriodFee

`func (o *DomainDetail) HasGracePeriodFee() bool`

HasGracePeriodFee returns a boolean if a field has been set.

### GetRedemptionPeriodFee

`func (o *DomainDetail) GetRedemptionPeriodFee() int32`

GetRedemptionPeriodFee returns the RedemptionPeriodFee field if non-nil, zero value otherwise.

### GetRedemptionPeriodFeeOk

`func (o *DomainDetail) GetRedemptionPeriodFeeOk() (*int32, bool)`

GetRedemptionPeriodFeeOk returns a tuple with the RedemptionPeriodFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedemptionPeriodFee

`func (o *DomainDetail) SetRedemptionPeriodFee(v int32)`

SetRedemptionPeriodFee sets RedemptionPeriodFee field to given value.

### HasRedemptionPeriodFee

`func (o *DomainDetail) HasRedemptionPeriodFee() bool`

HasRedemptionPeriodFee returns a boolean if a field has been set.

### GetInGrace

`func (o *DomainDetail) GetInGrace() bool`

GetInGrace returns the InGrace field if non-nil, zero value otherwise.

### GetInGraceOk

`func (o *DomainDetail) GetInGraceOk() (*bool, bool)`

GetInGraceOk returns a tuple with the InGrace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInGrace

`func (o *DomainDetail) SetInGrace(v bool)`

SetInGrace sets InGrace field to given value.

### HasInGrace

`func (o *DomainDetail) HasInGrace() bool`

HasInGrace returns a boolean if a field has been set.

### GetInRedemption

`func (o *DomainDetail) GetInRedemption() bool`

GetInRedemption returns the InRedemption field if non-nil, zero value otherwise.

### GetInRedemptionOk

`func (o *DomainDetail) GetInRedemptionOk() (*bool, bool)`

GetInRedemptionOk returns a tuple with the InRedemption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInRedemption

`func (o *DomainDetail) SetInRedemption(v bool)`

SetInRedemption sets InRedemption field to given value.

### HasInRedemption

`func (o *DomainDetail) HasInRedemption() bool`

HasInRedemption returns a boolean if a field has been set.

### GetExpirydate

`func (o *DomainDetail) GetExpirydate() string`

GetExpirydate returns the Expirydate field if non-nil, zero value otherwise.

### GetExpirydateOk

`func (o *DomainDetail) GetExpirydateOk() (*string, bool)`

GetExpirydateOk returns a tuple with the Expirydate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirydate

`func (o *DomainDetail) SetExpirydate(v string)`

SetExpirydate sets Expirydate field to given value.

### HasExpirydate

`func (o *DomainDetail) HasExpirydate() bool`

HasExpirydate returns a boolean if a field has been set.

### GetNextinvoicedate

`func (o *DomainDetail) GetNextinvoicedate() string`

GetNextinvoicedate returns the Nextinvoicedate field if non-nil, zero value otherwise.

### GetNextinvoicedateOk

`func (o *DomainDetail) GetNextinvoicedateOk() (*string, bool)`

GetNextinvoicedateOk returns a tuple with the Nextinvoicedate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextinvoicedate

`func (o *DomainDetail) SetNextinvoicedate(v string)`

SetNextinvoicedate sets Nextinvoicedate field to given value.

### HasNextinvoicedate

`func (o *DomainDetail) HasNextinvoicedate() bool`

HasNextinvoicedate returns a boolean if a field has been set.

### GetNextduedate

`func (o *DomainDetail) GetNextduedate() string`

GetNextduedate returns the Nextduedate field if non-nil, zero value otherwise.

### GetNextduedateOk

`func (o *DomainDetail) GetNextduedateOk() (*string, bool)`

GetNextduedateOk returns a tuple with the Nextduedate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextduedate

`func (o *DomainDetail) SetNextduedate(v string)`

SetNextduedate sets Nextduedate field to given value.

### HasNextduedate

`func (o *DomainDetail) HasNextduedate() bool`

HasNextduedate returns a boolean if a field has been set.

### GetAvailableFeatures

`func (o *DomainDetail) GetAvailableFeatures() DomainAvailableFeatures`

GetAvailableFeatures returns the AvailableFeatures field if non-nil, zero value otherwise.

### GetAvailableFeaturesOk

`func (o *DomainDetail) GetAvailableFeaturesOk() (*DomainAvailableFeatures, bool)`

GetAvailableFeaturesOk returns a tuple with the AvailableFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableFeatures

`func (o *DomainDetail) SetAvailableFeatures(v DomainAvailableFeatures)`

SetAvailableFeatures sets AvailableFeatures field to given value.

### HasAvailableFeatures

`func (o *DomainDetail) HasAvailableFeatures() bool`

HasAvailableFeatures returns a boolean if a field has been set.

### GetDomainNameservers

`func (o *DomainDetail) GetDomainNameservers() []string`

GetDomainNameservers returns the DomainNameservers field if non-nil, zero value otherwise.

### GetDomainNameserversOk

`func (o *DomainDetail) GetDomainNameserversOk() (*[]string, bool)`

GetDomainNameserversOk returns a tuple with the DomainNameservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainNameservers

`func (o *DomainDetail) SetDomainNameservers(v []string)`

SetDomainNameservers sets DomainNameservers field to given value.

### HasDomainNameservers

`func (o *DomainDetail) HasDomainNameservers() bool`

HasDomainNameservers returns a boolean if a field has been set.

### GetDefaultNameservers

`func (o *DomainDetail) GetDefaultNameservers() []string`

GetDefaultNameservers returns the DefaultNameservers field if non-nil, zero value otherwise.

### GetDefaultNameserversOk

`func (o *DomainDetail) GetDefaultNameserversOk() (*[]string, bool)`

GetDefaultNameserversOk returns a tuple with the DefaultNameservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultNameservers

`func (o *DomainDetail) SetDefaultNameservers(v []string)`

SetDefaultNameservers sets DefaultNameservers field to given value.

### HasDefaultNameservers

`func (o *DomainDetail) HasDefaultNameservers() bool`

HasDefaultNameservers returns a boolean if a field has been set.

### GetNsChanging

`func (o *DomainDetail) GetNsChanging() string`

GetNsChanging returns the NsChanging field if non-nil, zero value otherwise.

### GetNsChangingOk

`func (o *DomainDetail) GetNsChangingOk() (*string, bool)`

GetNsChangingOk returns a tuple with the NsChanging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsChanging

`func (o *DomainDetail) SetNsChanging(v string)`

SetNsChanging sets NsChanging field to given value.

### HasNsChanging

`func (o *DomainDetail) HasNsChanging() bool`

HasNsChanging returns a boolean if a field has been set.

### GetHasHosting

`func (o *DomainDetail) GetHasHosting() DomainHostingLink`

GetHasHosting returns the HasHosting field if non-nil, zero value otherwise.

### GetHasHostingOk

`func (o *DomainDetail) GetHasHostingOk() (*DomainHostingLink, bool)`

GetHasHostingOk returns a tuple with the HasHosting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasHosting

`func (o *DomainDetail) SetHasHosting(v DomainHostingLink)`

SetHasHosting sets HasHosting field to given value.

### HasHasHosting

`func (o *DomainDetail) HasHasHosting() bool`

HasHasHosting returns a boolean if a field has been set.

### GetHasDnsManagerZone

`func (o *DomainDetail) GetHasDnsManagerZone() bool`

GetHasDnsManagerZone returns the HasDnsManagerZone field if non-nil, zero value otherwise.

### GetHasDnsManagerZoneOk

`func (o *DomainDetail) GetHasDnsManagerZoneOk() (*bool, bool)`

GetHasDnsManagerZoneOk returns a tuple with the HasDnsManagerZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDnsManagerZone

`func (o *DomainDetail) SetHasDnsManagerZone(v bool)`

SetHasDnsManagerZone sets HasDnsManagerZone field to given value.


### GetExpiryCountdown

`func (o *DomainDetail) GetExpiryCountdown() DomainExpiryCountdown`

GetExpiryCountdown returns the ExpiryCountdown field if non-nil, zero value otherwise.

### GetExpiryCountdownOk

`func (o *DomainDetail) GetExpiryCountdownOk() (*DomainExpiryCountdown, bool)`

GetExpiryCountdownOk returns a tuple with the ExpiryCountdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryCountdown

`func (o *DomainDetail) SetExpiryCountdown(v DomainExpiryCountdown)`

SetExpiryCountdown sets ExpiryCountdown field to given value.

### HasExpiryCountdown

`func (o *DomainDetail) HasExpiryCountdown() bool`

HasExpiryCountdown returns a boolean if a field has been set.

### GetEvaluation

`func (o *DomainDetail) GetEvaluation() interface{}`

GetEvaluation returns the Evaluation field if non-nil, zero value otherwise.

### GetEvaluationOk

`func (o *DomainDetail) GetEvaluationOk() (*interface{}, bool)`

GetEvaluationOk returns a tuple with the Evaluation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluation

`func (o *DomainDetail) SetEvaluation(v interface{})`

SetEvaluation sets Evaluation field to given value.

### HasEvaluation

`func (o *DomainDetail) HasEvaluation() bool`

HasEvaluation returns a boolean if a field has been set.

### SetEvaluationNil

`func (o *DomainDetail) SetEvaluationNil(b bool)`

 SetEvaluationNil sets the value for Evaluation to be an explicit nil

### UnsetEvaluation
`func (o *DomainDetail) UnsetEvaluation()`

UnsetEvaluation ensures that no value is present for Evaluation, not even an explicit nil
### GetDomainEvaluationAvailable

`func (o *DomainDetail) GetDomainEvaluationAvailable() bool`

GetDomainEvaluationAvailable returns the DomainEvaluationAvailable field if non-nil, zero value otherwise.

### GetDomainEvaluationAvailableOk

`func (o *DomainDetail) GetDomainEvaluationAvailableOk() (*bool, bool)`

GetDomainEvaluationAvailableOk returns a tuple with the DomainEvaluationAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainEvaluationAvailable

`func (o *DomainDetail) SetDomainEvaluationAvailable(v bool)`

SetDomainEvaluationAvailable sets DomainEvaluationAvailable field to given value.

### HasDomainEvaluationAvailable

`func (o *DomainDetail) HasDomainEvaluationAvailable() bool`

HasDomainEvaluationAvailable returns a boolean if a field has been set.

### GetHasRedirect

`func (o *DomainDetail) GetHasRedirect() interface{}`

GetHasRedirect returns the HasRedirect field if non-nil, zero value otherwise.

### GetHasRedirectOk

`func (o *DomainDetail) GetHasRedirectOk() (*interface{}, bool)`

GetHasRedirectOk returns a tuple with the HasRedirect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasRedirect

`func (o *DomainDetail) SetHasRedirect(v interface{})`

SetHasRedirect sets HasRedirect field to given value.

### HasHasRedirect

`func (o *DomainDetail) HasHasRedirect() bool`

HasHasRedirect returns a boolean if a field has been set.

### SetHasRedirectNil

`func (o *DomainDetail) SetHasRedirectNil(b bool)`

 SetHasRedirectNil sets the value for HasRedirect to be an explicit nil

### UnsetHasRedirect
`func (o *DomainDetail) UnsetHasRedirect()`

UnsetHasRedirect ensures that no value is present for HasRedirect, not even an explicit nil
### GetNoEpp

`func (o *DomainDetail) GetNoEpp() bool`

GetNoEpp returns the NoEpp field if non-nil, zero value otherwise.

### GetNoEppOk

`func (o *DomainDetail) GetNoEppOk() (*bool, bool)`

GetNoEppOk returns a tuple with the NoEpp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoEpp

`func (o *DomainDetail) SetNoEpp(v bool)`

SetNoEpp sets NoEpp field to given value.

### HasNoEpp

`func (o *DomainDetail) HasNoEpp() bool`

HasNoEpp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


