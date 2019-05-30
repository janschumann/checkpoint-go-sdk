# PortRanges

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExcludedOthers** | [**[]interface{}**](interface{}.md) | Objects which are not represented as port numbers and are negated in the given rule - for example if negate is set for the service of this rule. The details-level parameter of the request determines whether they are displayed as UID&#39;s or objects. | [optional] [default to null]
**Others** | [**[]interface{}**](interface{}.md) | Objects which are not represented as port numbers and match the given rule.The details-level parameter of the request determines whether they are displayed as UID&#39;s or objects. | [optional] [default to null]
**Tcp** | [**[]ModelRange**](Range.md) | Range of TCP ports that match in the given rule. | [optional] [default to null]
**Udp** | [**[]ModelRange**](Range.md) | Range of UDP ports that match in the given rule. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


