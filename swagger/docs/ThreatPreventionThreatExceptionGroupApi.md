# \ThreatPreventionThreatExceptionGroupApi

All URIs are relative to *https://virtserver.swaggerhub.com/Schumann-IT/checkpoint-management-api/v1.5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddExceptionGroupPost**](ThreatPreventionThreatExceptionGroupApi.md#AddExceptionGroupPost) | **Post** /add-exception-group | 
[**DeleteExceptionGroupPost**](ThreatPreventionThreatExceptionGroupApi.md#DeleteExceptionGroupPost) | **Post** /delete-exception-group | 
[**SetExceptionGroupPost**](ThreatPreventionThreatExceptionGroupApi.md#SetExceptionGroupPost) | **Post** /set-exception-group | 
[**ShowExceptionGroupPost**](ThreatPreventionThreatExceptionGroupApi.md#ShowExceptionGroupPost) | **Post** /show-exception-group | 
[**ShowExceptionGroupsPost**](ThreatPreventionThreatExceptionGroupApi.md#ShowExceptionGroupsPost) | **Post** /show-exception-groups | 


# **AddExceptionGroupPost**
> ThreatExceptionGroupReply AddExceptionGroupPost(ctx, threatExceptionGroupRequestNew)


Create new object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **threatExceptionGroupRequestNew** | [**ThreatExceptionGroupRequestNew**](ThreatExceptionGroupRequestNew.md)|  | 

### Return type

[**ThreatExceptionGroupReply**](ThreatExceptionGroupReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteExceptionGroupPost**
> ApiOkReply DeleteExceptionGroupPost(ctx, apiVisualCPObjectIdentifierRequestDelete)


Delete existing object using object name or uid.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiVisualCPObjectIdentifierRequestDelete** | [**ApiVisualCpObjectIdentifierRequestDelete**](ApiVisualCpObjectIdentifierRequestDelete.md)|  | 

### Return type

[**ApiOkReply**](ApiOkReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetExceptionGroupPost**
> ThreatExceptionGroupReply SetExceptionGroupPost(ctx, threatExceptionGroupRequestEdit)


Edit existing object using object name or uid.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **threatExceptionGroupRequestEdit** | [**ThreatExceptionGroupRequestEdit**](ThreatExceptionGroupRequestEdit.md)|  | 

### Return type

[**ThreatExceptionGroupReply**](ThreatExceptionGroupReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowExceptionGroupPost**
> ThreatExceptionGroupReply ShowExceptionGroupPost(ctx, apiVisualCPObjectIdentifierRequestShow)


Retrieve existing object using object name or uid.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiVisualCPObjectIdentifierRequestShow** | [**ApiVisualCpObjectIdentifierRequestShow**](ApiVisualCpObjectIdentifierRequestShow.md)|  | 

### Return type

[**ThreatExceptionGroupReply**](ThreatExceptionGroupReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowExceptionGroupsPost**
> ApiQueryObjectReply ShowExceptionGroupsPost(ctx, apiQueryRequest)


Retrieve all objects.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiQueryRequest** | [**ApiQueryRequest**](ApiQueryRequest.md)|  | 

### Return type

[**ApiQueryObjectReply**](ApiQueryObjectReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

