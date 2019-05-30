/*
 * Check Point Management API
 *
 * Generated by https://github.com/janschumann/checkpoint-go-sdk.
 *
 * API version: v1.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Servers Configuration.
type HostServersReply struct {
	// Gets True if this server is a DNS Server.
	DnsServer bool `json:"dns-server,omitempty" xml:"dns-server"`
	// Gets True if this server is a Mail Server.
	MailServer bool `json:"mail-server,omitempty" xml:"mail-server"`
	// Gets True if this server is a Web Server.
	WebServer bool `json:"web-server,omitempty" xml:"web-server"`
	WebServerConfig *WebServerReply `json:"web-server-config,omitempty" xml:"web-server-config"`
}
