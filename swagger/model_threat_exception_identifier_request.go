/*
 * Check Point Management API
 *
 * Generated by https://github.com/janschumann/checkpoint-go-sdk.
 *
 * API version: v1.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ThreatExceptionIdentifierRequest struct {
	// The level of detail for some of the fields in the response can vary from showing only the UID value of the object to a fully detailed representation of the object.
	DetailsLevel string `json:"details-level,omitempty" xml:"details-level"`
	// The name of the exception-group.
	ExceptionGroupName string `json:"exception-group-name,omitempty" xml:"exception-group-name"`
	// The UID of the exception-group.
	ExceptionGroupUid string `json:"exception-group-uid" xml:"exception-group-uid"`
	// The exception's number in the rulebase.
	ExceptionNumber int32 `json:"exception-number,omitempty" xml:"exception-number"`
	// Layer that the rule belongs to identified by the name or UID.
	Layer string `json:"layer" xml:"layer"`
	// The name of the exception.
	Name string `json:"name,omitempty" xml:"name"`
	// The name of the parent rule.
	RuleName string `json:"rule-name,omitempty" xml:"rule-name"`
	// The position in the rulebase of the parent rule.
	RuleNumber int32 `json:"rule-number,omitempty" xml:"rule-number"`
	// The UID of the parent rule.
	RuleUid string `json:"rule-uid" xml:"rule-uid"`
	// Object unique identifier.
	Uid string `json:"uid" xml:"uid"`
}