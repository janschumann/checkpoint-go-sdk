/*
 * Check Point Management API
 *
 * Generated by https://github.com/janschumann/checkpoint-go-sdk.
 *
 * API version: v1.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type AdministratorReply struct {
	// Authentication method.
	AuthenticationMethod string `json:"authentication-method,omitempty" xml:"authentication-method"`
	// Color of the object. Should be one of existing colors.
	Color string `json:"color,omitempty" xml:"color"`
	// Comments string.
	Comments string `json:"comments,omitempty" xml:"comments"`
	Domain *ApiDomainIdentifier `json:"domain,omitempty" xml:"domain"`
	// Administrator email.
	Email string `json:"email,omitempty" xml:"email"`
	ExpirationDate *ApiDateReply `json:"expiration-date,omitempty" xml:"expiration-date"`
	// Object icon.
	Icon string `json:"icon,omitempty" xml:"icon"`
	MetaInfo *MetaInfoForTopLevelReply `json:"meta-info,omitempty" xml:"meta-info"`
	// Administrator multi-domain profile. How much details are returned depends on the details-level field of the request. This table shows the level of detail shown when details-level is set to standard.
	MultiDomainProfile []ApiObjectStandardIdentifier `json:"multi-domain-profile,omitempty" xml:"multi-domain-profile"`
	// True if administrator must change password on the next login.
	MustChangePassword bool `json:"must-change-password,omitempty" xml:"must-change-password"`
	// Object name. Should be unique in the domain.
	Name string `json:"name,omitempty" xml:"name"`
	// Administrator permissions profile. How much details are returned depends on the details-level field of the request. This table shows the level of detail shown when details-level is set to standard.
	PermissionsProfile []ApiObjectStandardIdentifier `json:"permissions-profile,omitempty" xml:"permissions-profile"`
	// Administrator phone number.
	PhoneNumber string `json:"phone-number,omitempty" xml:"phone-number"`
	RadiusServer *ApiObjectStandardIdentifier `json:"radius-server,omitempty" xml:"radius-server"`
	// Indicates whether the object is read-only.
	ReadOnly bool `json:"read-only,omitempty" xml:"read-only"`
	// Name of the Secure Internal Connection Trust.
	SicName string `json:"sic-name,omitempty" xml:"sic-name"`
	TacacsServer *ApiObjectStandardIdentifier `json:"tacacs-server,omitempty" xml:"tacacs-server"`
	// Collection of tag objects identified by the name or UID. How much details are returned depends on the details-level field of the request. This table shows the level of detail shown when details-level is set to standard.
	Tags []ApiObjectStandardIdentifier `json:"tags,omitempty" xml:"tags"`
	// Type of the object.
	Type_ string `json:"type,omitempty" xml:"type"`
	// Object unique identifier.
	Uid string `json:"uid,omitempty" xml:"uid"`
}
