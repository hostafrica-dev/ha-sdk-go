# DeleteDnssecRecordRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainId** | **string** | Domain service id - must be sent as a string | 
**KeyTag** | **int32** | DNSSEC key tag (0-65535) | 
**Alg** | **int32** | DNSSEC algorithm (1-255) | 
**DigestType** | **int32** | DNSSEC digest type (1-255) | 
**Digest** | **string** | Hex digest identifying the DS record to remove | 

## Methods

### NewDeleteDnssecRecordRequestContent

`func NewDeleteDnssecRecordRequestContent(domainId string, keyTag int32, alg int32, digestType int32, digest string, ) *DeleteDnssecRecordRequestContent`

NewDeleteDnssecRecordRequestContent instantiates a new DeleteDnssecRecordRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteDnssecRecordRequestContentWithDefaults

`func NewDeleteDnssecRecordRequestContentWithDefaults() *DeleteDnssecRecordRequestContent`

NewDeleteDnssecRecordRequestContentWithDefaults instantiates a new DeleteDnssecRecordRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomainId

`func (o *DeleteDnssecRecordRequestContent) GetDomainId() string`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *DeleteDnssecRecordRequestContent) GetDomainIdOk() (*string, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *DeleteDnssecRecordRequestContent) SetDomainId(v string)`

SetDomainId sets DomainId field to given value.


### GetKeyTag

`func (o *DeleteDnssecRecordRequestContent) GetKeyTag() int32`

GetKeyTag returns the KeyTag field if non-nil, zero value otherwise.

### GetKeyTagOk

`func (o *DeleteDnssecRecordRequestContent) GetKeyTagOk() (*int32, bool)`

GetKeyTagOk returns a tuple with the KeyTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyTag

`func (o *DeleteDnssecRecordRequestContent) SetKeyTag(v int32)`

SetKeyTag sets KeyTag field to given value.


### GetAlg

`func (o *DeleteDnssecRecordRequestContent) GetAlg() int32`

GetAlg returns the Alg field if non-nil, zero value otherwise.

### GetAlgOk

`func (o *DeleteDnssecRecordRequestContent) GetAlgOk() (*int32, bool)`

GetAlgOk returns a tuple with the Alg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlg

`func (o *DeleteDnssecRecordRequestContent) SetAlg(v int32)`

SetAlg sets Alg field to given value.


### GetDigestType

`func (o *DeleteDnssecRecordRequestContent) GetDigestType() int32`

GetDigestType returns the DigestType field if non-nil, zero value otherwise.

### GetDigestTypeOk

`func (o *DeleteDnssecRecordRequestContent) GetDigestTypeOk() (*int32, bool)`

GetDigestTypeOk returns a tuple with the DigestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigestType

`func (o *DeleteDnssecRecordRequestContent) SetDigestType(v int32)`

SetDigestType sets DigestType field to given value.


### GetDigest

`func (o *DeleteDnssecRecordRequestContent) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *DeleteDnssecRecordRequestContent) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *DeleteDnssecRecordRequestContent) SetDigest(v string)`

SetDigest sets Digest field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


