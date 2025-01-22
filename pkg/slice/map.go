package slice

func Map[I, O any](slice []I, fn func(value I) O) []O {
	result := make([]O, len(slice))
	for index, value := range slice {
		result[index] = fn(value)
	}
	return result
}

func MapToPtr[I, O any](slice []I, fn func(value I) O) []*O {
	result := make([]*O, len(slice))
	for index, value := range slice {
		mapped := fn(value)
		result[index] = &mapped
	}
	return result
}

func MapNonPtrToPtr[I, O any](slice []I, fn func(*I) *O) []*O {
	result := make([]*O, len(slice))
	for index, value := range slice {
		v := value // Create new variable to get correct pointer
		result[index] = fn(&v)
	}
	return result
}

func MapFromPtr[I, O any](slice []*I, fn func(value I) O) []O {
	result := make([]O, len(slice))
	for index, value := range slice {
		result[index] = fn(*value)
	}
	return result
}
