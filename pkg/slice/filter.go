package slice

func Filter[T any](s []T, f func(T) bool) []T {
	var result []T
	for _, v := range s {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

func FilterInstanceOf[T any](s []any) []T {
	var result []T
	for _, v := range s {
		if _, ok := v.(T); ok {
			result = append(result, v.(T))
		}
	}
	return result
}

func FilterNotNil[T any](s []*T) []T {
	var result []T
	for _, v := range s {
		if v != nil {
			result = append(result, *v)
		}
	}
	return result
}
