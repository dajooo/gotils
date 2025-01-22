package pointer

func Of[T any](value T) *T {
	return &value
}

func OfOk[T any](value T, ok bool) *T {
	if !ok {
		return nil
	}
	return &value
}

func Resolve[T any](value *T) T {
	return *value
}

func ResolveOrDefault[T any](value *T) T {
	if value == nil {
		var defaultVal T
		return defaultVal
	}
	return *value
}

func ResolveOr[T any](value *T, defaultValue T) T {
	if value == nil {
		return defaultValue
	}
	return *value
}
