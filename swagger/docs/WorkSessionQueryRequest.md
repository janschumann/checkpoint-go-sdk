# WorkSessionQueryRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DetailsLevel** | **string** | The level of detail for some of the fields in the response can vary from showing only the UID value of the object to a fully detailed representation of the object. | [optional] [default to null]
**Limit** | **int32** | No more than that many results will be returned. | [optional] [default to null]
**Offset** | **int32** | Skip that many results before beginning to return them. | [optional] [default to null]
**Order** | [**[]ApiQueryOrderRequest**](ApiQueryOrderRequest.md) | Sorts results by the given field. By default the results are sorted in the descending order by the session publish time. | [optional] [default to null]
**ViewPublishedSessions** | **bool** | Show a list of published sessions. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


