/*
 * Check Point Management API
 *
 * Generated by https://github.com/janschumann/checkpoint-go-sdk.
 *
 * API version: v1.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type GlobalAssignmentIdentifierRequest struct {
	// N/A
	DependentDomain string `json:"dependent-domain,omitempty" xml:"dependent-domain"`
	// The level of detail for some of the fields in the response can vary from showing only the UID value of the object to a fully detailed representation of the object.
	DetailsLevel string `json:"details-level,omitempty" xml:"details-level"`
	// N/A
	GlobalDomain string `json:"global-domain,omitempty" xml:"global-domain"`
	// Object unique identifier.
	Uid string `json:"uid" xml:"uid"`
}
