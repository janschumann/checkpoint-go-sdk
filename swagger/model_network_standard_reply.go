/*
 * Check Point Management API
 *
 * Generated by https://github.com/janschumann/checkpoint-go-sdk.
 *
 * API version: v1.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type NetworkStandardReply struct {
	Domain *ApiDomainIdentifier `json:"domain,omitempty" xml:"domain"`
	// IPv4 network mask length.
	MaskLength4 int32 `json:"mask-length4,omitempty" xml:"mask-length4"`
	// IPv6 network mask length.
	MaskLength6 int32 `json:"mask-length6,omitempty" xml:"mask-length6"`
	// Object name. Should be unique in the domain.
	Name string `json:"name,omitempty" xml:"name"`
	// IPv4 network mask.
	SubnetMask string `json:"subnet-mask,omitempty" xml:"subnet-mask"`
	// IPv4 network address.
	Subnet4 string `json:"subnet4,omitempty" xml:"subnet4"`
	// IPv6 network address.
	Subnet6 string `json:"subnet6,omitempty" xml:"subnet6"`
	// Type of the object.
	Type_ string `json:"type,omitempty" xml:"type"`
	// Object unique identifier.
	Uid string `json:"uid,omitempty" xml:"uid"`
}
