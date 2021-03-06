/*
 * Check Point Management API
 *
 * Generated by https://github.com/janschumann/checkpoint-go-sdk.
 *
 * API version: v1.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ProtectionReply struct {
	// Protection comments.
	Comments string `json:"comments,omitempty" xml:"comments"`
	// How confident IPS is that recognized attacks are actually undesirable traffic.
	ConfidenceLevel string `json:"confidence-level,omitempty" xml:"confidence-level"`
	Domain *ApiDomainIdentifier `json:"domain,omitempty" xml:"domain"`
	// Tag the protection with pre-defined follow-up flag.
	FollowUp bool `json:"follow-up,omitempty" xml:"follow-up"`
	// International CVE or CVE candidate name for attack.
	IndustryReference []string `json:"industry-reference,omitempty" xml:"industry-reference"`
	// IPS protection extended attributes.
	IpsAdditionalProperties []IpsAdditionalPropertiesReply `json:"ipsAdditionalProperties,omitempty" xml:"ipsAdditionalProperties"`
	// Object name. Should be unique in the domain.
	Name string `json:"name,omitempty" xml:"name"`
	// How much this protection affects the performance of a Security Gateway.
	PerformanceImpact string `json:"performance-impact,omitempty" xml:"performance-impact"`
	// Protection settings per profile.
	Profiles []AllActivationsByProfileReply `json:"profiles,omitempty" xml:"profiles"`
	// The protection's type (Core or Threat Cloud).
	ProtectionType string `json:"protection-type,omitempty" xml:"protection-type"`
	// The date in which the protection was releases by Check Point.
	ReleaseDate string `json:"release-date,omitempty" xml:"release-date"`
	// The intrusion severity.
	Severity string `json:"severity,omitempty" xml:"severity"`
	// Type of the object.
	Type_ string `json:"type,omitempty" xml:"type"`
	// Object unique identifier.
	Uid string `json:"uid,omitempty" xml:"uid"`
	// The date in which the protection was updated by Check Point.
	UpdateDate string `json:"update-date,omitempty" xml:"update-date"`
}
