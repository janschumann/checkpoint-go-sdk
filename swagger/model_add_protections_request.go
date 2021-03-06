/*
 * Check Point Management API
 *
 * Generated by https://github.com/janschumann/checkpoint-go-sdk.
 *
 * API version: v1.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type AddProtectionsRequest struct {
	// Protections package format.
	PackageFormat string `json:"package-format,omitempty" xml:"package-format"`
	// Protections package path.
	PackagePath string `json:"package-path,omitempty" xml:"package-path"`
}
