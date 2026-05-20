# NoVncConsoleDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Port** | **string** | VNC port number | 
**Upid** | **string** | Proxmox Unique Process ID for the VNC proxy | 
**User** | **string** | Proxmox user for authentication | 
**Ticket** | **string** | Authentication ticket for VNC connection | 
**Cert** | **string** | Certificate for secure connection | 

## Methods

### NewNoVncConsoleDetails

`func NewNoVncConsoleDetails(port string, upid string, user string, ticket string, cert string, ) *NoVncConsoleDetails`

NewNoVncConsoleDetails instantiates a new NoVncConsoleDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNoVncConsoleDetailsWithDefaults

`func NewNoVncConsoleDetailsWithDefaults() *NoVncConsoleDetails`

NewNoVncConsoleDetailsWithDefaults instantiates a new NoVncConsoleDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPort

`func (o *NoVncConsoleDetails) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *NoVncConsoleDetails) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *NoVncConsoleDetails) SetPort(v string)`

SetPort sets Port field to given value.


### GetUpid

`func (o *NoVncConsoleDetails) GetUpid() string`

GetUpid returns the Upid field if non-nil, zero value otherwise.

### GetUpidOk

`func (o *NoVncConsoleDetails) GetUpidOk() (*string, bool)`

GetUpidOk returns a tuple with the Upid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpid

`func (o *NoVncConsoleDetails) SetUpid(v string)`

SetUpid sets Upid field to given value.


### GetUser

`func (o *NoVncConsoleDetails) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *NoVncConsoleDetails) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *NoVncConsoleDetails) SetUser(v string)`

SetUser sets User field to given value.


### GetTicket

`func (o *NoVncConsoleDetails) GetTicket() string`

GetTicket returns the Ticket field if non-nil, zero value otherwise.

### GetTicketOk

`func (o *NoVncConsoleDetails) GetTicketOk() (*string, bool)`

GetTicketOk returns a tuple with the Ticket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicket

`func (o *NoVncConsoleDetails) SetTicket(v string)`

SetTicket sets Ticket field to given value.


### GetCert

`func (o *NoVncConsoleDetails) GetCert() string`

GetCert returns the Cert field if non-nil, zero value otherwise.

### GetCertOk

`func (o *NoVncConsoleDetails) GetCertOk() (*string, bool)`

GetCertOk returns a tuple with the Cert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCert

`func (o *NoVncConsoleDetails) SetCert(v string)`

SetCert sets Cert field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


