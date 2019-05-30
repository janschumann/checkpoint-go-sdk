# \AccessControlNATAccessRuleApi

All URIs are relative to *https://virtserver.swaggerhub.com/Schumann-IT/checkpoint-management-api/v1.5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAccessRulePost**](AccessControlNATAccessRuleApi.md#AddAccessRulePost) | **Post** /add-access-rule | 
[**DeleteAccessRulePost**](AccessControlNATAccessRuleApi.md#DeleteAccessRulePost) | **Post** /delete-access-rule | 
[**SetAccessRulePost**](AccessControlNATAccessRuleApi.md#SetAccessRulePost) | **Post** /set-access-rule | 
[**ShowAccessRulePost**](AccessControlNATAccessRuleApi.md#ShowAccessRulePost) | **Post** /show-access-rule | 
[**ShowAccessRulebasePost**](AccessControlNATAccessRuleApi.md#ShowAccessRulebasePost) | **Post** /show-access-rulebase | 


# **AddAccessRulePost**
> AccessRuleReply AddAccessRulePost(ctx, accessRuleRequestNew)


Create new object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessRuleRequestNew** | [**AccessRuleRequestNew**](AccessRuleRequestNew.md)|  | 

### Return type

[**AccessRuleReply**](AccessRuleReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAccessRulePost**
> ApiOkReply DeleteAccessRulePost(ctx, accessRuleIdentifierRequest)


Delete existing object using object name or uid.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessRuleIdentifierRequest** | [**AccessRuleIdentifierRequest**](AccessRuleIdentifierRequest.md)|  | 

### Return type

[**ApiOkReply**](ApiOkReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetAccessRulePost**
> AccessRuleReply SetAccessRulePost(ctx, accessRuleRequestEdit)


Edit existing object using object name or uid.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessRuleRequestEdit** | [**AccessRuleRequestEdit**](AccessRuleRequestEdit.md)|  | 

### Return type

[**AccessRuleReply**](AccessRuleReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowAccessRulePost**
> AccessRuleReply ShowAccessRulePost(ctx, accessRuleIdentifierRequestShow)


Retrieve existing object using object name or uid.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessRuleIdentifierRequestShow** | [**AccessRuleIdentifierRequestShow**](AccessRuleIdentifierRequestShow.md)|  | 

### Return type

[**AccessRuleReply**](AccessRuleReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowAccessRulebasePost**
> QueryAccessRulebaseReply ShowAccessRulebasePost(ctx, queryAccessRulebaseRequest)


Shows the entire Access Rules layer.  This layer is divided into sections. An Access Rule may be within a section, or independent of a section (in which case it is said to be under the \"global\" section). The reply features a list of objects. Each object may be a section of the layer, with all its rules in, or a rule itself, for the case of rules which are under the global section. An optional \"filter\" field may be added in order to filter out only those rules that match a search criteria.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **queryAccessRulebaseRequest** | [**QueryAccessRulebaseRequest**](QueryAccessRulebaseRequest.md)|  | 

### Return type

[**QueryAccessRulebaseReply**](QueryAccessRulebaseReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

