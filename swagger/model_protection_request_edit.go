/*
 * Check Point Management API
 *
 * Generated by https://github.com/janschumann/checkpoint-go-sdk.
 *
 * API version: v1.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ProtectionRequestEdit struct {
	// Protection comments.
	Comments string `json:"comments,omitempty" xml:"comments"`
	// The level of detail for some of the fields in the response can vary from showing only the UID value of the object to a fully detailed representation of the object.
	DetailsLevel string `json:"details-level,omitempty" xml:"details-level"`
	// Tag the protection with pre-defined follow-up flag.
	FollowUp bool `json:"follow-up,omitempty" xml:"follow-up"`
	// Object name.
	Name string `json:"name,omitempty" xml:"name"`
	Overrides *Remove `json:"overrides,omitempty" xml:"overrides"`
	// Object unique identifier.
	Uid string `json:"uid" xml:"uid"`
}
