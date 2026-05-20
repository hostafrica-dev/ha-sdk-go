# ShutdownVpsRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceId** | **string** | Service ID - must be sent as a string | 

## Methods

### NewShutdownVpsRequestContent

`func NewShutdownVpsRequestContent(serviceId string, ) *ShutdownVpsRequestContent`

NewShutdownVpsRequestContent instantiates a new ShutdownVpsRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShutdownVpsRequestContentWithDefaults

`func NewShutdownVpsRequestContentWithDefaults() *ShutdownVpsRequestContent`

NewShutdownVpsRequestContentWithDefaults instantiates a new ShutdownVpsRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceId

`func (o *ShutdownVpsRequestContent) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ShutdownVpsRequestContent) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ShutdownVpsRequestContent) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


