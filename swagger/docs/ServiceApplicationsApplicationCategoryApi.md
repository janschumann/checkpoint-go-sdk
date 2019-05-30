# \ServiceApplicationsApplicationCategoryApi

All URIs are relative to *https://virtserver.swaggerhub.com/Schumann-IT/checkpoint-management-api/v1.5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddApplicationSiteCategoryPost**](ServiceApplicationsApplicationCategoryApi.md#AddApplicationSiteCategoryPost) | **Post** /add-application-site-category | 
[**DeleteApplicationSiteCategoryPost**](ServiceApplicationsApplicationCategoryApi.md#DeleteApplicationSiteCategoryPost) | **Post** /delete-application-site-category | 
[**SetApplicationSiteCategoryPost**](ServiceApplicationsApplicationCategoryApi.md#SetApplicationSiteCategoryPost) | **Post** /set-application-site-category | 
[**ShowApplicationSiteCategoriesPost**](ServiceApplicationsApplicationCategoryApi.md#ShowApplicationSiteCategoriesPost) | **Post** /show-application-site-categories | 
[**ShowApplicationSiteCategoryPost**](ServiceApplicationsApplicationCategoryApi.md#ShowApplicationSiteCategoryPost) | **Post** /show-application-site-category | 


# **AddApplicationSiteCategoryPost**
> ApplicationSiteCategoryReply AddApplicationSiteCategoryPost(ctx, applicationSiteCategoryRequestNew)


Create new object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationSiteCategoryRequestNew** | [**ApplicationSiteCategoryRequestNew**](ApplicationSiteCategoryRequestNew.md)|  | 

### Return type

[**ApplicationSiteCategoryReply**](ApplicationSiteCategoryReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteApplicationSiteCategoryPost**
> ApiOkReply DeleteApplicationSiteCategoryPost(ctx, apiVisualCPObjectIdentifierRequestDelete)


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

# **SetApplicationSiteCategoryPost**
> ApplicationSiteCategoryReply SetApplicationSiteCategoryPost(ctx, applicationSiteCategoryRequestEdit)


Edit existing object using object name or uid.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationSiteCategoryRequestEdit** | [**ApplicationSiteCategoryRequestEdit**](ApplicationSiteCategoryRequestEdit.md)|  | 

### Return type

[**ApplicationSiteCategoryReply**](ApplicationSiteCategoryReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowApplicationSiteCategoriesPost**
> ApiQueryObjectReply ShowApplicationSiteCategoriesPost(ctx, apiQueryRequest)


Retrieve all objects.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiQueryRequest** | [**ApiQueryRequest**](ApiQueryRequest.md)|  | 

### Return type

[**ApiQueryObjectReply**](ApiQueryObjectReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowApplicationSiteCategoryPost**
> ApplicationSiteCategoryReply ShowApplicationSiteCategoryPost(ctx, apiVisualCPObjectIdentifierRequestShow)


Retrieve existing object using object name or uid.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiVisualCPObjectIdentifierRequestShow** | [**ApiVisualCpObjectIdentifierRequestShow**](ApiVisualCpObjectIdentifierRequestShow.md)|  | 

### Return type

[**ApplicationSiteCategoryReply**](ApplicationSiteCategoryReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

