# DomainAvailableFeatures

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Eppcode** | Pointer to **bool** | Whether EPP/auth code retrieval is available | [optional] 
**Dnssec** | Pointer to **bool** | Whether DNSSEC management is available | [optional] 
**PrivateNameservers** | Pointer to **bool** | Whether private/custom nameservers are supported | [optional] 
**Redirector** | Pointer to **bool** | Whether domain redirect is available | [optional] 
**Dnsmanagement** | Pointer to [**DomainAddonFeature**](DomainAddonFeature.md) |  | [optional] 
**Emailforwarding** | Pointer to [**DomainAddonFeature**](DomainAddonFeature.md) |  | [optional] 
**Idprotection** | Pointer to [**DomainAddonFeature**](DomainAddonFeature.md) |  | [optional] 

## Methods

### NewDomainAvailableFeatures

`func NewDomainAvailableFeatures() *DomainAvailableFeatures`

NewDomainAvailableFeatures instantiates a new DomainAvailableFeatures object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainAvailableFeaturesWithDefaults

`func NewDomainAvailableFeaturesWithDefaults() *DomainAvailableFeatures`

NewDomainAvailableFeaturesWithDefaults instantiates a new DomainAvailableFeatures object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEppcode

`func (o *DomainAvailableFeatures) GetEppcode() bool`

GetEppcode returns the Eppcode field if non-nil, zero value otherwise.

### GetEppcodeOk

`func (o *DomainAvailableFeatures) GetEppcodeOk() (*bool, bool)`

GetEppcodeOk returns a tuple with the Eppcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEppcode

`func (o *DomainAvailableFeatures) SetEppcode(v bool)`

SetEppcode sets Eppcode field to given value.

### HasEppcode

`func (o *DomainAvailableFeatures) HasEppcode() bool`

HasEppcode returns a boolean if a field has been set.

### GetDnssec

`func (o *DomainAvailableFeatures) GetDnssec() bool`

GetDnssec returns the Dnssec field if non-nil, zero value otherwise.

### GetDnssecOk

`func (o *DomainAvailableFeatures) GetDnssecOk() (*bool, bool)`

GetDnssecOk returns a tuple with the Dnssec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssec

`func (o *DomainAvailableFeatures) SetDnssec(v bool)`

SetDnssec sets Dnssec field to given value.

### HasDnssec

`func (o *DomainAvailableFeatures) HasDnssec() bool`

HasDnssec returns a boolean if a field has been set.

### GetPrivateNameservers

`func (o *DomainAvailableFeatures) GetPrivateNameservers() bool`

GetPrivateNameservers returns the PrivateNameservers field if non-nil, zero value otherwise.

### GetPrivateNameserversOk

`func (o *DomainAvailableFeatures) GetPrivateNameserversOk() (*bool, bool)`

GetPrivateNameserversOk returns a tuple with the PrivateNameservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateNameservers

`func (o *DomainAvailableFeatures) SetPrivateNameservers(v bool)`

SetPrivateNameservers sets PrivateNameservers field to given value.

### HasPrivateNameservers

`func (o *DomainAvailableFeatures) HasPrivateNameservers() bool`

HasPrivateNameservers returns a boolean if a field has been set.

### GetRedirector

`func (o *DomainAvailableFeatures) GetRedirector() bool`

GetRedirector returns the Redirector field if non-nil, zero value otherwise.

### GetRedirectorOk

`func (o *DomainAvailableFeatures) GetRedirectorOk() (*bool, bool)`

GetRedirectorOk returns a tuple with the Redirector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirector

`func (o *DomainAvailableFeatures) SetRedirector(v bool)`

SetRedirector sets Redirector field to given value.

### HasRedirector

`func (o *DomainAvailableFeatures) HasRedirector() bool`

HasRedirector returns a boolean if a field has been set.

### GetDnsmanagement

`func (o *DomainAvailableFeatures) GetDnsmanagement() DomainAddonFeature`

GetDnsmanagement returns the Dnsmanagement field if non-nil, zero value otherwise.

### GetDnsmanagementOk

`func (o *DomainAvailableFeatures) GetDnsmanagementOk() (*DomainAddonFeature, bool)`

GetDnsmanagementOk returns a tuple with the Dnsmanagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsmanagement

`func (o *DomainAvailableFeatures) SetDnsmanagement(v DomainAddonFeature)`

SetDnsmanagement sets Dnsmanagement field to given value.

### HasDnsmanagement

`func (o *DomainAvailableFeatures) HasDnsmanagement() bool`

HasDnsmanagement returns a boolean if a field has been set.

### GetEmailforwarding

`func (o *DomainAvailableFeatures) GetEmailforwarding() DomainAddonFeature`

GetEmailforwarding returns the Emailforwarding field if non-nil, zero value otherwise.

### GetEmailforwardingOk

`func (o *DomainAvailableFeatures) GetEmailforwardingOk() (*DomainAddonFeature, bool)`

GetEmailforwardingOk returns a tuple with the Emailforwarding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailforwarding

`func (o *DomainAvailableFeatures) SetEmailforwarding(v DomainAddonFeature)`

SetEmailforwarding sets Emailforwarding field to given value.

### HasEmailforwarding

`func (o *DomainAvailableFeatures) HasEmailforwarding() bool`

HasEmailforwarding returns a boolean if a field has been set.

### GetIdprotection

`func (o *DomainAvailableFeatures) GetIdprotection() DomainAddonFeature`

GetIdprotection returns the Idprotection field if non-nil, zero value otherwise.

### GetIdprotectionOk

`func (o *DomainAvailableFeatures) GetIdprotectionOk() (*DomainAddonFeature, bool)`

GetIdprotectionOk returns a tuple with the Idprotection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdprotection

`func (o *DomainAvailableFeatures) SetIdprotection(v DomainAddonFeature)`

SetIdprotection sets Idprotection field to given value.

### HasIdprotection

`func (o *DomainAvailableFeatures) HasIdprotection() bool`

HasIdprotection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


