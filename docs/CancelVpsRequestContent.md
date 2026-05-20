# CancelVpsRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceId** | **string** | Service ID - must be sent as a string | 
**Type** | Pointer to [**VpsCancelType**](VpsCancelType.md) |  | [optional] 

## Methods

### NewCancelVpsRequestContent

`func NewCancelVpsRequestContent(serviceId string, ) *CancelVpsRequestContent`

NewCancelVpsRequestContent instantiates a new CancelVpsRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCancelVpsRequestContentWithDefaults

`func NewCancelVpsRequestContentWithDefaults() *CancelVpsRequestContent`

NewCancelVpsRequestContentWithDefaults instantiates a new CancelVpsRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceId

`func (o *CancelVpsRequestContent) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *CancelVpsRequestContent) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *CancelVpsRequestContent) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetType

`func (o *CancelVpsRequestContent) GetType() VpsCancelType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CancelVpsRequestContent) GetTypeOk() (*VpsCancelType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CancelVpsRequestContent) SetType(v VpsCancelType)`

SetType sets Type field to given value.

### HasType

`func (o *CancelVpsRequestContent) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


