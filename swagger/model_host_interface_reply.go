/*
 * Check Point Management API
 *
 * Generated by https://github.com/janschumann/checkpoint-go-sdk.
 *
 * API version: v1.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type HostInterfaceReply struct {
	// Color of the object. Should be one of existing colors.
	Color string `json:"color,omitempty" xml:"color"`
	// Comments string.
	Comments string `json:"comments,omitempty" xml:"comments"`
	Domain *ApiDomainIdentifier `json:"domain,omitempty" xml:"domain"`
	// Object icon.
	Icon string `json:"icon,omitempty" xml:"icon"`
	// IPv4 network mask length.
	MaskLength4 int32 `json:"mask-length4,omitempty" xml:"mask-length4"`
	// IPv6 network mask length.
	MaskLength6 int32 `json:"mask-length6,omitempty" xml:"mask-length6"`
	MetaInfo *MetaInfoForTopLevelReply `json:"meta-info,omitempty" xml:"meta-info"`
	// Object name. Should be unique in the domain.
	Name string `json:"name,omitempty" xml:"name"`
	// Indicates whether the object is read-only.
	ReadOnly bool `json:"read-only,omitempty" xml:"read-only"`
	// IPv4 network mask.
	SubnetMask string `json:"subnet-mask,omitempty" xml:"subnet-mask"`
	// IPv4 network address.
	Subnet4 string `json:"subnet4,omitempty" xml:"subnet4"`
	// IPv6 network address.
	Subnet6 string `json:"subnet6,omitempty" xml:"subnet6"`
	// Collection of tag objects identified by the name or UID. How much details are returned depends on the details-level field of the request. This table shows the level of detail shown when details-level is set to standard.
	Tags []ApiObjectStandardIdentifier `json:"tags,omitempty" xml:"tags"`
	// Type of the object.
	Type_ string `json:"type,omitempty" xml:"type"`
	// Object unique identifier.
	Uid string `json:"uid,omitempty" xml:"uid"`
}
