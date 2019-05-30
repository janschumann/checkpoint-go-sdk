# ThreatRuleReply

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | [***ApiObjectStandardIdentifier**](ApiObjectStandardIdentifier.md) |  | [optional] [default to null]
**Comments** | **string** | Comments string. | [optional] [default to null]
**Destination** | [**[]ApiObjectStandardIdentifier**](ApiObjectStandardIdentifier.md) | Collection of Network objects identified by the name or UID. How much details are returned depends on the details-level field of the request. This table shows the level of detail shown when details-level is set to standard. | [optional] [default to null]
**DestinationNegate** | **bool** | True if negate is set for destination. | [optional] [default to null]
**Domain** | [***ApiDomainIdentifier**](ApiDomainIdentifier.md) |  | [optional] [default to null]
**Enabled** | **bool** | Enable/Disable the rule. | [optional] [default to null]
**Exceptions** | [***interface{}**](interface{}.md) | The rule&#39;s exceptions. | [optional] [default to null]
**InstallOn** | [**[]ApiObjectStandardIdentifier**](ApiObjectStandardIdentifier.md) | Which gateway, identified by the name or UID, to install the policy. How much details are returned depends on the details-level field of the request. This table shows the level of detail shown when details-level is set to standard. | [optional] [default to null]
**Layer** | **string** | N/A | [optional] [default to null]
**MetaInfo** | [***MetaInfoForTopLevelReply**](MetaInfoForTopLevelReply.md) |  | [optional] [default to null]
**Name** | **string** | Object name. Should be unique in the domain. | [optional] [default to null]
**ProtectedScope** | [**[]ApiObjectStandardIdentifier**](ApiObjectStandardIdentifier.md) | Collection of network objects defining Protection Scope identified by the name or UID. How much details are returned depends on the details-level field of the request. This table shows the level of detail shown when details-level is set to standard. | [optional] [default to null]
**ProtectedScopeNegate** | **bool** | True if negate is set for Protected Scope. | [optional] [default to null]
**Service** | [**[]ApiObjectStandardIdentifier**](ApiObjectStandardIdentifier.md) | Collection of network objects defining Service identified by the name or UID. How much details are returned depends on the details-level field of the request. This table shows the level of detail shown when details-level is set to standard. | [optional] [default to null]
**ServiceNegate** | **bool** | True if negate is set for Service. | [optional] [default to null]
**Source** | [**[]ApiObjectStandardIdentifier**](ApiObjectStandardIdentifier.md) | Collection of network objects defining Source identified by the name or UID. How much details are returned depends on the details-level field of the request. This table shows the level of detail shown when details-level is set to standard. | [optional] [default to null]
**SourceNegate** | **bool** | True if negate is set for source. | [optional] [default to null]
**Track** | [***ApiObjectStandardIdentifier**](ApiObjectStandardIdentifier.md) |  | [optional] [default to null]
**TrackSettings** | [***ThreatRuleTrackSettingsReply**](ThreatRuleTrackSettingsReply.md) |  | [optional] [default to null]
**Type_** | **string** | Type of the object. | [optional] [default to null]
**Uid** | **string** | Object unique identifier. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


