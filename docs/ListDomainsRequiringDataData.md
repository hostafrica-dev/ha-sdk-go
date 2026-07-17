# ListDomainsRequiringDataData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Status message indicating the result | 
**Domains** | [**[]DomainRequiringDataInfo**](DomainRequiringDataInfo.md) | Pending domains needing additional registrar or contact data | 
**TotalCount** | **int32** | Count of domains returned | 
**RequiresAdditionalFields** | **bool** | True when any domain has missing required fields | 

## Methods

### NewListDomainsRequiringDataData

`func NewListDomainsRequiringDataData(message string, domains []DomainRequiringDataInfo, totalCount int32, requiresAdditionalFields bool, ) *ListDomainsRequiringDataData`

NewListDomainsRequiringDataData instantiates a new ListDomainsRequiringDataData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListDomainsRequiringDataDataWithDefaults

`func NewListDomainsRequiringDataDataWithDefaults() *ListDomainsRequiringDataData`

NewListDomainsRequiringDataDataWithDefaults instantiates a new ListDomainsRequiringDataData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ListDomainsRequiringDataData) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ListDomainsRequiringDataData) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ListDomainsRequiringDataData) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetDomains

`func (o *ListDomainsRequiringDataData) GetDomains() []DomainRequiringDataInfo`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *ListDomainsRequiringDataData) GetDomainsOk() (*[]DomainRequiringDataInfo, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *ListDomainsRequiringDataData) SetDomains(v []DomainRequiringDataInfo)`

SetDomains sets Domains field to given value.


### GetTotalCount

`func (o *ListDomainsRequiringDataData) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ListDomainsRequiringDataData) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ListDomainsRequiringDataData) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetRequiresAdditionalFields

`func (o *ListDomainsRequiringDataData) GetRequiresAdditionalFields() bool`

GetRequiresAdditionalFields returns the RequiresAdditionalFields field if non-nil, zero value otherwise.

### GetRequiresAdditionalFieldsOk

`func (o *ListDomainsRequiringDataData) GetRequiresAdditionalFieldsOk() (*bool, bool)`

GetRequiresAdditionalFieldsOk returns a tuple with the RequiresAdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresAdditionalFields

`func (o *ListDomainsRequiringDataData) SetRequiresAdditionalFields(v bool)`

SetRequiresAdditionalFields sets RequiresAdditionalFields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


