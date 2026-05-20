# ListVpsServicesResponseContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**OperationStatus**](OperationStatus.md) |  | 
**Data** | [**ListVpsServicesData**](ListVpsServicesData.md) |  | 

## Methods

### NewListVpsServicesResponseContent

`func NewListVpsServicesResponseContent(status OperationStatus, data ListVpsServicesData, ) *ListVpsServicesResponseContent`

NewListVpsServicesResponseContent instantiates a new ListVpsServicesResponseContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListVpsServicesResponseContentWithDefaults

`func NewListVpsServicesResponseContentWithDefaults() *ListVpsServicesResponseContent`

NewListVpsServicesResponseContentWithDefaults instantiates a new ListVpsServicesResponseContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ListVpsServicesResponseContent) GetStatus() OperationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListVpsServicesResponseContent) GetStatusOk() (*OperationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListVpsServicesResponseContent) SetStatus(v OperationStatus)`

SetStatus sets Status field to given value.


### GetData

`func (o *ListVpsServicesResponseContent) GetData() ListVpsServicesData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListVpsServicesResponseContent) GetDataOk() (*ListVpsServicesData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListVpsServicesResponseContent) SetData(v ListVpsServicesData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


