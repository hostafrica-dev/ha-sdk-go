# ListRdnsResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Records** | [**[]RdnsRecord**](RdnsRecord.md) | All PTR records owned by the client | 
**PtrCount** | **int32** | Number of PTR records | 
**PtrLimit** | **int32** | Lowest finite per-package PTR limit (-1 &#x3D; unlimited / not configured) | 
**CustomIpMode** | **bool** | Any IP is allowed in the PTR add form | 
**SubnetCustomIpMode** | **bool** | Free-text IP allowed inside an offered subnet pool | 
**ServiceOnlyIps** | **bool** | UI hint: only show IPs tied to the related service | 
**AvailableItems** | [**[]RdnsAvailableItem**](RdnsAvailableItem.md) | Services the client can manage PTRs for | 

## Methods

### NewListRdnsResponseData

`func NewListRdnsResponseData(records []RdnsRecord, ptrCount int32, ptrLimit int32, customIpMode bool, subnetCustomIpMode bool, serviceOnlyIps bool, availableItems []RdnsAvailableItem, ) *ListRdnsResponseData`

NewListRdnsResponseData instantiates a new ListRdnsResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListRdnsResponseDataWithDefaults

`func NewListRdnsResponseDataWithDefaults() *ListRdnsResponseData`

NewListRdnsResponseDataWithDefaults instantiates a new ListRdnsResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecords

`func (o *ListRdnsResponseData) GetRecords() []RdnsRecord`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *ListRdnsResponseData) GetRecordsOk() (*[]RdnsRecord, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *ListRdnsResponseData) SetRecords(v []RdnsRecord)`

SetRecords sets Records field to given value.


### GetPtrCount

`func (o *ListRdnsResponseData) GetPtrCount() int32`

GetPtrCount returns the PtrCount field if non-nil, zero value otherwise.

### GetPtrCountOk

`func (o *ListRdnsResponseData) GetPtrCountOk() (*int32, bool)`

GetPtrCountOk returns a tuple with the PtrCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPtrCount

`func (o *ListRdnsResponseData) SetPtrCount(v int32)`

SetPtrCount sets PtrCount field to given value.


### GetPtrLimit

`func (o *ListRdnsResponseData) GetPtrLimit() int32`

GetPtrLimit returns the PtrLimit field if non-nil, zero value otherwise.

### GetPtrLimitOk

`func (o *ListRdnsResponseData) GetPtrLimitOk() (*int32, bool)`

GetPtrLimitOk returns a tuple with the PtrLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPtrLimit

`func (o *ListRdnsResponseData) SetPtrLimit(v int32)`

SetPtrLimit sets PtrLimit field to given value.


### GetCustomIpMode

`func (o *ListRdnsResponseData) GetCustomIpMode() bool`

GetCustomIpMode returns the CustomIpMode field if non-nil, zero value otherwise.

### GetCustomIpModeOk

`func (o *ListRdnsResponseData) GetCustomIpModeOk() (*bool, bool)`

GetCustomIpModeOk returns a tuple with the CustomIpMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomIpMode

`func (o *ListRdnsResponseData) SetCustomIpMode(v bool)`

SetCustomIpMode sets CustomIpMode field to given value.


### GetSubnetCustomIpMode

`func (o *ListRdnsResponseData) GetSubnetCustomIpMode() bool`

GetSubnetCustomIpMode returns the SubnetCustomIpMode field if non-nil, zero value otherwise.

### GetSubnetCustomIpModeOk

`func (o *ListRdnsResponseData) GetSubnetCustomIpModeOk() (*bool, bool)`

GetSubnetCustomIpModeOk returns a tuple with the SubnetCustomIpMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetCustomIpMode

`func (o *ListRdnsResponseData) SetSubnetCustomIpMode(v bool)`

SetSubnetCustomIpMode sets SubnetCustomIpMode field to given value.


### GetServiceOnlyIps

`func (o *ListRdnsResponseData) GetServiceOnlyIps() bool`

GetServiceOnlyIps returns the ServiceOnlyIps field if non-nil, zero value otherwise.

### GetServiceOnlyIpsOk

`func (o *ListRdnsResponseData) GetServiceOnlyIpsOk() (*bool, bool)`

GetServiceOnlyIpsOk returns a tuple with the ServiceOnlyIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceOnlyIps

`func (o *ListRdnsResponseData) SetServiceOnlyIps(v bool)`

SetServiceOnlyIps sets ServiceOnlyIps field to given value.


### GetAvailableItems

`func (o *ListRdnsResponseData) GetAvailableItems() []RdnsAvailableItem`

GetAvailableItems returns the AvailableItems field if non-nil, zero value otherwise.

### GetAvailableItemsOk

`func (o *ListRdnsResponseData) GetAvailableItemsOk() (*[]RdnsAvailableItem, bool)`

GetAvailableItemsOk returns a tuple with the AvailableItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableItems

`func (o *ListRdnsResponseData) SetAvailableItems(v []RdnsAvailableItem)`

SetAvailableItems sets AvailableItems field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


