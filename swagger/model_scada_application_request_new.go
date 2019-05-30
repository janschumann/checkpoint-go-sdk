/*
 * Check Point Management API
 *
 * Generated by https://github.com/janschumann/checkpoint-go-sdk.
 *
 * API version: v1.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ScadaApplicationRequestNew struct {
	// Used to configure or edit the additional categories of a custom application / site used in the Application and URL Filtering or Threat Prevention.
	AdditionalCategories string `json:"additional-categories,omitempty" xml:"additional-categories"`
	// N/A
	ApplicationSignature string `json:"application-signature,omitempty" xml:"application-signature"`
	// SCADA application category.
	Category string `json:"category" xml:"category"`
	// Color of the object. Should be one of existing colors.
	Color string `json:"color,omitempty" xml:"color"`
	// Comments string.
	Comments string `json:"comments,omitempty" xml:"comments"`
	// A description for the application.
	Description string `json:"description,omitempty" xml:"description"`
	// The level of detail for some of the fields in the response can vary from showing only the UID value of the object to a fully detailed representation of the object.
	DetailsLevel string `json:"details-level,omitempty" xml:"details-level"`
	// Collection of group identifiers.
	Groups string `json:"groups,omitempty" xml:"groups"`
	// Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.
	IgnoreErrors bool `json:"ignore-errors,omitempty" xml:"ignore-errors"`
	// Apply changes ignoring warnings.
	IgnoreWarnings bool `json:"ignore-warnings,omitempty" xml:"ignore-warnings"`
	// Object name. Should be unique in the domain.
	Name string `json:"name" xml:"name"`
	// SCADA application properties.
	ScadaProperties []ScadaApplicationEntryItem `json:"scada-properties,omitempty" xml:"scada-properties"`
	// Collection of tag identifiers.
	Tags string `json:"tags,omitempty" xml:"tags"`
}
