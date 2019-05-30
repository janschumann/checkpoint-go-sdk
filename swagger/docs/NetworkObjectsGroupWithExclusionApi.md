# \NetworkObjectsGroupWithExclusionApi

All URIs are relative to *https://virtserver.swaggerhub.com/Schumann-IT/checkpoint-management-api/v1.5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddGroupWithExclusionPost**](NetworkObjectsGroupWithExclusionApi.md#AddGroupWithExclusionPost) | **Post** /add-group-with-exclusion | 
[**DeleteGroupWithExclusionPost**](NetworkObjectsGroupWithExclusionApi.md#DeleteGroupWithExclusionPost) | **Post** /delete-group-with-exclusion | 
[**SetGroupWithExclusionPost**](NetworkObjectsGroupWithExclusionApi.md#SetGroupWithExclusionPost) | **Post** /set-group-with-exclusion | 
[**ShowGroupWithExclusionPost**](NetworkObjectsGroupWithExclusionApi.md#ShowGroupWithExclusionPost) | **Post** /show-group-with-exclusion | 
[**ShowGroupsWithExclusionPost**](NetworkObjectsGroupWithExclusionApi.md#ShowGroupsWithExclusionPost) | **Post** /show-groups-with-exclusion | 


# **AddGroupWithExclusionPost**
> GroupWithExclusionReply AddGroupWithExclusionPost(ctx, groupWithExclusionRequestNew)


Create new object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupWithExclusionRequestNew** | [**GroupWithExclusionRequestNew**](GroupWithExclusionRequestNew.md)|  | 

### Return type

[**GroupWithExclusionReply**](GroupWithExclusionReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteGroupWithExclusionPost**
> ApiOkReply DeleteGroupWithExclusionPost(ctx, apiVisualCPObjectIdentifierRequestDelete)


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

# **SetGroupWithExclusionPost**
> GroupWithExclusionReply SetGroupWithExclusionPost(ctx, groupWithExclusionRequestEdit)


Edit existing object using object name or uid.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupWithExclusionRequestEdit** | [**GroupWithExclusionRequestEdit**](GroupWithExclusionRequestEdit.md)|  | 

### Return type

[**GroupWithExclusionReply**](GroupWithExclusionReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowGroupWithExclusionPost**
> GroupWithExclusionReply ShowGroupWithExclusionPost(ctx, groupWithExclusionRequestShow)


Retrieve existing object using object name or uid.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupWithExclusionRequestShow** | [**GroupWithExclusionRequestShow**](GroupWithExclusionRequestShow.md)|  | 

### Return type

[**GroupWithExclusionReply**](GroupWithExclusionReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowGroupsWithExclusionPost**
> ApiQueryObjectReply ShowGroupsWithExclusionPost(ctx, groupWithExclusionRequestQuery)


Retrieve all objects.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupWithExclusionRequestQuery** | [**GroupWithExclusionRequestQuery**](GroupWithExclusionRequestQuery.md)|  | 

### Return type

[**ApiQueryObjectReply**](ApiQueryObjectReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

