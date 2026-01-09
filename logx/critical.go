package logx

import (
	"fmt"
)

// Critical logs a message indicating a serious failure requiring immediate attention.
func Critical(msg string) {
	Criticalx(msg, nil)
}

// Critical logs a message indicating a serious failure requiring immediate attention.
// Same as [Critical] but provides a message formatting API.
func Criticalf(format string, values ...any) {
	Criticalx(fmt.Sprintf(format, values...), nil)
}

// Critical logs a message indicating a serious failure requiring immediate attention.
// Same as [Critical] but includes structured context.
func Criticalx(msg string, extras map[string]any) {
	for _, l := range loggers {
		l.Critical(msg, extras)
	}
}

// Critical logs a message indicating a serious failure requiring immediate attention.
// Same as [Critical] but includes an error in structured context with key "error".
func Criticale(msg string, err error) {
	for _, l := range loggers {
		l.Critical(msg, map[string]any{"error": err})
	}
}
