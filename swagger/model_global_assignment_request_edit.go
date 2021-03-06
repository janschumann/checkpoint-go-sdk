/*
 * Check Point Management API
 *
 * Generated by https://github.com/janschumann/checkpoint-go-sdk.
 *
 * API version: v1.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type GlobalAssignmentRequestEdit struct {
	// N/A
	DependentDomain string `json:"dependent-domain,omitempty" xml:"dependent-domain"`
	// The level of detail for some of the fields in the response can vary from showing only the UID value of the object to a fully detailed representation of the object.
	DetailsLevel string `json:"details-level,omitempty" xml:"details-level"`
	// Global domain access policy that is assigned to a dependent domain.
	GlobalAccessPolicy string `json:"global-access-policy,omitempty" xml:"global-access-policy"`
	// Global domain name or UID.
	GlobalDomain string `json:"global-domain,omitempty" xml:"global-domain"`
	// Global domain threat prevention policy that is assigned to a dependent domain.
	GlobalThreatPreventionPolicy string `json:"global-threat-prevention-policy,omitempty" xml:"global-threat-prevention-policy"`
	// Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.
	IgnoreErrors bool `json:"ignore-errors,omitempty" xml:"ignore-errors"`
	// Apply changes ignoring warnings.
	IgnoreWarnings bool `json:"ignore-warnings,omitempty" xml:"ignore-warnings"`
	// N/A
	ManageProtectionActions bool `json:"manage-protection-actions,omitempty" xml:"manage-protection-actions"`
	// Object unique identifier.
	Uid string `json:"uid" xml:"uid"`
}
