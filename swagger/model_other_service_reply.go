/*
 * Check Point Management API
 *
 * Generated by https://github.com/janschumann/checkpoint-go-sdk.
 *
 * API version: v1.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type OtherServiceReply struct {
	// Specifies whether Other Service replies are to be accepted.
	AcceptReplies bool `json:"accept-replies,omitempty" xml:"accept-replies"`
	// Contains an INSPECT expression that defines the action to take if a rule containing this service is matched. Example: set r_mhandler &open_ssl_handler sets a handler on the connection.
	Action string `json:"action,omitempty" xml:"action"`
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
	// IP protocol number.
	IpProtocol int32 `json:"ip-protocol,omitempty" xml:"ip-protocol"`
	// Keep connections open after policy has been installed even if they are not allowed under the new policy. This overrides the settings in the Connection Persistence page. If you change this property, the change will not affect open connections, but only future connections.
	KeepConnectionsOpenAfterPolicyInstallation bool `json:"keep-connections-open-after-policy-installation,omitempty" xml:"keep-connections-open-after-policy-installation"`
	// Contains an INSPECT expression that defines the matching criteria. The connection is examined against the expression during the first packet. Example: tcp, dport = 21, direction = 0 matches incoming FTP control connections.
	Match string `json:"match,omitempty" xml:"match"`
	// Indicates whether this service is used when 'Any' is set as the rule's service and there are several service objects with the same source port and protocol.
	MatchForAny bool `json:"match-for-any,omitempty" xml:"match-for-any"`
	MetaInfo *MetaInfoForTopLevelReply `json:"meta-info,omitempty" xml:"meta-info"`
	// Object name. Should be unique in the domain.
	Name string `json:"name,omitempty" xml:"name"`
	// Indicates whether this service is a Data Domain service which has been overridden.
	OverrideDefaultSettings bool `json:"override-default-settings,omitempty" xml:"override-default-settings"`
	// Indicates whether the object is read-only.
	ReadOnly bool `json:"read-only,omitempty" xml:"read-only"`
	// Time (in seconds) before the session times out.
	SessionTimeout int32 `json:"session-timeout,omitempty" xml:"session-timeout"`
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
