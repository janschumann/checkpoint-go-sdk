# \NetworkObjectsSimpleGatewayApi

All URIs are relative to *https://virtserver.swaggerhub.com/Schumann-IT/checkpoint-management-api/v1.5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddSimpleGatewayPost**](NetworkObjectsSimpleGatewayApi.md#AddSimpleGatewayPost) | **Post** /add-simple-gateway | 
[**DeleteSimpleGatewayPost**](NetworkObjectsSimpleGatewayApi.md#DeleteSimpleGatewayPost) | **Post** /delete-simple-gateway | 
[**SetSimpleGatewayPost**](NetworkObjectsSimpleGatewayApi.md#SetSimpleGatewayPost) | **Post** /set-simple-gateway | 
[**ShowSimpleGatewayPost**](NetworkObjectsSimpleGatewayApi.md#ShowSimpleGatewayPost) | **Post** /show-simple-gateway | 
[**ShowSimpleGatewaysPost**](NetworkObjectsSimpleGatewayApi.md#ShowSimpleGatewaysPost) | **Post** /show-simple-gateways | 


# **AddSimpleGatewayPost**
> GatewayCkpReply AddSimpleGatewayPost(ctx, gatewayCkpRequestNew)


Create new object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gatewayCkpRequestNew** | [**GatewayCkpRequestNew**](GatewayCkpRequestNew.md)|  | 

### Return type

[**GatewayCkpReply**](GatewayCkpReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSimpleGatewayPost**
> ApiOkReply DeleteSimpleGatewayPost(ctx, apiVisualCPObjectIdentifierRequestDelete)


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

# **SetSimpleGatewayPost**
> GatewayCkpReply SetSimpleGatewayPost(ctx, gatewayCkpRequestEdit)


Edit existing object using object name or uid.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gatewayCkpRequestEdit** | [**GatewayCkpRequestEdit**](GatewayCkpRequestEdit.md)|  | 

### Return type

[**GatewayCkpReply**](GatewayCkpReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowSimpleGatewayPost**
> GatewayCkpReply ShowSimpleGatewayPost(ctx, apiVisualCPObjectIdentifierRequestShow)


Retrieve existing object using object name or uid.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiVisualCPObjectIdentifierRequestShow** | [**ApiVisualCpObjectIdentifierRequestShow**](ApiVisualCpObjectIdentifierRequestShow.md)|  | 

### Return type

[**GatewayCkpReply**](GatewayCkpReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowSimpleGatewaysPost**
> ApiQueryObjectReply ShowSimpleGatewaysPost(ctx, objectInGroupQueryRequest)


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

