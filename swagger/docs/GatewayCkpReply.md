# GatewayCkpReply

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AntiBot** | **bool** | Anti-Bot blade enabled. | [optional] [default to null]
**AntiVirus** | **bool** | Anti-Virus blade enabled. | [optional] [default to null]
**ApplicationControl** | **bool** | Application Control blade enabled. | [optional] [default to null]
**Color** | **string** | Color of the object. Should be one of existing colors. | [optional] [default to null]
**Comments** | **string** | Comments string. | [optional] [default to null]
**ContentAwareness** | **bool** | Content Awareness blade enabled. | [optional] [default to null]
**Domain** | [***ApiDomainIdentifier**](ApiDomainIdentifier.md) |  | [optional] [default to null]
**DynamicIp** | **bool** | Dynamic IP address. | [optional] [default to null]
**Firewall** | **bool** | Firewall blade enabled. | [optional] [default to null]
**FirewallSettings** | [***FirewallSettingsReply**](FirewallSettingsReply.md) |  | [optional] [default to null]
**Groups** | [**[]ApiObjectStandardIdentifier**](ApiObjectStandardIdentifier.md) | How much details are returned depends on the details-level field of the request. This table shows the level of detail shown when details-level is set to standard. | [optional] [default to null]
**Hardware** | **string** | Gateway platform hardware type. | [optional] [default to null]
**Icon** | **string** | Object icon. | [optional] [default to null]
**Interfaces** | [**[]NetworkInterfaceReply**](NetworkInterfaceReply.md) | Network interfaces. | [optional] [default to null]
**Ips** | **bool** | Intrusion Prevention System blade enabled. | [optional] [default to null]
**Ipv4Address** | **string** | IPv4 address. | [optional] [default to null]
**Ipv6Address** | **string** | IPv6 address. | [optional] [default to null]
**LogsSettings** | [***LogsSettingsReply**](LogsSettingsReply.md) |  | [optional] [default to null]
**MetaInfo** | [***MetaInfoForTopLevelReply**](MetaInfoForTopLevelReply.md) |  | [optional] [default to null]
**Name** | **string** | Object name. Should be unique in the domain. | [optional] [default to null]
**OsName** | **string** | Gateway platform operating system. | [optional] [default to null]
**ReadOnly** | **bool** | Indicates whether the object is read-only. | [optional] [default to null]
**SaveLogsLocally** | **bool** | Save logs locally on the gateway. | [optional] [default to null]
**SendAlertsToServer** | **[]string** | Server(s) to send alerts to. | [optional] [default to null]
**SendLogsToBackupServer** | **[]string** | Backup server(s) to send logs to. | [optional] [default to null]
**SendLogsToServer** | **[]string** | Servers(s) to send logs to. | [optional] [default to null]
**SicName** | **string** | Secure Internal Communication name. | [optional] [default to null]
**SicState** | **string** | Secure Internal Communication state. | [optional] [default to null]
**Tags** | [**[]ApiObjectStandardIdentifier**](ApiObjectStandardIdentifier.md) | Collection of tag objects identified by the name or UID. How much details are returned depends on the details-level field of the request. This table shows the level of detail shown when details-level is set to standard. | [optional] [default to null]
**ThreatEmulation** | **bool** | Threat Emulation blade enabled. | [optional] [default to null]
**ThreatExtraction** | **bool** | Threat Extraction blade enabled. | [optional] [default to null]
**Type_** | **string** | Type of the object. | [optional] [default to null]
**Uid** | **string** | Object unique identifier. | [optional] [default to null]
**UrlFiltering** | **bool** | URL Filtering blade enabled. | [optional] [default to null]
**Version** | **string** | Gateway platform version. | [optional] [default to null]
**Vpn** | **bool** | VPN blade enabled. | [optional] [default to null]
**VpnSettings** | [***VpnSettingsReply**](VpnSettingsReply.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


