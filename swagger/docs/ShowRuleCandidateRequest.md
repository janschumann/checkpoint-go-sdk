# ShowRuleCandidateRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DereferenceGroupMembers** | **bool** | Indicates whether to dereference \&quot;members\&quot; field by details level for every object in reply. | [optional] [default to null]
**DetailsLevel** | **string** | The level of detail for some of the fields in the response can vary from showing only the UID value of the object to a fully detailed representation of the object. | [optional] [default to null]
**Field** | **string** | Policy field name.&lt;br&gt;For access layer: destination, source, service, install-on, action, time, vpn, track, content&lt;br&gt;For nat layer: original-destination, original-source, original-service, translated-destination, translated-source, translated-service, install-on&lt;br&gt;For threat layer: destination, source, service, install-on, protected-scope, action, track&lt;br&gt;For exception: destination, source, service, install-on, protected-scope, action, track, protection-or-site. | [default to null]
**Filter** | **string** | Textual search expression to filter objects by. | [default to null]
**Layer** | **string** | Layer or exception group name or uid. | [default to null]
**Limit** | **int32** | No more than that many results will be returned. | [optional] [default to null]
**Offset** | **int32** | Skip that many results before beginning to return them. | [optional] [default to null]
**Order** | [**[]ApiQueryOrderRequest**](ApiQueryOrderRequest.md) | Sorts results by the given field. By default the results are sorted in the ascending order by name. | [optional] [default to null]
**ShowMembership** | **bool** | Indicates whether to calculate and show \&quot;groups\&quot; field for every object in reply. | [optional] [default to null]
**SubField** | **string** | Optional sub-field to the policy field.&lt;br&gt;For \&quot;vpn\&quot; field in access layer: all-gw-to-gw, specific&lt;br&gt;For \&quot;service\&quot; field in access layer: services, application-and-sites, categories, mobile-application&lt;br&gt;For \&quot;content\&quot; field in access layer: file-types&lt;br&gt;For \&quot;protection-or-site\&quot; field in exception: whitelist-files, ips-protections, anti-protection, user-applicatoin, blades. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


