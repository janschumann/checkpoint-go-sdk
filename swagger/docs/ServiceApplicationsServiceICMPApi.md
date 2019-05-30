# \ServiceApplicationsServiceICMPApi

All URIs are relative to *https://virtserver.swaggerhub.com/Schumann-IT/checkpoint-management-api/v1.5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddServiceIcmpPost**](ServiceApplicationsServiceICMPApi.md#AddServiceIcmpPost) | **Post** /add-service-icmp | 
[**DeleteServiceIcmpPost**](ServiceApplicationsServiceICMPApi.md#DeleteServiceIcmpPost) | **Post** /delete-service-icmp | 
[**SetServiceIcmpPost**](ServiceApplicationsServiceICMPApi.md#SetServiceIcmpPost) | **Post** /set-service-icmp | 
[**ShowServiceIcmpPost**](ServiceApplicationsServiceICMPApi.md#ShowServiceIcmpPost) | **Post** /show-service-icmp | 
[**ShowServicesIcmpPost**](ServiceApplicationsServiceICMPApi.md#ShowServicesIcmpPost) | **Post** /show-services-icmp | 


# **AddServiceIcmpPost**
> ServiceIcmpReply AddServiceIcmpPost(ctx, serviceIcmpRequestNew)


Create new object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceIcmpRequestNew** | [**ServiceIcmpRequestNew**](ServiceIcmpRequestNew.md)|  | 

### Return type

[**ServiceIcmpReply**](ServiceIcmpReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteServiceIcmpPost**
> ApiOkReply DeleteServiceIcmpPost(ctx, apiVisualCPObjectIdentifierRequestDelete)


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

# **SetServiceIcmpPost**
> ServiceIcmpReply SetServiceIcmpPost(ctx, serviceIcmpRequestEdit)


Edit existing object using object name or uid.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceIcmpRequestEdit** | [**ServiceIcmpRequestEdit**](ServiceIcmpRequestEdit.md)|  | 

### Return type

[**ServiceIcmpReply**](ServiceIcmpReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowServiceIcmpPost**
> ServiceIcmpReply ShowServiceIcmpPost(ctx, apiVisualCPObjectIdentifierRequestShow)


Retrieve existing object using object name or uid.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiVisualCPObjectIdentifierRequestShow** | [**ApiVisualCpObjectIdentifierRequestShow**](ApiVisualCpObjectIdentifierRequestShow.md)|  | 

### Return type

[**ServiceIcmpReply**](ServiceIcmpReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowServicesIcmpPost**
> ApiQueryObjectReply ShowServicesIcmpPost(ctx, objectInGroupQueryRequest)


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

