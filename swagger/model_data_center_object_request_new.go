/*
 * Check Point Management API
 *
 * Generated by https://github.com/janschumann/checkpoint-go-sdk.
 *
 * API version: v1.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type DataCenterObjectRequestNew struct {
	// Color of the object. Should be one of existing colors.
	Color string `json:"color,omitempty" xml:"color"`
	// Comments string.
	Comments string `json:"comments,omitempty" xml:"comments"`
	// Name of the Data Center the object is in.
	DataCenterName string `json:"data-center-name" xml:"data-center-name"`
	// Unique identifier of the Data Center the object is in.
	DataCenterUid string `json:"data-center-uid,omitempty" xml:"data-center-uid"`
	// The level of detail for some of the fields in the response can vary from showing only the UID value of the object to a fully detailed representation of the object.
	DetailsLevel string `json:"details-level,omitempty" xml:"details-level"`
	// Collection of group identifiers.
	Groups *interface{} `json:"groups,omitempty" xml:"groups"`
	// Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.
	IgnoreErrors bool `json:"ignore-errors,omitempty" xml:"ignore-errors"`
	// Apply changes ignoring warnings.
	IgnoreWarnings bool `json:"ignore-warnings,omitempty" xml:"ignore-warnings"`
	// Override default name on data-center.
	Name string `json:"name,omitempty" xml:"name"`
	// Collection of tag identifiers.
	Tags string `json:"tags,omitempty" xml:"tags"`
	// Unique identifier of the object in the Data Center.
	UidInDataCenter string `json:"uid-in-data-center,omitempty" xml:"uid-in-data-center"`
	// URI of the object in the Data Center.
	Uri string `json:"uri" xml:"uri"`
}
