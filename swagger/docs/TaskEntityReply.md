# TaskEntityReply

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Color** | **string** | Color of the object. Should be one of existing colors. | [optional] [default to null]
**Comments** | **string** | Comments string. | [optional] [default to null]
**Domain** | [***ApiDomainIdentifier**](ApiDomainIdentifier.md) |  | [optional] [default to null]
**Icon** | **string** | Object icon. | [optional] [default to null]
**LastUpdateTime** | [***ApiDateReply**](ApiDateReply.md) |  | [optional] [default to null]
**MetaInfo** | [***MetaInfoForTopLevelReply**](MetaInfoForTopLevelReply.md) |  | [optional] [default to null]
**Name** | **string** | Object name. Should be unique in the domain. | [optional] [default to null]
**ProgressDescription** | **string** | N/A | [optional] [default to null]
**ProgressPercentage** | **int32** | N/A | [optional] [default to null]
**ReadOnly** | **bool** | Indicates whether the object is read-only. | [optional] [default to null]
**StartTime** | [***ApiDateReply**](ApiDateReply.md) |  | [optional] [default to null]
**Status** | **string** | Task status. | [optional] [default to null]
**Suppressed** | **bool** | N/A | [optional] [default to null]
**Tags** | [**[]ApiObjectStandardIdentifier**](ApiObjectStandardIdentifier.md) | Collection of tag objects identified by the name or UID. How much details are returned depends on the details-level field of the request. This table shows the level of detail shown when details-level is set to standard. | [optional] [default to null]
**TaskDetails** | [**[]interface{}**](interface{}.md) | Task-specific details according to the requested task type. | [optional] [default to null]
**TaskId** | **string** | Asynchronous task unique identifier. | [optional] [default to null]
**TaskName** | **string** | N/A | [optional] [default to null]
**Type_** | **string** | Type of the object. | [optional] [default to null]
**Uid** | **string** | Object unique identifier. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


