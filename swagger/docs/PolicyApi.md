# \PolicyApi

All URIs are relative to *https://virtserver.swaggerhub.com/Schumann-IT/checkpoint-management-api/v1.5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InstallPolicyPost**](PolicyApi.md#InstallPolicyPost) | **Post** /install-policy | 
[**VerifyPolicyPost**](PolicyApi.md#VerifyPolicyPost) | **Post** /verify-policy | 


# **InstallPolicyPost**
> PolicyInstallationReply InstallPolicyPost(ctx, policyInstallationRequest)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **policyInstallationRequest** | [**PolicyInstallationRequest**](PolicyInstallationRequest.md)|  | 

### Return type

[**PolicyInstallationReply**](PolicyInstallationReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VerifyPolicyPost**
> VerifyPolicyReply VerifyPolicyPost(ctx, verifyPolicyRequest)


Verifies the policy of the selected package.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **verifyPolicyRequest** | [**VerifyPolicyRequest**](VerifyPolicyRequest.md)|  | 

### Return type

[**VerifyPolicyReply**](VerifyPolicyReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

