# OrderLastAttempt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**At** | **string** | Timestamp of the attempt (ISO 8601) | 
**Status** | **string** | Attempt status (e.g. paid, failed) | 
**Code** | Pointer to **string** | Error code from the payment provider, if applicable | [optional] 
**Message** | Pointer to **string** | Human-readable message from the payment provider, if applicable | [optional] 

## Methods

### NewOrderLastAttempt

`func NewOrderLastAttempt(at string, status string, ) *OrderLastAttempt`

NewOrderLastAttempt instantiates a new OrderLastAttempt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderLastAttemptWithDefaults

`func NewOrderLastAttemptWithDefaults() *OrderLastAttempt`

NewOrderLastAttemptWithDefaults instantiates a new OrderLastAttempt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAt

`func (o *OrderLastAttempt) GetAt() string`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *OrderLastAttempt) GetAtOk() (*string, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *OrderLastAttempt) SetAt(v string)`

SetAt sets At field to given value.


### GetStatus

`func (o *OrderLastAttempt) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OrderLastAttempt) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OrderLastAttempt) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCode

`func (o *OrderLastAttempt) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *OrderLastAttempt) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *OrderLastAttempt) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *OrderLastAttempt) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *OrderLastAttempt) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *OrderLastAttempt) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *OrderLastAttempt) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *OrderLastAttempt) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


