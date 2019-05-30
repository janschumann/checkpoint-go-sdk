/*
 * Check Point Management API
 *
 * Generated by https://github.com/janschumann/checkpoint-go-sdk.
 *
 * API version: v1.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type AutoCompleteRequest struct {
	// N/A
	CursorPosition int32 `json:"cursor-position,omitempty" xml:"cursor-position"`
	// N/A
	Prefix string `json:"prefix,omitempty" xml:"prefix"`
}
