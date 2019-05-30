# TaskQueryRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DetailsLevel** | **string** | The level of detail for some of the fields in the response can vary from showing only the UID value of the object to a fully detailed representation of the object. | [optional] [default to null]
**FromDate** | **string** | The date from which tracking tasks is to be performed, by the task&#39;s last update date. ISO 8601. If timezone isn&#39;t specified in the input, the Management server&#39;s timezone is used. | [optional] [default to null]
**Initiator** | **string** | Initiator&#39;s name. If name isn&#39;t specified, tasks from all initiators will be shown. | [optional] [default to null]
**Limit** | **int32** | No more than that many results will be returned. | [optional] [default to null]
**Offset** | **int32** | Skip that many results before beginning to return them. | [optional] [default to null]
**Order** | [**[]ApiQueryOrderRequest**](ApiQueryOrderRequest.md) | Sorts results by the given field. By default the results are sorted in the descending order by the task&#39;s last update date. | [optional] [default to null]
**Status** | **string** | Status. | [optional] [default to null]
**ToDate** | **string** | The date until which tracking tasks is to be performed, by the task&#39;s last update date. ISO 8601. If timezone isn&#39;t specified in the input, the Management server&#39;s timezone is used. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


