package errorsx

// Must returns the provided value or panics if the given error is non-nil.
// It is a convenience for simplifying error handling in initialization code.
func Must[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}

	return value
}
