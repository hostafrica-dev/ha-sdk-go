# GetDomainRedirectDetailsResponseContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**OperationStatus**](OperationStatus.md) |  | 
**Data** | [**GetDomainRedirectDetailsData**](GetDomainRedirectDetailsData.md) |  | 

## Methods

### NewGetDomainRedirectDetailsResponseContent

`func NewGetDomainRedirectDetailsResponseContent(status OperationStatus, data GetDomainRedirectDetailsData, ) *GetDomainRedirectDetailsResponseContent`

NewGetDomainRedirectDetailsResponseContent instantiates a new GetDomainRedirectDetailsResponseContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDomainRedirectDetailsResponseContentWithDefaults

`func NewGetDomainRedirectDetailsResponseContentWithDefaults() *GetDomainRedirectDetailsResponseContent`

NewGetDomainRedirectDetailsResponseContentWithDefaults instantiates a new GetDomainRedirectDetailsResponseContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetDomainRedirectDetailsResponseContent) GetStatus() OperationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetDomainRedirectDetailsResponseContent) GetStatusOk() (*OperationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetDomainRedirectDetailsResponseContent) SetStatus(v OperationStatus)`

SetStatus sets Status field to given value.


### GetData

`func (o *GetDomainRedirectDetailsResponseContent) GetData() GetDomainRedirectDetailsData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetDomainRedirectDetailsResponseContent) GetDataOk() (*GetDomainRedirectDetailsData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetDomainRedirectDetailsResponseContent) SetData(v GetDomainRedirectDetailsData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


