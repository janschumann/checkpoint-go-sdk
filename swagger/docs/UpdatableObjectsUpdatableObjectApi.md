# \UpdatableObjectsUpdatableObjectApi

All URIs are relative to *https://virtserver.swaggerhub.com/Schumann-IT/checkpoint-management-api/v1.5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddUpdatableObjectPost**](UpdatableObjectsUpdatableObjectApi.md#AddUpdatableObjectPost) | **Post** /add-updatable-object | 
[**DeleteUpdatableObjectPost**](UpdatableObjectsUpdatableObjectApi.md#DeleteUpdatableObjectPost) | **Post** /delete-updatable-object | 
[**ShowUpdatableObjectPost**](UpdatableObjectsUpdatableObjectApi.md#ShowUpdatableObjectPost) | **Post** /show-updatable-object | 
[**ShowUpdatableObjectsPost**](UpdatableObjectsUpdatableObjectApi.md#ShowUpdatableObjectsPost) | **Post** /show-updatable-objects | 


# **AddUpdatableObjectPost**
> UpdatableObjectReply AddUpdatableObjectPost(ctx, updatableObjectRequestNew)


Import an updatable object from the repository to the management server. This operation takes effect immediately and doesn't require publishing.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **updatableObjectRequestNew** | [**UpdatableObjectRequestNew**](UpdatableObjectRequestNew.md)|  | 

### Return type

[**UpdatableObjectReply**](UpdatableObjectReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteUpdatableObjectPost**
> ApiOkReply DeleteUpdatableObjectPost(ctx, apiVisualCPObjectIdentifierRequestDelete)


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

# **ShowUpdatableObjectPost**
> UpdatableObjectReply ShowUpdatableObjectPost(ctx, apiVisualCPObjectIdentifierRequestShow)


Retrieves an existing Updatable Object from the Management server.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiVisualCPObjectIdentifierRequestShow** | [**ApiVisualCpObjectIdentifierRequestShow**](ApiVisualCpObjectIdentifierRequestShow.md)|  | 

### Return type

[**UpdatableObjectReply**](UpdatableObjectReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowUpdatableObjectsPost**
> ApiQueryObjectReply ShowUpdatableObjectsPost(ctx, objectInGroupQueryRequest)


Retrieves all Updatable Objects that were imported to the Management Server.

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

