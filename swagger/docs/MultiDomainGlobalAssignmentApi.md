# \MultiDomainGlobalAssignmentApi

All URIs are relative to *https://virtserver.swaggerhub.com/Schumann-IT/checkpoint-management-api/v1.5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddGlobalAssignmentPost**](MultiDomainGlobalAssignmentApi.md#AddGlobalAssignmentPost) | **Post** /add-global-assignment | 
[**AssignGlobalAssignmentPost**](MultiDomainGlobalAssignmentApi.md#AssignGlobalAssignmentPost) | **Post** /assign-global-assignment | 
[**DeleteGlobalAssignmentPost**](MultiDomainGlobalAssignmentApi.md#DeleteGlobalAssignmentPost) | **Post** /delete-global-assignment | 
[**SetGlobalAssignmentPost**](MultiDomainGlobalAssignmentApi.md#SetGlobalAssignmentPost) | **Post** /set-global-assignment | 
[**ShowGlobalAssignmentPost**](MultiDomainGlobalAssignmentApi.md#ShowGlobalAssignmentPost) | **Post** /show-global-assignment | 
[**ShowGlobalAssignmentsPost**](MultiDomainGlobalAssignmentApi.md#ShowGlobalAssignmentsPost) | **Post** /show-global-assignments | 


# **AddGlobalAssignmentPost**
> GlobalAssignmentReply AddGlobalAssignmentPost(ctx, globalAssignmentRequestNew)


Create new object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **globalAssignmentRequestNew** | [**GlobalAssignmentRequestNew**](GlobalAssignmentRequestNew.md)|  | 

### Return type

[**GlobalAssignmentReply**](GlobalAssignmentReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssignGlobalAssignmentPost**
> AssignGlobalPolicyReply AssignGlobalAssignmentPost(ctx, assignGlobalPolicyRequest)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **assignGlobalPolicyRequest** | [**AssignGlobalPolicyRequest**](AssignGlobalPolicyRequest.md)|  | 

### Return type

[**AssignGlobalPolicyReply**](AssignGlobalPolicyReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteGlobalAssignmentPost**
> ApiTaskReply DeleteGlobalAssignmentPost(ctx, globalAssignmentIdentifierRequest)


Delete existing object using object name or uid.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **globalAssignmentIdentifierRequest** | [**GlobalAssignmentIdentifierRequest**](GlobalAssignmentIdentifierRequest.md)|  | 

### Return type

[**ApiTaskReply**](ApiTaskReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetGlobalAssignmentPost**
> GlobalAssignmentReply SetGlobalAssignmentPost(ctx, globalAssignmentRequestEdit)


Edit existing object using object name or uid.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **globalAssignmentRequestEdit** | [**GlobalAssignmentRequestEdit**](GlobalAssignmentRequestEdit.md)|  | 

### Return type

[**GlobalAssignmentReply**](GlobalAssignmentReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowGlobalAssignmentPost**
> GlobalAssignmentReply ShowGlobalAssignmentPost(ctx, globalAssignmentIdentifierRequest)


Retrieve existing object using object name or uid.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **globalAssignmentIdentifierRequest** | [**GlobalAssignmentIdentifierRequest**](GlobalAssignmentIdentifierRequest.md)|  | 

### Return type

[**GlobalAssignmentReply**](GlobalAssignmentReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowGlobalAssignmentsPost**
> ApiQueryObjectReply ShowGlobalAssignmentsPost(ctx, apiQueryRequest)


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

