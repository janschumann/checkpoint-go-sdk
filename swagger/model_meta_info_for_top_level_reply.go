/*
 * Check Point Management API
 *
 * Generated by https://github.com/janschumann/checkpoint-go-sdk.
 *
 * API version: v1.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Object metadata.
type MetaInfoForTopLevelReply struct {
	CreationTime *ApiDateReply `json:"creation-time,omitempty" xml:"creation-time"`
	// N/A
	Creator string `json:"creator,omitempty" xml:"creator"`
	// N/A
	LastModifier string `json:"last-modifier,omitempty" xml:"last-modifier"`
	LastModifyTime *ApiDateReply `json:"last-modify-time,omitempty" xml:"last-modify-time"`
	// Object lock state. It's not allowed to edit objects locked by other session.
	Lock string `json:"lock,omitempty" xml:"lock"`
	// N/A
	ValidationState string `json:"validation-state,omitempty" xml:"validation-state"`
}
