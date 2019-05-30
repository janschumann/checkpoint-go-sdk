# DiffRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DereferenceGroupMembers** | **bool** | Indicates whether to dereference \&quot;members\&quot; field by details level for every object in reply. | [optional] [default to null]
**DetailsLevel** | **string** | The level of detail for some of the fields in the response can vary from showing only the UID value of the object to a fully detailed representation of the object. | [optional] [default to null]
**FromDate** | **string** | The date from which tracking changes is to be performed. ISO 8601. If timezone isn&#39;t specified in the input, the Management server&#39;s timezone is used. | [optional] [default to null]
**FromSession** | **string** | The session UID from which tracking changes is to be performed. | [optional] [default to null]
**Limit** | **int32** | Maximum number of sessions to analyze. | [optional] [default to null]
**Offset** | **int32** | Number of sessions to skip (beginning with from-session). | [optional] [default to null]
**ShowMembership** | **bool** | Indicates whether to calculate and show \&quot;groups\&quot; field for every object in reply. | [optional] [default to null]
**ToDate** | **string** | The date until which tracking changes is to be performed. ISO 8601. If timezone isn&#39;t specified in the input, the Management server&#39;s timezone is used. | [optional] [default to null]
**ToSession** | **string** | The session UID until which tracking changes is to be performed. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


