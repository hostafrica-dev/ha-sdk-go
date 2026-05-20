# PowerTaskDialogRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequiredOnCreate** | **[]string** | Fields required when creating a new power task | 
**JobTypeValues** | **[]string** | Valid job type values | 
**ActionValues** | **[]string** | Valid action values | 
**WeekdayValues** | **[]string** | Valid weekday abbreviations | 
**JobMinutesStep** | **int32** | Minute step increment for job time selection | 

## Methods

### NewPowerTaskDialogRules

`func NewPowerTaskDialogRules(requiredOnCreate []string, jobTypeValues []string, actionValues []string, weekdayValues []string, jobMinutesStep int32, ) *PowerTaskDialogRules`

NewPowerTaskDialogRules instantiates a new PowerTaskDialogRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPowerTaskDialogRulesWithDefaults

`func NewPowerTaskDialogRulesWithDefaults() *PowerTaskDialogRules`

NewPowerTaskDialogRulesWithDefaults instantiates a new PowerTaskDialogRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequiredOnCreate

`func (o *PowerTaskDialogRules) GetRequiredOnCreate() []string`

GetRequiredOnCreate returns the RequiredOnCreate field if non-nil, zero value otherwise.

### GetRequiredOnCreateOk

`func (o *PowerTaskDialogRules) GetRequiredOnCreateOk() (*[]string, bool)`

GetRequiredOnCreateOk returns a tuple with the RequiredOnCreate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredOnCreate

`func (o *PowerTaskDialogRules) SetRequiredOnCreate(v []string)`

SetRequiredOnCreate sets RequiredOnCreate field to given value.


### GetJobTypeValues

`func (o *PowerTaskDialogRules) GetJobTypeValues() []string`

GetJobTypeValues returns the JobTypeValues field if non-nil, zero value otherwise.

### GetJobTypeValuesOk

`func (o *PowerTaskDialogRules) GetJobTypeValuesOk() (*[]string, bool)`

GetJobTypeValuesOk returns a tuple with the JobTypeValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTypeValues

`func (o *PowerTaskDialogRules) SetJobTypeValues(v []string)`

SetJobTypeValues sets JobTypeValues field to given value.


### GetActionValues

`func (o *PowerTaskDialogRules) GetActionValues() []string`

GetActionValues returns the ActionValues field if non-nil, zero value otherwise.

### GetActionValuesOk

`func (o *PowerTaskDialogRules) GetActionValuesOk() (*[]string, bool)`

GetActionValuesOk returns a tuple with the ActionValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionValues

`func (o *PowerTaskDialogRules) SetActionValues(v []string)`

SetActionValues sets ActionValues field to given value.


### GetWeekdayValues

`func (o *PowerTaskDialogRules) GetWeekdayValues() []string`

GetWeekdayValues returns the WeekdayValues field if non-nil, zero value otherwise.

### GetWeekdayValuesOk

`func (o *PowerTaskDialogRules) GetWeekdayValuesOk() (*[]string, bool)`

GetWeekdayValuesOk returns a tuple with the WeekdayValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekdayValues

`func (o *PowerTaskDialogRules) SetWeekdayValues(v []string)`

SetWeekdayValues sets WeekdayValues field to given value.


### GetJobMinutesStep

`func (o *PowerTaskDialogRules) GetJobMinutesStep() int32`

GetJobMinutesStep returns the JobMinutesStep field if non-nil, zero value otherwise.

### GetJobMinutesStepOk

`func (o *PowerTaskDialogRules) GetJobMinutesStepOk() (*int32, bool)`

GetJobMinutesStepOk returns a tuple with the JobMinutesStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobMinutesStep

`func (o *PowerTaskDialogRules) SetJobMinutesStep(v int32)`

SetJobMinutesStep sets JobMinutesStep field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


