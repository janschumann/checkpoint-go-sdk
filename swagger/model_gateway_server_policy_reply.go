/*
 * Check Point Management API
 *
 * Generated by https://github.com/janschumann/checkpoint-go-sdk.
 *
 * API version: v1.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Installed policy package.
type GatewayServerPolicyReply struct {
	AccessPolicyInstallationDate *ApiDateReply `json:"access-policy-installation-date,omitempty" xml:"access-policy-installation-date"`
	// Gets true if access-policy was installed.
	AccessPolicyInstalled bool `json:"access-policy-installed,omitempty" xml:"access-policy-installed"`
	// Name of the access-policy.
	AccessPolicyName string `json:"access-policy-name,omitempty" xml:"access-policy-name"`
	AccessPolicyRevision *ApiObjectStandardIdentifier `json:"access-policy-revision,omitempty" xml:"access-policy-revision"`
	// Revisions of the access policy installed on each cluster member.
	ClusterMembersAccessPolicyRevision []GatewayServerPolicyReplyClusterMemberReply `json:"cluster-members-access-policy-revision,omitempty" xml:"cluster-members-access-policy-revision"`
	// Revisions of the threat policy installed on each cluster member.
	ClusterMembersThreatPolicyRevision []GatewayServerPolicyReplyClusterMemberReply `json:"cluster-members-threat-policy-revision,omitempty" xml:"cluster-members-threat-policy-revision"`
	ThreatPolicyInstallationDate *ApiDateReply `json:"threat-policy-installation-date,omitempty" xml:"threat-policy-installation-date"`
	// Gets true if threat-policy was installed.
	ThreatPolicyInstalled bool `json:"threat-policy-installed,omitempty" xml:"threat-policy-installed"`
	// Name of the threat-policy.
	ThreatPolicyName string `json:"threat-policy-name,omitempty" xml:"threat-policy-name"`
	ThreatPolicyRevision *ApiObjectStandardIdentifier `json:"threat-policy-revision,omitempty" xml:"threat-policy-revision"`
}