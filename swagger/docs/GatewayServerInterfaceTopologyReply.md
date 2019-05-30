# GatewayServerInterfaceTopologyReply

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpAddressBehindThisInterface** | **string** | If the interface is internal, this field specifies to which network it leads. | [optional] [default to null]
**LeadsToDmz** | **bool** | Gets true if the interface leads to DMZ. | [optional] [default to null]
**LeadsToInternet** | **bool** | Gets true if the interface is external. | [optional] [default to null]
**LeadsToSpecificNetwork** | [***ApiObjectStandardIdentifier**](ApiObjectStandardIdentifier.md) |  | [optional] [default to null]
**SecurityZone** | [***ApiObjectStandardIdentifier**](ApiObjectStandardIdentifier.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


