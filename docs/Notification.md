# Notification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Notification ID | 
**HostingId** | **int32** | Hosting/service ID | 
**Name** | **string** | Notification name/title | 
**ConditionKey** | **string** | Condition key (cpu_usage, memory_usage, network_traffic, disk_read, disk_write) | 
**Threshold** | **int32** | Threshold value | 
**Period** | **int32** | Period in minutes | 
**Timeframe** | **int32** | Timeframe in minutes | 
**Enabled** | Pointer to **bool** | Whether the notification is enabled | [optional] 
**CreatedAt** | Pointer to **string** | Creation timestamp | [optional] 
**UpdatedAt** | Pointer to **string** | Last updated timestamp | [optional] 

## Methods

### NewNotification

`func NewNotification(id int32, hostingId int32, name string, conditionKey string, threshold int32, period int32, timeframe int32, ) *Notification`

NewNotification instantiates a new Notification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationWithDefaults

`func NewNotificationWithDefaults() *Notification`

NewNotificationWithDefaults instantiates a new Notification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Notification) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Notification) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Notification) SetId(v int32)`

SetId sets Id field to given value.


### GetHostingId

`func (o *Notification) GetHostingId() int32`

GetHostingId returns the HostingId field if non-nil, zero value otherwise.

### GetHostingIdOk

`func (o *Notification) GetHostingIdOk() (*int32, bool)`

GetHostingIdOk returns a tuple with the HostingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostingId

`func (o *Notification) SetHostingId(v int32)`

SetHostingId sets HostingId field to given value.


### GetName

`func (o *Notification) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Notification) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Notification) SetName(v string)`

SetName sets Name field to given value.


### GetConditionKey

`func (o *Notification) GetConditionKey() string`

GetConditionKey returns the ConditionKey field if non-nil, zero value otherwise.

### GetConditionKeyOk

`func (o *Notification) GetConditionKeyOk() (*string, bool)`

GetConditionKeyOk returns a tuple with the ConditionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionKey

`func (o *Notification) SetConditionKey(v string)`

SetConditionKey sets ConditionKey field to given value.


### GetThreshold

`func (o *Notification) GetThreshold() int32`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *Notification) GetThresholdOk() (*int32, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *Notification) SetThreshold(v int32)`

SetThreshold sets Threshold field to given value.


### GetPeriod

`func (o *Notification) GetPeriod() int32`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *Notification) GetPeriodOk() (*int32, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *Notification) SetPeriod(v int32)`

SetPeriod sets Period field to given value.


### GetTimeframe

`func (o *Notification) GetTimeframe() int32`

GetTimeframe returns the Timeframe field if non-nil, zero value otherwise.

### GetTimeframeOk

`func (o *Notification) GetTimeframeOk() (*int32, bool)`

GetTimeframeOk returns a tuple with the Timeframe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeframe

`func (o *Notification) SetTimeframe(v int32)`

SetTimeframe sets Timeframe field to given value.


### GetEnabled

`func (o *Notification) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Notification) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Notification) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *Notification) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Notification) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Notification) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Notification) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Notification) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Notification) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Notification) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Notification) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Notification) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


