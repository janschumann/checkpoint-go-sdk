/*
 * Check Point Management API
 *
 * Generated by https://github.com/janschumann/checkpoint-go-sdk.
 *
 * API version: v1.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type IpsUpdateScheduleReply struct {
	// IPS Update Schedule status.
	Enabled bool `json:"enabled,omitempty" xml:"enabled"`
	Recurrence *IpsUpdateScheduleDayRecurrence `json:"recurrence,omitempty" xml:"recurrence"`
	// Time in format HH:mm.
	Time string `json:"time,omitempty" xml:"time"`
}
