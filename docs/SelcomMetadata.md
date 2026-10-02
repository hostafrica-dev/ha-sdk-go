# SelcomMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **string** | Due amount formatted for display (e.g. \&quot;TZS 29,000.00\&quot;). Non-TZS amounts are converted using the WHMCS currency rate. | 
**PesaAccount** | **string** | Selcom Pesa account: 5-digit padded client id plus 6-digit padded invoice id | 
**MobileReference** | **string** | Mobile payment reference: 61107370 plus the same padded client and invoice ids | 
**UssdCode** | **string** | USSD shortcode for Selcom (always *150*50*1#) | 
**Instructions** | **string** | WHMCS Selcom alt-gateway instructions HTML with placeholders filled | 

## Methods

### NewSelcomMetadata

`func NewSelcomMetadata(amount string, pesaAccount string, mobileReference string, ussdCode string, instructions string, ) *SelcomMetadata`

NewSelcomMetadata instantiates a new SelcomMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelcomMetadataWithDefaults

`func NewSelcomMetadataWithDefaults() *SelcomMetadata`

NewSelcomMetadataWithDefaults instantiates a new SelcomMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *SelcomMetadata) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *SelcomMetadata) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *SelcomMetadata) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetPesaAccount

`func (o *SelcomMetadata) GetPesaAccount() string`

GetPesaAccount returns the PesaAccount field if non-nil, zero value otherwise.

### GetPesaAccountOk

`func (o *SelcomMetadata) GetPesaAccountOk() (*string, bool)`

GetPesaAccountOk returns a tuple with the PesaAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPesaAccount

`func (o *SelcomMetadata) SetPesaAccount(v string)`

SetPesaAccount sets PesaAccount field to given value.


### GetMobileReference

`func (o *SelcomMetadata) GetMobileReference() string`

GetMobileReference returns the MobileReference field if non-nil, zero value otherwise.

### GetMobileReferenceOk

`func (o *SelcomMetadata) GetMobileReferenceOk() (*string, bool)`

GetMobileReferenceOk returns a tuple with the MobileReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileReference

`func (o *SelcomMetadata) SetMobileReference(v string)`

SetMobileReference sets MobileReference field to given value.


### GetUssdCode

`func (o *SelcomMetadata) GetUssdCode() string`

GetUssdCode returns the UssdCode field if non-nil, zero value otherwise.

### GetUssdCodeOk

`func (o *SelcomMetadata) GetUssdCodeOk() (*string, bool)`

GetUssdCodeOk returns a tuple with the UssdCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUssdCode

`func (o *SelcomMetadata) SetUssdCode(v string)`

SetUssdCode sets UssdCode field to given value.


### GetInstructions

`func (o *SelcomMetadata) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *SelcomMetadata) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *SelcomMetadata) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


