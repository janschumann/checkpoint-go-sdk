/*
 * Check Point Management API
 *
 * Generated by https://github.com/janschumann/checkpoint-go-sdk.
 *
 * API version: v1.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// N/A
type SecurityZoneSettingsRequest struct {
	// Security Zone is calculated according to where the interface leads to.
	AutoCalculated bool `json:"auto-calculated,omitempty" xml:"auto-calculated"`
	// Security Zone specified manually.
	SpecificZone string `json:"specific-zone,omitempty" xml:"specific-zone"`
}
