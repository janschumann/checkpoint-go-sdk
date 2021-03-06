/*
 * Check Point Management API
 *
 * Generated by https://github.com/janschumann/checkpoint-go-sdk.
 *
 * API version: v1.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type DisconnectRequest struct {
	// Discard all changes committed during the session.
	Discard bool `json:"discard,omitempty" xml:"discard"`
	// Session unique identifier.
	Uid string `json:"uid" xml:"uid"`
}
