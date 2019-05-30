# MailSettingsRequestEdit

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddCustomizedTextToEmailBody** | **bool** | Add customized text to the malicious email body. | [optional] [default to null]
**AddEmailSubjectPrefix** | **bool** | Add a prefix to the malicious email subject. | [optional] [default to null]
**AddXHeaderToEmail** | **bool** | Add an X-Header to the malicious email. | [optional] [default to null]
**EmailAction** | **string** | Block - block the entire malicious email&lt;br&gt;Allow - pass the malicious email and apply email changes (like: remove attachments and links, add x-header, etc...). | [optional] [default to null]
**EmailBodyCustomizedText** | **string** | Customized text for the malicious email body.&lt;br&gt; Available predefined fields:&lt;br&gt; $verdicts$ - the malicious/error attachments/links verdict. | [optional] [default to null]
**EmailSubjectPrefixText** | **string** | Prefix for the malicious email subject. | [optional] [default to null]
**FailedToScanAttachmentsText** | **string** | Replace attachments that failed to be scanned with this text.&lt;br&gt; Available predefined fields:&lt;br&gt; $filename$ - the malicious file name.&lt;br&gt; $md5$ - MD5 of the malicious file. | [optional] [default to null]
**MaliciousAttachmentsText** | **string** | Replace malicious attachments with this text.&lt;br&gt; Available predefined fields:&lt;br&gt; $filename$ - the malicious file name.&lt;br&gt; $md5$ - MD5 of the malicious file. | [optional] [default to null]
**MaliciousLinksText** | **string** | Replace malicious links with this text.&lt;br&gt; Available predefined fields:&lt;br&gt; $neutralized_url$ - neutralized malicious link. | [optional] [default to null]
**RemoveAttachmentsAndLinks** | **bool** | Remove attachments and links from the malicious email. | [optional] [default to null]
**SendCopy** | **bool** | Send a copy of the malicious email to the recipient list. | [optional] [default to null]
**SendCopyList** | [***Add**](add.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


