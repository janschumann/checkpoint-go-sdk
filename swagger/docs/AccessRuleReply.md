# AccessRuleReply

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | [***ApiObjectStandardIdentifier**](ApiObjectStandardIdentifier.md) |  | [optional] [default to null]
**ActionSettings** | [***AdvancedActionSettingsReply**](AdvancedActionSettingsReply.md) |  | [optional] [default to null]
**Comments** | **string** | Comments string. | [optional] [default to null]
**Content** | [**[]ApiObjectStandardIdentifier**](ApiObjectStandardIdentifier.md) | How much details are returned depends on the details-level field of the request. This table shows the level of detail shown when details-level is set to standard. | [optional] [default to null]
**ContentDirection** | **string** | On which direction the file types processing is applied. | [optional] [default to null]
**ContentNegate** | **bool** | True if negate is set for data. | [optional] [default to null]
**CustomFields** | [***CustomSummaryFieldsReply**](CustomSummaryFieldsReply.md) |  | [optional] [default to null]
**Destination** | [**[]ApiObjectStandardIdentifier**](ApiObjectStandardIdentifier.md) | Collection of Network objects identified by the name or UID. How much details are returned depends on the details-level field of the request. This table shows the level of detail shown when details-level is set to standard. | [optional] [default to null]
**DestinationNegate** | **bool** | True if negate is set for destination. | [optional] [default to null]
**DestinationRanges** | [***IpRanges**](IpRanges.md) |  | [optional] [default to null]
**Domain** | [***ApiDomainIdentifier**](ApiDomainIdentifier.md) |  | [optional] [default to null]
**Enabled** | **bool** | Enable/Disable the rule. | [optional] [default to null]
**Hits** | [***HitsReply**](HitsReply.md) |  | [optional] [default to null]
**InlineLayer** | [***ApiObjectStandardIdentifier**](ApiObjectStandardIdentifier.md) |  | [optional] [default to null]
**InstallOn** | [**[]ApiObjectStandardIdentifier**](ApiObjectStandardIdentifier.md) | Which gateway, identified by the name or UID, to install the policy. How much details are returned depends on the details-level field of the request. This table shows the level of detail shown when details-level is set to standard. | [optional] [default to null]
**Layer** | **string** | N/A | [optional] [default to null]
**MetaInfo** | [***MetaInfoForTopLevelReply**](MetaInfoForTopLevelReply.md) |  | [optional] [default to null]
**Name** | **string** | Object name. Should be unique in the domain. | [optional] [default to null]
**Service** | [**[]ApiObjectStandardIdentifier**](ApiObjectStandardIdentifier.md) | Collection of Network objects identified by the name or UID. How much details are returned depends on the details-level field of the request. This table shows the level of detail shown when details-level is set to standard. | [optional] [default to null]
**ServiceNegate** | **bool** | True if negate is set for service. | [optional] [default to null]
**ServiceRanges** | [***PortRanges**](PortRanges.md) |  | [optional] [default to null]
**Source** | [**[]ApiObjectStandardIdentifier**](ApiObjectStandardIdentifier.md) | Collection of Network objects identified by the name or UID. How much details are returned depends on the details-level field of the request. This table shows the level of detail shown when details-level is set to standard. | [optional] [default to null]
**SourceNegate** | **bool** | True if negate is set for source. | [optional] [default to null]
**SourceRanges** | [***IpRanges**](IpRanges.md) |  | [optional] [default to null]
**Time** | [**[]ApiObjectStandardIdentifier**](ApiObjectStandardIdentifier.md) | List of time objects. For example: \&quot;Weekend\&quot;, \&quot;Off-Work\&quot;, \&quot;Every-Day\&quot;. How much details are returned depends on the details-level field of the request. This table shows the level of detail shown when details-level is set to standard. | [optional] [default to null]
**Track** | [***interface{}**](interface{}.md) | Track Settings. | [optional] [default to null]
**Type_** | **string** | Type of the object. | [optional] [default to null]
**Uid** | **string** | Object unique identifier. | [optional] [default to null]
**UserCheck** | [***UserCheckReply**](UserCheckReply.md) |  | [optional] [default to null]
**Vpn** | [**[]ApiObjectStandardIdentifier**](ApiObjectStandardIdentifier.md) | VPN settings. How much details are returned depends on the details-level field of the request. This table shows the level of detail shown when details-level is set to standard. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


