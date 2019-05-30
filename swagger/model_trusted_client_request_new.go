/*
 * Check Point Management API
 *
 * Generated by https://github.com/janschumann/checkpoint-go-sdk.
 *
 * API version: v1.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type TrustedClientRequestNew struct {
	// Color of the object. Should be one of existing colors.
	Color string `json:"color,omitempty" xml:"color"`
	// Comments string.
	Comments string `json:"comments,omitempty" xml:"comments"`
	// The level of detail for some of the fields in the response can vary from showing only the UID value of the object to a fully detailed representation of the object.
	DetailsLevel string `json:"details-level,omitempty" xml:"details-level"`
	// Domains to be added to this profile.
	DomainsAssignment string `json:"domains-assignment,omitempty" xml:"domains-assignment"`
	// Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.
	IgnoreErrors bool `json:"ignore-errors,omitempty" xml:"ignore-errors"`
	// Apply changes ignoring warnings.
	IgnoreWarnings bool `json:"ignore-warnings,omitempty" xml:"ignore-warnings"`
	// IPv4 or IPv6 address. If both addresses are required use ipv4-address and ipv6-address fields explicitly.
	IpAddress string `json:"ip-address" xml:"ip-address"`
	// First IP address in the range. If both IPv4 and IPv6 address ranges are required, use the ipv4-address-first and the ipv6-address-first fields instead.
	IpAddressFirst string `json:"ip-address-first,omitempty" xml:"ip-address-first"`
	// Last IP address in the range. If both IPv4 and IPv6 address ranges are required, use the ipv4-address-first and the ipv6-address-first fields instead.
	IpAddressLast string `json:"ip-address-last,omitempty" xml:"ip-address-last"`
	// IPv4 address.
	Ipv4Address string `json:"ipv4-address,omitempty" xml:"ipv4-address"`
	// First IPv4 address in the range.
	Ipv4AddressFirst string `json:"ipv4-address-first,omitempty" xml:"ipv4-address-first"`
	// Last IPv4 address in the range.
	Ipv4AddressLast string `json:"ipv4-address-last,omitempty" xml:"ipv4-address-last"`
	// IPv6 address.
	Ipv6Address string `json:"ipv6-address,omitempty" xml:"ipv6-address"`
	// First IPv6 address in the range.
	Ipv6AddressFirst string `json:"ipv6-address-first,omitempty" xml:"ipv6-address-first"`
	// Last IPv6 address in the range.
	Ipv6AddressLast string `json:"ipv6-address-last,omitempty" xml:"ipv6-address-last"`
	// IPv4 or IPv6 mask length. If both masks are required use mask-length4 and mask-length6 fields explicitly.
	MaskLength int32 `json:"mask-length,omitempty" xml:"mask-length"`
	// IPv4 mask length.
	MaskLength4 int32 `json:"mask-length4,omitempty" xml:"mask-length4"`
	// IPv6 mask length.
	MaskLength6 int32 `json:"mask-length6,omitempty" xml:"mask-length6"`
	// Let this trusted client connect to all Multi-Domain Servers in the deployment.
	MultiDomainServerTrustedClient bool `json:"multi-domain-server-trusted-client,omitempty" xml:"multi-domain-server-trusted-client"`
	// Object name. Should be unique in the domain.
	Name string `json:"name" xml:"name"`
	// Collection of tag identifiers.
	Tags string `json:"tags,omitempty" xml:"tags"`
	// Trusted client type.
	Type_ string `json:"type,omitempty" xml:"type"`
	// IP wild card (e.g. 192.0.2.*).
	WildCard string `json:"wild-card,omitempty" xml:"wild-card"`
}
