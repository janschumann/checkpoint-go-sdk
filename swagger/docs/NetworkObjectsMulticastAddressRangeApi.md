# \NetworkObjectsMulticastAddressRangeApi

All URIs are relative to *https://virtserver.swaggerhub.com/Schumann-IT/checkpoint-management-api/v1.5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddMulticastAddressRangePost**](NetworkObjectsMulticastAddressRangeApi.md#AddMulticastAddressRangePost) | **Post** /add-multicast-address-range | 
[**DeleteMulticastAddressRangePost**](NetworkObjectsMulticastAddressRangeApi.md#DeleteMulticastAddressRangePost) | **Post** /delete-multicast-address-range | 
[**SetMulticastAddressRangePost**](NetworkObjectsMulticastAddressRangeApi.md#SetMulticastAddressRangePost) | **Post** /set-multicast-address-range | 
[**ShowMulticastAddressRangePost**](NetworkObjectsMulticastAddressRangeApi.md#ShowMulticastAddressRangePost) | **Post** /show-multicast-address-range | 
[**ShowMulticastAddressRangesPost**](NetworkObjectsMulticastAddressRangeApi.md#ShowMulticastAddressRangesPost) | **Post** /show-multicast-address-ranges | 


# **AddMulticastAddressRangePost**
> MultiCastAddressRangeReply AddMulticastAddressRangePost(ctx, multiCastAddressRangeRequestNew)


Create new object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **multiCastAddressRangeRequestNew** | [**MultiCastAddressRangeRequestNew**](MultiCastAddressRangeRequestNew.md)|  | 

### Return type

[**MultiCastAddressRangeReply**](MultiCastAddressRangeReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMulticastAddressRangePost**
> ApiOkReply DeleteMulticastAddressRangePost(ctx, apiVisualCPObjectIdentifierRequestDelete)


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

# **SetMulticastAddressRangePost**
> MultiCastAddressRangeReply SetMulticastAddressRangePost(ctx, multiCastAddressRangeRequestEdit)


Edit existing object using object name or uid.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **multiCastAddressRangeRequestEdit** | [**MultiCastAddressRangeRequestEdit**](MultiCastAddressRangeRequestEdit.md)|  | 

### Return type

[**MultiCastAddressRangeReply**](MultiCastAddressRangeReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowMulticastAddressRangePost**
> MultiCastAddressRangeReply ShowMulticastAddressRangePost(ctx, apiVisualCPObjectIdentifierRequestShow)


Retrieve existing object using object name or uid.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiVisualCPObjectIdentifierRequestShow** | [**ApiVisualCpObjectIdentifierRequestShow**](ApiVisualCpObjectIdentifierRequestShow.md)|  | 

### Return type

[**MultiCastAddressRangeReply**](MultiCastAddressRangeReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowMulticastAddressRangesPost**
> MulticastAddressRangeQueryReply ShowMulticastAddressRangesPost(ctx, objectInGroupQueryRequest)


Retrieve all objects.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectInGroupQueryRequest** | [**ObjectInGroupQueryRequest**](ObjectInGroupQueryRequest.md)|  | 

### Return type

[**MulticastAddressRangeQueryReply**](MulticastAddressRangeQueryReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

