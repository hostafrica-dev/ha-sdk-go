# ListDnsZonesData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Status message indicating the result | 
**Zones** | [**[]DnsZoneInfo**](DnsZoneInfo.md) | DNS zones owned by the authenticated client | 
**Total** | **int32** | Total number of zones | 

## Methods

### NewListDnsZonesData

`func NewListDnsZonesData(message string, zones []DnsZoneInfo, total int32, ) *ListDnsZonesData`

NewListDnsZonesData instantiates a new ListDnsZonesData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListDnsZonesDataWithDefaults

`func NewListDnsZonesDataWithDefaults() *ListDnsZonesData`

NewListDnsZonesDataWithDefaults instantiates a new ListDnsZonesData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ListDnsZonesData) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ListDnsZonesData) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ListDnsZonesData) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetZones

`func (o *ListDnsZonesData) GetZones() []DnsZoneInfo`

GetZones returns the Zones field if non-nil, zero value otherwise.

### GetZonesOk

`func (o *ListDnsZonesData) GetZonesOk() (*[]DnsZoneInfo, bool)`

GetZonesOk returns a tuple with the Zones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZones

`func (o *ListDnsZonesData) SetZones(v []DnsZoneInfo)`

SetZones sets Zones field to given value.


### GetTotal

`func (o *ListDnsZonesData) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ListDnsZonesData) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ListDnsZonesData) SetTotal(v int32)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


