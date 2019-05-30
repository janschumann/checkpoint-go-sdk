# \AccessControlNATNATSectionApi

All URIs are relative to *https://virtserver.swaggerhub.com/Schumann-IT/checkpoint-management-api/v1.5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddNatSectionPost**](AccessControlNATNATSectionApi.md#AddNatSectionPost) | **Post** /add-nat-section | 
[**DeleteNatSectionPost**](AccessControlNATNATSectionApi.md#DeleteNatSectionPost) | **Post** /delete-nat-section | 
[**SetNatSectionPost**](AccessControlNATNATSectionApi.md#SetNatSectionPost) | **Post** /set-nat-section | 
[**ShowNatSectionPost**](AccessControlNATNATSectionApi.md#ShowNatSectionPost) | **Post** /show-nat-section | 


# **AddNatSectionPost**
> NatSectionReply AddNatSectionPost(ctx, natSectionRequestNew)


Create new object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **natSectionRequestNew** | [**NatSectionRequestNew**](NatSectionRequestNew.md)|  | 

### Return type

[**NatSectionReply**](NatSectionReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteNatSectionPost**
> ApiOkReply DeleteNatSectionPost(ctx, natSectionIdentifierRequest)


Delete existing object using object name or uid.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **natSectionIdentifierRequest** | [**NatSectionIdentifierRequest**](NatSectionIdentifierRequest.md)|  | 

### Return type

[**ApiOkReply**](ApiOkReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetNatSectionPost**
> NatSectionReply SetNatSectionPost(ctx, natSectionRequestEdit)


Edit existing object using object name or uid.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **natSectionRequestEdit** | [**NatSectionRequestEdit**](NatSectionRequestEdit.md)|  | 

### Return type

[**NatSectionReply**](NatSectionReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowNatSectionPost**
> NatSectionReply ShowNatSectionPost(ctx, natSectionIdentifierRequest)


Retrieve existing object using object name or uid.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **natSectionIdentifierRequest** | [**NatSectionIdentifierRequest**](NatSectionIdentifierRequest.md)|  | 

### Return type

[**NatSectionReply**](NatSectionReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

