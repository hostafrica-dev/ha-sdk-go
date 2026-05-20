# UpdateFirewallRuleRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceId** | **string** | Service ID - must be sent as a string | 
**Pos** | **int32** | Position/index of the rule to update | 
**Comment** | Pointer to **string** | Comment/description for the rule | [optional] 
**RuleAction** | Pointer to [**FirewallRuleAction**](FirewallRuleAction.md) |  | [optional] 

## Methods

### NewUpdateFirewallRuleRequestContent

`func NewUpdateFirewallRuleRequestContent(serviceId string, pos int32, ) *UpdateFirewallRuleRequestContent`

NewUpdateFirewallRuleRequestContent instantiates a new UpdateFirewallRuleRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateFirewallRuleRequestContentWithDefaults

`func NewUpdateFirewallRuleRequestContentWithDefaults() *UpdateFirewallRuleRequestContent`

NewUpdateFirewallRuleRequestContentWithDefaults instantiates a new UpdateFirewallRuleRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceId

`func (o *UpdateFirewallRuleRequestContent) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *UpdateFirewallRuleRequestContent) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *UpdateFirewallRuleRequestContent) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetPos

`func (o *UpdateFirewallRuleRequestContent) GetPos() int32`

GetPos returns the Pos field if non-nil, zero value otherwise.

### GetPosOk

`func (o *UpdateFirewallRuleRequestContent) GetPosOk() (*int32, bool)`

GetPosOk returns a tuple with the Pos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPos

`func (o *UpdateFirewallRuleRequestContent) SetPos(v int32)`

SetPos sets Pos field to given value.


### GetComment

`func (o *UpdateFirewallRuleRequestContent) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *UpdateFirewallRuleRequestContent) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *UpdateFirewallRuleRequestContent) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *UpdateFirewallRuleRequestContent) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetRuleAction

`func (o *UpdateFirewallRuleRequestContent) GetRuleAction() FirewallRuleAction`

GetRuleAction returns the RuleAction field if non-nil, zero value otherwise.

### GetRuleActionOk

`func (o *UpdateFirewallRuleRequestContent) GetRuleActionOk() (*FirewallRuleAction, bool)`

GetRuleActionOk returns a tuple with the RuleAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleAction

`func (o *UpdateFirewallRuleRequestContent) SetRuleAction(v FirewallRuleAction)`

SetRuleAction sets RuleAction field to given value.

### HasRuleAction

`func (o *UpdateFirewallRuleRequestContent) HasRuleAction() bool`

HasRuleAction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


