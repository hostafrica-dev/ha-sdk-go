# UpdateSshKeysRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceId** | **string** | Service ID - must be sent as a string | 
**SshKeys** | **string** | SSH public key in OpenSSH format (must start with ssh-rsa, ssh-dss, ssh-ed25519, or ssh-ecdsa) | 

## Methods

### NewUpdateSshKeysRequestContent

`func NewUpdateSshKeysRequestContent(serviceId string, sshKeys string, ) *UpdateSshKeysRequestContent`

NewUpdateSshKeysRequestContent instantiates a new UpdateSshKeysRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSshKeysRequestContentWithDefaults

`func NewUpdateSshKeysRequestContentWithDefaults() *UpdateSshKeysRequestContent`

NewUpdateSshKeysRequestContentWithDefaults instantiates a new UpdateSshKeysRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceId

`func (o *UpdateSshKeysRequestContent) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *UpdateSshKeysRequestContent) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *UpdateSshKeysRequestContent) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetSshKeys

`func (o *UpdateSshKeysRequestContent) GetSshKeys() string`

GetSshKeys returns the SshKeys field if non-nil, zero value otherwise.

### GetSshKeysOk

`func (o *UpdateSshKeysRequestContent) GetSshKeysOk() (*string, bool)`

GetSshKeysOk returns a tuple with the SshKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeys

`func (o *UpdateSshKeysRequestContent) SetSshKeys(v string)`

SetSshKeys sets SshKeys field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


