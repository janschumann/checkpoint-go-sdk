/*
 * Check Point Management API
 *
 * Generated by https://github.com/janschumann/checkpoint-go-sdk.
 *
 * API version: v1.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ServiceIcmpRequestNew struct {
	// Color of the object. Should be one of existing colors.
	Color string `json:"color,omitempty" xml:"color"`
	// Comments string.
	Comments string `json:"comments,omitempty" xml:"comments"`
	// The level of detail for some of the fields in the response can vary from showing only the UID value of the object to a fully detailed representation of the object.
	DetailsLevel string `json:"details-level,omitempty" xml:"details-level"`
	// Collection of group identifiers.
	Groups string `json:"groups,omitempty" xml:"groups"`
	// As listed in: <a href=\"http://www.iana.org/assignments/icmp-parameters\" target=\"_blank\">RFC 792</a>.
	IcmpCode int32 `json:"icmp-code,omitempty" xml:"icmp-code"`
	// As listed in: <a href=\"http://www.iana.org/assignments/icmp-parameters\" target=\"_blank\">RFC 792</a>.
	IcmpType int32 `json:"icmp-type,omitempty" xml:"icmp-type"`
	// Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.
	IgnoreErrors bool `json:"ignore-errors,omitempty" xml:"ignore-errors"`
	// Apply changes ignoring warnings.
	IgnoreWarnings bool `json:"ignore-warnings,omitempty" xml:"ignore-warnings"`
	// Keep connections open after policy has been installed even if they are not allowed under the new policy. This overrides the settings in the Connection Persistence page. If you change this property, the change will not affect open connections, but only future connections.
	KeepConnectionsOpenAfterPolicyInstallation bool `json:"keep-connections-open-after-policy-installation,omitempty" xml:"keep-connections-open-after-policy-installation"`
	// Object name. Should be unique in the domain.
	Name string `json:"name" xml:"name"`
	// If another object with the same identifier already exists, it will be updated. The command behaviour will be the same as if originally a set command was called. Pay attention that original object's fields will be overwritten by the fields provided in the request payload!
	SetIfExists bool `json:"set-if-exists,omitempty" xml:"set-if-exists"`
	// Collection of tag identifiers.
	Tags string `json:"tags,omitempty" xml:"tags"`
}
