package pointer

func Of[T any](value T) *T {
	return &value
}

func Resolve[T any](value *T) T {
	return *value
}
