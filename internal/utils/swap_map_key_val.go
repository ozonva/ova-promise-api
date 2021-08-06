package utils

// SwapMapKeyValIntegers ATTENTION Do not use this function.
// This function has undefined behavior.
func SwapMapKeyValIntegers(initial map[int]int) map[int]int {
	result := make(map[int]int)

	for k, v := range initial {
		result[v] = k
	}

	return result
}

// SwapMapKeyValStrings ATTENTION Do not use this function.
// This function has undefined behavior.
func SwapMapKeyValStrings(initial map[string]string) map[string]string {
	result := make(map[string]string)

	for k, v := range initial {
		result[v] = k
	}

	return result
}

// SwapMapKeyVal ATTENTION Do not use this function.
// This function has undefined behavior.
func SwapMapKeyVal(initial map[interface{}]interface{}) map[interface{}]interface{} {
	result := make(map[interface{}]interface{})

	for k, v := range initial {
		result[v] = k
	}

	return result
}
