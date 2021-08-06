package utils

// FilterSliceByExcludedIntegers returns new slice without excluded elements.
func FilterSliceByExcludedIntegers(initial, excluded []int) []int {
	// make and fill a set
	excludedSet := make(map[int]interface{})
	for _, val := range excluded {
		excludedSet[val] = nil
	}

	var result []int

	for _, val := range initial {
		if _, ok := excludedSet[val]; !ok {
			result = append(result, val)
		}
	}

	return result
}

// FilterSliceByExcludedStrings returns new slice without excluded elements.
func FilterSliceByExcludedStrings(initial, excluded []string) []string {
	// make and fill a set
	excludedSet := make(map[string]interface{})

	for _, val := range excluded {
		excludedSet[val] = nil
	}

	var result []string

	for _, val := range initial {
		if _, ok := excludedSet[val]; !ok {
			result = append(result, val)
		}
	}

	return result
}

// FilterSliceByExcluded returns new slice without excluded elements.
func FilterSliceByExcluded(initial, excluded []interface{}) []interface{} {
	// make and fill a set
	excludedSet := make(map[interface{}]interface{})

	for _, val := range excluded {
		excludedSet[val] = nil
	}

	var result []interface{}

	for _, val := range initial {
		if _, ok := excludedSet[val]; !ok {
			result = append(result, val)
		}
	}

	return result
}
