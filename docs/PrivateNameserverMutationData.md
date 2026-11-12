# PrivateNameserverMutationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Status message indicating the result | 
**Nameserver** | Pointer to **string** | Child nameserver hostname | [optional] 
**IpAddress** | Pointer to **string** | Glue IP address when returned (register) | [optional] 

## Methods

### NewPrivateNameserverMutationData

`func NewPrivateNameserverMutationData(message string, ) *PrivateNameserverMutationData`

NewPrivateNameserverMutationData instantiates a new PrivateNameserverMutationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateNameserverMutationDataWithDefaults

`func NewPrivateNameserverMutationDataWithDefaults() *PrivateNameserverMutationData`

NewPrivateNameserverMutationDataWithDefaults instantiates a new PrivateNameserverMutationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *PrivateNameserverMutationData) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PrivateNameserverMutationData) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PrivateNameserverMutationData) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetNameserver

`func (o *PrivateNameserverMutationData) GetNameserver() string`

GetNameserver returns the Nameserver field if non-nil, zero value otherwise.

### GetNameserverOk

`func (o *PrivateNameserverMutationData) GetNameserverOk() (*string, bool)`

GetNameserverOk returns a tuple with the Nameserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameserver

`func (o *PrivateNameserverMutationData) SetNameserver(v string)`

SetNameserver sets Nameserver field to given value.

### HasNameserver

`func (o *PrivateNameserverMutationData) HasNameserver() bool`

HasNameserver returns a boolean if a field has been set.

### GetIpAddress

`func (o *PrivateNameserverMutationData) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *PrivateNameserverMutationData) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *PrivateNameserverMutationData) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *PrivateNameserverMutationData) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


