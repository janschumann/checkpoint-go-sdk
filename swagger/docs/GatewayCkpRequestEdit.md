# GatewayCkpRequestEdit

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AntiBot** | **bool** | Anti-Bot blade enabled. | [optional] [default to null]
**AntiVirus** | **bool** | Anti-Virus blade enabled. | [optional] [default to null]
**ApplicationControl** | **bool** | Application Control blade enabled. | [optional] [default to null]
**Color** | **string** | Color of the object. Should be one of existing colors. | [optional] [default to null]
**Comments** | **string** | Comments string. | [optional] [default to null]
**ContentAwareness** | **bool** | Content Awareness blade enabled. | [optional] [default to null]
**DetailsLevel** | **string** | The level of detail for some of the fields in the response can vary from showing only the UID value of the object to a fully detailed representation of the object. | [optional] [default to null]
**Firewall** | **bool** | Firewall blade enabled. | [optional] [default to null]
**FirewallSettings** | [***FirewallSettingsRequestEdit**](FirewallSettingsRequestEdit.md) |  | [optional] [default to null]
**Groups** | [***Add**](add.md) |  | [optional] [default to null]
**IgnoreWarnings** | **bool** | Apply changes ignoring warnings. | [optional] [default to null]
**Interfaces** | [**[]NetworkInterfaceRequestEdit**](NetworkInterfaceRequestEdit.md) | Network interfaces. When a gateway is updated with a new interfaces, the existing interfaces are removed. | [optional] [default to null]
**IpAddress** | **string** | IPv4 or IPv6 address. If both addresses are required use ipv4-address and ipv6-address fields explicitly. | [optional] [default to null]
**Ips** | **bool** | Intrusion Prevention System blade enabled. | [optional] [default to null]
**Ipv4Address** | **string** | IPv4 address. | [optional] [default to null]
**Ipv6Address** | **string** | IPv6 address. | [optional] [default to null]
**LogsSettings** | [***LogsSettingsRequest**](LogsSettingsRequest.md) |  | [optional] [default to null]
**Name** | **string** | Object name. | [optional] [default to null]
**NewName** | **string** | New name of the object. | [optional] [default to null]
**OneTimePassword** | **string** | N/A | [optional] [default to null]
**OsName** | **string** | Gateway platform operating system. | [optional] [default to null]
**SaveLogsLocally** | **bool** | Save logs locally on the gateway. | [optional] [default to null]
**SendAlertsToServer** | [***Add**](add.md) |  | [optional] [default to null]
**SendLogsToBackupServer** | [***Add**](add.md) |  | [optional] [default to null]
**SendLogsToServer** | [***Add**](add.md) |  | [optional] [default to null]
**Tags** | [***Add**](add.md) |  | [optional] [default to null]
**ThreatEmulation** | **bool** | Threat Emulation blade enabled. | [optional] [default to null]
**ThreatExtraction** | **bool** | Threat Extraction blade enabled. | [optional] [default to null]
**Uid** | **string** | Object unique identifier. | [default to null]
**UrlFiltering** | **bool** | URL Filtering blade enabled. | [optional] [default to null]
**Version** | **string** | Gateway platform version. | [optional] [default to null]
**Vpn** | **bool** | VPN blade enabled. | [optional] [default to null]
**VpnSettings** | [***VpnSettingsRequest**](VpnSettingsRequest.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


