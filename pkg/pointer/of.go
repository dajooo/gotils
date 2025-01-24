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
