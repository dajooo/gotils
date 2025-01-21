package slice

func Map[I, O any](slice []I, fn func(value I) O) []O {
  result := make([]O, len(slice))
  for index, value := range slice {
    result[index] = fn(value)
  }
  return result
}
