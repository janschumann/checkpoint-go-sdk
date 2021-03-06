/*
 * Check Point Management API
 *
 * Generated by https://github.com/janschumann/checkpoint-go-sdk.
 *
 * API version: v1.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Administrator permissions profile. Permissions profile should not be provided when multi-domain-profile is set to \"Multi-Domain Super User\" or \"Domain Super User\".
type PermissionsRoleRequest struct {
	// Permission profile.
	Profile string `json:"profile,omitempty" xml:"profile"`
}
