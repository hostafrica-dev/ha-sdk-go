# CreateFirewallRuleRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceId** | **string** | Service ID - must be sent as a string | 
**Type** | [**FirewallRuleType**](FirewallRuleType.md) |  | 
**RuleAction** | [**FirewallRuleAction**](FirewallRuleAction.md) |  | 
**Enable** | Pointer to **int32** | Whether the rule is enabled (1&#x3D;enabled, 0&#x3D;disabled) | [optional] 
**Comment** | Pointer to **string** | Comment/description for the rule | [optional] 
**Source** | Pointer to **string** | Source address/network | [optional] 
**Dest** | Pointer to **string** | Destination address/network | [optional] 
**Proto** | Pointer to **string** | Protocol (tcp, udp, icmp, etc). Use &#39;0&#39; or omit for none | [optional] 
**Dport** | Pointer to **string** | Destination port | [optional] 
**Sport** | Pointer to **string** | Source port | [optional] 
**Macro** | Pointer to **string** | Firewall macro name | [optional] 
**Iface** | Pointer to **string** | Network interface name | [optional] 
**Log** | Pointer to **string** | Log level | [optional] 
**Pos** | Pointer to **int32** | Position/priority of the rule | [optional] 

## Methods

### NewCreateFirewallRuleRequestContent

`func NewCreateFirewallRuleRequestContent(serviceId string, type_ FirewallRuleType, ruleAction FirewallRuleAction, ) *CreateFirewallRuleRequestContent`

NewCreateFirewallRuleRequestContent instantiates a new CreateFirewallRuleRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFirewallRuleRequestContentWithDefaults

`func NewCreateFirewallRuleRequestContentWithDefaults() *CreateFirewallRuleRequestContent`

NewCreateFirewallRuleRequestContentWithDefaults instantiates a new CreateFirewallRuleRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceId

`func (o *CreateFirewallRuleRequestContent) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *CreateFirewallRuleRequestContent) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *CreateFirewallRuleRequestContent) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetType

`func (o *CreateFirewallRuleRequestContent) GetType() FirewallRuleType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateFirewallRuleRequestContent) GetTypeOk() (*FirewallRuleType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateFirewallRuleRequestContent) SetType(v FirewallRuleType)`

SetType sets Type field to given value.


### GetRuleAction

`func (o *CreateFirewallRuleRequestContent) GetRuleAction() FirewallRuleAction`

GetRuleAction returns the RuleAction field if non-nil, zero value otherwise.

### GetRuleActionOk

`func (o *CreateFirewallRuleRequestContent) GetRuleActionOk() (*FirewallRuleAction, bool)`

GetRuleActionOk returns a tuple with the RuleAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleAction

`func (o *CreateFirewallRuleRequestContent) SetRuleAction(v FirewallRuleAction)`

SetRuleAction sets RuleAction field to given value.


### GetEnable

`func (o *CreateFirewallRuleRequestContent) GetEnable() int32`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *CreateFirewallRuleRequestContent) GetEnableOk() (*int32, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *CreateFirewallRuleRequestContent) SetEnable(v int32)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *CreateFirewallRuleRequestContent) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetComment

`func (o *CreateFirewallRuleRequestContent) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *CreateFirewallRuleRequestContent) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *CreateFirewallRuleRequestContent) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *CreateFirewallRuleRequestContent) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetSource

`func (o *CreateFirewallRuleRequestContent) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CreateFirewallRuleRequestContent) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CreateFirewallRuleRequestContent) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *CreateFirewallRuleRequestContent) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetDest

`func (o *CreateFirewallRuleRequestContent) GetDest() string`

GetDest returns the Dest field if non-nil, zero value otherwise.

### GetDestOk

`func (o *CreateFirewallRuleRequestContent) GetDestOk() (*string, bool)`

GetDestOk returns a tuple with the Dest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDest

`func (o *CreateFirewallRuleRequestContent) SetDest(v string)`

SetDest sets Dest field to given value.

### HasDest

`func (o *CreateFirewallRuleRequestContent) HasDest() bool`

HasDest returns a boolean if a field has been set.

### GetProto

`func (o *CreateFirewallRuleRequestContent) GetProto() string`

GetProto returns the Proto field if non-nil, zero value otherwise.

### GetProtoOk

`func (o *CreateFirewallRuleRequestContent) GetProtoOk() (*string, bool)`

GetProtoOk returns a tuple with the Proto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProto

`func (o *CreateFirewallRuleRequestContent) SetProto(v string)`

SetProto sets Proto field to given value.

### HasProto

`func (o *CreateFirewallRuleRequestContent) HasProto() bool`

HasProto returns a boolean if a field has been set.

### GetDport

`func (o *CreateFirewallRuleRequestContent) GetDport() string`

GetDport returns the Dport field if non-nil, zero value otherwise.

### GetDportOk

`func (o *CreateFirewallRuleRequestContent) GetDportOk() (*string, bool)`

GetDportOk returns a tuple with the Dport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDport

`func (o *CreateFirewallRuleRequestContent) SetDport(v string)`

SetDport sets Dport field to given value.

### HasDport

`func (o *CreateFirewallRuleRequestContent) HasDport() bool`

HasDport returns a boolean if a field has been set.

### GetSport

`func (o *CreateFirewallRuleRequestContent) GetSport() string`

GetSport returns the Sport field if non-nil, zero value otherwise.

### GetSportOk

`func (o *CreateFirewallRuleRequestContent) GetSportOk() (*string, bool)`

GetSportOk returns a tuple with the Sport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSport

`func (o *CreateFirewallRuleRequestContent) SetSport(v string)`

SetSport sets Sport field to given value.

### HasSport

`func (o *CreateFirewallRuleRequestContent) HasSport() bool`

HasSport returns a boolean if a field has been set.

### GetMacro

`func (o *CreateFirewallRuleRequestContent) GetMacro() string`

GetMacro returns the Macro field if non-nil, zero value otherwise.

### GetMacroOk

`func (o *CreateFirewallRuleRequestContent) GetMacroOk() (*string, bool)`

GetMacroOk returns a tuple with the Macro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacro

`func (o *CreateFirewallRuleRequestContent) SetMacro(v string)`

SetMacro sets Macro field to given value.

### HasMacro

`func (o *CreateFirewallRuleRequestContent) HasMacro() bool`

HasMacro returns a boolean if a field has been set.

### GetIface

`func (o *CreateFirewallRuleRequestContent) GetIface() string`

GetIface returns the Iface field if non-nil, zero value otherwise.

### GetIfaceOk

`func (o *CreateFirewallRuleRequestContent) GetIfaceOk() (*string, bool)`

GetIfaceOk returns a tuple with the Iface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIface

`func (o *CreateFirewallRuleRequestContent) SetIface(v string)`

SetIface sets Iface field to given value.

### HasIface

`func (o *CreateFirewallRuleRequestContent) HasIface() bool`

HasIface returns a boolean if a field has been set.

### GetLog

`func (o *CreateFirewallRuleRequestContent) GetLog() string`

GetLog returns the Log field if non-nil, zero value otherwise.

### GetLogOk

`func (o *CreateFirewallRuleRequestContent) GetLogOk() (*string, bool)`

GetLogOk returns a tuple with the Log field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLog

`func (o *CreateFirewallRuleRequestContent) SetLog(v string)`

SetLog sets Log field to given value.

### HasLog

`func (o *CreateFirewallRuleRequestContent) HasLog() bool`

HasLog returns a boolean if a field has been set.

### GetPos

`func (o *CreateFirewallRuleRequestContent) GetPos() int32`

GetPos returns the Pos field if non-nil, zero value otherwise.

### GetPosOk

`func (o *CreateFirewallRuleRequestContent) GetPosOk() (*int32, bool)`

GetPosOk returns a tuple with the Pos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPos

`func (o *CreateFirewallRuleRequestContent) SetPos(v int32)`

SetPos sets Pos field to given value.

### HasPos

`func (o *CreateFirewallRuleRequestContent) HasPos() bool`

HasPos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


