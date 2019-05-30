# \Misc_Api

All URIs are relative to *https://virtserver.swaggerhub.com/Schumann-IT/checkpoint-management-api/v1.5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExportPost**](Misc_Api.md#ExportPost) | **Post** /export | 
[**PutFilePost**](Misc_Api.md#PutFilePost) | **Post** /put-file | 
[**RunScriptPost**](Misc_Api.md#RunScriptPost) | **Post** /run-script | 
[**ShowApiVersionsPost**](Misc_Api.md#ShowApiVersionsPost) | **Post** /show-api-versions | 
[**ShowChangesPost**](Misc_Api.md#ShowChangesPost) | **Post** /show-changes | 
[**ShowCommandsPost**](Misc_Api.md#ShowCommandsPost) | **Post** /show-commands | 
[**ShowGatewaysAndServersPost**](Misc_Api.md#ShowGatewaysAndServersPost) | **Post** /show-gateways-and-servers | 
[**ShowObjectPost**](Misc_Api.md#ShowObjectPost) | **Post** /show-object | 
[**ShowObjectsPost**](Misc_Api.md#ShowObjectsPost) | **Post** /show-objects | 
[**ShowTaskPost**](Misc_Api.md#ShowTaskPost) | **Post** /show-task | 
[**ShowTasksPost**](Misc_Api.md#ShowTasksPost) | **Post** /show-tasks | 
[**ShowUnusedObjectsPost**](Misc_Api.md#ShowUnusedObjectsPost) | **Post** /show-unused-objects | 
[**ShowValidationsPost**](Misc_Api.md#ShowValidationsPost) | **Post** /show-validations | 
[**WhereUsedPost**](Misc_Api.md#WhereUsedPost) | **Post** /where-used | 


# **ExportPost**
> WebApiExportReply ExportPost(ctx, webApiExportRequest)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **webApiExportRequest** | [**WebApiExportRequest**](WebApiExportRequest.md)|  | 

### Return type

[**WebApiExportReply**](WebApiExportReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutFilePost**
> CdmCommandReply PutFilePost(ctx, putFileRequest)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **putFileRequest** | [**PutFileRequest**](PutFileRequest.md)|  | 

### Return type

[**CdmCommandReply**](CdmCommandReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RunScriptPost**
> CdmCommandReply RunScriptPost(ctx, runScriptRequest)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **runScriptRequest** | [**RunScriptRequest**](RunScriptRequest.md)|  | 

### Return type

[**CdmCommandReply**](CdmCommandReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowApiVersionsPost**
> ApiVersionsReply ShowApiVersionsPost(ctx, emptyRequest)


Shows all supported API versions and current API version (the latest one).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **emptyRequest** | [**EmptyRequest**](EmptyRequest.md)|  | 

### Return type

[**ApiVersionsReply**](ApiVersionsReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowChangesPost**
> DiffReplyTask ShowChangesPost(ctx, diffRequest)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **diffRequest** | [**DiffRequest**](DiffRequest.md)|  | 

### Return type

[**DiffReplyTask**](DiffReplyTask.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowCommandsPost**
> ShowCommandsReply ShowCommandsPost(ctx, showCommandsRequest)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **showCommandsRequest** | [**ShowCommandsRequest**](ShowCommandsRequest.md)|  | 

### Return type

[**ShowCommandsReply**](ShowCommandsReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowGatewaysAndServersPost**
> QueryGatewaysServersReply ShowGatewaysAndServersPost(ctx, apiQueryRequest)


Shows list of Gateways & Servers sorted by name.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiQueryRequest** | [**ApiQueryRequest**](ApiQueryRequest.md)|  | 

### Return type

[**QueryGatewaysServersReply**](QueryGatewaysServersReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowObjectPost**
> ShowObjectReply ShowObjectPost(ctx, showObjectRequest)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **showObjectRequest** | [**ShowObjectRequest**](ShowObjectRequest.md)|  | 

### Return type

[**ShowObjectReply**](ShowObjectReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowObjectsPost**
> QueryObjectsReply ShowObjectsPost(ctx, queryObjectsRequest)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **queryObjectsRequest** | [**QueryObjectsRequest**](QueryObjectsRequest.md)|  | 

### Return type

[**QueryObjectsReply**](QueryObjectsReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowTaskPost**
> TaskReply ShowTaskPost(ctx, taskRequest)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **taskRequest** | [**TaskRequest**](TaskRequest.md)|  | 

### Return type

[**TaskReply**](TaskReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowTasksPost**
> TaskQueryReply ShowTasksPost(ctx, taskQueryRequest)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **taskQueryRequest** | [**TaskQueryRequest**](TaskQueryRequest.md)|  | 

### Return type

[**TaskQueryReply**](TaskQueryReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowUnusedObjectsPost**
> ApiQueryObjectReply ShowUnusedObjectsPost(ctx, objectInGroupWithMembersQueryRequest)


Retrieve all unused objects.

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

# **ShowValidationsPost**
> ValidationsReply ShowValidationsPost(ctx, emptyRequest)


Show all validation incidents limited to 500.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **emptyRequest** | [**EmptyRequest**](EmptyRequest.md)|  | 

### Return type

[**ValidationsReply**](ValidationsReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WhereUsedPost**
> WhereUsedObjectReply WhereUsedPost(ctx, whereUsedObjectRequest)


Searches for usage of the target object in other objects and rules.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **whereUsedObjectRequest** | [**WhereUsedObjectRequest**](WhereUsedObjectRequest.md)|  | 

### Return type

[**WhereUsedObjectReply**](WhereUsedObjectReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

