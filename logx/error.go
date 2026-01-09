package logx

// Error logs a program error.
func Error(err error) {
	Errorx(err, nil)
}

// Error logs a program error.
// Same as [Error] but includes structured context.
func Errorx(err error, extras map[string]any) {
	for _, l := range loggers {
		l.Error(err, extras)
	}
}

// Error logs a program error.
// Same as [Error] but includes the error in structured context with key "error".
func Errore(err error) {
	for _, l := range loggers {
		l.Error(err, map[string]any{"error": err})
	}
}
