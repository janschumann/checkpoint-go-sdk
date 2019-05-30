# VpnMeshedCommunityRequestNew

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Color** | **string** | Color of the object. Should be one of existing colors. | [optional] [default to null]
**Comments** | **string** | Comments string. | [optional] [default to null]
**DetailsLevel** | **string** | The level of detail for some of the fields in the response can vary from showing only the UID value of the object to a fully detailed representation of the object. | [optional] [default to null]
**EncryptionMethod** | **string** | The encryption method to be used. | [optional] [default to null]
**EncryptionSuite** | **string** | The encryption suite to be used. | [optional] [default to null]
**Gateways** | **string** | Collection of Gateway objects identified by the name or UID. | [optional] [default to null]
**IgnoreErrors** | **bool** | Apply changes ignoring errors. You won&#39;t be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. | [optional] [default to null]
**IgnoreWarnings** | **bool** | Apply changes ignoring warnings. | [optional] [default to null]
**IkePhase1** | [***IkeP1Request**](IkeP1Request.md) |  | [optional] [default to null]
**IkePhase2** | [***IkeP2Request**](IkeP2Request.md) |  | [optional] [default to null]
**Name** | **string** | Object name. Should be unique in the domain. | [default to null]
**SharedSecrets** | [***SharedSecretRequest**](SharedSecretRequest.md) |  | [optional] [default to null]
**Tags** | **string** | Collection of tag identifiers. | [optional] [default to null]
**UseSharedSecret** | **bool** | Indicates whether the shared secret should be used for all external gateways. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


