# PolicyPackageReply

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Access** | **bool** | True - enables, False - disables access policy, empty - nothing is changed. | [optional] [default to null]
**AccessLayers** | [**[]ApiObjectStandardIdentifier**](ApiObjectStandardIdentifier.md) | Access policy layers. How much details are returned depends on the details-level field of the request. This table shows the level of detail shown when details-level is set to standard. | [optional] [default to null]
**Color** | **string** | Color of the object. Should be one of existing colors. | [optional] [default to null]
**Comments** | **string** | Comments string. | [optional] [default to null]
**DesktopSecurity** | **bool** | True - enables, False - disables Desktop security policy, empty - nothing is changed. | [optional] [default to null]
**Domain** | [***ApiDomainIdentifier**](ApiDomainIdentifier.md) |  | [optional] [default to null]
**Icon** | **string** | Object icon. | [optional] [default to null]
**InstallationTargets** | [***ApiObjectStandardIdentifier**](ApiObjectStandardIdentifier.md) |  | [optional] [default to null]
**InstallationTargetsRevision** | [**[]InstallationTargetRevisionFullReply**](InstallationTargetRevisionFullReply.md) | List of installation targets and revisions on which this policy package was installed. | [optional] [default to null]
**MetaInfo** | [***MetaInfoForTopLevelReply**](MetaInfoForTopLevelReply.md) |  | [optional] [default to null]
**Name** | **string** | Object name. Should be unique in the domain. | [optional] [default to null]
**NatPolicy** | **bool** | True - enables, False - disables NAT policy, empty - nothing is changed. | [optional] [default to null]
**Qos** | **bool** | True - enables, False - disables QoS policy, empty - nothing is changed. | [optional] [default to null]
**QosPolicyType** | **string** | QoS policy type. | [optional] [default to null]
**ReadOnly** | **bool** | Indicates whether the object is read-only. | [optional] [default to null]
**Tags** | [**[]ApiObjectStandardIdentifier**](ApiObjectStandardIdentifier.md) | Collection of tag objects identified by the name or UID. How much details are returned depends on the details-level field of the request. This table shows the level of detail shown when details-level is set to standard. | [optional] [default to null]
**ThreatLayers** | [**[]ApiObjectStandardIdentifier**](ApiObjectStandardIdentifier.md) | Threat policy layers. How much details are returned depends on the details-level field of the request. This table shows the level of detail shown when details-level is set to standard. | [optional] [default to null]
**ThreatPrevention** | **bool** | True - enables, False - disables Threat policy, empty - nothing is changed. | [optional] [default to null]
**Type_** | **string** | Type of the object. | [optional] [default to null]
**Uid** | **string** | Object unique identifier. | [optional] [default to null]
**VpnTraditionalMode** | **bool** | True - enables, False - disables VPN traditional mode, empty - nothing is changed. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


