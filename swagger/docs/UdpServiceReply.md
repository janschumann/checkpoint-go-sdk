# UdpServiceReply

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptReplies** | **bool** | N/A | [optional] [default to null]
**AggressiveAging** | [***AggressiveAgingReply**](AggressiveAgingReply.md) |  | [optional] [default to null]
**Color** | **string** | Color of the object. Should be one of existing colors. | [optional] [default to null]
**Comments** | **string** | Comments string. | [optional] [default to null]
**Domain** | [***ApiDomainIdentifier**](ApiDomainIdentifier.md) |  | [optional] [default to null]
**Groups** | [**[]ApiObjectStandardIdentifier**](ApiObjectStandardIdentifier.md) | How much details are returned depends on the details-level field of the request. This table shows the level of detail shown when details-level is set to standard. | [optional] [default to null]
**Icon** | **string** | Object icon. | [optional] [default to null]
**KeepConnectionsOpenAfterPolicyInstallation** | **bool** | Keep connections open after policy has been installed even if they are not allowed under the new policy. This overrides the settings in the Connection Persistence page. If you change this property, the change will not affect open connections, but only future connections. | [optional] [default to null]
**MatchByProtocolSignature** | **bool** | A value of true enables matching by the selected protocol&#39;s signature - The signature identifies the protocol as genuine. | [optional] [default to null]
**MatchForAny** | **bool** | Indicates whether this service is used when &#39;Any&#39; is set as the rule&#39;s service and there are several service objects with the same source port and protocol. | [optional] [default to null]
**MetaInfo** | [***MetaInfoForTopLevelReply**](MetaInfoForTopLevelReply.md) |  | [optional] [default to null]
**Name** | **string** | Object name. Should be unique in the domain. | [optional] [default to null]
**OverrideDefaultSettings** | **bool** | Indicates whether this service is a Data Domain service which has been overridden. | [optional] [default to null]
**Port** | **string** | The number of the port used to provide this service. | [optional] [default to null]
**Protocol** | **string** | The protocol type associated with the service, and by implication, the management server (if any) that enforces Content Security and Authentication for the service. | [optional] [default to null]
**ReadOnly** | **bool** | Indicates whether the object is read-only. | [optional] [default to null]
**SessionTimeout** | **int32** | Time (in seconds) before the session times out. | [optional] [default to null]
**SourcePort** | **string** | Port number for the client side service. If specified, only those Source port Numbers will be Accepted, Dropped, or Rejected during packet inspection. Otherwise, the source port is not inspected. | [optional] [default to null]
**SyncConnectionsOnCluster** | **bool** | Enables state-synchronized High Availability or Load Sharing on a ClusterXL or OPSEC-certified cluster. | [optional] [default to null]
**Tags** | [**[]ApiObjectStandardIdentifier**](ApiObjectStandardIdentifier.md) | Collection of tag objects identified by the name or UID. How much details are returned depends on the details-level field of the request. This table shows the level of detail shown when details-level is set to standard. | [optional] [default to null]
**Type_** | **string** | Type of the object. | [optional] [default to null]
**Uid** | **string** | Object unique identifier. | [optional] [default to null]
**UseDefaultSessionTimeout** | **bool** | Use default virtual session timeout. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


