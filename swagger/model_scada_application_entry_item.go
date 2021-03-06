/*
 * Check Point Management API
 *
 * Generated by https://github.com/janschumann/checkpoint-go-sdk.
 *
 * API version: v1.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ScadaApplicationEntryItem struct {
	// SCADA property key.
	Key string `json:"key,omitempty" xml:"key"`
	// SCADA property value.
	Value string `json:"value,omitempty" xml:"value"`
}
