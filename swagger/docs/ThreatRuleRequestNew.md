# ThreatRuleRequestNew

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | Action-the enforced profile. | [optional] [default to null]
**Comments** | **string** | Comments string. | [optional] [default to null]
**Destination** | **string** | Collection of Network objects identified by the name or UID. | [optional] [default to null]
**DestinationNegate** | **bool** | True if negate is set for destination. | [optional] [default to null]
**DetailsLevel** | **string** | The level of detail for some of the fields in the response can vary from showing only the UID value of the object to a fully detailed representation of the object. | [optional] [default to null]
**Enabled** | **bool** | Enable/Disable the rule. | [optional] [default to null]
**IgnoreErrors** | **bool** | Apply changes ignoring errors. You won&#39;t be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. | [optional] [default to null]
**IgnoreWarnings** | **bool** | Apply changes ignoring warnings. | [optional] [default to null]
**InstallOn** | **string** | Which Gateways identified by the name or UID to install the policy on. | [optional] [default to null]
**Layer** | **string** | Layer that the rule belongs to identified by the name or UID. | [default to null]
**Name** | **string** | Rule name. | [optional] [default to null]
**Position** | **int32** | Position in the rulebase. | [default to null]
**ProtectedScope** | **string** | Collection of objects defining Protected Scope identified by the name or UID. | [optional] [default to null]
**ProtectedScopeNegate** | **bool** | True if negate is set for Protected Scope. | [optional] [default to null]
**Service** | **string** | Collection of Network objects identified by the name or UID. | [optional] [default to null]
**ServiceNegate** | **bool** | True if negate is set for Service. | [optional] [default to null]
**Source** | **string** | Collection of Network objects identified by the name or UID. | [optional] [default to null]
**SourceNegate** | **bool** | True if negate is set for source. | [optional] [default to null]
**Track** | **string** | Packet tracking. | [optional] [default to null]
**TrackSettings** | [***ThreatRuleTrackSettingsRequest**](ThreatRuleTrackSettingsRequest.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


