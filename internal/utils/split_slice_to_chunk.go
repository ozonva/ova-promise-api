package utils

// SplitSliceToChunkStrings splits a slice of strings into a slice of slices of strings of the chunkSize length.
// If chunkSize < 1 then chunkSize = 1.
func SplitSliceToChunkStrings(initial []string, chunkSize int) [][]string {
	var result [][]string

	if len(initial) == 0 {
		return result
	}

	// it's not ub, it's documented behaviour
	if chunkSize < 1 {
		chunkSize = 1
	}

	for i := 0; i < len(initial); i += chunkSize {
		end := i + chunkSize

		if end > len(initial) {
			end = len(initial)
		}

		result = append(result, initial[i:end])
	}

	return result
}

// SplitSliceToChunkIntegers splits a slice of integers into a slice of slices of integers of the chunkSize length.
// If chunkSize < 1 then chunkSize = 1.
func SplitSliceToChunkIntegers(initial []int, chunkSize int) [][]int {
	var result [][]int

	if len(initial) == 0 {
		return result
	}

	// it's not ub, it's documented behaviour
	if chunkSize < 1 {
		chunkSize = 1
	}

	for i := 0; i < len(initial); i += chunkSize {
		end := i + chunkSize

		if end > len(initial) {
			end = len(initial)
		}

		result = append(result, initial[i:end])
	}

	return result
}

// SplitSliceToChunk splits a slice of interface into a slice of slices of interface of the chunkSize length.
// If chunkSize < 1 then chunkSize = 1.
func SplitSliceToChunk(initial []interface{}, chunkSize int) [][]interface{} {
	var result [][]interface{}

	if len(initial) == 0 {
		return result
	}

	// it's not ub, it's documented behaviour
	if chunkSize < 1 {
		chunkSize = 1
	}

	for i := 0; i < len(initial); i += chunkSize {
		end := i + chunkSize

		if end > len(initial) {
			end = len(initial)
		}

		result = append(result, initial[i:end])
	}

	return result
}
