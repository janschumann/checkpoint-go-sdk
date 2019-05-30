# \MiscAdministratorApi

All URIs are relative to *https://virtserver.swaggerhub.com/Schumann-IT/checkpoint-management-api/v1.5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAdministratorPost**](MiscAdministratorApi.md#AddAdministratorPost) | **Post** /add-administrator | 
[**DeleteAdministratorPost**](MiscAdministratorApi.md#DeleteAdministratorPost) | **Post** /delete-administrator | 
[**SetAdministratorPost**](MiscAdministratorApi.md#SetAdministratorPost) | **Post** /set-administrator | 
[**ShowAdministratorPost**](MiscAdministratorApi.md#ShowAdministratorPost) | **Post** /show-administrator | 
[**ShowAdministratorsPost**](MiscAdministratorApi.md#ShowAdministratorsPost) | **Post** /show-administrators | 
[**UnlockAdministratorPost**](MiscAdministratorApi.md#UnlockAdministratorPost) | **Post** /unlock-administrator | 


# **AddAdministratorPost**
> AdministratorReply AddAdministratorPost(ctx, administratorRequestNew)


Create new administrator.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **administratorRequestNew** | [**AdministratorRequestNew**](AdministratorRequestNew.md)|  | 

### Return type

[**AdministratorReply**](AdministratorReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAdministratorPost**
> ApiOkReply DeleteAdministratorPost(ctx, apiVisualCPObjectIdentifierRequestDelete)


Delete existing administrator using object name or uid.

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

# **SetAdministratorPost**
> AdministratorReply SetAdministratorPost(ctx, administratorRequestEdit)


Edit existing administrator using object name or uid.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **administratorRequestEdit** | [**AdministratorRequestEdit**](AdministratorRequestEdit.md)|  | 

### Return type

[**AdministratorReply**](AdministratorReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowAdministratorPost**
> AdministratorReply ShowAdministratorPost(ctx, apiVisualCPObjectIdentifierRequestShow)


Retrieve existing administrator using object name or uid.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiVisualCPObjectIdentifierRequestShow** | [**ApiVisualCpObjectIdentifierRequestShow**](ApiVisualCpObjectIdentifierRequestShow.md)|  | 

### Return type

[**AdministratorReply**](AdministratorReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowAdministratorsPost**
> ApiQueryObjectReply ShowAdministratorsPost(ctx, apiQueryRequest)


Retrieve all administrators.

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

# **UnlockAdministratorPost**
> ApiOkReply UnlockAdministratorPost(ctx, apiVisualCPObjectIdentifierRequestShow)


Unlock administrator.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiVisualCPObjectIdentifierRequestShow** | [**ApiVisualCpObjectIdentifierRequestShow**](ApiVisualCpObjectIdentifierRequestShow.md)|  | 

### Return type

[**ApiOkReply**](ApiOkReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

