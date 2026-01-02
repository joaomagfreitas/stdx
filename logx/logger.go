package logx

// Logger defines a minimal structured logging interface with common severity levels.
// Implementations are expected to handle message formatting, severity filtering,
// and optional structured context provided via [extras].
type Logger interface {
	// Trace logs a low-level diagnostic message.
	Trace(msg string, extras map[string]any)

	// Info logs a general informational message.
	Info(msg string, extras map[string]any)

	// Warning logs a message indicating a potential issue.
	Warning(msg string, extras map[string]any)

	// Critical logs a message indicating a serious failure requiring immediate attention.
	Critical(msg string, extras map[string]any)

	// Error logs a program error.
	Error(err error, extras map[string]any)
}
