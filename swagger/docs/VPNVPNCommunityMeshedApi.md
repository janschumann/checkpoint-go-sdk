# \VPNVPNCommunityMeshedApi

All URIs are relative to *https://virtserver.swaggerhub.com/Schumann-IT/checkpoint-management-api/v1.5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddVpnCommunityMeshedPost**](VPNVPNCommunityMeshedApi.md#AddVpnCommunityMeshedPost) | **Post** /add-vpn-community-meshed | 
[**DeleteVpnCommunityMeshedPost**](VPNVPNCommunityMeshedApi.md#DeleteVpnCommunityMeshedPost) | **Post** /delete-vpn-community-meshed | 
[**SetVpnCommunityMeshedPost**](VPNVPNCommunityMeshedApi.md#SetVpnCommunityMeshedPost) | **Post** /set-vpn-community-meshed | 
[**ShowVpnCommunitiesMeshedPost**](VPNVPNCommunityMeshedApi.md#ShowVpnCommunitiesMeshedPost) | **Post** /show-vpn-communities-meshed | 
[**ShowVpnCommunityMeshedPost**](VPNVPNCommunityMeshedApi.md#ShowVpnCommunityMeshedPost) | **Post** /show-vpn-community-meshed | 


# **AddVpnCommunityMeshedPost**
> VpnMeshedCommunityReply AddVpnCommunityMeshedPost(ctx, vpnMeshedCommunityRequestNew)


Create new object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vpnMeshedCommunityRequestNew** | [**VpnMeshedCommunityRequestNew**](VpnMeshedCommunityRequestNew.md)|  | 

### Return type

[**VpnMeshedCommunityReply**](VpnMeshedCommunityReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteVpnCommunityMeshedPost**
> ApiOkReply DeleteVpnCommunityMeshedPost(ctx, apiVisualCPObjectIdentifierRequestDelete)


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

# **SetVpnCommunityMeshedPost**
> VpnMeshedCommunityReply SetVpnCommunityMeshedPost(ctx, vpnMeshedCommunityRequestEdit)


Edit existing object using object name or uid.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vpnMeshedCommunityRequestEdit** | [**VpnMeshedCommunityRequestEdit**](VpnMeshedCommunityRequestEdit.md)|  | 

### Return type

[**VpnMeshedCommunityReply**](VpnMeshedCommunityReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowVpnCommunitiesMeshedPost**
> ApiQueryObjectReply ShowVpnCommunitiesMeshedPost(ctx, apiQueryRequest)


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

# **ShowVpnCommunityMeshedPost**
> VpnMeshedCommunityReply ShowVpnCommunityMeshedPost(ctx, apiVisualCPObjectIdentifierRequestShow)


Retrieve existing object using object name or uid.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiVisualCPObjectIdentifierRequestShow** | [**ApiVisualCpObjectIdentifierRequestShow**](ApiVisualCpObjectIdentifierRequestShow.md)|  | 

### Return type

[**VpnMeshedCommunityReply**](VpnMeshedCommunityReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

