# NoVncConsoleDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NovncRedirectUrl** | Pointer to **string** | Redirect URL for the noVNC console | [optional] 
**Mode** | Pointer to **string** | Console connection mode | [optional] 
**WebsocketUrl** | Pointer to **string** | WebSocket URL for proxied console access | [optional] 
**Password** | Pointer to **string** | Password for proxied console access | [optional] 

## Methods

### NewNoVncConsoleDetails

`func NewNoVncConsoleDetails() *NoVncConsoleDetails`

NewNoVncConsoleDetails instantiates a new NoVncConsoleDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNoVncConsoleDetailsWithDefaults

`func NewNoVncConsoleDetailsWithDefaults() *NoVncConsoleDetails`

NewNoVncConsoleDetailsWithDefaults instantiates a new NoVncConsoleDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNovncRedirectUrl

`func (o *NoVncConsoleDetails) GetNovncRedirectUrl() string`

GetNovncRedirectUrl returns the NovncRedirectUrl field if non-nil, zero value otherwise.

### GetNovncRedirectUrlOk

`func (o *NoVncConsoleDetails) GetNovncRedirectUrlOk() (*string, bool)`

GetNovncRedirectUrlOk returns a tuple with the NovncRedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNovncRedirectUrl

`func (o *NoVncConsoleDetails) SetNovncRedirectUrl(v string)`

SetNovncRedirectUrl sets NovncRedirectUrl field to given value.

### HasNovncRedirectUrl

`func (o *NoVncConsoleDetails) HasNovncRedirectUrl() bool`

HasNovncRedirectUrl returns a boolean if a field has been set.

### GetMode

`func (o *NoVncConsoleDetails) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *NoVncConsoleDetails) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *NoVncConsoleDetails) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *NoVncConsoleDetails) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetWebsocketUrl

`func (o *NoVncConsoleDetails) GetWebsocketUrl() string`

GetWebsocketUrl returns the WebsocketUrl field if non-nil, zero value otherwise.

### GetWebsocketUrlOk

`func (o *NoVncConsoleDetails) GetWebsocketUrlOk() (*string, bool)`

GetWebsocketUrlOk returns a tuple with the WebsocketUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsocketUrl

`func (o *NoVncConsoleDetails) SetWebsocketUrl(v string)`

SetWebsocketUrl sets WebsocketUrl field to given value.

### HasWebsocketUrl

`func (o *NoVncConsoleDetails) HasWebsocketUrl() bool`

HasWebsocketUrl returns a boolean if a field has been set.

### GetPassword

`func (o *NoVncConsoleDetails) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *NoVncConsoleDetails) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *NoVncConsoleDetails) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *NoVncConsoleDetails) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


