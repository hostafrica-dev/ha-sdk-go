# ListBackupsResponseContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**OperationStatus**](OperationStatus.md) |  | 
**Data** | [**ServiceBackupsResponseData**](ServiceBackupsResponseData.md) |  | 

## Methods

### NewListBackupsResponseContent

`func NewListBackupsResponseContent(status OperationStatus, data ServiceBackupsResponseData, ) *ListBackupsResponseContent`

NewListBackupsResponseContent instantiates a new ListBackupsResponseContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListBackupsResponseContentWithDefaults

`func NewListBackupsResponseContentWithDefaults() *ListBackupsResponseContent`

NewListBackupsResponseContentWithDefaults instantiates a new ListBackupsResponseContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ListBackupsResponseContent) GetStatus() OperationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListBackupsResponseContent) GetStatusOk() (*OperationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListBackupsResponseContent) SetStatus(v OperationStatus)`

SetStatus sets Status field to given value.


### GetData

`func (o *ListBackupsResponseContent) GetData() ServiceBackupsResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListBackupsResponseContent) GetDataOk() (*ServiceBackupsResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListBackupsResponseContent) SetData(v ServiceBackupsResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


