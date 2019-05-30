# \NetworkObjectsOPSECApplicationApi

All URIs are relative to *https://virtserver.swaggerhub.com/Schumann-IT/checkpoint-management-api/v1.5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddOpsecApplicationPost**](NetworkObjectsOPSECApplicationApi.md#AddOpsecApplicationPost) | **Post** /add-opsec-application | 
[**DeleteOpsecApplicationPost**](NetworkObjectsOPSECApplicationApi.md#DeleteOpsecApplicationPost) | **Post** /delete-opsec-application | 
[**SetOpsecApplicationPost**](NetworkObjectsOPSECApplicationApi.md#SetOpsecApplicationPost) | **Post** /set-opsec-application | 
[**ShowOpsecApplicationPost**](NetworkObjectsOPSECApplicationApi.md#ShowOpsecApplicationPost) | **Post** /show-opsec-application | 
[**ShowOpsecApplicationsPost**](NetworkObjectsOPSECApplicationApi.md#ShowOpsecApplicationsPost) | **Post** /show-opsec-applications | 


# **AddOpsecApplicationPost**
> OpsecApplicationReply AddOpsecApplicationPost(ctx, opsecApplicationRequestNew)


Create a new OPSEC Application.  At least one client entity (LEA, CPMI) must be supplied.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **opsecApplicationRequestNew** | [**OpsecApplicationRequestNew**](OpsecApplicationRequestNew.md)|  | 

### Return type

[**OpsecApplicationReply**](OpsecApplicationReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOpsecApplicationPost**
> ApiOkReply DeleteOpsecApplicationPost(ctx, apiVisualCPObjectIdentifierRequestDelete)


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

# **SetOpsecApplicationPost**
> OpsecApplicationReply SetOpsecApplicationPost(ctx, opsecApplicationRequestEdit)


Edit existing object using object name or uid.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **opsecApplicationRequestEdit** | [**OpsecApplicationRequestEdit**](OpsecApplicationRequestEdit.md)|  | 

### Return type

[**OpsecApplicationReply**](OpsecApplicationReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowOpsecApplicationPost**
> OpsecApplicationReply ShowOpsecApplicationPost(ctx, apiVisualCPObjectIdentifierRequestShow)


Retrieve existing object using object name or uid.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiVisualCPObjectIdentifierRequestShow** | [**ApiVisualCpObjectIdentifierRequestShow**](ApiVisualCpObjectIdentifierRequestShow.md)|  | 

### Return type

[**OpsecApplicationReply**](OpsecApplicationReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowOpsecApplicationsPost**
> ApiQueryObjectReply ShowOpsecApplicationsPost(ctx, objectInGroupQueryRequest)


Retrieve all objects.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectInGroupQueryRequest** | [**ObjectInGroupQueryRequest**](ObjectInGroupQueryRequest.md)|  | 

### Return type

[**ApiQueryObjectReply**](ApiQueryObjectReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

