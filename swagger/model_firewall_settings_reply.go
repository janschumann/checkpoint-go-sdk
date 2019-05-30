/*
 * Check Point Management API
 *
 * Generated by https://github.com/janschumann/checkpoint-go-sdk.
 *
 * API version: v1.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// N/A
type FirewallSettingsReply struct {
	// N/A
	AutoCalculateConnectionsHashTableSizeAndMemoryPool bool `json:"auto-calculate-connections-hash-table-size-and-memory-pool,omitempty" xml:"auto-calculate-connections-hash-table-size-and-memory-pool"`
	// N/A
	AutoMaximumLimitForConcurrentConnections bool `json:"auto-maximum-limit-for-concurrent-connections,omitempty" xml:"auto-maximum-limit-for-concurrent-connections"`
	// N/A
	ConnectionsHashSize int32 `json:"connections-hash-size,omitempty" xml:"connections-hash-size"`
	// N/A
	MaximumLimitForConcurrentConnections int32 `json:"maximum-limit-for-concurrent-connections,omitempty" xml:"maximum-limit-for-concurrent-connections"`
	// N/A
	MaximumMemoryPoolSize int32 `json:"maximum-memory-pool-size,omitempty" xml:"maximum-memory-pool-size"`
	// N/A
	MemoryPoolSize int32 `json:"memory-pool-size,omitempty" xml:"memory-pool-size"`
}