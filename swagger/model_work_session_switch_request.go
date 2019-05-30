/*
 * Check Point Management API
 *
 * Generated by https://github.com/janschumann/checkpoint-go-sdk.
 *
 * API version: v1.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type WorkSessionSwitchRequest struct {
	// Session unique identifier. It should belong to the current administrator. Switching to the sessions opened in SmartConsole is not supported.
	Uid string `json:"uid" xml:"uid"`
}