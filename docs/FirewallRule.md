# FirewallRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Digest** | **string** | SHA1 digest of the rule | 
**Pos** | **int32** | Position/priority of the rule | 
**Type** | **string** | Rule type (in, out) | 
**Action** | **string** | Action to perform (ACCEPT, REJECT, DROP) | 
**Iface** | **string** | Network interface name | 
**Enable** | **int32** | Whether the rule is enabled (1&#x3D;enabled, 0&#x3D;disabled) | 
**Comment** | Pointer to **string** | Comment/description for the rule | [optional] 
**Source** | Pointer to **string** | Source address/network | [optional] 
**Dest** | Pointer to **string** | Destination address/network | [optional] 
**Proto** | Pointer to **string** | Protocol (tcp, udp, icmp, etc) | [optional] 
**Dport** | Pointer to **string** | Destination port | [optional] 
**Sport** | Pointer to **string** | Source port | [optional] 
**Macro** | Pointer to **string** | Firewall macro name | [optional] 
**Log** | Pointer to **string** | Log level | [optional] 

## Methods

### NewFirewallRule

`func NewFirewallRule(digest string, pos int32, type_ string, action string, iface string, enable int32, ) *FirewallRule`

NewFirewallRule instantiates a new FirewallRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallRuleWithDefaults

`func NewFirewallRuleWithDefaults() *FirewallRule`

NewFirewallRuleWithDefaults instantiates a new FirewallRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDigest

`func (o *FirewallRule) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *FirewallRule) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *FirewallRule) SetDigest(v string)`

SetDigest sets Digest field to given value.


### GetPos

`func (o *FirewallRule) GetPos() int32`

GetPos returns the Pos field if non-nil, zero value otherwise.

### GetPosOk

`func (o *FirewallRule) GetPosOk() (*int32, bool)`

GetPosOk returns a tuple with the Pos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPos

`func (o *FirewallRule) SetPos(v int32)`

SetPos sets Pos field to given value.


### GetType

`func (o *FirewallRule) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FirewallRule) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FirewallRule) SetType(v string)`

SetType sets Type field to given value.


### GetAction

`func (o *FirewallRule) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *FirewallRule) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *FirewallRule) SetAction(v string)`

SetAction sets Action field to given value.


### GetIface

`func (o *FirewallRule) GetIface() string`

GetIface returns the Iface field if non-nil, zero value otherwise.

### GetIfaceOk

`func (o *FirewallRule) GetIfaceOk() (*string, bool)`

GetIfaceOk returns a tuple with the Iface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIface

`func (o *FirewallRule) SetIface(v string)`

SetIface sets Iface field to given value.


### GetEnable

`func (o *FirewallRule) GetEnable() int32`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *FirewallRule) GetEnableOk() (*int32, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *FirewallRule) SetEnable(v int32)`

SetEnable sets Enable field to given value.


### GetComment

`func (o *FirewallRule) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *FirewallRule) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *FirewallRule) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *FirewallRule) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetSource

`func (o *FirewallRule) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *FirewallRule) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *FirewallRule) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *FirewallRule) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetDest

`func (o *FirewallRule) GetDest() string`

GetDest returns the Dest field if non-nil, zero value otherwise.

### GetDestOk

`func (o *FirewallRule) GetDestOk() (*string, bool)`

GetDestOk returns a tuple with the Dest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDest

`func (o *FirewallRule) SetDest(v string)`

SetDest sets Dest field to given value.

### HasDest

`func (o *FirewallRule) HasDest() bool`

HasDest returns a boolean if a field has been set.

### GetProto

`func (o *FirewallRule) GetProto() string`

GetProto returns the Proto field if non-nil, zero value otherwise.

### GetProtoOk

`func (o *FirewallRule) GetProtoOk() (*string, bool)`

GetProtoOk returns a tuple with the Proto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProto

`func (o *FirewallRule) SetProto(v string)`

SetProto sets Proto field to given value.

### HasProto

`func (o *FirewallRule) HasProto() bool`

HasProto returns a boolean if a field has been set.

### GetDport

`func (o *FirewallRule) GetDport() string`

GetDport returns the Dport field if non-nil, zero value otherwise.

### GetDportOk

`func (o *FirewallRule) GetDportOk() (*string, bool)`

GetDportOk returns a tuple with the Dport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDport

`func (o *FirewallRule) SetDport(v string)`

SetDport sets Dport field to given value.

### HasDport

`func (o *FirewallRule) HasDport() bool`

HasDport returns a boolean if a field has been set.

### GetSport

`func (o *FirewallRule) GetSport() string`

GetSport returns the Sport field if non-nil, zero value otherwise.

### GetSportOk

`func (o *FirewallRule) GetSportOk() (*string, bool)`

GetSportOk returns a tuple with the Sport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSport

`func (o *FirewallRule) SetSport(v string)`

SetSport sets Sport field to given value.

### HasSport

`func (o *FirewallRule) HasSport() bool`

HasSport returns a boolean if a field has been set.

### GetMacro

`func (o *FirewallRule) GetMacro() string`

GetMacro returns the Macro field if non-nil, zero value otherwise.

### GetMacroOk

`func (o *FirewallRule) GetMacroOk() (*string, bool)`

GetMacroOk returns a tuple with the Macro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacro

`func (o *FirewallRule) SetMacro(v string)`

SetMacro sets Macro field to given value.

### HasMacro

`func (o *FirewallRule) HasMacro() bool`

HasMacro returns a boolean if a field has been set.

### GetLog

`func (o *FirewallRule) GetLog() string`

GetLog returns the Log field if non-nil, zero value otherwise.

### GetLogOk

`func (o *FirewallRule) GetLogOk() (*string, bool)`

GetLogOk returns a tuple with the Log field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLog

`func (o *FirewallRule) SetLog(v string)`

SetLog sets Log field to given value.

### HasLog

`func (o *FirewallRule) HasLog() bool`

HasLog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


