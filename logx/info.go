package logx

import (
	"fmt"
)

// Info logs a general informational message.
func Info(msg string) {
	Infox(msg, nil)
}

// Infof logs a general informational message.
// Same as [Info] but provides a message formatting API.
func Infof(format string, values ...any) {
	Infox(fmt.Sprintf(format, values...), nil)
}

// Infox logs a general informational message.
// Same as [Info] but includes structured context.
func Infox(msg string, extras map[string]any) {
	for _, l := range loggers {
		l.Info(msg, extras)
	}
}

// Info logs a general informational message.
// Same as [Info] but includes an error in structured context with key "error".
func Infoe(msg string, err error) {
	for _, l := range loggers {
		l.Info(msg, map[string]any{"error": err})
	}
}
