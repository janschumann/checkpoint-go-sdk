# VpnMeshedCommunityReply

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Color** | **string** | Color of the object. Should be one of existing colors. | [optional] [default to null]
**Comments** | **string** | Comments string. | [optional] [default to null]
**Domain** | [***ApiDomainIdentifier**](ApiDomainIdentifier.md) |  | [optional] [default to null]
**EncryptionMethod** | **string** | The encryption method to be used. | [optional] [default to null]
**EncryptionSuite** | **string** | The encryption suite to be used. | [optional] [default to null]
**Gateways** | [**[]ApiObjectStandardIdentifier**](ApiObjectStandardIdentifier.md) | Collection of Gateway objects identified by the name or UID. How much details are returned depends on the details-level field of the request. This table shows the level of detail shown when details-level is set to standard. | [optional] [default to null]
**Icon** | **string** | Object icon. | [optional] [default to null]
**IkePhase1** | [***IkeP1Reply**](IkeP1Reply.md) |  | [optional] [default to null]
**IkePhase2** | [***IkeP2Reply**](IkeP2Reply.md) |  | [optional] [default to null]
**MetaInfo** | [***MetaInfoForTopLevelReply**](MetaInfoForTopLevelReply.md) |  | [optional] [default to null]
**Name** | **string** | Object name. Should be unique in the domain. | [optional] [default to null]
**ReadOnly** | **bool** | Indicates whether the object is read-only. | [optional] [default to null]
**SharedSecrets** | [**[]SharedSecretReply**](SharedSecretReply.md) | Shared secrets for external gateways. | [optional] [default to null]
**Tags** | [**[]ApiObjectStandardIdentifier**](ApiObjectStandardIdentifier.md) | Collection of tag objects identified by the name or UID. How much details are returned depends on the details-level field of the request. This table shows the level of detail shown when details-level is set to standard. | [optional] [default to null]
**Type_** | **string** | Type of the object. | [optional] [default to null]
**Uid** | **string** | Object unique identifier. | [optional] [default to null]
**UseSharedSecret** | **bool** | Indicates whether the shared secret should be used for all external gateways. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


