# VpsCancelResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Status message indicating the result of the operation | 
**ServiceId** | **int32** | ID of the service being cancelled | 
**CancellationType** | **string** | The cancellation type that was applied - &#39;Immediate&#39; or &#39;End of Billing Period&#39; | 
**Status** | **string** | Current status of the cancellation request | 

## Methods

### NewVpsCancelResponseData

`func NewVpsCancelResponseData(message string, serviceId int32, cancellationType string, status string, ) *VpsCancelResponseData`

NewVpsCancelResponseData instantiates a new VpsCancelResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpsCancelResponseDataWithDefaults

`func NewVpsCancelResponseDataWithDefaults() *VpsCancelResponseData`

NewVpsCancelResponseDataWithDefaults instantiates a new VpsCancelResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *VpsCancelResponseData) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *VpsCancelResponseData) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *VpsCancelResponseData) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetServiceId

`func (o *VpsCancelResponseData) GetServiceId() int32`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *VpsCancelResponseData) GetServiceIdOk() (*int32, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *VpsCancelResponseData) SetServiceId(v int32)`

SetServiceId sets ServiceId field to given value.


### GetCancellationType

`func (o *VpsCancelResponseData) GetCancellationType() string`

GetCancellationType returns the CancellationType field if non-nil, zero value otherwise.

### GetCancellationTypeOk

`func (o *VpsCancelResponseData) GetCancellationTypeOk() (*string, bool)`

GetCancellationTypeOk returns a tuple with the CancellationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancellationType

`func (o *VpsCancelResponseData) SetCancellationType(v string)`

SetCancellationType sets CancellationType field to given value.


### GetStatus

`func (o *VpsCancelResponseData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VpsCancelResponseData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VpsCancelResponseData) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


