# TrustedClientRequestEdit

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Color** | **string** | Color of the object. Should be one of existing colors. | [optional] [default to null]
**Comments** | **string** | Comments string. | [optional] [default to null]
**DetailsLevel** | **string** | The level of detail for some of the fields in the response can vary from showing only the UID value of the object to a fully detailed representation of the object. | [optional] [default to null]
**DomainsAssignment** | [***Add**](add.md) |  | [optional] [default to null]
**IgnoreErrors** | **bool** | Apply changes ignoring errors. You won&#39;t be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. | [optional] [default to null]
**IgnoreWarnings** | **bool** | Apply changes ignoring warnings. | [optional] [default to null]
**IpAddress** | **string** | IPv4 or IPv6 address. If both addresses are required use ipv4-address and ipv6-address fields explicitly. | [optional] [default to null]
**IpAddressFirst** | **string** | First IP address in the range. If both IPv4 and IPv6 address ranges are required, use the ipv4-address-first and the ipv6-address-first fields instead. | [optional] [default to null]
**IpAddressLast** | **string** | Last IP address in the range. If both IPv4 and IPv6 address ranges are required, use the ipv4-address-first and the ipv6-address-first fields instead. | [optional] [default to null]
**Ipv4Address** | **string** | IPv4 address. | [optional] [default to null]
**Ipv4AddressFirst** | **string** | First IPv4 address in the range. | [optional] [default to null]
**Ipv4AddressLast** | **string** | Last IPv4 address in the range. | [optional] [default to null]
**Ipv6Address** | **string** | IPv6 address. | [optional] [default to null]
**Ipv6AddressFirst** | **string** | First IPv6 address in the range. | [optional] [default to null]
**Ipv6AddressLast** | **string** | Last IPv6 address in the range. | [optional] [default to null]
**MaskLength** | **int32** | IPv4 or IPv6 mask length. If both masks are required use mask-length4 and mask-length6 fields explicitly. | [optional] [default to null]
**MaskLength4** | **int32** | IPv4 mask length. | [optional] [default to null]
**MaskLength6** | **int32** | IPv6 mask length. | [optional] [default to null]
**MultiDomainServerTrustedClient** | **bool** | Let this trusted client connect to all Multi-Domain Servers in the deployment. | [optional] [default to null]
**Name** | **string** | Object name. | [optional] [default to null]
**NewName** | **string** | New name of the object. | [optional] [default to null]
**Tags** | [***Add**](add.md) |  | [optional] [default to null]
**Type_** | **string** | Trusted client type. | [optional] [default to null]
**Uid** | **string** | Object unique identifier. | [default to null]
**WildCard** | **string** | IP wild card (e.g. 192.0.2.*). | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


