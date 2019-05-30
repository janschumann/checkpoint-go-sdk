# ApiRulebaseFilterSettingsRequestPacketSearchSettings

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpandGroupMembers** | **bool** | When true, if the search expression contains a UID or a name of a group object, results will include rules that match on at least one member of the group. | [optional] [default to null]
**ExpandGroupWithExclusionMembers** | **bool** | When true, if the search expression contains a UID or a name of a group-with-exclusion object, results will include rules that match at least one member of the \&quot;include\&quot; part and is not a member of the \&quot;except\&quot; part. | [optional] [default to null]
**MatchOnAny** | **bool** | Whether to match on &#39;Any&#39; object. | [optional] [default to null]
**MatchOnGroupWithExclusion** | **bool** | Whether to match on a group-with-exclusion. | [optional] [default to null]
**MatchOnNegate** | **bool** | Whether to match on a negated cell. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


