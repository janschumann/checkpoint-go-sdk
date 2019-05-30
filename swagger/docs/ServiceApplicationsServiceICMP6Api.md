# \ServiceApplicationsServiceICMP6Api

All URIs are relative to *https://virtserver.swaggerhub.com/Schumann-IT/checkpoint-management-api/v1.5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddServiceIcmp6Post**](ServiceApplicationsServiceICMP6Api.md#AddServiceIcmp6Post) | **Post** /add-service-icmp6 | 
[**DeleteServiceIcmp6Post**](ServiceApplicationsServiceICMP6Api.md#DeleteServiceIcmp6Post) | **Post** /delete-service-icmp6 | 
[**SetServiceIcmp6Post**](ServiceApplicationsServiceICMP6Api.md#SetServiceIcmp6Post) | **Post** /set-service-icmp6 | 
[**ShowServiceIcmp6Post**](ServiceApplicationsServiceICMP6Api.md#ShowServiceIcmp6Post) | **Post** /show-service-icmp6 | 
[**ShowServicesIcmp6Post**](ServiceApplicationsServiceICMP6Api.md#ShowServicesIcmp6Post) | **Post** /show-services-icmp6 | 


# **AddServiceIcmp6Post**
> ServiceIcmp6Reply AddServiceIcmp6Post(ctx, serviceIcmp6RequestNew)


Create new object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceIcmp6RequestNew** | [**ServiceIcmp6RequestNew**](ServiceIcmp6RequestNew.md)|  | 

### Return type

[**ServiceIcmp6Reply**](ServiceIcmp6Reply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteServiceIcmp6Post**
> ApiOkReply DeleteServiceIcmp6Post(ctx, apiVisualCPObjectIdentifierRequestDelete)


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

# **SetServiceIcmp6Post**
> ServiceIcmp6Reply SetServiceIcmp6Post(ctx, serviceIcmp6RequestEdit)


Edit existing object using object name or uid.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceIcmp6RequestEdit** | [**ServiceIcmp6RequestEdit**](ServiceIcmp6RequestEdit.md)|  | 

### Return type

[**ServiceIcmp6Reply**](ServiceIcmp6Reply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowServiceIcmp6Post**
> ServiceIcmp6Reply ShowServiceIcmp6Post(ctx, apiVisualCPObjectIdentifierRequestShow)


Retrieve existing object using object name or uid.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiVisualCPObjectIdentifierRequestShow** | [**ApiVisualCpObjectIdentifierRequestShow**](ApiVisualCpObjectIdentifierRequestShow.md)|  | 

### Return type

[**ServiceIcmp6Reply**](ServiceIcmp6Reply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowServicesIcmp6Post**
> ApiQueryObjectReply ShowServicesIcmp6Post(ctx, objectInGroupQueryRequest)


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

