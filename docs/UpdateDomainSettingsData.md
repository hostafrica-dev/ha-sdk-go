# UpdateDomainSettingsData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Status message indicating the result | 
**DomainId** | **string** | Domain service id as returned by upstream | 
**Setting** | [**DomainSettingKey**](DomainSettingKey.md) |  | 
**Value** | **bool** | Boolean value requested | 
**RequiresPayment** | **bool** | True when payment must be completed before a paid addon is enabled | 
**InvoiceId** | Pointer to **int32** | Invoice id when requires_payment is true | [optional] 
**Amount** | Pointer to **float64** | Invoice amount in client currency when requires_payment is true | [optional] 

## Methods

### NewUpdateDomainSettingsData

`func NewUpdateDomainSettingsData(message string, domainId string, setting DomainSettingKey, value bool, requiresPayment bool, ) *UpdateDomainSettingsData`

NewUpdateDomainSettingsData instantiates a new UpdateDomainSettingsData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDomainSettingsDataWithDefaults

`func NewUpdateDomainSettingsDataWithDefaults() *UpdateDomainSettingsData`

NewUpdateDomainSettingsDataWithDefaults instantiates a new UpdateDomainSettingsData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *UpdateDomainSettingsData) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *UpdateDomainSettingsData) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *UpdateDomainSettingsData) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetDomainId

`func (o *UpdateDomainSettingsData) GetDomainId() string`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *UpdateDomainSettingsData) GetDomainIdOk() (*string, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *UpdateDomainSettingsData) SetDomainId(v string)`

SetDomainId sets DomainId field to given value.


### GetSetting

`func (o *UpdateDomainSettingsData) GetSetting() DomainSettingKey`

GetSetting returns the Setting field if non-nil, zero value otherwise.

### GetSettingOk

`func (o *UpdateDomainSettingsData) GetSettingOk() (*DomainSettingKey, bool)`

GetSettingOk returns a tuple with the Setting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetting

`func (o *UpdateDomainSettingsData) SetSetting(v DomainSettingKey)`

SetSetting sets Setting field to given value.


### GetValue

`func (o *UpdateDomainSettingsData) GetValue() bool`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UpdateDomainSettingsData) GetValueOk() (*bool, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UpdateDomainSettingsData) SetValue(v bool)`

SetValue sets Value field to given value.


### GetRequiresPayment

`func (o *UpdateDomainSettingsData) GetRequiresPayment() bool`

GetRequiresPayment returns the RequiresPayment field if non-nil, zero value otherwise.

### GetRequiresPaymentOk

`func (o *UpdateDomainSettingsData) GetRequiresPaymentOk() (*bool, bool)`

GetRequiresPaymentOk returns a tuple with the RequiresPayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresPayment

`func (o *UpdateDomainSettingsData) SetRequiresPayment(v bool)`

SetRequiresPayment sets RequiresPayment field to given value.


### GetInvoiceId

`func (o *UpdateDomainSettingsData) GetInvoiceId() int32`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *UpdateDomainSettingsData) GetInvoiceIdOk() (*int32, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *UpdateDomainSettingsData) SetInvoiceId(v int32)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *UpdateDomainSettingsData) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### GetAmount

`func (o *UpdateDomainSettingsData) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *UpdateDomainSettingsData) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *UpdateDomainSettingsData) SetAmount(v float64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *UpdateDomainSettingsData) HasAmount() bool`

HasAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


