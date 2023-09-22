package maps

func Keys[K comparable, V any](m map[K]V) []K {
	var result []K
	for k := range m {
		result = append(result, k)
	}
	return result
}

func Values[K comparable, V any](m map[K]V) []V {
	var result []V
	for _, v := range m {
		result = append(result, v)
	}
	return result
}
