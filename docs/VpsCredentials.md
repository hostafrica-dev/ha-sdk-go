# VpsCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** | Username for VPS access | 
**Password** | **string** | Password for VPS access. Always returned as \&quot;&lt;redacted&gt;\&quot; from get-details; plaintext passwords are never included in API responses. | 

## Methods

### NewVpsCredentials

`func NewVpsCredentials(username string, password string, ) *VpsCredentials`

NewVpsCredentials instantiates a new VpsCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpsCredentialsWithDefaults

`func NewVpsCredentialsWithDefaults() *VpsCredentials`

NewVpsCredentialsWithDefaults instantiates a new VpsCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *VpsCredentials) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *VpsCredentials) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *VpsCredentials) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *VpsCredentials) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *VpsCredentials) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *VpsCredentials) SetPassword(v string)`

SetPassword sets Password field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


