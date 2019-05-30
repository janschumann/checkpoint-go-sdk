# \AccessControlNATAccessSectionApi

All URIs are relative to *https://virtserver.swaggerhub.com/Schumann-IT/checkpoint-management-api/v1.5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAccessSectionPost**](AccessControlNATAccessSectionApi.md#AddAccessSectionPost) | **Post** /add-access-section | 
[**DeleteAccessSectionPost**](AccessControlNATAccessSectionApi.md#DeleteAccessSectionPost) | **Post** /delete-access-section | 
[**SetAccessSectionPost**](AccessControlNATAccessSectionApi.md#SetAccessSectionPost) | **Post** /set-access-section | 
[**ShowAccessSectionPost**](AccessControlNATAccessSectionApi.md#ShowAccessSectionPost) | **Post** /show-access-section | 


# **AddAccessSectionPost**
> AccessSectionReply AddAccessSectionPost(ctx, accessSectionRequestNew)


Create new object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessSectionRequestNew** | [**AccessSectionRequestNew**](AccessSectionRequestNew.md)|  | 

### Return type

[**AccessSectionReply**](AccessSectionReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAccessSectionPost**
> ApiOkReply DeleteAccessSectionPost(ctx, accessSectionIdentifierRequest)


Delete existing object using object name or uid.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessSectionIdentifierRequest** | [**AccessSectionIdentifierRequest**](AccessSectionIdentifierRequest.md)|  | 

### Return type

[**ApiOkReply**](ApiOkReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetAccessSectionPost**
> AccessSectionReply SetAccessSectionPost(ctx, accessSectionRequestEdit)


Edit existing object using object name or uid.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessSectionRequestEdit** | [**AccessSectionRequestEdit**](AccessSectionRequestEdit.md)|  | 

### Return type

[**AccessSectionReply**](AccessSectionReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowAccessSectionPost**
> AccessSectionReply ShowAccessSectionPost(ctx, accessSectionIdentifierRequest)


Retrieve existing object using object name or uid.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessSectionIdentifierRequest** | [**AccessSectionIdentifierRequest**](AccessSectionIdentifierRequest.md)|  | 

### Return type

[**AccessSectionReply**](AccessSectionReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

