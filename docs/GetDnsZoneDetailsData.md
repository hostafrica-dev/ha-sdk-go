# GetDnsZoneDetailsData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**DomainId** | Pointer to **string** | Domain service id | [optional] 
**ZoneId** | Pointer to **string** | DNS zone identifier | [optional] 
**ZoneExists** | Pointer to **bool** | True when a DNS zone exists for this domain | [optional] 
**ManagementAvailable** | Pointer to **bool** | True when DNS management is available for this domain | [optional] 
**DomainNameservers** | Pointer to **[]string** | Configured domain nameservers | [optional] 
**NsChanging** | Pointer to **string** | Nameserver change state (e.g. own, pending) | [optional] 
**PackageSettings** | Pointer to **interface{}** | Package quotas and allowed record types from upstream | [optional] 
**Records** | Pointer to [**[]DnsRecordInfo**](DnsRecordInfo.md) | DNS records in the zone | [optional] 

## Methods

### NewGetDnsZoneDetailsData

`func NewGetDnsZoneDetailsData(message string, ) *GetDnsZoneDetailsData`

NewGetDnsZoneDetailsData instantiates a new GetDnsZoneDetailsData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDnsZoneDetailsDataWithDefaults

`func NewGetDnsZoneDetailsDataWithDefaults() *GetDnsZoneDetailsData`

NewGetDnsZoneDetailsDataWithDefaults instantiates a new GetDnsZoneDetailsData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *GetDnsZoneDetailsData) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetDnsZoneDetailsData) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetDnsZoneDetailsData) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetDomainId

`func (o *GetDnsZoneDetailsData) GetDomainId() string`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *GetDnsZoneDetailsData) GetDomainIdOk() (*string, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *GetDnsZoneDetailsData) SetDomainId(v string)`

SetDomainId sets DomainId field to given value.

### HasDomainId

`func (o *GetDnsZoneDetailsData) HasDomainId() bool`

HasDomainId returns a boolean if a field has been set.

### GetZoneId

`func (o *GetDnsZoneDetailsData) GetZoneId() string`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *GetDnsZoneDetailsData) GetZoneIdOk() (*string, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *GetDnsZoneDetailsData) SetZoneId(v string)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *GetDnsZoneDetailsData) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

### GetZoneExists

`func (o *GetDnsZoneDetailsData) GetZoneExists() bool`

GetZoneExists returns the ZoneExists field if non-nil, zero value otherwise.

### GetZoneExistsOk

`func (o *GetDnsZoneDetailsData) GetZoneExistsOk() (*bool, bool)`

GetZoneExistsOk returns a tuple with the ZoneExists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneExists

`func (o *GetDnsZoneDetailsData) SetZoneExists(v bool)`

SetZoneExists sets ZoneExists field to given value.

### HasZoneExists

`func (o *GetDnsZoneDetailsData) HasZoneExists() bool`

HasZoneExists returns a boolean if a field has been set.

### GetManagementAvailable

`func (o *GetDnsZoneDetailsData) GetManagementAvailable() bool`

GetManagementAvailable returns the ManagementAvailable field if non-nil, zero value otherwise.

### GetManagementAvailableOk

`func (o *GetDnsZoneDetailsData) GetManagementAvailableOk() (*bool, bool)`

GetManagementAvailableOk returns a tuple with the ManagementAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementAvailable

`func (o *GetDnsZoneDetailsData) SetManagementAvailable(v bool)`

SetManagementAvailable sets ManagementAvailable field to given value.

### HasManagementAvailable

`func (o *GetDnsZoneDetailsData) HasManagementAvailable() bool`

HasManagementAvailable returns a boolean if a field has been set.

### GetDomainNameservers

`func (o *GetDnsZoneDetailsData) GetDomainNameservers() []string`

GetDomainNameservers returns the DomainNameservers field if non-nil, zero value otherwise.

### GetDomainNameserversOk

`func (o *GetDnsZoneDetailsData) GetDomainNameserversOk() (*[]string, bool)`

GetDomainNameserversOk returns a tuple with the DomainNameservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainNameservers

`func (o *GetDnsZoneDetailsData) SetDomainNameservers(v []string)`

SetDomainNameservers sets DomainNameservers field to given value.

### HasDomainNameservers

`func (o *GetDnsZoneDetailsData) HasDomainNameservers() bool`

HasDomainNameservers returns a boolean if a field has been set.

### GetNsChanging

`func (o *GetDnsZoneDetailsData) GetNsChanging() string`

GetNsChanging returns the NsChanging field if non-nil, zero value otherwise.

### GetNsChangingOk

`func (o *GetDnsZoneDetailsData) GetNsChangingOk() (*string, bool)`

GetNsChangingOk returns a tuple with the NsChanging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsChanging

`func (o *GetDnsZoneDetailsData) SetNsChanging(v string)`

SetNsChanging sets NsChanging field to given value.

### HasNsChanging

`func (o *GetDnsZoneDetailsData) HasNsChanging() bool`

HasNsChanging returns a boolean if a field has been set.

### GetPackageSettings

`func (o *GetDnsZoneDetailsData) GetPackageSettings() interface{}`

GetPackageSettings returns the PackageSettings field if non-nil, zero value otherwise.

### GetPackageSettingsOk

`func (o *GetDnsZoneDetailsData) GetPackageSettingsOk() (*interface{}, bool)`

GetPackageSettingsOk returns a tuple with the PackageSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageSettings

`func (o *GetDnsZoneDetailsData) SetPackageSettings(v interface{})`

SetPackageSettings sets PackageSettings field to given value.

### HasPackageSettings

`func (o *GetDnsZoneDetailsData) HasPackageSettings() bool`

HasPackageSettings returns a boolean if a field has been set.

### SetPackageSettingsNil

`func (o *GetDnsZoneDetailsData) SetPackageSettingsNil(b bool)`

 SetPackageSettingsNil sets the value for PackageSettings to be an explicit nil

### UnsetPackageSettings
`func (o *GetDnsZoneDetailsData) UnsetPackageSettings()`

UnsetPackageSettings ensures that no value is present for PackageSettings, not even an explicit nil
### GetRecords

`func (o *GetDnsZoneDetailsData) GetRecords() []DnsRecordInfo`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *GetDnsZoneDetailsData) GetRecordsOk() (*[]DnsRecordInfo, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *GetDnsZoneDetailsData) SetRecords(v []DnsRecordInfo)`

SetRecords sets Records field to given value.

### HasRecords

`func (o *GetDnsZoneDetailsData) HasRecords() bool`

HasRecords returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


