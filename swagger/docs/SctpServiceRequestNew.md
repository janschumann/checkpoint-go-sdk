# SctpServiceRequestNew

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AggressiveAging** | [***AggressiveAgingRequest**](AggressiveAgingRequest.md) |  | [optional] [default to null]
**Color** | **string** | Color of the object. Should be one of existing colors. | [optional] [default to null]
**Comments** | **string** | Comments string. | [optional] [default to null]
**DetailsLevel** | **string** | The level of detail for some of the fields in the response can vary from showing only the UID value of the object to a fully detailed representation of the object. | [optional] [default to null]
**Groups** | **string** | Collection of group identifiers. | [optional] [default to null]
**IgnoreErrors** | **bool** | Apply changes ignoring errors. You won&#39;t be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. | [optional] [default to null]
**IgnoreWarnings** | **bool** | Apply changes ignoring warnings. | [optional] [default to null]
**KeepConnectionsOpenAfterPolicyInstallation** | **bool** | Keep connections open after policy has been installed even if they are not allowed under the new policy. This overrides the settings in the Connection Persistence page. If you change this property, the change will not affect open connections, but only future connections. | [optional] [default to null]
**MatchForAny** | **bool** | Indicates whether this service is used when &#39;Any&#39; is set as the rule&#39;s service and there are several service objects with the same source port and protocol. | [optional] [default to null]
**Name** | **string** | Object name. Should be unique in the domain. | [default to null]
**Port** | **string** | Port number. To specify a port range add a hyphen between the lowest and the highest port numbers, for example 44-45. | [default to null]
**SessionTimeout** | **int32** | Time (in seconds) before the session times out. | [optional] [default to null]
**SetIfExists** | **bool** | If another object with the same identifier already exists, it will be updated. The command behaviour will be the same as if originally a set command was called. Pay attention that original object&#39;s fields will be overwritten by the fields provided in the request payload! | [optional] [default to null]
**SourcePort** | **string** | Source port number. To specify a port range add a hyphen between the lowest and the highest port numbers, for example 44-45. | [optional] [default to null]
**SyncConnectionsOnCluster** | **bool** | Enables state-synchronized High Availability or Load Sharing on a ClusterXL or OPSEC-certified cluster. | [optional] [default to null]
**Tags** | **string** | Collection of tag identifiers. | [optional] [default to null]
**UseDefaultSessionTimeout** | **bool** | Use default virtual session timeout. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


