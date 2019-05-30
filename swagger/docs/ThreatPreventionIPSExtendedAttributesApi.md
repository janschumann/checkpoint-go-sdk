# \ThreatPreventionIPSExtendedAttributesApi

All URIs are relative to *https://virtserver.swaggerhub.com/Schumann-IT/checkpoint-management-api/v1.5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ShowIpsProtectionExtendedAttributePost**](ThreatPreventionIPSExtendedAttributesApi.md#ShowIpsProtectionExtendedAttributePost) | **Post** /show-ips-protection-extended-attribute | 
[**ShowIpsProtectionExtendedAttributesPost**](ThreatPreventionIPSExtendedAttributesApi.md#ShowIpsProtectionExtendedAttributesPost) | **Post** /show-ips-protection-extended-attributes | 


# **ShowIpsProtectionExtendedAttributePost**
> IpsAdditionalPropertiesShowReply ShowIpsProtectionExtendedAttributePost(ctx, apiVisualCPObjectIdentifierRequestShow)


Retrieve existing object using object name or uid.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiVisualCPObjectIdentifierRequestShow** | [**ApiVisualCpObjectIdentifierRequestShow**](ApiVisualCpObjectIdentifierRequestShow.md)|  | 

### Return type

[**IpsAdditionalPropertiesShowReply**](IpsAdditionalPropertiesShowReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowIpsProtectionExtendedAttributesPost**
> IpsAdditionalPropertiesQueryReply ShowIpsProtectionExtendedAttributesPost(ctx, apiQueryRequest)


Retrieve all objects.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiQueryRequest** | [**ApiQueryRequest**](ApiQueryRequest.md)|  | 

### Return type

[**IpsAdditionalPropertiesQueryReply**](IpsAdditionalPropertiesQueryReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

