package standards

type LoggerNull struct {
}

// Emergency System is unusable
func (l *LoggerNull) Emergency(message string, context []any) error {
	return l.Log(EMERGENCY, message, context)
}

// Alert Action must be taken immediately.
func (l *LoggerNull) Alert(message string, context []any) error {
	return l.Log(ALERT, message, context)
}

// Critical Critical conditions.
func (l *LoggerNull) Critical(message string, context []any) error {
	return l.Log(CRITICAL, message, context)
}

// Error Runtime errors that do not require immediate action but should typically
func (l *LoggerNull) Error(message string, context []any) error {
	return l.Log(ERROR, message, context)
}

// Warning Exceptional occurrences that are not errors.
func (l *LoggerNull) Warning(message string, context []any) error {
	return l.Log(WARNING, message, context)
}

// Notice Normal but significant events.
func (l *LoggerNull) Notice(message string, context []any) error {
	return l.Log(NOTICE, message, context)
}

// Info Interesting events
func (l *LoggerNull) Info(message string, context []any) error {
	return l.Log(INFO, message, context)
}

// Debug Detailed debug information.
func (l *LoggerNull) Debug(message string, context []any) error {
	return l.Log(DEBUG, message, context)
}

// Log Logs with an arbitrary level.
func (l *LoggerNull) Log(level LogLevel, message string, context []any) error {
	return nil
}
