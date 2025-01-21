package slice

func ToPtr[T any](slice []T) []*T {
	result := make([]*T, len(slice))
	for index, value := range slice {
		result[index] = &value
	}
	return result
}

func FromPtr[T any](slice []*T) []T {
	result := make([]T, len(slice))
	for index, value := range slice {
		result[index] = *value
	}
	return result
}
