/*
 * Check Point Management API
 *
 * Generated by https://github.com/janschumann/checkpoint-go-sdk.
 *
 * API version: v1.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ApiVersionsReply struct {
	// N/A
	CurrentVersion string `json:"current-version,omitempty" xml:"current-version"`
	// N/A
	SupportedVersions []string `json:"supported-versions,omitempty" xml:"supported-versions"`
}