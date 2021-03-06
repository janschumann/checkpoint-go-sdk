/*
 * Check Point Management API
 *
 * Generated by https://github.com/janschumann/checkpoint-go-sdk.
 *
 * API version: v1.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type QueryGatewaysServersReply struct {
	// From which element number the query was done.
	From int32 `json:"from,omitempty" xml:"from"`
	// How much detail is returned depends on the detail-level field of the request. This table shows the level of detail shown when details-level is set to full.
	Objects []GatewayServerReply `json:"objects,omitempty" xml:"objects"`
	// To which element number the query was done.
	To int32 `json:"to,omitempty" xml:"to"`
	// Total number of elements returned by the query.
	Total int32 `json:"total,omitempty" xml:"total"`
}
