# IpRanges

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExcludedOthers** | [**[]interface{}**](interface{}.md) | Objects which are not represented as IP addresses and are negated in the given rule - for example if negate is set for the source or destination of this rule, or if they appear in the &#39;exclude&#39; member of a group-with-exclusion object. The details-level parameter of the request determines whether they are displayed as UID&#39;s or objects. | [optional] [default to null]
**Ipv4** | [**[]ModelRange**](Range.md) | Range of IPv4 addresses that match in the given rule. | [optional] [default to null]
**Ipv6** | [**[]ModelRange**](Range.md) | Range of IPv6 addresses that match in the given rule. | [optional] [default to null]
**Others** | [**[]interface{}**](interface{}.md) | Objects which are not represented as IP addresses and match the given rule. The details-level parameter of the request determines whether they are displayed as UID&#39;s or objects. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


