# OpsecApplicationRequestEdit

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Color** | **string** | Color of the object. Should be one of existing colors. | [optional] [default to null]
**Comments** | **string** | Comments string. | [optional] [default to null]
**Cpmi** | [***CpmiRequest**](CpmiRequest.md) |  | [optional] [default to null]
**DetailsLevel** | **string** | The level of detail for some of the fields in the response can vary from showing only the UID value of the object to a fully detailed representation of the object. | [optional] [default to null]
**Host** | **string** | The host where the server is running. Pre-define the host as a network object. | [optional] [default to null]
**IgnoreErrors** | **bool** | Apply changes ignoring errors. You won&#39;t be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. | [optional] [default to null]
**IgnoreWarnings** | **bool** | Apply changes ignoring warnings. | [optional] [default to null]
**Lea** | [***LeaRequest**](LeaRequest.md) |  | [optional] [default to null]
**Name** | **string** | Object name. | [optional] [default to null]
**NewName** | **string** | New name of the object. | [optional] [default to null]
**OneTimePassword** | **string** | A password required for establishing a Secure Internal Communication (SIC). | [optional] [default to null]
**Tags** | [***interface{}**](interface{}.md) | Collection of tag identifiers. | [optional] [default to null]
**Uid** | **string** | Object unique identifier. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


