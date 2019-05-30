/*
 * Check Point Management API
 *
 * Generated by https://github.com/janschumann/checkpoint-go-sdk.
 *
 * API version: v1.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// User check settings.
type UserCheckRequest struct {
	// N/A
	Confirm string `json:"confirm,omitempty" xml:"confirm"`
	CustomFrequency *CustomFrequencySettings `json:"custom-frequency,omitempty" xml:"custom-frequency"`
	// N/A
	Frequency string `json:"frequency,omitempty" xml:"frequency"`
	// N/A
	Interaction string `json:"interaction,omitempty" xml:"interaction"`
}