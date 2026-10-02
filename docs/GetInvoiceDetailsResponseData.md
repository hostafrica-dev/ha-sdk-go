# GetInvoiceDetailsResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Invoice** | [**InvoiceDetails**](InvoiceDetails.md) |  | 
**SelcomMetadata** | Pointer to [**SelcomMetadata**](SelcomMetadata.md) |  | [optional] 

## Methods

### NewGetInvoiceDetailsResponseData

`func NewGetInvoiceDetailsResponseData(invoice InvoiceDetails, ) *GetInvoiceDetailsResponseData`

NewGetInvoiceDetailsResponseData instantiates a new GetInvoiceDetailsResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInvoiceDetailsResponseDataWithDefaults

`func NewGetInvoiceDetailsResponseDataWithDefaults() *GetInvoiceDetailsResponseData`

NewGetInvoiceDetailsResponseDataWithDefaults instantiates a new GetInvoiceDetailsResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvoice

`func (o *GetInvoiceDetailsResponseData) GetInvoice() InvoiceDetails`

GetInvoice returns the Invoice field if non-nil, zero value otherwise.

### GetInvoiceOk

`func (o *GetInvoiceDetailsResponseData) GetInvoiceOk() (*InvoiceDetails, bool)`

GetInvoiceOk returns a tuple with the Invoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoice

`func (o *GetInvoiceDetailsResponseData) SetInvoice(v InvoiceDetails)`

SetInvoice sets Invoice field to given value.


### GetSelcomMetadata

`func (o *GetInvoiceDetailsResponseData) GetSelcomMetadata() SelcomMetadata`

GetSelcomMetadata returns the SelcomMetadata field if non-nil, zero value otherwise.

### GetSelcomMetadataOk

`func (o *GetInvoiceDetailsResponseData) GetSelcomMetadataOk() (*SelcomMetadata, bool)`

GetSelcomMetadataOk returns a tuple with the SelcomMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelcomMetadata

`func (o *GetInvoiceDetailsResponseData) SetSelcomMetadata(v SelcomMetadata)`

SetSelcomMetadata sets SelcomMetadata field to given value.

### HasSelcomMetadata

`func (o *GetInvoiceDetailsResponseData) HasSelcomMetadata() bool`

HasSelcomMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


