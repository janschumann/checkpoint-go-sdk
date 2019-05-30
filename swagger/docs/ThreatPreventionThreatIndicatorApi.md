# \ThreatPreventionThreatIndicatorApi

All URIs are relative to *https://virtserver.swaggerhub.com/Schumann-IT/checkpoint-management-api/v1.5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddThreatIndicatorPost**](ThreatPreventionThreatIndicatorApi.md#AddThreatIndicatorPost) | **Post** /add-threat-indicator | 
[**DeleteThreatIndicatorPost**](ThreatPreventionThreatIndicatorApi.md#DeleteThreatIndicatorPost) | **Post** /delete-threat-indicator | 
[**SetThreatIndicatorPost**](ThreatPreventionThreatIndicatorApi.md#SetThreatIndicatorPost) | **Post** /set-threat-indicator | 
[**ShowThreatIndicatorPost**](ThreatPreventionThreatIndicatorApi.md#ShowThreatIndicatorPost) | **Post** /show-threat-indicator | 
[**ShowThreatIndicatorsPost**](ThreatPreventionThreatIndicatorApi.md#ShowThreatIndicatorsPost) | **Post** /show-threat-indicators | 


# **AddThreatIndicatorPost**
> ApiTaskReply AddThreatIndicatorPost(ctx, indicatorRequestNew)


Create a new Threat-Indicator.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **indicatorRequestNew** | [**IndicatorRequestNew**](IndicatorRequestNew.md)|  | 

### Return type

[**ApiTaskReply**](ApiTaskReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteThreatIndicatorPost**
> ApiOkReply DeleteThreatIndicatorPost(ctx, apiVisualCPObjectIdentifierRequestShow)


Delete a Threat-Indicator.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiVisualCPObjectIdentifierRequestShow** | [**ApiVisualCpObjectIdentifierRequestShow**](ApiVisualCpObjectIdentifierRequestShow.md)|  | 

### Return type

[**ApiOkReply**](ApiOkReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetThreatIndicatorPost**
> IndicatorReply SetThreatIndicatorPost(ctx, indicatorRequestEdit)


Edit an existing Threat-Indicator.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **indicatorRequestEdit** | [**IndicatorRequestEdit**](IndicatorRequestEdit.md)|  | 

### Return type

[**IndicatorReply**](IndicatorReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowThreatIndicatorPost**
> IndicatorReply ShowThreatIndicatorPost(ctx, apiVisualCPObjectIdentifierRequestShow)


Displays a Threat-Indicator.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiVisualCPObjectIdentifierRequestShow** | [**ApiVisualCpObjectIdentifierRequestShow**](ApiVisualCpObjectIdentifierRequestShow.md)|  | 

### Return type

[**IndicatorReply**](IndicatorReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowThreatIndicatorsPost**
> IndicatorQueryReply ShowThreatIndicatorsPost(ctx, apiQueryRequest)


Display a list of Threat-Indicators.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiQueryRequest** | [**ApiQueryRequest**](ApiQueryRequest.md)|  | 

### Return type

[**IndicatorQueryReply**](IndicatorQueryReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

