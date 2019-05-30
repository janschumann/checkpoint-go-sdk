# WorkSessionReply

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Application** | **string** | The name of the application serving the web_api requests. | [optional] [default to null]
**Changes** | **int32** | Number of pending changes. | [optional] [default to null]
**Color** | **string** | Color of the object. Should be one of existing colors. | [optional] [default to null]
**Comments** | **string** | Comments string. | [optional] [default to null]
**ConnectionMode** | **string** | N/A | [optional] [default to null]
**Description** | **string** | Session description. | [optional] [default to null]
**Domain** | [***ApiDomainIdentifier**](ApiDomainIdentifier.md) |  | [optional] [default to null]
**Email** | **string** | Administrator email. | [optional] [default to null]
**ExpiredSession** | **bool** | True if the session is expired. | [optional] [default to null]
**Icon** | **string** | Object icon. | [optional] [default to null]
**InWork** | **bool** | True if the session is in work state. | [optional] [default to null]
**IpAddress** | **string** | IP address from which the session was initiated. | [optional] [default to null]
**LastLoginTime** | [***ApiDateReply**](ApiDateReply.md) |  | [optional] [default to null]
**LastLogoutTime** | [***ApiDateReply**](ApiDateReply.md) |  | [optional] [default to null]
**Locks** | **int32** | Number of locked objects. | [optional] [default to null]
**MetaInfo** | [***MetaInfoForTopLevelReply**](MetaInfoForTopLevelReply.md) |  | [optional] [default to null]
**Name** | **string** | Object name. Should be unique in the domain. | [optional] [default to null]
**PhoneNumber** | **string** | Administrator phone number. | [optional] [default to null]
**PublishTime** | [***ApiDateReply**](ApiDateReply.md) |  | [optional] [default to null]
**ReadOnly** | **bool** | Indicates whether the object is read-only. | [optional] [default to null]
**SessionTimeout** | **int32** | Session expiration timeout in seconds. | [optional] [default to null]
**State** | **string** | Session state. | [optional] [default to null]
**Tags** | [**[]ApiObjectStandardIdentifier**](ApiObjectStandardIdentifier.md) | Collection of tag objects identified by the name or UID. How much details are returned depends on the details-level field of the request. This table shows the level of detail shown when details-level is set to standard. | [optional] [default to null]
**Type_** | **string** | Type of the object. | [optional] [default to null]
**Uid** | **string** | Object unique identifier. | [optional] [default to null]
**UserName** | **string** | The name of the logged in user. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


