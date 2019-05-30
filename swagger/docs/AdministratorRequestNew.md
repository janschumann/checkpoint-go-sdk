# AdministratorRequestNew

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationMethod** | **string** | Authentication method. | [optional] [default to null]
**Color** | **string** | Color of the object. Should be one of existing colors. | [optional] [default to null]
**Comments** | **string** | Comments string. | [optional] [default to null]
**DetailsLevel** | **string** | The level of detail for some of the fields in the response can vary from showing only the UID value of the object to a fully detailed representation of the object. | [optional] [default to null]
**Email** | **string** | Administrator email. | [optional] [default to null]
**ExpirationDate** | **string** | Format: YYYY-MM-DD, YYYY-mm-ddThh:mm:ss. | [optional] [default to null]
**IgnoreErrors** | **bool** | Apply changes ignoring errors. You won&#39;t be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. | [optional] [default to null]
**IgnoreWarnings** | **bool** | Apply changes ignoring warnings. | [optional] [default to null]
**MultiDomainProfile** | **string** | Administrator multi-domain profile. | [optional] [default to null]
**MustChangePassword** | **bool** | True if administrator must change password on the next login. | [optional] [default to null]
**Name** | **string** | Object name. Should be unique in the domain. | [default to null]
**Password** | **string** | Administrator password. | [optional] [default to null]
**PasswordHash** | **string** | Administrator password hash. | [optional] [default to null]
**PermissionsProfile** | [***PermissionsRoleRequest**](PermissionsRoleRequest.md) |  | [optional] [default to null]
**PhoneNumber** | **string** | Administrator phone number. | [optional] [default to null]
**RadiusServer** | **string** | RADIUS server object identified by the name or UID. Must be set when \&quot;authentication-method\&quot; was selected to be \&quot;RADIUS\&quot;. | [optional] [default to null]
**TacacsServer** | **string** | TACACS server object identified by the name or UID. Must be set when \&quot;authentication-method\&quot; was selected to be \&quot;TACACS\&quot;. | [optional] [default to null]
**Tags** | **string** | Collection of tag identifiers. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


