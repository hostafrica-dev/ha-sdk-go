# ListVpsServicesData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Services** | [**[]VpsServiceInfo**](VpsServiceInfo.md) | Array of VPS services | 

## Methods

### NewListVpsServicesData

`func NewListVpsServicesData(services []VpsServiceInfo, ) *ListVpsServicesData`

NewListVpsServicesData instantiates a new ListVpsServicesData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListVpsServicesDataWithDefaults

`func NewListVpsServicesDataWithDefaults() *ListVpsServicesData`

NewListVpsServicesDataWithDefaults instantiates a new ListVpsServicesData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServices

`func (o *ListVpsServicesData) GetServices() []VpsServiceInfo`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *ListVpsServicesData) GetServicesOk() (*[]VpsServiceInfo, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *ListVpsServicesData) SetServices(v []VpsServiceInfo)`

SetServices sets Services field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


