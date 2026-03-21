package consolex

import (
	"fmt"
	"net/http"
	"net/url"
	"slices"
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

	ks := []string{}
	for k := range extras {
		ks = append(ks, k)
	}

	slices.Sort(ks)

	var sb strings.Builder
	for _, k := range ks {
		v := extras[k]
		fmt.Fprintf(&sb, "\n » %s: ", k)

		switch t := v.(type) {
		case http.Header:
			for k, v := range t {
				fmt.Fprintf(&sb, "\n  » %v: %v", k, v)
			}
		case url.Values:
			for k, v := range t {
				fmt.Fprintf(&sb, "\n  » %v: %v", k, v)
			}
		case map[string]any:
			for k, v := range t {
				fmt.Fprintf(&sb, "\n  » %v: %v", k, v)
			}
		default:
			fmt.Fprintf(&sb, "%v", t)
		}
	}

	return sb.String()
}
