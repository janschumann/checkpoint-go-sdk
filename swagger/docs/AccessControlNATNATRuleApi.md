# \AccessControlNATNATRuleApi

All URIs are relative to *https://virtserver.swaggerhub.com/Schumann-IT/checkpoint-management-api/v1.5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddNatRulePost**](AccessControlNATNATRuleApi.md#AddNatRulePost) | **Post** /add-nat-rule | 
[**DeleteNatRulePost**](AccessControlNATNATRuleApi.md#DeleteNatRulePost) | **Post** /delete-nat-rule | 
[**SetNatRulePost**](AccessControlNATNATRuleApi.md#SetNatRulePost) | **Post** /set-nat-rule | 
[**ShowNatRulePost**](AccessControlNATNATRuleApi.md#ShowNatRulePost) | **Post** /show-nat-rule | 
[**ShowNatRulebasePost**](AccessControlNATNATRuleApi.md#ShowNatRulebasePost) | **Post** /show-nat-rulebase | 


# **AddNatRulePost**
> NatRuleReply AddNatRulePost(ctx, natRuleRequestNew)


Create new object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **natRuleRequestNew** | [**NatRuleRequestNew**](NatRuleRequestNew.md)|  | 

### Return type

[**NatRuleReply**](NatRuleReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteNatRulePost**
> ApiOkReply DeleteNatRulePost(ctx, natRuleIdentifierRequest)


Delete existing object using object name or uid.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **natRuleIdentifierRequest** | [**NatRuleIdentifierRequest**](NatRuleIdentifierRequest.md)|  | 

### Return type

[**ApiOkReply**](ApiOkReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetNatRulePost**
> NatRuleReply SetNatRulePost(ctx, natRuleRequestEdit)


Edit existing object using object name or uid.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **natRuleRequestEdit** | [**NatRuleRequestEdit**](NatRuleRequestEdit.md)|  | 

### Return type

[**NatRuleReply**](NatRuleReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowNatRulePost**
> NatRuleReply ShowNatRulePost(ctx, natRuleIdentifierRequest)


Retrieve existing object using object name or uid.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **natRuleIdentifierRequest** | [**NatRuleIdentifierRequest**](NatRuleIdentifierRequest.md)|  | 

### Return type

[**NatRuleReply**](NatRuleReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowNatRulebasePost**
> QueryNatRulebaseReply ShowNatRulebasePost(ctx, queryNatRulebaseRequest)


Shows the entire NAT Rules layer.  This layer is divided into sections. A NAT Rule may be within a section, or independent of a section (in which case it is said to be under the \"global\" section). There are two types of sections: auto generated read only sections and general sections which are created manually. The reply features a list of objects. Each object may be a section of the layer, within which its rules may be found, or a rule itself, for the case of rules which are under the global section. An optional \"filter\" field may be added in order to filter out only those rules that match a search criteria.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **queryNatRulebaseRequest** | [**QueryNatRulebaseRequest**](QueryNatRulebaseRequest.md)|  | 

### Return type

[**QueryNatRulebaseReply**](QueryNatRulebaseReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

