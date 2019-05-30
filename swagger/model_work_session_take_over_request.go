/*
 * Check Point Management API
 *
 * Generated by https://github.com/janschumann/checkpoint-go-sdk.
 *
 * API version: v1.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type WorkSessionTakeOverRequest struct {
	// Allows taking over of an active session, currently executed by another administrator.
	DisconnectActiveSession bool `json:"disconnect-active-session,omitempty" xml:"disconnect-active-session"`
	// Session unique identifier.
	Uid string `json:"uid" xml:"uid"`
}