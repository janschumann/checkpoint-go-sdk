# TrustedClientReply

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Color** | **string** | Color of the object. Should be one of existing colors. | [optional] [default to null]
**Comments** | **string** | Comments string. | [optional] [default to null]
**Domain** | [***ApiDomainIdentifier**](ApiDomainIdentifier.md) |  | [optional] [default to null]
**DomainsAssignment** | [**[]ApiObjectStandardIdentifier**](ApiObjectStandardIdentifier.md) | Domains assignment. How much details are returned depends on the details-level field of the request. This table shows the level of detail shown when details-level is set to standard. | [optional] [default to null]
**Icon** | **string** | Object icon. | [optional] [default to null]
**Ipv4Address** | **string** | IPv4 address. | [optional] [default to null]
**Ipv4AddressFirst** | **string** | First IPv4 address in the range. | [optional] [default to null]
**Ipv4AddressLast** | **string** | Last IPv4 address in the range. | [optional] [default to null]
**Ipv6Address** | **string** | IPv6 address. | [optional] [default to null]
**Ipv6AddressFirst** | **string** | First IPv6 address in the range. | [optional] [default to null]
**Ipv6AddressLast** | **string** | Last IPv6 address in the range. | [optional] [default to null]
**MaskLength4** | **int32** | IPv4 mask length. | [optional] [default to null]
**MaskLength6** | **int32** | IPv6 mask length. | [optional] [default to null]
**MetaInfo** | [***MetaInfoForTopLevelReply**](MetaInfoForTopLevelReply.md) |  | [optional] [default to null]
**MultiDomainServerTrustedClient** | **bool** | Let this trusted client connect to all Multi-Domain Servers in the deployment. | [optional] [default to null]
**Name** | **string** | Object name. Should be unique in the domain. | [optional] [default to null]
**ReadOnly** | **bool** | Indicates whether the object is read-only. | [optional] [default to null]
**SubnetMask4** | **string** | IPv4 mask. | [optional] [default to null]
**Tags** | [**[]ApiObjectStandardIdentifier**](ApiObjectStandardIdentifier.md) | Collection of tag objects identified by the name or UID. How much details are returned depends on the details-level field of the request. This table shows the level of detail shown when details-level is set to standard. | [optional] [default to null]
**Type_** | **string** | Trusted client type. | [optional] [default to null]
**Uid** | **string** | Object unique identifier. | [optional] [default to null]
**WildCard** | **string** | IP wild card (e.g. 192.0.2.*). | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


