/*
 * Check Point Management API
 *
 * Generated by https://github.com/janschumann/checkpoint-go-sdk.
 *
 * API version: v1.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type WebApiLoginRequest struct {
	// When 'continue-last-session' is set to 'True', the new session would continue where the last session was stopped. This option is available when the administrator has only one session that can be continued. If there is more than one session, see 'switch-session' API.
	ContinueLastSession bool `json:"continue-last-session,omitempty" xml:"continue-last-session"`
	// Use domain to login to specific domain. Domain can be identified by name or UID.
	Domain string `json:"domain,omitempty" xml:"domain"`
	// Login to the last published session. Such login is done with the Read Only permissions.
	EnterLastPublishedSession bool `json:"enter-last-published-session,omitempty" xml:"enter-last-published-session"`
	// Administrator password.
	Password string `json:"password" xml:"password"`
	// Login with Read Only permissions. This parameter is not considered in case continue-last-session is true.
	ReadOnly bool `json:"read-only,omitempty" xml:"read-only"`
	// Session comments.
	SessionComments string `json:"session-comments,omitempty" xml:"session-comments"`
	// Session description.
	SessionDescription string `json:"session-description,omitempty" xml:"session-description"`
	// Session unique name.
	SessionName string `json:"session-name,omitempty" xml:"session-name"`
	// Session expiration timeout in seconds. Default 600 seconds.
	SessionTimeout int32 `json:"session-timeout,omitempty" xml:"session-timeout"`
	// Administrator user name.
	User string `json:"user" xml:"user"`
}
