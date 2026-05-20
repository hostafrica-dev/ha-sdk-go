# MoveFirewallRuleRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceId** | **string** | Service ID - must be sent as a string | 
**Pos** | **int32** | Current position/index of the rule to move | 
**TargetPos** | Pointer to **int32** | Target position/index for the rule (mutually exclusive with direction) | [optional] 
**Direction** | Pointer to [**FirewallMoveDirection**](FirewallMoveDirection.md) |  | [optional] 

## Methods

### NewMoveFirewallRuleRequestContent

`func NewMoveFirewallRuleRequestContent(serviceId string, pos int32, ) *MoveFirewallRuleRequestContent`

NewMoveFirewallRuleRequestContent instantiates a new MoveFirewallRuleRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoveFirewallRuleRequestContentWithDefaults

`func NewMoveFirewallRuleRequestContentWithDefaults() *MoveFirewallRuleRequestContent`

NewMoveFirewallRuleRequestContentWithDefaults instantiates a new MoveFirewallRuleRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceId

`func (o *MoveFirewallRuleRequestContent) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *MoveFirewallRuleRequestContent) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *MoveFirewallRuleRequestContent) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetPos

`func (o *MoveFirewallRuleRequestContent) GetPos() int32`

GetPos returns the Pos field if non-nil, zero value otherwise.

### GetPosOk

`func (o *MoveFirewallRuleRequestContent) GetPosOk() (*int32, bool)`

GetPosOk returns a tuple with the Pos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPos

`func (o *MoveFirewallRuleRequestContent) SetPos(v int32)`

SetPos sets Pos field to given value.


### GetTargetPos

`func (o *MoveFirewallRuleRequestContent) GetTargetPos() int32`

GetTargetPos returns the TargetPos field if non-nil, zero value otherwise.

### GetTargetPosOk

`func (o *MoveFirewallRuleRequestContent) GetTargetPosOk() (*int32, bool)`

GetTargetPosOk returns a tuple with the TargetPos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPos

`func (o *MoveFirewallRuleRequestContent) SetTargetPos(v int32)`

SetTargetPos sets TargetPos field to given value.

### HasTargetPos

`func (o *MoveFirewallRuleRequestContent) HasTargetPos() bool`

HasTargetPos returns a boolean if a field has been set.

### GetDirection

`func (o *MoveFirewallRuleRequestContent) GetDirection() FirewallMoveDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *MoveFirewallRuleRequestContent) GetDirectionOk() (*FirewallMoveDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *MoveFirewallRuleRequestContent) SetDirection(v FirewallMoveDirection)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *MoveFirewallRuleRequestContent) HasDirection() bool`

HasDirection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


