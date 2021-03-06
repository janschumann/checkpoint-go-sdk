# AdministratorReply

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationMethod** | **string** | Authentication method. | [optional] [default to null]
**Color** | **string** | Color of the object. Should be one of existing colors. | [optional] [default to null]
**Comments** | **string** | Comments string. | [optional] [default to null]
**Domain** | [***ApiDomainIdentifier**](ApiDomainIdentifier.md) |  | [optional] [default to null]
**Email** | **string** | Administrator email. | [optional] [default to null]
**ExpirationDate** | [***ApiDateReply**](ApiDateReply.md) |  | [optional] [default to null]
**Icon** | **string** | Object icon. | [optional] [default to null]
**MetaInfo** | [***MetaInfoForTopLevelReply**](MetaInfoForTopLevelReply.md) |  | [optional] [default to null]
**MultiDomainProfile** | [**[]ApiObjectStandardIdentifier**](ApiObjectStandardIdentifier.md) | Administrator multi-domain profile. How much details are returned depends on the details-level field of the request. This table shows the level of detail shown when details-level is set to standard. | [optional] [default to null]
**MustChangePassword** | **bool** | True if administrator must change password on the next login. | [optional] [default to null]
**Name** | **string** | Object name. Should be unique in the domain. | [optional] [default to null]
**PermissionsProfile** | [**[]ApiObjectStandardIdentifier**](ApiObjectStandardIdentifier.md) | Administrator permissions profile. How much details are returned depends on the details-level field of the request. This table shows the level of detail shown when details-level is set to standard. | [optional] [default to null]
**PhoneNumber** | **string** | Administrator phone number. | [optional] [default to null]
**RadiusServer** | [***ApiObjectStandardIdentifier**](ApiObjectStandardIdentifier.md) |  | [optional] [default to null]
**ReadOnly** | **bool** | Indicates whether the object is read-only. | [optional] [default to null]
**SicName** | **string** | Name of the Secure Internal Connection Trust. | [optional] [default to null]
**TacacsServer** | [***ApiObjectStandardIdentifier**](ApiObjectStandardIdentifier.md) |  | [optional] [default to null]
**Tags** | [**[]ApiObjectStandardIdentifier**](ApiObjectStandardIdentifier.md) | Collection of tag objects identified by the name or UID. How much details are returned depends on the details-level field of the request. This table shows the level of detail shown when details-level is set to standard. | [optional] [default to null]
**Type_** | **string** | Type of the object. | [optional] [default to null]
**Uid** | **string** | Object unique identifier. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


