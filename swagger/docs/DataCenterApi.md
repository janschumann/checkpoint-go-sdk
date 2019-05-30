# \DataCenterApi

All URIs are relative to *https://virtserver.swaggerhub.com/Schumann-IT/checkpoint-management-api/v1.5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ShowDataCenterContentPost**](DataCenterApi.md#ShowDataCenterContentPost) | **Post** /show-data-center-content | 


# **ShowDataCenterContentPost**
> DataCenterContentQueryReply ShowDataCenterContentPost(ctx, dataCenterContentRequest)


Retrieve Data Center Objects from a Data Center.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dataCenterContentRequest** | [**DataCenterContentRequest**](DataCenterContentRequest.md)|  | 

### Return type

[**DataCenterContentQueryReply**](DataCenterContentQueryReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

