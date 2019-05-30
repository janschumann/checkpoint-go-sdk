/*
 * Check Point Management API
 *
 * Generated by https://github.com/janschumann/checkpoint-go-sdk.
 *
 * API version: v1.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type DnsDomainRequestEdit struct {
	// Color of the object. Should be one of existing colors.
	Color string `json:"color,omitempty" xml:"color"`
	// Comments string.
	Comments string `json:"comments,omitempty" xml:"comments"`
	// The level of detail for some of the fields in the response can vary from showing only the UID value of the object to a fully detailed representation of the object.
	DetailsLevel string `json:"details-level,omitempty" xml:"details-level"`
	// Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.
	IgnoreErrors bool `json:"ignore-errors,omitempty" xml:"ignore-errors"`
	// Apply changes ignoring warnings.
	IgnoreWarnings bool `json:"ignore-warnings,omitempty" xml:"ignore-warnings"`
	// Whether to match sub-domains in addition to the domain itself.
	IsSubDomain bool `json:"is-sub-domain,omitempty" xml:"is-sub-domain"`
	// Object name.
	Name string `json:"name,omitempty" xml:"name"`
	// New name of the DNS domain. Should always start with a '.' character. Should be unique in the domain.
	NewName string `json:"new-name,omitempty" xml:"new-name"`
	Tags *Add `json:"tags,omitempty" xml:"tags"`
	// Object unique identifier.
	Uid string `json:"uid" xml:"uid"`
}
