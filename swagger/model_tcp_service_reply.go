/*
 * Check Point Management API
 *
 * Generated by https://github.com/janschumann/checkpoint-go-sdk.
 *
 * API version: v1.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type TcpServiceReply struct {
	AggressiveAging *AggressiveAgingReply `json:"aggressive-aging,omitempty" xml:"aggressive-aging"`
	// Color of the object. Should be one of existing colors.
	Color string `json:"color,omitempty" xml:"color"`
	// Comments string.
	Comments string `json:"comments,omitempty" xml:"comments"`
	Domain *ApiDomainIdentifier `json:"domain,omitempty" xml:"domain"`
	// How much details are returned depends on the details-level field of the request. This table shows the level of detail shown when details-level is set to standard.
	Groups []ApiObjectStandardIdentifier `json:"groups,omitempty" xml:"groups"`
	// Object icon.
	Icon string `json:"icon,omitempty" xml:"icon"`
	// Keep connections open after policy has been installed even if they are not allowed under the new policy. This overrides the settings in the Connection Persistence page. If you change this property, the change will not affect open connections, but only future connections.
	KeepConnectionsOpenAfterPolicyInstallation bool `json:"keep-connections-open-after-policy-installation,omitempty" xml:"keep-connections-open-after-policy-installation"`
	// A value of true enables matching by the selected protocol's signature - The signature identifies the protocol as genuine.
	MatchByProtocolSignature bool `json:"match-by-protocol-signature,omitempty" xml:"match-by-protocol-signature"`
	// Indicates whether this service is used when 'Any' is set as the rule's service and there are several service objects with the same source port and protocol.
	MatchForAny bool `json:"match-for-any,omitempty" xml:"match-for-any"`
	MetaInfo *MetaInfoForTopLevelReply `json:"meta-info,omitempty" xml:"meta-info"`
	// Object name. Should be unique in the domain.
	Name string `json:"name,omitempty" xml:"name"`
	// Indicates whether this service is a Data Domain service which has been overridden.
	OverrideDefaultSettings bool `json:"override-default-settings,omitempty" xml:"override-default-settings"`
	// The number of the port used to provide this service.
	Port string `json:"port,omitempty" xml:"port"`
	// The protocol type associated with the service, and by implication, the management server (if any) that enforces Content Security and Authentication for the service.
	Protocol string `json:"protocol,omitempty" xml:"protocol"`
	// Indicates whether the object is read-only.
	ReadOnly bool `json:"read-only,omitempty" xml:"read-only"`
	// Time (in seconds) before the session times out.
	SessionTimeout int32 `json:"session-timeout,omitempty" xml:"session-timeout"`
	// Port number for the client side service. If specified, only those Source port Numbers will be Accepted, Dropped, or Rejected during packet inspection. Otherwise, the source port is not inspected.
	SourcePort string `json:"source-port,omitempty" xml:"source-port"`
	// Enables state-synchronized High Availability or Load Sharing on a ClusterXL or OPSEC-certified cluster.
	SyncConnectionsOnCluster bool `json:"sync-connections-on-cluster,omitempty" xml:"sync-connections-on-cluster"`
	// Collection of tag objects identified by the name or UID. How much details are returned depends on the details-level field of the request. This table shows the level of detail shown when details-level is set to standard.
	Tags []ApiObjectStandardIdentifier `json:"tags,omitempty" xml:"tags"`
	// Type of the object.
	Type_ string `json:"type,omitempty" xml:"type"`
	// Object unique identifier.
	Uid string `json:"uid,omitempty" xml:"uid"`
	// Use default virtual session timeout.
	UseDefaultSessionTimeout bool `json:"use-default-session-timeout,omitempty" xml:"use-default-session-timeout"`
}
