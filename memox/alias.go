package memox

// Shorthand for a function that does not accept or return values.
type func0 = func()

// Shorthand for a function that does not accept arguments, but returns
// a single value.
type func0R1[R1 any] = func() R1

// Shorthand for a function that does not accept arguments, but returns
// a two values.
type func0R2[R1, R2 any] = func() (R1, R2)
