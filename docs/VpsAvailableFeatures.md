# VpsAvailableFeatures

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PowerStart** | **bool** | Can start the VPS | 
**PowerStop** | **bool** | Can stop the VPS | 
**PowerReboot** | **bool** | Can reboot the VPS | 
**PowerShutdown** | **bool** | Can shutdown the VPS | 
**NovncConsole** | **bool** | Has noVNC console access | 
**Backups** | **bool** | Has backup capabilities | 
**BackupJobs** | **bool** | Has backup jobs feature | 
**BackupSchedule** | **bool** | Has backup schedule feature | 
**Firewall** | **bool** | Has firewall feature | 
**Reinstall** | **bool** | Can reinstall OS | 
**ChangeHostname** | **bool** | Can change hostname | 
**ChangeIsoImage** | **bool** | Can change ISO image | 
**NetworkStats** | **bool** | Has network statistics feature | 
**Graphs** | **bool** | Has graphs feature | 
**OsTemplates** | **[]string** | List of available OS templates | 

## Methods

### NewVpsAvailableFeatures

`func NewVpsAvailableFeatures(powerStart bool, powerStop bool, powerReboot bool, powerShutdown bool, novncConsole bool, backups bool, backupJobs bool, backupSchedule bool, firewall bool, reinstall bool, changeHostname bool, changeIsoImage bool, networkStats bool, graphs bool, osTemplates []string, ) *VpsAvailableFeatures`

NewVpsAvailableFeatures instantiates a new VpsAvailableFeatures object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpsAvailableFeaturesWithDefaults

`func NewVpsAvailableFeaturesWithDefaults() *VpsAvailableFeatures`

NewVpsAvailableFeaturesWithDefaults instantiates a new VpsAvailableFeatures object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPowerStart

`func (o *VpsAvailableFeatures) GetPowerStart() bool`

GetPowerStart returns the PowerStart field if non-nil, zero value otherwise.

### GetPowerStartOk

`func (o *VpsAvailableFeatures) GetPowerStartOk() (*bool, bool)`

GetPowerStartOk returns a tuple with the PowerStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerStart

`func (o *VpsAvailableFeatures) SetPowerStart(v bool)`

SetPowerStart sets PowerStart field to given value.


### GetPowerStop

`func (o *VpsAvailableFeatures) GetPowerStop() bool`

GetPowerStop returns the PowerStop field if non-nil, zero value otherwise.

### GetPowerStopOk

`func (o *VpsAvailableFeatures) GetPowerStopOk() (*bool, bool)`

GetPowerStopOk returns a tuple with the PowerStop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerStop

`func (o *VpsAvailableFeatures) SetPowerStop(v bool)`

SetPowerStop sets PowerStop field to given value.


### GetPowerReboot

`func (o *VpsAvailableFeatures) GetPowerReboot() bool`

GetPowerReboot returns the PowerReboot field if non-nil, zero value otherwise.

### GetPowerRebootOk

`func (o *VpsAvailableFeatures) GetPowerRebootOk() (*bool, bool)`

GetPowerRebootOk returns a tuple with the PowerReboot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerReboot

`func (o *VpsAvailableFeatures) SetPowerReboot(v bool)`

SetPowerReboot sets PowerReboot field to given value.


### GetPowerShutdown

`func (o *VpsAvailableFeatures) GetPowerShutdown() bool`

GetPowerShutdown returns the PowerShutdown field if non-nil, zero value otherwise.

### GetPowerShutdownOk

`func (o *VpsAvailableFeatures) GetPowerShutdownOk() (*bool, bool)`

GetPowerShutdownOk returns a tuple with the PowerShutdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerShutdown

`func (o *VpsAvailableFeatures) SetPowerShutdown(v bool)`

SetPowerShutdown sets PowerShutdown field to given value.


### GetNovncConsole

`func (o *VpsAvailableFeatures) GetNovncConsole() bool`

GetNovncConsole returns the NovncConsole field if non-nil, zero value otherwise.

### GetNovncConsoleOk

`func (o *VpsAvailableFeatures) GetNovncConsoleOk() (*bool, bool)`

GetNovncConsoleOk returns a tuple with the NovncConsole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNovncConsole

`func (o *VpsAvailableFeatures) SetNovncConsole(v bool)`

SetNovncConsole sets NovncConsole field to given value.


### GetBackups

`func (o *VpsAvailableFeatures) GetBackups() bool`

GetBackups returns the Backups field if non-nil, zero value otherwise.

### GetBackupsOk

`func (o *VpsAvailableFeatures) GetBackupsOk() (*bool, bool)`

GetBackupsOk returns a tuple with the Backups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackups

`func (o *VpsAvailableFeatures) SetBackups(v bool)`

SetBackups sets Backups field to given value.


### GetBackupJobs

`func (o *VpsAvailableFeatures) GetBackupJobs() bool`

GetBackupJobs returns the BackupJobs field if non-nil, zero value otherwise.

### GetBackupJobsOk

`func (o *VpsAvailableFeatures) GetBackupJobsOk() (*bool, bool)`

GetBackupJobsOk returns a tuple with the BackupJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupJobs

`func (o *VpsAvailableFeatures) SetBackupJobs(v bool)`

SetBackupJobs sets BackupJobs field to given value.


### GetBackupSchedule

`func (o *VpsAvailableFeatures) GetBackupSchedule() bool`

GetBackupSchedule returns the BackupSchedule field if non-nil, zero value otherwise.

### GetBackupScheduleOk

`func (o *VpsAvailableFeatures) GetBackupScheduleOk() (*bool, bool)`

GetBackupScheduleOk returns a tuple with the BackupSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupSchedule

`func (o *VpsAvailableFeatures) SetBackupSchedule(v bool)`

SetBackupSchedule sets BackupSchedule field to given value.


### GetFirewall

`func (o *VpsAvailableFeatures) GetFirewall() bool`

GetFirewall returns the Firewall field if non-nil, zero value otherwise.

### GetFirewallOk

`func (o *VpsAvailableFeatures) GetFirewallOk() (*bool, bool)`

GetFirewallOk returns a tuple with the Firewall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewall

`func (o *VpsAvailableFeatures) SetFirewall(v bool)`

SetFirewall sets Firewall field to given value.


### GetReinstall

`func (o *VpsAvailableFeatures) GetReinstall() bool`

GetReinstall returns the Reinstall field if non-nil, zero value otherwise.

### GetReinstallOk

`func (o *VpsAvailableFeatures) GetReinstallOk() (*bool, bool)`

GetReinstallOk returns a tuple with the Reinstall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReinstall

`func (o *VpsAvailableFeatures) SetReinstall(v bool)`

SetReinstall sets Reinstall field to given value.


### GetChangeHostname

`func (o *VpsAvailableFeatures) GetChangeHostname() bool`

GetChangeHostname returns the ChangeHostname field if non-nil, zero value otherwise.

### GetChangeHostnameOk

`func (o *VpsAvailableFeatures) GetChangeHostnameOk() (*bool, bool)`

GetChangeHostnameOk returns a tuple with the ChangeHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeHostname

`func (o *VpsAvailableFeatures) SetChangeHostname(v bool)`

SetChangeHostname sets ChangeHostname field to given value.


### GetChangeIsoImage

`func (o *VpsAvailableFeatures) GetChangeIsoImage() bool`

GetChangeIsoImage returns the ChangeIsoImage field if non-nil, zero value otherwise.

### GetChangeIsoImageOk

`func (o *VpsAvailableFeatures) GetChangeIsoImageOk() (*bool, bool)`

GetChangeIsoImageOk returns a tuple with the ChangeIsoImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeIsoImage

`func (o *VpsAvailableFeatures) SetChangeIsoImage(v bool)`

SetChangeIsoImage sets ChangeIsoImage field to given value.


### GetNetworkStats

`func (o *VpsAvailableFeatures) GetNetworkStats() bool`

GetNetworkStats returns the NetworkStats field if non-nil, zero value otherwise.

### GetNetworkStatsOk

`func (o *VpsAvailableFeatures) GetNetworkStatsOk() (*bool, bool)`

GetNetworkStatsOk returns a tuple with the NetworkStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkStats

`func (o *VpsAvailableFeatures) SetNetworkStats(v bool)`

SetNetworkStats sets NetworkStats field to given value.


### GetGraphs

`func (o *VpsAvailableFeatures) GetGraphs() bool`

GetGraphs returns the Graphs field if non-nil, zero value otherwise.

### GetGraphsOk

`func (o *VpsAvailableFeatures) GetGraphsOk() (*bool, bool)`

GetGraphsOk returns a tuple with the Graphs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraphs

`func (o *VpsAvailableFeatures) SetGraphs(v bool)`

SetGraphs sets Graphs field to given value.


### GetOsTemplates

`func (o *VpsAvailableFeatures) GetOsTemplates() []string`

GetOsTemplates returns the OsTemplates field if non-nil, zero value otherwise.

### GetOsTemplatesOk

`func (o *VpsAvailableFeatures) GetOsTemplatesOk() (*[]string, bool)`

GetOsTemplatesOk returns a tuple with the OsTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsTemplates

`func (o *VpsAvailableFeatures) SetOsTemplates(v []string)`

SetOsTemplates sets OsTemplates field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


