# \ServiceApplicationsServiceDCERPCApi

All URIs are relative to *https://virtserver.swaggerhub.com/Schumann-IT/checkpoint-management-api/v1.5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddServiceDceRpcPost**](ServiceApplicationsServiceDCERPCApi.md#AddServiceDceRpcPost) | **Post** /add-service-dce-rpc | 
[**DeleteServiceDceRpcPost**](ServiceApplicationsServiceDCERPCApi.md#DeleteServiceDceRpcPost) | **Post** /delete-service-dce-rpc | 
[**SetServiceDceRpcPost**](ServiceApplicationsServiceDCERPCApi.md#SetServiceDceRpcPost) | **Post** /set-service-dce-rpc | 
[**ShowServiceDceRpcPost**](ServiceApplicationsServiceDCERPCApi.md#ShowServiceDceRpcPost) | **Post** /show-service-dce-rpc | 
[**ShowServicesDceRpcPost**](ServiceApplicationsServiceDCERPCApi.md#ShowServicesDceRpcPost) | **Post** /show-services-dce-rpc | 


# **AddServiceDceRpcPost**
> DcerpcServiceReply AddServiceDceRpcPost(ctx, dcerpcServiceRequestNew)


Create new object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dcerpcServiceRequestNew** | [**DcerpcServiceRequestNew**](DcerpcServiceRequestNew.md)|  | 

### Return type

[**DcerpcServiceReply**](DcerpcServiceReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteServiceDceRpcPost**
> ApiOkReply DeleteServiceDceRpcPost(ctx, apiVisualCPObjectIdentifierRequestDelete)


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

# **SetServiceDceRpcPost**
> DcerpcServiceReply SetServiceDceRpcPost(ctx, dcerpcServiceRequestEdit)


Edit existing object using object name or uid.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dcerpcServiceRequestEdit** | [**DcerpcServiceRequestEdit**](DcerpcServiceRequestEdit.md)|  | 

### Return type

[**DcerpcServiceReply**](DcerpcServiceReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowServiceDceRpcPost**
> DcerpcServiceReply ShowServiceDceRpcPost(ctx, apiVisualCPObjectIdentifierRequestShow)


Retrieve existing object using object name or uid.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiVisualCPObjectIdentifierRequestShow** | [**ApiVisualCpObjectIdentifierRequestShow**](ApiVisualCpObjectIdentifierRequestShow.md)|  | 

### Return type

[**DcerpcServiceReply**](DcerpcServiceReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowServicesDceRpcPost**
> ApiQueryObjectReply ShowServicesDceRpcPost(ctx, objectInGroupQueryRequest)


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

