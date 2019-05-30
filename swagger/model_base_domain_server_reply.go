/*
 * Check Point Management API
 *
 * Generated by https://github.com/janschumann/checkpoint-go-sdk.
 *
 * API version: v1.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type BaseDomainServerReply struct {
	// Domain server status.
	Active bool `json:"active,omitempty" xml:"active"`
	// IPv4 address.
	Ipv4Address string `json:"ipv4-address,omitempty" xml:"ipv4-address"`
	// IPv6 address.
	Ipv6Address string `json:"ipv6-address,omitempty" xml:"ipv6-address"`
	// Multi Domain server name or UID.
	MultiDomainServer string `json:"multi-domain-server,omitempty" xml:"multi-domain-server"`
	// Object name. Should be unique in the domain.
	Name string `json:"name,omitempty" xml:"name"`
	// Set this value to be true to prevent starting the new created domain.
	SkipStartDomainServer bool `json:"skip-start-domain-server,omitempty" xml:"skip-start-domain-server"`
	// Domain server type.
	Type_ string `json:"type,omitempty" xml:"type"`
}