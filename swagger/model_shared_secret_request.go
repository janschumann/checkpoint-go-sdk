/*
 * Check Point Management API
 *
 * Generated by https://github.com/janschumann/checkpoint-go-sdk.
 *
 * API version: v1.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Shared secrets for external gateways.
type SharedSecretRequest struct {
	// External gateway identified by the name or UID.
	ExternalGateway string `json:"external-gateway,omitempty" xml:"external-gateway"`
	// Shared secret.
	SharedSecret string `json:"shared-secret,omitempty" xml:"shared-secret"`
}
