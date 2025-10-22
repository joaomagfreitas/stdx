package slicesx

// First applies the predicate function to each element of input and
// returns the first item that matches the predicate.
func First[I any](input []I, predicate func(I) bool) (I, bool) {
	for _, item := range input {
		if !predicate(item) {
			continue
		}

		return item, true
	}

	var item I
	return item, false
}

// Last applies the predicate function to each element of input and
// returns the last item that matches the predicate.
func Last[I any](input []I, predicate func(I) bool) (I, bool) {
	l := len(input)
	for l > 0 {
		l--

		if !predicate(input[l]) {
			continue
		}

		return input[l], true
	}

	var item I
	return item, false
}
