# ThreatExceptionGroupRequestNew

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppliedProfile** | **string** | The threat profile to apply this group to in the case of apply-on threat-rules-with-specific-profile. | [optional] [default to null]
**AppliedThreatRules** | [***ManualAttachment**](ManualAttachment.md) |  | [optional] [default to null]
**ApplyOn** | **string** | An exception group can be set to apply on all threat rules, all threat rules which have a specific profile, or those rules manually chosen by the user. | [optional] [default to null]
**Color** | **string** | Color of the object. Should be one of existing colors. | [optional] [default to null]
**Comments** | **string** | Comments string. | [optional] [default to null]
**DetailsLevel** | **string** | The level of detail for some of the fields in the response can vary from showing only the UID value of the object to a fully detailed representation of the object. | [optional] [default to null]
**IgnoreErrors** | **bool** | Apply changes ignoring errors. You won&#39;t be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. | [optional] [default to null]
**IgnoreWarnings** | **bool** | Apply changes ignoring warnings. | [optional] [default to null]
**Name** | **string** | Object name. Should be unique in the domain. | [default to null]
**Tags** | **string** | Collection of tag identifiers. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


