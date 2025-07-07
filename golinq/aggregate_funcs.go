package golinq

// Map mapping each element due transform function []T -> []R
func Map[T any, R any](input []T, transform func(T) R) []R {
	output := make([]R, 0, len(input))
	for _, v := range input {
		output = append(output, transform(v))
	}
	return output
}

func Where[T any](input []T, predicate func(T) bool) []T {
	output := make([]T, 0, len(input))
	for _, v := range input {
		if predicate(v) {
			output = append(output, v)
		}
	}
	return output
}

func Any[T any](input []T, predicate func(T) bool) bool {
	for _, v := range input {
		if predicate(v) {
			return true
		}
	}
	return false
}
