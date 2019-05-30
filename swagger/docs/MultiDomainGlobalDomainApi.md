# \MultiDomainGlobalDomainApi

All URIs are relative to *https://virtserver.swaggerhub.com/Schumann-IT/checkpoint-management-api/v1.5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SetGlobalDomainPost**](MultiDomainGlobalDomainApi.md#SetGlobalDomainPost) | **Post** /set-global-domain | 
[**ShowGlobalDomainPost**](MultiDomainGlobalDomainApi.md#ShowGlobalDomainPost) | **Post** /show-global-domain | 


# **SetGlobalDomainPost**
> GlobalDomainEditReply SetGlobalDomainPost(ctx, globalDomainRequestEdit)


Edit Global domain object using domain name or UID. When the list of Multi Domain Server is edited, the command is handled asynchronously. A list of task identifiers is returned to a user. In this case, the changes to the Global domain object are done in a public session and so should not be published. If the domain is changed in other parameters than the Multi Domain Servers, i.e.: comments, color or tags, such changes are done in the user's private session and therefore should be published. In this case, the returned command output is similar to the one of 'show-global-domain'.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **globalDomainRequestEdit** | [**GlobalDomainRequestEdit**](GlobalDomainRequestEdit.md)|  | 

### Return type

[**GlobalDomainEditReply**](GlobalDomainEditReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowGlobalDomainPost**
> GlobalDomainReply ShowGlobalDomainPost(ctx, apiVisualCPObjectIdentifierRequestShow)


Retrieve existing object using object name or uid.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiVisualCPObjectIdentifierRequestShow** | [**ApiVisualCpObjectIdentifierRequestShow**](ApiVisualCpObjectIdentifierRequestShow.md)|  | 

### Return type

[**GlobalDomainReply**](GlobalDomainReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

