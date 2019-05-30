/*
 * Check Point Management API
 *
 * Generated by https://github.com/janschumann/checkpoint-go-sdk.
 *
 * API version: v1.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Overrides per profile for this protection<br> Note: Remove override for Core protections removes only the action�s override. Remove override for Threat Cloud protections removes the action, track and packet captures.
type Remove struct {
	// Removes from collection of values
	Remove string `json:"remove,omitempty" xml:"remove"`
}