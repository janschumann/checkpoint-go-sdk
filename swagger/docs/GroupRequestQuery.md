# GroupRequestQuery

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DereferenceGroupMembers** | **bool** | Indicates whether to dereference \&quot;members\&quot; field by details level for every object in reply. | [optional] [default to null]
**DetailsLevel** | **string** | The level of detail for some of the fields in the response can vary from showing only the UID value of the object to a fully detailed representation of the object. | [optional] [default to null]
**Limit** | **int32** | No more than that many results will be returned. | [optional] [default to null]
**Offset** | **int32** | Skip that many results before beginning to return them. | [optional] [default to null]
**Order** | [**[]ApiQueryOrderRequest**](ApiQueryOrderRequest.md) | Sorts results by the given field. By default the results are sorted in the ascending order by name. | [optional] [default to null]
**ShowAsRanges** | **bool** | When true, the group&#39;s matched content is displayed as ranges of IP addresses rather than network objects.&lt;br /&gt;Objects that are not represented using IP addresses are presented as objects.&lt;br /&gt;The &#39;members&#39; parameter is omitted from the response and instead the &#39;ranges&#39; parameter is displayed. | [optional] [default to null]
**ShowMembership** | **bool** | Indicates whether to calculate and show \&quot;groups\&quot; field for every object in reply. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


