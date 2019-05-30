/*
 * Check Point Management API
 *
 * Generated by https://github.com/janschumann/checkpoint-go-sdk.
 *
 * API version: v1.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Network security blades.
type GatewayServerNetworkBladesReply struct {
	// Anti-Bot blade.
	AntiBot bool `json:"anti-bot,omitempty" xml:"anti-bot"`
	// Anti-Spam & Email Security blade.
	AntiSpam bool `json:"anti-spam,omitempty" xml:"anti-spam"`
	// Anti-Virus blade.
	AntiVirus bool `json:"anti-virus,omitempty" xml:"anti-virus"`
	// Application control blade.
	ApplicationControl bool `json:"application-control,omitempty" xml:"application-control"`
	// Content awareness blade.
	ContentAwareness bool `json:"content-awareness,omitempty" xml:"content-awareness"`
	// Data loss prevention blade.
	DataLossPrevention bool `json:"data-loss-prevention,omitempty" xml:"data-loss-prevention"`
	// Firewall blade.
	Firewall bool `json:"firewall,omitempty" xml:"firewall"`
	// Identity awareness blade.
	IdentityAwareness bool `json:"identity-awareness,omitempty" xml:"identity-awareness"`
	// IPS blade.
	Ips bool `json:"ips,omitempty" xml:"ips"`
	// Mobile access blade.
	MobileAccess bool `json:"mobile-access,omitempty" xml:"mobile-access"`
	// Monitoring blade.
	Monitoring bool `json:"monitoring,omitempty" xml:"monitoring"`
	// QoS blade.
	Qos bool `json:"qos,omitempty" xml:"qos"`
	// Site to site VPN blade.
	SiteToSiteVpn bool `json:"site-to-site-vpn,omitempty" xml:"site-to-site-vpn"`
	// Threat emulation blade.
	ThreatEmulation bool `json:"threat-emulation,omitempty" xml:"threat-emulation"`
	// Threat extraction blade.
	ThreatExtraction bool `json:"threat-extraction,omitempty" xml:"threat-extraction"`
	// Traditional Anti-Virus blade.
	TraditionalAntiVirus bool `json:"traditional-anti-virus,omitempty" xml:"traditional-anti-virus"`
	// URL filtering blade.
	UrlFiltering bool `json:"url-filtering,omitempty" xml:"url-filtering"`
}
