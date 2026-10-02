# InvoiceIssuedTo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyName** | **string** | Company name on the invoice | 
**FullName** | **string** | Full name of the billed contact | 
**Address** | **string** | Street address | 
**CityStateZip** | **string** | City, state/province, and postal code | 
**Country** | **string** | ISO country code | 
**CountryName** | **string** | Country display name | 
**Tax** | Pointer to [**InvoiceIssuedToTax**](InvoiceIssuedToTax.md) |  | [optional] 

## Methods

### NewInvoiceIssuedTo

`func NewInvoiceIssuedTo(companyName string, fullName string, address string, cityStateZip string, country string, countryName string, ) *InvoiceIssuedTo`

NewInvoiceIssuedTo instantiates a new InvoiceIssuedTo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceIssuedToWithDefaults

`func NewInvoiceIssuedToWithDefaults() *InvoiceIssuedTo`

NewInvoiceIssuedToWithDefaults instantiates a new InvoiceIssuedTo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyName

`func (o *InvoiceIssuedTo) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *InvoiceIssuedTo) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *InvoiceIssuedTo) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.


### GetFullName

`func (o *InvoiceIssuedTo) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *InvoiceIssuedTo) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *InvoiceIssuedTo) SetFullName(v string)`

SetFullName sets FullName field to given value.


### GetAddress

`func (o *InvoiceIssuedTo) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *InvoiceIssuedTo) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *InvoiceIssuedTo) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetCityStateZip

`func (o *InvoiceIssuedTo) GetCityStateZip() string`

GetCityStateZip returns the CityStateZip field if non-nil, zero value otherwise.

### GetCityStateZipOk

`func (o *InvoiceIssuedTo) GetCityStateZipOk() (*string, bool)`

GetCityStateZipOk returns a tuple with the CityStateZip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityStateZip

`func (o *InvoiceIssuedTo) SetCityStateZip(v string)`

SetCityStateZip sets CityStateZip field to given value.


### GetCountry

`func (o *InvoiceIssuedTo) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *InvoiceIssuedTo) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *InvoiceIssuedTo) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetCountryName

`func (o *InvoiceIssuedTo) GetCountryName() string`

GetCountryName returns the CountryName field if non-nil, zero value otherwise.

### GetCountryNameOk

`func (o *InvoiceIssuedTo) GetCountryNameOk() (*string, bool)`

GetCountryNameOk returns a tuple with the CountryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryName

`func (o *InvoiceIssuedTo) SetCountryName(v string)`

SetCountryName sets CountryName field to given value.


### GetTax

`func (o *InvoiceIssuedTo) GetTax() InvoiceIssuedToTax`

GetTax returns the Tax field if non-nil, zero value otherwise.

### GetTaxOk

`func (o *InvoiceIssuedTo) GetTaxOk() (*InvoiceIssuedToTax, bool)`

GetTaxOk returns a tuple with the Tax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTax

`func (o *InvoiceIssuedTo) SetTax(v InvoiceIssuedToTax)`

SetTax sets Tax field to given value.

### HasTax

`func (o *InvoiceIssuedTo) HasTax() bool`

HasTax returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


