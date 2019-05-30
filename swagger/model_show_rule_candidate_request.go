/*
 * Check Point Management API
 *
 * Generated by https://github.com/janschumann/checkpoint-go-sdk.
 *
 * API version: v1.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ShowRuleCandidateRequest struct {
	// Indicates whether to dereference \"members\" field by details level for every object in reply.
	DereferenceGroupMembers bool `json:"dereference-group-members,omitempty" xml:"dereference-group-members"`
	// The level of detail for some of the fields in the response can vary from showing only the UID value of the object to a fully detailed representation of the object.
	DetailsLevel string `json:"details-level,omitempty" xml:"details-level"`
	// Policy field name.<br>For access layer: destination, source, service, install-on, action, time, vpn, track, content<br>For nat layer: original-destination, original-source, original-service, translated-destination, translated-source, translated-service, install-on<br>For threat layer: destination, source, service, install-on, protected-scope, action, track<br>For exception: destination, source, service, install-on, protected-scope, action, track, protection-or-site.
	Field string `json:"field" xml:"field"`
	// Textual search expression to filter objects by.
	Filter string `json:"filter" xml:"filter"`
	// Layer or exception group name or uid.
	Layer string `json:"layer" xml:"layer"`
	// No more than that many results will be returned.
	Limit int32 `json:"limit,omitempty" xml:"limit"`
	// Skip that many results before beginning to return them.
	Offset int32 `json:"offset,omitempty" xml:"offset"`
	// Sorts results by the given field. By default the results are sorted in the ascending order by name.
	Order []ApiQueryOrderRequest `json:"order,omitempty" xml:"order"`
	// Indicates whether to calculate and show \"groups\" field for every object in reply.
	ShowMembership bool `json:"show-membership,omitempty" xml:"show-membership"`
	// Optional sub-field to the policy field.<br>For \"vpn\" field in access layer: all-gw-to-gw, specific<br>For \"service\" field in access layer: services, application-and-sites, categories, mobile-application<br>For \"content\" field in access layer: file-types<br>For \"protection-or-site\" field in exception: whitelist-files, ips-protections, anti-protection, user-applicatoin, blades.
	SubField string `json:"sub-field,omitempty" xml:"sub-field"`
}
