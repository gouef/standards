package standards

type LogLevel string

var (
	EMERGENCY LogLevel = "emergency"
	ALERT     LogLevel = "alert"
	CRITICAL  LogLevel = "critical"
	ERROR     LogLevel = "error"
	WARNING   LogLevel = "warning"
	NOTICE    LogLevel = "notice"
	INFO      LogLevel = "info"
	DEBUG     LogLevel = "debug"
)
