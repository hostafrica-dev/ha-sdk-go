# ListFirewallRulesResponseContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**OperationStatus**](OperationStatus.md) |  | 
**Data** | [**FirewallListResponseData**](FirewallListResponseData.md) |  | 

## Methods

### NewListFirewallRulesResponseContent

`func NewListFirewallRulesResponseContent(status OperationStatus, data FirewallListResponseData, ) *ListFirewallRulesResponseContent`

NewListFirewallRulesResponseContent instantiates a new ListFirewallRulesResponseContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListFirewallRulesResponseContentWithDefaults

`func NewListFirewallRulesResponseContentWithDefaults() *ListFirewallRulesResponseContent`

NewListFirewallRulesResponseContentWithDefaults instantiates a new ListFirewallRulesResponseContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ListFirewallRulesResponseContent) GetStatus() OperationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListFirewallRulesResponseContent) GetStatusOk() (*OperationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListFirewallRulesResponseContent) SetStatus(v OperationStatus)`

SetStatus sets Status field to given value.


### GetData

`func (o *ListFirewallRulesResponseContent) GetData() FirewallListResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListFirewallRulesResponseContent) GetDataOk() (*FirewallListResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListFirewallRulesResponseContent) SetData(v FirewallListResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


