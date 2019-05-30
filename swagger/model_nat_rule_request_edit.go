/*
 * Check Point Management API
 *
 * Generated by https://github.com/janschumann/checkpoint-go-sdk.
 *
 * API version: v1.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type NatRuleRequestEdit struct {
	// Comments string.
	Comments string `json:"comments,omitempty" xml:"comments"`
	// The level of detail for some of the fields in the response can vary from showing only the UID value of the object to a fully detailed representation of the object.
	DetailsLevel string `json:"details-level,omitempty" xml:"details-level"`
	// Enable/Disable the rule.
	Enabled bool `json:"enabled,omitempty" xml:"enabled"`
	// Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.
	IgnoreErrors bool `json:"ignore-errors,omitempty" xml:"ignore-errors"`
	// Apply changes ignoring warnings.
	IgnoreWarnings bool `json:"ignore-warnings,omitempty" xml:"ignore-warnings"`
	InstallOn *Add `json:"install-on,omitempty" xml:"install-on"`
	// Nat method.
	Method string `json:"method,omitempty" xml:"method"`
	// New position in the rulebase.
	NewPosition int32 `json:"new-position,omitempty" xml:"new-position"`
	// Original destination.
	OriginalDestination string `json:"original-destination,omitempty" xml:"original-destination"`
	// Original service.
	OriginalService string `json:"original-service,omitempty" xml:"original-service"`
	// Original source.
	OriginalSource string `json:"original-source,omitempty" xml:"original-source"`
	// Name of the package.
	Package_ string `json:"package" xml:"package"`
	// Rule number.
	RuleNumber string `json:"rule-number,omitempty" xml:"rule-number"`
	// Translated  destination.
	TranslatedDestination string `json:"translated-destination,omitempty" xml:"translated-destination"`
	// Translated  service.
	TranslatedService string `json:"translated-service,omitempty" xml:"translated-service"`
	// Translated  source.
	TranslatedSource string `json:"translated-source,omitempty" xml:"translated-source"`
	// Object unique identifier.
	Uid string `json:"uid" xml:"uid"`
}
