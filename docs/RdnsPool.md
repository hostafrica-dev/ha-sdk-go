# RdnsPool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pool** | **string** | Network address of the pool | 
**Mask** | **int32** | CIDR prefix length | 

## Methods

### NewRdnsPool

`func NewRdnsPool(pool string, mask int32, ) *RdnsPool`

NewRdnsPool instantiates a new RdnsPool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRdnsPoolWithDefaults

`func NewRdnsPoolWithDefaults() *RdnsPool`

NewRdnsPoolWithDefaults instantiates a new RdnsPool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPool

`func (o *RdnsPool) GetPool() string`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *RdnsPool) GetPoolOk() (*string, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *RdnsPool) SetPool(v string)`

SetPool sets Pool field to given value.


### GetMask

`func (o *RdnsPool) GetMask() int32`

GetMask returns the Mask field if non-nil, zero value otherwise.

### GetMaskOk

`func (o *RdnsPool) GetMaskOk() (*int32, bool)`

GetMaskOk returns a tuple with the Mask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMask

`func (o *RdnsPool) SetMask(v int32)`

SetMask sets Mask field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


