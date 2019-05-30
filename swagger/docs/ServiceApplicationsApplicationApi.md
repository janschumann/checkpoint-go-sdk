# \ServiceApplicationsApplicationApi

All URIs are relative to *https://virtserver.swaggerhub.com/Schumann-IT/checkpoint-management-api/v1.5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddApplicationSitePost**](ServiceApplicationsApplicationApi.md#AddApplicationSitePost) | **Post** /add-application-site | 
[**DeleteApplicationSitePost**](ServiceApplicationsApplicationApi.md#DeleteApplicationSitePost) | **Post** /delete-application-site | 
[**SetApplicationSitePost**](ServiceApplicationsApplicationApi.md#SetApplicationSitePost) | **Post** /set-application-site | 
[**ShowApplicationSitePost**](ServiceApplicationsApplicationApi.md#ShowApplicationSitePost) | **Post** /show-application-site | 
[**ShowApplicationSitesPost**](ServiceApplicationsApplicationApi.md#ShowApplicationSitesPost) | **Post** /show-application-sites | 


# **AddApplicationSitePost**
> ApplicationSiteReply AddApplicationSitePost(ctx, applicationSiteRequestNew)


Creates new application site, which can be initialized with 'url-list' or 'application-signature' (not both of them).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationSiteRequestNew** | [**ApplicationSiteRequestNew**](ApplicationSiteRequestNew.md)|  | 

### Return type

[**ApplicationSiteReply**](ApplicationSiteReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteApplicationSitePost**
> ApiOkReply DeleteApplicationSitePost(ctx, apiVisualCPObjectIdentifierRequestDelete)


Delete existing object using object name or uid.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiVisualCPObjectIdentifierRequestDelete** | [**ApiVisualCpObjectIdentifierRequestDelete**](ApiVisualCpObjectIdentifierRequestDelete.md)|  | 

### Return type

[**ApiOkReply**](ApiOkReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetApplicationSitePost**
> ApplicationSiteReply SetApplicationSitePost(ctx, applicationSiteRequestEdit)


Edit existing application using object name or uid. It's impossible to set 'application-signature' when the application was initialized with 'url-list' and vice-verse.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationSiteRequestEdit** | [**ApplicationSiteRequestEdit**](ApplicationSiteRequestEdit.md)|  | 

### Return type

[**ApplicationSiteReply**](ApplicationSiteReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowApplicationSitePost**
> ApplicationSiteReply ShowApplicationSitePost(ctx, applicationSiteIdentifierRequestShow)


Retrieve existing object using object name or uid.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationSiteIdentifierRequestShow** | [**ApplicationSiteIdentifierRequestShow**](ApplicationSiteIdentifierRequestShow.md)|  | 

### Return type

[**ApplicationSiteReply**](ApplicationSiteReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowApplicationSitesPost**
> ApiQueryObjectReply ShowApplicationSitesPost(ctx, objectInGroupQueryRequest)


Retrieve all objects.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectInGroupQueryRequest** | [**ObjectInGroupQueryRequest**](ObjectInGroupQueryRequest.md)|  | 

### Return type

[**ApiQueryObjectReply**](ApiQueryObjectReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

