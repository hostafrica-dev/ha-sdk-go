# ListDnsCreateCandidatesResponseContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**OperationStatus**](OperationStatus.md) |  | 
**Data** | [**ListDnsCreateCandidatesData**](ListDnsCreateCandidatesData.md) |  | 

## Methods

### NewListDnsCreateCandidatesResponseContent

`func NewListDnsCreateCandidatesResponseContent(status OperationStatus, data ListDnsCreateCandidatesData, ) *ListDnsCreateCandidatesResponseContent`

NewListDnsCreateCandidatesResponseContent instantiates a new ListDnsCreateCandidatesResponseContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListDnsCreateCandidatesResponseContentWithDefaults

`func NewListDnsCreateCandidatesResponseContentWithDefaults() *ListDnsCreateCandidatesResponseContent`

NewListDnsCreateCandidatesResponseContentWithDefaults instantiates a new ListDnsCreateCandidatesResponseContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ListDnsCreateCandidatesResponseContent) GetStatus() OperationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListDnsCreateCandidatesResponseContent) GetStatusOk() (*OperationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListDnsCreateCandidatesResponseContent) SetStatus(v OperationStatus)`

SetStatus sets Status field to given value.


### GetData

`func (o *ListDnsCreateCandidatesResponseContent) GetData() ListDnsCreateCandidatesData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListDnsCreateCandidatesResponseContent) GetDataOk() (*ListDnsCreateCandidatesData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListDnsCreateCandidatesResponseContent) SetData(v ListDnsCreateCandidatesData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


