# ServiceIcmpReply

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Color** | **string** | Color of the object. Should be one of existing colors. | [optional] [default to null]
**Comments** | **string** | Comments string. | [optional] [default to null]
**Domain** | [***ApiDomainIdentifier**](ApiDomainIdentifier.md) |  | [optional] [default to null]
**Groups** | [**[]ApiObjectStandardIdentifier**](ApiObjectStandardIdentifier.md) | How much details are returned depends on the details-level field of the request. This table shows the level of detail shown when details-level is set to standard. | [optional] [default to null]
**IcmpCode** | **int32** | As listed in: &lt;a href&#x3D;\&quot;http://www.iana.org/assignments/icmp-parameters\&quot; target&#x3D;\&quot;_blank\&quot;&gt;RFC 792&lt;/a&gt;. | [optional] [default to null]
**IcmpType** | **int32** | As listed in: &lt;a href&#x3D;\&quot;http://www.iana.org/assignments/icmp-parameters\&quot; target&#x3D;\&quot;_blank\&quot;&gt;RFC 792&lt;/a&gt;. | [optional] [default to null]
**Icon** | **string** | Object icon. | [optional] [default to null]
**KeepConnectionsOpenAfterPolicyInstallation** | **bool** | Keep connections open after policy has been installed even if they are not allowed under the new policy. This overrides the settings in the Connection Persistence page. If you change this property, the change will not affect open connections, but only future connections. | [optional] [default to null]
**MetaInfo** | [***MetaInfoForTopLevelReply**](MetaInfoForTopLevelReply.md) |  | [optional] [default to null]
**Name** | **string** | Object name. Should be unique in the domain. | [optional] [default to null]
**ReadOnly** | **bool** | Indicates whether the object is read-only. | [optional] [default to null]
**Tags** | [**[]ApiObjectStandardIdentifier**](ApiObjectStandardIdentifier.md) | Collection of tag objects identified by the name or UID. How much details are returned depends on the details-level field of the request. This table shows the level of detail shown when details-level is set to standard. | [optional] [default to null]
**Type_** | **string** | Type of the object. | [optional] [default to null]
**Uid** | **string** | Object unique identifier. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


