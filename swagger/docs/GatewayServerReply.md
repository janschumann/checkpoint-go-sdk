# GatewayServerReply

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterMemberNames** | **[]string** | Names of cluster members. | [optional] [default to null]
**Color** | **string** | Color of the object. Should be one of existing colors. | [optional] [default to null]
**Comments** | **string** | Comments string. | [optional] [default to null]
**Domain** | [***ApiDomainIdentifier**](ApiDomainIdentifier.md) |  | [optional] [default to null]
**Groups** | [**[]ApiObjectStandardIdentifier**](ApiObjectStandardIdentifier.md) | How much details are returned depends on the details-level field of the request. This table shows the level of detail shown when details-level is set to standard. | [optional] [default to null]
**Hardware** | **string** | Appliance type. | [optional] [default to null]
**Icon** | **string** | Object icon. | [optional] [default to null]
**Interfaces** | [**[]GatewayServerInterfaceReply**](GatewayServerInterfaceReply.md) | Network interfaces. | [optional] [default to null]
**Ipv4Address** | **string** | IPv4 address. | [optional] [default to null]
**Ipv6Address** | **string** | IPv6 address. | [optional] [default to null]
**ManagementBlades** | [***GatewayServerManagementBladesReply**](GatewayServerManagementBladesReply.md) |  | [optional] [default to null]
**MetaInfo** | [***MetaInfoForTopLevelReply**](MetaInfoForTopLevelReply.md) |  | [optional] [default to null]
**Name** | **string** | Object name. Should be unique in the domain. | [optional] [default to null]
**NetworkSecurityBlades** | [***GatewayServerNetworkBladesReply**](GatewayServerNetworkBladesReply.md) |  | [optional] [default to null]
**OperatingSystem** | **string** | Operating System. | [optional] [default to null]
**Policy** | [***GatewayServerPolicyReply**](GatewayServerPolicyReply.md) |  | [optional] [default to null]
**ReadOnly** | **bool** | Indicates whether the object is read-only. | [optional] [default to null]
**SicStatus** | **string** | Secure Internal Communication status. | [optional] [default to null]
**Tags** | **[]string** | Collection of tag objects identified by the name or UID. | [optional] [default to null]
**Type_** | **string** | Object type. | [optional] [default to null]
**Uid** | **string** | Object unique identifier. | [optional] [default to null]
**Version** | **string** | Version. | [optional] [default to null]
**VpnEncryptionDomain** | **string** | VPN domain. | [optional] [default to null]
**VpnEncryptionDomainManuallyDefined** | [***ApiObjectStandardIdentifier**](ApiObjectStandardIdentifier.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


