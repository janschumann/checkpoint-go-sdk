/*
 * Check Point Management API
 *
 * Generated by https://github.com/janschumann/checkpoint-go-sdk.
 *
 * API version: v1.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type NetworkRequestNew struct {
	// Allow broadcast address inclusion.
	Broadcast string `json:"broadcast,omitempty" xml:"broadcast"`
	// Color of the object. Should be one of existing colors.
	Color string `json:"color,omitempty" xml:"color"`
	// Comments string.
	Comments string `json:"comments,omitempty" xml:"comments"`
	// The level of detail for some of the fields in the response can vary from showing only the UID value of the object to a fully detailed representation of the object.
	DetailsLevel string `json:"details-level,omitempty" xml:"details-level"`
	// Collection of group identifiers.
	Groups string `json:"groups,omitempty" xml:"groups"`
	// Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.
	IgnoreErrors bool `json:"ignore-errors,omitempty" xml:"ignore-errors"`
	// Apply changes ignoring warnings.
	IgnoreWarnings bool `json:"ignore-warnings,omitempty" xml:"ignore-warnings"`
	// IPv4 or IPv6 network mask length. If both masks are required use mask-length4 and mask-length6 fields explicitly. Instead of IPv4 mask length it is possible to specify IPv4 mask itself in subnet-mask field.
	MaskLength int32 `json:"mask-length" xml:"mask-length"`
	// IPv4 network mask length.
	MaskLength4 int32 `json:"mask-length4,omitempty" xml:"mask-length4"`
	// IPv6 network mask length.
	MaskLength6 int32 `json:"mask-length6,omitempty" xml:"mask-length6"`
	// Object name. Should be unique in the domain.
	Name string `json:"name" xml:"name"`
	NatSettings *NatSettingsRequest `json:"nat-settings,omitempty" xml:"nat-settings"`
	// If another object with the same identifier already exists, it will be updated. The command behaviour will be the same as if originally a set command was called. Pay attention that original object's fields will be overwritten by the fields provided in the request payload!
	SetIfExists bool `json:"set-if-exists,omitempty" xml:"set-if-exists"`
	// IPv4 or IPv6 network address. If both addresses are required use subnet4 and subnet6 fields explicitly.
	Subnet string `json:"subnet" xml:"subnet"`
	// IPv4 network mask.
	SubnetMask string `json:"subnet-mask,omitempty" xml:"subnet-mask"`
	// IPv4 network address.
	Subnet4 string `json:"subnet4,omitempty" xml:"subnet4"`
	// IPv6 network address.
	Subnet6 string `json:"subnet6,omitempty" xml:"subnet6"`
	// Collection of tag identifiers.
	Tags string `json:"tags,omitempty" xml:"tags"`
}