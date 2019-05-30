/*
 * Check Point Management API
 *
 * Generated by https://github.com/janschumann/checkpoint-go-sdk.
 *
 * API version: v1.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type UpdatableObjectsRepositoryContentReply struct {
	// From which element number the query was done.
	From int32 `json:"from,omitempty" xml:"from"`
	// Remote view of the retrieved Updatable Objects.
	Objects []UpdatableObjectsRepositoryContentObjectReply `json:"objects,omitempty" xml:"objects"`
	// To which element number the query was done.
	To int32 `json:"to,omitempty" xml:"to"`
	// Total number of elements returned by the query.
	Total int32 `json:"total,omitempty" xml:"total"`
}
