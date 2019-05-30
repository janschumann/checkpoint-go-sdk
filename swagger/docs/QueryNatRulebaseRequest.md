# QueryNatRulebaseRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DereferenceGroupMembers** | **bool** | Indicates whether to dereference \&quot;members\&quot; field by details level for every object in reply. | [optional] [default to null]
**DetailsLevel** | **string** | The level of detail for some of the fields in the response can vary from showing only the UID value of the object to a fully detailed representation of the object. | [optional] [default to null]
**Filter** | **string** | Search expression to filter the rulebase. The provided text should be exactly the same as it would be given in Smart Console. The logical operators in the expression (&#39;AND&#39;, &#39;OR&#39;) should be provided in capital letters. If an operator is not used, the default OR operator applies. | [optional] [default to null]
**FilterSettings** | [***ApiRulebaseFilterSettingsRequest**](ApiRulebaseFilterSettingsRequest.md) |  | [optional] [default to null]
**Limit** | **int32** | No more than that many results will be returned. | [optional] [default to null]
**Offset** | **int32** | Skip that many results before beginning to return them. | [optional] [default to null]
**Order** | [**[]ApiQueryOrderRequest**](ApiQueryOrderRequest.md) | Sorts results by the given field. By default the results are sorted in the ascending order by name. | [optional] [default to null]
**Package_** | **string** | Name of the package. | [default to null]
**ShowMembership** | **bool** | Indicates whether to calculate and show \&quot;groups\&quot; field for every object in reply. | [optional] [default to null]
**UseObjectDictionary** | **bool** | N/A | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


