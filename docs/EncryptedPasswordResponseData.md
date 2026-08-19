# EncryptedPasswordResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** | Username for VPS access (plaintext) | 
**Password** | **string** | Base64-encoded ciphertext of the VPS password. Produced with RSA-OAEP (SHA-256) and the request public_key. Decode from base64, then decrypt with the matching private key via openssl pkeyutl -decrypt -pkeyopt rsa_padding_mode:oaep -pkeyopt rsa_oaep_md:sha256 -pkeyopt rsa_mgf1_md:sha256. | 
**Encryption** | [**PasswordEncryptionInfo**](PasswordEncryptionInfo.md) |  | 

## Methods

### NewEncryptedPasswordResponseData

`func NewEncryptedPasswordResponseData(username string, password string, encryption PasswordEncryptionInfo, ) *EncryptedPasswordResponseData`

NewEncryptedPasswordResponseData instantiates a new EncryptedPasswordResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEncryptedPasswordResponseDataWithDefaults

`func NewEncryptedPasswordResponseDataWithDefaults() *EncryptedPasswordResponseData`

NewEncryptedPasswordResponseDataWithDefaults instantiates a new EncryptedPasswordResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *EncryptedPasswordResponseData) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *EncryptedPasswordResponseData) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *EncryptedPasswordResponseData) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *EncryptedPasswordResponseData) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *EncryptedPasswordResponseData) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *EncryptedPasswordResponseData) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetEncryption

`func (o *EncryptedPasswordResponseData) GetEncryption() PasswordEncryptionInfo`

GetEncryption returns the Encryption field if non-nil, zero value otherwise.

### GetEncryptionOk

`func (o *EncryptedPasswordResponseData) GetEncryptionOk() (*PasswordEncryptionInfo, bool)`

GetEncryptionOk returns a tuple with the Encryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryption

`func (o *EncryptedPasswordResponseData) SetEncryption(v PasswordEncryptionInfo)`

SetEncryption sets Encryption field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


