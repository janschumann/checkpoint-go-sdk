# \VPNVPNCommunityStarApi

All URIs are relative to *https://virtserver.swaggerhub.com/Schumann-IT/checkpoint-management-api/v1.5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddVpnCommunityStarPost**](VPNVPNCommunityStarApi.md#AddVpnCommunityStarPost) | **Post** /add-vpn-community-star | 
[**DeleteVpnCommunityStarPost**](VPNVPNCommunityStarApi.md#DeleteVpnCommunityStarPost) | **Post** /delete-vpn-community-star | 
[**SetVpnCommunityStarPost**](VPNVPNCommunityStarApi.md#SetVpnCommunityStarPost) | **Post** /set-vpn-community-star | 
[**ShowVpnCommunitiesStarPost**](VPNVPNCommunityStarApi.md#ShowVpnCommunitiesStarPost) | **Post** /show-vpn-communities-star | 
[**ShowVpnCommunityStarPost**](VPNVPNCommunityStarApi.md#ShowVpnCommunityStarPost) | **Post** /show-vpn-community-star | 


# **AddVpnCommunityStarPost**
> VpnStarCommunityReply AddVpnCommunityStarPost(ctx, vpnStarCommunityRequestNew)


Create new object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vpnStarCommunityRequestNew** | [**VpnStarCommunityRequestNew**](VpnStarCommunityRequestNew.md)|  | 

### Return type

[**VpnStarCommunityReply**](VpnStarCommunityReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteVpnCommunityStarPost**
> ApiOkReply DeleteVpnCommunityStarPost(ctx, apiVisualCPObjectIdentifierRequestDelete)


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

# **SetVpnCommunityStarPost**
> VpnStarCommunityReply SetVpnCommunityStarPost(ctx, vpnStarCommunityRequestEdit)


Edit existing object using object name or uid.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vpnStarCommunityRequestEdit** | [**VpnStarCommunityRequestEdit**](VpnStarCommunityRequestEdit.md)|  | 

### Return type

[**VpnStarCommunityReply**](VpnStarCommunityReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowVpnCommunitiesStarPost**
> ApiQueryObjectReply ShowVpnCommunitiesStarPost(ctx, apiQueryRequest)


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

# **ShowVpnCommunityStarPost**
> VpnStarCommunityReply ShowVpnCommunityStarPost(ctx, apiVisualCPObjectIdentifierRequestShow)


Retrieve existing object using object name or uid.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiVisualCPObjectIdentifierRequestShow** | [**ApiVisualCpObjectIdentifierRequestShow**](ApiVisualCpObjectIdentifierRequestShow.md)|  | 

### Return type

[**VpnStarCommunityReply**](VpnStarCommunityReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

