# ServiceIcmpRequestEdit

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Color** | **string** | Color of the object. Should be one of existing colors. | [optional] [default to null]
**Comments** | **string** | Comments string. | [optional] [default to null]
**DetailsLevel** | **string** | The level of detail for some of the fields in the response can vary from showing only the UID value of the object to a fully detailed representation of the object. | [optional] [default to null]
**Groups** | [***Add**](add.md) |  | [optional] [default to null]
**IcmpCode** | **int32** | As listed in: &lt;a href&#x3D;\&quot;http://www.iana.org/assignments/icmp-parameters\&quot; target&#x3D;\&quot;_blank\&quot;&gt;RFC 792&lt;/a&gt;. | [optional] [default to null]
**IcmpType** | **int32** | As listed in: &lt;a href&#x3D;\&quot;http://www.iana.org/assignments/icmp-parameters\&quot; target&#x3D;\&quot;_blank\&quot;&gt;RFC 792&lt;/a&gt;. | [optional] [default to null]
**IgnoreErrors** | **bool** | Apply changes ignoring errors. You won&#39;t be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. | [optional] [default to null]
**IgnoreWarnings** | **bool** | Apply changes ignoring warnings. | [optional] [default to null]
**KeepConnectionsOpenAfterPolicyInstallation** | **bool** | Keep connections open after policy has been installed even if they are not allowed under the new policy. This overrides the settings in the Connection Persistence page. If you change this property, the change will not affect open connections, but only future connections. | [optional] [default to null]
**Name** | **string** | Object name. | [optional] [default to null]
**NewName** | **string** | New name of the object. | [optional] [default to null]
**Tags** | [***Add**](add.md) |  | [optional] [default to null]
**Uid** | **string** | Object unique identifier. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


