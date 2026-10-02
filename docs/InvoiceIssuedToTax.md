# InvoiceIssuedToTax

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Tax name or region label (e.g. SA) | 
**Rate** | **float64** | Tax rate percentage | 
**Level** | **int32** | Tax level | 

## Methods

### NewInvoiceIssuedToTax

`func NewInvoiceIssuedToTax(name string, rate float64, level int32, ) *InvoiceIssuedToTax`

NewInvoiceIssuedToTax instantiates a new InvoiceIssuedToTax object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceIssuedToTaxWithDefaults

`func NewInvoiceIssuedToTaxWithDefaults() *InvoiceIssuedToTax`

NewInvoiceIssuedToTaxWithDefaults instantiates a new InvoiceIssuedToTax object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InvoiceIssuedToTax) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InvoiceIssuedToTax) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InvoiceIssuedToTax) SetName(v string)`

SetName sets Name field to given value.


### GetRate

`func (o *InvoiceIssuedToTax) GetRate() float64`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *InvoiceIssuedToTax) GetRateOk() (*float64, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *InvoiceIssuedToTax) SetRate(v float64)`

SetRate sets Rate field to given value.


### GetLevel

`func (o *InvoiceIssuedToTax) GetLevel() int32`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *InvoiceIssuedToTax) GetLevelOk() (*int32, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *InvoiceIssuedToTax) SetLevel(v int32)`

SetLevel sets Level field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


