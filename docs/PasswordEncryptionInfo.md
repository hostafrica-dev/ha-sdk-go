# PasswordEncryptionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Algorithm** | **string** | Asymmetric encryption algorithm. Always \&quot;RSA-OAEP\&quot;. | 
**Hash** | **string** | OAEP hash / MGF1 hash function. Always \&quot;SHA-256\&quot;. | 
**KeySize** | **int32** | RSA modulus size in bits. Always 4096. | 
**Encoding** | **string** | Encoding of the password ciphertext field. Always \&quot;base64\&quot; (standard alphabet, not URL-safe). | 

## Methods

### NewPasswordEncryptionInfo

`func NewPasswordEncryptionInfo(algorithm string, hash string, keySize int32, encoding string, ) *PasswordEncryptionInfo`

NewPasswordEncryptionInfo instantiates a new PasswordEncryptionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPasswordEncryptionInfoWithDefaults

`func NewPasswordEncryptionInfoWithDefaults() *PasswordEncryptionInfo`

NewPasswordEncryptionInfoWithDefaults instantiates a new PasswordEncryptionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgorithm

`func (o *PasswordEncryptionInfo) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *PasswordEncryptionInfo) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *PasswordEncryptionInfo) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.


### GetHash

`func (o *PasswordEncryptionInfo) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *PasswordEncryptionInfo) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *PasswordEncryptionInfo) SetHash(v string)`

SetHash sets Hash field to given value.


### GetKeySize

`func (o *PasswordEncryptionInfo) GetKeySize() int32`

GetKeySize returns the KeySize field if non-nil, zero value otherwise.

### GetKeySizeOk

`func (o *PasswordEncryptionInfo) GetKeySizeOk() (*int32, bool)`

GetKeySizeOk returns a tuple with the KeySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySize

`func (o *PasswordEncryptionInfo) SetKeySize(v int32)`

SetKeySize sets KeySize field to given value.


### GetEncoding

`func (o *PasswordEncryptionInfo) GetEncoding() string`

GetEncoding returns the Encoding field if non-nil, zero value otherwise.

### GetEncodingOk

`func (o *PasswordEncryptionInfo) GetEncodingOk() (*string, bool)`

GetEncodingOk returns a tuple with the Encoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoding

`func (o *PasswordEncryptionInfo) SetEncoding(v string)`

SetEncoding sets Encoding field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


