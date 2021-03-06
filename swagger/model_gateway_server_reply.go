/*
 * Check Point Management API
 *
 * Generated by https://github.com/janschumann/checkpoint-go-sdk.
 *
 * API version: v1.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type GatewayServerReply struct {
	// Names of cluster members.
	ClusterMemberNames []string `json:"cluster-member-names,omitempty" xml:"cluster-member-names"`
	// Color of the object. Should be one of existing colors.
	Color string `json:"color,omitempty" xml:"color"`
	// Comments string.
	Comments string `json:"comments,omitempty" xml:"comments"`
	Domain *ApiDomainIdentifier `json:"domain,omitempty" xml:"domain"`
	// How much details are returned depends on the details-level field of the request. This table shows the level of detail shown when details-level is set to standard.
	Groups []ApiObjectStandardIdentifier `json:"groups,omitempty" xml:"groups"`
	// Appliance type.
	Hardware string `json:"hardware,omitempty" xml:"hardware"`
	// Object icon.
	Icon string `json:"icon,omitempty" xml:"icon"`
	// Network interfaces.
	Interfaces []GatewayServerInterfaceReply `json:"interfaces,omitempty" xml:"interfaces"`
	// IPv4 address.
	Ipv4Address string `json:"ipv4-address,omitempty" xml:"ipv4-address"`
	// IPv6 address.
	Ipv6Address string `json:"ipv6-address,omitempty" xml:"ipv6-address"`
	ManagementBlades *GatewayServerManagementBladesReply `json:"management-blades,omitempty" xml:"management-blades"`
	MetaInfo *MetaInfoForTopLevelReply `json:"meta-info,omitempty" xml:"meta-info"`
	// Object name. Should be unique in the domain.
	Name string `json:"name,omitempty" xml:"name"`
	NetworkSecurityBlades *GatewayServerNetworkBladesReply `json:"network-security-blades,omitempty" xml:"network-security-blades"`
	// Operating System.
	OperatingSystem string `json:"operating-system,omitempty" xml:"operating-system"`
	Policy *GatewayServerPolicyReply `json:"policy,omitempty" xml:"policy"`
	// Indicates whether the object is read-only.
	ReadOnly bool `json:"read-only,omitempty" xml:"read-only"`
	// Secure Internal Communication status.
	SicStatus string `json:"sic-status,omitempty" xml:"sic-status"`
	// Collection of tag objects identified by the name or UID.
	Tags []string `json:"tags,omitempty" xml:"tags"`
	// Object type.
	Type_ string `json:"type,omitempty" xml:"type"`
	// Object unique identifier.
	Uid string `json:"uid,omitempty" xml:"uid"`
	// Version.
	Version string `json:"version,omitempty" xml:"version"`
	// VPN domain.
	VpnEncryptionDomain string `json:"vpn-encryption-domain,omitempty" xml:"vpn-encryption-domain"`
	VpnEncryptionDomainManuallyDefined *ApiObjectStandardIdentifier `json:"vpn-encryption-domain-manually-defined,omitempty" xml:"vpn-encryption-domain-manually-defined"`
}
