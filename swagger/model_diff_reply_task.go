/*
 * Check Point Management API
 *
 * Generated by https://github.com/janschumann/checkpoint-go-sdk.
 *
 * API version: v1.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type DiffReplyTask struct {
	// Diff task UID. Use show-task command to check the progress of the task.
	TaskId string `json:"task-id,omitempty" xml:"task-id"`
}