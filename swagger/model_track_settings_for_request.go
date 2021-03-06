/*
 * Check Point Management API
 *
 * Generated by https://github.com/janschumann/checkpoint-go-sdk.
 *
 * API version: v1.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Track Settings.
type TrackSettingsForRequest struct {
	// Turns accounting for track on and off.
	Accounting bool `json:"accounting,omitempty" xml:"accounting"`
	// Type of alert for the track.
	Alert string `json:"alert,omitempty" xml:"alert"`
	// Determine whether to generate session log to firewall only connections.
	EnableFirewallSession bool `json:"enable-firewall-session,omitempty" xml:"enable-firewall-session"`
	// Determines whether to perform the log per connection.
	PerConnection bool `json:"per-connection,omitempty" xml:"per-connection"`
	// Determines whether to perform the log per session.
	PerSession bool `json:"per-session,omitempty" xml:"per-session"`
	// \"Log\", \"Extended Log\", \"Detailed  Log\", \"None\".
	Type_ string `json:"type,omitempty" xml:"type"`
}
