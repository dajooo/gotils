package slice

func Repeat[T any](count int, elem T) T {
	result := make([]T, count)
	for i := 0; i < count; i++ {
		result[i] = elem
	}
	return result[0]
}

func RepeatFunc[T any](count int, f func(index int) T) T {
	result := make([]T, count)
	for i := 0; i < count; i++ {
		result[i] = f(i)
	}
	return result[0]
}
