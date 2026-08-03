# GetDnsZoneDetailsResponseContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**OperationStatus**](OperationStatus.md) |  | 
**Data** | [**GetDnsZoneDetailsData**](GetDnsZoneDetailsData.md) |  | 

## Methods

### NewGetDnsZoneDetailsResponseContent

`func NewGetDnsZoneDetailsResponseContent(status OperationStatus, data GetDnsZoneDetailsData, ) *GetDnsZoneDetailsResponseContent`

NewGetDnsZoneDetailsResponseContent instantiates a new GetDnsZoneDetailsResponseContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDnsZoneDetailsResponseContentWithDefaults

`func NewGetDnsZoneDetailsResponseContentWithDefaults() *GetDnsZoneDetailsResponseContent`

NewGetDnsZoneDetailsResponseContentWithDefaults instantiates a new GetDnsZoneDetailsResponseContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetDnsZoneDetailsResponseContent) GetStatus() OperationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetDnsZoneDetailsResponseContent) GetStatusOk() (*OperationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetDnsZoneDetailsResponseContent) SetStatus(v OperationStatus)`

SetStatus sets Status field to given value.


### GetData

`func (o *GetDnsZoneDetailsResponseContent) GetData() GetDnsZoneDetailsData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetDnsZoneDetailsResponseContent) GetDataOk() (*GetDnsZoneDetailsData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetDnsZoneDetailsResponseContent) SetData(v GetDnsZoneDetailsData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


