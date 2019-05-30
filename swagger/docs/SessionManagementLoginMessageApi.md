# \SessionManagementLoginMessageApi

All URIs are relative to *https://virtserver.swaggerhub.com/Schumann-IT/checkpoint-management-api/v1.5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SetLoginMessagePost**](SessionManagementLoginMessageApi.md#SetLoginMessagePost) | **Post** /set-login-message | 
[**ShowLoginMessagePost**](SessionManagementLoginMessageApi.md#ShowLoginMessagePost) | **Post** /show-login-message | 


# **SetLoginMessagePost**
> LoginMessageReply SetLoginMessagePost(ctx, loginMessageRequestSet)


Edit existing object using object name or uid.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **loginMessageRequestSet** | [**LoginMessageRequestSet**](LoginMessageRequestSet.md)|  | 

### Return type

[**LoginMessageReply**](LoginMessageReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowLoginMessagePost**
> LoginMessageReply ShowLoginMessagePost(ctx, loginMessageRequestShow)


Retrieve existing object using object name or uid.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **loginMessageRequestShow** | [**LoginMessageRequestShow**](LoginMessageRequestShow.md)|  | 

### Return type

[**LoginMessageReply**](LoginMessageReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

