# DataCenterContentRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataCenterName** | **string** | Name of the Data Center where to search for objects. | [default to null]
**DataCenterUid** | **string** | Unique identifier of the Data Center where to search for objects. | [optional] [default to null]
**DetailsLevel** | **string** | Standard and Full description are the same. | [optional] [default to null]
**Filter** | [***DataCenterContentObjectFilter**](DataCenterContentObjectFilter.md) |  | [optional] [default to null]
**Limit** | **int32** | No more than that many results will be returned. | [optional] [default to null]
**Offset** | **int32** | Skip that many results before beginning to return them. | [optional] [default to null]
**Order** | [**[]ApiQueryOrderRequest**](ApiQueryOrderRequest.md) | Sorts results by the given field. By default the results are sorted in the ascending order by name. | [optional] [default to null]
**UidInDataCenter** | **string** | Return result matching the unique identifier of the object on the Data Center. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


