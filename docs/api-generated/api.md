# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [api.proto](#api.proto)
    - [Endpoint](#v1.Endpoint)
    - [GetServerInformationMsg](#v1.GetServerInformationMsg)
    - [GetServerInformationReply](#v1.GetServerInformationReply)
    - [GetServerInformationReply.Buckets](#v1.GetServerInformationReply.Buckets)
    - [GetServerInformationReply.ErasureBackendType](#v1.GetServerInformationReply.ErasureBackendType)
    - [GetServerInformationReply.FSBackendType](#v1.GetServerInformationReply.FSBackendType)
    - [GetServerInformationReply.Objects](#v1.GetServerInformationReply.Objects)
    - [GetServerInformationReply.ServerProperties](#v1.GetServerInformationReply.ServerProperties)
    - [GetServerInformationReply.ServerProperties.Disk](#v1.GetServerInformationReply.ServerProperties.Disk)
    - [GetServerInformationReply.ServerProperties.Network](#v1.GetServerInformationReply.ServerProperties.Network)
    - [GetServerInformationReply.Services](#v1.GetServerInformationReply.Services)
    - [GetServerInformationReply.Services.Audit](#v1.GetServerInformationReply.Services.Audit)
    - [GetServerInformationReply.Services.Audit.AuditMap](#v1.GetServerInformationReply.Services.Audit.AuditMap)
    - [GetServerInformationReply.Services.LDAP](#v1.GetServerInformationReply.Services.LDAP)
    - [GetServerInformationReply.Services.Logger](#v1.GetServerInformationReply.Services.Logger)
    - [GetServerInformationReply.Services.Logger.LoggerMap](#v1.GetServerInformationReply.Services.Logger.LoggerMap)
    - [GetServerInformationReply.Services.Notifications](#v1.GetServerInformationReply.Services.Notifications)
    - [GetServerInformationReply.Services.Notifications.NotificationMap](#v1.GetServerInformationReply.Services.Notifications.NotificationMap)
    - [GetServerInformationReply.Services.Notifications.NotificationMap.TargetIDStatus](#v1.GetServerInformationReply.Services.Notifications.NotificationMap.TargetIDStatus)
    - [GetServerInformationReply.Services.Notifications.NotificationMap.TargetIDStatus.TargetIDStatusMap](#v1.GetServerInformationReply.Services.Notifications.NotificationMap.TargetIDStatus.TargetIDStatusMap)
    - [GetServerInformationReply.Services.Vault](#v1.GetServerInformationReply.Services.Vault)
    - [GetServerInformationReply.Usage](#v1.GetServerInformationReply.Usage)
    - [GetVersionMsg](#v1.GetVersionMsg)
    - [GetVersionReply](#v1.GetVersionReply)
    - [GetVersionReply.VersionInformation](#v1.GetVersionReply.VersionInformation)
  
    - [minioadmin](#v1.minioadmin)
  
- [api.proto](#api.proto)
    - [Endpoint](#v1.Endpoint)
    - [GetServerInformationMsg](#v1.GetServerInformationMsg)
    - [GetServerInformationReply](#v1.GetServerInformationReply)
    - [GetServerInformationReply.Buckets](#v1.GetServerInformationReply.Buckets)
    - [GetServerInformationReply.ErasureBackendType](#v1.GetServerInformationReply.ErasureBackendType)
    - [GetServerInformationReply.FSBackendType](#v1.GetServerInformationReply.FSBackendType)
    - [GetServerInformationReply.Objects](#v1.GetServerInformationReply.Objects)
    - [GetServerInformationReply.ServerProperties](#v1.GetServerInformationReply.ServerProperties)
    - [GetServerInformationReply.ServerProperties.Disk](#v1.GetServerInformationReply.ServerProperties.Disk)
    - [GetServerInformationReply.ServerProperties.Network](#v1.GetServerInformationReply.ServerProperties.Network)
    - [GetServerInformationReply.Services](#v1.GetServerInformationReply.Services)
    - [GetServerInformationReply.Services.Audit](#v1.GetServerInformationReply.Services.Audit)
    - [GetServerInformationReply.Services.Audit.AuditMap](#v1.GetServerInformationReply.Services.Audit.AuditMap)
    - [GetServerInformationReply.Services.LDAP](#v1.GetServerInformationReply.Services.LDAP)
    - [GetServerInformationReply.Services.Logger](#v1.GetServerInformationReply.Services.Logger)
    - [GetServerInformationReply.Services.Logger.LoggerMap](#v1.GetServerInformationReply.Services.Logger.LoggerMap)
    - [GetServerInformationReply.Services.Notifications](#v1.GetServerInformationReply.Services.Notifications)
    - [GetServerInformationReply.Services.Notifications.NotificationMap](#v1.GetServerInformationReply.Services.Notifications.NotificationMap)
    - [GetServerInformationReply.Services.Notifications.NotificationMap.TargetIDStatus](#v1.GetServerInformationReply.Services.Notifications.NotificationMap.TargetIDStatus)
    - [GetServerInformationReply.Services.Notifications.NotificationMap.TargetIDStatus.TargetIDStatusMap](#v1.GetServerInformationReply.Services.Notifications.NotificationMap.TargetIDStatus.TargetIDStatusMap)
    - [GetServerInformationReply.Services.Vault](#v1.GetServerInformationReply.Services.Vault)
    - [GetServerInformationReply.Usage](#v1.GetServerInformationReply.Usage)
    - [GetVersionMsg](#v1.GetVersionMsg)
    - [GetVersionReply](#v1.GetVersionReply)
    - [GetVersionReply.VersionInformation](#v1.GetVersionReply.VersionInformation)
  
    - [minioadmin](#v1.minioadmin)
  
- [Scalar Value Types](#scalar-value-types)



<a name="api.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api.proto



<a name="v1.Endpoint"></a>

### Endpoint



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| endpoint | [string](#string) |  |  |
| access_key_id | [string](#string) |  |  |
| secret_access_key | [string](#string) |  |  |
| ssl | [bool](#bool) |  |  |






<a name="v1.GetServerInformationMsg"></a>

### GetServerInformationMsg



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| endpoint | [Endpoint](#v1.Endpoint) |  |  |






<a name="v1.GetServerInformationReply"></a>

### GetServerInformationReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| mode | [string](#string) |  |  |
| domain | [string](#string) | repeated |  |
| region | [string](#string) |  |  |
| sqs_arn | [string](#string) | repeated |  |
| deployment_id | [string](#string) |  |  |
| buckets | [GetServerInformationReply.Buckets](#v1.GetServerInformationReply.Buckets) |  |  |
| objects | [GetServerInformationReply.Objects](#v1.GetServerInformationReply.Objects) |  |  |
| usage | [GetServerInformationReply.Usage](#v1.GetServerInformationReply.Usage) |  |  |
| services | [GetServerInformationReply.Services](#v1.GetServerInformationReply.Services) |  |  |
| fs_backend | [GetServerInformationReply.FSBackendType](#v1.GetServerInformationReply.FSBackendType) |  |  |
| erasure_backend | [GetServerInformationReply.ErasureBackendType](#v1.GetServerInformationReply.ErasureBackendType) |  |  |
| servers | [GetServerInformationReply.ServerProperties](#v1.GetServerInformationReply.ServerProperties) | repeated |  |






<a name="v1.GetServerInformationReply.Buckets"></a>

### GetServerInformationReply.Buckets



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| count | [uint64](#uint64) |  |  |






<a name="v1.GetServerInformationReply.ErasureBackendType"></a>

### GetServerInformationReply.ErasureBackendType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| backend_type | [string](#string) |  |  |
| online_disks | [int32](#int32) |  |  |
| offline_disks | [int32](#int32) |  |  |
| standard_sc_data | [int32](#int32) |  |  |
| standard_sc_parity | [int32](#int32) |  |  |
| rrsc_data | [int32](#int32) |  |  |
| rrsc_parity | [int32](#int32) |  |  |






<a name="v1.GetServerInformationReply.FSBackendType"></a>

### GetServerInformationReply.FSBackendType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| backend_type | [string](#string) |  |  |






<a name="v1.GetServerInformationReply.Objects"></a>

### GetServerInformationReply.Objects



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| count | [uint64](#uint64) |  |  |






<a name="v1.GetServerInformationReply.ServerProperties"></a>

### GetServerInformationReply.ServerProperties



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| state | [string](#string) |  |  |
| endpoint | [string](#string) |  |  |
| uptime | [int64](#int64) |  |  |
| version | [string](#string) |  |  |
| commit_id | [string](#string) |  |  |
| network | [GetServerInformationReply.ServerProperties.Network](#v1.GetServerInformationReply.ServerProperties.Network) | repeated |  |
| disks | [GetServerInformationReply.ServerProperties.Disk](#v1.GetServerInformationReply.ServerProperties.Disk) | repeated |  |






<a name="v1.GetServerInformationReply.ServerProperties.Disk"></a>

### GetServerInformationReply.ServerProperties.Disk



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| endpoint | [string](#string) |  |  |
| root_disk | [bool](#bool) |  |  |
| drive_path | [string](#string) |  |  |
| healing | [bool](#bool) |  |  |
| state | [string](#string) |  |  |
| uuid | [string](#string) |  |  |
| model | [string](#string) |  |  |
| total_space | [uint64](#uint64) |  |  |
| used_space | [uint64](#uint64) |  |  |
| available_space | [uint64](#uint64) |  |  |
| read_throughput | [float](#float) |  |  |
| write_throughput | [float](#float) |  |  |
| read_latency | [float](#float) |  |  |
| write_latency | [float](#float) |  |  |
| utilization | [float](#float) |  |  |






<a name="v1.GetServerInformationReply.ServerProperties.Network"></a>

### GetServerInformationReply.ServerProperties.Network



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="v1.GetServerInformationReply.Services"></a>

### GetServerInformationReply.Services



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| vault | [GetServerInformationReply.Services.Vault](#v1.GetServerInformationReply.Services.Vault) |  |  |
| ldap | [GetServerInformationReply.Services.LDAP](#v1.GetServerInformationReply.Services.LDAP) |  |  |
| loggers | [GetServerInformationReply.Services.Logger](#v1.GetServerInformationReply.Services.Logger) | repeated |  |
| audits | [GetServerInformationReply.Services.Audit](#v1.GetServerInformationReply.Services.Audit) | repeated |  |
| notifications | [GetServerInformationReply.Services.Notifications](#v1.GetServerInformationReply.Services.Notifications) | repeated |  |






<a name="v1.GetServerInformationReply.Services.Audit"></a>

### GetServerInformationReply.Services.Audit



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| audit_map | [GetServerInformationReply.Services.Audit.AuditMap](#v1.GetServerInformationReply.Services.Audit.AuditMap) | repeated |  |






<a name="v1.GetServerInformationReply.Services.Audit.AuditMap"></a>

### GetServerInformationReply.Services.Audit.AuditMap



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| status | [string](#string) |  |  |






<a name="v1.GetServerInformationReply.Services.LDAP"></a>

### GetServerInformationReply.Services.LDAP



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [string](#string) |  |  |






<a name="v1.GetServerInformationReply.Services.Logger"></a>

### GetServerInformationReply.Services.Logger



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| logger_map | [GetServerInformationReply.Services.Logger.LoggerMap](#v1.GetServerInformationReply.Services.Logger.LoggerMap) | repeated |  |






<a name="v1.GetServerInformationReply.Services.Logger.LoggerMap"></a>

### GetServerInformationReply.Services.Logger.LoggerMap



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| status | [string](#string) |  |  |






<a name="v1.GetServerInformationReply.Services.Notifications"></a>

### GetServerInformationReply.Services.Notifications



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| notification_map | [GetServerInformationReply.Services.Notifications.NotificationMap](#v1.GetServerInformationReply.Services.Notifications.NotificationMap) | repeated |  |






<a name="v1.GetServerInformationReply.Services.Notifications.NotificationMap"></a>

### GetServerInformationReply.Services.Notifications.NotificationMap



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| target_id_status | [GetServerInformationReply.Services.Notifications.NotificationMap.TargetIDStatus](#v1.GetServerInformationReply.Services.Notifications.NotificationMap.TargetIDStatus) | repeated |  |






<a name="v1.GetServerInformationReply.Services.Notifications.NotificationMap.TargetIDStatus"></a>

### GetServerInformationReply.Services.Notifications.NotificationMap.TargetIDStatus



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| target_id_status_map | [GetServerInformationReply.Services.Notifications.NotificationMap.TargetIDStatus.TargetIDStatusMap](#v1.GetServerInformationReply.Services.Notifications.NotificationMap.TargetIDStatus.TargetIDStatusMap) | repeated |  |






<a name="v1.GetServerInformationReply.Services.Notifications.NotificationMap.TargetIDStatus.TargetIDStatusMap"></a>

### GetServerInformationReply.Services.Notifications.NotificationMap.TargetIDStatus.TargetIDStatusMap



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| status | [string](#string) |  |  |






<a name="v1.GetServerInformationReply.Services.Vault"></a>

### GetServerInformationReply.Services.Vault



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [string](#string) |  |  |
| encrypt | [string](#string) |  |  |
| decrypt | [string](#string) |  |  |






<a name="v1.GetServerInformationReply.Usage"></a>

### GetServerInformationReply.Usage



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| size | [uint64](#uint64) |  |  |






<a name="v1.GetVersionMsg"></a>

### GetVersionMsg







<a name="v1.GetVersionReply"></a>

### GetVersionReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | If operation was OK |
| version_information | [GetVersionReply.VersionInformation](#v1.GetVersionReply.VersionInformation) |  | Version Information |






<a name="v1.GetVersionReply.VersionInformation"></a>

### GetVersionReply.VersionInformation



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| git_version | [string](#string) |  | The tag on the git repository |
| git_commit | [string](#string) |  | The hash of the git commit |
| git_tree_state | [string](#string) |  | Whether or not the tree was clean when built |
| build_date | [string](#string) |  | Date of build |
| go_version | [string](#string) |  | Version of go used to compile |
| compiler | [string](#string) |  | Compiler used |
| platform | [string](#string) |  | Platform it was compiled for / running on |





 

 

 


<a name="v1.minioadmin"></a>

### minioadmin


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetVersionInformation | [GetVersionMsg](#v1.GetVersionMsg) | [GetVersionReply](#v1.GetVersionReply) | Will return version information about api server |
| GetServerInformation | [GetServerInformationMsg](#v1.GetServerInformationMsg) | [GetServerInformationReply](#v1.GetServerInformationReply) |  |

 



<a name="api.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api.proto



<a name="v1.Endpoint"></a>

### Endpoint



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| endpoint | [string](#string) |  |  |
| access_key_id | [string](#string) |  |  |
| secret_access_key | [string](#string) |  |  |
| ssl | [bool](#bool) |  |  |






<a name="v1.GetServerInformationMsg"></a>

### GetServerInformationMsg



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| endpoint | [Endpoint](#v1.Endpoint) |  |  |






<a name="v1.GetServerInformationReply"></a>

### GetServerInformationReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| mode | [string](#string) |  |  |
| domain | [string](#string) | repeated |  |
| region | [string](#string) |  |  |
| sqs_arn | [string](#string) | repeated |  |
| deployment_id | [string](#string) |  |  |
| buckets | [GetServerInformationReply.Buckets](#v1.GetServerInformationReply.Buckets) |  |  |
| objects | [GetServerInformationReply.Objects](#v1.GetServerInformationReply.Objects) |  |  |
| usage | [GetServerInformationReply.Usage](#v1.GetServerInformationReply.Usage) |  |  |
| services | [GetServerInformationReply.Services](#v1.GetServerInformationReply.Services) |  |  |
| fs_backend | [GetServerInformationReply.FSBackendType](#v1.GetServerInformationReply.FSBackendType) |  |  |
| erasure_backend | [GetServerInformationReply.ErasureBackendType](#v1.GetServerInformationReply.ErasureBackendType) |  |  |
| servers | [GetServerInformationReply.ServerProperties](#v1.GetServerInformationReply.ServerProperties) | repeated |  |






<a name="v1.GetServerInformationReply.Buckets"></a>

### GetServerInformationReply.Buckets



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| count | [uint64](#uint64) |  |  |






<a name="v1.GetServerInformationReply.ErasureBackendType"></a>

### GetServerInformationReply.ErasureBackendType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| backend_type | [string](#string) |  |  |
| online_disks | [int32](#int32) |  |  |
| offline_disks | [int32](#int32) |  |  |
| standard_sc_data | [int32](#int32) |  |  |
| standard_sc_parity | [int32](#int32) |  |  |
| rrsc_data | [int32](#int32) |  |  |
| rrsc_parity | [int32](#int32) |  |  |






<a name="v1.GetServerInformationReply.FSBackendType"></a>

### GetServerInformationReply.FSBackendType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| backend_type | [string](#string) |  |  |






<a name="v1.GetServerInformationReply.Objects"></a>

### GetServerInformationReply.Objects



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| count | [uint64](#uint64) |  |  |






<a name="v1.GetServerInformationReply.ServerProperties"></a>

### GetServerInformationReply.ServerProperties



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| state | [string](#string) |  |  |
| endpoint | [string](#string) |  |  |
| uptime | [int64](#int64) |  |  |
| version | [string](#string) |  |  |
| commit_id | [string](#string) |  |  |
| network | [GetServerInformationReply.ServerProperties.Network](#v1.GetServerInformationReply.ServerProperties.Network) | repeated |  |
| disks | [GetServerInformationReply.ServerProperties.Disk](#v1.GetServerInformationReply.ServerProperties.Disk) | repeated |  |






<a name="v1.GetServerInformationReply.ServerProperties.Disk"></a>

### GetServerInformationReply.ServerProperties.Disk



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| endpoint | [string](#string) |  |  |
| root_disk | [bool](#bool) |  |  |
| drive_path | [string](#string) |  |  |
| healing | [bool](#bool) |  |  |
| state | [string](#string) |  |  |
| uuid | [string](#string) |  |  |
| model | [string](#string) |  |  |
| total_space | [uint64](#uint64) |  |  |
| used_space | [uint64](#uint64) |  |  |
| available_space | [uint64](#uint64) |  |  |
| read_throughput | [float](#float) |  |  |
| write_throughput | [float](#float) |  |  |
| read_latency | [float](#float) |  |  |
| write_latency | [float](#float) |  |  |
| utilization | [float](#float) |  |  |






<a name="v1.GetServerInformationReply.ServerProperties.Network"></a>

### GetServerInformationReply.ServerProperties.Network



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="v1.GetServerInformationReply.Services"></a>

### GetServerInformationReply.Services



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| vault | [GetServerInformationReply.Services.Vault](#v1.GetServerInformationReply.Services.Vault) |  |  |
| ldap | [GetServerInformationReply.Services.LDAP](#v1.GetServerInformationReply.Services.LDAP) |  |  |
| loggers | [GetServerInformationReply.Services.Logger](#v1.GetServerInformationReply.Services.Logger) | repeated |  |
| audits | [GetServerInformationReply.Services.Audit](#v1.GetServerInformationReply.Services.Audit) | repeated |  |
| notifications | [GetServerInformationReply.Services.Notifications](#v1.GetServerInformationReply.Services.Notifications) | repeated |  |






<a name="v1.GetServerInformationReply.Services.Audit"></a>

### GetServerInformationReply.Services.Audit



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| audit_map | [GetServerInformationReply.Services.Audit.AuditMap](#v1.GetServerInformationReply.Services.Audit.AuditMap) | repeated |  |






<a name="v1.GetServerInformationReply.Services.Audit.AuditMap"></a>

### GetServerInformationReply.Services.Audit.AuditMap



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| status | [string](#string) |  |  |






<a name="v1.GetServerInformationReply.Services.LDAP"></a>

### GetServerInformationReply.Services.LDAP



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [string](#string) |  |  |






<a name="v1.GetServerInformationReply.Services.Logger"></a>

### GetServerInformationReply.Services.Logger



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| logger_map | [GetServerInformationReply.Services.Logger.LoggerMap](#v1.GetServerInformationReply.Services.Logger.LoggerMap) | repeated |  |






<a name="v1.GetServerInformationReply.Services.Logger.LoggerMap"></a>

### GetServerInformationReply.Services.Logger.LoggerMap



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| status | [string](#string) |  |  |






<a name="v1.GetServerInformationReply.Services.Notifications"></a>

### GetServerInformationReply.Services.Notifications



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| notification_map | [GetServerInformationReply.Services.Notifications.NotificationMap](#v1.GetServerInformationReply.Services.Notifications.NotificationMap) | repeated |  |






<a name="v1.GetServerInformationReply.Services.Notifications.NotificationMap"></a>

### GetServerInformationReply.Services.Notifications.NotificationMap



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| target_id_status | [GetServerInformationReply.Services.Notifications.NotificationMap.TargetIDStatus](#v1.GetServerInformationReply.Services.Notifications.NotificationMap.TargetIDStatus) | repeated |  |






<a name="v1.GetServerInformationReply.Services.Notifications.NotificationMap.TargetIDStatus"></a>

### GetServerInformationReply.Services.Notifications.NotificationMap.TargetIDStatus



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| target_id_status_map | [GetServerInformationReply.Services.Notifications.NotificationMap.TargetIDStatus.TargetIDStatusMap](#v1.GetServerInformationReply.Services.Notifications.NotificationMap.TargetIDStatus.TargetIDStatusMap) | repeated |  |






<a name="v1.GetServerInformationReply.Services.Notifications.NotificationMap.TargetIDStatus.TargetIDStatusMap"></a>

### GetServerInformationReply.Services.Notifications.NotificationMap.TargetIDStatus.TargetIDStatusMap



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| status | [string](#string) |  |  |






<a name="v1.GetServerInformationReply.Services.Vault"></a>

### GetServerInformationReply.Services.Vault



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [string](#string) |  |  |
| encrypt | [string](#string) |  |  |
| decrypt | [string](#string) |  |  |






<a name="v1.GetServerInformationReply.Usage"></a>

### GetServerInformationReply.Usage



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| size | [uint64](#uint64) |  |  |






<a name="v1.GetVersionMsg"></a>

### GetVersionMsg







<a name="v1.GetVersionReply"></a>

### GetVersionReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | If operation was OK |
| version_information | [GetVersionReply.VersionInformation](#v1.GetVersionReply.VersionInformation) |  | Version Information |






<a name="v1.GetVersionReply.VersionInformation"></a>

### GetVersionReply.VersionInformation



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| git_version | [string](#string) |  | The tag on the git repository |
| git_commit | [string](#string) |  | The hash of the git commit |
| git_tree_state | [string](#string) |  | Whether or not the tree was clean when built |
| build_date | [string](#string) |  | Date of build |
| go_version | [string](#string) |  | Version of go used to compile |
| compiler | [string](#string) |  | Compiler used |
| platform | [string](#string) |  | Platform it was compiled for / running on |





 

 

 


<a name="v1.minioadmin"></a>

### minioadmin


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetVersionInformation | [GetVersionMsg](#v1.GetVersionMsg) | [GetVersionReply](#v1.GetVersionReply) | Will return version information about api server |
| GetServerInformation | [GetServerInformationMsg](#v1.GetServerInformationMsg) | [GetServerInformationReply](#v1.GetServerInformationReply) |  |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

