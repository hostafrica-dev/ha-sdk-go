# CreateNotificationRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceId** | **string** | Service ID - must be sent as a string | 
**Name** | **string** | Name/title for the notification | 
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

### NewCreateNotificationRequestContent

`func NewCreateNotificationRequestContent(serviceId string, name string, ) *CreateNotificationRequestContent`

NewCreateNotificationRequestContent instantiates a new CreateNotificationRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNotificationRequestContentWithDefaults

`func NewCreateNotificationRequestContentWithDefaults() *CreateNotificationRequestContent`

NewCreateNotificationRequestContentWithDefaults instantiates a new CreateNotificationRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceId

`func (o *CreateNotificationRequestContent) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *CreateNotificationRequestContent) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *CreateNotificationRequestContent) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetName

`func (o *CreateNotificationRequestContent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateNotificationRequestContent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateNotificationRequestContent) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *CreateNotificationRequestContent) GetStatus() NotificationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateNotificationRequestContent) GetStatusOk() (*NotificationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateNotificationRequestContent) SetStatus(v NotificationStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CreateNotificationRequestContent) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetNotificationInterval

`func (o *CreateNotificationRequestContent) GetNotificationInterval() int32`

GetNotificationInterval returns the NotificationInterval field if non-nil, zero value otherwise.

### GetNotificationIntervalOk

`func (o *CreateNotificationRequestContent) GetNotificationIntervalOk() (*int32, bool)`

GetNotificationIntervalOk returns a tuple with the NotificationInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationInterval

`func (o *CreateNotificationRequestContent) SetNotificationInterval(v int32)`

SetNotificationInterval sets NotificationInterval field to given value.

### HasNotificationInterval

`func (o *CreateNotificationRequestContent) HasNotificationInterval() bool`

HasNotificationInterval returns a boolean if a field has been set.

### GetDataTimeframe

`func (o *CreateNotificationRequestContent) GetDataTimeframe() int32`

GetDataTimeframe returns the DataTimeframe field if non-nil, zero value otherwise.

### GetDataTimeframeOk

`func (o *CreateNotificationRequestContent) GetDataTimeframeOk() (*int32, bool)`

GetDataTimeframeOk returns a tuple with the DataTimeframe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTimeframe

`func (o *CreateNotificationRequestContent) SetDataTimeframe(v int32)`

SetDataTimeframe sets DataTimeframe field to given value.

### HasDataTimeframe

`func (o *CreateNotificationRequestContent) HasDataTimeframe() bool`

HasDataTimeframe returns a boolean if a field has been set.

### GetExceedAll

`func (o *CreateNotificationRequestContent) GetExceedAll() bool`

GetExceedAll returns the ExceedAll field if non-nil, zero value otherwise.

### GetExceedAllOk

`func (o *CreateNotificationRequestContent) GetExceedAllOk() (*bool, bool)`

GetExceedAllOk returns a tuple with the ExceedAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceedAll

`func (o *CreateNotificationRequestContent) SetExceedAll(v bool)`

SetExceedAll sets ExceedAll field to given value.

### HasExceedAll

`func (o *CreateNotificationRequestContent) HasExceedAll() bool`

HasExceedAll returns a boolean if a field has been set.

### GetEmailAddress

`func (o *CreateNotificationRequestContent) GetEmailAddress() []string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *CreateNotificationRequestContent) GetEmailAddressOk() (*[]string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *CreateNotificationRequestContent) SetEmailAddress(v []string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *CreateNotificationRequestContent) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### GetCpuUsage

`func (o *CreateNotificationRequestContent) GetCpuUsage() int32`

GetCpuUsage returns the CpuUsage field if non-nil, zero value otherwise.

### GetCpuUsageOk

`func (o *CreateNotificationRequestContent) GetCpuUsageOk() (*int32, bool)`

GetCpuUsageOk returns a tuple with the CpuUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUsage

`func (o *CreateNotificationRequestContent) SetCpuUsage(v int32)`

SetCpuUsage sets CpuUsage field to given value.

### HasCpuUsage

`func (o *CreateNotificationRequestContent) HasCpuUsage() bool`

HasCpuUsage returns a boolean if a field has been set.

### GetMemoryUsage

`func (o *CreateNotificationRequestContent) GetMemoryUsage() int32`

GetMemoryUsage returns the MemoryUsage field if non-nil, zero value otherwise.

### GetMemoryUsageOk

`func (o *CreateNotificationRequestContent) GetMemoryUsageOk() (*int32, bool)`

GetMemoryUsageOk returns a tuple with the MemoryUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryUsage

`func (o *CreateNotificationRequestContent) SetMemoryUsage(v int32)`

SetMemoryUsage sets MemoryUsage field to given value.

### HasMemoryUsage

`func (o *CreateNotificationRequestContent) HasMemoryUsage() bool`

HasMemoryUsage returns a boolean if a field has been set.

### GetNetworkTraffic

`func (o *CreateNotificationRequestContent) GetNetworkTraffic() int32`

GetNetworkTraffic returns the NetworkTraffic field if non-nil, zero value otherwise.

### GetNetworkTrafficOk

`func (o *CreateNotificationRequestContent) GetNetworkTrafficOk() (*int32, bool)`

GetNetworkTrafficOk returns a tuple with the NetworkTraffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkTraffic

`func (o *CreateNotificationRequestContent) SetNetworkTraffic(v int32)`

SetNetworkTraffic sets NetworkTraffic field to given value.

### HasNetworkTraffic

`func (o *CreateNotificationRequestContent) HasNetworkTraffic() bool`

HasNetworkTraffic returns a boolean if a field has been set.

### GetDiskRead

`func (o *CreateNotificationRequestContent) GetDiskRead() int32`

GetDiskRead returns the DiskRead field if non-nil, zero value otherwise.

### GetDiskReadOk

`func (o *CreateNotificationRequestContent) GetDiskReadOk() (*int32, bool)`

GetDiskReadOk returns a tuple with the DiskRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskRead

`func (o *CreateNotificationRequestContent) SetDiskRead(v int32)`

SetDiskRead sets DiskRead field to given value.

### HasDiskRead

`func (o *CreateNotificationRequestContent) HasDiskRead() bool`

HasDiskRead returns a boolean if a field has been set.

### GetDiskWrite

`func (o *CreateNotificationRequestContent) GetDiskWrite() int32`

GetDiskWrite returns the DiskWrite field if non-nil, zero value otherwise.

### GetDiskWriteOk

`func (o *CreateNotificationRequestContent) GetDiskWriteOk() (*int32, bool)`

GetDiskWriteOk returns a tuple with the DiskWrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskWrite

`func (o *CreateNotificationRequestContent) SetDiskWrite(v int32)`

SetDiskWrite sets DiskWrite field to given value.

### HasDiskWrite

`func (o *CreateNotificationRequestContent) HasDiskWrite() bool`

HasDiskWrite returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


