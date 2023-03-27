package sliceutils

func Unique[E comparable](input []E) []E {

	u := make([]E, 0, len(input))
	m := make(map[E]bool)

	for _, val := range input {
		if _, ok := m[val]; !ok {
			m[val] = true
			u = append(u, val)
		}
	}

	return u
}

func Index[E comparable](s []E, v E) int {
	for i, vs := range s {
		if v == vs {
			return i
		}
	}
	return -1
}

func Map[E any, V any](slice []E, mapper func(E) V) []V {
	result := make([]V, 0, len(slice))
	for _, value := range slice {
		result = append(result, mapper(value))
	}
	return result
}

func Find[V any](slice []V, predicate func(V) bool, defaultValue V) V {
	for _, value := range slice {
		if predicate(value) {
			return value
		}
	}
	return defaultValue
}

func Filter[V any](slice []V, predicate func(V) bool) []V {
	result := make([]V, 0)
	for _, value := range slice {
		if predicate(value) {
			result = append(result, value)
		}
	}
	return result
}

func Contains[V comparable](slice []V, value V) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}
