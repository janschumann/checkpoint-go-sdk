# \ThreatPreventionThreatExceptionApi

All URIs are relative to *https://virtserver.swaggerhub.com/Schumann-IT/checkpoint-management-api/v1.5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddThreatExceptionPost**](ThreatPreventionThreatExceptionApi.md#AddThreatExceptionPost) | **Post** /add-threat-exception | 
[**DeleteThreatExceptionPost**](ThreatPreventionThreatExceptionApi.md#DeleteThreatExceptionPost) | **Post** /delete-threat-exception | 
[**SetThreatExceptionPost**](ThreatPreventionThreatExceptionApi.md#SetThreatExceptionPost) | **Post** /set-threat-exception | 
[**ShowThreatExceptionPost**](ThreatPreventionThreatExceptionApi.md#ShowThreatExceptionPost) | **Post** /show-threat-exception | 
[**ShowThreatRuleExceptionRulebasePost**](ThreatPreventionThreatExceptionApi.md#ShowThreatRuleExceptionRulebasePost) | **Post** /show-threat-rule-exception-rulebase | 


# **AddThreatExceptionPost**
> ThreatExceptionReply AddThreatExceptionPost(ctx, threatExceptionRequestNew)


Create new object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **threatExceptionRequestNew** | [**ThreatExceptionRequestNew**](ThreatExceptionRequestNew.md)|  | 

### Return type

[**ThreatExceptionReply**](ThreatExceptionReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteThreatExceptionPost**
> ApiOkReply DeleteThreatExceptionPost(ctx, threatExceptionIdentifierRequest)


Delete existing object using object name or uid.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **threatExceptionIdentifierRequest** | [**ThreatExceptionIdentifierRequest**](ThreatExceptionIdentifierRequest.md)|  | 

### Return type

[**ApiOkReply**](ApiOkReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetThreatExceptionPost**
> ThreatExceptionReply SetThreatExceptionPost(ctx, threatExceptionRequestEdit)


Edit existing object using object name or uid.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **threatExceptionRequestEdit** | [**ThreatExceptionRequestEdit**](ThreatExceptionRequestEdit.md)|  | 

### Return type

[**ThreatExceptionReply**](ThreatExceptionReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowThreatExceptionPost**
> ThreatExceptionReply ShowThreatExceptionPost(ctx, threatExceptionIdentifierRequest)


Retrieve existing object using object name or uid.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **threatExceptionIdentifierRequest** | [**ThreatExceptionIdentifierRequest**](ThreatExceptionIdentifierRequest.md)|  | 

### Return type

[**ThreatExceptionReply**](ThreatExceptionReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowThreatRuleExceptionRulebasePost**
> QueryThreatExceptionRulebaseReply ShowThreatRuleExceptionRulebasePost(ctx, queryThreatExceptionRulebaseRequest)


Shows the entire Threat Exceptions layer  generated for a given threat rule.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **queryThreatExceptionRulebaseRequest** | [**QueryThreatExceptionRulebaseRequest**](QueryThreatExceptionRulebaseRequest.md)|  | 

### Return type

[**QueryThreatExceptionRulebaseReply**](QueryThreatExceptionRulebaseReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

