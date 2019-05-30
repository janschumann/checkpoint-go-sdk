# LocalDomainServerRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | **bool** | Activate domain server. Only one domain server is allowed to be active. | [optional] [default to null]
**IpAddress** | **string** | IPv4 or IPv6 address. If both addresses are required use ipv4-address and ipv6-address fields explicitly. | [default to null]
**Ipv4Address** | **string** | IPv4 address. | [optional] [default to null]
**Ipv6Address** | **string** | IPv6 address. | [optional] [default to null]
**MultiDomainServer** | **string** | Multi Domain server name or UID. | [default to null]
**Name** | **string** | Object name. Should be unique in the domain. | [default to null]
**SkipStartDomainServer** | **bool** | Set this value to be true to prevent starting the new created domain. | [optional] [default to null]
**Type_** | **string** | Domain server type. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


