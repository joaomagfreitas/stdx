package slicesx

// Map applies the transform function to each element of input and
// returns a new slice containing the transformed results.
//
// It preserves the order of elements and allocates a new slice
// of the same length as input.
func Map[I, O any](input []I, transform func(I) O) []O {
	if input == nil {
		return nil
	}

	output := make([]O, len(input))
	for i := range input {
		output[i] = transform(input[i])
	}

	return output
}
