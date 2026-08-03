# DomainHostingLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostingId** | Pointer to **int32** | Linked hosting service id; omitted when no hosting is linked | [optional] 
**Module** | Pointer to **string** | Hosting module name (e.g. cpanel); omitted when no hosting is linked | [optional] 

## Methods

### NewDomainHostingLink

`func NewDomainHostingLink() *DomainHostingLink`

NewDomainHostingLink instantiates a new DomainHostingLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainHostingLinkWithDefaults

`func NewDomainHostingLinkWithDefaults() *DomainHostingLink`

NewDomainHostingLinkWithDefaults instantiates a new DomainHostingLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostingId

`func (o *DomainHostingLink) GetHostingId() int32`

GetHostingId returns the HostingId field if non-nil, zero value otherwise.

### GetHostingIdOk

`func (o *DomainHostingLink) GetHostingIdOk() (*int32, bool)`

GetHostingIdOk returns a tuple with the HostingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostingId

`func (o *DomainHostingLink) SetHostingId(v int32)`

SetHostingId sets HostingId field to given value.

### HasHostingId

`func (o *DomainHostingLink) HasHostingId() bool`

HasHostingId returns a boolean if a field has been set.

### GetModule

`func (o *DomainHostingLink) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *DomainHostingLink) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *DomainHostingLink) SetModule(v string)`

SetModule sets Module field to given value.

### HasModule

`func (o *DomainHostingLink) HasModule() bool`

HasModule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


