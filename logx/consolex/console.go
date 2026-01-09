package consolex

import "encoding/json"

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

	b, err := json.Marshal(extras)
	if err != nil {
		return err.Error()
	}

	return string(b)
}
