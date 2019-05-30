# \UpdatableObjectsUpdatableObjectsRepositoryApi

All URIs are relative to *https://virtserver.swaggerhub.com/Schumann-IT/checkpoint-management-api/v1.5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ShowUpdatableObjectsRepositoryContentPost**](UpdatableObjectsUpdatableObjectsRepositoryApi.md#ShowUpdatableObjectsRepositoryContentPost) | **Post** /show-updatable-objects-repository-content | 
[**UpdateUpdatableObjectsRepositoryContentPost**](UpdatableObjectsUpdatableObjectsRepositoryApi.md#UpdateUpdatableObjectsRepositoryContentPost) | **Post** /update-updatable-objects-repository-content | 


# **ShowUpdatableObjectsRepositoryContentPost**
> UpdatableObjectsRepositoryContentReply ShowUpdatableObjectsRepositoryContentPost(ctx, updatableObjectsRepositoryContentRequest)


Shows the content of the available updatable objects from the Check Point User Center.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **updatableObjectsRepositoryContentRequest** | [**UpdatableObjectsRepositoryContentRequest**](UpdatableObjectsRepositoryContentRequest.md)|  | 

### Return type

[**UpdatableObjectsRepositoryContentReply**](UpdatableObjectsRepositoryContentReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateUpdatableObjectsRepositoryContentPost**
> ApiTaskReply UpdateUpdatableObjectsRepositoryContentPost(ctx, emptyRequest)


Updates the content of the Updatable Objects repository from the Check Point User Center.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **emptyRequest** | [**EmptyRequest**](EmptyRequest.md)|  | 

### Return type

[**ApiTaskReply**](ApiTaskReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

