# IndicatorObservableRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comments** | **string** | Comments string. | [optional] [default to null]
**Confidence** | **string** | The confidence level the indicator has that a real threat has been uncovered. | [optional] [default to null]
**Domain** | **string** | The name of a domain. | [optional] [default to null]
**IgnoreErrors** | **bool** | Apply changes ignoring errors. You won&#39;t be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. | [optional] [default to null]
**IgnoreWarnings** | **bool** | Apply changes ignoring warnings. | [optional] [default to null]
**IpAddress** | **string** | A valid IP-Address. | [optional] [default to null]
**IpAddressFirst** | **string** | A valid IP-Address, the beginning of the range. If you configure this parameter with a value, you must also configure the value of the &#39;ip-address-last&#39; parameter. | [optional] [default to null]
**IpAddressLast** | **string** | A valid IP-Address, the end of the range. If you configure this parameter with a value, you must also configure the value of the &#39;ip-address-first&#39; parameter. | [optional] [default to null]
**MailCc** | **string** | A valid E-Mail address, cc field. | [optional] [default to null]
**MailFrom** | **string** | A valid E-Mail address, sender field. | [optional] [default to null]
**MailReplyTo** | **string** | A valid E-Mail address, reply-to field. | [optional] [default to null]
**MailSubject** | **string** | Subject of E-Mail. | [optional] [default to null]
**MailTo** | **string** | A valid E-Mail address, recipient filed. | [optional] [default to null]
**Md5** | **string** | A valid MD5 sequence. | [default to null]
**Name** | **string** | Object name. Should be unique in the domain. | [default to null]
**Product** | **string** | The software blade that processes the observable: AV - AntiVirus, AB - AntiBot. | [optional] [default to null]
**Severity** | **string** | The severity level of the threat. | [optional] [default to null]
**Url** | **string** | A valid URL. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


