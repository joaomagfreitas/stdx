package slicesx

// Gen creates a new fixed-size slice capped to [count], and applies the
// [generate] function to resolve the element of each index.
func Gen[O any](count int, generate func(index int) O) []O {
	if count == 0 {
		return []O{}
	}

	output := make([]O, count)
	for i := range count {
		output[i] = generate(i)
	}

	return output
}
