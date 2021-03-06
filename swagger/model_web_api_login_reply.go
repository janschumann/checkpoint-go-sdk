/*
 * Check Point Management API
 *
 * Generated by https://github.com/janschumann/checkpoint-go-sdk.
 *
 * API version: v1.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type WebApiLoginReply struct {
	// API Server version.
	ApiServerVersion string `json:"api-server-version,omitempty" xml:"api-server-version"`
	// Information about the available disk space on the management server.
	DiskSpaceMessage string `json:"disk-space-message,omitempty" xml:"disk-space-message"`
	LastLoginWasAt *ApiDateReply `json:"last-login-was-at,omitempty" xml:"last-login-was-at"`
	LoginMessage *LoginMessageReply `json:"login-message,omitempty" xml:"login-message"`
	// True if this session is read only.
	ReadOnly bool `json:"read-only,omitempty" xml:"read-only"`
	// Session expiration timeout in seconds.
	SessionTimeout int32 `json:"session-timeout,omitempty" xml:"session-timeout"`
	// Session unique identifier. Enter this session unique identifier in the 'X-chkp-sid' header of each request.
	Sid string `json:"sid,omitempty" xml:"sid"`
	// True if this management server is in the standby mode.
	Standby bool `json:"standby,omitempty" xml:"standby"`
	// Session object unique identifier. This identifier may be used in the discard API to discard changes that were made in this session, when administrator is working from another session, or in the 'switch-session' API.
	Uid string `json:"uid,omitempty" xml:"uid"`
	// URL that was used to reach the API server.
	Url string `json:"url,omitempty" xml:"url"`
}
