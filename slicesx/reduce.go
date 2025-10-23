package slicesx

// Reduce applies the combine function to each element of input.
func Reduce[I any](input []I, combine func(previous I, current I) I) I {
	if len(input) > 0 {
		return Fold(input[1:], input[0], combine)
	}

	var empty I
	return empty
}
