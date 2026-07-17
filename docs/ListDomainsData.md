# ListDomainsData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Status message indicating the result | 
**Domains** | [**[]DomainInfo**](DomainInfo.md) | Domains owned by the authenticated client | 
**DomainWarranty** | **bool** | Whether ID-protection warranty is enabled for the account | 
**DomainWarrantyRename** | Pointer to **string** | Domain warranty rename setting when configured | [optional] 
**DomainEvaluationAvailable** | **bool** | Whether domain evaluation is available for the account | 
**TotalCount** | **int32** | Total number of domains returned | 

## Methods

### NewListDomainsData

`func NewListDomainsData(message string, domains []DomainInfo, domainWarranty bool, domainEvaluationAvailable bool, totalCount int32, ) *ListDomainsData`

NewListDomainsData instantiates a new ListDomainsData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListDomainsDataWithDefaults

`func NewListDomainsDataWithDefaults() *ListDomainsData`

NewListDomainsDataWithDefaults instantiates a new ListDomainsData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ListDomainsData) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ListDomainsData) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ListDomainsData) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetDomains

`func (o *ListDomainsData) GetDomains() []DomainInfo`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *ListDomainsData) GetDomainsOk() (*[]DomainInfo, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *ListDomainsData) SetDomains(v []DomainInfo)`

SetDomains sets Domains field to given value.


### GetDomainWarranty

`func (o *ListDomainsData) GetDomainWarranty() bool`

GetDomainWarranty returns the DomainWarranty field if non-nil, zero value otherwise.

### GetDomainWarrantyOk

`func (o *ListDomainsData) GetDomainWarrantyOk() (*bool, bool)`

GetDomainWarrantyOk returns a tuple with the DomainWarranty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainWarranty

`func (o *ListDomainsData) SetDomainWarranty(v bool)`

SetDomainWarranty sets DomainWarranty field to given value.


### GetDomainWarrantyRename

`func (o *ListDomainsData) GetDomainWarrantyRename() string`

GetDomainWarrantyRename returns the DomainWarrantyRename field if non-nil, zero value otherwise.

### GetDomainWarrantyRenameOk

`func (o *ListDomainsData) GetDomainWarrantyRenameOk() (*string, bool)`

GetDomainWarrantyRenameOk returns a tuple with the DomainWarrantyRename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainWarrantyRename

`func (o *ListDomainsData) SetDomainWarrantyRename(v string)`

SetDomainWarrantyRename sets DomainWarrantyRename field to given value.

### HasDomainWarrantyRename

`func (o *ListDomainsData) HasDomainWarrantyRename() bool`

HasDomainWarrantyRename returns a boolean if a field has been set.

### GetDomainEvaluationAvailable

`func (o *ListDomainsData) GetDomainEvaluationAvailable() bool`

GetDomainEvaluationAvailable returns the DomainEvaluationAvailable field if non-nil, zero value otherwise.

### GetDomainEvaluationAvailableOk

`func (o *ListDomainsData) GetDomainEvaluationAvailableOk() (*bool, bool)`

GetDomainEvaluationAvailableOk returns a tuple with the DomainEvaluationAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainEvaluationAvailable

`func (o *ListDomainsData) SetDomainEvaluationAvailable(v bool)`

SetDomainEvaluationAvailable sets DomainEvaluationAvailable field to given value.


### GetTotalCount

`func (o *ListDomainsData) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ListDomainsData) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ListDomainsData) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


