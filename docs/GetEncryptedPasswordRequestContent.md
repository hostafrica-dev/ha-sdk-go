# GetEncryptedPasswordRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceId** | **string** | Service ID - must be sent as a string | 
**PublicKey** | **string** | PEM-encoded RSA public key only (never the private key). Accepts SPKI (&#x60;-----BEGIN PUBLIC KEY-----&#x60;) or PKCS#1 (&#x60;-----BEGIN RSA PUBLIC KEY-----&#x60;). Must be exactly 4096-bit. Used with RSA-OAEP and SHA-256 to encrypt the password. | 

## Methods

### NewGetEncryptedPasswordRequestContent

`func NewGetEncryptedPasswordRequestContent(serviceId string, publicKey string, ) *GetEncryptedPasswordRequestContent`

NewGetEncryptedPasswordRequestContent instantiates a new GetEncryptedPasswordRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEncryptedPasswordRequestContentWithDefaults

`func NewGetEncryptedPasswordRequestContentWithDefaults() *GetEncryptedPasswordRequestContent`

NewGetEncryptedPasswordRequestContentWithDefaults instantiates a new GetEncryptedPasswordRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceId

`func (o *GetEncryptedPasswordRequestContent) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *GetEncryptedPasswordRequestContent) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *GetEncryptedPasswordRequestContent) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetPublicKey

`func (o *GetEncryptedPasswordRequestContent) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *GetEncryptedPasswordRequestContent) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *GetEncryptedPasswordRequestContent) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


