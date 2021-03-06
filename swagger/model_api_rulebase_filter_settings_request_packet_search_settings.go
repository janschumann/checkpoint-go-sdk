/*
 * Check Point Management API
 *
 * Generated by https://github.com/janschumann/checkpoint-go-sdk.
 *
 * API version: v1.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// When 'search-mode' is set to 'packet', this object allows to set the packet search preferences.
type ApiRulebaseFilterSettingsRequestPacketSearchSettings struct {
	// When true, if the search expression contains a UID or a name of a group object, results will include rules that match on at least one member of the group.
	ExpandGroupMembers bool `json:"expand-group-members,omitempty" xml:"expand-group-members"`
	// When true, if the search expression contains a UID or a name of a group-with-exclusion object, results will include rules that match at least one member of the \"include\" part and is not a member of the \"except\" part.
	ExpandGroupWithExclusionMembers bool `json:"expand-group-with-exclusion-members,omitempty" xml:"expand-group-with-exclusion-members"`
	// Whether to match on 'Any' object.
	MatchOnAny bool `json:"match-on-any,omitempty" xml:"match-on-any"`
	// Whether to match on a group-with-exclusion.
	MatchOnGroupWithExclusion bool `json:"match-on-group-with-exclusion,omitempty" xml:"match-on-group-with-exclusion"`
	// Whether to match on a negated cell.
	MatchOnNegate bool `json:"match-on-negate,omitempty" xml:"match-on-negate"`
}
