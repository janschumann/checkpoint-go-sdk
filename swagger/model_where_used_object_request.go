/*
 * Check Point Management API
 *
 * Generated by https://github.com/janschumann/checkpoint-go-sdk.
 *
 * API version: v1.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type WhereUsedObjectRequest struct {
	// Indicates whether to dereference \"members\" field by details level for every object in reply.
	DereferenceGroupMembers bool `json:"dereference-group-members,omitempty" xml:"dereference-group-members"`
	// The level of detail for some of the fields in the response can vary from showing only the UID value of the object to a fully detailed representation of the object.
	DetailsLevel string `json:"details-level,omitempty" xml:"details-level"`
	// Search for indirect usage.
	Indirect bool `json:"indirect,omitempty" xml:"indirect"`
	// Maximum nesting level during indirect usage search.
	IndirectMaxDepth int32 `json:"indirect-max-depth,omitempty" xml:"indirect-max-depth"`
	// Object name.
	Name string `json:"name,omitempty" xml:"name"`
	// Indicates whether to calculate and show \"groups\" field for every object in reply.
	ShowMembership bool `json:"show-membership,omitempty" xml:"show-membership"`
	// Object unique identifier.
	Uid string `json:"uid" xml:"uid"`
}
