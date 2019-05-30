# AccessRuleRequestEdit

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | \&quot;Accept\&quot;, \&quot;Drop\&quot;, \&quot;Ask\&quot;, \&quot;Inform\&quot;, \&quot;Reject\&quot;, \&quot;User Auth\&quot;, \&quot;Client Auth\&quot;, \&quot;Apply Layer\&quot;. | [optional] [default to null]
**ActionSettings** | [***AdvancedActionSettingsRequest**](AdvancedActionSettingsRequest.md) |  | [optional] [default to null]
**Comments** | **string** | Comments string. | [optional] [default to null]
**Content** | [***interface{}**](interface{}.md) | List of processed file types that this rule applies on. | [optional] [default to null]
**ContentDirection** | **string** | On which direction the file types processing is applied. | [optional] [default to null]
**ContentNegate** | **bool** | True if negate is set for data. | [optional] [default to null]
**CustomFields** | [***CustomSummaryFieldsRequest**](CustomSummaryFieldsRequest.md) |  | [optional] [default to null]
**Destination** | [***Add**](add.md) |  | [optional] [default to null]
**DestinationNegate** | **bool** | True if negate is set for destination. | [optional] [default to null]
**DetailsLevel** | **string** | The level of detail for some of the fields in the response can vary from showing only the UID value of the object to a fully detailed representation of the object. | [optional] [default to null]
**Enabled** | **bool** | Enable/Disable the rule. | [optional] [default to null]
**IgnoreErrors** | **bool** | Apply changes ignoring errors. You won&#39;t be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. | [optional] [default to null]
**IgnoreWarnings** | **bool** | Apply changes ignoring warnings. | [optional] [default to null]
**InlineLayer** | **string** | Inline Layer identified by the name or UID. Relevant only if \&quot;Action\&quot; was set to \&quot;Apply Layer\&quot;. | [optional] [default to null]
**InstallOn** | [***Add**](add.md) |  | [optional] [default to null]
**Layer** | **string** | Layer that the rule belongs to identified by the name or UID. | [default to null]
**Name** | **string** | Object name. | [optional] [default to null]
**NewName** | **string** | New name of the object. | [optional] [default to null]
**NewPosition** | **int32** | New position in the rulebase. | [optional] [default to null]
**RuleNumber** | **int32** | Rule number. | [optional] [default to null]
**Service** | [***Add**](add.md) |  | [optional] [default to null]
**ServiceNegate** | **bool** | True if negate is set for service. | [optional] [default to null]
**Source** | [***Add**](add.md) |  | [optional] [default to null]
**SourceNegate** | **bool** | True if negate is set for source. | [optional] [default to null]
**Time** | [***Add**](add.md) |  | [optional] [default to null]
**Track** | [***TrackSettingsForRequest**](TrackSettingsForRequest.md) |  | [optional] [default to null]
**Uid** | **string** | Object unique identifier. | [default to null]
**UserCheck** | [***UserCheckRequest**](UserCheckRequest.md) |  | [optional] [default to null]
**Vpn** | **string** | Communities or Directional. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


