/*
 * Check Point Management API
 *
 * Generated by https://github.com/janschumann/checkpoint-go-sdk.
 *
 * API version: v1.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Starting time. Note: Each gateway may interpret this time differently according to its time zone.
type TimeObject struct {
	// Date in format dd-MMM-yyyy.
	Date string `json:"date,omitempty" xml:"date"`
	// Date and time represented in international ISO 8601 format.
	Iso8601 string `json:"iso-8601,omitempty" xml:"iso-8601"`
	// Number of milliseconds that have elapsed since 00:00:00, 1 January 1970.
	Posix int32 `json:"posix,omitempty" xml:"posix"`
	// Time in format HH:mm.
	Time string `json:"time,omitempty" xml:"time"`
}
