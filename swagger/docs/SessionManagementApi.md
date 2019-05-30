# \SessionManagementApi

All URIs are relative to *https://virtserver.swaggerhub.com/Schumann-IT/checkpoint-management-api/v1.5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DiscardPost**](SessionManagementApi.md#DiscardPost) | **Post** /discard | 
[**DisconnectPost**](SessionManagementApi.md#DisconnectPost) | **Post** /disconnect | 
[**KeepalivePost**](SessionManagementApi.md#KeepalivePost) | **Post** /keepalive | 
[**LoginPost**](SessionManagementApi.md#LoginPost) | **Post** /login | 
[**LogoutPost**](SessionManagementApi.md#LogoutPost) | **Post** /logout | 
[**PublishPost**](SessionManagementApi.md#PublishPost) | **Post** /publish | 


# **DiscardPost**
> DiscardReply DiscardPost(ctx, discardRequest)


All changes done by user are discarded and removed from database.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **discardRequest** | [**DiscardRequest**](DiscardRequest.md)|  | 

### Return type

[**DiscardReply**](DiscardReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DisconnectPost**
> DisconnectReply DisconnectPost(ctx, disconnectRequest)


Disconnect a private session.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **disconnectRequest** | [**DisconnectRequest**](DisconnectRequest.md)|  | 

### Return type

[**DisconnectReply**](DisconnectReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **KeepalivePost**
> KeepAliveReply KeepalivePost(ctx, keepAliveRequest)


Keep the session valid/alive.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **keepAliveRequest** | [**KeepAliveRequest**](KeepAliveRequest.md)|  | 

### Return type

[**KeepAliveReply**](KeepAliveReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LoginPost**
> WebApiLoginReply LoginPost(ctx, webApiLoginRequest)


Log in to the server with username and password. The server shows your session unique identifier. Enter this session unique identifier in the 'X-chkp-sid' header of each request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **webApiLoginRequest** | [**WebApiLoginRequest**](WebApiLoginRequest.md)|  | 

### Return type

[**WebApiLoginReply**](WebApiLoginReply.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LogoutPost**
> WebApiLogoutReply LogoutPost(ctx, webApiLogoutRequest)


Log out from the current session. After logging out the session id is not valid any more.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **webApiLogoutRequest** | [**WebApiLogoutRequest**](WebApiLogoutRequest.md)|  | 

### Return type

[**WebApiLogoutReply**](WebApiLogoutReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PublishPost**
> PublishReply PublishPost(ctx, publishRequest)


All the changes done by this user will be seen by all users only after publish is called.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **publishRequest** | [**PublishRequest**](PublishRequest.md)|  | 

### Return type

[**PublishReply**](PublishReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

