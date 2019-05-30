# GatewayServerPolicyReply

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessPolicyInstallationDate** | [***ApiDateReply**](ApiDateReply.md) |  | [optional] [default to null]
**AccessPolicyInstalled** | **bool** | Gets true if access-policy was installed. | [optional] [default to null]
**AccessPolicyName** | **string** | Name of the access-policy. | [optional] [default to null]
**AccessPolicyRevision** | [***ApiObjectStandardIdentifier**](ApiObjectStandardIdentifier.md) |  | [optional] [default to null]
**ClusterMembersAccessPolicyRevision** | [**[]GatewayServerPolicyReplyClusterMemberReply**](GatewayServerPolicyReplyClusterMemberReply.md) | Revisions of the access policy installed on each cluster member. | [optional] [default to null]
**ClusterMembersThreatPolicyRevision** | [**[]GatewayServerPolicyReplyClusterMemberReply**](GatewayServerPolicyReplyClusterMemberReply.md) | Revisions of the threat policy installed on each cluster member. | [optional] [default to null]
**ThreatPolicyInstallationDate** | [***ApiDateReply**](ApiDateReply.md) |  | [optional] [default to null]
**ThreatPolicyInstalled** | **bool** | Gets true if threat-policy was installed. | [optional] [default to null]
**ThreatPolicyName** | **string** | Name of the threat-policy. | [optional] [default to null]
**ThreatPolicyRevision** | [***ApiObjectStandardIdentifier**](ApiObjectStandardIdentifier.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


