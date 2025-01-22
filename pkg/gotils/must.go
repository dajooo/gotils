package gotils

func Must[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}

func MustOk[T any](value T, ok bool) T {
	if !ok {
		panic("function's 'ok' value returned false")
	}
	return value
}
