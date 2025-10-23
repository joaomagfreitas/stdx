package slicesx

// Fold applies the combine function to each element of input from an
// initial value.
func Fold[I, O any](input []I, initial O, combine func(O, I) O) O {
	output := initial
	for i := range input {
		output = combine(output, input[i])
	}

	return output
}
