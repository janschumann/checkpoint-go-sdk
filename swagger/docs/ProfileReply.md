# ProfileReply

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveProtectionsPerformanceImpact** | **string** | Protections with this performance impact only will be activated in the profile. | [optional] [default to null]
**ActiveProtectionsSeverity** | **string** | Protections with this severity only will be activated in the profile. | [optional] [default to null]
**AntiBot** | **bool** | Is Anti-Bot blade activated. | [optional] [default to null]
**AntiVirus** | **bool** | Is Anti-Virus blade activated. | [optional] [default to null]
**Color** | **string** | Color of the object. Should be one of existing colors. | [optional] [default to null]
**Comments** | **string** | Comments string. | [optional] [default to null]
**ConfidenceLevelHigh** | **string** | Action for protections with high confidence level. | [optional] [default to null]
**ConfidenceLevelLow** | **string** | Action for protections with low confidence level. | [optional] [default to null]
**ConfidenceLevelMedium** | **string** | Action for protections with medium confidence level. | [optional] [default to null]
**Domain** | [***ApiDomainIdentifier**](ApiDomainIdentifier.md) |  | [optional] [default to null]
**ExtendedAttributesToActivate** | [**[]IpsAdditionalPropertiesReply**](IpsAdditionalPropertiesReply.md) | Activate protections by these extended attributes. | [optional] [default to null]
**ExtendedAttributesToDeactivate** | [**[]IpsAdditionalPropertiesReply**](IpsAdditionalPropertiesReply.md) | Deactivate protections by these extended attributes. | [optional] [default to null]
**Icon** | **string** | Object icon. | [optional] [default to null]
**IndicatorOverrides** | [**[]ProfileIndicatorOverrideRequest**](ProfileIndicatorOverrideRequest.md) | Indicators whose action will be overridden in this profile. | [optional] [default to null]
**Ips** | **bool** | Is IPS blade activated. | [optional] [default to null]
**IpsSettings** | [***IpsSettingsReply**](IpsSettingsReply.md) |  | [optional] [default to null]
**MaliciousMailPolicySettings** | [***MailSettingsReply**](MailSettingsReply.md) |  | [optional] [default to null]
**MetaInfo** | [***MetaInfoForTopLevelReply**](MetaInfoForTopLevelReply.md) |  | [optional] [default to null]
**Name** | **string** | Object name. Should be unique in the domain. | [optional] [default to null]
**Overrides** | [**[]AllActivationsByProtectionReply**](AllActivationsByProtectionReply.md) | Overrides per profile for this protection. | [optional] [default to null]
**ReadOnly** | **bool** | Indicates whether the object is read-only. | [optional] [default to null]
**Tags** | [**[]ApiObjectStandardIdentifier**](ApiObjectStandardIdentifier.md) | Collection of tag objects identified by the name or UID. How much details are returned depends on the details-level field of the request. This table shows the level of detail shown when details-level is set to standard. | [optional] [default to null]
**ThreatEmulation** | **bool** | Is Threat Emulation blade activated. | [optional] [default to null]
**Type_** | **string** | Type of the object. | [optional] [default to null]
**Uid** | **string** | Object unique identifier. | [optional] [default to null]
**UseExtendedAttributes** | **bool** | Whether to activate/deactivate IPS protections according to the extended attributes. | [optional] [default to null]
**UseIndicators** | **bool** | Indicates whether the profile should make use of indicators. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


