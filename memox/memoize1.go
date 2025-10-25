package memox

// Memoize1 creates a copy of input function that memorizes its invocation.
func Memoize1[A1 any](fn func1[A1]) func1[A1] {
	c := map[uint64]any{}

	return func(a1 A1) {
		h := fnv64a(a1)
		_, ok := c[h]

		if !ok {
			fn(a1)
			c[h] = nil
		}
	}
}

// Applies [Memoize1] for a function that returns a single value.
func Memoize1R1[A1, R1 any](fn func1R1[A1, R1]) func1R1[A1, R1] {
	c := map[uint64]R1{}

	return func(a1 A1) R1 {
		h := fnv64a(a1)
		r1, ok := c[h]

		if !ok {
			r1 = fn(a1)
			c[h] = r1
		}

		return r1
	}
}

// Applies [Memoize2] for a function that returns a single value.
func Memoize1R2[A1, R1, R2 any](fn func1R2[A1, R1, R2]) func1R2[A1, R1, R2] {
	c := map[uint64]tuple[R1, R2]{}

	return func(a1 A1) (R1, R2) {
		h := fnv64a(a1)
		r, ok := c[h]

		if !ok {
			r1, r2 := fn(a1)
			r = tuple[R1, R2]{A: r1, B: r2}
			c[h] = r
		}

		return r.A, r.B
	}
}
