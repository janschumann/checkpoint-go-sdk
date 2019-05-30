/*
 * Check Point Management API
 *
 * Generated by https://github.com/janschumann/checkpoint-go-sdk.
 *
 * API version: v1.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// The Data Center the object is on.
type DataCenterServerReply struct {
	// Color of the object. Should be one of existing colors.
	Color string `json:"color,omitempty" xml:"color"`
	// Comments string.
	Comments string `json:"comments,omitempty" xml:"comments"`
	// Data Center type.
	DataCenterType string `json:"data-center-type,omitempty" xml:"data-center-type"`
	Domain *ApiDomainIdentifier `json:"domain,omitempty" xml:"domain"`
	// Object icon.
	Icon string `json:"icon,omitempty" xml:"icon"`
	MetaInfo *MetaInfoForTopLevelReply `json:"meta-info,omitempty" xml:"meta-info"`
	// Object name. Should be unique in the domain.
	Name string `json:"name,omitempty" xml:"name"`
	// Data Center properties.
	Properties []DataCenterProperty `json:"properties,omitempty" xml:"properties"`
	// Indicates whether the object is read-only.
	ReadOnly bool `json:"read-only,omitempty" xml:"read-only"`
	// Collection of tag objects identified by the name or UID. How much details are returned depends on the details-level field of the request. This table shows the level of detail shown when details-level is set to standard.
	Tags []ApiObjectStandardIdentifier `json:"tags,omitempty" xml:"tags"`
	// Type of the object.
	Type_ string `json:"type,omitempty" xml:"type"`
	// Object unique identifier.
	Uid string `json:"uid,omitempty" xml:"uid"`
}
