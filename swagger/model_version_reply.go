/*
 * Check Point Management API
 *
 * Generated by https://github.com/janschumann/checkpoint-go-sdk.
 *
 * API version: v1.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type VersionReply struct {
	// N/A
	OsBuild string `json:"os-build,omitempty" xml:"os-build"`
	// N/A
	OsEdition string `json:"os-edition,omitempty" xml:"os-edition"`
	// N/A
	OsKernelVersion string `json:"os-kernel-version,omitempty" xml:"os-kernel-version"`
	// N/A
	ProductVersion string `json:"product-version,omitempty" xml:"product-version"`
}