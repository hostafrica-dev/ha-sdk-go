# FirewallMoveResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Status message indicating the result | 
**FromPos** | **int32** | Original position of the rule | 
**TargetPos** | **int32** | New position of the rule | 

## Methods

### NewFirewallMoveResponseData

`func NewFirewallMoveResponseData(message string, fromPos int32, targetPos int32, ) *FirewallMoveResponseData`

NewFirewallMoveResponseData instantiates a new FirewallMoveResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallMoveResponseDataWithDefaults

`func NewFirewallMoveResponseDataWithDefaults() *FirewallMoveResponseData`

NewFirewallMoveResponseDataWithDefaults instantiates a new FirewallMoveResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *FirewallMoveResponseData) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *FirewallMoveResponseData) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *FirewallMoveResponseData) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetFromPos

`func (o *FirewallMoveResponseData) GetFromPos() int32`

GetFromPos returns the FromPos field if non-nil, zero value otherwise.

### GetFromPosOk

`func (o *FirewallMoveResponseData) GetFromPosOk() (*int32, bool)`

GetFromPosOk returns a tuple with the FromPos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromPos

`func (o *FirewallMoveResponseData) SetFromPos(v int32)`

SetFromPos sets FromPos field to given value.


### GetTargetPos

`func (o *FirewallMoveResponseData) GetTargetPos() int32`

GetTargetPos returns the TargetPos field if non-nil, zero value otherwise.

### GetTargetPosOk

`func (o *FirewallMoveResponseData) GetTargetPosOk() (*int32, bool)`

GetTargetPosOk returns a tuple with the TargetPos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPos

`func (o *FirewallMoveResponseData) SetTargetPos(v int32)`

SetTargetPos sets TargetPos field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


