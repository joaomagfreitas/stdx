package logx

import (
	"fmt"
)

// Info logs a general informational message.
func Info(msg string) {
	Infox(msg, nil)
}

// Info logs a general informational message.
// Same as [Info] but provides a message formatting API.
func Infof(format string, values ...any) {
	Infox(fmt.Sprintf(format, values...), nil)
}

// Info logs a general informational message.
// Same as [Info] but includes structured context.
func Infox(msg string, extras map[string]any) {
	for _, l := range loggers {
		l.Info(msg, extras)
	}
}
