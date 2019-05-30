# \MultiDomainDomainApi

All URIs are relative to *https://virtserver.swaggerhub.com/Schumann-IT/checkpoint-management-api/v1.5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddDomainPost**](MultiDomainDomainApi.md#AddDomainPost) | **Post** /add-domain | 
[**DeleteDomainPost**](MultiDomainDomainApi.md#DeleteDomainPost) | **Post** /delete-domain | 
[**SetDomainPost**](MultiDomainDomainApi.md#SetDomainPost) | **Post** /set-domain | 
[**ShowDomainPost**](MultiDomainDomainApi.md#ShowDomainPost) | **Post** /show-domain | 
[**ShowDomainsPost**](MultiDomainDomainApi.md#ShowDomainsPost) | **Post** /show-domains | 


# **AddDomainPost**
> ApiTasksReply AddDomainPost(ctx, localDomainRequestNew)


Create a new domain in a Multi-Domain-Management environment. In order to allow administrators to connect to this domain using SmartConsole, use add-trusted-client command.<br> Note: This operation is not part of session and will take effect immediately.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **localDomainRequestNew** | [**LocalDomainRequestNew**](LocalDomainRequestNew.md)|  | 

### Return type

[**ApiTasksReply**](ApiTasksReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDomainPost**
> ApiTaskReply DeleteDomainPost(ctx, apiVisualCPObjectIdentifierRequestDelete)


Delete existing object using object name or uid.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiVisualCPObjectIdentifierRequestDelete** | [**ApiVisualCpObjectIdentifierRequestDelete**](ApiVisualCpObjectIdentifierRequestDelete.md)|  | 

### Return type

[**ApiTaskReply**](ApiTaskReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetDomainPost**
> LocalDomainEditReply SetDomainPost(ctx, localDomainRequestEdit)


Edit domain object using domain name or UID. When the list of domain servers is edited, the command is handled asynchronously. A list of task identifiers is returned to a user. In this case, the changes to the domain object are done in a public session and so should not be published. If the domain is changed in other parameters than the domain servers, i.e.: comments, color or tags, such changes are done in the user's private session and therefore should be published. In this case, the returned command output is similar to the one of 'show-domain'.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **localDomainRequestEdit** | [**LocalDomainRequestEdit**](LocalDomainRequestEdit.md)|  | 

### Return type

[**LocalDomainEditReply**](LocalDomainEditReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowDomainPost**
> LocalDomainReply ShowDomainPost(ctx, apiVisualCPObjectIdentifierRequestShow)


Retrieve existing object using object name or uid.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiVisualCPObjectIdentifierRequestShow** | [**ApiVisualCpObjectIdentifierRequestShow**](ApiVisualCpObjectIdentifierRequestShow.md)|  | 

### Return type

[**LocalDomainReply**](LocalDomainReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowDomainsPost**
> ApiQueryObjectReply ShowDomainsPost(ctx, apiQueryRequest)


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

