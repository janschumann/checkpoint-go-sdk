# \ServiceApplicationsApplicationGroupApi

All URIs are relative to *https://virtserver.swaggerhub.com/Schumann-IT/checkpoint-management-api/v1.5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddApplicationSiteGroupPost**](ServiceApplicationsApplicationGroupApi.md#AddApplicationSiteGroupPost) | **Post** /add-application-site-group | 
[**DeleteApplicationSiteGroupPost**](ServiceApplicationsApplicationGroupApi.md#DeleteApplicationSiteGroupPost) | **Post** /delete-application-site-group | 
[**SetApplicationSiteGroupPost**](ServiceApplicationsApplicationGroupApi.md#SetApplicationSiteGroupPost) | **Post** /set-application-site-group | 
[**ShowApplicationSiteGroupPost**](ServiceApplicationsApplicationGroupApi.md#ShowApplicationSiteGroupPost) | **Post** /show-application-site-group | 
[**ShowApplicationSiteGroupsPost**](ServiceApplicationsApplicationGroupApi.md#ShowApplicationSiteGroupsPost) | **Post** /show-application-site-groups | 


# **AddApplicationSiteGroupPost**
> ApplicationSiteGroupReply AddApplicationSiteGroupPost(ctx, applicationSiteGroupRequestNew)


Create new object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationSiteGroupRequestNew** | [**ApplicationSiteGroupRequestNew**](ApplicationSiteGroupRequestNew.md)|  | 

### Return type

[**ApplicationSiteGroupReply**](ApplicationSiteGroupReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteApplicationSiteGroupPost**
> ApiOkReply DeleteApplicationSiteGroupPost(ctx, apiVisualCPObjectIdentifierRequestDelete)


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

# **SetApplicationSiteGroupPost**
> ApplicationSiteGroupReply SetApplicationSiteGroupPost(ctx, applicationSiteGroupRequestEdit)


Edit existing object using object name or uid.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationSiteGroupRequestEdit** | [**ApplicationSiteGroupRequestEdit**](ApplicationSiteGroupRequestEdit.md)|  | 

### Return type

[**ApplicationSiteGroupReply**](ApplicationSiteGroupReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowApplicationSiteGroupPost**
> ApplicationSiteGroupReply ShowApplicationSiteGroupPost(ctx, apiVisualCPObjectIdentifierRequestShow)


Retrieve existing object using object name or uid.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiVisualCPObjectIdentifierRequestShow** | [**ApiVisualCpObjectIdentifierRequestShow**](ApiVisualCpObjectIdentifierRequestShow.md)|  | 

### Return type

[**ApplicationSiteGroupReply**](ApplicationSiteGroupReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowApplicationSiteGroupsPost**
> ApiQueryObjectReply ShowApplicationSiteGroupsPost(ctx, objectInGroupWithMembersQueryRequest)


Retrieve all objects.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectInGroupWithMembersQueryRequest** | [**ObjectInGroupWithMembersQueryRequest**](ObjectInGroupWithMembersQueryRequest.md)|  | 

### Return type

[**ApiQueryObjectReply**](ApiQueryObjectReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

