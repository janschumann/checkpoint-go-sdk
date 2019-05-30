# NetworkInterfaceRequestNew

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AntiSpoofing** | **bool** | N/A | [optional] [default to null]
**AntiSpoofingSettings** | [***AntiSpoofingSettingsRequest**](AntiSpoofingSettingsRequest.md) |  | [optional] [default to null]
**Color** | **string** | Color of the object. Should be one of existing colors. | [optional] [default to null]
**Comments** | **string** | Comments string. | [optional] [default to null]
**DetailsLevel** | **string** | The level of detail for some of the fields in the response can vary from showing only the UID value of the object to a fully detailed representation of the object. | [optional] [default to null]
**IgnoreErrors** | **bool** | Apply changes ignoring errors. You won&#39;t be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. | [optional] [default to null]
**IgnoreWarnings** | **bool** | Apply changes ignoring warnings. | [optional] [default to null]
**IpAddress** | **string** | IPv4 or IPv6 address. If both addresses are required use ipv4-address and ipv6-address fields explicitly. | [default to null]
**Ipv4Address** | **string** | IPv4 address. | [optional] [default to null]
**Ipv4MaskLength** | **string** | IPv4 network mask length. | [optional] [default to null]
**Ipv4NetworkMask** | **string** | IPv4 network address. | [optional] [default to null]
**Ipv6Address** | **string** | IPv6 address. | [optional] [default to null]
**Ipv6MaskLength** | **string** | IPv6 network mask length. | [optional] [default to null]
**Ipv6NetworkMask** | **string** | IPv6 network address. | [optional] [default to null]
**MaskLength** | **string** | IPv4 or IPv6 network mask length. | [optional] [default to null]
**Name** | **string** | Object name. Should be unique in the domain. | [default to null]
**NetworkMask** | **string** | IPv4 or IPv6 network mask. If both masks are required use ipv4-network-mask and ipv6-network-mask fields explicitly. Instead of providing mask itself it is possible to specify IPv4 or IPv6 mask length in mask-length field. If both masks length are required use ipv4-mask-length and  ipv6-mask-length fields explicitly. | [default to null]
**SecurityZone** | **bool** | N/A | [optional] [default to null]
**SecurityZoneSettings** | [***SecurityZoneSettingsRequest**](SecurityZoneSettingsRequest.md) |  | [optional] [default to null]
**Tags** | [***interface{}**](interface{}.md) | Collection of tag identifiers. | [optional] [default to null]
**Topology** | **string** | N/A | [optional] [default to null]
**TopologySettings** | [***InternalTopologySettingsRequest**](InternalTopologySettingsRequest.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


