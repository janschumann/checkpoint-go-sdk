/*
 * Check Point Management API
 *
 * Generated by https://github.com/janschumann/checkpoint-go-sdk.
 *
 * API version: v1.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type TaskEntityReply struct {
	// Color of the object. Should be one of existing colors.
	Color string `json:"color,omitempty" xml:"color"`
	// Comments string.
	Comments string `json:"comments,omitempty" xml:"comments"`
	Domain *ApiDomainIdentifier `json:"domain,omitempty" xml:"domain"`
	// Object icon.
	Icon string `json:"icon,omitempty" xml:"icon"`
	LastUpdateTime *ApiDateReply `json:"last-update-time,omitempty" xml:"last-update-time"`
	MetaInfo *MetaInfoForTopLevelReply `json:"meta-info,omitempty" xml:"meta-info"`
	// Object name. Should be unique in the domain.
	Name string `json:"name,omitempty" xml:"name"`
	// N/A
	ProgressDescription string `json:"progress-description,omitempty" xml:"progress-description"`
	// N/A
	ProgressPercentage int32 `json:"progress-percentage,omitempty" xml:"progress-percentage"`
	// Indicates whether the object is read-only.
	ReadOnly bool `json:"read-only,omitempty" xml:"read-only"`
	StartTime *ApiDateReply `json:"start-time,omitempty" xml:"start-time"`
	// Task status.
	Status string `json:"status,omitempty" xml:"status"`
	// N/A
	Suppressed bool `json:"suppressed,omitempty" xml:"suppressed"`
	// Collection of tag objects identified by the name or UID. How much details are returned depends on the details-level field of the request. This table shows the level of detail shown when details-level is set to standard.
	Tags []ApiObjectStandardIdentifier `json:"tags,omitempty" xml:"tags"`
	// Task-specific details according to the requested task type.
	TaskDetails []interface{} `json:"task-details,omitempty" xml:"task-details"`
	// Asynchronous task unique identifier.
	TaskId string `json:"task-id,omitempty" xml:"task-id"`
	// N/A
	TaskName string `json:"task-name,omitempty" xml:"task-name"`
	// Type of the object.
	Type_ string `json:"type,omitempty" xml:"type"`
	// Object unique identifier.
	Uid string `json:"uid,omitempty" xml:"uid"`
}
