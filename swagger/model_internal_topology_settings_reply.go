/*
 * Check Point Management API
 *
 * Generated by https://github.com/janschumann/checkpoint-go-sdk.
 *
 * API version: v1.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Internal topology settings.
type InternalTopologySettingsReply struct {
	// Whether this interface leads to demilitarized zone (perimeter network).
	InterfaceLeadsToDmz bool `json:"interface-leads-to-dmz,omitempty" xml:"interface-leads-to-dmz"`
	// N/A
	IpAddressBehindThisInterface string `json:"ip-address-behind-this-interface,omitempty" xml:"ip-address-behind-this-interface"`
	// Network behind this interface.
	SpecificNetwork string `json:"specific-network,omitempty" xml:"specific-network"`
}
