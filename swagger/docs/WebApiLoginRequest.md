# WebApiLoginRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContinueLastSession** | **bool** | When &#39;continue-last-session&#39; is set to &#39;True&#39;, the new session would continue where the last session was stopped. This option is available when the administrator has only one session that can be continued. If there is more than one session, see &#39;switch-session&#39; API. | [optional] [default to null]
**Domain** | **string** | Use domain to login to specific domain. Domain can be identified by name or UID. | [optional] [default to null]
**EnterLastPublishedSession** | **bool** | Login to the last published session. Such login is done with the Read Only permissions. | [optional] [default to null]
**Password** | **string** | Administrator password. | [default to null]
**ReadOnly** | **bool** | Login with Read Only permissions. This parameter is not considered in case continue-last-session is true. | [optional] [default to null]
**SessionComments** | **string** | Session comments. | [optional] [default to null]
**SessionDescription** | **string** | Session description. | [optional] [default to null]
**SessionName** | **string** | Session unique name. | [optional] [default to null]
**SessionTimeout** | **int32** | Session expiration timeout in seconds. Default 600 seconds. | [optional] [default to null]
**User** | **string** | Administrator user name. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


