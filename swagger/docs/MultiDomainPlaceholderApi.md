# \MultiDomainPlaceholderApi

All URIs are relative to *https://virtserver.swaggerhub.com/Schumann-IT/checkpoint-management-api/v1.5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ShowPlaceHolderPost**](MultiDomainPlaceholderApi.md#ShowPlaceHolderPost) | **Post** /show-place-holder | 


# **ShowPlaceHolderPost**
> ApiPlaceHolderReply ShowPlaceHolderPost(ctx, apiPlaceHolderIdentifierRequest)


Retrieve existing object using object name or uid.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiPlaceHolderIdentifierRequest** | [**ApiPlaceHolderIdentifierRequest**](ApiPlaceHolderIdentifierRequest.md)|  | 

### Return type

[**ApiPlaceHolderReply**](ApiPlaceHolderReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

