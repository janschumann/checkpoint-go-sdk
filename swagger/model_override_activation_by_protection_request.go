/*
 * Check Point Management API
 *
 * Generated by https://github.com/janschumann/checkpoint-go-sdk.
 *
 * API version: v1.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Overrides per profile for this protection.
type OverrideActivationByProtectionRequest struct {
	// Protection action.
	Action string `json:"action" xml:"action"`
	// Capture packets.
	CapturePackets bool `json:"capture-packets,omitempty" xml:"capture-packets"`
	// IPS protection identified by name or UID.
	Protection string `json:"protection" xml:"protection"`
	// Tracking method for protection.
	Track string `json:"track,omitempty" xml:"track"`
}
