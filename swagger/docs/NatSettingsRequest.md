# NatSettingsRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoRule** | **bool** | Whether to add automatic address translation rules. | [default to null]
**HideBehind** | **string** | Hide behind method. This parameter is not required in case \&quot;method\&quot; parameter is \&quot;static\&quot;. | [optional] [default to null]
**InstallOn** | **string** | Which gateway should apply the NAT translation. | [optional] [default to null]
**IpAddress** | **string** | IPv4 or IPv6 address. If both addresses are required use ipv4-address and ipv6-address fields explicitly. This parameter is not required in case \&quot;method\&quot; parameter is \&quot;hide\&quot; and \&quot;hide-behind\&quot; parameter is \&quot;gateway\&quot;. | [default to null]
**Ipv4Address** | **string** | IPv4 address. | [optional] [default to null]
**Ipv6Address** | **string** | IPv6 address. | [optional] [default to null]
**Method** | **string** | NAT translation method. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


