# TriggerReinstallRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceId** | **string** | Service ID - must be sent as a string | 
**Template** | **string** | Template identifier for the OS to reinstall | 

## Methods

### NewTriggerReinstallRequestContent

`func NewTriggerReinstallRequestContent(serviceId string, template string, ) *TriggerReinstallRequestContent`

NewTriggerReinstallRequestContent instantiates a new TriggerReinstallRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTriggerReinstallRequestContentWithDefaults

`func NewTriggerReinstallRequestContentWithDefaults() *TriggerReinstallRequestContent`

NewTriggerReinstallRequestContentWithDefaults instantiates a new TriggerReinstallRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceId

`func (o *TriggerReinstallRequestContent) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *TriggerReinstallRequestContent) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *TriggerReinstallRequestContent) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetTemplate

`func (o *TriggerReinstallRequestContent) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *TriggerReinstallRequestContent) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *TriggerReinstallRequestContent) SetTemplate(v string)`

SetTemplate sets Template field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


