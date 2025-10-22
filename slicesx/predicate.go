package slicesx

// Any applies the predicate function to each element of input and
// returns true if at least one of the items match the predicate.
func Any[I any](input []I, predicate func(I) bool) bool {
	for _, item := range input {
		if !predicate(item) {
			continue
		}

		return true
	}

	return false
}

// All applies the predicate function to each element of input and
// returns true if all the items match the predicate.
func All[I any](input []I, predicate func(I) bool) bool {
	if len(input) == 0 {
		return false
	}

	for _, item := range input {
		if !predicate(item) {
			return false
		}
	}

	return true
}
