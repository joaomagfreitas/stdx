package memox

// Memoize creates a copy of input function that memorizes its invocation.
func Memoize(fn func0) func0 {
	ok := false
	return func() {
		if !ok {
			fn()
			ok = true
		}
	}
}

// Applies [Memoize] for a function that returns a single value.
func MemoizeR1[R1 any](fn func0R1[R1]) func0R1[R1] {
	var r1 R1
	ok := false

	return func() R1 {
		if !ok {
			r1 = fn()
			ok = true
		}

		return r1
	}
}

// Applies [Memoize] for a function that returns two values.
func MemoizeR2[R1, R2 any](fn func0R2[R1, R2]) func0R2[R1, R2] {
	var r1 R1
	var r2 R2
	ok := false

	return func() (R1, R2) {
		if !ok {
			r1, r2 = fn()
			ok = true
		}

		return r1, r2
	}
}
