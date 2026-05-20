# SshKeyDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PublicKey** | **string** | The public SSH key in OpenSSH format | 
**Type** | **string** | Type of SSH key configuration - &#39;sshkeys&#39; for KVM cloud-init or &#39;keypair&#39; for traditional SSH (LXC) | 

## Methods

### NewSshKeyDetails

`func NewSshKeyDetails(publicKey string, type_ string, ) *SshKeyDetails`

NewSshKeyDetails instantiates a new SshKeyDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSshKeyDetailsWithDefaults

`func NewSshKeyDetailsWithDefaults() *SshKeyDetails`

NewSshKeyDetailsWithDefaults instantiates a new SshKeyDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublicKey

`func (o *SshKeyDetails) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *SshKeyDetails) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *SshKeyDetails) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.


### GetType

`func (o *SshKeyDetails) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SshKeyDetails) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SshKeyDetails) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


