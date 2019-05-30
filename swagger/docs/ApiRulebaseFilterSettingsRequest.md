# ApiRulebaseFilterSettingsRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PacketSearchSettings** | [***ApiRulebaseFilterSettingsRequestPacketSearchSettings**](ApiRulebaseFilterSettingsRequestPacketSearchSettings.md) |  | [optional] [default to null]
**SearchMode** | **string** | When set to &#39;general&#39;, both the Full Text Search and Packet Search are enabled. In this mode, Packet Search will not match on &#39;Any&#39; object, a negated cell or a group-with-exclusion. When the search-mode is set to &#39;packet&#39;, by default, the match on &#39;Any&#39; object, a negated cell or a group-with-exclusion are enabled. packet-search-settings may be provided to change the default behavior. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


