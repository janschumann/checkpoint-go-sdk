/*
 * Check Point Management API
 *
 * Generated by https://github.com/janschumann/checkpoint-go-sdk.
 *
 * API version: v1.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Additional properties on the object.
type UpdatableObjectAdditionalProperties struct {
	// Description of retrieved Updatable Object.
	Description string `json:"description,omitempty" xml:"description"`
	// Information about the Updatable Object IP ranges source.
	InfoText string `json:"info-text,omitempty" xml:"info-text"`
	// URL of the Updatable Object IP ranges source.
	InfoUrl string `json:"info-url,omitempty" xml:"info-url"`
	// URI of the Updatable Object under the Updatable Objects Repository.
	Uri string `json:"uri,omitempty" xml:"uri"`
}
