# \ThreatPreventionThreatRuleApi

All URIs are relative to *https://virtserver.swaggerhub.com/Schumann-IT/checkpoint-management-api/v1.5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddThreatRulePost**](ThreatPreventionThreatRuleApi.md#AddThreatRulePost) | **Post** /add-threat-rule | 
[**DeleteThreatRulePost**](ThreatPreventionThreatRuleApi.md#DeleteThreatRulePost) | **Post** /delete-threat-rule | 
[**SetThreatRulePost**](ThreatPreventionThreatRuleApi.md#SetThreatRulePost) | **Post** /set-threat-rule | 
[**ShowThreatRulePost**](ThreatPreventionThreatRuleApi.md#ShowThreatRulePost) | **Post** /show-threat-rule | 
[**ShowThreatRulebasePost**](ThreatPreventionThreatRuleApi.md#ShowThreatRulebasePost) | **Post** /show-threat-rulebase | 


# **AddThreatRulePost**
> ThreatRuleReply AddThreatRulePost(ctx, threatRuleRequestNew)


Create new object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **threatRuleRequestNew** | [**ThreatRuleRequestNew**](ThreatRuleRequestNew.md)|  | 

### Return type

[**ThreatRuleReply**](ThreatRuleReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteThreatRulePost**
> ApiOkReply DeleteThreatRulePost(ctx, threatRuleIdentifierRequest)


Delete existing object using object name or uid.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **threatRuleIdentifierRequest** | [**ThreatRuleIdentifierRequest**](ThreatRuleIdentifierRequest.md)|  | 

### Return type

[**ApiOkReply**](ApiOkReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetThreatRulePost**
> ThreatRuleReply SetThreatRulePost(ctx, threatRuleRequestEdit)


Edit existing object using object name or uid.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **threatRuleRequestEdit** | [**ThreatRuleRequestEdit**](ThreatRuleRequestEdit.md)|  | 

### Return type

[**ThreatRuleReply**](ThreatRuleReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowThreatRulePost**
> ThreatRuleReply ShowThreatRulePost(ctx, threatRuleIdentifierRequest)


Retrieve existing object using object name or uid.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **threatRuleIdentifierRequest** | [**ThreatRuleIdentifierRequest**](ThreatRuleIdentifierRequest.md)|  | 

### Return type

[**ThreatRuleReply**](ThreatRuleReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowThreatRulebasePost**
> QueryThreatRulebaseReply ShowThreatRulebasePost(ctx, queryThreatRulebaseRequest)


Shows the entire Threat Prevention Rules layer. The reply features a list of rules. Each rule has the Global Exceptions Group attached and may have any number of an Exceptions Group attached. An optional \"filter\" field may be added in order to filter out only those rules that match a search criteria.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **queryThreatRulebaseRequest** | [**QueryThreatRulebaseRequest**](QueryThreatRulebaseRequest.md)|  | 

### Return type

[**QueryThreatRulebaseReply**](QueryThreatRulebaseReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

