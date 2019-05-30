# \ThreatPreventionIPSApi

All URIs are relative to *https://virtserver.swaggerhub.com/Schumann-IT/checkpoint-management-api/v1.5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RunIpsUpdatePost**](ThreatPreventionIPSApi.md#RunIpsUpdatePost) | **Post** /run-ips-update | 
[**SetIpsUpdateSchedulePost**](ThreatPreventionIPSApi.md#SetIpsUpdateSchedulePost) | **Post** /set-ips-update-schedule | 
[**ShowIpsStatusPost**](ThreatPreventionIPSApi.md#ShowIpsStatusPost) | **Post** /show-ips-status | 
[**ShowIpsUpdateSchedulePost**](ThreatPreventionIPSApi.md#ShowIpsUpdateSchedulePost) | **Post** /show-ips-update-schedule | 


# **RunIpsUpdatePost**
> IpsUpdateReply RunIpsUpdatePost(ctx, ipsUpdateRequest)


Runs IPS database update. If \"package-path\" is not provided server will try to get the latest package from the User Center.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ipsUpdateRequest** | [**IpsUpdateRequest**](IpsUpdateRequest.md)|  | 

### Return type

[**IpsUpdateReply**](IpsUpdateReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetIpsUpdateSchedulePost**
> IpsUpdateScheduleReply SetIpsUpdateSchedulePost(ctx, ipsUpdateScheduleRequestEdit)


Edit existing object using object name or uid.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ipsUpdateScheduleRequestEdit** | [**IpsUpdateScheduleRequestEdit**](IpsUpdateScheduleRequestEdit.md)|  | 

### Return type

[**IpsUpdateScheduleReply**](IpsUpdateScheduleReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowIpsStatusPost**
> IpsStatusReply ShowIpsStatusPost(ctx, ipsStatusRequest)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ipsStatusRequest** | [**IpsStatusRequest**](IpsStatusRequest.md)|  | 

### Return type

[**IpsStatusReply**](IpsStatusReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowIpsUpdateSchedulePost**
> IpsUpdateScheduleReply ShowIpsUpdateSchedulePost(ctx, ipsUpdateScheduleRequestShow)


Retrieve existing object using object name or uid.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ipsUpdateScheduleRequestShow** | [**IpsUpdateScheduleRequestShow**](IpsUpdateScheduleRequestShow.md)|  | 

### Return type

[**IpsUpdateScheduleReply**](IpsUpdateScheduleReply.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

