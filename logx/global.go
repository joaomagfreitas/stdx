package logx

// Holds the registered [Logger] implementations that will be used on the top level logging calls.
var loggers = []Logger{}

// WithLoggers replaces the current global set of [Logger] used in top level logging calls.
func WithLoggers(
	logger ...Logger,
) {
	loggers = append([]Logger{}, logger...)
}
