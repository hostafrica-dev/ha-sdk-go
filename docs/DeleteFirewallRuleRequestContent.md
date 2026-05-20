# DeleteFirewallRuleRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceId** | **string** | Service ID - must be sent as a string | 
**Pos** | **int32** | Position/index of the rule to delete | 

## Methods

### NewDeleteFirewallRuleRequestContent

`func NewDeleteFirewallRuleRequestContent(serviceId string, pos int32, ) *DeleteFirewallRuleRequestContent`

NewDeleteFirewallRuleRequestContent instantiates a new DeleteFirewallRuleRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteFirewallRuleRequestContentWithDefaults

`func NewDeleteFirewallRuleRequestContentWithDefaults() *DeleteFirewallRuleRequestContent`

NewDeleteFirewallRuleRequestContentWithDefaults instantiates a new DeleteFirewallRuleRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceId

`func (o *DeleteFirewallRuleRequestContent) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *DeleteFirewallRuleRequestContent) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *DeleteFirewallRuleRequestContent) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetPos

`func (o *DeleteFirewallRuleRequestContent) GetPos() int32`

GetPos returns the Pos field if non-nil, zero value otherwise.

### GetPosOk

`func (o *DeleteFirewallRuleRequestContent) GetPosOk() (*int32, bool)`

GetPosOk returns a tuple with the Pos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPos

`func (o *DeleteFirewallRuleRequestContent) SetPos(v int32)`

SetPos sets Pos field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


