# SuggestDomainsRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Prompt** | **string** | Natural-language description of the business or idea, e.g. \&quot;my coffee shop\&quot; | 
**Currency** | Pointer to **string** | Currency for pricing. Accepts ISO code, symbol, country name, or numeric ID. Defaults to ZAR. | [optional] 

## Methods

### NewSuggestDomainsRequestContent

`func NewSuggestDomainsRequestContent(prompt string, ) *SuggestDomainsRequestContent`

NewSuggestDomainsRequestContent instantiates a new SuggestDomainsRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuggestDomainsRequestContentWithDefaults

`func NewSuggestDomainsRequestContentWithDefaults() *SuggestDomainsRequestContent`

NewSuggestDomainsRequestContentWithDefaults instantiates a new SuggestDomainsRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrompt

`func (o *SuggestDomainsRequestContent) GetPrompt() string`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *SuggestDomainsRequestContent) GetPromptOk() (*string, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *SuggestDomainsRequestContent) SetPrompt(v string)`

SetPrompt sets Prompt field to given value.


### GetCurrency

`func (o *SuggestDomainsRequestContent) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *SuggestDomainsRequestContent) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *SuggestDomainsRequestContent) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *SuggestDomainsRequestContent) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


