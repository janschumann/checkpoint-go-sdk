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
type UserCheckReply struct {
	// N/A
	Confirm string `json:"confirm,omitempty" xml:"confirm"`
	CustomFrequency *CustomFrequencySettings `json:"custom-frequency,omitempty" xml:"custom-frequency"`
	// N/A
	Frequency string `json:"frequency,omitempty" xml:"frequency"`
	// How much details are returned depends on the details-level field of the request. This table shows the level of detail shown when details-level is set to standard.
	Interaction []ApiObjectStandardIdentifier `json:"interaction,omitempty" xml:"interaction"`
}
