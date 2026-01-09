package logx

import (
	"fmt"
)

// Trace logs a low-level diagnostic message.
func Trace(msg string) {
	Tracex(msg, nil)
}

// Trace logs a low-level diagnostic message.
// Same as [Trace] but provides a message formatting API.
func Tracef(format string, values ...any) {
	Tracex(fmt.Sprintf(format, values...), nil)
}

// Trace logs a low-level diagnostic message.
// Same as [Trace] but includes structured context.
func Tracex(msg string, extras map[string]any) {
	for _, l := range loggers {
		l.Trace(msg, extras)
	}
}

// Trace logs a low-level diagnostic message.
// Same as [Trace] but includes an error in structured context with key "error".
func Tracee(msg string, err error) {
	for _, l := range loggers {
		l.Trace(msg, map[string]any{"error": err})
	}
}
