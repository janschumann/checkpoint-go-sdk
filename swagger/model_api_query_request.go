/*
 * Check Point Management API
 *
 * Generated by https://github.com/janschumann/checkpoint-go-sdk.
 *
 * API version: v1.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ApiQueryRequest struct {
	// The level of detail for some of the fields in the response can vary from showing only the UID value of the object to a fully detailed representation of the object.
	DetailsLevel string `json:"details-level,omitempty" xml:"details-level"`
	// No more than that many results will be returned.
	Limit int32 `json:"limit,omitempty" xml:"limit"`
	// Skip that many results before beginning to return them.
	Offset int32 `json:"offset,omitempty" xml:"offset"`
	// Sorts results by the given field. By default the results are sorted in the ascending order by name.
	Order []ApiQueryOrderRequest `json:"order,omitempty" xml:"order"`
}
