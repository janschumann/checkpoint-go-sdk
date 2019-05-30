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
type LogsSettingsRequest struct {
	// N/A
	AlertWhenFreeDiskSpaceBelow bool `json:"alert-when-free-disk-space-below,omitempty" xml:"alert-when-free-disk-space-below"`
	// N/A
	AlertWhenFreeDiskSpaceBelowThreshold int32 `json:"alert-when-free-disk-space-below-threshold,omitempty" xml:"alert-when-free-disk-space-below-threshold"`
	// N/A
	AlertWhenFreeDiskSpaceBelowType string `json:"alert-when-free-disk-space-below-type,omitempty" xml:"alert-when-free-disk-space-below-type"`
	// N/A
	BeforeDeleteKeepLogsFromTheLastDays bool `json:"before-delete-keep-logs-from-the-last-days,omitempty" xml:"before-delete-keep-logs-from-the-last-days"`
	// N/A
	BeforeDeleteKeepLogsFromTheLastDaysThreshold int32 `json:"before-delete-keep-logs-from-the-last-days-threshold,omitempty" xml:"before-delete-keep-logs-from-the-last-days-threshold"`
	// N/A
	BeforeDeleteRunScript bool `json:"before-delete-run-script,omitempty" xml:"before-delete-run-script"`
	// N/A
	BeforeDeleteRunScriptCommand string `json:"before-delete-run-script-command,omitempty" xml:"before-delete-run-script-command"`
	// N/A
	DeleteIndexFilesOlderThanDays bool `json:"delete-index-files-older-than-days,omitempty" xml:"delete-index-files-older-than-days"`
	// N/A
	DeleteIndexFilesOlderThanDaysThreshold int32 `json:"delete-index-files-older-than-days-threshold,omitempty" xml:"delete-index-files-older-than-days-threshold"`
	// N/A
	DeleteIndexFilesWhenIndexSizeAbove bool `json:"delete-index-files-when-index-size-above,omitempty" xml:"delete-index-files-when-index-size-above"`
	// N/A
	DeleteIndexFilesWhenIndexSizeAboveThreshold int32 `json:"delete-index-files-when-index-size-above-threshold,omitempty" xml:"delete-index-files-when-index-size-above-threshold"`
	// N/A
	DeleteWhenFreeDiskSpaceBelow bool `json:"delete-when-free-disk-space-below,omitempty" xml:"delete-when-free-disk-space-below"`
	// N/A
	DeleteWhenFreeDiskSpaceBelowThreshold int32 `json:"delete-when-free-disk-space-below-threshold,omitempty" xml:"delete-when-free-disk-space-below-threshold"`
	// N/A
	DetectNewCitrixIcaApplicationNames bool `json:"detect-new-citrix-ica-application-names,omitempty" xml:"detect-new-citrix-ica-application-names"`
	// N/A
	ForwardLogsToLogServer bool `json:"forward-logs-to-log-server,omitempty" xml:"forward-logs-to-log-server"`
	// N/A
	ForwardLogsToLogServerName string `json:"forward-logs-to-log-server-name,omitempty" xml:"forward-logs-to-log-server-name"`
	// N/A
	ForwardLogsToLogServerScheduleName string `json:"forward-logs-to-log-server-schedule-name,omitempty" xml:"forward-logs-to-log-server-schedule-name"`
	// N/A
	FreeDiskSpaceMetrics string `json:"free-disk-space-metrics,omitempty" xml:"free-disk-space-metrics"`
	// N/A
	PerformLogRotateBeforeLogForwarding bool `json:"perform-log-rotate-before-log-forwarding,omitempty" xml:"perform-log-rotate-before-log-forwarding"`
	// N/A
	RejectConnectionsWhenFreeDiskSpaceBelowThreshold bool `json:"reject-connections-when-free-disk-space-below-threshold,omitempty" xml:"reject-connections-when-free-disk-space-below-threshold"`
	// N/A
	ReserveForPacketCaptureMetrics string `json:"reserve-for-packet-capture-metrics,omitempty" xml:"reserve-for-packet-capture-metrics"`
	// N/A
	ReserveForPacketCaptureThreshold int32 `json:"reserve-for-packet-capture-threshold,omitempty" xml:"reserve-for-packet-capture-threshold"`
	// N/A
	RotateLogByFileSize bool `json:"rotate-log-by-file-size,omitempty" xml:"rotate-log-by-file-size"`
	// N/A
	RotateLogFileSizeThreshold int32 `json:"rotate-log-file-size-threshold,omitempty" xml:"rotate-log-file-size-threshold"`
	// N/A
	RotateLogOnSchedule bool `json:"rotate-log-on-schedule,omitempty" xml:"rotate-log-on-schedule"`
	// N/A
	RotateLogScheduleName string `json:"rotate-log-schedule-name,omitempty" xml:"rotate-log-schedule-name"`
	// N/A
	StopLoggingWhenFreeDiskSpaceBelow bool `json:"stop-logging-when-free-disk-space-below,omitempty" xml:"stop-logging-when-free-disk-space-below"`
	// N/A
	StopLoggingWhenFreeDiskSpaceBelowThreshold int32 `json:"stop-logging-when-free-disk-space-below-threshold,omitempty" xml:"stop-logging-when-free-disk-space-below-threshold"`
	// N/A
	TurnOnQosLogging bool `json:"turn-on-qos-logging,omitempty" xml:"turn-on-qos-logging"`
	// N/A
	UpdateAccountLogEvery int32 `json:"update-account-log-every,omitempty" xml:"update-account-log-every"`
}