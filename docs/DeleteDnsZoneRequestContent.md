# DeleteDnsZoneRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ZoneId** | **string** | DNS Manager zone id from list-dns-zones. Required by upstream deletezone. | 
**DomainId** | Pointer to **string** | Optional domain id from list-dns-zones when available. | [optional] 
**ServiceId** | Pointer to **int32** | Optional hosting service id (list-dns-zones hosting_id) when available. | [optional] 

## Methods

### NewDeleteDnsZoneRequestContent

`func NewDeleteDnsZoneRequestContent(zoneId string, ) *DeleteDnsZoneRequestContent`

NewDeleteDnsZoneRequestContent instantiates a new DeleteDnsZoneRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteDnsZoneRequestContentWithDefaults

`func NewDeleteDnsZoneRequestContentWithDefaults() *DeleteDnsZoneRequestContent`

NewDeleteDnsZoneRequestContentWithDefaults instantiates a new DeleteDnsZoneRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetZoneId

`func (o *DeleteDnsZoneRequestContent) GetZoneId() string`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *DeleteDnsZoneRequestContent) GetZoneIdOk() (*string, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *DeleteDnsZoneRequestContent) SetZoneId(v string)`

SetZoneId sets ZoneId field to given value.


### GetDomainId

`func (o *DeleteDnsZoneRequestContent) GetDomainId() string`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *DeleteDnsZoneRequestContent) GetDomainIdOk() (*string, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *DeleteDnsZoneRequestContent) SetDomainId(v string)`

SetDomainId sets DomainId field to given value.

### HasDomainId

`func (o *DeleteDnsZoneRequestContent) HasDomainId() bool`

HasDomainId returns a boolean if a field has been set.

### GetServiceId

`func (o *DeleteDnsZoneRequestContent) GetServiceId() int32`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *DeleteDnsZoneRequestContent) GetServiceIdOk() (*int32, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *DeleteDnsZoneRequestContent) SetServiceId(v int32)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *DeleteDnsZoneRequestContent) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


