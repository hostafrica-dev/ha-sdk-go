# ListDnsCreateCandidatesData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**Candidates** | [**[]DnsCreateCandidateInfo**](DnsCreateCandidateInfo.md) | Domains and services eligible for new DNS zones | 
**Total** | **int32** |  | 

## Methods

### NewListDnsCreateCandidatesData

`func NewListDnsCreateCandidatesData(message string, candidates []DnsCreateCandidateInfo, total int32, ) *ListDnsCreateCandidatesData`

NewListDnsCreateCandidatesData instantiates a new ListDnsCreateCandidatesData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListDnsCreateCandidatesDataWithDefaults

`func NewListDnsCreateCandidatesDataWithDefaults() *ListDnsCreateCandidatesData`

NewListDnsCreateCandidatesDataWithDefaults instantiates a new ListDnsCreateCandidatesData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ListDnsCreateCandidatesData) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ListDnsCreateCandidatesData) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ListDnsCreateCandidatesData) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetCandidates

`func (o *ListDnsCreateCandidatesData) GetCandidates() []DnsCreateCandidateInfo`

GetCandidates returns the Candidates field if non-nil, zero value otherwise.

### GetCandidatesOk

`func (o *ListDnsCreateCandidatesData) GetCandidatesOk() (*[]DnsCreateCandidateInfo, bool)`

GetCandidatesOk returns a tuple with the Candidates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCandidates

`func (o *ListDnsCreateCandidatesData) SetCandidates(v []DnsCreateCandidateInfo)`

SetCandidates sets Candidates field to given value.


### GetTotal

`func (o *ListDnsCreateCandidatesData) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ListDnsCreateCandidatesData) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ListDnsCreateCandidatesData) SetTotal(v int32)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


