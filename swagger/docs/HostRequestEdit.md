# HostRequestEdit

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Color** | **string** | Color of the object. Should be one of existing colors. | [optional] [default to null]
**Comments** | **string** | Comments string. | [optional] [default to null]
**DetailsLevel** | **string** | The level of detail for some of the fields in the response can vary from showing only the UID value of the object to a fully detailed representation of the object. | [optional] [default to null]
**Groups** | [***Add**](add.md) |  | [optional] [default to null]
**HostServers** | [***HostServersRequestEdit**](HostServersRequestEdit.md) |  | [optional] [default to null]
**IgnoreErrors** | **bool** | Apply changes ignoring errors. You won&#39;t be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. | [optional] [default to null]
**IgnoreWarnings** | **bool** | Apply changes ignoring warnings. | [optional] [default to null]
**Interfaces** | [***Add**](add.md) |  | [optional] [default to null]
**IpAddress** | **string** | IPv4 or IPv6 address. If both addresses are required use ipv4-address and ipv6-address fields explicitly. | [optional] [default to null]
**Ipv4Address** | **string** | IPv4 address. | [optional] [default to null]
**Ipv6Address** | **string** | IPv6 address. | [optional] [default to null]
**Name** | **string** | Object name. | [optional] [default to null]
**NatSettings** | [***NatSettingsRequest**](NatSettingsRequest.md) |  | [optional] [default to null]
**NewName** | **string** | New name of the object. | [optional] [default to null]
**Tags** | [***Add**](add.md) |  | [optional] [default to null]
**Uid** | **string** | Object unique identifier. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


