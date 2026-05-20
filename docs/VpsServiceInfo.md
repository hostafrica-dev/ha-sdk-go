# VpsServiceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Service ID (integer) | 
**Product** | **string** | Product name | 
**Hostname** | **string** | Hostname of the VPS | 
**Status** | **string** | Service status | 
**Ips** | **[]string** | List of IP addresses assigned to this VPS | 
**CreatedAt** | **string** | Service creation date (YYYY-MM-DD format) | 

## Methods

### NewVpsServiceInfo

`func NewVpsServiceInfo(id int32, product string, hostname string, status string, ips []string, createdAt string, ) *VpsServiceInfo`

NewVpsServiceInfo instantiates a new VpsServiceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpsServiceInfoWithDefaults

`func NewVpsServiceInfoWithDefaults() *VpsServiceInfo`

NewVpsServiceInfoWithDefaults instantiates a new VpsServiceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VpsServiceInfo) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VpsServiceInfo) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VpsServiceInfo) SetId(v int32)`

SetId sets Id field to given value.


### GetProduct

`func (o *VpsServiceInfo) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *VpsServiceInfo) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *VpsServiceInfo) SetProduct(v string)`

SetProduct sets Product field to given value.


### GetHostname

`func (o *VpsServiceInfo) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *VpsServiceInfo) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *VpsServiceInfo) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetStatus

`func (o *VpsServiceInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VpsServiceInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VpsServiceInfo) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetIps

`func (o *VpsServiceInfo) GetIps() []string`

GetIps returns the Ips field if non-nil, zero value otherwise.

### GetIpsOk

`func (o *VpsServiceInfo) GetIpsOk() (*[]string, bool)`

GetIpsOk returns a tuple with the Ips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIps

`func (o *VpsServiceInfo) SetIps(v []string)`

SetIps sets Ips field to given value.


### GetCreatedAt

`func (o *VpsServiceInfo) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VpsServiceInfo) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VpsServiceInfo) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


