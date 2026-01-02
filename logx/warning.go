package logx

import (
	"fmt"
)

// Warning logs a message indicating a potential issue.
func Warning(msg string) {
	Warningx(msg, nil)
}

// Warning logs a message indicating a potential issue.
// Same as [Warning] but provides a message formatting API.
func Warningf(format string, values ...any) {
	Warningx(fmt.Sprintf(format, values...), nil)
}

// Warning logs a message indicating a potential issue.
// Same as [Warning] but includes structured context.
func Warningx(msg string, extras map[string]any) {
	for _, l := range loggers {
		l.Warning(msg, extras)
	}
}
