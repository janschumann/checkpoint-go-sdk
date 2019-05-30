# WebApiLoginReply

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiServerVersion** | **string** | API Server version. | [optional] [default to null]
**DiskSpaceMessage** | **string** | Information about the available disk space on the management server. | [optional] [default to null]
**LastLoginWasAt** | [***ApiDateReply**](ApiDateReply.md) |  | [optional] [default to null]
**LoginMessage** | [***LoginMessageReply**](LoginMessageReply.md) |  | [optional] [default to null]
**ReadOnly** | **bool** | True if this session is read only. | [optional] [default to null]
**SessionTimeout** | **int32** | Session expiration timeout in seconds. | [optional] [default to null]
**Sid** | **string** | Session unique identifier. Enter this session unique identifier in the &#39;X-chkp-sid&#39; header of each request. | [optional] [default to null]
**Standby** | **bool** | True if this management server is in the standby mode. | [optional] [default to null]
**Uid** | **string** | Session object unique identifier. This identifier may be used in the discard API to discard changes that were made in this session, when administrator is working from another session, or in the &#39;switch-session&#39; API. | [optional] [default to null]
**Url** | **string** | URL that was used to reach the API server. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


