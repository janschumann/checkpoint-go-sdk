# ProtectionReply

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comments** | **string** | Protection comments. | [optional] [default to null]
**ConfidenceLevel** | **string** | How confident IPS is that recognized attacks are actually undesirable traffic. | [optional] [default to null]
**Domain** | [***ApiDomainIdentifier**](ApiDomainIdentifier.md) |  | [optional] [default to null]
**FollowUp** | **bool** | Tag the protection with pre-defined follow-up flag. | [optional] [default to null]
**IndustryReference** | **[]string** | International CVE or CVE candidate name for attack. | [optional] [default to null]
**IpsAdditionalProperties** | [**[]IpsAdditionalPropertiesReply**](IpsAdditionalPropertiesReply.md) | IPS protection extended attributes. | [optional] [default to null]
**Name** | **string** | Object name. Should be unique in the domain. | [optional] [default to null]
**PerformanceImpact** | **string** | How much this protection affects the performance of a Security Gateway. | [optional] [default to null]
**Profiles** | [**[]AllActivationsByProfileReply**](AllActivationsByProfileReply.md) | Protection settings per profile. | [optional] [default to null]
**ProtectionType** | **string** | The protection&#39;s type (Core or Threat Cloud). | [optional] [default to null]
**ReleaseDate** | **string** | The date in which the protection was releases by Check Point. | [optional] [default to null]
**Severity** | **string** | The intrusion severity. | [optional] [default to null]
**Type_** | **string** | Type of the object. | [optional] [default to null]
**Uid** | **string** | Object unique identifier. | [optional] [default to null]
**UpdateDate** | **string** | The date in which the protection was updated by Check Point. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


