package memox

// Shorthand for functions that do not receive arguments.
type func0 = func()
type func0R1[R1 any] = func() R1
type func0R2[R1, R2 any] = func() (R1, R2)

// Shorthand for functions that receive a single argument.
type func1[A1 any] = func(A1)
type func1R1[A1, R1 any] = func(A1) R1
type func1R2[A1, R1, R2 any] = func(A1) (R1, R2)

// Used to store the result of functions that return two values.
type tuple[A, B any] struct {
	A A
	B B
}
