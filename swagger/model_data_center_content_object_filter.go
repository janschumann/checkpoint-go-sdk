/*
 * Check Point Management API
 *
 * Generated by https://github.com/janschumann/checkpoint-go-sdk.
 *
 * API version: v1.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Return results matching the specified filter.
type DataCenterContentObjectFilter struct {
	// Return results under the specified Data Center Object (identified by UID).
	ParentUidInDataCenter string `json:"parent-uid-in-data-center,omitempty" xml:"parent-uid-in-data-center"`
	// Return results containing the specified text value.
	Text string `json:"text,omitempty" xml:"text"`
	// Return results under the specified Data Center Object (identified by URI).
	Uri string `json:"uri,omitempty" xml:"uri"`
}
