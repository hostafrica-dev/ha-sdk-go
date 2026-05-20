# FirewallListResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Status message | 
**Rules** | [**[]FirewallRule**](FirewallRule.md) | List of firewall rules | 
**AvailableInterfaces** | [**[]FirewallOption**](FirewallOption.md) | Available network interfaces | 
**AvailableMacros** | [**[]FirewallOption**](FirewallOption.md) | Available firewall macros | 
**AvailableProtocols** | [**[]FirewallOption**](FirewallOption.md) | Available network protocols | 

## Methods

### NewFirewallListResponseData

`func NewFirewallListResponseData(message string, rules []FirewallRule, availableInterfaces []FirewallOption, availableMacros []FirewallOption, availableProtocols []FirewallOption, ) *FirewallListResponseData`

NewFirewallListResponseData instantiates a new FirewallListResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallListResponseDataWithDefaults

`func NewFirewallListResponseDataWithDefaults() *FirewallListResponseData`

NewFirewallListResponseDataWithDefaults instantiates a new FirewallListResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *FirewallListResponseData) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *FirewallListResponseData) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *FirewallListResponseData) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetRules

`func (o *FirewallListResponseData) GetRules() []FirewallRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *FirewallListResponseData) GetRulesOk() (*[]FirewallRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *FirewallListResponseData) SetRules(v []FirewallRule)`

SetRules sets Rules field to given value.


### GetAvailableInterfaces

`func (o *FirewallListResponseData) GetAvailableInterfaces() []FirewallOption`

GetAvailableInterfaces returns the AvailableInterfaces field if non-nil, zero value otherwise.

### GetAvailableInterfacesOk

`func (o *FirewallListResponseData) GetAvailableInterfacesOk() (*[]FirewallOption, bool)`

GetAvailableInterfacesOk returns a tuple with the AvailableInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableInterfaces

`func (o *FirewallListResponseData) SetAvailableInterfaces(v []FirewallOption)`

SetAvailableInterfaces sets AvailableInterfaces field to given value.


### GetAvailableMacros

`func (o *FirewallListResponseData) GetAvailableMacros() []FirewallOption`

GetAvailableMacros returns the AvailableMacros field if non-nil, zero value otherwise.

### GetAvailableMacrosOk

`func (o *FirewallListResponseData) GetAvailableMacrosOk() (*[]FirewallOption, bool)`

GetAvailableMacrosOk returns a tuple with the AvailableMacros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableMacros

`func (o *FirewallListResponseData) SetAvailableMacros(v []FirewallOption)`

SetAvailableMacros sets AvailableMacros field to given value.


### GetAvailableProtocols

`func (o *FirewallListResponseData) GetAvailableProtocols() []FirewallOption`

GetAvailableProtocols returns the AvailableProtocols field if non-nil, zero value otherwise.

### GetAvailableProtocolsOk

`func (o *FirewallListResponseData) GetAvailableProtocolsOk() (*[]FirewallOption, bool)`

GetAvailableProtocolsOk returns a tuple with the AvailableProtocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableProtocols

`func (o *FirewallListResponseData) SetAvailableProtocols(v []FirewallOption)`

SetAvailableProtocols sets AvailableProtocols field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


