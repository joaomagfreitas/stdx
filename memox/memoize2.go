package memox

// Memoize2 creates a copy of input function that memorizes its invocation.
func Memoize2[A1, A2 any](fn func2[A1, A2]) func2[A1, A2] {
	c := map[uint64]any{}

	return func(a1 A1, a2 A2) {
		h := hash2(a1, a2)
		_, ok := c[h]

		if !ok {
			fn(a1, a2)
			c[h] = nil
		}
	}
}

// Applies [Memoize2] for a function that returns a single value.
func Memoize2R1[A1, A2, R1 any](fn func2R1[A1, A2, R1]) func2R1[A1, A2, R1] {
	c := map[uint64]R1{}

	return func(a1 A1, a2 A2) R1 {
		h := hash2(a1, a2)
		r1, ok := c[h]

		if !ok {
			r1 = fn(a1, a2)
			c[h] = r1
		}

		return r1
	}
}

// Applies [Memoize2] for a function that returns a single value.
func Memoize2R2[A1, A2, R1, R2 any](fn func2R2[A1, A2, R1, R2]) func2R2[A1, A2, R1, R2] {
	c := map[uint64]tuple[R1, R2]{}

	return func(a1 A1, a2 A2) (R1, R2) {
		h := hash2(a1, a2)
		r, ok := c[h]

		if !ok {
			r1, r2 := fn(a1, a2)
			r = tuple[R1, R2]{A: r1, B: r2}
			c[h] = r
		}

		return r.A, r.B
	}
}
