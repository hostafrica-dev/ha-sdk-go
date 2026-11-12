# CreateDnsZoneResponseContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**OperationStatus**](OperationStatus.md) |  | 
**Data** | [**DnsZoneMutationData**](DnsZoneMutationData.md) |  | 

## Methods

### NewCreateDnsZoneResponseContent

`func NewCreateDnsZoneResponseContent(status OperationStatus, data DnsZoneMutationData, ) *CreateDnsZoneResponseContent`

NewCreateDnsZoneResponseContent instantiates a new CreateDnsZoneResponseContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDnsZoneResponseContentWithDefaults

`func NewCreateDnsZoneResponseContentWithDefaults() *CreateDnsZoneResponseContent`

NewCreateDnsZoneResponseContentWithDefaults instantiates a new CreateDnsZoneResponseContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *CreateDnsZoneResponseContent) GetStatus() OperationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateDnsZoneResponseContent) GetStatusOk() (*OperationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateDnsZoneResponseContent) SetStatus(v OperationStatus)`

SetStatus sets Status field to given value.


### GetData

`func (o *CreateDnsZoneResponseContent) GetData() DnsZoneMutationData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateDnsZoneResponseContent) GetDataOk() (*DnsZoneMutationData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateDnsZoneResponseContent) SetData(v DnsZoneMutationData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


