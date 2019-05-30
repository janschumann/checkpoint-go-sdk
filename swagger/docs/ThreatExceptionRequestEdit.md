# ThreatExceptionRequestEdit

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | Action-the enforced profile. | [optional] [default to null]
**Comments** | **string** | Comments string. | [optional] [default to null]
**Destination** | [***Add**](add.md) |  | [optional] [default to null]
**DestinationNegate** | **bool** | True if negate is set for destination. | [optional] [default to null]
**DetailsLevel** | **string** | The level of detail for some of the fields in the response can vary from showing only the UID value of the object to a fully detailed representation of the object. | [optional] [default to null]
**Enabled** | **bool** | Enable/Disable the rule. | [optional] [default to null]
**ExceptionGroupName** | **string** | The name of the exception-group. | [optional] [default to null]
**ExceptionGroupUid** | **string** | The UID of the exception-group. | [default to null]
**ExceptionNumber** | **int32** | N/A | [optional] [default to null]
**IgnoreErrors** | **bool** | Apply changes ignoring errors. You won&#39;t be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. | [optional] [default to null]
**IgnoreWarnings** | **bool** | Apply changes ignoring warnings. | [optional] [default to null]
**InstallOn** | [***Add**](add.md) |  | [optional] [default to null]
**Layer** | **string** | Layer that the rule belongs to identified by the name or UID. | [default to null]
**Name** | **string** | The name of the exception. | [default to null]
**NewName** | **string** | New name of the object. | [optional] [default to null]
**NewPosition** | **int32** | New position in the rulebase. | [optional] [default to null]
**ProtectedScope** | [***Add**](add.md) |  | [optional] [default to null]
**ProtectedScopeNegate** | **bool** | True if negate is set for Protected Scope. | [optional] [default to null]
**ProtectionOrSite** | [***Add**](add.md) |  | [optional] [default to null]
**RuleName** | **string** | The name of the parent rule. | [optional] [default to null]
**RuleNumber** | **int32** | The position in the rulebase of the parent rule. | [optional] [default to null]
**RuleUid** | **string** | The UID of the parent rule. | [default to null]
**Service** | [***Add**](add.md) |  | [optional] [default to null]
**ServiceNegate** | **bool** | True if negate is set for Service. | [optional] [default to null]
**Source** | [***Add**](add.md) |  | [optional] [default to null]
**SourceNegate** | **bool** | True if negate is set for source. | [optional] [default to null]
**Track** | **string** | Packet tracking. | [optional] [default to null]
**Uid** | **string** | Object unique identifier. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


