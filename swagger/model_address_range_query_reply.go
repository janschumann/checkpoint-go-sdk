/*
 * Check Point Management API
 *
 * Generated by https://github.com/janschumann/checkpoint-go-sdk.
 *
 * API version: v1.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type AddressRangeQueryReply struct {
	// From which element number the query was done.
	From int32 `json:"from,omitempty" xml:"from"`
	// How much details are returned depends on the details-level field of the request. This table shows the level of detail shown when details-level is set to standard.
	Objects []AddressRangeStandardReply `json:"objects,omitempty" xml:"objects"`
	// To which element number the query was done.
	To int32 `json:"to,omitempty" xml:"to"`
	// Total number of elements returned by the query.
	Total int32 `json:"total,omitempty" xml:"total"`
}
