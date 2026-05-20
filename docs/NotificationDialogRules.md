# NotificationDialogRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequiredOnCreate** | **[]string** | Fields required when creating a new notification | 
**ConditionKeys** | **[]string** | Valid condition key values | 
**ThresholdRange** | **[]int32** | Threshold value range [min, max] | 
**PeriodMin** | **int32** | Minimum period value in minutes | 
**TimeframeMin** | **int32** | Minimum timeframe value in minutes | 

## Methods

### NewNotificationDialogRules

`func NewNotificationDialogRules(requiredOnCreate []string, conditionKeys []string, thresholdRange []int32, periodMin int32, timeframeMin int32, ) *NotificationDialogRules`

NewNotificationDialogRules instantiates a new NotificationDialogRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationDialogRulesWithDefaults

`func NewNotificationDialogRulesWithDefaults() *NotificationDialogRules`

NewNotificationDialogRulesWithDefaults instantiates a new NotificationDialogRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequiredOnCreate

`func (o *NotificationDialogRules) GetRequiredOnCreate() []string`

GetRequiredOnCreate returns the RequiredOnCreate field if non-nil, zero value otherwise.

### GetRequiredOnCreateOk

`func (o *NotificationDialogRules) GetRequiredOnCreateOk() (*[]string, bool)`

GetRequiredOnCreateOk returns a tuple with the RequiredOnCreate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredOnCreate

`func (o *NotificationDialogRules) SetRequiredOnCreate(v []string)`

SetRequiredOnCreate sets RequiredOnCreate field to given value.


### GetConditionKeys

`func (o *NotificationDialogRules) GetConditionKeys() []string`

GetConditionKeys returns the ConditionKeys field if non-nil, zero value otherwise.

### GetConditionKeysOk

`func (o *NotificationDialogRules) GetConditionKeysOk() (*[]string, bool)`

GetConditionKeysOk returns a tuple with the ConditionKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionKeys

`func (o *NotificationDialogRules) SetConditionKeys(v []string)`

SetConditionKeys sets ConditionKeys field to given value.


### GetThresholdRange

`func (o *NotificationDialogRules) GetThresholdRange() []int32`

GetThresholdRange returns the ThresholdRange field if non-nil, zero value otherwise.

### GetThresholdRangeOk

`func (o *NotificationDialogRules) GetThresholdRangeOk() (*[]int32, bool)`

GetThresholdRangeOk returns a tuple with the ThresholdRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdRange

`func (o *NotificationDialogRules) SetThresholdRange(v []int32)`

SetThresholdRange sets ThresholdRange field to given value.


### GetPeriodMin

`func (o *NotificationDialogRules) GetPeriodMin() int32`

GetPeriodMin returns the PeriodMin field if non-nil, zero value otherwise.

### GetPeriodMinOk

`func (o *NotificationDialogRules) GetPeriodMinOk() (*int32, bool)`

GetPeriodMinOk returns a tuple with the PeriodMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodMin

`func (o *NotificationDialogRules) SetPeriodMin(v int32)`

SetPeriodMin sets PeriodMin field to given value.


### GetTimeframeMin

`func (o *NotificationDialogRules) GetTimeframeMin() int32`

GetTimeframeMin returns the TimeframeMin field if non-nil, zero value otherwise.

### GetTimeframeMinOk

`func (o *NotificationDialogRules) GetTimeframeMinOk() (*int32, bool)`

GetTimeframeMinOk returns a tuple with the TimeframeMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeframeMin

`func (o *NotificationDialogRules) SetTimeframeMin(v int32)`

SetTimeframeMin sets TimeframeMin field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


