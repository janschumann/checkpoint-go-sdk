# ProfileRequestEdit

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivateProtectionsByExtendedAttributes** | [***Add**](add.md) |  | [optional] [default to null]
**ActiveProtectionsPerformanceImpact** | **string** | Protections with this performance impact only will be activated in the profile. | [optional] [default to null]
**ActiveProtectionsSeverity** | **string** | Protections with this severity only will be activated in the profile. | [optional] [default to null]
**AntiBot** | **bool** | Is Anti-Bot blade activated. | [optional] [default to null]
**AntiVirus** | **bool** | Is Anti-Virus blade activated. | [optional] [default to null]
**Color** | **string** | Color of the object. Should be one of existing colors. | [optional] [default to null]
**Comments** | **string** | Comments string. | [optional] [default to null]
**ConfidenceLevelHigh** | **string** | Action for protections with high confidence level. | [optional] [default to null]
**ConfidenceLevelLow** | **string** | Action for protections with low confidence level. | [optional] [default to null]
**ConfidenceLevelMedium** | **string** | Action for protections with medium confidence level. | [optional] [default to null]
**DeactivateProtectionsByExtendedAttributes** | [***Add**](add.md) |  | [optional] [default to null]
**DetailsLevel** | **string** | The level of detail for some of the fields in the response can vary from showing only the UID value of the object to a fully detailed representation of the object. | [optional] [default to null]
**IgnoreErrors** | **bool** | Apply changes ignoring errors. You won&#39;t be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. | [optional] [default to null]
**IgnoreWarnings** | **bool** | Apply changes ignoring warnings. | [optional] [default to null]
**IndicatorOverrides** | [***Add**](add.md) |  | [optional] [default to null]
**Ips** | **bool** | Is IPS blade activated. | [optional] [default to null]
**IpsSettings** | [***IpsSettingsRequest**](IpsSettingsRequest.md) |  | [optional] [default to null]
**MaliciousMailPolicySettings** | [***MailSettingsRequestEdit**](MailSettingsRequestEdit.md) |  | [optional] [default to null]
**Name** | **string** | Object name. | [optional] [default to null]
**NewName** | **string** | New name of the object. | [optional] [default to null]
**Overrides** | [***Add**](add.md) |  | [optional] [default to null]
**Tags** | [***Add**](add.md) |  | [optional] [default to null]
**ThreatEmulation** | **bool** | Is Threat Emulation blade activated. | [optional] [default to null]
**Uid** | **string** | Object unique identifier. | [default to null]
**UseExtendedAttributes** | **bool** | Whether to activate/deactivate IPS protections according to the extended attributes. | [optional] [default to null]
**UseIndicators** | **bool** | Indicates whether the profile should make use of indicators. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


