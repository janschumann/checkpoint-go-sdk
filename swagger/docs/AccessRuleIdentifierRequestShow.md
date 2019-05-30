# AccessRuleIdentifierRequestShow

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DetailsLevel** | **string** | The level of detail for some of the fields in the response can vary from showing only the UID value of the object to a fully detailed representation of the object. | [optional] [default to null]
**HitsSettings** | [***HitsSettingsRequest**](HitsSettingsRequest.md) |  | [optional] [default to null]
**Layer** | **string** | Layer that the rule belongs to identified by the name or UID. | [default to null]
**Name** | **string** | Object name. | [optional] [default to null]
**RuleNumber** | **int32** | Rule number. | [optional] [default to null]
**ShowAsRanges** | **bool** | When true, the source, destination and services &amp; applications parameters are displayed as ranges of IP addresses and port numbers rather than network objects.&lt;br /&gt; Objects that are not represented using IP addresses or port numbers are presented as objects.&lt;br /&gt; In addition, the response of each rule does not contain the parameters: source, source-negate, destination, destination-negate, service and service-negate, but instead it contains the parameters: source-ranges, destination-ranges and service-ranges.&lt;br /&gt;&lt;br /&gt; Note: Requesting to show rules as ranges is limited up to 20 rules per request, otherwise an error is returned. If you wish to request more rules, use the offset and limit parameters to limit your request. | [optional] [default to null]
**ShowHits** | **bool** | N/A | [optional] [default to null]
**Uid** | **string** | Object unique identifier. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


