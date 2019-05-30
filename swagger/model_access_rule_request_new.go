/*
 * Check Point Management API
 *
 * Generated by https://github.com/janschumann/checkpoint-go-sdk.
 *
 * API version: v1.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type AccessRuleRequestNew struct {
	// \"Accept\", \"Drop\", \"Ask\", \"Inform\", \"Reject\", \"User Auth\", \"Client Auth\", \"Apply Layer\".
	Action string `json:"action,omitempty" xml:"action"`
	ActionSettings *AdvancedActionSettingsRequest `json:"action-settings,omitempty" xml:"action-settings"`
	// Comments string.
	Comments string `json:"comments,omitempty" xml:"comments"`
	// List of processed file types that this rule applies on.
	Content *interface{} `json:"content,omitempty" xml:"content"`
	// On which direction the file types processing is applied.
	ContentDirection string `json:"content-direction,omitempty" xml:"content-direction"`
	// True if negate is set for data.
	ContentNegate bool `json:"content-negate,omitempty" xml:"content-negate"`
	CustomFields *CustomSummaryFieldsRequest `json:"custom-fields,omitempty" xml:"custom-fields"`
	// Collection of Network objects identified by the name or UID.
	Destination string `json:"destination,omitempty" xml:"destination"`
	// True if negate is set for destination.
	DestinationNegate bool `json:"destination-negate,omitempty" xml:"destination-negate"`
	// The level of detail for some of the fields in the response can vary from showing only the UID value of the object to a fully detailed representation of the object.
	DetailsLevel string `json:"details-level,omitempty" xml:"details-level"`
	// Enable/Disable the rule.
	Enabled bool `json:"enabled,omitempty" xml:"enabled"`
	// Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.
	IgnoreErrors bool `json:"ignore-errors,omitempty" xml:"ignore-errors"`
	// Apply changes ignoring warnings.
	IgnoreWarnings bool `json:"ignore-warnings,omitempty" xml:"ignore-warnings"`
	// Inline Layer identified by the name or UID. Relevant only if \"Action\" was set to \"Apply Layer\".
	InlineLayer string `json:"inline-layer,omitempty" xml:"inline-layer"`
	// Which Gateways identified by the name or UID to install the policy on.
	InstallOn string `json:"install-on,omitempty" xml:"install-on"`
	// Layer that the rule belongs to identified by the name or UID.
	Layer string `json:"layer" xml:"layer"`
	// Rule name.
	Name string `json:"name,omitempty" xml:"name"`
	// Position in the rulebase.
	Position int32 `json:"position" xml:"position"`
	// Collection of Network objects identified by the name or UID.
	Service string `json:"service,omitempty" xml:"service"`
	// True if negate is set for service.
	ServiceNegate bool `json:"service-negate,omitempty" xml:"service-negate"`
	// Collection of Network objects identified by the name or UID.
	Source string `json:"source,omitempty" xml:"source"`
	// True if negate is set for source.
	SourceNegate bool `json:"source-negate,omitempty" xml:"source-negate"`
	// List of time objects. For example: \"Weekend\", \"Off-Work\", \"Every-Day\".
	Time string `json:"time,omitempty" xml:"time"`
	Track *TrackSettingsForRequest `json:"track,omitempty" xml:"track"`
	UserCheck *UserCheckRequest `json:"user-check,omitempty" xml:"user-check"`
	// Communities or Directional.
	Vpn string `json:"vpn,omitempty" xml:"vpn"`
}
