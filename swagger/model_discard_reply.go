/*
 * Check Point Management API
 *
 * Generated by https://github.com/janschumann/checkpoint-go-sdk.
 *
 * API version: v1.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type DiscardReply struct {
	// Operation status.
	Message string `json:"message,omitempty" xml:"message"`
	// Number of discarded changes.
	NumberOfDiscardedChanges int32 `json:"number-of-discarded-changes,omitempty" xml:"number-of-discarded-changes"`
}
