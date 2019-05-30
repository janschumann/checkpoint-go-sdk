# HostRequestNew

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Color** | **string** | Color of the object. Should be one of existing colors. | [optional] [default to null]
**Comments** | **string** | Comments string. | [optional] [default to null]
**DetailsLevel** | **string** | The level of detail for some of the fields in the response can vary from showing only the UID value of the object to a fully detailed representation of the object. | [optional] [default to null]
**Groups** | **string** | Collection of group identifiers. | [optional] [default to null]
**HostServers** | [***HostServersRequestNew**](HostServersRequestNew.md) |  | [optional] [default to null]
**IgnoreErrors** | **bool** | Apply changes ignoring errors. You won&#39;t be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. | [optional] [default to null]
**IgnoreWarnings** | **bool** | Apply changes ignoring warnings. | [optional] [default to null]
**Interfaces** | [***HostInterfaceRequestNew**](HostInterfaceRequestNew.md) |  | [optional] [default to null]
**IpAddress** | **string** | IPv4 or IPv6 address. If both addresses are required use ipv4-address and ipv6-address fields explicitly. | [default to null]
**Ipv4Address** | **string** | IPv4 address. | [optional] [default to null]
**Ipv6Address** | **string** | IPv6 address. | [optional] [default to null]
**Name** | **string** | Object name. Should be unique in the domain. | [default to null]
**NatSettings** | [***NatSettingsRequest**](NatSettingsRequest.md) |  | [optional] [default to null]
**SetIfExists** | **bool** | If another object with the same identifier already exists, it will be updated. The command behaviour will be the same as if originally a set command was called. Pay attention that original object&#39;s fields will be overwritten by the fields provided in the request payload! | [optional] [default to null]
**Tags** | **string** | Collection of tag identifiers. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


