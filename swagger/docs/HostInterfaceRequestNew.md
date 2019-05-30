# HostInterfaceRequestNew

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Color** | **string** | Color of the object. Should be one of existing colors. | [optional] [default to null]
**Comments** | **string** | Comments string. | [optional] [default to null]
**DetailsLevel** | **string** | The level of detail for some of the fields in the response can vary from showing only the UID value of the object to a fully detailed representation of the object. | [optional] [default to null]
**IgnoreErrors** | **bool** | Apply changes ignoring errors. You won&#39;t be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. | [optional] [default to null]
**IgnoreWarnings** | **bool** | Apply changes ignoring warnings. | [optional] [default to null]
**MaskLength** | **int32** | IPv4 or IPv6 network mask length. If both masks are required use mask-length4 and mask-length6 fields explicitly. Instead of IPv4 mask length it is possible to specify IPv4 mask itself in subnet-mask field. | [default to null]
**MaskLength4** | **int32** | IPv4 network mask length. | [optional] [default to null]
**MaskLength6** | **int32** | IPv6 network mask length. | [optional] [default to null]
**Name** | **string** | Interface name. | [default to null]
**Subnet** | **string** | IPv4 or IPv6 network address. If both addresses are required use subnet4 and subnet6 fields explicitly. | [default to null]
**SubnetMask** | **string** | IPv4 network mask. | [optional] [default to null]
**Subnet4** | **string** | IPv4 network address. | [optional] [default to null]
**Subnet6** | **string** | IPv6 network address. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


