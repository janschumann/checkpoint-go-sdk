# \ThreatPreventionThreatProtectionApi

All URIs are relative to *https://virtserver.swaggerhub.com/Schumann-IT/checkpoint-management-api/v1.5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddThreatProtectionsPost**](ThreatPreventionThreatProtectionApi.md#AddThreatProtectionsPost) | **Post** /add-threat-protections | 
[**DeleteThreatProtectionsPost**](ThreatPreventionThreatProtectionApi.md#DeleteThreatProtectionsPost) | **Post** /delete-threat-protections | 
[**SetThreatProtectionPost**](ThreatPreventionThreatProtectionApi.md#SetThreatProtectionPost) | **Post** /set-threat-protection | 
[**ShowThreatProtectionPost**](ThreatPreventionThreatProtectionApi.md#ShowThreatProtectionPost) | **Post** /show-threat-protection | 
[**ShowThreatProtectionsPost**](ThreatPreventionThreatProtectionApi.md#ShowThreatProtectionsPost) | **Post** /show-threat-protections | 


# **AddThreatProtectionsPost**
> ApiTaskReply AddThreatProtectionsPost(ctx, addProtectionsRequest)


Adds threat protections.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **addProtectionsRequest** | [**AddProtectionsRequest**](AddProtectionsRequest.md)|  | 

### Return type

[**ApiTaskReply**](ApiTaskReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteThreatProtectionsPost**
> ApiTaskReply DeleteThreatProtectionsPost(ctx, deleteProtectionsRequest)


Deletes threat protections.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deleteProtectionsRequest** | [**DeleteProtectionsRequest**](DeleteProtectionsRequest.md)|  | 

### Return type

[**ApiTaskReply**](ApiTaskReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetThreatProtectionPost**
> ProtectionReply SetThreatProtectionPost(ctx, protectionRequestEdit)


Edit existing object using object name or uid.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **protectionRequestEdit** | [**ProtectionRequestEdit**](ProtectionRequestEdit.md)|  | 

### Return type

[**ProtectionReply**](ProtectionReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowThreatProtectionPost**
> ProtectionReply ShowThreatProtectionPost(ctx, apiVisualCPObjectIdentifierRequestShow)


Retrieve existing object using object name or uid.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiVisualCPObjectIdentifierRequestShow** | [**ApiVisualCpObjectIdentifierRequestShow**](ApiVisualCpObjectIdentifierRequestShow.md)|  | 

### Return type

[**ProtectionReply**](ProtectionReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowThreatProtectionsPost**
> ProtectionQueryReply ShowThreatProtectionsPost(ctx, apiQueryRequest)


Retrieve all objects.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiQueryRequest** | [**ApiQueryRequest**](ApiQueryRequest.md)|  | 

### Return type

[**ProtectionQueryReply**](ProtectionQueryReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

