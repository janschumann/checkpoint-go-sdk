# \UnknownApi

All URIs are relative to *https://virtserver.swaggerhub.com/Schumann-IT/checkpoint-management-api/v1.5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddScadaApplicationPost**](UnknownApi.md#AddScadaApplicationPost) | **Post** /add-scada-application | 
[**AutoCompletePost**](UnknownApi.md#AutoCompletePost) | **Post** /auto-complete | 
[**DeleteScadaApplicationPost**](UnknownApi.md#DeleteScadaApplicationPost) | **Post** /delete-scada-application | 
[**GetInterfacesSyncPost**](UnknownApi.md#GetInterfacesSyncPost) | **Post** /get-interfaces-sync | 
[**MakeServerActivePost**](UnknownApi.md#MakeServerActivePost) | **Post** /make-server-active | 
[**ResultLinkPost**](UnknownApi.md#ResultLinkPost) | **Post** /result-link | 
[**RunInitInterfacesPost**](UnknownApi.md#RunInitInterfacesPost) | **Post** /run-init-interfaces | 
[**SetScadaApplicationPost**](UnknownApi.md#SetScadaApplicationPost) | **Post** /set-scada-application | 
[**ShowApiStatusPost**](UnknownApi.md#ShowApiStatusPost) | **Post** /show-api-status | 
[**ShowInternalVersionPost**](UnknownApi.md#ShowInternalVersionPost) | **Post** /show-internal-version | 
[**ShowRuleCandidatesPost**](UnknownApi.md#ShowRuleCandidatesPost) | **Post** /show-rule-candidates | 
[**ShowScadaApplicationPost**](UnknownApi.md#ShowScadaApplicationPost) | **Post** /show-scada-application | 
[**ShowScadaApplicationsPost**](UnknownApi.md#ShowScadaApplicationsPost) | **Post** /show-scada-applications | 
[**ShowVersionPost**](UnknownApi.md#ShowVersionPost) | **Post** /show-version | 


# **AddScadaApplicationPost**
> ScadaApplicationReply AddScadaApplicationPost(ctx, scadaApplicationRequestNew)


Create new object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **scadaApplicationRequestNew** | [**ScadaApplicationRequestNew**](ScadaApplicationRequestNew.md)|  | 

### Return type

[**ScadaApplicationReply**](ScadaApplicationReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AutoCompletePost**
> AutoCompleteReply AutoCompletePost(ctx, autoCompleteRequest)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **autoCompleteRequest** | [**AutoCompleteRequest**](AutoCompleteRequest.md)|  | 

### Return type

[**AutoCompleteReply**](AutoCompleteReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteScadaApplicationPost**
> ApiOkReply DeleteScadaApplicationPost(ctx, apiVisualCPObjectIdentifierRequestDelete)


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

# **GetInterfacesSyncPost**
> ApiOkReply GetInterfacesSyncPost(ctx, getInterfacesRequest)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **getInterfacesRequest** | [**GetInterfacesRequest**](GetInterfacesRequest.md)|  | 

### Return type

[**ApiOkReply**](ApiOkReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MakeServerActivePost**
> SetActiveReply MakeServerActivePost(ctx, setActiveRequest)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **setActiveRequest** | [**SetActiveRequest**](SetActiveRequest.md)|  | 

### Return type

[**SetActiveReply**](SetActiveReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResultLinkPost**
> WebApiResultLinkReply ResultLinkPost(ctx, webApiResultLinkRequest)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **webApiResultLinkRequest** | [**WebApiResultLinkRequest**](WebApiResultLinkRequest.md)|  | 

### Return type

[**WebApiResultLinkReply**](WebApiResultLinkReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RunInitInterfacesPost**
> RunInitInterfacesReply RunInitInterfacesPost(ctx, runInitInterfacesRequest)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **runInitInterfacesRequest** | [**RunInitInterfacesRequest**](RunInitInterfacesRequest.md)|  | 

### Return type

[**RunInitInterfacesReply**](RunInitInterfacesReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetScadaApplicationPost**
> ScadaApplicationReply SetScadaApplicationPost(ctx, scadaApplicationRequestEdit)


Edit existing object using object name or uid.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **scadaApplicationRequestEdit** | [**ScadaApplicationRequestEdit**](ScadaApplicationRequestEdit.md)|  | 

### Return type

[**ScadaApplicationReply**](ScadaApplicationReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowApiStatusPost**
> WebApiServerStatusReply ShowApiStatusPost(ctx, webApiServerStatusRequest)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **webApiServerStatusRequest** | [**WebApiServerStatusRequest**](WebApiServerStatusRequest.md)|  | 

### Return type

[**WebApiServerStatusReply**](WebApiServerStatusReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowInternalVersionPost**
> VersionInternalReply ShowInternalVersionPost(ctx, versionInternalRequest)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **versionInternalRequest** | [**VersionInternalRequest**](VersionInternalRequest.md)|  | 

### Return type

[**VersionInternalReply**](VersionInternalReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowRuleCandidatesPost**
> ShowRuleCandidateReply ShowRuleCandidatesPost(ctx, showRuleCandidateRequest)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **showRuleCandidateRequest** | [**ShowRuleCandidateRequest**](ShowRuleCandidateRequest.md)|  | 

### Return type

[**ShowRuleCandidateReply**](ShowRuleCandidateReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowScadaApplicationPost**
> ScadaApplicationReply ShowScadaApplicationPost(ctx, applicationSiteIdentifierRequestShow)


Retrieve existing object using object name or uid.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationSiteIdentifierRequestShow** | [**ApplicationSiteIdentifierRequestShow**](ApplicationSiteIdentifierRequestShow.md)|  | 

### Return type

[**ScadaApplicationReply**](ScadaApplicationReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowScadaApplicationsPost**
> ApiQueryObjectReply ShowScadaApplicationsPost(ctx, apiQueryRequest)


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

# **ShowVersionPost**
> VersionReply ShowVersionPost(ctx, versionRequest)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **versionRequest** | [**VersionRequest**](VersionRequest.md)|  | 

### Return type

[**VersionReply**](VersionReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

