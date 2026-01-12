package consolex

import (
	"fmt"
	"strings"
)

type ConsoleLogger struct{}

func (ConsoleLogger) Trace(msg string, extras map[string]any) {
	println("[TRACE]", msg, formatExtras(extras))
}

func (ConsoleLogger) Info(msg string, extras map[string]any) {
	println("[INFO]", msg, formatExtras(extras))
}

func (ConsoleLogger) Warning(msg string, extras map[string]any) {
	println("[WARN]", msg, formatExtras(extras))
}

func (ConsoleLogger) Critical(msg string, extras map[string]any) {
	println("[CRITICAL]", msg, formatExtras(extras))
}

func (ConsoleLogger) Error(err error, extras map[string]any) {
	println("[ERROR]", err.Error(), formatExtras(extras))
}

func formatExtras(extras map[string]any) string {
	if len(extras) == 0 {
		return ""
	}

	var sb strings.Builder
	for k, v := range extras {
		sb.WriteString(fmt.Sprintf("\n » %s: \"%s\"", k, v))
	}

	return sb.String()
}
