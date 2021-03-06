/*
 * Check Point Management API
 *
 * Generated by https://github.com/janschumann/checkpoint-go-sdk.
 *
 * API version: v1.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type PolicyPackageRequestEdit struct {
	// True - enables, False - disables access & NAT policies, empty - nothing is changed.
	Access bool `json:"access,omitempty" xml:"access"`
	AccessLayers *MultiValueAccessLayer `json:"access-layers,omitempty" xml:"access-layers"`
	// Color of the object. Should be one of existing colors.
	Color string `json:"color,omitempty" xml:"color"`
	// Comments string.
	Comments string `json:"comments,omitempty" xml:"comments"`
	// True - enables, False - disables Desktop security policy, empty - nothing is changed.
	DesktopSecurity bool `json:"desktop-security,omitempty" xml:"desktop-security"`
	// The level of detail for some of the fields in the response can vary from showing only the UID value of the object to a fully detailed representation of the object.
	DetailsLevel string `json:"details-level,omitempty" xml:"details-level"`
	// Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.
	IgnoreErrors bool `json:"ignore-errors,omitempty" xml:"ignore-errors"`
	// Apply changes ignoring warnings.
	IgnoreWarnings bool `json:"ignore-warnings,omitempty" xml:"ignore-warnings"`
	InstallationTargets *Add `json:"installation-targets,omitempty" xml:"installation-targets"`
	// Object name.
	Name string `json:"name,omitempty" xml:"name"`
	// New name of the object.
	NewName string `json:"new-name,omitempty" xml:"new-name"`
	// True - enables, False - disables QoS policy, empty - nothing is changed.
	Qos bool `json:"qos,omitempty" xml:"qos"`
	// QoS policy type.
	QosPolicyType string `json:"qos-policy-type,omitempty" xml:"qos-policy-type"`
	Tags *Add `json:"tags,omitempty" xml:"tags"`
	ThreatLayers *MultiValueThreatLayer `json:"threat-layers,omitempty" xml:"threat-layers"`
	// True - enables, False - disables Threat policy, empty - nothing is changed.
	ThreatPrevention bool `json:"threat-prevention,omitempty" xml:"threat-prevention"`
	// Object unique identifier.
	Uid string `json:"uid" xml:"uid"`
	// True - enables, False - disables VPN traditional mode, empty - nothing is changed.
	VpnTraditionalMode bool `json:"vpn-traditional-mode,omitempty" xml:"vpn-traditional-mode"`
}
