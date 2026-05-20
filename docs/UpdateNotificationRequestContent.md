# UpdateNotificationRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceId** | **string** | Service ID - must be sent as a string | 
**NotificationId** | **int32** | Notification ID to update | 
**Name** | Pointer to **string** | Name/title for the notification | [optional] 
**Status** | Pointer to [**NotificationStatus**](NotificationStatus.md) |  | [optional] 
**NotificationInterval** | Pointer to **int32** | Notification interval in minutes (must be &gt; 0) | [optional] 
**DataTimeframe** | Pointer to **int32** | Data timeframe in minutes (must be &gt; 0) | [optional] 
**ExceedAll** | Pointer to **bool** | Whether all thresholds must be exceeded (true) or any one (false) | [optional] 
**EmailAddress** | Pointer to **[]string** | Email addresses for notifications (comma-separated string or array) | [optional] 
**CpuUsage** | Pointer to **int32** | CPU usage threshold percentage (0-100) | [optional] 
**MemoryUsage** | Pointer to **int32** | Memory usage threshold percentage (0-100) | [optional] 
**NetworkTraffic** | Pointer to **int32** | Network traffic threshold percentage (0-100) | [optional] 
**DiskRead** | Pointer to **int32** | Disk read threshold percentage (0-100) | [optional] 
**DiskWrite** | Pointer to **int32** | Disk write threshold percentage (0-100) | [optional] 

## Methods

### NewUpdateNotificationRequestContent

`func NewUpdateNotificationRequestContent(serviceId string, notificationId int32, ) *UpdateNotificationRequestContent`

NewUpdateNotificationRequestContent instantiates a new UpdateNotificationRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNotificationRequestContentWithDefaults

`func NewUpdateNotificationRequestContentWithDefaults() *UpdateNotificationRequestContent`

NewUpdateNotificationRequestContentWithDefaults instantiates a new UpdateNotificationRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceId

`func (o *UpdateNotificationRequestContent) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *UpdateNotificationRequestContent) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *UpdateNotificationRequestContent) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetNotificationId

`func (o *UpdateNotificationRequestContent) GetNotificationId() int32`

GetNotificationId returns the NotificationId field if non-nil, zero value otherwise.

### GetNotificationIdOk

`func (o *UpdateNotificationRequestContent) GetNotificationIdOk() (*int32, bool)`

GetNotificationIdOk returns a tuple with the NotificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationId

`func (o *UpdateNotificationRequestContent) SetNotificationId(v int32)`

SetNotificationId sets NotificationId field to given value.


### GetName

`func (o *UpdateNotificationRequestContent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateNotificationRequestContent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateNotificationRequestContent) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateNotificationRequestContent) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *UpdateNotificationRequestContent) GetStatus() NotificationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateNotificationRequestContent) GetStatusOk() (*NotificationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateNotificationRequestContent) SetStatus(v NotificationStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UpdateNotificationRequestContent) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetNotificationInterval

`func (o *UpdateNotificationRequestContent) GetNotificationInterval() int32`

GetNotificationInterval returns the NotificationInterval field if non-nil, zero value otherwise.

### GetNotificationIntervalOk

`func (o *UpdateNotificationRequestContent) GetNotificationIntervalOk() (*int32, bool)`

GetNotificationIntervalOk returns a tuple with the NotificationInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationInterval

`func (o *UpdateNotificationRequestContent) SetNotificationInterval(v int32)`

SetNotificationInterval sets NotificationInterval field to given value.

### HasNotificationInterval

`func (o *UpdateNotificationRequestContent) HasNotificationInterval() bool`

HasNotificationInterval returns a boolean if a field has been set.

### GetDataTimeframe

`func (o *UpdateNotificationRequestContent) GetDataTimeframe() int32`

GetDataTimeframe returns the DataTimeframe field if non-nil, zero value otherwise.

### GetDataTimeframeOk

`func (o *UpdateNotificationRequestContent) GetDataTimeframeOk() (*int32, bool)`

GetDataTimeframeOk returns a tuple with the DataTimeframe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTimeframe

`func (o *UpdateNotificationRequestContent) SetDataTimeframe(v int32)`

SetDataTimeframe sets DataTimeframe field to given value.

### HasDataTimeframe

`func (o *UpdateNotificationRequestContent) HasDataTimeframe() bool`

HasDataTimeframe returns a boolean if a field has been set.

### GetExceedAll

`func (o *UpdateNotificationRequestContent) GetExceedAll() bool`

GetExceedAll returns the ExceedAll field if non-nil, zero value otherwise.

### GetExceedAllOk

`func (o *UpdateNotificationRequestContent) GetExceedAllOk() (*bool, bool)`

GetExceedAllOk returns a tuple with the ExceedAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceedAll

`func (o *UpdateNotificationRequestContent) SetExceedAll(v bool)`

SetExceedAll sets ExceedAll field to given value.

### HasExceedAll

`func (o *UpdateNotificationRequestContent) HasExceedAll() bool`

HasExceedAll returns a boolean if a field has been set.

### GetEmailAddress

`func (o *UpdateNotificationRequestContent) GetEmailAddress() []string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *UpdateNotificationRequestContent) GetEmailAddressOk() (*[]string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *UpdateNotificationRequestContent) SetEmailAddress(v []string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *UpdateNotificationRequestContent) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### GetCpuUsage

`func (o *UpdateNotificationRequestContent) GetCpuUsage() int32`

GetCpuUsage returns the CpuUsage field if non-nil, zero value otherwise.

### GetCpuUsageOk

`func (o *UpdateNotificationRequestContent) GetCpuUsageOk() (*int32, bool)`

GetCpuUsageOk returns a tuple with the CpuUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUsage

`func (o *UpdateNotificationRequestContent) SetCpuUsage(v int32)`

SetCpuUsage sets CpuUsage field to given value.

### HasCpuUsage

`func (o *UpdateNotificationRequestContent) HasCpuUsage() bool`

HasCpuUsage returns a boolean if a field has been set.

### GetMemoryUsage

`func (o *UpdateNotificationRequestContent) GetMemoryUsage() int32`

GetMemoryUsage returns the MemoryUsage field if non-nil, zero value otherwise.

### GetMemoryUsageOk

`func (o *UpdateNotificationRequestContent) GetMemoryUsageOk() (*int32, bool)`

GetMemoryUsageOk returns a tuple with the MemoryUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryUsage

`func (o *UpdateNotificationRequestContent) SetMemoryUsage(v int32)`

SetMemoryUsage sets MemoryUsage field to given value.

### HasMemoryUsage

`func (o *UpdateNotificationRequestContent) HasMemoryUsage() bool`

HasMemoryUsage returns a boolean if a field has been set.

### GetNetworkTraffic

`func (o *UpdateNotificationRequestContent) GetNetworkTraffic() int32`

GetNetworkTraffic returns the NetworkTraffic field if non-nil, zero value otherwise.

### GetNetworkTrafficOk

`func (o *UpdateNotificationRequestContent) GetNetworkTrafficOk() (*int32, bool)`

GetNetworkTrafficOk returns a tuple with the NetworkTraffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkTraffic

`func (o *UpdateNotificationRequestContent) SetNetworkTraffic(v int32)`

SetNetworkTraffic sets NetworkTraffic field to given value.

### HasNetworkTraffic

`func (o *UpdateNotificationRequestContent) HasNetworkTraffic() bool`

HasNetworkTraffic returns a boolean if a field has been set.

### GetDiskRead

`func (o *UpdateNotificationRequestContent) GetDiskRead() int32`

GetDiskRead returns the DiskRead field if non-nil, zero value otherwise.

### GetDiskReadOk

`func (o *UpdateNotificationRequestContent) GetDiskReadOk() (*int32, bool)`

GetDiskReadOk returns a tuple with the DiskRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskRead

`func (o *UpdateNotificationRequestContent) SetDiskRead(v int32)`

SetDiskRead sets DiskRead field to given value.

### HasDiskRead

`func (o *UpdateNotificationRequestContent) HasDiskRead() bool`

HasDiskRead returns a boolean if a field has been set.

### GetDiskWrite

`func (o *UpdateNotificationRequestContent) GetDiskWrite() int32`

GetDiskWrite returns the DiskWrite field if non-nil, zero value otherwise.

### GetDiskWriteOk

`func (o *UpdateNotificationRequestContent) GetDiskWriteOk() (*int32, bool)`

GetDiskWriteOk returns a tuple with the DiskWrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskWrite

`func (o *UpdateNotificationRequestContent) SetDiskWrite(v int32)`

SetDiskWrite sets DiskWrite field to given value.

### HasDiskWrite

`func (o *UpdateNotificationRequestContent) HasDiskWrite() bool`

HasDiskWrite returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


