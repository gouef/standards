package standards

type Logger interface {
	// Emergency System is unusable
	Emergency(message string, context []any) error
	// Alert Action must be taken immediately.
	Alert(message string, context []any) error
	// Critical Critical conditions.
	Critical(message string, context []any) error
	// Error Runtime errors that do not require immediate action but should typically
	Error(message string, context []any) error
	// Warning Exceptional occurrences that are not errors.
	Warning(message string, context []any) error
	// Notice Normal but significant events.
	Notice(message string, context []any) error
	// Info Interesting events
	Info(message string, context []any) error
	// Debug Detailed debug information.
	Debug(message string, context []any) error
	// Log Logs with an arbitrary level.
	Log(level LogLevel, message string, context []any) error
}
