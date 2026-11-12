# RegisterPrivateNameserverResponseContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**OperationStatus**](OperationStatus.md) |  | 
**Data** | [**PrivateNameserverMutationData**](PrivateNameserverMutationData.md) |  | 

## Methods

### NewRegisterPrivateNameserverResponseContent

`func NewRegisterPrivateNameserverResponseContent(status OperationStatus, data PrivateNameserverMutationData, ) *RegisterPrivateNameserverResponseContent`

NewRegisterPrivateNameserverResponseContent instantiates a new RegisterPrivateNameserverResponseContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterPrivateNameserverResponseContentWithDefaults

`func NewRegisterPrivateNameserverResponseContentWithDefaults() *RegisterPrivateNameserverResponseContent`

NewRegisterPrivateNameserverResponseContentWithDefaults instantiates a new RegisterPrivateNameserverResponseContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *RegisterPrivateNameserverResponseContent) GetStatus() OperationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RegisterPrivateNameserverResponseContent) GetStatusOk() (*OperationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RegisterPrivateNameserverResponseContent) SetStatus(v OperationStatus)`

SetStatus sets Status field to given value.


### GetData

`func (o *RegisterPrivateNameserverResponseContent) GetData() PrivateNameserverMutationData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RegisterPrivateNameserverResponseContent) GetDataOk() (*PrivateNameserverMutationData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RegisterPrivateNameserverResponseContent) SetData(v PrivateNameserverMutationData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


