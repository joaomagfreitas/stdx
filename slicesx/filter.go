package slicesx

// Filter applies the predicate function to each element of input and
// returns a new slice containing only the items that match the predicate.
func Filter[I any](input []I, predicate func(I) bool) []I {
	if input == nil {
		return nil
	}

	output := []I{}
	for _, item := range input {
		if !predicate(item) {
			continue
		}

		output = append(output, item)
	}

	return output
}
