package errorsx

// Swallow marks input error as consumed, ignoring the "err not checked" warning
// reported by linters.
// Use it wisely.
func Swallow(err error) { /* NO-OP */ }
