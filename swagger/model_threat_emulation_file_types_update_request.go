/*
 * Check Point Management API
 *
 * Generated by https://github.com/janschumann/checkpoint-go-sdk.
 *
 * API version: v1.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ThreatEmulationFileTypesUpdateRequest struct {
	// File path for offline update of Threat Emulation file types, the file path should be on the management machine.
	FilePath string `json:"file-path" xml:"file-path"`
	// The contents of a file containing the Threat Emulation file types.
	FileRawData string `json:"file-raw-data,omitempty" xml:"file-raw-data"`
}
