# UpdateFirewallRuleResponseContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**OperationStatus**](OperationStatus.md) |  | 
**Data** | [**VpsSimpleActionResponseData**](VpsSimpleActionResponseData.md) |  | 

## Methods

### NewUpdateFirewallRuleResponseContent

`func NewUpdateFirewallRuleResponseContent(status OperationStatus, data VpsSimpleActionResponseData, ) *UpdateFirewallRuleResponseContent`

NewUpdateFirewallRuleResponseContent instantiates a new UpdateFirewallRuleResponseContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateFirewallRuleResponseContentWithDefaults

`func NewUpdateFirewallRuleResponseContentWithDefaults() *UpdateFirewallRuleResponseContent`

NewUpdateFirewallRuleResponseContentWithDefaults instantiates a new UpdateFirewallRuleResponseContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *UpdateFirewallRuleResponseContent) GetStatus() OperationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateFirewallRuleResponseContent) GetStatusOk() (*OperationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateFirewallRuleResponseContent) SetStatus(v OperationStatus)`

SetStatus sets Status field to given value.


### GetData

`func (o *UpdateFirewallRuleResponseContent) GetData() VpsSimpleActionResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UpdateFirewallRuleResponseContent) GetDataOk() (*VpsSimpleActionResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UpdateFirewallRuleResponseContent) SetData(v VpsSimpleActionResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


