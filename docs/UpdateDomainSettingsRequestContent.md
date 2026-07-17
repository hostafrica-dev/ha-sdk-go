# UpdateDomainSettingsRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainId** | **string** | Domain service id - must be sent as a string | 
**Setting** | [**DomainSettingKey**](DomainSettingKey.md) |  | 
**Value** | **bool** | New boolean value | 
**Gateway** | Pointer to **string** | Payment gateway slug when enabling a paid addon (e.g. stripe) | [optional] 

## Methods

### NewUpdateDomainSettingsRequestContent

`func NewUpdateDomainSettingsRequestContent(domainId string, setting DomainSettingKey, value bool, ) *UpdateDomainSettingsRequestContent`

NewUpdateDomainSettingsRequestContent instantiates a new UpdateDomainSettingsRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDomainSettingsRequestContentWithDefaults

`func NewUpdateDomainSettingsRequestContentWithDefaults() *UpdateDomainSettingsRequestContent`

NewUpdateDomainSettingsRequestContentWithDefaults instantiates a new UpdateDomainSettingsRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomainId

`func (o *UpdateDomainSettingsRequestContent) GetDomainId() string`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *UpdateDomainSettingsRequestContent) GetDomainIdOk() (*string, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *UpdateDomainSettingsRequestContent) SetDomainId(v string)`

SetDomainId sets DomainId field to given value.


### GetSetting

`func (o *UpdateDomainSettingsRequestContent) GetSetting() DomainSettingKey`

GetSetting returns the Setting field if non-nil, zero value otherwise.

### GetSettingOk

`func (o *UpdateDomainSettingsRequestContent) GetSettingOk() (*DomainSettingKey, bool)`

GetSettingOk returns a tuple with the Setting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetting

`func (o *UpdateDomainSettingsRequestContent) SetSetting(v DomainSettingKey)`

SetSetting sets Setting field to given value.


### GetValue

`func (o *UpdateDomainSettingsRequestContent) GetValue() bool`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UpdateDomainSettingsRequestContent) GetValueOk() (*bool, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UpdateDomainSettingsRequestContent) SetValue(v bool)`

SetValue sets Value field to given value.


### GetGateway

`func (o *UpdateDomainSettingsRequestContent) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *UpdateDomainSettingsRequestContent) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *UpdateDomainSettingsRequestContent) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *UpdateDomainSettingsRequestContent) HasGateway() bool`

HasGateway returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


