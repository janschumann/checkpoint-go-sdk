# \SessionManagementSessionApi

All URIs are relative to *https://virtserver.swaggerhub.com/Schumann-IT/checkpoint-management-api/v1.5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignSessionPost**](SessionManagementSessionApi.md#AssignSessionPost) | **Post** /assign-session | 
[**ContinueSessionInSmartconsolePost**](SessionManagementSessionApi.md#ContinueSessionInSmartconsolePost) | **Post** /continue-session-in-smartconsole | 
[**PurgePublishedSessionsPost**](SessionManagementSessionApi.md#PurgePublishedSessionsPost) | **Post** /purge-published-sessions | 
[**SetSessionPost**](SessionManagementSessionApi.md#SetSessionPost) | **Post** /set-session | 
[**ShowLastPublishedSessionPost**](SessionManagementSessionApi.md#ShowLastPublishedSessionPost) | **Post** /show-last-published-session | 
[**ShowSessionPost**](SessionManagementSessionApi.md#ShowSessionPost) | **Post** /show-session | 
[**ShowSessionsPost**](SessionManagementSessionApi.md#ShowSessionsPost) | **Post** /show-sessions | 
[**SwitchSessionPost**](SessionManagementSessionApi.md#SwitchSessionPost) | **Post** /switch-session | 
[**TakeOverSessionPost**](SessionManagementSessionApi.md#TakeOverSessionPost) | **Post** /take-over-session | 


# **AssignSessionPost**
> ApiOkReply AssignSessionPost(ctx, workSessionAssignRequest)


Assign a session ownership to another administrator.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workSessionAssignRequest** | [**WorkSessionAssignRequest**](WorkSessionAssignRequest.md)|  | 

### Return type

[**ApiOkReply**](ApiOkReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContinueSessionInSmartconsolePost**
> ApiOkReply ContinueSessionInSmartconsolePost(ctx, workSessionObjectIdentifierRequest)


Logout from existing session. The session will be continued next time your open SmartConsole. In case 'uid' is not provided, use current session. In order for the session to pass successfully to SmartConsole, make sure you don't have any other active GUI sessions.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workSessionObjectIdentifierRequest** | [**WorkSessionObjectIdentifierRequest**](WorkSessionObjectIdentifierRequest.md)|  | 

### Return type

[**ApiOkReply**](ApiOkReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PurgePublishedSessionsPost**
> ApiTaskReply PurgePublishedSessionsPost(ctx, workSessionPurgeRequest)


Permanently deletes all data which belongs to the published sessions not selected for preservation. This operation is irreversible.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workSessionPurgeRequest** | [**WorkSessionPurgeRequest**](WorkSessionPurgeRequest.md)|  | 

### Return type

[**ApiTaskReply**](ApiTaskReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetSessionPost**
> WorkSessionReply SetSessionPost(ctx, workSessionRequestEdit)


Edit user's current session.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workSessionRequestEdit** | [**WorkSessionRequestEdit**](WorkSessionRequestEdit.md)|  | 

### Return type

[**WorkSessionReply**](WorkSessionReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowLastPublishedSessionPost**
> WorkSessionReply ShowLastPublishedSessionPost(ctx, emptyRequest)


Shows the last published session.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **emptyRequest** | [**EmptyRequest**](EmptyRequest.md)|  | 

### Return type

[**WorkSessionReply**](WorkSessionReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowSessionPost**
> WorkSessionReply ShowSessionPost(ctx, workSessionObjectIdentifierRequest)


Show session.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workSessionObjectIdentifierRequest** | [**WorkSessionObjectIdentifierRequest**](WorkSessionObjectIdentifierRequest.md)|  | 

### Return type

[**WorkSessionReply**](WorkSessionReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowSessionsPost**
> ApiQueryObjectReply ShowSessionsPost(ctx, workSessionQueryRequest)


Retrieve all objects.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workSessionQueryRequest** | [**WorkSessionQueryRequest**](WorkSessionQueryRequest.md)|  | 

### Return type

[**ApiQueryObjectReply**](ApiQueryObjectReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SwitchSessionPost**
> WorkSessionReply SwitchSessionPost(ctx, workSessionSwitchRequest)


Switch to another session.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workSessionSwitchRequest** | [**WorkSessionSwitchRequest**](WorkSessionSwitchRequest.md)|  | 

### Return type

[**WorkSessionReply**](WorkSessionReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TakeOverSessionPost**
> WorkSessionReply TakeOverSessionPost(ctx, workSessionTakeOverRequest)


Take ownership of another session and start working on it.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workSessionTakeOverRequest** | [**WorkSessionTakeOverRequest**](WorkSessionTakeOverRequest.md)|  | 

### Return type

[**WorkSessionReply**](WorkSessionReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

