# TimeRequestNew

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Color** | **string** | Color of the object. Should be one of existing colors. | [optional] [default to null]
**Comments** | **string** | Comments string. | [optional] [default to null]
**DetailsLevel** | **string** | The level of detail for some of the fields in the response can vary from showing only the UID value of the object to a fully detailed representation of the object. | [optional] [default to null]
**End** | [***TimeObjectForRequest**](TimeObjectForRequest.md) |  | [optional] [default to null]
**EndNever** | **bool** | End never. | [optional] [default to null]
**Groups** | **string** | Collection of group identifiers. | [optional] [default to null]
**HoursRanges** | [**[]HourRange**](HourRange.md) | Hours recurrence. Note: Each gateway may interpret this time differently according to its time zone. | [optional] [default to null]
**IgnoreErrors** | **bool** | Apply changes ignoring errors. You won&#39;t be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. | [optional] [default to null]
**IgnoreWarnings** | **bool** | Apply changes ignoring warnings. | [optional] [default to null]
**Name** | **string** | Time object name. Cannot be more than 11 characters. Should be unique in the domain. | [default to null]
**Recurrence** | [***DayRecurrence**](DayRecurrence.md) |  | [optional] [default to null]
**Start** | [***TimeObjectForRequest**](TimeObjectForRequest.md) |  | [optional] [default to null]
**StartNow** | **bool** | Start immediately. | [optional] [default to null]
**Tags** | **string** | Collection of tag identifiers. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


