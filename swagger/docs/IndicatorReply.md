# IndicatorReply

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | The indicator&#39;s action. | [optional] [default to null]
**Color** | **string** | Color of the object. Should be one of existing colors. | [optional] [default to null]
**Comments** | **string** | Comments string. | [optional] [default to null]
**Domain** | [***ApiDomainIdentifier**](ApiDomainIdentifier.md) |  | [optional] [default to null]
**Icon** | **string** | Object icon. | [optional] [default to null]
**MetaInfo** | [***MetaInfoForTopLevelReply**](MetaInfoForTopLevelReply.md) |  | [optional] [default to null]
**Name** | **string** | Object name. Should be unique in the domain. | [optional] [default to null]
**NumberOfObservables** | **int32** | The amount of observables in the indicator. | [optional] [default to null]
**ProfileOverrides** | [**[]IndicatorOverrideRequest**](IndicatorOverrideRequest.md) | A list of profiles where the default action of this indicator is overridden, along with the overriding action. | [optional] [default to null]
**ReadOnly** | **bool** | Indicates whether the object is read-only. | [optional] [default to null]
**Tags** | [**[]ApiObjectStandardIdentifier**](ApiObjectStandardIdentifier.md) | Collection of tag objects identified by the name or UID. How much details are returned depends on the details-level field of the request. This table shows the level of detail shown when details-level is set to standard. | [optional] [default to null]
**Type_** | **string** | Type of the object. | [optional] [default to null]
**Uid** | **string** | Object unique identifier. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


