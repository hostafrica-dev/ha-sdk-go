# GetVpsDetailsResponseContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**OperationStatus**](OperationStatus.md) |  | 
**Data** | [**VpsDetailsResponse**](VpsDetailsResponse.md) |  | 

## Methods

### NewGetVpsDetailsResponseContent

`func NewGetVpsDetailsResponseContent(status OperationStatus, data VpsDetailsResponse, ) *GetVpsDetailsResponseContent`

NewGetVpsDetailsResponseContent instantiates a new GetVpsDetailsResponseContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetVpsDetailsResponseContentWithDefaults

`func NewGetVpsDetailsResponseContentWithDefaults() *GetVpsDetailsResponseContent`

NewGetVpsDetailsResponseContentWithDefaults instantiates a new GetVpsDetailsResponseContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetVpsDetailsResponseContent) GetStatus() OperationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetVpsDetailsResponseContent) GetStatusOk() (*OperationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetVpsDetailsResponseContent) SetStatus(v OperationStatus)`

SetStatus sets Status field to given value.


### GetData

`func (o *GetVpsDetailsResponseContent) GetData() VpsDetailsResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetVpsDetailsResponseContent) GetDataOk() (*VpsDetailsResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetVpsDetailsResponseContent) SetData(v VpsDetailsResponse)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


