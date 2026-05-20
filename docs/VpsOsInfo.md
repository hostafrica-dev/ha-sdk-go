# VpsOsInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | OS name or identifier | [optional] 
**Version** | Pointer to **string** | OS version | [optional] 

## Methods

### NewVpsOsInfo

`func NewVpsOsInfo() *VpsOsInfo`

NewVpsOsInfo instantiates a new VpsOsInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpsOsInfoWithDefaults

`func NewVpsOsInfoWithDefaults() *VpsOsInfo`

NewVpsOsInfoWithDefaults instantiates a new VpsOsInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VpsOsInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VpsOsInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VpsOsInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VpsOsInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVersion

`func (o *VpsOsInfo) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *VpsOsInfo) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *VpsOsInfo) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *VpsOsInfo) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


